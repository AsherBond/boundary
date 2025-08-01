// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/storage/iam/store/v1/role_global_individual_org_grant_scope.proto

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

type GlobalRoleIndividualOrgGrantScope struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// role_id is the ID of the role this is a part of
	// @inject_tag: `gorm:"primary_key"`
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" gorm:"primary_key"`
	// scope_id is the string grant scope value as provided by the user
	//
	// @inject_tag: `gorm:"primary_key"`
	ScopeId string `protobuf:"bytes,3,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"primary_key"`
	// grant_scope control type of grant scope granted to this role ['individual']
	//
	// @inject_tag: `gorm:"default:null"`
	GrantScope    string `protobuf:"bytes,4,opt,name=grant_scope,json=grantScope,proto3" json:"grant_scope,omitempty" gorm:"default:null"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GlobalRoleIndividualOrgGrantScope) Reset() {
	*x = GlobalRoleIndividualOrgGrantScope{}
	mi := &file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GlobalRoleIndividualOrgGrantScope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlobalRoleIndividualOrgGrantScope) ProtoMessage() {}

func (x *GlobalRoleIndividualOrgGrantScope) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlobalRoleIndividualOrgGrantScope.ProtoReflect.Descriptor instead.
func (*GlobalRoleIndividualOrgGrantScope) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDescGZIP(), []int{0}
}

func (x *GlobalRoleIndividualOrgGrantScope) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *GlobalRoleIndividualOrgGrantScope) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *GlobalRoleIndividualOrgGrantScope) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *GlobalRoleIndividualOrgGrantScope) GetGrantScope() string {
	if x != nil {
		return x.GrantScope
	}
	return ""
}

var File_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto protoreflect.FileDescriptor

const file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDesc = "" +
	"\n" +
	"Lcontroller/storage/iam/store/v1/role_global_individual_org_grant_scope.proto\x12\x1fcontroller.storage.iam.store.v1\x1a/controller/storage/timestamp/v1/timestamp.proto\"\xc5\x01\n" +
	"!GlobalRoleIndividualOrgGrantScope\x12K\n" +
	"\vcreate_time\x18\x01 \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\x12\x17\n" +
	"\arole_id\x18\x02 \x01(\tR\x06roleId\x12\x19\n" +
	"\bscope_id\x18\x03 \x01(\tR\ascopeId\x12\x1f\n" +
	"\vgrant_scope\x18\x04 \x01(\tR\n" +
	"grantScopeB8Z6github.com/hashicorp/boundary/internal/iam/store;storeb\x06proto3"

var (
	file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDescOnce sync.Once
	file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDescData []byte
)

func file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDescGZIP() []byte {
	file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDescOnce.Do(func() {
		file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDesc), len(file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDesc)))
	})
	return file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDescData
}

var file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_goTypes = []any{
	(*GlobalRoleIndividualOrgGrantScope)(nil), // 0: controller.storage.iam.store.v1.GlobalRoleIndividualOrgGrantScope
	(*timestamp.Timestamp)(nil),               // 1: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_depIdxs = []int32{
	1, // 0: controller.storage.iam.store.v1.GlobalRoleIndividualOrgGrantScope.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_init() }
func file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_init() {
	if File_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDesc), len(file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_goTypes,
		DependencyIndexes: file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_depIdxs,
		MessageInfos:      file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_msgTypes,
	}.Build()
	File_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto = out.File
	file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_goTypes = nil
	file_controller_storage_iam_store_v1_role_global_individual_org_grant_scope_proto_depIdxs = nil
}
