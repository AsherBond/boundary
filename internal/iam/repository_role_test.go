// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package iam

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/boundary/internal/db"
	dbassert "github.com/hashicorp/boundary/internal/db/assert"
	"github.com/hashicorp/boundary/internal/errors"
	"github.com/hashicorp/boundary/internal/iam/store"
	"github.com/hashicorp/boundary/internal/kms"
	"github.com/hashicorp/boundary/internal/oplog"
	"github.com/hashicorp/go-uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestRepository_CreateRole(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	conn, _ := db.TestSetup(t, "postgres")
	rw := db.New(conn)
	wrapper := db.TestWrapper(t)
	repo := TestRepo(t, conn, wrapper)
	id := testId(t)

	org, proj := TestScopes(t, repo)

	type args struct {
		role *Role
		opt  []Option
	}
	tests := []struct {
		name        string
		args        args
		wantDup     bool
		wantErr     bool
		wantErrMsg  string
		wantIsError errors.Code
	}{
		{
			name: "valid-org",
			args: args{
				role: func() *Role {
					r, err := NewRole(ctx, org.PublicId, WithName("valid-org"+id), WithDescription(id))
					assert.NoError(t, err)
					return r
				}(),
			},
			wantErr: false,
		},
		{
			name: "valid-proj",
			args: args{
				role: func() *Role {
					r, err := NewRole(ctx, proj.PublicId, WithName("valid-proj"+id), WithDescription(id))
					assert.NoError(t, err)
					return r
				}(),
			},
			wantErr: false,
		},
		{
			name: "bad-public-id",
			args: args{
				role: func() *Role {
					r, err := NewRole(ctx, proj.PublicId, WithName("valid-proj"+id), WithDescription(id))
					assert.NoError(t, err)
					r.PublicId = id
					return r
				}(),
			},
			wantErrMsg:  "am.(Repository).CreateRole: public id not empty: parameter violation: error #100",
			wantIsError: errors.InvalidParameter,
			wantErr:     true,
		},
		{
			name: "nil-role",
			args: args{
				role: nil,
			},
			wantErr:     true,
			wantErrMsg:  "iam.(Repository).CreateRole: missing role: parameter violation: error #100",
			wantIsError: errors.InvalidParameter,
		},
		{
			name: "nil-store",
			args: args{
				role: func() *Role {
					return &Role{
						Role: nil,
					}
				}(),
			},
			wantErr:     true,
			wantErrMsg:  "iam.(Repository).CreateRole: missing role store: parameter violation: error #100",
			wantIsError: errors.InvalidParameter,
		},
		{
			name: "bad-scope-id",
			args: args{
				role: func() *Role {
					r, err := NewRole(ctx, id)
					assert.NoError(t, err)
					return r
				}(),
			},
			wantErr:     true,
			wantErrMsg:  "iam.(Repository).create: error getting metadata: iam.(Repository).stdMetadata: unable to get scope: iam.LookupScope: db.LookupWhere: record not found, search issue: error #1100",
			wantIsError: errors.RecordNotFound,
		},
		{
			name: "dup-name",
			args: args{
				role: func() *Role {
					r, err := NewRole(ctx, org.PublicId, WithName("dup-name"+id), WithDescription(id))
					assert.NoError(t, err)
					return r
				}(),
				opt: []Option{WithName("dup-name" + id)},
			},
			wantDup:     true,
			wantErr:     true,
			wantErrMsg:  "already exists in scope ",
			wantIsError: errors.NotUnique,
		},
		{
			name: "dup-name-but-diff-scope",
			args: args{
				role: func() *Role {
					r, err := NewRole(ctx, proj.PublicId, WithName("dup-name-but-diff-scope"+id), WithDescription(id))
					assert.NoError(t, err)
					return r
				}(),
				opt: []Option{WithName("dup-name-but-diff-scope" + id)},
			},
			wantDup: true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			if tt.wantDup {
				dup, err := NewRole(ctx, org.PublicId, tt.args.opt...)
				assert.NoError(err)
				dup, _, _, _, err = repo.CreateRole(context.Background(), dup, tt.args.opt...)
				assert.NoError(err)
				assert.NotNil(dup)
			}
			grp, _, _, _, err := repo.CreateRole(context.Background(), tt.args.role, tt.args.opt...)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(grp)
				assert.Contains(err.Error(), tt.wantErrMsg)
				assert.True(errors.Match(errors.T(tt.wantIsError), err))
				return
			}
			assert.NoError(err)
			assert.NotNil(grp.CreateTime)
			assert.NotNil(grp.UpdateTime)

			foundGrp, _, _, _, err := repo.LookupRole(context.Background(), grp.PublicId)
			assert.NoError(err)
			assert.True(proto.Equal(foundGrp, grp))

			err = db.TestVerifyOplog(t, rw, grp.PublicId, db.WithOperation(oplog.OpType_OP_TYPE_CREATE), db.WithCreateNotBefore(10*time.Second))
			assert.NoError(err)
		})
	}
}

