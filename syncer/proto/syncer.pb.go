// Code generated by protoc-gen-go. DO NOT EDIT.
// source: syncer.proto

/*
Package proto is a generated protocol buffer package.

import "../../server/core/proto/services.proto";

It is generated from these files:
	syncer.proto

It has these top-level messages:
	PullRequest
	SyncService
	SyncData
	MappingEntry
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto2 "github.com/apache/servicecomb-service-center/server/core/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type PullRequest struct {
	ServiceName string `protobuf:"bytes,1,opt,name=serviceName" json:"serviceName,omitempty"`
	Options     string `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
	Time        string `protobuf:"bytes,3,opt,name=time" json:"time,omitempty"`
}

func (m *PullRequest) Reset()                    { *m = PullRequest{} }
func (m *PullRequest) String() string            { return proto1.CompactTextString(m) }
func (*PullRequest) ProtoMessage()               {}
func (*PullRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PullRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *PullRequest) GetOptions() string {
	if m != nil {
		return m.Options
	}
	return ""
}

func (m *PullRequest) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type SyncService struct {
	DomainProject string                         `protobuf:"bytes,1,opt,name=domainProject" json:"domainProject,omitempty"`
	Service       *proto2.MicroService           `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Instances     []*proto2.MicroServiceInstance `protobuf:"bytes,3,rep,name=instances" json:"instances,omitempty"`
}

func (m *SyncService) Reset()                    { *m = SyncService{} }
func (m *SyncService) String() string            { return proto1.CompactTextString(m) }
func (*SyncService) ProtoMessage()               {}
func (*SyncService) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SyncService) GetDomainProject() string {
	if m != nil {
		return m.DomainProject
	}
	return ""
}

func (m *SyncService) GetService() *proto2.MicroService {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *SyncService) GetInstances() []*proto2.MicroServiceInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type SyncData struct {
	Services []*SyncService `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *SyncData) Reset()                    { *m = SyncData{} }
func (m *SyncData) String() string            { return proto1.CompactTextString(m) }
func (*SyncData) ProtoMessage()               {}
func (*SyncData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SyncData) GetServices() []*SyncService {
	if m != nil {
		return m.Services
	}
	return nil
}

type MappingEntry struct {
	ClusterName   string `protobuf:"bytes,1,opt,name=clusterName" json:"clusterName,omitempty"`
	DomainProject string `protobuf:"bytes,2,opt,name=domainProject" json:"domainProject,omitempty"`
	OrgServiceID  string `protobuf:"bytes,3,opt,name=orgServiceID" json:"orgServiceID,omitempty"`
	OrgInstanceID string `protobuf:"bytes,4,opt,name=orgInstanceID" json:"orgInstanceID,omitempty"`
	CurServiceID  string `protobuf:"bytes,5,opt,name=curServiceID" json:"curServiceID,omitempty"`
	CurInstanceID string `protobuf:"bytes,6,opt,name=curInstanceID" json:"curInstanceID,omitempty"`
}

func (m *MappingEntry) Reset()                    { *m = MappingEntry{} }
func (m *MappingEntry) String() string            { return proto1.CompactTextString(m) }
func (*MappingEntry) ProtoMessage()               {}
func (*MappingEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MappingEntry) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *MappingEntry) GetDomainProject() string {
	if m != nil {
		return m.DomainProject
	}
	return ""
}

func (m *MappingEntry) GetOrgServiceID() string {
	if m != nil {
		return m.OrgServiceID
	}
	return ""
}

func (m *MappingEntry) GetOrgInstanceID() string {
	if m != nil {
		return m.OrgInstanceID
	}
	return ""
}

func (m *MappingEntry) GetCurServiceID() string {
	if m != nil {
		return m.CurServiceID
	}
	return ""
}

func (m *MappingEntry) GetCurInstanceID() string {
	if m != nil {
		return m.CurInstanceID
	}
	return ""
}

func init() {
	proto1.RegisterType((*PullRequest)(nil), "proto.PullRequest")
	proto1.RegisterType((*SyncService)(nil), "proto.SyncService")
	proto1.RegisterType((*SyncData)(nil), "proto.SyncData")
	proto1.RegisterType((*MappingEntry)(nil), "proto.MappingEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sync service

type SyncClient interface {
	Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*SyncData, error)
}

type syncClient struct {
	cc *grpc.ClientConn
}

func NewSyncClient(cc *grpc.ClientConn) SyncClient {
	return &syncClient{cc}
}

func (c *syncClient) Pull(ctx context.Context, in *PullRequest, opts ...grpc.CallOption) (*SyncData, error) {
	out := new(SyncData)
	err := grpc.Invoke(ctx, "/proto.Sync/Pull", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sync service

type SyncServer interface {
	Pull(context.Context, *PullRequest) (*SyncData, error)
}

func RegisterSyncServer(s *grpc.Server, srv SyncServer) {
	s.RegisterService(&_Sync_serviceDesc, srv)
}

func _Sync_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Sync/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServer).Pull(ctx, req.(*PullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Sync",
	HandlerType: (*SyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pull",
			Handler:    _Sync_Pull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "syncer.proto",
}

func init() { proto1.RegisterFile("syncer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x49, 0x9b, 0xfe, 0xbb, 0x14, 0x90, 0x8e, 0xc5, 0x2a, 0x4b, 0x15, 0x31, 0x74, 0x69,
	0x87, 0x22, 0x06, 0x98, 0xcb, 0xd0, 0xa1, 0xa8, 0x4a, 0x67, 0x86, 0x60, 0xac, 0xca, 0xa8, 0xb5,
	0x83, 0xed, 0x20, 0xe5, 0x61, 0x78, 0x41, 0x9e, 0x02, 0xd9, 0x71, 0xa8, 0x23, 0x3a, 0x25, 0xf9,
	0x72, 0xfe, 0xee, 0xee, 0x27, 0xc3, 0x58, 0x57, 0x82, 0x32, 0xb5, 0x28, 0x94, 0x34, 0x12, 0x7b,
	0xee, 0x31, 0xb9, 0xd2, 0x4c, 0x7d, 0x71, 0xca, 0x74, 0x8d, 0xd3, 0x57, 0x48, 0xb6, 0xe5, 0xe1,
	0x90, 0xb1, 0xcf, 0x92, 0x69, 0x83, 0x53, 0x48, 0x7c, 0xc1, 0x4b, 0x7e, 0x64, 0x24, 0x9a, 0x46,
	0xb3, 0x51, 0x16, 0x22, 0x24, 0x30, 0x90, 0x85, 0xe1, 0x52, 0x68, 0xd2, 0x71, 0x7f, 0x9b, 0x4f,
	0x44, 0x88, 0x0d, 0x3f, 0x32, 0xd2, 0x75, 0xd8, 0xbd, 0xa7, 0xdf, 0x11, 0x24, 0xbb, 0x4a, 0xd0,
	0x5d, 0x6d, 0xc0, 0x3b, 0xb8, 0x7c, 0x97, 0xc7, 0x9c, 0x8b, 0xad, 0x92, 0x1f, 0x8c, 0x1a, 0xdf,
	0xa1, 0x0d, 0x71, 0x0e, 0x03, 0xdf, 0xd2, 0xf5, 0x48, 0x96, 0x37, 0xf5, 0xb4, 0x8b, 0x0d, 0xa7,
	0x4a, 0x7a, 0x57, 0xd6, 0xd4, 0xe0, 0x23, 0x8c, 0xb8, 0xd0, 0x26, 0x17, 0x94, 0x69, 0xd2, 0x9d,
	0x76, 0x67, 0xc9, 0xf2, 0xf6, 0xcc, 0x81, 0xb5, 0xaf, 0xc9, 0x4e, 0xd5, 0xe9, 0x13, 0x0c, 0xed,
	0x78, 0xab, 0xdc, 0xe4, 0xb8, 0x80, 0x61, 0x13, 0x0e, 0x89, 0x9c, 0x05, 0xbd, 0x25, 0xd8, 0x20,
	0xfb, 0xab, 0x49, 0x7f, 0x22, 0x18, 0x6f, 0xf2, 0xa2, 0xe0, 0x62, 0xff, 0x2c, 0x8c, 0xaa, 0x6c,
	0x78, 0xf4, 0x50, 0x6a, 0xc3, 0x54, 0x18, 0x5e, 0x80, 0xfe, 0xaf, 0xdf, 0x39, 0xb7, 0x7e, 0x0a,
	0x63, 0xa9, 0xf6, 0xcd, 0xd4, 0x2b, 0x1f, 0x68, 0x8b, 0x59, 0x93, 0x54, 0xfb, 0x66, 0xa5, 0xf5,
	0x8a, 0xc4, 0xb5, 0xa9, 0x05, 0xad, 0x89, 0x96, 0xea, 0x64, 0xea, 0xd5, 0xa6, 0x90, 0x59, 0x13,
	0x2d, 0x55, 0x60, 0xea, 0xd7, 0xa6, 0x16, 0x5c, 0x3e, 0x40, 0x6c, 0x53, 0xc0, 0x39, 0xc4, 0xf6,
	0xbe, 0x60, 0x13, 0x4d, 0x70, 0x79, 0x26, 0xd7, 0x41, 0x5c, 0x36, 0xd1, 0xf4, 0xe2, 0xad, 0xef,
	0xc8, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x2b, 0x2b, 0x6b, 0x8c, 0x02, 0x00, 0x00,
}