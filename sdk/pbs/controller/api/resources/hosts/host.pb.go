// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/api/resources/hosts/v1/host.proto

package hosts

import (
	plugins "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/plugins"
	scopes "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// Host contains all fields related to a Host resource
type Host struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The ID of the Host.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// The Host Catalog of which this Host is a part.
	HostCatalogId string `protobuf:"bytes,20,opt,name=host_catalog_id,proto3" json:"host_catalog_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. Scope information for this resource.
	Scope *scopes.ScopeInfo `protobuf:"bytes,30,opt,name=scope,proto3" json:"scope,omitempty"`
	// Output only. Plugin information for this resource.
	Plugin *plugins.PluginInfo `protobuf:"bytes,35,opt,name=plugin,proto3" json:"plugin,omitempty"`
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
	// The type of the resource.
	Type string `protobuf:"bytes,90,opt,name=type,proto3" json:"type,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. A list of Host Sets containing this Host.
	HostSetIds []string `protobuf:"bytes,100,rep,name=host_set_ids,proto3" json:"host_set_ids,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Types that are valid to be assigned to Attrs:
	//
	//	*Host_Attributes
	//	*Host_StaticHostAttributes
	Attrs isHost_Attrs `protobuf_oneof:"attrs"`
	// Output only.  The list of ip addresses associated with this host.
	IpAddresses []string `protobuf:"bytes,120,rep,name=ip_addresses,json=ipAddresses,proto3" json:"ip_addresses,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only.  The list of dns addresses associated with this host.
	DnsNames []string `protobuf:"bytes,130,rep,name=dns_names,json=dnsNames,proto3" json:"dns_names,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The external ID of the host, if any.
	ExternalId string `protobuf:"bytes,140,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. Refers to the name for a given host provided by the plugin enabled backing service.
	ExternalName string `protobuf:"bytes,150,opt,name=external_name,json=externalName,proto3" json:"external_name,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The available actions on this resource for this user.
	AuthorizedActions []string `protobuf:"bytes,300,rep,name=authorized_actions,proto3" json:"authorized_actions,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *Host) Reset() {
	*x = Host{}
	mi := &file_controller_api_resources_hosts_v1_host_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_hosts_v1_host_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_hosts_v1_host_proto_rawDescGZIP(), []int{0}
}

func (x *Host) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Host) GetHostCatalogId() string {
	if x != nil {
		return x.HostCatalogId
	}
	return ""
}

func (x *Host) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *Host) GetPlugin() *plugins.PluginInfo {
	if x != nil {
		return x.Plugin
	}
	return nil
}

func (x *Host) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *Host) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Host) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *Host) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *Host) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Host) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Host) GetHostSetIds() []string {
	if x != nil {
		return x.HostSetIds
	}
	return nil
}

func (x *Host) GetAttrs() isHost_Attrs {
	if x != nil {
		return x.Attrs
	}
	return nil
}

func (x *Host) GetAttributes() *structpb.Struct {
	if x != nil {
		if x, ok := x.Attrs.(*Host_Attributes); ok {
			return x.Attributes
		}
	}
	return nil
}

func (x *Host) GetStaticHostAttributes() *StaticHostAttributes {
	if x != nil {
		if x, ok := x.Attrs.(*Host_StaticHostAttributes); ok {
			return x.StaticHostAttributes
		}
	}
	return nil
}

func (x *Host) GetIpAddresses() []string {
	if x != nil {
		return x.IpAddresses
	}
	return nil
}

func (x *Host) GetDnsNames() []string {
	if x != nil {
		return x.DnsNames
	}
	return nil
}

func (x *Host) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *Host) GetExternalName() string {
	if x != nil {
		return x.ExternalName
	}
	return ""
}

func (x *Host) GetAuthorizedActions() []string {
	if x != nil {
		return x.AuthorizedActions
	}
	return nil
}

type isHost_Attrs interface {
	isHost_Attrs()
}

type Host_Attributes struct {
	// The attributes that are applicable to the specific Host type.
	Attributes *structpb.Struct `protobuf:"bytes,110,opt,name=attributes,proto3,oneof"`
}

