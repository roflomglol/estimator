// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estimation.proto

/*
Package estimation is a generated protocol buffer package.

It is generated from these files:
	estimation.proto

It has these top-level messages:
	PointsInfo
	DirectionsInfo
*/
package estimation

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

type PointsInfo struct {
	Lat1  string `protobuf:"bytes,1,opt,name=lat1" json:"lat1,omitempty"`
	Long1 string `protobuf:"bytes,2,opt,name=long1" json:"long1,omitempty"`
	Lat2  string `protobuf:"bytes,3,opt,name=lat2" json:"lat2,omitempty"`
	Long2 string `protobuf:"bytes,4,opt,name=long2" json:"long2,omitempty"`
}

func (m *PointsInfo) Reset()                    { *m = PointsInfo{} }
func (m *PointsInfo) String() string            { return proto.CompactTextString(m) }
func (*PointsInfo) ProtoMessage()               {}
func (*PointsInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PointsInfo) GetLat1() string {
	if m != nil {
		return m.Lat1
	}
	return ""
}

func (m *PointsInfo) GetLong1() string {
	if m != nil {
		return m.Long1
	}
	return ""
}

func (m *PointsInfo) GetLat2() string {
	if m != nil {
		return m.Lat2
	}
	return ""
}

func (m *PointsInfo) GetLong2() string {
	if m != nil {
		return m.Long2
	}
	return ""
}

type DirectionsInfo struct {
	Time     int32 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Distance int32 `protobuf:"varint,2,opt,name=distance" json:"distance,omitempty"`
}

func (m *DirectionsInfo) Reset()                    { *m = DirectionsInfo{} }
func (m *DirectionsInfo) String() string            { return proto.CompactTextString(m) }
func (*DirectionsInfo) ProtoMessage()               {}
func (*DirectionsInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DirectionsInfo) GetTime() int32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *DirectionsInfo) GetDistance() int32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*PointsInfo)(nil), "estimation.PointsInfo")
	proto.RegisterType((*DirectionsInfo)(nil), "estimation.DirectionsInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Estimation service

type EstimationClient interface {
	TimeAndDistanceBetweenPoints(ctx context.Context, in *PointsInfo, opts ...grpc.CallOption) (*DirectionsInfo, error)
}

type estimationClient struct {
	cc *grpc.ClientConn
}

func NewEstimationClient(cc *grpc.ClientConn) EstimationClient {
	return &estimationClient{cc}
}

func (c *estimationClient) TimeAndDistanceBetweenPoints(ctx context.Context, in *PointsInfo, opts ...grpc.CallOption) (*DirectionsInfo, error) {
	out := new(DirectionsInfo)
	err := grpc.Invoke(ctx, "/estimation.Estimation/TimeAndDistanceBetweenPoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Estimation service

type EstimationServer interface {
	TimeAndDistanceBetweenPoints(context.Context, *PointsInfo) (*DirectionsInfo, error)
}

func RegisterEstimationServer(s *grpc.Server, srv EstimationServer) {
	s.RegisterService(&_Estimation_serviceDesc, srv)
}

func _Estimation_TimeAndDistanceBetweenPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstimationServer).TimeAndDistanceBetweenPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/estimation.Estimation/TimeAndDistanceBetweenPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstimationServer).TimeAndDistanceBetweenPoints(ctx, req.(*PointsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Estimation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "estimation.Estimation",
	HandlerType: (*EstimationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TimeAndDistanceBetweenPoints",
			Handler:    _Estimation_TimeAndDistanceBetweenPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "estimation.proto",
}

func init() { proto.RegisterFile("estimation.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0xcb, 0x82, 0x40,
	0x10, 0xc6, 0x5f, 0xdf, 0x34, 0x6a, 0x0e, 0x11, 0x4b, 0xc4, 0x22, 0x1d, 0xc2, 0x53, 0x27, 0xc1,
	0xed, 0x0b, 0x54, 0xd8, 0xa1, 0x5b, 0x88, 0x1f, 0x20, 0xff, 0x4c, 0xb1, 0xa0, 0xbb, 0xa1, 0x03,
	0x7d, 0xfd, 0x68, 0xb5, 0xd5, 0x6e, 0x33, 0x0f, 0x3f, 0xf8, 0xcd, 0x33, 0xb0, 0xc4, 0x96, 0x64,
	0x9d, 0x91, 0xd4, 0x2a, 0x7c, 0x36, 0x9a, 0x34, 0x83, 0x21, 0x09, 0x6e, 0x00, 0x57, 0x2d, 0x15,
	0xb5, 0x17, 0x75, 0xd7, 0x8c, 0x81, 0x5b, 0x65, 0x14, 0x71, 0x67, 0xeb, 0xec, 0xe6, 0x89, 0x99,
	0xd9, 0x0a, 0xbc, 0x4a, 0xab, 0x47, 0xc4, 0xff, 0x4d, 0xd8, 0x2d, 0x3d, 0x29, 0xf8, 0xc4, 0x92,
	0xe2, 0x4b, 0x0a, 0xee, 0x0e, 0xa4, 0x08, 0x0e, 0xb0, 0x88, 0x65, 0x83, 0xc5, 0x47, 0x67, 0x2d,
	0x24, 0x6b, 0x34, 0x16, 0x2f, 0x31, 0x33, 0xf3, 0x61, 0x56, 0xca, 0x96, 0x32, 0x55, 0xa0, 0x11,
	0x79, 0x89, 0xdd, 0x45, 0x0e, 0x70, 0xb6, 0x17, 0xb3, 0x14, 0x36, 0xa9, 0xac, 0xf1, 0xa8, 0xca,
	0xb8, 0x07, 0x4e, 0x48, 0x2f, 0x44, 0xd5, 0xf5, 0x60, 0xeb, 0x70, 0x54, 0x78, 0xe8, 0xe6, 0xfb,
	0xe3, 0xfc, 0xf7, 0xa2, 0xe0, 0x2f, 0x9f, 0x9a, 0xd7, 0xec, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x2c, 0x2f, 0x59, 0x8b, 0x2e, 0x01, 0x00, 0x00,
}
