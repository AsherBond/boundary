// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/api/resources/roles/v1/role.proto

package roles

import (
	scopes "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Principal struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The ID of the principal.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The type of the principal.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The Scope of the principal.
	ScopeId       string `protobuf:"bytes,3,opt,name=scope_id,proto3" json:"scope_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Principal) Reset() {
	*x = Principal{}
	mi := &file_controller_api_resources_roles_v1_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Principal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Principal) ProtoMessage() {}

func (x *Principal) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_roles_v1_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Principal.ProtoReflect.Descriptor instead.
func (*Principal) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_roles_v1_role_proto_rawDescGZIP(), []int{0}
}

func (x *Principal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Principal) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Principal) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

type GrantJson struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The ID, if set.
	// Deprecated: use "ids" instead.
	//
	// Deprecated: Marked as deprecated in controller/api/resources/roles/v1/role.proto.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The IDs, if set.
	Ids []string `protobuf:"bytes,4,rep,name=ids,proto3" json:"ids,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The type, if set.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The actions.
	Actions       []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantJson) Reset() {
	*x = GrantJson{}
	mi := &file_controller_api_resources_roles_v1_role_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantJson) ProtoMessage() {}

func (x *GrantJson) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_roles_v1_role_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantJson.ProtoReflect.Descriptor instead.
func (*GrantJson) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_roles_v1_role_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in controller/api/resources/roles/v1/role.proto.
func (x *GrantJson) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GrantJson) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GrantJson) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GrantJson) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type Grant struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The original user-supplied string.
	Raw string `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The canonically-formatted string.
	Canonical string `protobuf:"bytes,2,opt,name=canonical,proto3" json:"canonical,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The JSON representation of the grant.
	Json          *GrantJson `protobuf:"bytes,3,opt,name=json,proto3" json:"json,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Grant) Reset() {
	*x = Grant{}
	mi := &file_controller_api_resources_roles_v1_role_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Grant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grant) ProtoMessage() {}

func (x *Grant) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_roles_v1_role_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grant.ProtoReflect.Descriptor instead.
func (*Grant) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_roles_v1_role_proto_rawDescGZIP(), []int{2}
}

func (x *Grant) GetRaw() string {
	if x != nil {
		return x.Raw
	}
	return ""
}

func (x *Grant) GetCanonical() string {
	if x != nil {
		return x.Canonical
	}
	return ""
}

func (x *Grant) GetJson() *GrantJson {
	if x != nil {
		return x.Json
	}
	return nil
}

// Role contains all fields related to a Role resource
type Role struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The ID of the Role.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// The ID of the Scope containing this Role.
	ScopeId string `protobuf:"bytes,20,opt,name=scope_id,proto3" json:"scope_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. Scope information for this resource.
	Scope *scopes.ScopeInfo `protobuf:"bytes,30,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional name for identification purposes.
	Name *wrapperspb.StringValue `protobuf:"bytes,40,opt,name=name,proto3" json:"name,omitempty" class:"public"` // @gotags: `class:"public"`
	// Optional user-set description for identification purposes.
	Description *wrapperspb.StringValue `protobuf:"bytes,50,opt,name=description,proto3" json:"description,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The time this resource was created.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,60,opt,name=created_time,proto3" json:"created_time,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The time this resource was last updated.
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,70,opt,name=updated_time,proto3" json:"updated_time,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Version is used in mutation requests, after the initial creation, to ensure this resource has not changed.
	// The mutation will fail if the version does not match the latest known good version.
	Version uint32 `protobuf:"varint,80,opt,name=version,proto3" json:"version,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The IDs of Scopes the grants will apply to. This can include
	// the role's own scope ID, or "this" for the same behavior; specific IDs of
	// scopes that are children of the role's scope; the value "children" to match
	// all direct child scopes of the role's scope; or the value "descendants" to
	// match all descendant scopes (e.g. child scopes, children of child scopes;
	// only valid at "global" scope since it is the only one with children of
	// children).
	GrantScopeIds []string `protobuf:"bytes,91,rep,name=grant_scope_ids,proto3" json:"grant_scope_ids,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The IDs (only) of principals that are assigned to this role.
	PrincipalIds []string `protobuf:"bytes,100,rep,name=principal_ids,proto3" json:"principal_ids,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The principals that are assigned to this role.
	Principals []*Principal `protobuf:"bytes,110,rep,name=principals,proto3" json:"principals,omitempty"`
	// Output only. The grants that this role provides for its principals.
	GrantStrings []string `protobuf:"bytes,120,rep,name=grant_strings,proto3" json:"grant_strings,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The parsed grant information.
	Grants []*Grant `protobuf:"bytes,130,rep,name=grants,proto3" json:"grants,omitempty"`
	// Output only. The available actions on this resource for this user.
	AuthorizedActions []string `protobuf:"bytes,300,rep,name=authorized_actions,proto3" json:"authorized_actions,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_controller_api_resources_roles_v1_role_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_roles_v1_role_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_roles_v1_role_proto_rawDescGZIP(), []int{3}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *Role) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Role) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Role) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Role) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Role) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Role) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Role) GetGrantScopeIds() []string {
	if x != nil {
		return x.GrantScopeIds
	}
	return nil
}

func (x *Role) GetPrincipalIds() []string {
	if x != nil {
		return x.PrincipalIds
	}
	return nil
}

func (x *Role) GetPrincipals() []*Principal {
	if x != nil {
		return x.Principals
	}
	return nil
}

func (x *Role) GetGrantStrings() []string {
	if x != nil {
		return x.GrantStrings
	}
	return nil
}

func (x *Role) GetGrants() []*Grant {
	if x != nil {
		return x.Grants
	}
	return nil
}

func (x *Role) GetAuthorizedActions() []string {
	if x != nil {
		return x.AuthorizedActions
	}
	return nil
}

var File_controller_api_resources_roles_v1_role_proto protoreflect.FileDescriptor

const file_controller_api_resources_roles_v1_role_proto_rawDesc = "" +
	"\n" +
	",controller/api/resources/roles/v1/role.proto\x12!controller.api.resources.roles.v1\x1a.controller/api/resources/scopes/v1/scope.proto\x1a*controller/custom_options/v1/options.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"K\n" +
	"\tPrincipal\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x1a\n" +
	"\bscope_id\x18\x03 \x01(\tR\bscope_id\"_\n" +
	"\tGrantJson\x12\x12\n" +
	"\x02id\x18\x01 \x01(\tB\x02\x18\x01R\x02id\x12\x10\n" +
	"\x03ids\x18\x04 \x03(\tR\x03ids\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x18\n" +
	"\aactions\x18\x03 \x03(\tR\aactions\"y\n" +
	"\x05Grant\x12\x10\n" +
	"\x03raw\x18\x01 \x01(\tR\x03raw\x12\x1c\n" +
	"\tcanonical\x18\x02 \x01(\tR\tcanonical\x12@\n" +
	"\x04json\x18\x03 \x01(\v2,.controller.api.resources.roles.v1.GrantJsonR\x04json\"\x8b\x06\n" +
	"\x04Role\x12\x0e\n" +
	"\x02id\x18\n" +
	" \x01(\tR\x02id\x12\x1a\n" +
	"\bscope_id\x18\x14 \x01(\tR\bscope_id\x12C\n" +
	"\x05scope\x18\x1e \x01(\v2-.controller.api.resources.scopes.v1.ScopeInfoR\x05scope\x12F\n" +
	"\x04name\x18( \x01(\v2\x1c.google.protobuf.StringValueB\x14\xa0\xda)\x01\xc2\xdd)\f\n" +
	"\x04name\x12\x04nameR\x04name\x12b\n" +
	"\vdescription\x182 \x01(\v2\x1c.google.protobuf.StringValueB\"\xa0\xda)\x01\xc2\xdd)\x1a\n" +
	"\vdescription\x12\vdescriptionR\vdescription\x12>\n" +
	"\fcreated_time\x18< \x01(\v2\x1a.google.protobuf.TimestampR\fcreated_time\x12>\n" +
	"\fupdated_time\x18F \x01(\v2\x1a.google.protobuf.TimestampR\fupdated_time\x12\x18\n" +
	"\aversion\x18P \x01(\rR\aversion\x12(\n" +
	"\x0fgrant_scope_ids\x18[ \x03(\tR\x0fgrant_scope_ids\x12$\n" +
	"\rprincipal_ids\x18d \x03(\tR\rprincipal_ids\x12L\n" +
	"\n" +
	"principals\x18n \x03(\v2,.controller.api.resources.roles.v1.PrincipalR\n" +
	"principals\x12$\n" +
	"\rgrant_strings\x18x \x03(\tR\rgrant_strings\x12A\n" +
	"\x06grants\x18\x82\x01 \x03(\v2(.controller.api.resources.roles.v1.GrantR\x06grants\x12/\n" +
	"\x12authorized_actions\x18\xac\x02 \x03(\tR\x12authorized_actionsJ\x04\bZ\x10[R\x0egrant_scope_idBLZJgithub.com/hashicorp/boundary/sdk/pbs/controller/api/resources/roles;rolesb\x06proto3"

var (
	file_controller_api_resources_roles_v1_role_proto_rawDescOnce sync.Once
	file_controller_api_resources_roles_v1_role_proto_rawDescData []byte
)

func file_controller_api_resources_roles_v1_role_proto_rawDescGZIP() []byte {
	file_controller_api_resources_roles_v1_role_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_roles_v1_role_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_api_resources_roles_v1_role_proto_rawDesc), len(file_controller_api_resources_roles_v1_role_proto_rawDesc)))
	})
	return file_controller_api_resources_roles_v1_role_proto_rawDescData
}

var file_controller_api_resources_roles_v1_role_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controller_api_resources_roles_v1_role_proto_goTypes = []any{
	(*Principal)(nil),              // 0: controller.api.resources.roles.v1.Principal
	(*GrantJson)(nil),              // 1: controller.api.resources.roles.v1.GrantJson
	(*Grant)(nil),                  // 2: controller.api.resources.roles.v1.Grant
	(*Role)(nil),                   // 3: controller.api.resources.roles.v1.Role
	(*scopes.ScopeInfo)(nil),       // 4: controller.api.resources.scopes.v1.ScopeInfo
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
}
var file_controller_api_resources_roles_v1_role_proto_depIdxs = []int32{
	1, // 0: controller.api.resources.roles.v1.Grant.json:type_name -> controller.api.resources.roles.v1.GrantJson
	4, // 1: controller.api.resources.roles.v1.Role.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	5, // 2: controller.api.resources.roles.v1.Role.name:type_name -> google.protobuf.StringValue
	5, // 3: controller.api.resources.roles.v1.Role.description:type_name -> google.protobuf.StringValue
	6, // 4: controller.api.resources.roles.v1.Role.created_time:type_name -> google.protobuf.Timestamp
	6, // 5: controller.api.resources.roles.v1.Role.updated_time:type_name -> google.protobuf.Timestamp
	0, // 6: controller.api.resources.roles.v1.Role.principals:type_name -> controller.api.resources.roles.v1.Principal
	2, // 7: controller.api.resources.roles.v1.Role.grants:type_name -> controller.api.resources.roles.v1.Grant
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_controller_api_resources_roles_v1_role_proto_init() }
func file_controller_api_resources_roles_v1_role_proto_init() {
	if File_controller_api_resources_roles_v1_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_api_resources_roles_v1_role_proto_rawDesc), len(file_controller_api_resources_roles_v1_role_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_roles_v1_role_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_roles_v1_role_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_roles_v1_role_proto_msgTypes,
	}.Build()
	File_controller_api_resources_roles_v1_role_proto = out.File
	file_controller_api_resources_roles_v1_role_proto_goTypes = nil
	file_controller_api_resources_roles_v1_role_proto_depIdxs = nil
}