type Host_StaticHostAttributes struct {
	StaticHostAttributes *StaticHostAttributes `protobuf:"bytes,111,opt,name=static_host_attributes,json=staticHostAttributes,proto3,oneof"`
}

func (*Host_Attributes) isHost_Attrs() {}

func (*Host_StaticHostAttributes) isHost_Attrs() {}

type StaticHostAttributes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The address (DNS or IP name) used to reach the Host.
	Address       *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty" class:"public"` // @gotags: `class:"public"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StaticHostAttributes) Reset() {
	*x = StaticHostAttributes{}
	mi := &file_controller_api_resources_hosts_v1_host_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StaticHostAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticHostAttributes) ProtoMessage() {}

func (x *StaticHostAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_hosts_v1_host_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticHostAttributes.ProtoReflect.Descriptor instead.
func (*StaticHostAttributes) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_hosts_v1_host_proto_rawDescGZIP(), []int{1}
}

func (x *StaticHostAttributes) GetAddress() *wrapperspb.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

var File_controller_api_resources_hosts_v1_host_proto protoreflect.FileDescriptor

const file_controller_api_resources_hosts_v1_host_proto_rawDesc = "" +
	"\n" +
	",controller/api/resources/hosts/v1/host.proto\x12!controller.api.resources.hosts.v1\x1a0controller/api/resources/plugins/v1/plugin.proto\x1a.controller/api/resources/scopes/v1/scope.proto\x1a*controller/custom_options/v1/options.proto\x1a\x1bgoogle/api/visibility.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\xf3\a\n" +
	"\x04Host\x12\x0e\n" +
	"\x02id\x18\n" +
	" \x01(\tR\x02id\x12.\n" +
	"\x0fhost_catalog_id\x18\x14 \x01(\tB\x04\xa0\xe3)\x01R\x0fhost_catalog_id\x12C\n" +
	"\x05scope\x18\x1e \x01(\v2-.controller.api.resources.scopes.v1.ScopeInfoR\x05scope\x12G\n" +
	"\x06plugin\x18# \x01(\v2/.controller.api.resources.plugins.v1.PluginInfoR\x06plugin\x12F\n" +
	"\x04name\x18( \x01(\v2\x1c.google.protobuf.StringValueB\x14\xa0\xda)\x01\xc2\xdd)\f\n" +
	"\x04name\x12\x04nameR\x04name\x12b\n" +
	"\vdescription\x182 \x01(\v2\x1c.google.protobuf.StringValueB\"\xa0\xda)\x01\xc2\xdd)\x1a\n" +
	"\vdescription\x12\vdescriptionR\vdescription\x12>\n" +
	"\fcreated_time\x18< \x01(\v2\x1a.google.protobuf.TimestampR\fcreated_time\x12>\n" +
	"\fupdated_time\x18F \x01(\v2\x1a.google.protobuf.TimestampR\fupdated_time\x12\x18\n" +
	"\aversion\x18P \x01(\rR\aversion\x12\x12\n" +
	"\x04type\x18Z \x01(\tR\x04type\x12\"\n" +
	"\fhost_set_ids\x18d \x03(\tR\fhost_set_ids\x12J\n" +
	"\n" +
	"attributes\x18n \x01(\v2\x17.google.protobuf.StructB\x0f\xa0\xda)\x01\x9a\xe3)\adefaultH\x00R\n" +
	"attributes\x12\x8f\x01\n" +
	"\x16static_host_attributes\x18o \x01(\v27.controller.api.resources.hosts.v1.StaticHostAttributesB\x1e\xa0\xda)\x01\x9a\xe3)\x06static\xfa\xd2\xe4\x93\x02\n" +
	"\x12\bINTERNALH\x00R\x14staticHostAttributes\x12!\n" +
	"\fip_addresses\x18x \x03(\tR\vipAddresses\x12\x1c\n" +
	"\tdns_names\x18\x82\x01 \x03(\tR\bdnsNames\x12 \n" +
	"\vexternal_id\x18\x8c\x01 \x01(\tR\n" +
	"externalId\x12$\n" +
	"\rexternal_name\x18\x96\x01 \x01(\tR\fexternalName\x12/\n" +
	"\x12authorized_actions\x18\xac\x02 \x03(\tR\x12authorized_actionsB\a\n" +
	"\x05attrs\"u\n" +
	"\x14StaticHostAttributes\x12]\n" +
	"\aaddress\x18\n" +
	" \x01(\v2\x1c.google.protobuf.StringValueB%\xa0\xda)\x01\xc2\xdd)\x1d\n" +
	"\x12attributes.address\x12\aaddressR\aaddressBLZJgithub.com/hashicorp/boundary/sdk/pbs/controller/api/resources/hosts;hostsb\x06proto3"

