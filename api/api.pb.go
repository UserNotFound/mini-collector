// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// NOTE: Protocol buffers used variable-length encoding, so even though uint64
// is arguably much bigger than we'd need, it's simpler to just use that. We
// do, however, use signed ints for disk usage and limit because those fields
// are optional (and should not be reported if they are < 0).
type PublishRequest struct {
	UnixTime             int64    `protobuf:"varint,1,opt,name=unix_time,json=unixTime,proto3" json:"unix_time,omitempty"`
	Running              bool     `protobuf:"varint,2,opt,name=running,proto3" json:"running,omitempty"`
	MilliCpuUsage        uint64   `protobuf:"varint,3,opt,name=milli_cpu_usage,json=milliCpuUsage,proto3" json:"milli_cpu_usage,omitempty"`
	MemoryTotalMb        uint64   `protobuf:"varint,4,opt,name=memory_total_mb,json=memoryTotalMb,proto3" json:"memory_total_mb,omitempty"`
	MemoryRssMb          uint64   `protobuf:"varint,5,opt,name=memory_rss_mb,json=memoryRssMb,proto3" json:"memory_rss_mb,omitempty"`
	MemoryLimitMb        uint64   `protobuf:"varint,6,opt,name=memory_limit_mb,json=memoryLimitMb,proto3" json:"memory_limit_mb,omitempty"`
	DiskUsageMb          int64    `protobuf:"zigzag64,7,opt,name=disk_usage_mb,json=diskUsageMb,proto3" json:"disk_usage_mb,omitempty"`
	DiskLimitMb          int64    `protobuf:"zigzag64,8,opt,name=disk_limit_mb,json=diskLimitMb,proto3" json:"disk_limit_mb,omitempty"`
	DiskReadKbps         uint64   `protobuf:"varint,9,opt,name=disk_read_kbps,json=diskReadKbps,proto3" json:"disk_read_kbps,omitempty"`
	DiskWriteKbps        uint64   `protobuf:"varint,10,opt,name=disk_write_kbps,json=diskWriteKbps,proto3" json:"disk_write_kbps,omitempty"`
	DiskReadIops         uint64   `protobuf:"varint,11,opt,name=disk_read_iops,json=diskReadIops,proto3" json:"disk_read_iops,omitempty"`
	DiskWriteIops        uint64   `protobuf:"varint,12,opt,name=disk_write_iops,json=diskWriteIops,proto3" json:"disk_write_iops,omitempty"`
	PidsCurrent          uint64   `protobuf:"varint,13,opt,name=pids_current,json=pidsCurrent,proto3" json:"pids_current,omitempty"`
	PidsLimit            uint64   `protobuf:"varint,14,opt,name=pids_limit,json=pidsLimit,proto3" json:"pids_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d57130a52f246854, []int{0}
}
func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (dst *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(dst, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetUnixTime() int64 {
	if m != nil {
		return m.UnixTime
	}
	return 0
}

func (m *PublishRequest) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

func (m *PublishRequest) GetMilliCpuUsage() uint64 {
	if m != nil {
		return m.MilliCpuUsage
	}
	return 0
}

func (m *PublishRequest) GetMemoryTotalMb() uint64 {
	if m != nil {
		return m.MemoryTotalMb
	}
	return 0
}

func (m *PublishRequest) GetMemoryRssMb() uint64 {
	if m != nil {
		return m.MemoryRssMb
	}
	return 0
}

func (m *PublishRequest) GetMemoryLimitMb() uint64 {
	if m != nil {
		return m.MemoryLimitMb
	}
	return 0
}

func (m *PublishRequest) GetDiskUsageMb() int64 {
	if m != nil {
		return m.DiskUsageMb
	}
	return 0
}

func (m *PublishRequest) GetDiskLimitMb() int64 {
	if m != nil {
		return m.DiskLimitMb
	}
	return 0
}

func (m *PublishRequest) GetDiskReadKbps() uint64 {
	if m != nil {
		return m.DiskReadKbps
	}
	return 0
}

func (m *PublishRequest) GetDiskWriteKbps() uint64 {
	if m != nil {
		return m.DiskWriteKbps
	}
	return 0
}

func (m *PublishRequest) GetDiskReadIops() uint64 {
	if m != nil {
		return m.DiskReadIops
	}
	return 0
}

func (m *PublishRequest) GetDiskWriteIops() uint64 {
	if m != nil {
		return m.DiskWriteIops
	}
	return 0
}

func (m *PublishRequest) GetPidsCurrent() uint64 {
	if m != nil {
		return m.PidsCurrent
	}
	return 0
}

func (m *PublishRequest) GetPidsLimit() uint64 {
	if m != nil {
		return m.PidsLimit
	}
	return 0
}

type PublishResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_d57130a52f246854, []int{1}
}
func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (dst *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(dst, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PublishRequest)(nil), "PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "PublishResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AggregatorClient is the client API for Aggregator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AggregatorClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
}

type aggregatorClient struct {
	cc *grpc.ClientConn
}

func NewAggregatorClient(cc *grpc.ClientConn) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/Aggregator/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregatorServer is the server API for Aggregator service.
type AggregatorServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
}

func RegisterAggregatorServer(s *grpc.Server, srv AggregatorServer) {
	s.RegisterService(&_Aggregator_serviceDesc, srv)
}

func _Aggregator_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Aggregator/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Aggregator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Aggregator_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_d57130a52f246854) }

var fileDescriptor_api_d57130a52f246854 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0x4b, 0xe3, 0x40,
	0x14, 0x80, 0x37, 0xdb, 0x6e, 0xdb, 0xbc, 0xfe, 0xda, 0x9d, 0xd3, 0xb0, 0xcb, 0x42, 0x0c, 0x22,
	0x39, 0xe5, 0xa0, 0x57, 0x2f, 0xd2, 0x93, 0x68, 0x41, 0x42, 0xc5, 0xe3, 0x90, 0x34, 0x43, 0x1c,
	0x9a, 0x64, 0xc6, 0xf9, 0x81, 0xfa, 0x27, 0xf8, 0x5f, 0xcb, 0xbc, 0xb4, 0xb5, 0xad, 0xc7, 0x7c,
	0x7c, 0xf9, 0xf2, 0x5e, 0x78, 0x10, 0xe6, 0x4a, 0xa4, 0x4a, 0x4b, 0x2b, 0xe3, 0x8f, 0x3e, 0xcc,
	0x1e, 0x5c, 0x51, 0x0b, 0xf3, 0x9c, 0xf1, 0x17, 0xc7, 0x8d, 0x25, 0xff, 0x20, 0x74, 0xad, 0x78,
	0x63, 0x56, 0x34, 0x9c, 0x06, 0x51, 0x90, 0xf4, 0xb2, 0x91, 0x07, 0x2b, 0xd1, 0x70, 0x42, 0x61,
	0xa8, 0x5d, 0xdb, 0x8a, 0xb6, 0xa2, 0x3f, 0xa3, 0x20, 0x19, 0x65, 0xbb, 0x47, 0x72, 0x01, 0xf3,
	0x46, 0xd4, 0xb5, 0x60, 0x6b, 0xe5, 0x98, 0x33, 0x79, 0xc5, 0x69, 0x2f, 0x0a, 0x92, 0x7e, 0x36,
	0x45, 0xbc, 0x50, 0xee, 0xd1, 0x43, 0xf4, 0x78, 0x23, 0xf5, 0x3b, 0xb3, 0xd2, 0xe6, 0x35, 0x6b,
	0x0a, 0xda, 0xdf, 0x7a, 0x88, 0x57, 0x9e, 0x2e, 0x0b, 0x12, 0xc3, 0x16, 0x30, 0x6d, 0x8c, 0xb7,
	0x7e, 0xa1, 0x35, 0xee, 0x60, 0x66, 0xcc, 0xb2, 0x38, 0x68, 0xd5, 0xa2, 0x11, 0xd6, 0x5b, 0x83,
	0xc3, 0xd6, 0xbd, 0xa7, 0x5d, 0xab, 0x14, 0x66, 0xd3, 0x8d, 0xe5, 0xad, 0x61, 0x14, 0x24, 0x24,
	0x1b, 0x7b, 0x88, 0x53, 0x1d, 0x38, 0xfb, 0xd2, 0xe8, 0xcb, 0xd9, 0x75, 0xce, 0x61, 0x86, 0x8e,
	0xe6, 0x79, 0xc9, 0x36, 0x85, 0x32, 0x34, 0xc4, 0xcf, 0x4d, 0x3c, 0xcd, 0x78, 0x5e, 0xde, 0x15,
	0xca, 0xf8, 0xa9, 0xd0, 0x7a, 0xd5, 0xc2, 0xf2, 0x4e, 0x83, 0x6e, 0x2a, 0x8f, 0x9f, 0x3c, 0x45,
	0xef, 0xa8, 0x26, 0xa4, 0x32, 0x74, 0x7c, 0x5c, 0xbb, 0x95, 0xdf, 0x6a, 0xa8, 0x4d, 0x4e, 0x6a,
	0xe8, 0x9d, 0xc1, 0x44, 0x89, 0xd2, 0xb0, 0xb5, 0xd3, 0x9a, 0xb7, 0x96, 0x4e, 0xbb, 0xdf, 0xe5,
	0xd9, 0xa2, 0x43, 0xe4, 0x3f, 0x00, 0x2a, 0xb8, 0x22, 0x9d, 0xa1, 0x10, 0x7a, 0x82, 0xfb, 0xc5,
	0x7f, 0x60, 0xbe, 0x3f, 0x05, 0xa3, 0x64, 0x6b, 0xf8, 0xe5, 0x35, 0xc0, 0x4d, 0x55, 0x69, 0x5e,
	0xe5, 0x56, 0x6a, 0x92, 0xc2, 0x70, 0x2b, 0x90, 0x79, 0x7a, 0x7c, 0x35, 0x7f, 0x7f, 0xa7, 0x27,
	0xef, 0xc6, 0x3f, 0x8a, 0x01, 0xde, 0xd8, 0xd5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x40,
	0xc6, 0x69, 0x70, 0x02, 0x00, 0x00,
}
