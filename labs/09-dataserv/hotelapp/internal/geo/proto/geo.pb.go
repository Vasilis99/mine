// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mine/labs/09-dataserv/hotelapp/internal/geo/proto/geo.proto

/*
Package geo is a generated protocol buffer package.

It is generated from these files:
        mine/labs/09-dataserv/hotelapp/internal/geo/proto/geo.proto

It has these top-level messages:
        Request
        Result
*/
package geo

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

type Request struct {
        Lat float32 `protobuf:"fixed32,1,opt,name=Lat,json=lat" json:"Lat,omitempty"`
        Lon float32 `protobuf:"fixed32,2,opt,name=Lon,json=lon" json:"Lon,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetLat() float32 {
        if m != nil {
                return m.Lat
        }
        return 0
}

func (m *Request) GetLon() float32 {
        if m != nil {
                return m.Lon
        }
        return 0
}

type Result struct {
        HotelIds []string `protobuf:"bytes,1,rep,name=HotelIds,json=hotelIds" json:"HotelIds,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetHotelIds() []string {
        if m != nil {
                return m.HotelIds
        }
        return nil
}

func init() {
        proto.RegisterType((*Request)(nil), "geo.Request")
        proto.RegisterType((*Result)(nil), "geo.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Geo service

type GeoClient interface {
        Nearby(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type geoClient struct {
        cc *grpc.ClientConn
}

func NewGeoClient(cc *grpc.ClientConn) GeoClient {
        return &geoClient{cc}
}

func (c *geoClient) Nearby(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
        out := new(Result)
        err := grpc.Invoke(ctx, "/geo.Geo/Nearby", in, out, c.cc, opts...)
        if err != nil {
                return nil, err
        }
        return out, nil
}

// Server API for Geo service

type GeoServer interface {
        Nearby(context.Context, *Request) (*Result, error)
}

func RegisterGeoServer(s *grpc.Server, srv GeoServer) {
        s.RegisterService(&_Geo_serviceDesc, srv)
}

func _Geo_Nearby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
        in := new(Request)
        if err := dec(in); err != nil {
                return nil, err
        }
        if interceptor == nil {
                return srv.(GeoServer).Nearby(ctx, in)
        }
        info := &grpc.UnaryServerInfo{
                Server:     srv,
                FullMethod: "/geo.Geo/Nearby",
        }
        handler := func(ctx context.Context, req interface{}) (interface{}, error) {
                return srv.(GeoServer).Nearby(ctx, req.(*Request))
        }
        return interceptor(ctx, in, info, handler)
}

var _Geo_serviceDesc = grpc.ServiceDesc{
        ServiceName: "geo.Geo",
        HandlerType: (*GeoServer)(nil),
        Methods: []grpc.MethodDesc{
                {
                        MethodName: "Nearby",
                        Handler:    _Geo_Nearby_Handler,
                },
        },
        Streams:  []grpc.StreamDesc{},
        Metadata: "mine/labs/09-dataserv/hotelapp/internal/geo/proto/geo.proto",
}

func init() {
        proto.RegisterFile("mine/labs/09-dataserv/hotelapp/internal/geo/proto/geo.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
        // 188 bytes of a gzipped FileDescriptorProto
        0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0xb1, 0x6a, 0xc3, 0x30,
        0x10, 0x86, 0xb1, 0x05, 0xae, 0xab, 0x76, 0x28, 0x9a, 0x8c, 0x27, 0xe3, 0x76, 0x30, 0x05, 0x5b,
        0xa5, 0x9d, 0x4a, 0x1e, 0x20, 0x09, 0x84, 0x0c, 0x7a, 0x83, 0x33, 0x3e, 0x6c, 0x83, 0xa2, 0x73,
        0x24, 0x39, 0x90, 0xb7, 0x0f, 0x72, 0xb4, 0xfd, 0xdf, 0x37, 0xdc, 0x77, 0x7c, 0x77, 0x99, 0x0d,
        0x4a, 0x0d, 0xbd, 0x93, 0x3f, 0xff, 0xed, 0x00, 0x1e, 0x1c, 0xda, 0x9b, 0x9c, 0xc8, 0xa3, 0x86,
        0x65, 0x91, 0xb3, 0xf1, 0x68, 0x0d, 0x68, 0x39, 0x22, 0xc9, 0xc5, 0x92, 0xa7, 0xb0, 0xba, 0x6d,
        0x09, 0x36, 0x22, 0xd5, 0x2d, 0x7f, 0x51, 0x78, 0x5d, 0xd1, 0x79, 0xf1, 0xc1, 0xd9, 0x09, 0x7c,
        0x91, 0x54, 0x49, 0x93, 0x2a, 0xa6, 0xe1, 0x69, 0xc8, 0x14, 0x69, 0x34, 0x64, 0xea, 0x2f, 0x9e,
        0x29, 0x74, 0xab, 0xf6, 0xa2, 0xe4, 0xf9, 0x21, 0x64, 0x8e, 0x83, 0x2b, 0x92, 0x8a, 0x35, 0xaf,
        0x2a, 0x9f, 0x22, 0xff, 0x7e, 0x73, 0xb6, 0x47, 0x12, 0x9f, 0x3c, 0x3b, 0x23, 0xd8, 0xfe, 0x2e,
        0xde, 0xbb, 0x90, 0x8d, 0xa1, 0xf2, 0x2d, 0x52, 0xb8, 0xd3, 0x67, 0xdb, 0x33, 0x7f, 0x8f, 0x00,
        0x00, 0x00, 0xff, 0xff, 0x44, 0x8d, 0x78, 0x12, 0xcb, 0x00, 0x00, 0x00,
}
