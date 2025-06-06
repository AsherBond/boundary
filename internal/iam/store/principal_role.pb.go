// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/storage/iam/store/v1/principal_role.proto

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type UserRole struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// role_id is the role of this principal.
	// @inject_tag: gorm:"primary_key"
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" gorm:"primary_key"`
	// principal_id is the public_id of the user (which is the principal)
	// @inject_tag: gorm:"primary_key"
	PrincipalId   string `protobuf:"bytes,3,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty" gorm:"primary_key"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRole) Reset() {
	*x = UserRole{}
	mi := &file_controller_storage_iam_store_v1_principal_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRole) ProtoMessage() {}

func (x *UserRole) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_principal_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRole.ProtoReflect.Descriptor instead.
func (*UserRole) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_principal_role_proto_rawDescGZIP(), []int{0}
}

func (x *UserRole) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *UserRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *UserRole) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

type GroupRole struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// role_id is the role of this principal.
	// @inject_tag: gorm:"primary_key"
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" gorm:"primary_key"`
	// principal_id is the public_id of the group (which is the principal)
	// @inject_tag: gorm:"primary_key"
	PrincipalId   string `protobuf:"bytes,3,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty" gorm:"primary_key"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupRole) Reset() {
	*x = GroupRole{}
	mi := &file_controller_storage_iam_store_v1_principal_role_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRole) ProtoMessage() {}

func (x *GroupRole) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_principal_role_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRole.ProtoReflect.Descriptor instead.
func (*GroupRole) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_principal_role_proto_rawDescGZIP(), []int{1}
}

func (x *GroupRole) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GroupRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *GroupRole) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

type ManagedGroupRole struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// role_id is the role of this principal.
	// @inject_tag: gorm:"primary_key"
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" gorm:"primary_key"`
	// principal_id is the public_id of the managed group (which is the principal)
	// @inject_tag: gorm:"primary_key"
	PrincipalId   string `protobuf:"bytes,3,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty" gorm:"primary_key"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ManagedGroupRole) Reset() {
	*x = ManagedGroupRole{}
	mi := &file_controller_storage_iam_store_v1_principal_role_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ManagedGroupRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagedGroupRole) ProtoMessage() {}

func (x *ManagedGroupRole) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_principal_role_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagedGroupRole.ProtoReflect.Descriptor instead.
func (*ManagedGroupRole) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_principal_role_proto_rawDescGZIP(), []int{2}
}

func (x *ManagedGroupRole) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ManagedGroupRole) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *ManagedGroupRole) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

type PrincipalRoleView struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// role_id is the role of this principal.
	// @inject_tag: gorm:"primary_key"
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" gorm:"primary_key"`
	// Principal type (User or Group)
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// @inject_tag: gorm:"primary_key"
	PrincipalId string `protobuf:"bytes,4,opt,name=principal_id,json=principalId,proto3" json:"principal_id,omitempty" gorm:"primary_key"`
	// principal_scope_id of the principal
	// @inject_tag: `gorm:"default:null"`
	PrincipalScopeId string `protobuf:"bytes,5,opt,name=principal_scope_id,json=principalScopeId,proto3" json:"principal_scope_id,omitempty" gorm:"default:null"`
	// role_scope_id of the role
	// @inject_tag: `gorm:"default:null"`
	RoleScopeId string `protobuf:"bytes,6,opt,name=role_scope_id,json=roleScopeId,proto3" json:"role_scope_id,omitempty" gorm:"default:null"`
	// scoped_principal_id of the principal
	// @inject_tag: `gorm:"default:null"`
	ScopedPrincipalId string `protobuf:"bytes,7,opt,name=scoped_principal_id,json=scopedPrincipalId,proto3" json:"scoped_principal_id,omitempty" gorm:"default:null"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *PrincipalRoleView) Reset() {
	*x = PrincipalRoleView{}
	mi := &file_controller_storage_iam_store_v1_principal_role_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrincipalRoleView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrincipalRoleView) ProtoMessage() {}

func (x *PrincipalRoleView) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_principal_role_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrincipalRoleView.ProtoReflect.Descriptor instead.
func (*PrincipalRoleView) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_principal_role_proto_rawDescGZIP(), []int{3}
}

func (x *PrincipalRoleView) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *PrincipalRoleView) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *PrincipalRoleView) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PrincipalRoleView) GetPrincipalId() string {
	if x != nil {
		return x.PrincipalId
	}
	return ""
}

func (x *PrincipalRoleView) GetPrincipalScopeId() string {
	if x != nil {
		return x.PrincipalScopeId
	}
	return ""
}

func (x *PrincipalRoleView) GetRoleScopeId() string {
	if x != nil {
		return x.RoleScopeId
	}
	return ""
}

func (x *PrincipalRoleView) GetScopedPrincipalId() string {
	if x != nil {
		return x.ScopedPrincipalId
	}
	return ""
}

var File_controller_storage_iam_store_v1_principal_role_proto protoreflect.FileDescriptor

const file_controller_storage_iam_store_v1_principal_role_proto_rawDesc = "" +
	"\n" +
	"4controller/storage/iam/store/v1/principal_role.proto\x12\x1fcontroller.storage.iam.store.v1\x1a/controller/storage/timestamp/v1/timestamp.proto\"\x93\x01\n" +
	"\bUserRole\x12K\n" +
	"\vcreate_time\x18\x01 \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\x12\x17\n" +
	"\arole_id\x18\x02 \x01(\tR\x06roleId\x12!\n" +
	"\fprincipal_id\x18\x03 \x01(\tR\vprincipalId\"\x94\x01\n" +
	"\tGroupRole\x12K\n" +
	"\vcreate_time\x18\x01 \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\x12\x17\n" +
	"\arole_id\x18\x02 \x01(\tR\x06roleId\x12!\n" +
	"\fprincipal_id\x18\x03 \x01(\tR\vprincipalId\"\x9b\x01\n" +
	"\x10ManagedGroupRole\x12K\n" +
	"\vcreate_time\x18\x01 \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\x12\x17\n" +
	"\arole_id\x18\x02 \x01(\tR\x06roleId\x12!\n" +
	"\fprincipal_id\x18\x03 \x01(\tR\vprincipalId\"\xb2\x02\n" +
	"\x11PrincipalRoleView\x12K\n" +
	"\vcreate_time\x18\x01 \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\x12\x17\n" +
	"\arole_id\x18\x02 \x01(\tR\x06roleId\x12\x12\n" +
	"\x04type\x18\x03 \x01(\tR\x04type\x12!\n" +
	"\fprincipal_id\x18\x04 \x01(\tR\vprincipalId\x12,\n" +
	"\x12principal_scope_id\x18\x05 \x01(\tR\x10principalScopeId\x12\"\n" +
	"\rrole_scope_id\x18\x06 \x01(\tR\vroleScopeId\x12.\n" +
	"\x13scoped_principal_id\x18\a \x01(\tR\x11scopedPrincipalIdB8Z6github.com/hashicorp/boundary/internal/iam/store;storeb\x06proto3"

var (
	file_controller_storage_iam_store_v1_principal_role_proto_rawDescOnce sync.Once
	file_controller_storage_iam_store_v1_principal_role_proto_rawDescData []byte
)

func file_controller_storage_iam_store_v1_principal_role_proto_rawDescGZIP() []byte {
	file_controller_storage_iam_store_v1_principal_role_proto_rawDescOnce.Do(func() {
		file_controller_storage_iam_store_v1_principal_role_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_storage_iam_store_v1_principal_role_proto_rawDesc), len(file_controller_storage_iam_store_v1_principal_role_proto_rawDesc)))
	})
	return file_controller_storage_iam_store_v1_principal_role_proto_rawDescData
}

var file_controller_storage_iam_store_v1_principal_role_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controller_storage_iam_store_v1_principal_role_proto_goTypes = []any{
	(*UserRole)(nil),            // 0: controller.storage.iam.store.v1.UserRole
	(*GroupRole)(nil),           // 1: controller.storage.iam.store.v1.GroupRole
	(*ManagedGroupRole)(nil),    // 2: controller.storage.iam.store.v1.ManagedGroupRole
	(*PrincipalRoleView)(nil),   // 3: controller.storage.iam.store.v1.PrincipalRoleView
	(*timestamp.Timestamp)(nil), // 4: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_iam_store_v1_principal_role_proto_depIdxs = []int32{
	4, // 0: controller.storage.iam.store.v1.UserRole.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	4, // 1: controller.storage.iam.store.v1.GroupRole.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	4, // 2: controller.storage.iam.store.v1.ManagedGroupRole.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	4, // 3: controller.storage.iam.store.v1.PrincipalRoleView.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_controller_storage_iam_store_v1_principal_role_proto_init() }
func file_controller_storage_iam_store_v1_principal_role_proto_init() {
	if File_controller_storage_iam_store_v1_principal_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_storage_iam_store_v1_principal_role_proto_rawDesc), len(file_controller_storage_iam_store_v1_principal_role_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_iam_store_v1_principal_role_proto_goTypes,
		DependencyIndexes: file_controller_storage_iam_store_v1_principal_role_proto_depIdxs,
		MessageInfos:      file_controller_storage_iam_store_v1_principal_role_proto_msgTypes,
	}.Build()
	File_controller_storage_iam_store_v1_principal_role_proto = out.File
	file_controller_storage_iam_store_v1_principal_role_proto_goTypes = nil
	file_controller_storage_iam_store_v1_principal_role_proto_depIdxs = nil
}
