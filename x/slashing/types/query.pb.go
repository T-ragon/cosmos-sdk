// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/slashing/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/T-ragon/cosmos-sdk/v3/types/query"
	_ "github.com/T-ragon/cosmos-sdk/v3/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QuerySigningInfoRequest is the request type for the Query/SigningInfo RPC
// method
type QuerySigningInfoRequest struct {
	// cons_address is the address to query signing info of
	ConsAddress string `protobuf:"bytes,1,opt,name=cons_address,json=consAddress,proto3" json:"cons_address,omitempty"`
}

func (m *QuerySigningInfoRequest) Reset()         { *m = QuerySigningInfoRequest{} }
func (m *QuerySigningInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfoRequest) ProtoMessage()    {}
func (*QuerySigningInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{2}
}
func (m *QuerySigningInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfoRequest.Merge(m, src)
}
func (m *QuerySigningInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfoRequest proto.InternalMessageInfo

func (m *QuerySigningInfoRequest) GetConsAddress() string {
	if m != nil {
		return m.ConsAddress
	}
	return ""
}

// QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
// method
type QuerySigningInfoResponse struct {
	// val_signing_info is the signing info of requested val cons address
	ValSigningInfo ValidatorSigningInfo `protobuf:"bytes,1,opt,name=val_signing_info,json=valSigningInfo,proto3" json:"val_signing_info"`
}

func (m *QuerySigningInfoResponse) Reset()         { *m = QuerySigningInfoResponse{} }
func (m *QuerySigningInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfoResponse) ProtoMessage()    {}
func (*QuerySigningInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{3}
}
func (m *QuerySigningInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfoResponse.Merge(m, src)
}
func (m *QuerySigningInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfoResponse proto.InternalMessageInfo

func (m *QuerySigningInfoResponse) GetValSigningInfo() ValidatorSigningInfo {
	if m != nil {
		return m.ValSigningInfo
	}
	return ValidatorSigningInfo{}
}

// QuerySigningInfosRequest is the request type for the Query/SigningInfos RPC
// method
type QuerySigningInfosRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySigningInfosRequest) Reset()         { *m = QuerySigningInfosRequest{} }
func (m *QuerySigningInfosRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfosRequest) ProtoMessage()    {}
func (*QuerySigningInfosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{4}
}
func (m *QuerySigningInfosRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfosRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfosRequest.Merge(m, src)
}
func (m *QuerySigningInfosRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfosRequest proto.InternalMessageInfo

func (m *QuerySigningInfosRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
// method
type QuerySigningInfosResponse struct {
	// info is the signing info of all validators
	Info       []ValidatorSigningInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info"`
	Pagination *query.PageResponse    `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QuerySigningInfosResponse) Reset()         { *m = QuerySigningInfosResponse{} }
func (m *QuerySigningInfosResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySigningInfosResponse) ProtoMessage()    {}
func (*QuerySigningInfosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_791b11d41a861ed0, []int{5}
}
func (m *QuerySigningInfosResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySigningInfosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySigningInfosResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySigningInfosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySigningInfosResponse.Merge(m, src)
}
func (m *QuerySigningInfosResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySigningInfosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySigningInfosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySigningInfosResponse proto.InternalMessageInfo

func (m *QuerySigningInfosResponse) GetInfo() []ValidatorSigningInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *QuerySigningInfosResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.slashing.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.slashing.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QuerySigningInfoRequest)(nil), "cosmos.slashing.v1beta1.QuerySigningInfoRequest")
	proto.RegisterType((*QuerySigningInfoResponse)(nil), "cosmos.slashing.v1beta1.QuerySigningInfoResponse")
	proto.RegisterType((*QuerySigningInfosRequest)(nil), "cosmos.slashing.v1beta1.QuerySigningInfosRequest")
	proto.RegisterType((*QuerySigningInfosResponse)(nil), "cosmos.slashing.v1beta1.QuerySigningInfosResponse")
}

func init() {
	proto.RegisterFile("cosmos/slashing/v1beta1/query.proto", fileDescriptor_791b11d41a861ed0)
}

var fileDescriptor_791b11d41a861ed0 = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0xad, 0x06, 0x3a, 0x29, 0xa2, 0x63, 0xa1, 0x6d, 0xd0, 0x8d, 0x5d, 0x21, 0x2d,
	0xd5, 0xee, 0x98, 0x88, 0xf4, 0xe4, 0xc1, 0x28, 0x8a, 0xe0, 0x41, 0x53, 0x10, 0xf4, 0x12, 0x66,
	0x9b, 0xe9, 0x38, 0xb8, 0x99, 0xd9, 0xee, 0x6c, 0x82, 0x45, 0xf4, 0xe0, 0xd9, 0x83, 0xe0, 0x67,
	0x10, 0x3c, 0xaa, 0xf8, 0x21, 0x7a, 0x2c, 0x7a, 0xf1, 0x24, 0x92, 0x08, 0xde, 0xfd, 0x04, 0xb2,
	0x33, 0x93, 0x64, 0x4a, 0x5c, 0x4d, 0xe9, 0x25, 0x0c, 0x6f, 0xde, 0xff, 0xfd, 0x7f, 0xef, 0xe5,
	0xcd, 0xc2, 0x8b, 0xdb, 0x52, 0x75, 0xa4, 0xc2, 0x2a, 0x22, 0xea, 0x09, 0x17, 0x0c, 0xf7, 0x6a,
	0x21, 0x4d, 0x49, 0x0d, 0xef, 0x76, 0x69, 0xb2, 0x17, 0xc4, 0x89, 0x4c, 0x25, 0x5a, 0x34, 0x49,
	0xc1, 0x30, 0x29, 0xb0, 0x49, 0xe5, 0x75, 0xab, 0x0e, 0x89, 0xa2, 0x46, 0x31, 0xd2, 0xc7, 0x84,
	0x71, 0x41, 0x52, 0x2e, 0x85, 0x29, 0x52, 0x5e, 0x60, 0x92, 0x49, 0x7d, 0xc4, 0xd9, 0xc9, 0x46,
	0xcf, 0x31, 0x29, 0x59, 0x44, 0x31, 0x89, 0x39, 0x26, 0x42, 0xc8, 0x54, 0x4b, 0x94, 0xbd, 0xad,
	0xe6, 0xd1, 0x8d, 0x48, 0x4c, 0xde, 0xb2, 0xc9, 0x6b, 0x99, 0xf2, 0x96, 0xd6, 0x5c, 0x9d, 0x21,
	0x1d, 0x2e, 0x24, 0xd6, 0xbf, 0x26, 0xe4, 0x2f, 0x40, 0xf4, 0x20, 0x63, 0xbd, 0x4f, 0x12, 0xd2,
	0x51, 0x4d, 0xba, 0xdb, 0xa5, 0x2a, 0xf5, 0x1f, 0xc1, 0xb3, 0x87, 0xa2, 0x2a, 0x96, 0x42, 0x51,
	0xd4, 0x80, 0xc5, 0x58, 0x47, 0x96, 0xc0, 0x05, 0xb0, 0x56, 0xaa, 0x57, 0x82, 0x9c, 0x61, 0x04,
	0x46, 0xd8, 0x98, 0xdb, 0xff, 0x5e, 0x29, 0xbc, 0xff, 0xf5, 0x61, 0x1d, 0x34, 0xad, 0xd2, 0x6f,
	0xc1, 0x45, 0x5d, 0x7a, 0x8b, 0x33, 0xc1, 0x05, 0xbb, 0x2b, 0x76, 0xa4, 0x75, 0x45, 0xb7, 0xe0,
	0xfc, 0xb6, 0x14, 0xaa, 0x45, 0xda, 0xed, 0x84, 0x2a, 0x63, 0x32, 0xd7, 0x58, 0xf9, 0xf2, 0x79,
	0xe3, 0xbc, 0xf5, 0xb9, 0x99, 0x61, 0x08, 0xd5, 0x55, 0x37, 0x4c, 0xca, 0x56, 0x9a, 0x70, 0xc1,
	0x9a, 0xa5, 0x4c, 0x66, 0x43, 0xfe, 0x4b, 0xb8, 0x34, 0x69, 0x60, 0x1b, 0x08, 0xe1, 0xe9, 0x1e,
	0x89, 0x5a, 0xca, 0x5c, 0xb5, 0xb8, 0xd8, 0x91, 0xb6, 0x95, 0x8d, 0xdc, 0x56, 0x1e, 0x92, 0x88,
	0xb7, 0x49, 0x2a, 0x13, 0xa7, 0xa0, 0xdb, 0xd8, 0xa9, 0x1e, 0x89, 0x9c, 0x2b, 0x3f, 0x9c, 0xf4,
	0x1f, 0xce, 0x15, 0xdd, 0x86, 0x70, 0xbc, 0x0b, 0xd6, 0xb9, 0x3a, 0x74, 0xce, 0x16, 0x27, 0x30,
	0xab, 0x36, 0x1e, 0x23, 0xa3, 0x56, 0xdb, 0x74, 0x94, 0xfe, 0x27, 0x00, 0x97, 0xff, 0x62, 0x62,
	0xbb, 0xbc, 0x07, 0x4f, 0xd8, 0xce, 0x66, 0x8f, 0xd5, 0x99, 0xae, 0x82, 0xee, 0x1c, 0x62, 0x9e,
	0xd1, 0xcc, 0xab, 0xff, 0x65, 0x36, 0x28, 0x2e, 0x74, 0xfd, 0xf7, 0x2c, 0x3c, 0xa9, 0xa1, 0xd1,
	0x6b, 0x00, 0x8b, 0x66, 0x43, 0xd0, 0xa5, 0x5c, 0xba, 0xc9, 0xb5, 0x2c, 0x5f, 0x9e, 0x2e, 0xd9,
	0x78, 0xfb, 0xab, 0xaf, 0xbe, 0xfe, 0x7c, 0x3b, 0xb3, 0x82, 0x2a, 0x38, 0xef, 0xe5, 0x98, 0x95,
	0x44, 0x1f, 0x01, 0x2c, 0x39, 0x23, 0x40, 0x57, 0xfe, 0x6d, 0x33, 0xb9, 0xb9, 0xe5, 0xda, 0x11,
	0x14, 0x96, 0xee, 0xba, 0xa6, 0xdb, 0x44, 0xd7, 0x72, 0xe9, 0xdc, 0x2d, 0x55, 0xf8, 0xb9, 0xfb,
	0x34, 0x5e, 0xa0, 0x77, 0x00, 0xce, 0xbb, 0x7f, 0x3e, 0x9a, 0x1e, 0x61, 0x34, 0xce, 0xfa, 0x51,
	0x24, 0x16, 0x3b, 0xd0, 0xd8, 0x6b, 0xa8, 0x3a, 0x1d, 0x76, 0x63, 0x73, 0xbf, 0xef, 0x81, 0x83,
	0xbe, 0x07, 0x7e, 0xf4, 0x3d, 0xf0, 0x66, 0xe0, 0x15, 0x0e, 0x06, 0x5e, 0xe1, 0xdb, 0xc0, 0x2b,
	0x3c, 0xb6, 0x6f, 0x5a, 0xb5, 0x9f, 0x06, 0x5c, 0xe2, 0x67, 0xe3, 0x42, 0xe9, 0x5e, 0x4c, 0x55,
	0x58, 0xd4, 0xdf, 0xa7, 0xab, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x12, 0x73, 0x05, 0x05, 0x95,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params queries the parameters of slashing module
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// SigningInfo queries the signing info of given cons address
	SigningInfo(ctx context.Context, in *QuerySigningInfoRequest, opts ...grpc.CallOption) (*QuerySigningInfoResponse, error)
	// SigningInfos queries signing info of all validators
	SigningInfos(ctx context.Context, in *QuerySigningInfosRequest, opts ...grpc.CallOption) (*QuerySigningInfosResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SigningInfo(ctx context.Context, in *QuerySigningInfoRequest, opts ...grpc.CallOption) (*QuerySigningInfoResponse, error) {
	out := new(QuerySigningInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v1beta1.Query/SigningInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SigningInfos(ctx context.Context, in *QuerySigningInfosRequest, opts ...grpc.CallOption) (*QuerySigningInfosResponse, error) {
	out := new(QuerySigningInfosResponse)
	err := c.cc.Invoke(ctx, "/cosmos.slashing.v1beta1.Query/SigningInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params queries the parameters of slashing module
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// SigningInfo queries the signing info of given cons address
	SigningInfo(context.Context, *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error)
	// SigningInfos queries signing info of all validators
	SigningInfos(context.Context, *QuerySigningInfosRequest) (*QuerySigningInfosResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) SigningInfo(ctx context.Context, req *QuerySigningInfoRequest) (*QuerySigningInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigningInfo not implemented")
}
func (*UnimplementedQueryServer) SigningInfos(ctx context.Context, req *QuerySigningInfosRequest) (*QuerySigningInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigningInfos not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SigningInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySigningInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SigningInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v1beta1.Query/SigningInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SigningInfo(ctx, req.(*QuerySigningInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SigningInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySigningInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SigningInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.slashing.v1beta1.Query/SigningInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SigningInfos(ctx, req.(*QuerySigningInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.slashing.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SigningInfo",
			Handler:    _Query_SigningInfo_Handler,
		},
		{
			MethodName: "SigningInfos",
			Handler:    _Query_SigningInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/slashing/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ConsAddress) > 0 {
		i -= len(m.ConsAddress)
		copy(dAtA[i:], m.ConsAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ConsAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ValSigningInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfosRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfosRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfosRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySigningInfosResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySigningInfosResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySigningInfosResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Info) > 0 {
		for iNdEx := len(m.Info) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Info[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySigningInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConsAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySigningInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ValSigningInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySigningInfosRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySigningInfosResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Info) > 0 {
		for _, e := range m.Info {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySigningInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySigningInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySigningInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySigningInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValSigningInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValSigningInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySigningInfosRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySigningInfosRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfosRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySigningInfosResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySigningInfosResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySigningInfosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = append(m.Info, ValidatorSigningInfo{})
			if err := m.Info[len(m.Info)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
