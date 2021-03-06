// Code generated by protoc-gen-go. DO NOT EDIT.
// source: estimator.proto

/*
Package estimator is a generated protocol buffer package.

It is generated from these files:
	estimator.proto

It has these top-level messages:
	PointsInfo
	Origin
	Destination
	DirectionsInfo
*/
package estimator

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
	Origin      *Origin      `protobuf:"bytes,1,opt,name=origin" json:"origin,omitempty"`
	Destination *Destination `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
}

func (m *PointsInfo) Reset()                    { *m = PointsInfo{} }
func (m *PointsInfo) String() string            { return proto.CompactTextString(m) }
func (*PointsInfo) ProtoMessage()               {}
func (*PointsInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PointsInfo) GetOrigin() *Origin {
	if m != nil {
		return m.Origin
	}
	return nil
}

func (m *PointsInfo) GetDestination() *Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

type Origin struct {
	Lat  string `protobuf:"bytes,1,opt,name=lat" json:"lat,omitempty"`
	Long string `protobuf:"bytes,2,opt,name=long" json:"long,omitempty"`
}

func (m *Origin) Reset()                    { *m = Origin{} }
func (m *Origin) String() string            { return proto.CompactTextString(m) }
func (*Origin) ProtoMessage()               {}
func (*Origin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Origin) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Origin) GetLong() string {
	if m != nil {
		return m.Long
	}
	return ""
}

type Destination struct {
	Lat  string `protobuf:"bytes,1,opt,name=lat" json:"lat,omitempty"`
	Long string `protobuf:"bytes,2,opt,name=long" json:"long,omitempty"`
}

func (m *Destination) Reset()                    { *m = Destination{} }
func (m *Destination) String() string            { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()               {}
func (*Destination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Destination) GetLat() string {
	if m != nil {
		return m.Lat
	}
	return ""
}

func (m *Destination) GetLong() string {
	if m != nil {
		return m.Long
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
func (*DirectionsInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	proto.RegisterType((*PointsInfo)(nil), "estimator.PointsInfo")
	proto.RegisterType((*Origin)(nil), "estimator.Origin")
	proto.RegisterType((*Destination)(nil), "estimator.Destination")
	proto.RegisterType((*DirectionsInfo)(nil), "estimator.DirectionsInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Estimate service

type EstimateClient interface {
	TimeAndDistanceBetweenPoints(ctx context.Context, in *PointsInfo, opts ...grpc.CallOption) (*DirectionsInfo, error)
}

type estimateClient struct {
	cc *grpc.ClientConn
}

func NewEstimateClient(cc *grpc.ClientConn) EstimateClient {
	return &estimateClient{cc}
}

func (c *estimateClient) TimeAndDistanceBetweenPoints(ctx context.Context, in *PointsInfo, opts ...grpc.CallOption) (*DirectionsInfo, error) {
	out := new(DirectionsInfo)
	err := grpc.Invoke(ctx, "/estimator.Estimate/timeAndDistanceBetweenPoints", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Estimate service

type EstimateServer interface {
	TimeAndDistanceBetweenPoints(context.Context, *PointsInfo) (*DirectionsInfo, error)
}

func RegisterEstimateServer(s *grpc.Server, srv EstimateServer) {
	s.RegisterService(&_Estimate_serviceDesc, srv)
}

func _Estimate_TimeAndDistanceBetweenPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PointsInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstimateServer).TimeAndDistanceBetweenPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/estimator.Estimate/TimeAndDistanceBetweenPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstimateServer).TimeAndDistanceBetweenPoints(ctx, req.(*PointsInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Estimate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "estimator.Estimate",
	HandlerType: (*EstimateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "timeAndDistanceBetweenPoints",
			Handler:    _Estimate_TimeAndDistanceBetweenPoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "estimator.proto",
}

func init() { proto.RegisterFile("estimator.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xb5, 0xda, 0x86, 0x66, 0x02, 0x7e, 0x0c, 0x28, 0xb5, 0x78, 0x90, 0x9c, 0xf4, 0x92, 0x43,
	0x7b, 0xf1, 0xa8, 0x12, 0x0f, 0x9e, 0x94, 0xfd, 0x01, 0x42, 0x6c, 0xc6, 0x32, 0xd0, 0xce, 0x68,
	0x32, 0xe0, 0xdf, 0x97, 0x4c, 0x4a, 0xb3, 0xde, 0x7a, 0x7b, 0xc9, 0xfb, 0xd8, 0x7d, 0x6f, 0xe1,
	0x8c, 0x5a, 0xe3, 0x6d, 0x65, 0xda, 0x14, 0xdf, 0x8d, 0x9a, 0x62, 0xba, 0xff, 0x91, 0xff, 0x00,
	0xbc, 0x2b, 0x8b, 0xb5, 0xaf, 0xf2, 0xa5, 0x78, 0x0f, 0x89, 0x36, 0xbc, 0x66, 0x99, 0x8d, 0x6e,
	0x47, 0x77, 0xd9, 0xe2, 0xa2, 0x18, 0xac, 0x6f, 0x4e, 0x84, 0x9d, 0x00, 0x1f, 0x20, 0xab, 0x3b,
	0x52, 0x2a, 0x63, 0x95, 0xd9, 0xb1, 0xeb, 0xaf, 0x22, 0x7d, 0x39, 0xb0, 0x21, 0x96, 0xe6, 0x05,
	0x24, 0x7d, 0x16, 0x9e, 0xc3, 0xc9, 0xa6, 0x32, 0x3f, 0x2b, 0x0d, 0x1d, 0x44, 0x84, 0xf1, 0x46,
	0x65, 0xed, 0x71, 0x69, 0x70, 0x9c, 0x2f, 0x21, 0x8b, 0xb2, 0x0e, 0x34, 0x3d, 0xc2, 0x69, 0xc9,
	0x0d, 0xad, 0x3a, 0x4b, 0xdf, 0x0d, 0x61, 0x6c, 0xbc, 0x25, 0x37, 0x4e, 0x82, 0x63, 0x9c, 0xc3,
	0xb4, 0xe6, 0xd6, 0x2a, 0x59, 0x91, 0xbb, 0x27, 0x61, 0xff, 0xbd, 0xf8, 0x80, 0xe9, 0x4b, 0x5f,
	0x86, 0x30, 0xc0, 0x4d, 0xa7, 0x7f, 0x92, 0xba, 0xdc, 0xd1, 0xcf, 0x64, 0xbf, 0x44, 0xd2, 0x6f,
	0x87, 0x97, 0x51, 0xef, 0x61, 0xce, 0xf9, 0x75, 0x3c, 0xc7, 0xbf, 0xdb, 0xe4, 0x47, 0x9f, 0x89,
	0xbf, 0xc5, 0xf2, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x29, 0xd5, 0x82, 0xe6, 0x9e, 0x01, 0x00, 0x00,
}