func TestRepository_UpdateRole(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	conn, _ := db.TestSetup(t, "postgres")
	rw := db.New(conn)
	wrapper := db.TestWrapper(t)
	repo := TestRepo(t, conn, wrapper)
	id, err := uuid.GenerateUUID()
	require.NoError(t, err)

	org, proj := TestScopes(t, repo)
	u := TestUser(t, repo, org.GetPublicId())

	pubId := func(s string) *string { return &s }

	type args struct {
		name           string
		description    string
		fieldMaskPaths []string
		opt            []Option
		ScopeId        string
		PublicId       *string
	}
	tests := []struct {
		name           string
		newScopeId     string
		newRoleOpts    []Option
		args           args
		wantRowsUpdate int
		wantErr        bool
		wantErrMsg     string
		wantIsError    errors.Code
		wantDup        bool
	}{
		{
			name: "valid",
			args: args{
				name:           "valid" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        org.PublicId,
			},
			newScopeId:     org.PublicId,
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "valid-no-op",
			args: args{
				name:           "valid-no-op" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        org.PublicId,
			},
			newScopeId:     org.PublicId,
			newRoleOpts:    []Option{WithName("valid-no-op" + id)},
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "not-found",
			args: args{
				name:           "not-found" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        org.PublicId,
				PublicId:       func() *string { s := "1"; return &s }(),
			},
			newScopeId:     org.PublicId,
			wantErr:        true,
			wantRowsUpdate: 0,
			wantErrMsg:     "error #1100",
			wantIsError:    errors.RecordNotFound,
		},
		{
			name: "null-name",
			args: args{
				name:           "",
				fieldMaskPaths: []string{"Name"},
				ScopeId:        org.PublicId,
			},
			newScopeId:     org.PublicId,
			newRoleOpts:    []Option{WithName("null-name" + id)},
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "null-description",
			args: args{
				name:           "",
				fieldMaskPaths: []string{"Description"},
				ScopeId:        org.PublicId,
			},
			newScopeId:     org.PublicId,
			newRoleOpts:    []Option{WithDescription("null-description" + id)},
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "empty-field-mask",
			args: args{
				name:           "valid" + id,
				fieldMaskPaths: []string{},
				ScopeId:        org.PublicId,
			},
			newScopeId:     org.PublicId,
			wantErr:        true,
			wantRowsUpdate: 0,
			wantErrMsg:     "iam.(Repository).UpdateRole: empty field mask, parameter violation: error #104",
			wantIsError:    errors.EmptyFieldMask,
		},
		{
			name: "nil-fieldmask",
			args: args{
				name:           "valid" + id,
				fieldMaskPaths: nil,
				ScopeId:        org.PublicId,
			},
			newScopeId:     org.PublicId,
			wantErr:        true,
			wantRowsUpdate: 0,
			wantErrMsg:     "iam.(Repository).UpdateRole: empty field mask, parameter violation: error #104",
			wantIsError:    errors.EmptyFieldMask,
		},
		{
			name: "read-only-fields",
			args: args{
				name:           "valid" + id,
				fieldMaskPaths: []string{"CreateTime"},
				ScopeId:        org.PublicId,
			},
			newScopeId:     org.PublicId,
			wantErr:        true,
			wantRowsUpdate: 0,
			wantErrMsg:     "iam.(Repository).UpdateRole: invalid field mask: CreateTime: parameter violation: error #103",
			wantIsError:    errors.InvalidFieldMask,
		},
		{
			name: "unknown-fields",
			args: args{
				name:           "valid" + id,
				fieldMaskPaths: []string{"Alice"},
				ScopeId:        org.PublicId,
			},
			newScopeId:     org.PublicId,
			wantErr:        true,
			wantRowsUpdate: 0,
			wantErrMsg:     "iam.(Repository).UpdateRole: invalid field mask: Alice: parameter violation: error #103",
			wantIsError:    errors.InvalidFieldMask,
		},
		{
			name: "no-public-id",
			args: args{
				name:           "valid" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        org.PublicId,
				PublicId:       pubId(""),
			},
			newScopeId:     org.PublicId,
			wantErr:        true,
			wantErrMsg:     "iam.(Repository).UpdateRole: missing public id: parameter violation: error #100",
			wantIsError:    errors.InvalidParameter,
			wantRowsUpdate: 0,
		},
		{
			name: "proj-scope-id-no-mask",
			args: args{
				name:    "proj-scope-id" + id,
				ScopeId: proj.PublicId,
			},
			newScopeId:  org.PublicId,
			wantErr:     true,
			wantErrMsg:  "iam.(Repository).UpdateRole: empty field mask, parameter violation: error #104",
			wantIsError: errors.EmptyFieldMask,
		},
		{
			name: "empty-scope-id-with-name-mask",
			args: args{
				name:           "empty-scope-id" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        "",
			},
			newScopeId:     org.PublicId,
			wantErr:        false,
			wantRowsUpdate: 1,
		},
		{
			name: "dup-name-in-diff-scope",
			args: args{
				name:           "dup-name-in-diff-scope" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        proj.PublicId,
			},
			newScopeId:     proj.PublicId,
			newRoleOpts:    []Option{WithName("dup-name-in-diff-scope-pre-update" + id)},
			wantErr:        false,
			wantRowsUpdate: 1,
			wantDup:        true,
		},
		{
			name: "dup-name",
			args: args{
				name:           "dup-name" + id,
				fieldMaskPaths: []string{"Name"},
				ScopeId:        org.PublicId,
			},
			newScopeId:  org.PublicId,
			wantErr:     true,
			wantDup:     true,
			wantErrMsg:  " already exists in org " + org.PublicId,
			wantIsError: errors.NotUnique,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require, assert := require.New(t), assert.New(t)
			if tt.wantDup {
				r := TestRole(t, conn, org.PublicId)
				_ = TestUserRole(t, conn, r.GetPublicId(), u.GetPublicId())
				_ = TestRoleGrant(t, conn, r.GetPublicId(), "ids=*;type=*;actions=*")
				r.Name = tt.args.name
				_, _, _, _, _, err := repo.UpdateRole(context.Background(), r, r.Version, tt.args.fieldMaskPaths, tt.args.opt...)
				assert.NoError(err)
			}

			r := TestRole(t, conn, tt.newScopeId, tt.newRoleOpts...)
			ur := TestUserRole(t, conn, r.GetPublicId(), u.GetPublicId())
			princRole := &PrincipalRole{PrincipalRoleView: &store.PrincipalRoleView{
				Type:              UserRoleType.String(),
				CreateTime:        ur.CreateTime,
				RoleId:            ur.RoleId,
				PrincipalId:       ur.PrincipalId,
				PrincipalScopeId:  org.GetPublicId(),
				ScopedPrincipalId: ur.PrincipalId,
				RoleScopeId:       org.GetPublicId(),
			}}
			if tt.newScopeId != org.GetPublicId() {
				// If the project is in a different scope from the created user we need to update the
				// scope specific fields.
				princRole.RoleScopeId = tt.newScopeId
				princRole.ScopedPrincipalId = fmt.Sprintf("%s:%s", org.PublicId, ur.PrincipalId)
			}
			rGrant := TestRoleGrant(t, conn, r.GetPublicId(), "ids=*;type=*;actions=*")

			updateRole := allocRole()
			updateRole.PublicId = r.PublicId
			if tt.args.PublicId != nil {
				updateRole.PublicId = *tt.args.PublicId
			}
			updateRole.ScopeId = tt.args.ScopeId
			updateRole.Name = tt.args.name
			updateRole.Description = tt.args.description

			roleAfterUpdate, principals, grants, _, updatedRows, err := repo.UpdateRole(context.Background(), &updateRole, r.Version, tt.args.fieldMaskPaths, tt.args.opt...)
			if tt.wantErr {
				require.Error(err)
				assert.True(errors.Match(errors.T(tt.wantIsError), err))
				assert.Nil(roleAfterUpdate)
				assert.Equal(0, updatedRows)
				assert.Contains(err.Error(), tt.wantErrMsg)
				err = db.TestVerifyOplog(t, rw, r.PublicId, db.WithOperation(oplog.OpType_OP_TYPE_UPDATE), db.WithCreateNotBefore(10*time.Second))
				assert.Error(err)
				assert.True(errors.IsNotFoundError(err))
				return
			}
			require.NoError(err)
			require.NotNil(roleAfterUpdate)
			assert.Equal(tt.wantRowsUpdate, updatedRows)
			assert.Equal([]*PrincipalRole{princRole}, principals)
			assert.Equal([]*RoleGrant{rGrant}, grants)
			switch tt.name {
			case "valid-no-op":
				assert.Equal(r.UpdateTime, roleAfterUpdate.UpdateTime)
			default:
				assert.NotEqual(r.UpdateTime, roleAfterUpdate.UpdateTime)
			}
			foundRole, _, _, _, err := repo.LookupRole(context.Background(), r.PublicId)
			assert.NoError(err)
			assert.True(proto.Equal(roleAfterUpdate, foundRole))
			underlyingDB, err := conn.SqlDB(ctx)
			require.NoError(err)
			dbassert := dbassert.New(t, underlyingDB)
			if tt.args.name == "" {
				assert.Equal(foundRole.Name, "")
				dbassert.IsNull(foundRole, "name")
			}
			if tt.args.description == "" {
				assert.Equal(foundRole.Description, "")
				dbassert.IsNull(foundRole, "description")
			}
			err = db.TestVerifyOplog(t, rw, r.PublicId, db.WithOperation(oplog.OpType_OP_TYPE_UPDATE), db.WithCreateNotBefore(10*time.Second))
			assert.NoError(err)
		})
	}
}

