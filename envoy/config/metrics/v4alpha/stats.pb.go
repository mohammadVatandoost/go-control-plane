// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/metrics/v4alpha/stats.proto

package envoy_config_metrics_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StatsSink struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*StatsSink_TypedConfig
	ConfigType           isStatsSink_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StatsSink) Reset()         { *m = StatsSink{} }
func (m *StatsSink) String() string { return proto.CompactTextString(m) }
func (*StatsSink) ProtoMessage()    {}
func (*StatsSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b37e1e73c1c12d, []int{0}
}

func (m *StatsSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsSink.Unmarshal(m, b)
}
func (m *StatsSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsSink.Marshal(b, m, deterministic)
}
func (m *StatsSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsSink.Merge(m, src)
}
func (m *StatsSink) XXX_Size() int {
	return xxx_messageInfo_StatsSink.Size(m)
}
func (m *StatsSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsSink.DiscardUnknown(m)
}

var xxx_messageInfo_StatsSink proto.InternalMessageInfo

func (m *StatsSink) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isStatsSink_ConfigType interface {
	isStatsSink_ConfigType()
}

type StatsSink_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*StatsSink_TypedConfig) isStatsSink_ConfigType() {}

func (m *StatsSink) GetConfigType() isStatsSink_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *StatsSink) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*StatsSink_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsSink_TypedConfig)(nil),
	}
}

type StatsConfig struct {
	StatsTags            []*TagSpecifier     `protobuf:"bytes,1,rep,name=stats_tags,json=statsTags,proto3" json:"stats_tags,omitempty"`
	UseAllDefaultTags    *wrappers.BoolValue `protobuf:"bytes,2,opt,name=use_all_default_tags,json=useAllDefaultTags,proto3" json:"use_all_default_tags,omitempty"`
	StatsMatcher         *StatsMatcher       `protobuf:"bytes,3,opt,name=stats_matcher,json=statsMatcher,proto3" json:"stats_matcher,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StatsConfig) Reset()         { *m = StatsConfig{} }
func (m *StatsConfig) String() string { return proto.CompactTextString(m) }
func (*StatsConfig) ProtoMessage()    {}
func (*StatsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b37e1e73c1c12d, []int{1}
}

func (m *StatsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsConfig.Unmarshal(m, b)
}
func (m *StatsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsConfig.Marshal(b, m, deterministic)
}
func (m *StatsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsConfig.Merge(m, src)
}
func (m *StatsConfig) XXX_Size() int {
	return xxx_messageInfo_StatsConfig.Size(m)
}
func (m *StatsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_StatsConfig proto.InternalMessageInfo

func (m *StatsConfig) GetStatsTags() []*TagSpecifier {
	if m != nil {
		return m.StatsTags
	}
	return nil
}

func (m *StatsConfig) GetUseAllDefaultTags() *wrappers.BoolValue {
	if m != nil {
		return m.UseAllDefaultTags
	}
	return nil
}

func (m *StatsConfig) GetStatsMatcher() *StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

type StatsMatcher struct {
	// Types that are valid to be assigned to StatsMatcher:
	//	*StatsMatcher_RejectAll
	//	*StatsMatcher_ExclusionList
	//	*StatsMatcher_InclusionList
	StatsMatcher         isStatsMatcher_StatsMatcher `protobuf_oneof:"stats_matcher"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *StatsMatcher) Reset()         { *m = StatsMatcher{} }
func (m *StatsMatcher) String() string { return proto.CompactTextString(m) }
func (*StatsMatcher) ProtoMessage()    {}
func (*StatsMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b37e1e73c1c12d, []int{2}
}

func (m *StatsMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsMatcher.Unmarshal(m, b)
}
func (m *StatsMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsMatcher.Marshal(b, m, deterministic)
}
func (m *StatsMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsMatcher.Merge(m, src)
}
func (m *StatsMatcher) XXX_Size() int {
	return xxx_messageInfo_StatsMatcher.Size(m)
}
func (m *StatsMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StatsMatcher proto.InternalMessageInfo

type isStatsMatcher_StatsMatcher interface {
	isStatsMatcher_StatsMatcher()
}

type StatsMatcher_RejectAll struct {
	RejectAll bool `protobuf:"varint,1,opt,name=reject_all,json=rejectAll,proto3,oneof"`
}

type StatsMatcher_ExclusionList struct {
	ExclusionList *v4alpha.ListStringMatcher `protobuf:"bytes,2,opt,name=exclusion_list,json=exclusionList,proto3,oneof"`
}

type StatsMatcher_InclusionList struct {
	InclusionList *v4alpha.ListStringMatcher `protobuf:"bytes,3,opt,name=inclusion_list,json=inclusionList,proto3,oneof"`
}

func (*StatsMatcher_RejectAll) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_ExclusionList) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_InclusionList) isStatsMatcher_StatsMatcher() {}

