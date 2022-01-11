// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/weather.proto

package weather

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetWeatherRequest struct {
	City                 string   `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWeatherRequest) Reset()         { *m = GetWeatherRequest{} }
func (m *GetWeatherRequest) String() string { return proto.CompactTextString(m) }
func (*GetWeatherRequest) ProtoMessage()    {}
func (*GetWeatherRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecdc6c2fb8d35c9c, []int{0}
}

func (m *GetWeatherRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWeatherRequest.Unmarshal(m, b)
}
func (m *GetWeatherRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWeatherRequest.Marshal(b, m, deterministic)
}
func (m *GetWeatherRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWeatherRequest.Merge(m, src)
}
func (m *GetWeatherRequest) XXX_Size() int {
	return xxx_messageInfo_GetWeatherRequest.Size(m)
}
func (m *GetWeatherRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWeatherRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetWeatherRequest proto.InternalMessageInfo

func (m *GetWeatherRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type GetWeatherReply struct {
	Tempreture           int32    `protobuf:"varint,1,opt,name=tempreture,proto3" json:"tempreture,omitempty"`
	ClothRecomendation   string   `protobuf:"bytes,2,opt,name=cloth_recomendation,json=clothRecomendation,proto3" json:"cloth_recomendation,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetWeatherReply) Reset()         { *m = GetWeatherReply{} }
func (m *GetWeatherReply) String() string { return proto.CompactTextString(m) }
func (*GetWeatherReply) ProtoMessage()    {}
func (*GetWeatherReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecdc6c2fb8d35c9c, []int{1}
}

func (m *GetWeatherReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetWeatherReply.Unmarshal(m, b)
}
func (m *GetWeatherReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetWeatherReply.Marshal(b, m, deterministic)
}
func (m *GetWeatherReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetWeatherReply.Merge(m, src)
}
func (m *GetWeatherReply) XXX_Size() int {
	return xxx_messageInfo_GetWeatherReply.Size(m)
}
func (m *GetWeatherReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetWeatherReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetWeatherReply proto.InternalMessageInfo

func (m *GetWeatherReply) GetTempreture() int32 {
	if m != nil {
		return m.Tempreture
	}
	return 0
}

func (m *GetWeatherReply) GetClothRecomendation() string {
	if m != nil {
		return m.ClothRecomendation
	}
	return ""
}

func (m *GetWeatherReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWeatherRequest)(nil), "weather.GetWeatherRequest")
	proto.RegisterType((*GetWeatherReply)(nil), "weather.GetWeatherReply")
}

func init() {
	proto.RegisterFile("proto/weather.proto", fileDescriptor_ecdc6c2fb8d35c9c)
}

var fileDescriptor_ecdc6c2fb8d35c9c = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0xd5, 0x16, 0xe7, 0x52, 0x3b, 0x05, 0x0d, 0x41, 0x44, 0x72, 0xd1, 0x53, 0x17,
	0xf5, 0x0d, 0xbc, 0x78, 0xcf, 0x45, 0xd0, 0x83, 0xac, 0x71, 0x4c, 0x03, 0xe9, 0xce, 0xba, 0x3b,
	0xa9, 0xc4, 0xa3, 0xaf, 0xe0, 0xa3, 0xf9, 0x0a, 0x3e, 0x88, 0x30, 0x4d, 0x30, 0x60, 0x6f, 0xfb,
	0xcf, 0xff, 0xb1, 0x33, 0xff, 0x0f, 0x0b, 0x1f, 0x58, 0xd8, 0xbc, 0x93, 0x95, 0x15, 0x85, 0xa5,
	0x2a, 0x9c, 0x76, 0x32, 0x3d, 0x2d, 0x99, 0xcb, 0x9a, 0x8c, 0xf5, 0x95, 0xb1, 0xce, 0xb1, 0x58,
	0xa9, 0xd8, 0xc5, 0x2d, 0x96, 0x5d, 0xc0, 0xfc, 0x8e, 0xe4, 0x7e, 0xcb, 0xe6, 0xf4, 0xd6, 0x50,
	0x14, 0x44, 0xd8, 0x2f, 0x2a, 0x69, 0x93, 0xd1, 0xf9, 0xe8, 0xf2, 0x30, 0xd7, 0x77, 0xf6, 0x01,
	0xb3, 0x21, 0xe8, 0xeb, 0x16, 0xcf, 0x00, 0x84, 0xd6, 0x3e, 0x90, 0x34, 0x81, 0x14, 0x3e, 0xc8,
	0x07, 0x13, 0x34, 0xb0, 0x28, 0x6a, 0x96, 0xd5, 0x53, 0xa0, 0x82, 0xd7, 0xe4, 0x5e, 0x74, 0x73,
	0x32, 0xd6, 0x5f, 0x51, 0xad, 0x7c, 0xe8, 0xe0, 0x31, 0x4c, 0xa2, 0x58, 0x69, 0x62, 0xb2, 0xa7,
	0x4c, 0xa7, 0xae, 0x5f, 0x61, 0xda, 0x2d, 0xc6, 0x47, 0x80, 0xbf, 0x33, 0x30, 0x5d, 0xf6, 0xa1,
	0xff, 0x85, 0x48, 0x93, 0x9d, 0x9e, 0xaf, 0xdb, 0xec, 0xe4, 0xf3, 0xfb, 0xe7, 0x6b, 0x3c, 0xc7,
	0x99, 0xd9, 0x5c, 0xf5, 0xad, 0x99, 0x92, 0xe4, 0x16, 0x1f, 0x8e, 0x22, 0x85, 0x4d, 0x55, 0x50,
	0xec, 0xe7, 0xcf, 0x13, 0xed, 0xe9, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x9d, 0xe7, 0x28,
	0x65, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WeatherClient is the client API for Weather service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WeatherClient interface {
	GetWeather(ctx context.Context, in *GetWeatherRequest, opts ...grpc.CallOption) (*GetWeatherReply, error)
}

type weatherClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherClient(cc grpc.ClientConnInterface) WeatherClient {
	return &weatherClient{cc}
}

func (c *weatherClient) GetWeather(ctx context.Context, in *GetWeatherRequest, opts ...grpc.CallOption) (*GetWeatherReply, error) {
	out := new(GetWeatherReply)
	err := c.cc.Invoke(ctx, "/weather.Weather/GetWeather", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeatherServer is the server API for Weather service.
type WeatherServer interface {
	GetWeather(context.Context, *GetWeatherRequest) (*GetWeatherReply, error)
}

// UnimplementedWeatherServer can be embedded to have forward compatible implementations.
type UnimplementedWeatherServer struct {
}

func (*UnimplementedWeatherServer) GetWeather(ctx context.Context, req *GetWeatherRequest) (*GetWeatherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeather not implemented")
}

func RegisterWeatherServer(s *grpc.Server, srv WeatherServer) {
	s.RegisterService(&_Weather_serviceDesc, srv)
}

func _Weather_GetWeather_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeatherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeatherServer).GetWeather(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weather.Weather/GetWeather",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeatherServer).GetWeather(ctx, req.(*GetWeatherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Weather_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weather.Weather",
	HandlerType: (*WeatherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWeather",
			Handler:    _Weather_GetWeather_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/weather.proto",
}