func TestRepository_DeleteRole(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	conn, _ := db.TestSetup(t, "postgres")
	rw := db.New(conn)
	wrapper := db.TestWrapper(t)
	repo := TestRepo(t, conn, wrapper)
	org, _ := TestScopes(t, repo)

	roleId, err := newRoleId(ctx)
	require.NoError(t, err)

	type args struct {
		role *Role
		opt  []Option
	}
	tests := []struct {
		name            string
		args            args
		wantRowsDeleted int
		wantErr         bool
		wantErrMsg      string
	}{
		{
			name: "valid",
			args: args{
				role: TestRole(t, conn, org.PublicId),
			},
			wantRowsDeleted: 1,
			wantErr:         false,
		},
		{
			name: "no-public-id",
			args: args{
				role: func() *Role {
					r := allocRole()
					return &r
				}(),
			},
			wantRowsDeleted: 0,
			wantErr:         true,
			wantErrMsg:      "iam.(Repository).DeleteRole: missing public id: parameter violation: error #100",
		},

		{
			name: "not-found",
			args: args{
				role: func() *Role {
					r, err := NewRole(ctx, org.PublicId)
					r.PublicId = roleId
					require.NoError(t, err)
					return r
				}(),
			},
			wantRowsDeleted: 0,
			wantErr:         true,
			wantErrMsg:      "iam.(Repository).DeleteRole: failed for " + roleId + ": db.LookupById: record not found, search issue: error #1100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			deletedRows, err := repo.DeleteRole(ctx, tt.args.role.PublicId, tt.args.opt...)
			if tt.wantErr {
				assert.Error(err)
				assert.Equal(0, deletedRows)
				assert.Contains(err.Error(), tt.wantErrMsg)
				err = db.TestVerifyOplog(t, rw, tt.args.role.PublicId, db.WithOperation(oplog.OpType_OP_TYPE_DELETE), db.WithCreateNotBefore(10*time.Second))
				assert.Error(err)
				assert.True(errors.IsNotFoundError(err))
				return
			}
			assert.NoError(err)
			assert.Equal(tt.wantRowsDeleted, deletedRows)
			foundRole, _, _, _, err := repo.LookupRole(ctx, tt.args.role.PublicId)
			assert.NoError(err)
			assert.Nil(foundRole)

			err = db.TestVerifyOplog(t, rw, tt.args.role.PublicId, db.WithOperation(oplog.OpType_OP_TYPE_DELETE), db.WithCreateNotBefore(10*time.Second))
			assert.NoError(err)
		})
	}
}