func (m *StatsMatcher) GetStatsMatcher() isStatsMatcher_StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

func (m *StatsMatcher) GetRejectAll() bool {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_RejectAll); ok {
		return x.RejectAll
	}
	return false
}

func (m *StatsMatcher) GetExclusionList() *v4alpha.ListStringMatcher {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_ExclusionList); ok {
		return x.ExclusionList
	}
	return nil
}

func (m *StatsMatcher) GetInclusionList() *v4alpha.ListStringMatcher {
	if x, ok := m.GetStatsMatcher().(*StatsMatcher_InclusionList); ok {
		return x.InclusionList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsMatcher_RejectAll)(nil),
		(*StatsMatcher_ExclusionList)(nil),
		(*StatsMatcher_InclusionList)(nil),
	}
}

type TagSpecifier struct {
	TagName string `protobuf:"bytes,1,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	// Types that are valid to be assigned to TagValue:
	//	*TagSpecifier_Regex
	//	*TagSpecifier_FixedValue
	TagValue             isTagSpecifier_TagValue `protobuf_oneof:"tag_value"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *TagSpecifier) Reset()         { *m = TagSpecifier{} }
func (m *TagSpecifier) String() string { return proto.CompactTextString(m) }
func (*TagSpecifier) ProtoMessage()    {}
func (*TagSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b37e1e73c1c12d, []int{3}
}

