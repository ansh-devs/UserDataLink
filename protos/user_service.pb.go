// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/user_service.proto

package protos

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

// GetUserById Messages definations
type GetUserByIdRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByIdRequest) Reset()         { *m = GetUserByIdRequest{} }
func (m *GetUserByIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByIdRequest) ProtoMessage()    {}
func (*GetUserByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e22b62bbe5b61e2e, []int{0}
}

func (m *GetUserByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIdRequest.Unmarshal(m, b)
}
func (m *GetUserByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIdRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIdRequest.Merge(m, src)
}
func (m *GetUserByIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByIdRequest.Size(m)
}
func (m *GetUserByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIdRequest proto.InternalMessageInfo

type GetUserByIdResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByIdResponse) Reset()         { *m = GetUserByIdResponse{} }
func (m *GetUserByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserByIdResponse) ProtoMessage()    {}
func (*GetUserByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e22b62bbe5b61e2e, []int{1}
}

func (m *GetUserByIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByIdResponse.Unmarshal(m, b)
}
func (m *GetUserByIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByIdResponse.Marshal(b, m, deterministic)
}
func (m *GetUserByIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByIdResponse.Merge(m, src)
}
func (m *GetUserByIdResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserByIdResponse.Size(m)
}
func (m *GetUserByIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByIdResponse proto.InternalMessageInfo

// GetUsersListByIds Messages definations
type GetUsersListByIdsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersListByIdsRequest) Reset()         { *m = GetUsersListByIdsRequest{} }
func (m *GetUsersListByIdsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUsersListByIdsRequest) ProtoMessage()    {}
func (*GetUsersListByIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e22b62bbe5b61e2e, []int{2}
}

func (m *GetUsersListByIdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersListByIdsRequest.Unmarshal(m, b)
}
func (m *GetUsersListByIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersListByIdsRequest.Marshal(b, m, deterministic)
}
func (m *GetUsersListByIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersListByIdsRequest.Merge(m, src)
}
func (m *GetUsersListByIdsRequest) XXX_Size() int {
	return xxx_messageInfo_GetUsersListByIdsRequest.Size(m)
}
func (m *GetUsersListByIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersListByIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersListByIdsRequest proto.InternalMessageInfo

type GetUsersListByIdsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersListByIdsResponse) Reset()         { *m = GetUsersListByIdsResponse{} }
func (m *GetUsersListByIdsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUsersListByIdsResponse) ProtoMessage()    {}
func (*GetUsersListByIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e22b62bbe5b61e2e, []int{3}
}

func (m *GetUsersListByIdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersListByIdsResponse.Unmarshal(m, b)
}
func (m *GetUsersListByIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersListByIdsResponse.Marshal(b, m, deterministic)
}
func (m *GetUsersListByIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersListByIdsResponse.Merge(m, src)
}
func (m *GetUsersListByIdsResponse) XXX_Size() int {
	return xxx_messageInfo_GetUsersListByIdsResponse.Size(m)
}
func (m *GetUsersListByIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersListByIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersListByIdsResponse proto.InternalMessageInfo

// GetUserByCriteria Messages definations
type GetUserByCriteriaRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByCriteriaRequest) Reset()         { *m = GetUserByCriteriaRequest{} }
func (m *GetUserByCriteriaRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserByCriteriaRequest) ProtoMessage()    {}
func (*GetUserByCriteriaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e22b62bbe5b61e2e, []int{4}
}

func (m *GetUserByCriteriaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByCriteriaRequest.Unmarshal(m, b)
}
func (m *GetUserByCriteriaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByCriteriaRequest.Marshal(b, m, deterministic)
}
func (m *GetUserByCriteriaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByCriteriaRequest.Merge(m, src)
}
func (m *GetUserByCriteriaRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserByCriteriaRequest.Size(m)
}
func (m *GetUserByCriteriaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByCriteriaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByCriteriaRequest proto.InternalMessageInfo

type GetUserByCriteriaResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserByCriteriaResponse) Reset()         { *m = GetUserByCriteriaResponse{} }
func (m *GetUserByCriteriaResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserByCriteriaResponse) ProtoMessage()    {}
func (*GetUserByCriteriaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e22b62bbe5b61e2e, []int{5}
}

func (m *GetUserByCriteriaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserByCriteriaResponse.Unmarshal(m, b)
}
func (m *GetUserByCriteriaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserByCriteriaResponse.Marshal(b, m, deterministic)
}
func (m *GetUserByCriteriaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserByCriteriaResponse.Merge(m, src)
}
func (m *GetUserByCriteriaResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserByCriteriaResponse.Size(m)
}
func (m *GetUserByCriteriaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserByCriteriaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserByCriteriaResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetUserByIdRequest)(nil), "protov3.GetUserByIdRequest")
	proto.RegisterType((*GetUserByIdResponse)(nil), "protov3.GetUserByIdResponse")
	proto.RegisterType((*GetUsersListByIdsRequest)(nil), "protov3.GetUsersListByIdsRequest")
	proto.RegisterType((*GetUsersListByIdsResponse)(nil), "protov3.GetUsersListByIdsResponse")
	proto.RegisterType((*GetUserByCriteriaRequest)(nil), "protov3.GetUserByCriteriaRequest")
	proto.RegisterType((*GetUserByCriteriaResponse)(nil), "protov3.GetUserByCriteriaResponse")
}

func init() {
	proto.RegisterFile("protos/user_service.proto", fileDescriptor_e22b62bbe5b61e2e)
}

var fileDescriptor_e22b62bbe5b61e2e = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5,
	0x03, 0x8b, 0x09, 0xb1, 0x83, 0xa9, 0x32, 0x63, 0x25, 0x11, 0x2e, 0x21, 0xf7, 0xd4, 0x92, 0xd0,
	0xe2, 0xd4, 0x22, 0xa7, 0x4a, 0xcf, 0x94, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x51,
	0x2e, 0x61, 0x14, 0xd1, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x25, 0x29, 0x2e, 0x09, 0xa8, 0x70,
	0xb1, 0x4f, 0x66, 0x71, 0x09, 0x48, 0xae, 0x18, 0xa6, 0x45, 0x9a, 0x4b, 0x12, 0x8b, 0x1c, 0x86,
	0x46, 0xa7, 0x4a, 0xe7, 0xa2, 0xcc, 0x92, 0xd4, 0xa2, 0xcc, 0x44, 0x4c, 0x8d, 0xc8, 0x72, 0x10,
	0x8d, 0x46, 0xfd, 0x4c, 0x5c, 0xdc, 0x20, 0xa9, 0x60, 0x88, 0xeb, 0x85, 0x3c, 0xb8, 0xb8, 0x91,
	0x1c, 0x26, 0x24, 0xad, 0x07, 0xf5, 0x87, 0x1e, 0xa6, 0x27, 0xa4, 0x64, 0xb0, 0x4b, 0x42, 0x4c,
	0x16, 0x8a, 0xe2, 0x12, 0xc4, 0x70, 0xaf, 0x90, 0x22, 0xba, 0x16, 0x0c, 0x7f, 0x4a, 0x29, 0xe1,
	0x53, 0x82, 0x61, 0x36, 0xc2, 0x4b, 0x98, 0x66, 0x63, 0x04, 0x05, 0xa6, 0xd9, 0x98, 0x21, 0xe2,
	0xc4, 0x15, 0xc5, 0xa1, 0xa7, 0x0f, 0x89, 0xd8, 0x24, 0x36, 0x30, 0x6d, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0xde, 0xf0, 0xeb, 0xe9, 0x01, 0x00, 0x00,
}
