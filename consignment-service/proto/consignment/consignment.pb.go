// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package go_micro_srv_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_8e545b56e4300d93, []int{0}
}
func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (dst *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(dst, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_8e545b56e4300d93, []int{1}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// Created a blank get request
type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_8e545b56e4300d93, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created     bool         `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment *Consignment `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	// Added a pluralised consignment to our generic response message
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_8e545b56e4300d93, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_consignment_8e545b56e4300d93)
}

var fileDescriptor_consignment_8e545b56e4300d93 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x49, 0xff, 0x33, 0xa9, 0x40, 0x78, 0x01, 0x16, 0x2c, 0x88, 0x52, 0x21, 0x75, 0x15,
	0xa4, 0x72, 0x84, 0x2c, 0xaa, 0x6c, 0xdd, 0x03, 0x40, 0x49, 0x46, 0xe9, 0x48, 0xc4, 0x0e, 0xb6,
	0x5b, 0xae, 0x06, 0xb7, 0x43, 0x75, 0x1a, 0x30, 0x42, 0x45, 0xdd, 0xf9, 0xcd, 0x9b, 0x37, 0xf9,
	0x32, 0x1a, 0x98, 0x35, 0x5a, 0x59, 0xf5, 0x50, 0x28, 0x69, 0xa8, 0x92, 0x35, 0x4a, 0xeb, 0xbf,
	0x53, 0xe7, 0x32, 0x5e, 0xa9, 0xb4, 0xa6, 0x42, 0xab, 0xd4, 0xe8, 0x5d, 0xea, 0xf9, 0xc9, 0x67,
	0x00, 0x51, 0xf6, 0xa3, 0xd9, 0x39, 0xf4, 0xa8, 0xe4, 0x41, 0x1c, 0xcc, 0x43, 0xd1, 0xa3, 0x92,
	0xc5, 0x10, 0x95, 0x68, 0x0a, 0x4d, 0x8d, 0x25, 0x25, 0x79, 0xcf, 0x19, 0x7e, 0x89, 0x5d, 0xc1,
	0xe8, 0x1d, 0xa9, 0xda, 0x58, 0xde, 0x8f, 0x83, 0xf9, 0x50, 0x1c, 0x14, 0xcb, 0x00, 0x0a, 0x25,
	0xed, 0x9a, 0x24, 0x6a, 0xc3, 0x07, 0x71, 0x7f, 0x1e, 0x2d, 0x66, 0xe9, 0x31, 0x90, 0x34, 0xeb,
	0x7a, 0x85, 0x17, 0x63, 0xb7, 0x10, 0xee, 0xd0, 0x18, 0x7c, 0x7d, 0xa2, 0x92, 0x0f, 0xdd, 0xc7,
	0x27, 0x6d, 0x21, 0x2f, 0x93, 0x1a, 0xc2, 0xef, 0xd4, 0x1f, 0xf0, 0x3b, 0x88, 0x8a, 0xad, 0xb1,
	0xaa, 0x46, 0xbd, 0xcf, 0xb6, 0xe0, 0xd0, 0x95, 0xf2, 0x72, 0xcf, 0xad, 0x34, 0x55, 0x24, 0x1d,
	0x77, 0x28, 0x0e, 0x8a, 0x5d, 0xc3, 0x78, 0x6b, 0xda, 0xd0, 0xa0, 0x35, 0xf6, 0x32, 0x2f, 0x93,
	0x29, 0xc0, 0x12, 0xad, 0xc0, 0xb7, 0x2d, 0x1a, 0x9b, 0x7c, 0x04, 0x30, 0x11, 0x68, 0x1a, 0x25,
	0x0d, 0x32, 0x0e, 0xe3, 0x42, 0xe3, 0xda, 0x62, 0x4b, 0x30, 0x11, 0x9d, 0x64, 0x4b, 0x88, 0xbc,
	0xbf, 0x74, 0x18, 0xd1, 0xe2, 0xfe, 0xdf, 0x35, 0x74, 0x6f, 0xe1, 0x27, 0x59, 0x0e, 0x53, 0x4f,
	0x1a, 0xde, 0x77, 0x0b, 0x3d, 0x71, 0xd2, 0xaf, 0xe8, 0xc2, 0xc0, 0xc5, 0x6a, 0x43, 0x4d, 0x43,
	0xb2, 0x5a, 0xa1, 0xde, 0x51, 0x81, 0xec, 0x19, 0x2e, 0x33, 0x47, 0xec, 0xdf, 0xc2, 0x69, 0xc3,
	0x6f, 0x92, 0xe3, 0x6d, 0xdd, 0x82, 0x92, 0xb3, 0x97, 0x91, 0xbb, 0xc4, 0xc7, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc9, 0xb8, 0xfd, 0x9e, 0xb0, 0x02, 0x00, 0x00,
}