func (m *TagSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagSpecifier.Unmarshal(m, b)
}
func (m *TagSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagSpecifier.Marshal(b, m, deterministic)
}
func (m *TagSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagSpecifier.Merge(m, src)
}
func (m *TagSpecifier) XXX_Size() int {
	return xxx_messageInfo_TagSpecifier.Size(m)
}
func (m *TagSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_TagSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_TagSpecifier proto.InternalMessageInfo

func (m *TagSpecifier) GetTagName() string {
	if m != nil {
		return m.TagName
	}
	return ""
}

type isTagSpecifier_TagValue interface {
	isTagSpecifier_TagValue()
}

type TagSpecifier_Regex struct {
	Regex string `protobuf:"bytes,2,opt,name=regex,proto3,oneof"`
}

type TagSpecifier_FixedValue struct {
	FixedValue string `protobuf:"bytes,3,opt,name=fixed_value,json=fixedValue,proto3,oneof"`
}

func (*TagSpecifier_Regex) isTagSpecifier_TagValue() {}

func (*TagSpecifier_FixedValue) isTagSpecifier_TagValue() {}

func (m *TagSpecifier) GetTagValue() isTagSpecifier_TagValue {
	if m != nil {
		return m.TagValue
	}
	return nil
}

func (m *TagSpecifier) GetRegex() string {
	if x, ok := m.GetTagValue().(*TagSpecifier_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *TagSpecifier) GetFixedValue() string {
	if x, ok := m.GetTagValue().(*TagSpecifier_FixedValue); ok {
		return x.FixedValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TagSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TagSpecifier_Regex)(nil),
		(*TagSpecifier_FixedValue)(nil),
	}
}

type StatsdSink struct {
	// Types that are valid to be assigned to StatsdSpecifier:
	//	*StatsdSink_Address
	//	*StatsdSink_TcpClusterName
	StatsdSpecifier      isStatsdSink_StatsdSpecifier `protobuf_oneof:"statsd_specifier"`
	Prefix               string                       `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StatsdSink) Reset()         { *m = StatsdSink{} }
func (m *StatsdSink) String() string { return proto.CompactTextString(m) }
func (*StatsdSink) ProtoMessage()    {}
func (*StatsdSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b37e1e73c1c12d, []int{4}
}

func (m *StatsdSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsdSink.Unmarshal(m, b)
}
func (m *StatsdSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsdSink.Marshal(b, m, deterministic)
}
func (m *StatsdSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsdSink.Merge(m, src)
}
func (m *StatsdSink) XXX_Size() int {
	return xxx_messageInfo_StatsdSink.Size(m)
}
func (m *StatsdSink) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsdSink.DiscardUnknown(m)
}

var xxx_messageInfo_StatsdSink proto.InternalMessageInfo

type isStatsdSink_StatsdSpecifier interface {
	isStatsdSink_StatsdSpecifier()
}

type StatsdSink_Address struct {
	Address *v4alpha1.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

type StatsdSink_TcpClusterName struct {
	TcpClusterName string `protobuf:"bytes,2,opt,name=tcp_cluster_name,json=tcpClusterName,proto3,oneof"`
}

func (*StatsdSink_Address) isStatsdSink_StatsdSpecifier() {}

func (*StatsdSink_TcpClusterName) isStatsdSink_StatsdSpecifier() {}

func (m *StatsdSink) GetStatsdSpecifier() isStatsdSink_StatsdSpecifier {
	if m != nil {
		return m.StatsdSpecifier
	}
	return nil
}

func (m *StatsdSink) GetAddress() *v4alpha1.Address {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *StatsdSink) GetTcpClusterName() string {
	if x, ok := m.GetStatsdSpecifier().(*StatsdSink_TcpClusterName); ok {
		return x.TcpClusterName
	}
	return ""
}

func (m *StatsdSink) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StatsdSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StatsdSink_Address)(nil),
		(*StatsdSink_TcpClusterName)(nil),
	}
}

type DogStatsdSink struct {
	// Types that are valid to be assigned to DogStatsdSpecifier:
	//	*DogStatsdSink_Address
	DogStatsdSpecifier   isDogStatsdSink_DogStatsdSpecifier `protobuf_oneof:"dog_statsd_specifier"`
	Prefix               string                             `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *DogStatsdSink) Reset()         { *m = DogStatsdSink{} }
func (m *DogStatsdSink) String() string { return proto.CompactTextString(m) }
func (*DogStatsdSink) ProtoMessage()    {}
func (*DogStatsdSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b37e1e73c1c12d, []int{5}
}

func (m *DogStatsdSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogStatsdSink.Unmarshal(m, b)
}
func (m *DogStatsdSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogStatsdSink.Marshal(b, m, deterministic)
}
func (m *DogStatsdSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogStatsdSink.Merge(m, src)
}
func (m *DogStatsdSink) XXX_Size() int {
	return xxx_messageInfo_DogStatsdSink.Size(m)
}
func (m *DogStatsdSink) XXX_DiscardUnknown() {
	xxx_messageInfo_DogStatsdSink.DiscardUnknown(m)
}

var xxx_messageInfo_DogStatsdSink proto.InternalMessageInfo

type isDogStatsdSink_DogStatsdSpecifier interface {
	isDogStatsdSink_DogStatsdSpecifier()
}

type DogStatsdSink_Address struct {
	Address *v4alpha1.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

func (*DogStatsdSink_Address) isDogStatsdSink_DogStatsdSpecifier() {}

func (m *DogStatsdSink) GetDogStatsdSpecifier() isDogStatsdSink_DogStatsdSpecifier {
	if m != nil {
		return m.DogStatsdSpecifier
	}
	return nil
}

func (m *DogStatsdSink) GetAddress() *v4alpha1.Address {
	if x, ok := m.GetDogStatsdSpecifier().(*DogStatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (m *DogStatsdSink) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DogStatsdSink) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DogStatsdSink_Address)(nil),
	}
}

