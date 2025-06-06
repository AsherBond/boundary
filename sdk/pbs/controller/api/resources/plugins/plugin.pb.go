// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/api/resources/plugins/v1/plugin.proto

package plugins

import (
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

type PluginInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The ID of the Plugin.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. The name of the plugin resource in boundary, if any.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The description of the plugin in boundary, if any.
	Description   string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PluginInfo) Reset() {
	*x = PluginInfo{}
	mi := &file_controller_api_resources_plugins_v1_plugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PluginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginInfo) ProtoMessage() {}

func (x *PluginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_plugins_v1_plugin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginInfo.ProtoReflect.Descriptor instead.
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_plugins_v1_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *PluginInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PluginInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PluginInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_controller_api_resources_plugins_v1_plugin_proto protoreflect.FileDescriptor

const file_controller_api_resources_plugins_v1_plugin_proto_rawDesc = "" +
	"\n" +
	"0controller/api/resources/plugins/v1/plugin.proto\x12#controller.api.resources.plugins.v1\"R\n" +
	"\n" +
	"PluginInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescriptionBPZNgithub.com/hashicorp/boundary/sdk/pbs/controller/api/resources/plugins;pluginsb\x06proto3"

var (
	file_controller_api_resources_plugins_v1_plugin_proto_rawDescOnce sync.Once
	file_controller_api_resources_plugins_v1_plugin_proto_rawDescData []byte
)

func file_controller_api_resources_plugins_v1_plugin_proto_rawDescGZIP() []byte {
	file_controller_api_resources_plugins_v1_plugin_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_plugins_v1_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_api_resources_plugins_v1_plugin_proto_rawDesc), len(file_controller_api_resources_plugins_v1_plugin_proto_rawDesc)))
	})
	return file_controller_api_resources_plugins_v1_plugin_proto_rawDescData
}

var file_controller_api_resources_plugins_v1_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_api_resources_plugins_v1_plugin_proto_goTypes = []any{
	(*PluginInfo)(nil), // 0: controller.api.resources.plugins.v1.PluginInfo
}
var file_controller_api_resources_plugins_v1_plugin_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_controller_api_resources_plugins_v1_plugin_proto_init() }
func file_controller_api_resources_plugins_v1_plugin_proto_init() {
	if File_controller_api_resources_plugins_v1_plugin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_api_resources_plugins_v1_plugin_proto_rawDesc), len(file_controller_api_resources_plugins_v1_plugin_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_plugins_v1_plugin_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_plugins_v1_plugin_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_plugins_v1_plugin_proto_msgTypes,
	}.Build()
	File_controller_api_resources_plugins_v1_plugin_proto = out.File
	file_controller_api_resources_plugins_v1_plugin_proto_goTypes = nil
	file_controller_api_resources_plugins_v1_plugin_proto_depIdxs = nil
}
