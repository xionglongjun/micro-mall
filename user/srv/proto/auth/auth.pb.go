// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package go_micro_srv_user

import (
	fmt "fmt"
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

type NameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRequest) Reset()         { *m = NameRequest{} }
func (m *NameRequest) String() string { return proto.CompactTextString(m) }
func (*NameRequest) ProtoMessage()    {}
func (*NameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *NameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRequest.Unmarshal(m, b)
}
func (m *NameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRequest.Marshal(b, m, deterministic)
}
func (m *NameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRequest.Merge(m, src)
}
func (m *NameRequest) XXX_Size() int {
	return xxx_messageInfo_NameRequest.Size(m)
}
func (m *NameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NameRequest proto.InternalMessageInfo

func (m *NameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type MobileRequest struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileRequest) Reset()         { *m = MobileRequest{} }
func (m *MobileRequest) String() string { return proto.CompactTextString(m) }
func (*MobileRequest) ProtoMessage()    {}
func (*MobileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *MobileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileRequest.Unmarshal(m, b)
}
func (m *MobileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileRequest.Marshal(b, m, deterministic)
}
func (m *MobileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileRequest.Merge(m, src)
}
func (m *MobileRequest) XXX_Size() int {
	return xxx_messageInfo_MobileRequest.Size(m)
}
func (m *MobileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MobileRequest proto.InternalMessageInfo

func (m *MobileRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *MobileRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type AuthResponse struct {
	Header               string   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ExpiresAt            string   `protobuf:"bytes,3,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResponse) GetExpiresAt() string {
	if m != nil {
		return m.ExpiresAt
	}
	return ""
}

type RegisterRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *RegisterRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type RegisterResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{4}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Error                []*Error `protobuf:"bytes,3,rep,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{5}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetError() []*Error {
	if m != nil {
		return m.Error
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
	return fileDescriptor_82b5829f48cfb8e5, []int{6}
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
	proto.RegisterType((*NameRequest)(nil), "go.micro.srv.user.NameRequest")
	proto.RegisterType((*MobileRequest)(nil), "go.micro.srv.user.MobileRequest")
	proto.RegisterType((*AuthResponse)(nil), "go.micro.srv.user.AuthResponse")
	proto.RegisterType((*RegisterRequest)(nil), "go.micro.srv.user.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "go.micro.srv.user.RegisterResponse")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.user.Error")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x3d, 0x4f, 0xe3, 0x40,
	0x10, 0x4d, 0x62, 0x3b, 0xe7, 0x4c, 0x2e, 0xba, 0xbb, 0x55, 0xee, 0x64, 0x59, 0x27, 0xb0, 0x96,
	0x26, 0x05, 0x32, 0x52, 0x28, 0x51, 0x8a, 0x14, 0x29, 0x28, 0x42, 0x61, 0x01, 0x05, 0x9d, 0x63,
	0x8f, 0x12, 0x8b, 0x38, 0x6b, 0x76, 0xd7, 0x81, 0x5f, 0xc7, 0x6f, 0x43, 0xde, 0xb5, 0x1d, 0x03,
	0x06, 0x21, 0x1a, 0x6b, 0xdf, 0x7c, 0xbc, 0x99, 0x37, 0x33, 0x86, 0xbf, 0x19, 0x67, 0x92, 0x9d,
	0x85, 0xb9, 0xdc, 0xa8, 0x8f, 0xaf, 0x30, 0xf9, 0xb3, 0x66, 0x7e, 0x9a, 0x44, 0x9c, 0xf9, 0x82,
	0xef, 0xfd, 0x5c, 0x20, 0xa7, 0x33, 0x18, 0x5e, 0x85, 0x29, 0x06, 0xf8, 0x90, 0xa3, 0x90, 0x84,
	0x80, 0xb9, 0x0b, 0x53, 0x74, 0xba, 0x5e, 0x77, 0x32, 0x08, 0xd4, 0x9b, 0xb8, 0x60, 0x67, 0xa1,
	0x10, 0x8f, 0x8c, 0xc7, 0x4e, 0x4f, 0xd9, 0x6b, 0x4c, 0x2f, 0x60, 0xb4, 0x64, 0xab, 0x64, 0x5b,
	0x13, 0xfc, 0x83, 0x7e, 0xaa, 0x0c, 0x25, 0x45, 0x89, 0x0a, 0xe2, 0x88, 0xc5, 0x58, 0x12, 0xa8,
	0x37, 0xbd, 0x83, 0x9f, 0xf3, 0x5c, 0x6e, 0x02, 0x14, 0x19, 0xdb, 0x09, 0x2c, 0x72, 0x37, 0x18,
	0xc6, 0xc8, 0xab, 0x5c, 0x8d, 0xc8, 0x18, 0x2c, 0xc9, 0xee, 0x71, 0x57, 0x26, 0x6b, 0x40, 0xfe,
	0xc3, 0x00, 0x9f, 0xb2, 0x84, 0xa3, 0x98, 0x4b, 0xc7, 0x50, 0x9e, 0x83, 0x81, 0xa6, 0xf0, 0x2b,
	0xc0, 0x75, 0x22, 0x24, 0xf2, 0x6f, 0x6a, 0x6b, 0x48, 0x31, 0x5a, 0xa5, 0x98, 0x0d, 0x29, 0xa7,
	0xf0, 0xfb, 0x50, 0xae, 0x94, 0xe3, 0xc0, 0x0f, 0x91, 0x47, 0x11, 0x0a, 0xa1, 0x4a, 0xda, 0x41,
	0x05, 0x69, 0x04, 0xd6, 0xb5, 0xd2, 0x50, 0x2b, 0xeb, 0x36, 0x95, 0x8d, 0xc1, 0xda, 0x87, 0xdb,
	0x44, 0x77, 0x64, 0x07, 0x1a, 0x10, 0x1f, 0x2c, 0xe4, 0x9c, 0x71, 0xc7, 0xf0, 0x8c, 0xc9, 0x70,
	0xea, 0xf8, 0xef, 0x96, 0xe9, 0x2f, 0x0a, 0x7f, 0xa0, 0xc3, 0xe8, 0x0c, 0x2c, 0x85, 0xeb, 0x7e,
	0x8b, 0x1a, 0x96, 0xee, 0x97, 0x78, 0x30, 0x8c, 0x51, 0x44, 0x3c, 0xc9, 0x64, 0xc2, 0xaa, 0xc1,
	0x36, 0x4d, 0xd3, 0xe7, 0x1e, 0x98, 0xc5, 0x76, 0xc8, 0x25, 0x98, 0xc5, 0x85, 0x90, 0xa3, 0x96,
	0x82, 0x8d, 0xd3, 0x71, 0x8f, 0x5b, 0xfc, 0xcd, 0xf5, 0xd2, 0x0e, 0x59, 0x42, 0x5f, 0x5f, 0x0b,
	0xf1, 0x5a, 0x82, 0x5f, 0x1d, 0xd2, 0x57, 0xe8, 0x6e, 0xc0, 0xae, 0x86, 0x4e, 0x68, 0x4b, 0xf8,
	0x9b, 0x03, 0x70, 0x4f, 0x3e, 0x8d, 0xa9, 0x69, 0x17, 0x30, 0xba, 0x2d, 0x26, 0x1e, 0x4a, 0xd4,
	0x5b, 0x6a, 0x1b, 0xb5, 0xf2, 0xb8, 0x1f, 0x7a, 0x68, 0x67, 0xd5, 0x57, 0xff, 0xdc, 0xf9, 0x4b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xde, 0xb7, 0xb9, 0x8c, 0x03, 0x00, 0x00,
}