func TestRepository_listRoles(t *testing.T) {
	t.Parallel()
	conn, _ := db.TestSetup(t, "postgres")
	rw := db.New(conn)
	const testLimit = 10
	wrapper := db.TestWrapper(t)
	kms := kms.TestKms(t, conn, wrapper)
	repo := TestRepo(t, conn, wrapper, WithLimit(testLimit))
	org, proj := TestScopes(t, repo, WithSkipDefaultRoleCreation(true))

	type args struct {
		withScopeId string
		opt         []Option
	}
	tests := []struct {
		name          string
		createCnt     int
		createScopeId string
		args          args
		wantCnt       int
		wantErr       bool
	}{
		{
			name:          "negative-limit",
			createCnt:     repo.defaultLimit + 1,
			createScopeId: org.PublicId,
			args: args{
				withScopeId: org.PublicId,
				opt:         []Option{WithLimit(-1)},
			},
			wantCnt: repo.defaultLimit + 1,
			wantErr: true,
		},
		{
			name:          "negative-limit-proj-role",
			createCnt:     repo.defaultLimit + 1,
			createScopeId: proj.PublicId,
			args: args{
				withScopeId: proj.PublicId,
				opt:         []Option{WithLimit(-1)},
			},
			wantCnt: repo.defaultLimit + 1,
			wantErr: true,
		},
		{
			name:          "default-limit",
			createCnt:     repo.defaultLimit + 1,
			createScopeId: org.PublicId,
			args: args{
				withScopeId: org.PublicId,
			},
			wantCnt: repo.defaultLimit,
			wantErr: false,
		},
		{
			name:          "custom-limit",
			createCnt:     repo.defaultLimit + 1,
			createScopeId: org.PublicId,
			args: args{
				withScopeId: org.PublicId,
				opt:         []Option{WithLimit(3)},
			},
			wantCnt: 3,
			wantErr: false,
		},
		{
			name:          "bad-org",
			createCnt:     1,
			createScopeId: org.PublicId,
			args: args{
				withScopeId: "bad-id",
			},
			wantCnt: 0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert, require := assert.New(t), require.New(t)
			t.Cleanup(func() {
				db.TestDeleteWhere(t, conn, func() any { r := allocRole(); return &r }(), "1=1")
			})
			testRoles := []*Role{}
			for i := 0; i < tt.createCnt; i++ {
				testRoles = append(testRoles, TestRole(t, conn, tt.createScopeId))
			}
			assert.Equal(tt.createCnt, len(testRoles))
			got, ttime, err := repo.listRoles(context.Background(), []string{tt.args.withScopeId}, tt.args.opt...)
			if tt.wantErr {
				require.Error(err)
				return
			}
			require.NoError(err)
			assert.Equal(tt.wantCnt, len(got))
			// Transaction timestamp should be within ~10 seconds of now
			assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
			assert.True(time.Now().After(ttime.Add(-10 * time.Second)))
		})
	}
	t.Run("withStartPageAfter", func(t *testing.T) {
		assert, require := assert.New(t), require.New(t)
		ctx := context.Background()

		for i := 0; i < 10; i++ {
			_ = TestRole(t, conn, org.GetPublicId())
		}

		repo, err := NewRepository(ctx, rw, rw, kms)
		require.NoError(err)
		page1, ttime, err := repo.listRoles(ctx, []string{org.GetPublicId()}, WithLimit(2))
		require.NoError(err)
		require.Len(page1, 2)
		// Transaction timestamp should be within ~10 seconds of now
		assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
		assert.True(time.Now().After(ttime.Add(-10 * time.Second)))
		page2, ttime, err := repo.listRoles(ctx, []string{org.GetPublicId()}, WithLimit(2), WithStartPageAfterItem(page1[1]))
		require.NoError(err)
		require.Len(page2, 2)
		assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
		assert.True(time.Now().After(ttime.Add(-10 * time.Second)))
		for _, item := range page1 {
			assert.NotEqual(item.GetPublicId(), page2[0].GetPublicId())
			assert.NotEqual(item.GetPublicId(), page2[1].GetPublicId())
		}
		page3, ttime, err := repo.listRoles(ctx, []string{org.GetPublicId()}, WithLimit(2), WithStartPageAfterItem(page2[1]))
		require.NoError(err)
		require.Len(page3, 2)
		assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
		assert.True(time.Now().After(ttime.Add(-10 * time.Second)))
		for _, item := range page2 {
			assert.NotEqual(item.GetPublicId(), page3[0].GetPublicId())
			assert.NotEqual(item.GetPublicId(), page3[1].GetPublicId())
		}
		page4, ttime, err := repo.listRoles(ctx, []string{org.GetPublicId()}, WithLimit(2), WithStartPageAfterItem(page3[1]))
		require.NoError(err)
		assert.Len(page4, 2)
		assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
		assert.True(time.Now().After(ttime.Add(-10 * time.Second)))
		for _, item := range page3 {
			assert.NotEqual(item.GetPublicId(), page4[0].GetPublicId())
			assert.NotEqual(item.GetPublicId(), page4[1].GetPublicId())
		}
		page5, ttime, err := repo.listRoles(ctx, []string{org.GetPublicId()}, WithLimit(2), WithStartPageAfterItem(page4[1]))
		require.NoError(err)
		assert.Len(page5, 2)
		assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
		assert.True(time.Now().After(ttime.Add(-10 * time.Second)))
		for _, item := range page4 {
			assert.NotEqual(item.GetPublicId(), page5[0].GetPublicId())
			assert.NotEqual(item.GetPublicId(), page5[1].GetPublicId())
		}
		page6, ttime, err := repo.listRoles(ctx, []string{org.GetPublicId()}, WithLimit(2), WithStartPageAfterItem(page5[1]))
		require.NoError(err)
		assert.Empty(page6)
		assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
		assert.True(time.Now().After(ttime.Add(-10 * time.Second)))

		// Create 2 new roles
		newR1 := TestRole(t, conn, org.GetPublicId())
		newR2 := TestRole(t, conn, org.GetPublicId())

		// since it will return newest to oldest, we get page1[1] first
		page7, ttime, err := repo.listRolesRefresh(
			ctx,
			time.Now().Add(-1*time.Second),
			[]string{org.GetPublicId()},
			WithLimit(1),
		)
		require.NoError(err)
		require.Len(page7, 1)
		require.Equal(page7[0].GetPublicId(), newR2.GetPublicId())
		assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
		assert.True(time.Now().After(ttime.Add(-10 * time.Second)))

		page8, ttime, err := repo.listRolesRefresh(
			context.Background(),
			time.Now().Add(-1*time.Second),
			[]string{org.GetPublicId()},
			WithLimit(1),
			WithStartPageAfterItem(page7[0]),
		)
		require.NoError(err)
		require.Len(page8, 1)
		require.Equal(page8[0].GetPublicId(), newR1.GetPublicId())
		assert.True(time.Now().Before(ttime.Add(10 * time.Second)))
		assert.True(time.Now().After(ttime.Add(-10 * time.Second)))
	})
}