type HystrixSink struct {
	NumBuckets           int64    `protobuf:"varint,1,opt,name=num_buckets,json=numBuckets,proto3" json:"num_buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HystrixSink) Reset()         { *m = HystrixSink{} }
func (m *HystrixSink) String() string { return proto.CompactTextString(m) }
func (*HystrixSink) ProtoMessage()    {}
func (*HystrixSink) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6b37e1e73c1c12d, []int{6}
}

func (m *HystrixSink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HystrixSink.Unmarshal(m, b)
}
func (m *HystrixSink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HystrixSink.Marshal(b, m, deterministic)
}
func (m *HystrixSink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HystrixSink.Merge(m, src)
}
func (m *HystrixSink) XXX_Size() int {
	return xxx_messageInfo_HystrixSink.Size(m)
}
func (m *HystrixSink) XXX_DiscardUnknown() {
	xxx_messageInfo_HystrixSink.DiscardUnknown(m)
}

var xxx_messageInfo_HystrixSink proto.InternalMessageInfo

func (m *HystrixSink) GetNumBuckets() int64 {
	if m != nil {
		return m.NumBuckets
	}
	return 0
}

func init() {
	proto.RegisterType((*StatsSink)(nil), "envoy.config.metrics.v4alpha.StatsSink")
	proto.RegisterType((*StatsConfig)(nil), "envoy.config.metrics.v4alpha.StatsConfig")
	proto.RegisterType((*StatsMatcher)(nil), "envoy.config.metrics.v4alpha.StatsMatcher")
	proto.RegisterType((*TagSpecifier)(nil), "envoy.config.metrics.v4alpha.TagSpecifier")
	proto.RegisterType((*StatsdSink)(nil), "envoy.config.metrics.v4alpha.StatsdSink")
	proto.RegisterType((*DogStatsdSink)(nil), "envoy.config.metrics.v4alpha.DogStatsdSink")
	proto.RegisterType((*HystrixSink)(nil), "envoy.config.metrics.v4alpha.HystrixSink")
}

func init() {
	proto.RegisterFile("envoy/config/metrics/v4alpha/stats.proto", fileDescriptor_c6b37e1e73c1c12d)
}

var fileDescriptor_c6b37e1e73c1c12d = []byte{
	// 800 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xeb, 0x44,
	0x14, 0x8d, 0xe3, 0xf7, 0xf2, 0x92, 0xeb, 0xe4, 0x29, 0x58, 0x11, 0xfd, 0xa0, 0xb4, 0x69, 0xda,
	0xd2, 0x10, 0xc0, 0x96, 0x5a, 0x36, 0x64, 0x81, 0x14, 0xb7, 0x8b, 0xf0, 0x5d, 0x39, 0x55, 0x17,
	0x6c, 0xac, 0xa9, 0x3d, 0x71, 0x4d, 0x1d, 0x8f, 0x35, 0x1e, 0x87, 0x64, 0xd7, 0x25, 0xbf, 0x81,
	0x1d, 0xac, 0xf8, 0x09, 0x88, 0x3d, 0x12, 0x5b, 0x24, 0x7e, 0x0c, 0x62, 0x85, 0xe6, 0x23, 0x26,
	0xe9, 0x47, 0x2a, 0xa1, 0xb7, 0xb3, 0xe7, 0xdc, 0x39, 0x3e, 0xe7, 0xdc, 0xeb, 0x0b, 0x5d, 0x9c,
	0x4c, 0xc9, 0xdc, 0xf6, 0x49, 0x32, 0x8e, 0x42, 0x7b, 0x82, 0x19, 0x8d, 0xfc, 0xcc, 0x9e, 0x7e,
	0x8c, 0xe2, 0xf4, 0x06, 0xd9, 0x19, 0x43, 0x2c, 0xb3, 0x52, 0x4a, 0x18, 0x31, 0x77, 0x44, 0xa5,
	0x25, 0x2b, 0x2d, 0x55, 0x69, 0xa9, 0xca, 0xed, 0xe3, 0x15, 0x1e, 0x9f, 0x50, 0x5c, 0x90, 0xa0,
	0x20, 0xa0, 0x38, 0x53, 0x34, 0x8b, 0x42, 0x36, 0x4f, 0xb1, 0x3d, 0x41, 0xcc, 0xbf, 0xc1, 0x74,
	0xe9, 0x73, 0x34, 0x4a, 0x42, 0x55, 0xb8, 0x15, 0x12, 0x12, 0xc6, 0xd8, 0x16, 0x6f, 0xd7, 0xf9,
	0xd8, 0x46, 0xc9, 0x5c, 0x41, 0x3b, 0xf7, 0xa1, 0x8c, 0xd1, 0xdc, 0x67, 0x0a, 0xdd, 0xbd, 0x8f,
	0x7e, 0x4f, 0x51, 0x9a, 0x62, 0xba, 0x50, 0xf0, 0x6e, 0x1e, 0xa4, 0xc8, 0x46, 0x49, 0x42, 0x18,
	0x62, 0x11, 0x49, 0x32, 0x61, 0x33, 0x5f, 0xc0, 0xfb, 0x0f, 0xe0, 0x29, 0xa6, 0x59, 0x44, 0x92,
	0xff, 0xa4, 0x6d, 0x4c, 0x51, 0x1c, 0x05, 0x88, 0x61, 0x7b, 0xf1, 0x20, 0x81, 0xce, 0x4f, 0x1a,
	0xd4, 0x46, 0x3c, 0xb3, 0x51, 0x94, 0xdc, 0x9a, 0x26, 0xbc, 0x48, 0xd0, 0x04, 0x6f, 0x6a, 0x6d,
	0xad, 0x5b, 0x73, 0xc5, 0xb3, 0xf9, 0x09, 0xd4, 0xb9, 0xf5, 0xc0, 0x93, 0x49, 0x6d, 0xea, 0x6d,
	0xad, 0x6b, 0x9c, 0xb4, 0x2c, 0xa9, 0xd9, 0x5a, 0x68, 0xb6, 0x06, 0xc9, 0x7c, 0x58, 0x72, 0x0d,
	0x51, 0x7b, 0x26, 0x4a, 0xfb, 0xdd, 0x1f, 0x7f, 0xff, 0x61, 0xf7, 0x00, 0xf6, 0x1f, 0xef, 0xc3,
	0xa9, 0x55, 0x7c, 0xd8, 0x69, 0x80, 0x21, 0x61, 0x8f, 0xdf, 0xff, 0xfc, 0x45, 0xb5, 0xdc, 0xd4,
	0xdd, 0x8a, 0x3c, 0xea, 0xfc, 0x5c, 0x06, 0x43, 0x94, 0x4a, 0x5a, 0xf3, 0x33, 0x00, 0xd1, 0x66,
	0x8f, 0xa1, 0x30, 0xdb, 0xd4, 0xda, 0x7a, 0xd7, 0x38, 0xe9, 0x59, 0xeb, 0x9a, 0x6d, 0x5d, 0xa2,
	0x70, 0x94, 0x62, 0x3f, 0x1a, 0x47, 0x98, 0xba, 0x35, 0x71, 0xfb, 0x12, 0x85, 0x99, 0xf9, 0x05,
	0xb4, 0xf2, 0x0c, 0x7b, 0x28, 0x8e, 0xbd, 0x00, 0x8f, 0x51, 0x1e, 0x33, 0x49, 0x5a, 0x16, 0x26,
	0xb7, 0x1f, 0x98, 0x74, 0x08, 0x89, 0xaf, 0x50, 0x9c, 0x63, 0xf7, 0xad, 0x3c, 0xc3, 0x83, 0x38,
	0x3e, 0x97, 0xb7, 0x04, 0xd9, 0x37, 0xd0, 0x90, 0xba, 0xd4, 0x94, 0xa8, 0xa8, 0x9e, 0x91, 0x26,
	0x9c, 0x7d, 0x25, 0x6f, 0xb8, 0xf5, 0x6c, 0xe9, 0xad, 0xdf, 0xe3, 0xf9, 0x1d, 0xc1, 0xc1, 0xda,
	0xfc, 0xce, 0x8a, 0x90, 0xea, 0xcb, 0x54, 0xe6, 0x1e, 0x00, 0xc5, 0xdf, 0x61, 0x9f, 0x71, 0x77,
	0xa2, 0xa3, 0xd5, 0x61, 0xc9, 0xad, 0xc9, 0xb3, 0x41, 0x1c, 0x9b, 0x57, 0xf0, 0x1a, 0xcf, 0xfc,
	0x38, 0xe7, 0x93, 0xe2, 0xc5, 0x51, 0xc6, 0x94, 0xeb, 0x8f, 0x94, 0x5e, 0xde, 0x09, 0x4b, 0x59,
	0x29, 0xd4, 0x7e, 0x19, 0x65, 0x6c, 0x24, 0x86, 0x5e, 0x7d, 0x67, 0x58, 0x72, 0x1b, 0x05, 0x0d,
	0x47, 0x39, 0x6f, 0x94, 0xac, 0xf0, 0xea, 0xff, 0x93, 0xb7, 0xa0, 0xe1, 0x68, 0xff, 0x03, 0x9e,
	0xc6, 0x7b, 0x70, 0xb8, 0x36, 0x0d, 0x75, 0xdb, 0x69, 0xdd, 0xeb, 0x85, 0xa9, 0xff, 0xed, 0x68,
	0x9d, 0x5f, 0x34, 0xa8, 0x2f, 0x8f, 0x82, 0xb9, 0x05, 0x55, 0x86, 0x42, 0x6f, 0x69, 0xe8, 0x5f,
	0x31, 0x14, 0x7e, 0xcd, 0xe7, 0xbe, 0x0d, 0x2f, 0x29, 0x0e, 0xf1, 0x4c, 0xa4, 0x52, 0x73, 0xaa,
	0xff, 0x38, 0x2f, 0xa9, 0xde, 0xbd, 0xe3, 0x21, 0x4a, 0xc0, 0xdc, 0x07, 0x63, 0x1c, 0xcd, 0x70,
	0xe0, 0x4d, 0xf9, 0x44, 0x08, 0x97, 0xb5, 0x61, 0xc9, 0x05, 0x71, 0x28, 0xa6, 0xe4, 0x59, 0xcd,
	0xcb, 0x62, 0x1c, 0x03, 0x6a, 0x5c, 0x8c, 0x60, 0xeb, 0xfc, 0xa5, 0x01, 0x08, 0x47, 0x81, 0xf8,
	0x33, 0x3f, 0x85, 0x57, 0x6a, 0x2b, 0x09, 0x9d, 0xc6, 0x49, 0x67, 0x75, 0xaa, 0xf8, 0xfe, 0x2a,
	0xc2, 0x1c, 0xc8, 0xca, 0x61, 0xc9, 0x5d, 0x5c, 0x32, 0x7b, 0xd0, 0x64, 0x7e, 0xea, 0xf1, 0x3c,
	0x19, 0xa6, 0xd2, 0x70, 0x59, 0x09, 0x7e, 0xcd, 0xfc, 0xf4, 0x4c, 0x02, 0xc2, 0xf9, 0xdb, 0x50,
	0x49, 0x29, 0x1e, 0x47, 0x33, 0x69, 0xc9, 0x55, 0x6f, 0xfd, 0xf7, 0xb9, 0x99, 0x43, 0xe8, 0xac,
	0x6d, 0x80, 0x90, 0xeb, 0x6c, 0x40, 0x53, 0xc4, 0x1f, 0x78, 0x59, 0x91, 0xb5, 0xe8, 0xc0, 0xaf,
	0x1a, 0x34, 0xce, 0x49, 0xf8, 0x06, 0x9d, 0x3d, 0xa5, 0xf6, 0x43, 0xae, 0xf6, 0x18, 0x8e, 0x9e,
	0x52, 0xbb, 0xa2, 0xc2, 0x79, 0x07, 0x5a, 0x01, 0x09, 0xbd, 0x47, 0x45, 0xcb, 0x75, 0xd4, 0xf9,
	0x16, 0x8c, 0xe1, 0x9c, 0x2f, 0xfc, 0x99, 0xd0, 0xbd, 0x07, 0x46, 0x92, 0x4f, 0xbc, 0xeb, 0xdc,
	0xbf, 0xc5, 0x4c, 0x6a, 0xd7, 0x5d, 0x48, 0xf2, 0x89, 0x23, 0x4f, 0x9e, 0xfd, 0x7b, 0x97, 0xc8,
	0x1c, 0xe7, 0xb7, 0xbb, 0x3f, 0xfe, 0xac, 0x94, 0x9b, 0x3a, 0xf4, 0x22, 0x22, 0xfd, 0xa7, 0x94,
	0xcc, 0xe6, 0x6b, 0x57, 0x87, 0x23, 0x07, 0xe4, 0x82, 0x2f, 0xa7, 0x0b, 0xed, 0xba, 0x22, 0xb6,
	0xd4, 0xe9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x63, 0x1b, 0x15, 0x21, 0x07, 0x00, 0x00,
}