var (
	file_controller_api_resources_hosts_v1_host_proto_rawDescOnce sync.Once
	file_controller_api_resources_hosts_v1_host_proto_rawDescData []byte
)

func file_controller_api_resources_hosts_v1_host_proto_rawDescGZIP() []byte {
	file_controller_api_resources_hosts_v1_host_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_hosts_v1_host_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_api_resources_hosts_v1_host_proto_rawDesc), len(file_controller_api_resources_hosts_v1_host_proto_rawDesc)))
	})
	return file_controller_api_resources_hosts_v1_host_proto_rawDescData
}

var file_controller_api_resources_hosts_v1_host_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_api_resources_hosts_v1_host_proto_goTypes = []any{
	(*Host)(nil),                   // 0: controller.api.resources.hosts.v1.Host
	(*StaticHostAttributes)(nil),   // 1: controller.api.resources.hosts.v1.StaticHostAttributes
	(*scopes.ScopeInfo)(nil),       // 2: controller.api.resources.scopes.v1.ScopeInfo
	(*plugins.PluginInfo)(nil),     // 3: controller.api.resources.plugins.v1.PluginInfo
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*structpb.Struct)(nil),        // 6: google.protobuf.Struct
}
var file_controller_api_resources_hosts_v1_host_proto_depIdxs = []int32{
	2, // 0: controller.api.resources.hosts.v1.Host.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	3, // 1: controller.api.resources.hosts.v1.Host.plugin:type_name -> controller.api.resources.plugins.v1.PluginInfo
	4, // 2: controller.api.resources.hosts.v1.Host.name:type_name -> google.protobuf.StringValue
	4, // 3: controller.api.resources.hosts.v1.Host.description:type_name -> google.protobuf.StringValue
	5, // 4: controller.api.resources.hosts.v1.Host.created_time:type_name -> google.protobuf.Timestamp
	5, // 5: controller.api.resources.hosts.v1.Host.updated_time:type_name -> google.protobuf.Timestamp
	6, // 6: controller.api.resources.hosts.v1.Host.attributes:type_name -> google.protobuf.Struct
	1, // 7: controller.api.resources.hosts.v1.Host.static_host_attributes:type_name -> controller.api.resources.hosts.v1.StaticHostAttributes
	4, // 8: controller.api.resources.hosts.v1.StaticHostAttributes.address:type_name -> google.protobuf.StringValue
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_controller_api_resources_hosts_v1_host_proto_init() }
func file_controller_api_resources_hosts_v1_host_proto_init() {
	if File_controller_api_resources_hosts_v1_host_proto != nil {
		return
	}
	file_controller_api_resources_hosts_v1_host_proto_msgTypes[0].OneofWrappers = []any{
		(*Host_Attributes)(nil),
		(*Host_StaticHostAttributes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_api_resources_hosts_v1_host_proto_rawDesc), len(file_controller_api_resources_hosts_v1_host_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_hosts_v1_host_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_hosts_v1_host_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_hosts_v1_host_proto_msgTypes,
	}.Build()
	File_controller_api_resources_hosts_v1_host_proto = out.File
	file_controller_api_resources_hosts_v1_host_proto_goTypes = nil
	file_controller_api_resources_hosts_v1_host_proto_depIdxs = nil
}
