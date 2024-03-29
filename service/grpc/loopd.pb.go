// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/grpc/loopd.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SetupRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetupRequest) Reset()         { *m = SetupRequest{} }
func (m *SetupRequest) String() string { return proto.CompactTextString(m) }
func (*SetupRequest) ProtoMessage()    {}
func (*SetupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73747ef4b3851b8, []int{0}
}

func (m *SetupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupRequest.Unmarshal(m, b)
}
func (m *SetupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupRequest.Marshal(b, m, deterministic)
}
func (m *SetupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupRequest.Merge(m, src)
}
func (m *SetupRequest) XXX_Size() int {
	return xxx_messageInfo_SetupRequest.Size(m)
}
func (m *SetupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetupRequest proto.InternalMessageInfo

func (m *SetupRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type SetupResponse struct {
	DeviceName           string   `protobuf:"bytes,1,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetupResponse) Reset()         { *m = SetupResponse{} }
func (m *SetupResponse) String() string { return proto.CompactTextString(m) }
func (*SetupResponse) ProtoMessage()    {}
func (*SetupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73747ef4b3851b8, []int{1}
}

func (m *SetupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetupResponse.Unmarshal(m, b)
}
func (m *SetupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetupResponse.Marshal(b, m, deterministic)
}
func (m *SetupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetupResponse.Merge(m, src)
}
func (m *SetupResponse) XXX_Size() int {
	return xxx_messageInfo_SetupResponse.Size(m)
}
func (m *SetupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetupResponse proto.InternalMessageInfo

func (m *SetupResponse) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

type TeardownRequest struct {
	DeviceName           string   `protobuf:"bytes,1,opt,name=deviceName,proto3" json:"deviceName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeardownRequest) Reset()         { *m = TeardownRequest{} }
func (m *TeardownRequest) String() string { return proto.CompactTextString(m) }
func (*TeardownRequest) ProtoMessage()    {}
func (*TeardownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73747ef4b3851b8, []int{2}
}

func (m *TeardownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeardownRequest.Unmarshal(m, b)
}
func (m *TeardownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeardownRequest.Marshal(b, m, deterministic)
}
func (m *TeardownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeardownRequest.Merge(m, src)
}
func (m *TeardownRequest) XXX_Size() int {
	return xxx_messageInfo_TeardownRequest.Size(m)
}
func (m *TeardownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeardownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeardownRequest proto.InternalMessageInfo

func (m *TeardownRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b73747ef4b3851b8, []int{3}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SetupRequest)(nil), "SetupRequest")
	proto.RegisterType((*SetupResponse)(nil), "SetupResponse")
	proto.RegisterType((*TeardownRequest)(nil), "TeardownRequest")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
}

func init() { proto.RegisterFile("service/grpc/loopd.proto", fileDescriptor_b73747ef4b3851b8) }

var fileDescriptor_b73747ef4b3851b8 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0xcf, 0xc9, 0xcf, 0x2f, 0x48, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe2, 0xe2, 0x09, 0x4e, 0x2d, 0x29, 0x2d, 0x08, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0x48, 0xcb, 0xcc, 0x49, 0xcd, 0x4b, 0xcc, 0x4d,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf3, 0x95, 0xf4, 0xb9, 0x78, 0xa1, 0x6a, 0x8b,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xe4, 0xb8, 0xb8, 0x52, 0x52, 0x41, 0xe6, 0xfa, 0x21, 0x94,
	0x23, 0x89, 0x28, 0x19, 0x72, 0xf1, 0x87, 0xa4, 0x26, 0x16, 0xa5, 0xe4, 0x97, 0xe7, 0xc1, 0xcc,
	0x27, 0xa4, 0x85, 0x9f, 0x8b, 0xd7, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0x66, 0x87, 0x51, 0x2c, 0x17,
	0xab, 0x0f, 0xc8, 0xbd, 0x42, 0x6a, 0x5c, 0xac, 0x60, 0xdb, 0x85, 0x78, 0xf5, 0x90, 0x5d, 0x2c,
	0xc5, 0xa7, 0x87, 0xea, 0x28, 0x1d, 0x2e, 0x0e, 0x98, 0xa5, 0x42, 0x02, 0x7a, 0x68, 0xf6, 0x4b,
	0xf1, 0xe9, 0xa1, 0x18, 0xef, 0xc4, 0x16, 0xc5, 0x02, 0x0a, 0x93, 0x24, 0x36, 0x70, 0x70, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x06, 0x2a, 0x56, 0x2a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoopdClient is the client API for Loopd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoopdClient interface {
	Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*SetupResponse, error)
	Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type loopdClient struct {
	cc *grpc.ClientConn
}

func NewLoopdClient(cc *grpc.ClientConn) LoopdClient {
	return &loopdClient{cc}
}

func (c *loopdClient) Setup(ctx context.Context, in *SetupRequest, opts ...grpc.CallOption) (*SetupResponse, error) {
	out := new(SetupResponse)
	err := c.cc.Invoke(ctx, "/Loopd/Setup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loopdClient) Teardown(ctx context.Context, in *TeardownRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/Loopd/Teardown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoopdServer is the server API for Loopd service.
type LoopdServer interface {
	Setup(context.Context, *SetupRequest) (*SetupResponse, error)
	Teardown(context.Context, *TeardownRequest) (*EmptyResponse, error)
}

// UnimplementedLoopdServer can be embedded to have forward compatible implementations.
type UnimplementedLoopdServer struct {
}

func (*UnimplementedLoopdServer) Setup(ctx context.Context, req *SetupRequest) (*SetupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Setup not implemented")
}
func (*UnimplementedLoopdServer) Teardown(ctx context.Context, req *TeardownRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Teardown not implemented")
}

func RegisterLoopdServer(s *grpc.Server, srv LoopdServer) {
	s.RegisterService(&_Loopd_serviceDesc, srv)
}

func _Loopd_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoopdServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Loopd/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoopdServer).Setup(ctx, req.(*SetupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loopd_Teardown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeardownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoopdServer).Teardown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Loopd/Teardown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoopdServer).Teardown(ctx, req.(*TeardownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Loopd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Loopd",
	HandlerType: (*LoopdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Setup",
			Handler:    _Loopd_Setup_Handler,
		},
		{
			MethodName: "Teardown",
			Handler:    _Loopd_Teardown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/grpc/loopd.proto",
}
