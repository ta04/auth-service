// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Auth2 struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	R                    string   `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
	C                    string   `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth2) Reset()         { *m = Auth2{} }
func (m *Auth2) String() string { return proto.CompactTextString(m) }
func (*Auth2) ProtoMessage()    {}
func (*Auth2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
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

func (m *Auth2) GetC() string {
	if m != nil {
		return m.C
	}
	return ""
}

type Auth1 struct {
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	T                    string   `protobuf:"bytes,3,opt,name=t,proto3" json:"t,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth1) Reset()         { *m = Auth1{} }
func (m *Auth1) String() string { return proto.CompactTextString(m) }
func (*Auth1) ProtoMessage()    {}
func (*Auth1) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
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

type C struct {
	C                    string   `protobuf:"bytes,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C) Reset()         { *m = C{} }
func (m *C) String() string { return proto.CompactTextString(m) }
func (*C) ProtoMessage()    {}
func (*C) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C.Unmarshal(m, b)
}
func (m *C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C.Marshal(b, m, deterministic)
}
func (m *C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C.Merge(m, src)
}
func (m *C) XXX_Size() int {
	return xxx_messageInfo_C.Size(m)
}
func (m *C) XXX_DiscardUnknown() {
	xxx_messageInfo_C.DiscardUnknown(m)
}

var xxx_messageInfo_C proto.InternalMessageInfo

func (m *C) GetC() string {
	if m != nil {
		return m.C
	}
	return ""
}

type Response struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
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

func (m *Response) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
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

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "com.ta04.srv.auth.Request")
	proto.RegisterType((*Auth2)(nil), "com.ta04.srv.auth.Auth2")
	proto.RegisterType((*Auth1)(nil), "com.ta04.srv.auth.Auth1")
	proto.RegisterType((*C)(nil), "com.ta04.srv.auth.C")
	proto.RegisterType((*Response)(nil), "com.ta04.srv.auth.Response")
	proto.RegisterType((*Error)(nil), "com.ta04.srv.auth.Error")
}

func init() {
	proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874)
}

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4f, 0x84, 0x30,
	0x10, 0x85, 0xb7, 0x8b, 0xac, 0xec, 0xe0, 0x65, 0x9b, 0x3d, 0x10, 0xbc, 0x90, 0x9e, 0xf6, 0xd4,
	0x2c, 0xe8, 0x51, 0x63, 0x0c, 0xf1, 0x6e, 0xea, 0x2f, 0xc0, 0xd2, 0x04, 0xa2, 0x4b, 0xb1, 0x2d,
	0xfc, 0x0e, 0x7f, 0xb2, 0x69, 0x0b, 0x1b, 0xcd, 0xe2, 0x8d, 0xf7, 0xde, 0xcc, 0x47, 0xde, 0x14,
	0xa0, 0x1a, 0x4c, 0x43, 0x7b, 0x25, 0x8d, 0xc4, 0x3b, 0x2e, 0x4f, 0xd4, 0x54, 0xc7, 0x7b, 0xaa,
	0xd5, 0x48, 0x6d, 0x40, 0xb6, 0x70, 0xcd, 0xc4, 0xd7, 0x20, 0xb4, 0x21, 0x4f, 0x10, 0x3e, 0x0f,
	0xa6, 0x29, 0x70, 0x0a, 0xd1, 0xa0, 0x85, 0xea, 0xaa, 0x93, 0x48, 0x50, 0x86, 0x0e, 0x5b, 0x76,
	0xd6, 0xf8, 0x06, 0x90, 0x4a, 0xd6, 0xce, 0x44, 0xca, 0x2a, 0x9e, 0x04, 0x5e, 0x71, 0x92, 0x7b,
	0x40, 0xfe, 0x07, 0xb0, 0xbe, 0x04, 0x98, 0x79, 0xc5, 0x90, 0x1d, 0xa0, 0xd2, 0x53, 0xd0, 0x4c,
	0x69, 0x20, 0x62, 0x42, 0xf7, 0xb2, 0xd3, 0x02, 0xef, 0x21, 0x34, 0xf2, 0x43, 0x74, 0x53, 0xea,
	0x85, 0x75, 0xc7, 0xea, 0xb3, 0xad, 0x1d, 0x3b, 0x62, 0x5e, 0xe0, 0x23, 0x6c, 0x84, 0x52, 0x52,
	0xe9, 0x24, 0xc8, 0x82, 0x43, 0x5c, 0x24, 0xf4, 0xa2, 0x2d, 0x7d, 0xb1, 0x03, 0x6c, 0x9a, 0x23,
	0x8f, 0x10, 0x3a, 0x03, 0x63, 0xb8, 0xe2, 0xb2, 0xf6, 0x65, 0x43, 0xe6, 0xbe, 0x71, 0x06, 0x71,
	0x2d, 0x34, 0x57, 0x6d, 0x6f, 0x5a, 0xd9, 0x4d, 0x35, 0x7e, 0x5b, 0xc5, 0x37, 0x82, 0xd8, 0xf6,
	0x7d, 0x13, 0x6a, 0x6c, 0xb9, 0xc0, 0x0f, 0x10, 0x59, 0xc9, 0x5e, 0xcb, 0x1c, 0x2f, 0xfd, 0xdc,
	0xdd, 0x26, 0xdd, 0x2f, 0x24, 0x25, 0x59, 0xe1, 0xf2, 0xbc, 0x5d, 0xfc, 0xbb, 0x5d, 0xa4, 0xb7,
	0x0b, 0xc9, 0x7c, 0x2d, 0xb2, 0x7a, 0xdf, 0xb8, 0x77, 0xbe, 0xfb, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x98, 0x06, 0x3f, 0x74, 0xf5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthService service

type AuthServiceClient interface {
	AuthRPC1(ctx context.Context, in *Auth1, opts ...client.CallOption) (*C, error)
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

func (c *authServiceClient) AuthRPC1(ctx context.Context, in *Auth1, opts ...client.CallOption) (*C, error) {
	req := c.c.NewRequest(c.serviceName, "AuthService.AuthRPC1", in)
	out := new(C)
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
	AuthRPC1(context.Context, *Auth1, *C) error
	AuthRPC2(context.Context, *Auth2, *Response) error
}

func RegisterAuthServiceHandler(s server.Server, hdlr AuthServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AuthService{hdlr}, opts...))
}

type AuthService struct {
	AuthServiceHandler
}

func (h *AuthService) AuthRPC1(ctx context.Context, in *Auth1, out *C) error {
	return h.AuthServiceHandler.AuthRPC1(ctx, in, out)
}

func (h *AuthService) AuthRPC2(ctx context.Context, in *Auth2, out *Response) error {
	return h.AuthServiceHandler.AuthRPC2(ctx, in, out)
}
