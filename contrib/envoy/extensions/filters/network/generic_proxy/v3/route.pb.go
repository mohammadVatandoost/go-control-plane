// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: contrib/envoy/extensions/filters/network/generic_proxy/v3/route.proto

package generic_proxyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/cncf/xds/go/xds/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type VirtualHost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the virtual host.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of hosts that will be matched to this virtual host. Wildcard hosts are supported in
	// the suffix or prefix form.
	//
	// Host search order:
	//  1. Exact names: ``www.foo.com``.
	//  2. Suffix wildcards: ``*.foo.com`` or ``*-bar.foo.com``.
	//  3. Prefix wildcards: ``foo.*`` or ``foo-*``.
	//  4. Special wildcard ``*`` matching any host and will be the default virtual host.
	//
	// .. note::
	//
	//   The wildcard will not match the empty string.
	//   e.g. ``*-bar.foo.com`` will match ``baz-bar.foo.com`` but not ``-bar.foo.com``.
	//   The longest wildcards match first.
	//   Only a single virtual host in the entire route configuration can match on ``*``. A domain
	//   must be unique across all virtual hosts or the config will fail to load.
	Hosts []string `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// The match tree to use when resolving route actions for incoming requests.
	Routes *v3.Matcher `protobuf:"bytes,3,opt,name=routes,proto3" json:"routes,omitempty"`
}

func (x *VirtualHost) Reset() {
	*x = VirtualHost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualHost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualHost) ProtoMessage() {}

func (x *VirtualHost) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualHost.ProtoReflect.Descriptor instead.
func (*VirtualHost) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescGZIP(), []int{0}
}

func (x *VirtualHost) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VirtualHost) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *VirtualHost) GetRoutes() *v3.Matcher {
	if x != nil {
		return x.Routes
	}
	return nil
}

// The generic proxy makes use of the `xds matching API` for routing configurations.
//
// In the below example, we combine a top level tree matcher with a linear matcher to match
// the incoming requests, and send the matching requests to v1 of the upstream service.
//
// .. code-block:: yaml
//
//   name: example
//   routes:
//     matcher_tree:
//       input:
//         name: request-service
//         typed_config:
//           "@type": type.googleapis.com/envoy.extensions.filters.network.generic_proxy.matcher.v3.ServiceMatchInput
//       exact_match_map:
//         map:
//           service_name_0:
//             matcher:
//               matcher_list:
//                 matchers:
//                 - predicate:
//                     and_matcher:
//                       predicate:
//                       - single_predicate:
//                           input:
//                             name: request-properties
//                             typed_config:
//                               "@type": type.googleapis.com/envoy.extensions.filters.network.generic_proxy.matcher.v3.PropertyMatchInput
//                               property_name: version
//                           value_match:
//                             exact: v1
//                       - single_predicate:
//                           input:
//                             name: request-properties
//                             typed_config:
//                               "@type": type.googleapis.com/envoy.extensions.filters.network.generic_proxy.matcher.v3.PropertyMatchInput
//                               property_name: user
//                           value_match:
//                             exact: john
//                   on_match:
//                     action:
//                       name: route
//                       typed_config:
//                         "@type": type.googleapis.com/envoy.extensions.filters.network.generic_proxy.action.v3.routeAction
//                         cluster: cluster_0
type RouteConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the route configuration. For example, it might match route_config_name in
	// envoy.extensions.filters.network.generic_proxy.v3.Rds.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The match tree to use when resolving route actions for incoming requests.
	// If no any virtual host is configured in the ``virtual_hosts`` field or no special wildcard
	// virtual host is configured, the ``routes`` field will be used as the default route table.
	// If both the wildcard virtual host and ``routes`` are configured, the configuration will fail
	// to load.
	Routes *v3.Matcher `protobuf:"bytes,2,opt,name=routes,proto3" json:"routes,omitempty"`
	// An array of virtual hosts that make up the route table.
	VirtualHosts []*VirtualHost `protobuf:"bytes,3,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
}

func (x *RouteConfiguration) Reset() {
	*x = RouteConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteConfiguration) ProtoMessage() {}

func (x *RouteConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteConfiguration.ProtoReflect.Descriptor instead.
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescGZIP(), []int{1}
}

func (x *RouteConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteConfiguration) GetRoutes() *v3.Matcher {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *RouteConfiguration) GetVirtualHosts() []*VirtualHost {
	if x != nil {
		return x.VirtualHosts
	}
	return nil
}

var File_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto protoreflect.FileDescriptor

var file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDesc = []byte{
	0x0a, 0x45, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x78, 0x64, 0x73,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0b, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x12, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x63, 0x0a,
	0x0d, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x0c, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f, 0x73,
	0x74, 0x73, 0x42, 0xc9, 0x01, 0x0a, 0x3f, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescOnce sync.Once
	file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescData = file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDesc
)

func file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescGZIP() []byte {
	file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescOnce.Do(func() {
		file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescData)
	})
	return file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDescData
}

var file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_goTypes = []interface{}{
	(*VirtualHost)(nil),        // 0: envoy.extensions.filters.network.generic_proxy.v3.VirtualHost
	(*RouteConfiguration)(nil), // 1: envoy.extensions.filters.network.generic_proxy.v3.RouteConfiguration
	(*v3.Matcher)(nil),         // 2: xds.type.matcher.v3.Matcher
}
var file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.network.generic_proxy.v3.VirtualHost.routes:type_name -> xds.type.matcher.v3.Matcher
	2, // 1: envoy.extensions.filters.network.generic_proxy.v3.RouteConfiguration.routes:type_name -> xds.type.matcher.v3.Matcher
	0, // 2: envoy.extensions.filters.network.generic_proxy.v3.RouteConfiguration.virtual_hosts:type_name -> envoy.extensions.filters.network.generic_proxy.v3.VirtualHost
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_init() }
func file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_init() {
	if File_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualHost); i {
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
		file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteConfiguration); i {
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
			RawDescriptor: file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_goTypes,
		DependencyIndexes: file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_depIdxs,
		MessageInfos:      file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_msgTypes,
	}.Build()
	File_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto = out.File
	file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_rawDesc = nil
	file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_goTypes = nil
	file_contrib_envoy_extensions_filters_network_generic_proxy_v3_route_proto_depIdxs = nil
}