func TestRepository_ListRoles_Multiple_Scopes(t *testing.T) {
	t.Parallel()
	conn, _ := db.TestSetup(t, "postgres")
	wrapper := db.TestWrapper(t)
	repo := TestRepo(t, conn, wrapper)
	org, proj := TestScopes(t, repo)

	db.TestDeleteWhere(t, conn, func() any { i := allocRole(); return &i }(), "1=1")

	const numPerScope = 10
	var total int
	for i := 0; i < numPerScope; i++ {
		TestRole(t, conn, "global")
		total++
		TestRole(t, conn, org.GetPublicId())
		total++
		TestRole(t, conn, proj.GetPublicId())
		total++
	}

	got, ttime, err := repo.listRoles(context.Background(), []string{"global", org.GetPublicId(), proj.GetPublicId()})
	require.NoError(t, err)
	assert.Equal(t, total, len(got))
	// Transaction timestamp should be within ~10 seconds of now
	assert.True(t, time.Now().Before(ttime.Add(10*time.Second)))
	assert.True(t, time.Now().After(ttime.Add(-10*time.Second)))
}

func Test_listRoleDeletedIds(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	conn, _ := db.TestSetup(t, "postgres")
	wrapper := db.TestWrapper(t)
	repo := TestRepo(t, conn, wrapper)
	org, _ := TestScopes(t, repo, WithSkipDefaultRoleCreation(true))
	r := TestRole(t, conn, org.GetPublicId())

	// Expect no entries at the start
	deletedIds, ttime, err := repo.listRoleDeletedIds(ctx, time.Now().AddDate(-1, 0, 0))
	require.NoError(t, err)
	require.Empty(t, deletedIds)
	// Transaction timestamp should be within ~10 seconds of now
	assert.True(t, time.Now().Before(ttime.Add(10*time.Second)))
	assert.True(t, time.Now().After(ttime.Add(-10*time.Second)))

	// Delete a role
	_, err = repo.DeleteRole(ctx, r.GetPublicId())
	require.NoError(t, err)

	// Expect a single entry
	deletedIds, ttime, err = repo.listRoleDeletedIds(ctx, time.Now().AddDate(-1, 0, 0))
	require.NoError(t, err)
	require.Equal(t, []string{r.GetPublicId()}, deletedIds)
	// Transaction timestamp should be within ~10 seconds of now
	assert.True(t, time.Now().Before(ttime.Add(10*time.Second)))
	assert.True(t, time.Now().After(ttime.Add(-10*time.Second)))

	// Try again with the time set to now, expect no entries
	deletedIds, ttime, err = repo.listRoleDeletedIds(ctx, time.Now())
	require.NoError(t, err)
	require.Empty(t, deletedIds)
	// Transaction timestamp should be within ~10 seconds of now
	assert.True(t, time.Now().Before(ttime.Add(10*time.Second)))
	assert.True(t, time.Now().After(ttime.Add(-10*time.Second)))
}

