// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: controller/api/resources/hostcatalogs/v1/host_catalog.proto

package hostcatalogs

import (
	plugins "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/plugins"
	scopes "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
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

// HostCatalog manages Hosts and Host Sets
type HostCatalog struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The ID of the host.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// The ID of the Scope of which this Host Catalog is a part.
	ScopeId string `protobuf:"bytes,20,opt,name=scope_id,proto3" json:"scope_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Output only. Scope information for this resource.
	Scope *scopes.ScopeInfo `protobuf:"bytes,30,opt,name=scope,proto3" json:"scope,omitempty"`
	// The ID of the plugin of which this catalog is created.
	PluginId string `protobuf:"bytes,34,opt,name=plugin_id,proto3" json:"plugin_id,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
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
	// The type of Host Catalog.
	Type string `protobuf:"bytes,90,opt,name=type,proto3" json:"type,omitempty" class:"public" eventstream:"observation"` // @gotags: `class:"public" eventstream:"observation"`
	// Types that are valid to be assigned to Attrs:
	//
	//	*HostCatalog_Attributes
	Attrs isHostCatalog_Attrs `protobuf_oneof:"attrs"`
	// Secrets specific to the catalog type. These are never output.
	Secrets *structpb.Struct `protobuf:"bytes,110,opt,name=secrets,proto3" json:"secrets,omitempty"`
	// Output only. The HMAC of the last secrets supplied via the API, if any.
	SecretsHmac string `protobuf:"bytes,120,opt,name=secrets_hmac,json=secretsHmac,proto3" json:"secrets_hmac,omitempty" class:"public"` // @gotags: `class:"public"`
	// Optional worker filter for plugin-subtype host catalogs. Boundary Enterprise only.
	WorkerFilter *wrapperspb.StringValue `protobuf:"bytes,130,opt,name=worker_filter,json=workerFilter,proto3" json:"worker_filter,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The available actions on this resource for this user.
	AuthorizedActions []string `protobuf:"bytes,300,rep,name=authorized_actions,proto3" json:"authorized_actions,omitempty" class:"public"` // @gotags: `class:"public"`
	// Output only. The authorized actions for the scope's collections.
	AuthorizedCollectionActions map[string]*structpb.ListValue `protobuf:"bytes,310,rep,name=authorized_collection_actions,proto3" json:"authorized_collection_actions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // classified as public via taggable implementation
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *HostCatalog) Reset() {
	*x = HostCatalog{}
	mi := &file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HostCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCatalog) ProtoMessage() {}

func (x *HostCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCatalog.ProtoReflect.Descriptor instead.
func (*HostCatalog) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *HostCatalog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HostCatalog) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *HostCatalog) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *HostCatalog) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *HostCatalog) GetPlugin() *plugins.PluginInfo {
	if x != nil {
		return x.Plugin
	}
	return nil
}

func (x *HostCatalog) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *HostCatalog) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *HostCatalog) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *HostCatalog) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *HostCatalog) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *HostCatalog) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *HostCatalog) GetAttrs() isHostCatalog_Attrs {
	if x != nil {
		return x.Attrs
	}
	return nil
}

func (x *HostCatalog) GetAttributes() *structpb.Struct {
	if x != nil {
		if x, ok := x.Attrs.(*HostCatalog_Attributes); ok {
			return x.Attributes
		}
	}
	return nil
}

func (x *HostCatalog) GetSecrets() *structpb.Struct {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *HostCatalog) GetSecretsHmac() string {
	if x != nil {
		return x.SecretsHmac
	}
	return ""
}

func (x *HostCatalog) GetWorkerFilter() *wrapperspb.StringValue {
	if x != nil {
		return x.WorkerFilter
	}
	return nil
}

func (x *HostCatalog) GetAuthorizedActions() []string {
	if x != nil {
		return x.AuthorizedActions
	}
	return nil
}

func (x *HostCatalog) GetAuthorizedCollectionActions() map[string]*structpb.ListValue {
	if x != nil {
		return x.AuthorizedCollectionActions
	}
	return nil
}

type isHostCatalog_Attrs interface {
	isHostCatalog_Attrs()
}

type HostCatalog_Attributes struct {
	// Attributes specific to the catalog type.
	Attributes *structpb.Struct `protobuf:"bytes,100,opt,name=attributes,proto3,oneof"`
}

func (*HostCatalog_Attributes) isHostCatalog_Attrs() {}

var File_controller_api_resources_hostcatalogs_v1_host_catalog_proto protoreflect.FileDescriptor

