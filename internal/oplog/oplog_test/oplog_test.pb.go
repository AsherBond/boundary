// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: controller/storage/oplog/test/v1/oplog_test.proto

// define a test proto package

package oplog_test

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TestUser for gorm test user model
type TestUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	// @inject_tag: gorm:"default:null"
	Version uint32 `protobuf:"varint,9,opt,name=version,proto3" json:"version,omitempty" gorm:"default:null"`
	// @inject_tag: gorm:"-" json:"-"
	Table string `protobuf:"bytes,7,opt,name=table,proto3" json:"-" gorm:"-"`
}

func (x *TestUser) Reset() {
	*x = TestUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestUser) ProtoMessage() {}

func (x *TestUser) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestUser.ProtoReflect.Descriptor instead.
func (*TestUser) Descriptor() ([]byte, []int) {
	return file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestUser) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestUser) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *TestUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *TestUser) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TestUser) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

// TestCar for gorm test car model
type TestCar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	Model string `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Mpg   int32  `protobuf:"varint,5,opt,name=mpg,proto3" json:"mpg,omitempty"`
	// @inject_tag: gorm:"-" json:"-"
	Table string `protobuf:"bytes,6,opt,name=table,proto3" json:"-" gorm:"-"`
}

func (x *TestCar) Reset() {
	*x = TestCar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCar) ProtoMessage() {}

func (x *TestCar) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCar.ProtoReflect.Descriptor instead.
func (*TestCar) Descriptor() ([]byte, []int) {
	return file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestCar) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestCar) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *TestCar) GetMpg() int32 {
	if x != nil {
		return x.Mpg
	}
	return 0
}

func (x *TestCar) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

// TestRental for gorm test rental model
type TestRental struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CarId  uint32 `protobuf:"varint,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	// @inject_tag: gorm:"-" json:"-"
	Table string `protobuf:"bytes,3,opt,name=table,proto3" json:"-" gorm:"-"`
}

func (x *TestRental) Reset() {
	*x = TestRental{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestRental) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestRental) ProtoMessage() {}

func (x *TestRental) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestRental.ProtoReflect.Descriptor instead.
func (*TestRental) Descriptor() ([]byte, []int) {
	return file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescGZIP(), []int{2}
}

func (x *TestRental) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TestRental) GetCarId() uint32 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *TestRental) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

// TestNonReplayableUser for negative test
type TestNonReplayableUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: gorm:"primary_key"
	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email       string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *TestNonReplayableUser) Reset() {
	*x = TestNonReplayableUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestNonReplayableUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestNonReplayableUser) ProtoMessage() {}

func (x *TestNonReplayableUser) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestNonReplayableUser.ProtoReflect.Descriptor instead.
func (*TestNonReplayableUser) Descriptor() ([]byte, []int) {
	return file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescGZIP(), []int{3}
}

func (x *TestNonReplayableUser) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestNonReplayableUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestNonReplayableUser) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *TestNonReplayableUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_controller_storage_oplog_test_v1_oplog_test_proto protoreflect.FileDescriptor

var file_controller_storage_oplog_test_v1_oplog_test_proto_rawDesc = []byte{
	0x0a, 0x31, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x97, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x57, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6d,
	0x70, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x52, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x74, 0x0a, 0x15,
	0x54, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x61, 0x62, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6f, 0x70, 0x6c,
	0x6f, 0x67, 0x2f, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x3b, 0x6f, 0x70,
	0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescOnce sync.Once
	file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescData = file_controller_storage_oplog_test_v1_oplog_test_proto_rawDesc
)

func file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescGZIP() []byte {
	file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescOnce.Do(func() {
		file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescData)
	})
	return file_controller_storage_oplog_test_v1_oplog_test_proto_rawDescData
}

var file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controller_storage_oplog_test_v1_oplog_test_proto_goTypes = []interface{}{
	(*TestUser)(nil),              // 0: controller.storage.oplog.test.v1.TestUser
	(*TestCar)(nil),               // 1: controller.storage.oplog.test.v1.TestCar
	(*TestRental)(nil),            // 2: controller.storage.oplog.test.v1.TestRental
	(*TestNonReplayableUser)(nil), // 3: controller.storage.oplog.test.v1.TestNonReplayableUser
}
var file_controller_storage_oplog_test_v1_oplog_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_controller_storage_oplog_test_v1_oplog_test_proto_init() }
func file_controller_storage_oplog_test_v1_oplog_test_proto_init() {
	if File_controller_storage_oplog_test_v1_oplog_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCar); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestRental); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestNonReplayableUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_storage_oplog_test_v1_oplog_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_oplog_test_v1_oplog_test_proto_goTypes,
		DependencyIndexes: file_controller_storage_oplog_test_v1_oplog_test_proto_depIdxs,
		MessageInfos:      file_controller_storage_oplog_test_v1_oplog_test_proto_msgTypes,
	}.Build()
	File_controller_storage_oplog_test_v1_oplog_test_proto = out.File
	file_controller_storage_oplog_test_v1_oplog_test_proto_rawDesc = nil
	file_controller_storage_oplog_test_v1_oplog_test_proto_goTypes = nil
	file_controller_storage_oplog_test_v1_oplog_test_proto_depIdxs = nil
}
