// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/aggregate/v2alpha/cluster.proto

package envoy_config_cluster_aggregate_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type ClusterConfig struct {
	Clusters             []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterConfig) Reset()         { *m = ClusterConfig{} }
func (m *ClusterConfig) String() string { return proto.CompactTextString(m) }
func (*ClusterConfig) ProtoMessage()    {}
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7f3700a1612359, []int{0}
}

func (m *ClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterConfig.Unmarshal(m, b)
}
func (m *ClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterConfig.Marshal(b, m, deterministic)
}
func (m *ClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfig.Merge(m, src)
}
func (m *ClusterConfig) XXX_Size() int {
	return xxx_messageInfo_ClusterConfig.Size(m)
}
func (m *ClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfig proto.InternalMessageInfo

func (m *ClusterConfig) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterConfig)(nil), "envoy.config.cluster.aggregate.v2alpha.ClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/aggregate/v2alpha/cluster.proto", fileDescriptor_2c7f3700a1612359)
}

var fileDescriptor_2c7f3700a1612359 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d,
	0xd2, 0x4f, 0x4c, 0x4f, 0x2f, 0x4a, 0x4d, 0x4f, 0x2c, 0x49, 0xd5, 0x2f, 0x33, 0x4a, 0xcc, 0x29,
	0xc8, 0x48, 0x84, 0xc9, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xa9, 0x81, 0x75, 0xe9, 0x41,
	0x74, 0xe9, 0xc1, 0xe4, 0xe0, 0xba, 0xf4, 0xa0, 0xba, 0xa4, 0xe4, 0x4a, 0x53, 0x0a, 0x12, 0xf5,
	0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x73, 0x33, 0xd3, 0x8b,
	0x40, 0x2a, 0xc0, 0xe6, 0x48, 0x89, 0x97, 0x25, 0xe6, 0x64, 0xa6, 0x80, 0xed, 0x81, 0x32, 0x20,
	0x12, 0x4a, 0xa6, 0x5c, 0xbc, 0xce, 0x10, 0x53, 0x9d, 0xc1, 0x76, 0x08, 0xa9, 0x70, 0x71, 0x40,
	0xad, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x74, 0xe2, 0xf8, 0xe5, 0xc4, 0x3a, 0x89, 0x91,
	0x89, 0x83, 0x31, 0x08, 0x2e, 0xe3, 0x54, 0xf2, 0x69, 0xc6, 0xbf, 0x7e, 0x56, 0x0d, 0x98, 0xfb,
	0x52, 0x2b, 0x4a, 0x52, 0xf3, 0x8a, 0x41, 0xf6, 0xc2, 0xdc, 0x58, 0x8c, 0xec, 0x48, 0x63, 0x2e,
	0x93, 0xcc, 0x7c, 0x3d, 0xb0, 0xd2, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0xe2, 0x7c, 0xe5, 0xc4,
	0x03, 0x75, 0x5a, 0x00, 0xc8, 0xa9, 0x01, 0x8c, 0x49, 0x6c, 0x60, 0x37, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x59, 0x51, 0x03, 0x56, 0x4c, 0x01, 0x00, 0x00,
}