const file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDesc = "" +
	"\n" +
	";controller/api/resources/hostcatalogs/v1/host_catalog.proto\x12(controller.api.resources.hostcatalogs.v1\x1a0controller/api/resources/plugins/v1/plugin.proto\x1a.controller/api/resources/scopes/v1/scope.proto\x1a*controller/custom_options/v1/options.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\x9e\t\n" +
	"\vHostCatalog\x12\x0e\n" +
	"\x02id\x18\n" +
	" \x01(\tR\x02id\x12\x1a\n" +
	"\bscope_id\x18\x14 \x01(\tR\bscope_id\x12C\n" +
	"\x05scope\x18\x1e \x01(\v2-.controller.api.resources.scopes.v1.ScopeInfoR\x05scope\x12\"\n" +
	"\tplugin_id\x18\" \x01(\tB\x04\xa0\xda)\x01R\tplugin_id\x12G\n" +
	"\x06plugin\x18# \x01(\v2/.controller.api.resources.plugins.v1.PluginInfoR\x06plugin\x12F\n" +
	"\x04name\x18( \x01(\v2\x1c.google.protobuf.StringValueB\x14\xa0\xda)\x01\xc2\xdd)\f\n" +
	"\x04name\x12\x04nameR\x04name\x12b\n" +
	"\vdescription\x182 \x01(\v2\x1c.google.protobuf.StringValueB\"\xa0\xda)\x01\xc2\xdd)\x1a\n" +
	"\vdescription\x12\vdescriptionR\vdescription\x12>\n" +
	"\fcreated_time\x18< \x01(\v2\x1a.google.protobuf.TimestampR\fcreated_time\x12>\n" +
	"\fupdated_time\x18F \x01(\v2\x1a.google.protobuf.TimestampR\fupdated_time\x12\x18\n" +
	"\aversion\x18P \x01(\rR\aversion\x12\x12\n" +
	"\x04type\x18Z \x01(\tR\x04type\x12J\n" +
	"\n" +
	"attributes\x18d \x01(\v2\x17.google.protobuf.StructB\x0f\xa0\xda)\x01\x9a\xe3)\adefaultH\x00R\n" +
	"attributes\x127\n" +
	"\asecrets\x18n \x01(\v2\x17.google.protobuf.StructB\x04\xa0\xda)\x01R\asecrets\x12!\n" +
	"\fsecrets_hmac\x18x \x01(\tR\vsecretsHmac\x12i\n" +
	"\rworker_filter\x18\x82\x01 \x01(\v2\x1c.google.protobuf.StringValueB%\xa0\xda)\x01\xc2\xdd)\x1d\n" +
	"\rworker_filter\x12\fWorkerFilterR\fworkerFilter\x12/\n" +
	"\x12authorized_actions\x18\xac\x02 \x03(\tR\x12authorized_actions\x12\x9d\x01\n" +
	"\x1dauthorized_collection_actions\x18\xb6\x02 \x03(\v2V.controller.api.resources.hostcatalogs.v1.HostCatalog.AuthorizedCollectionActionsEntryR\x1dauthorized_collection_actions\x1aj\n" +
	" AuthorizedCollectionActionsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x120\n" +
	"\x05value\x18\x02 \x01(\v2\x1a.google.protobuf.ListValueR\x05value:\x028\x01B\a\n" +
	"\x05attrsBZZXgithub.com/hashicorp/boundary/sdk/pbs/controller/api/resources/hostcatalogs;hostcatalogsb\x06proto3"

var (
	file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDescOnce sync.Once
	file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDescData []byte
)

func file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDescGZIP() []byte {
	file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDesc), len(file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDesc)))
	})
	return file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDescData
}

var file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_goTypes = []any{
	(*HostCatalog)(nil),            // 0: controller.api.resources.hostcatalogs.v1.HostCatalog
	nil,                            // 1: controller.api.resources.hostcatalogs.v1.HostCatalog.AuthorizedCollectionActionsEntry
	(*scopes.ScopeInfo)(nil),       // 2: controller.api.resources.scopes.v1.ScopeInfo
	(*plugins.PluginInfo)(nil),     // 3: controller.api.resources.plugins.v1.PluginInfo
	(*wrapperspb.StringValue)(nil), // 4: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*structpb.Struct)(nil),        // 6: google.protobuf.Struct
	(*structpb.ListValue)(nil),     // 7: google.protobuf.ListValue
}
var file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_depIdxs = []int32{
	2,  // 0: controller.api.resources.hostcatalogs.v1.HostCatalog.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	3,  // 1: controller.api.resources.hostcatalogs.v1.HostCatalog.plugin:type_name -> controller.api.resources.plugins.v1.PluginInfo
	4,  // 2: controller.api.resources.hostcatalogs.v1.HostCatalog.name:type_name -> google.protobuf.StringValue
	4,  // 3: controller.api.resources.hostcatalogs.v1.HostCatalog.description:type_name -> google.protobuf.StringValue
	5,  // 4: controller.api.resources.hostcatalogs.v1.HostCatalog.created_time:type_name -> google.protobuf.Timestamp
	5,  // 5: controller.api.resources.hostcatalogs.v1.HostCatalog.updated_time:type_name -> google.protobuf.Timestamp
	6,  // 6: controller.api.resources.hostcatalogs.v1.HostCatalog.attributes:type_name -> google.protobuf.Struct
	6,  // 7: controller.api.resources.hostcatalogs.v1.HostCatalog.secrets:type_name -> google.protobuf.Struct
	4,  // 8: controller.api.resources.hostcatalogs.v1.HostCatalog.worker_filter:type_name -> google.protobuf.StringValue
	1,  // 9: controller.api.resources.hostcatalogs.v1.HostCatalog.authorized_collection_actions:type_name -> controller.api.resources.hostcatalogs.v1.HostCatalog.AuthorizedCollectionActionsEntry
	7,  // 10: controller.api.resources.hostcatalogs.v1.HostCatalog.AuthorizedCollectionActionsEntry.value:type_name -> google.protobuf.ListValue
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_init() }
func file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_init() {
	if File_controller_api_resources_hostcatalogs_v1_host_catalog_proto != nil {
		return
	}
	file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_msgTypes[0].OneofWrappers = []any{
		(*HostCatalog_Attributes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDesc), len(file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_msgTypes,
	}.Build()
	File_controller_api_resources_hostcatalogs_v1_host_catalog_proto = out.File
	file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_goTypes = nil
	file_controller_api_resources_hostcatalogs_v1_host_catalog_proto_depIdxs = nil
}
