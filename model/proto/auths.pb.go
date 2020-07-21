// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auths.proto

package com_ta04_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Auth1 struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	T                    string   `protobuf:"bytes,2,opt,name=t,proto3" json:"t,omitempty"`
	C                    string   `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth1) Reset()         { *m = Auth1{} }
func (m *Auth1) String() string { return proto.CompactTextString(m) }
func (*Auth1) ProtoMessage()    {}
func (*Auth1) Descriptor() ([]byte, []int) {
	return fileDescriptor_697c400aa9441216, []int{0}
}

func (m *Auth1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth1.Unmarshal(m, b)
}
func (m *Auth1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth1.Marshal(b, m, deterministic)
}
func (m *Auth1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth1.Merge(m, src)
}
func (m *Auth1) XXX_Size() int {
	return xxx_messageInfo_Auth1.Size(m)
}
func (m *Auth1) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth1.DiscardUnknown(m)
}

var xxx_messageInfo_Auth1 proto.InternalMessageInfo

func (m *Auth1) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Auth1) GetT() string {
	if m != nil {
		return m.T
	}
	return ""
}

func (m *Auth1) GetC() string {
	if m != nil {
		return m.C
	}
	return ""
}

type Auth2 struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	R                    string   `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth2) Reset()         { *m = Auth2{} }
func (m *Auth2) String() string { return proto.CompactTextString(m) }
func (*Auth2) ProtoMessage()    {}
func (*Auth2) Descriptor() ([]byte, []int) {
	return fileDescriptor_697c400aa9441216, []int{1}
}

func (m *Auth2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth2.Unmarshal(m, b)
}
func (m *Auth2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth2.Marshal(b, m, deterministic)
}
func (m *Auth2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth2.Merge(m, src)
}
func (m *Auth2) XXX_Size() int {
	return xxx_messageInfo_Auth2.Size(m)
}
func (m *Auth2) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth2.DiscardUnknown(m)
}

var xxx_messageInfo_Auth2 proto.InternalMessageInfo

func (m *Auth2) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Auth2) GetR() string {
	if m != nil {
		return m.R
	}
	return ""
}

type Response struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	C                    string   `protobuf:"bytes,2,opt,name=c,proto3" json:"c,omitempty"`
	Error                *Error   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_697c400aa9441216, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Response) GetC() string {
	if m != nil {
		return m.C
	}
	return ""
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_697c400aa9441216, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Auth1)(nil), "com.ta04.srv.auth.Auth1")
	proto.RegisterType((*Auth2)(nil), "com.ta04.srv.auth.Auth2")
	proto.RegisterType((*Response)(nil), "com.ta04.srv.auth.Response")
	proto.RegisterType((*Error)(nil), "com.ta04.srv.auth.Error")
}

func init() {
	proto.RegisterFile("auths.proto", fileDescriptor_697c400aa9441216)
}

var fileDescriptor_697c400aa9441216 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xc4, 0x20,
	0x10, 0x86, 0xad, 0x8a, 0xd6, 0xa9, 0x17, 0x27, 0x1e, 0xc8, 0x7a, 0x31, 0x9c, 0x3c, 0x11, 0x8b,
	0x7a, 0x36, 0x66, 0xe3, 0xdd, 0xe0, 0xdd, 0xa4, 0xe2, 0xc4, 0x35, 0xa6, 0x65, 0x33, 0xd0, 0x7d,
	0x15, 0x5f, 0xd7, 0x00, 0xad, 0x31, 0x71, 0xcd, 0xde, 0xf8, 0x98, 0xe1, 0xfb, 0x19, 0x80, 0xa6,
	0x1b, 0xe3, 0x2a, 0xe8, 0x35, 0xfb, 0xe8, 0xf1, 0xcc, 0xf9, 0x5e, 0xc7, 0xee, 0xfa, 0x56, 0x07,
	0xde, 0xe8, 0x54, 0x51, 0xf7, 0x20, 0x1e, 0xc6, 0xb8, 0x6a, 0x71, 0x01, 0xf5, 0x18, 0x88, 0x87,
	0xae, 0x27, 0x59, 0x5d, 0x56, 0x57, 0x27, 0xf6, 0x87, 0xf1, 0x14, 0xaa, 0x28, 0xf7, 0xf3, 0x66,
	0x15, 0x13, 0x39, 0x79, 0x50, 0xc8, 0xa9, 0xb6, 0x08, 0xcc, 0x2e, 0x01, 0xcf, 0x02, 0x56, 0x2f,
	0x50, 0x5b, 0x0a, 0x6b, 0x3f, 0x04, 0xc2, 0x73, 0x10, 0xd1, 0x7f, 0xd2, 0x30, 0x1d, 0x29, 0x50,
	0x22, 0xa6, 0x7e, 0x87, 0x1a, 0x04, 0x31, 0x7b, 0xce, 0xa1, 0x8d, 0x91, 0xfa, 0xcf, 0x18, 0xfa,
	0x31, 0xd5, 0x6d, 0x69, 0x53, 0x77, 0x20, 0x32, 0x23, 0xc2, 0xa1, 0xf3, 0x6f, 0xe5, 0x3a, 0xc2,
	0xe6, 0x35, 0x4a, 0x38, 0xee, 0x29, 0x84, 0xee, 0x9d, 0xa6, 0x80, 0x19, 0xcd, 0x57, 0x05, 0x4d,
	0x1a, 0xe5, 0x99, 0x78, 0xf3, 0xe1, 0x08, 0x97, 0x50, 0x27, 0xb4, 0x4f, 0xcb, 0x16, 0xb7, 0x65,
	0xe6, 0x77, 0x5b, 0x5c, 0x6c, 0xa9, 0xcc, 0xd3, 0xa9, 0xbd, 0x5f, 0x12, 0xf3, 0xaf, 0xc4, 0xec,
	0x90, 0xbc, 0x1e, 0xe5, 0xef, 0xbb, 0xf9, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x47, 0x2b, 0x54, 0x53,
	0xcd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthService service

type AuthServiceClient interface {
	AuthRPC1(ctx context.Context, in *Auth1, opts ...client.CallOption) (*Response, error)
	AuthRPC2(ctx context.Context, in *Auth2, opts ...client.CallOption) (*Response, error)
}

type authServiceClient struct {
	c           client.Client
	serviceName string
}

func NewAuthServiceClient(serviceName string, c client.Client) AuthServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "com.ta04.srv.auth"
	}
	return &authServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authServiceClient) AuthRPC1(ctx context.Context, in *Auth1, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.AuthRPC1", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthRPC2(ctx context.Context, in *Auth2, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.AuthRPC2", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceHandler interface {
	AuthRPC1(context.Context, *Auth1, *Response) error
	AuthRPC2(context.Context, *Auth2, *Response) error
}

func RegisterAuthServiceHandler(s server.Server, hdlr AuthServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthService{hdlr}, opts...))
}

type AuthService struct {
	AuthServiceHandler
}

func (h *AuthService) AuthRPC1(ctx context.Context, in *Auth1, out *Response) error {
	return h.AuthServiceHandler.AuthRPC1(ctx, in, out)
}

func (h *AuthService) AuthRPC2(ctx context.Context, in *Auth2, out *Response) error {
	return h.AuthServiceHandler.AuthRPC2(ctx, in, out)
}
