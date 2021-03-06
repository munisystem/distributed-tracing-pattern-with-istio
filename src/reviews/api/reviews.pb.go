// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reviews.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	reviews.proto

It has these top-level messages:
	GetCustomerReviewRequest
	GetCustomerReviewResponse
*/
package pb

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

type GetCustomerReviewRequest struct {
	CustomerId int64 `protobuf:"varint,1,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
}

func (m *GetCustomerReviewRequest) Reset()                    { *m = GetCustomerReviewRequest{} }
func (m *GetCustomerReviewRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCustomerReviewRequest) ProtoMessage()               {}
func (*GetCustomerReviewRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetCustomerReviewRequest) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

type GetCustomerReviewResponse struct {
	Id          int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CustomerId  int64  `protobuf:"varint,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	ItemName    string `protobuf:"bytes,3,opt,name=item_name,json=itemName" json:"item_name,omitempty"`
	Rating      int64  `protobuf:"varint,4,opt,name=rating" json:"rating,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
}

func (m *GetCustomerReviewResponse) Reset()                    { *m = GetCustomerReviewResponse{} }
func (m *GetCustomerReviewResponse) String() string            { return proto.CompactTextString(m) }
func (*GetCustomerReviewResponse) ProtoMessage()               {}
func (*GetCustomerReviewResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetCustomerReviewResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetCustomerReviewResponse) GetCustomerId() int64 {
	if m != nil {
		return m.CustomerId
	}
	return 0
}

func (m *GetCustomerReviewResponse) GetItemName() string {
	if m != nil {
		return m.ItemName
	}
	return ""
}

func (m *GetCustomerReviewResponse) GetRating() int64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *GetCustomerReviewResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerReviewRequest)(nil), "review.GetCustomerReviewRequest")
	proto.RegisterType((*GetCustomerReviewResponse)(nil), "review.GetCustomerReviewResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Review service

type ReviewClient interface {
	GetCustomerReview(ctx context.Context, in *GetCustomerReviewRequest, opts ...grpc.CallOption) (*GetCustomerReviewResponse, error)
}

type reviewClient struct {
	cc *grpc.ClientConn
}

func NewReviewClient(cc *grpc.ClientConn) ReviewClient {
	return &reviewClient{cc}
}

func (c *reviewClient) GetCustomerReview(ctx context.Context, in *GetCustomerReviewRequest, opts ...grpc.CallOption) (*GetCustomerReviewResponse, error) {
	out := new(GetCustomerReviewResponse)
	err := grpc.Invoke(ctx, "/review.Review/GetCustomerReview", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Review service

type ReviewServer interface {
	GetCustomerReview(context.Context, *GetCustomerReviewRequest) (*GetCustomerReviewResponse, error)
}

func RegisterReviewServer(s *grpc.Server, srv ReviewServer) {
	s.RegisterService(&_Review_serviceDesc, srv)
}

func _Review_GetCustomerReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServer).GetCustomerReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/review.Review/GetCustomerReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServer).GetCustomerReview(ctx, req.(*GetCustomerReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Review_serviceDesc = grpc.ServiceDesc{
	ServiceName: "review.Review",
	HandlerType: (*ReviewServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerReview",
			Handler:    _Review_GetCustomerReview_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reviews.proto",
}

func init() { proto.RegisterFile("reviews.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0xcb,
	0x4c, 0x2d, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x70, 0x95, 0xac, 0xb9,
	0x24, 0xdc, 0x53, 0x4b, 0x9c, 0x4b, 0x8b, 0x4b, 0xf2, 0x73, 0x53, 0x8b, 0x82, 0xc0, 0x82, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xf2, 0x5c, 0xdc, 0xc9, 0x50, 0x89, 0xf8, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x2e, 0x98, 0x90, 0x67, 0x8a, 0xd2, 0x62, 0x46, 0x2e,
	0x49, 0x2c, 0xba, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xf8, 0xb8, 0x98, 0xe0, 0xba, 0x98,
	0x32, 0x53, 0xd0, 0x8d, 0x63, 0x42, 0x37, 0x4e, 0x48, 0x9a, 0x8b, 0x33, 0xb3, 0x24, 0x35, 0x37,
	0x3e, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0x33, 0x88, 0x03, 0x24, 0xe0, 0x97,
	0x98, 0x9b, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94, 0x58, 0x92, 0x99, 0x97, 0x2e, 0xc1, 0x02, 0xd6,
	0x08, 0xe5, 0x09, 0x29, 0x70, 0x71, 0xa7, 0xa4, 0x16, 0x27, 0x17, 0x65, 0x16, 0x94, 0x64, 0xe6,
	0xe7, 0x49, 0xb0, 0x82, 0xb5, 0x21, 0x0b, 0x19, 0xa5, 0x70, 0xb1, 0x41, 0x5c, 0x26, 0x14, 0xc5,
	0x25, 0x88, 0xe1, 0x5c, 0x21, 0x05, 0x3d, 0x48, 0x50, 0xe8, 0xe1, 0x0a, 0x07, 0x29, 0x45, 0x3c,
	0x2a, 0x20, 0x7e, 0x55, 0x62, 0x70, 0x62, 0x89, 0x62, 0x2a, 0x48, 0x4a, 0x62, 0x03, 0x87, 0xae,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x09, 0xcf, 0x4f, 0xe1, 0x6e, 0x01, 0x00, 0x00,
}
