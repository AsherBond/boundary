// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/storage/policy/storage/store/v1/policy.proto

// Package store provides protobufs for storing types in the static
// credential package.

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
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

type Policy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// public_id is a surrogate key suitable for use in a public API.
	// @inject_tag: `gorm:"primary_key"`
	PublicId string `protobuf:"bytes,1,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty" gorm:"primary_key"`
	// scope_id must be either global or an org scope.
	// @inject_tag: `gorm:"not_null"`
	ScopeId string `protobuf:"bytes,2,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty" gorm:"not_null"`
	// retain_for_days is the number of days for which a session recording will be
	// retained. Must be provided.
	// @inject_tag: `gorm:"not_null"`
	RetainForDays int32 `protobuf:"varint,3,opt,name=retain_for_days,json=retainForDays,proto3" json:"retain_for_days,omitempty" gorm:"not_null"`
	// retain_for_days_overridable signals whether this storage policy's retention
	// duration can be overridden.
	RetainForDaysOverridable bool `protobuf:"varint,4,opt,name=retain_for_days_overridable,json=retainForDaysOverridable,proto3" json:"retain_for_days_overridable,omitempty"`
	// delete_after_days is the number of days after which a session recording
	// will be automatically deleted.
	// @inject_tag: `gorm:"not_null"`
	DeleteAfterDays int32 `protobuf:"varint,5,opt,name=delete_after_days,json=deleteAfterDays,proto3" json:"delete_after_days,omitempty" gorm:"not_null"`
	// delete_after_days_overridable signals whether this storage policy's
	// deletion policy can be overridden.
	DeleteAfterDaysOverridable bool `protobuf:"varint,6,opt,name=delete_after_days_overridable,json=deleteAfterDaysOverridable,proto3" json:"delete_after_days_overridable,omitempty"`
	// name is optional.
	// @inject_tag: `gorm:"default:null"`
	Name string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty" gorm:"default:null"`
	// description is optional.
	// @inject_tag: `gorm:"default:null"`
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty" gorm:"default:null"`
	// The create_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// The update_time is set by the database.
	// @inject_tag: `gorm:"default:current_timestamp"`
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty" gorm:"default:current_timestamp"`
	// version allows optimistic locking of the resource.
	// @inject_tag: `gorm:"default:null"`
	Version       uint32 `protobuf:"varint,11,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Policy) Reset() {
	*x = Policy{}
	mi := &file_controller_storage_policy_storage_store_v1_policy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_policy_storage_store_v1_policy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_controller_storage_policy_storage_store_v1_policy_proto_rawDescGZIP(), []int{0}
}

func (x *Policy) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Policy) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *Policy) GetRetainForDays() int32 {
	if x != nil {
		return x.RetainForDays
	}
	return 0
}

func (x *Policy) GetRetainForDaysOverridable() bool {
	if x != nil {
		return x.RetainForDaysOverridable
	}
	return false
}

func (x *Policy) GetDeleteAfterDays() int32 {
	if x != nil {
		return x.DeleteAfterDays
	}
	return 0
}

func (x *Policy) GetDeleteAfterDaysOverridable() bool {
	if x != nil {
		return x.DeleteAfterDaysOverridable
	}
	return false
}

func (x *Policy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Policy) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Policy) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Policy) GetUpdateTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Policy) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_controller_storage_policy_storage_store_v1_policy_proto protoreflect.FileDescriptor

const file_controller_storage_policy_storage_store_v1_policy_proto_rawDesc = "" +
	"\n" +
	"7controller/storage/policy/storage/store/v1/policy.proto\x12*controller.storage.policy.storage.store.v1\x1a*controller/custom_options/v1/options.proto\x1a/controller/storage/timestamp/v1/timestamp.proto\"\xa4\x06\n" +
	"\x06Policy\x12\x1b\n" +
	"\tpublic_id\x18\x01 \x01(\tR\bpublicId\x12\x19\n" +
	"\bscope_id\x18\x02 \x01(\tR\ascopeId\x12W\n" +
	"\x0fretain_for_days\x18\x03 \x01(\x05B/\xc2\xdd)+\n" +
	"\rRetainForDays\x12\x1aattributes.retain_for.daysR\rretainForDays\x12\x80\x01\n" +
	"\x1bretain_for_days_overridable\x18\x04 \x01(\bBA\xc2\xdd)=\n" +
	"\x18RetainForDaysOverridable\x12!attributes.retain_for.overridableR\x18retainForDaysOverridable\x12_\n" +
	"\x11delete_after_days\x18\x05 \x01(\x05B3\xc2\xdd)/\n" +
	"\x0fDeleteAfterDays\x12\x1cattributes.delete_after.daysR\x0fdeleteAfterDays\x12\x88\x01\n" +
	"\x1ddelete_after_days_overridable\x18\x06 \x01(\bBE\xc2\xdd)A\n" +
	"\x1aDeleteAfterDaysOverridable\x12#attributes.delete_after.overridableR\x1adeleteAfterDaysOverridable\x12$\n" +
	"\x04name\x18\a \x01(\tB\x10\xc2\xdd)\f\n" +
	"\x04Name\x12\x04nameR\x04name\x12@\n" +
	"\vdescription\x18\b \x01(\tB\x1e\xc2\xdd)\x1a\n" +
	"\vDescription\x12\vdescriptionR\vdescription\x12K\n" +
	"\vcreate_time\x18\t \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"createTime\x12K\n" +
	"\vupdate_time\x18\n" +
	" \x01(\v2*.controller.storage.timestamp.v1.TimestampR\n" +
	"updateTime\x12\x18\n" +
	"\aversion\x18\v \x01(\rR\aversionBCZAgithub.com/hashicorp/boundary/internal/policy/storage/store;storeb\x06proto3"

var (
	file_controller_storage_policy_storage_store_v1_policy_proto_rawDescOnce sync.Once
	file_controller_storage_policy_storage_store_v1_policy_proto_rawDescData []byte
)

func file_controller_storage_policy_storage_store_v1_policy_proto_rawDescGZIP() []byte {
	file_controller_storage_policy_storage_store_v1_policy_proto_rawDescOnce.Do(func() {
		file_controller_storage_policy_storage_store_v1_policy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_storage_policy_storage_store_v1_policy_proto_rawDesc), len(file_controller_storage_policy_storage_store_v1_policy_proto_rawDesc)))
	})
	return file_controller_storage_policy_storage_store_v1_policy_proto_rawDescData
}

var file_controller_storage_policy_storage_store_v1_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_storage_policy_storage_store_v1_policy_proto_goTypes = []any{
	(*Policy)(nil),              // 0: controller.storage.policy.storage.store.v1.Policy
	(*timestamp.Timestamp)(nil), // 1: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_policy_storage_store_v1_policy_proto_depIdxs = []int32{
	1, // 0: controller.storage.policy.storage.store.v1.Policy.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	1, // 1: controller.storage.policy.storage.store.v1.Policy.update_time:type_name -> controller.storage.timestamp.v1.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_controller_storage_policy_storage_store_v1_policy_proto_init() }
func file_controller_storage_policy_storage_store_v1_policy_proto_init() {
	if File_controller_storage_policy_storage_store_v1_policy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_storage_policy_storage_store_v1_policy_proto_rawDesc), len(file_controller_storage_policy_storage_store_v1_policy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_policy_storage_store_v1_policy_proto_goTypes,
		DependencyIndexes: file_controller_storage_policy_storage_store_v1_policy_proto_depIdxs,
		MessageInfos:      file_controller_storage_policy_storage_store_v1_policy_proto_msgTypes,
	}.Build()
	File_controller_storage_policy_storage_store_v1_policy_proto = out.File
	file_controller_storage_policy_storage_store_v1_policy_proto_goTypes = nil
	file_controller_storage_policy_storage_store_v1_policy_proto_depIdxs = nil
}