func Test_estimatedRoleCount(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	conn, _ := db.TestSetup(t, "postgres")
	sqlDb, err := conn.SqlDB(ctx)
	require.NoError(t, err)
	rw := db.New(conn)
	wrapper := db.TestWrapper(t)
	kms := kms.TestKms(t, conn, wrapper)
	repo, err := NewRepository(ctx, rw, rw, kms)
	require.NoError(t, err)

	// Run analyze to update estimate
	_, err = sqlDb.ExecContext(ctx, "analyze")
	require.NoError(t, err)

	// Check total entries at start, expect 0
	numItems, err := repo.estimatedRoleCount(ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, numItems)

	iamRepo := TestRepo(t, conn, wrapper)
	org, _ := TestScopes(t, iamRepo, WithSkipDefaultRoleCreation(true))
	// Create a role, expect 1 entry
	r := TestRole(t, conn, org.GetPublicId())

	// Run analyze to update estimate
	_, err = sqlDb.ExecContext(ctx, "analyze")
	require.NoError(t, err)
	numItems, err = repo.estimatedRoleCount(ctx)
	require.NoError(t, err)
	assert.Equal(t, 1, numItems)

	// Delete the role, expect 0 again
	_, err = repo.DeleteRole(ctx, r.GetPublicId())
	require.NoError(t, err)

	// Run analyze to update estimate
	_, err = sqlDb.ExecContext(ctx, "analyze")
	require.NoError(t, err)
	numItems, err = repo.estimatedRoleCount(ctx)
	require.NoError(t, err)
	assert.Equal(t, 0, numItems)
}
