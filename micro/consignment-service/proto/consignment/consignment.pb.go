// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package micro_consignment_service

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

// 托运货物
type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Container            []*Container `protobuf:"bytes,4,rep,name=container,proto3" json:"container,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
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

func (m *Consignment) GetContainer() []*Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

// 集装箱
type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
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

// 查看货物信息的请求
type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

// 托运结果
type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{3}
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
	proto.RegisterType((*Consignment)(nil), "micro.consignment.service.Consignment")
	proto.RegisterType((*Container)(nil), "micro.consignment.service.Container")
	proto.RegisterType((*GetRequest)(nil), "micro.consignment.service.GetRequest")
	proto.RegisterType((*Response)(nil), "micro.consignment.service.Response")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4e, 0xeb, 0x30,
	0x14, 0x85, 0x5f, 0xfa, 0x9f, 0x9b, 0xea, 0x55, 0xcf, 0x83, 0x47, 0x80, 0x01, 0x51, 0x0a, 0xa8,
	0xa3, 0x20, 0x95, 0x1d, 0xd0, 0x41, 0x29, 0x43, 0x77, 0x01, 0x55, 0x9b, 0x5c, 0xa5, 0x57, 0x22,
	0x76, 0xb0, 0xdd, 0xb2, 0x35, 0xc4, 0x42, 0x58, 0x0f, 0xaa, 0xd3, 0x50, 0x23, 0x44, 0xd5, 0x59,
	0xce, 0xb9, 0xf7, 0xd8, 0x5f, 0x8e, 0x0c, 0xc3, 0x52, 0x49, 0x23, 0xef, 0x52, 0x29, 0x34, 0xe5,
	0xa2, 0x40, 0x61, 0xdc, 0xef, 0xc4, 0x4e, 0xd9, 0x79, 0x41, 0xa9, 0x92, 0x89, 0x3b, 0xd0, 0xa8,
	0xb6, 0x94, 0x62, 0xfc, 0xe6, 0x41, 0x30, 0x39, 0xf8, 0xec, 0x2f, 0x34, 0x28, 0x0b, 0xbd, 0xc8,
	0x1b, 0xf9, 0xbc, 0x41, 0x19, 0x8b, 0x20, 0xc8, 0x50, 0xa7, 0x8a, 0x4a, 0x43, 0x52, 0x84, 0x0d,
	0x3b, 0x70, 0x2d, 0xf6, 0x1f, 0x3a, 0xaf, 0x48, 0xf9, 0xda, 0x84, 0xcd, 0xc8, 0x1b, 0xb5, 0xf9,
	0x5e, 0xb1, 0x07, 0xf0, 0x53, 0x29, 0xcc, 0x92, 0x04, 0xaa, 0xb0, 0x15, 0x35, 0x47, 0xc1, 0xf8,
	0x3a, 0xf9, 0x15, 0x24, 0x99, 0xd4, 0xbb, 0xfc, 0x10, 0x63, 0x97, 0xe0, 0x6f, 0x51, 0x6b, 0x7c,
	0x5e, 0x50, 0x16, 0xb6, 0xed, 0xdd, 0xbd, 0xca, 0x98, 0x65, 0x71, 0x01, 0xfe, 0x57, 0xe8, 0x07,
	0xf7, 0x15, 0x04, 0xe9, 0x46, 0x1b, 0x59, 0xa0, 0xda, 0x65, 0x2b, 0x6e, 0xa8, 0xad, 0x59, 0xb6,
	0xc3, 0x96, 0x8a, 0x72, 0x12, 0x16, 0xdb, 0xe7, 0x7b, 0xc5, 0xce, 0xa0, 0xbb, 0xd1, 0x55, 0xa8,
	0x55, 0x0d, 0x76, 0x72, 0x96, 0xc5, 0x7d, 0x80, 0x29, 0x1a, 0x8e, 0x2f, 0x1b, 0xd4, 0x26, 0x7e,
	0xf7, 0xa0, 0xc7, 0x51, 0x97, 0x52, 0x68, 0x64, 0x21, 0x74, 0x53, 0x85, 0x4b, 0x83, 0x15, 0x41,
	0x8f, 0xd7, 0x92, 0x3d, 0x42, 0xe0, 0xfc, 0xac, 0xc5, 0x08, 0xc6, 0xb7, 0xc7, 0x6b, 0xa8, 0x3d,
	0xee, 0x46, 0xd9, 0x13, 0xf4, 0x1d, 0xa9, 0xc3, 0xa6, 0x6d, 0xf4, 0xd4, 0xa3, 0xbe, 0x65, 0xc7,
	0x1f, 0x1e, 0x0c, 0xe6, 0x6b, 0x2a, 0x4b, 0x12, 0xf9, 0xbc, 0xda, 0x66, 0x2b, 0xf8, 0x37, 0xb1,
	0xd0, 0xee, 0x6b, 0x38, 0xf1, 0xf8, 0x8b, 0xe1, 0x91, 0xbd, 0xba, 0xa5, 0xf8, 0x0f, 0x5b, 0xc0,
	0x60, 0x8a, 0xc6, 0x09, 0x6a, 0x76, 0x73, 0x24, 0x79, 0xa8, 0xfb, 0xc4, 0x0b, 0x56, 0x1d, 0xfb,
	0xde, 0xef, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xae, 0xa8, 0x5b, 0x16, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	// 创建托运货物
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	// 查看托运货物信息
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "micro.consignment.service"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	// 创建托运货物
	CreateConsignment(context.Context, *Consignment, *Response) error
	// 查看托运货物信息
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}
