// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/consensus/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	v11 "github.com/cometbft/cometbft/api/cometbft/abci/v1"
	v1 "github.com/cometbft/cometbft/api/cometbft/types/v1"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/T-ragon/cosmos-sdk/v3/types/msgservice"
	_ "github.com/T-ragon/cosmos-sdk/v3/types/tx/amino"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the x/consensus parameters to update.
	// VersionsParams is not included in this Msg because it is tracked
	// separarately in x/upgrade.
	//
	// NOTE: All parameters must be supplied.
	Block     *v1.BlockParams     `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Evidence  *v1.EvidenceParams  `protobuf:"bytes,3,opt,name=evidence,proto3" json:"evidence,omitempty"`
	Validator *v1.ValidatorParams `protobuf:"bytes,4,opt,name=validator,proto3" json:"validator,omitempty"`
	// Since: cosmos-sdk 0.51
	Abci      *v1.ABCIParams      `protobuf:"bytes,5,opt,name=abci,proto3" json:"abci,omitempty"` // Deprecated: Do not use.
	Synchrony *v1.SynchronyParams `protobuf:"bytes,6,opt,name=synchrony,proto3" json:"synchrony,omitempty"`
	Feature   *v1.FeatureParams   `protobuf:"bytes,7,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_2135c60575ab504d, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetBlock() *v1.BlockParams {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *MsgUpdateParams) GetEvidence() *v1.EvidenceParams {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *MsgUpdateParams) GetValidator() *v1.ValidatorParams {
	if m != nil {
		return m.Validator
	}
	return nil
}

// Deprecated: Do not use.
func (m *MsgUpdateParams) GetAbci() *v1.ABCIParams {
	if m != nil {
		return m.Abci
	}
	return nil
}

func (m *MsgUpdateParams) GetSynchrony() *v1.SynchronyParams {
	if m != nil {
		return m.Synchrony
	}
	return nil
}

func (m *MsgUpdateParams) GetFeature() *v1.FeatureParams {
	if m != nil {
		return m.Feature
	}
	return nil
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2135c60575ab504d, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

// MsgCometInfo is the Msg/CometInfo request type.
type MsgSetCometInfo struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// evidence is the misbehaviour evidence to submit.
	Evidence []*v11.Misbehavior `protobuf:"bytes,2,rep,name=evidence,proto3" json:"evidence,omitempty"`
	// validators_hash is the hash of the current validator set.
	ValidatorsHash []byte `protobuf:"bytes,3,opt,name=validators_hash,json=validatorsHash,proto3" json:"validators_hash,omitempty"`
	// proposer_address is the address of the current proposer.
	ProposerAddress []byte `protobuf:"bytes,4,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
	// last_commit is the last commit info.
	LastCommit *v11.CommitInfo `protobuf:"bytes,5,opt,name=last_commit,json=lastCommit,proto3" json:"last_commit,omitempty"`
}

func (m *MsgSetCometInfo) Reset()         { *m = MsgSetCometInfo{} }
func (m *MsgSetCometInfo) String() string { return proto.CompactTextString(m) }
func (*MsgSetCometInfo) ProtoMessage()    {}
func (*MsgSetCometInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2135c60575ab504d, []int{2}
}
func (m *MsgSetCometInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetCometInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetCometInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetCometInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetCometInfo.Merge(m, src)
}
func (m *MsgSetCometInfo) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetCometInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetCometInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetCometInfo proto.InternalMessageInfo

func (m *MsgSetCometInfo) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSetCometInfo) GetEvidence() []*v11.Misbehavior {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *MsgSetCometInfo) GetValidatorsHash() []byte {
	if m != nil {
		return m.ValidatorsHash
	}
	return nil
}

func (m *MsgSetCometInfo) GetProposerAddress() []byte {
	if m != nil {
		return m.ProposerAddress
	}
	return nil
}

func (m *MsgSetCometInfo) GetLastCommit() *v11.CommitInfo {
	if m != nil {
		return m.LastCommit
	}
	return nil
}

// MsgCometInfoResponse defines the response
type MsgSetCometInfoResponse struct {
}

func (m *MsgSetCometInfoResponse) Reset()         { *m = MsgSetCometInfoResponse{} }
func (m *MsgSetCometInfoResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetCometInfoResponse) ProtoMessage()    {}
func (*MsgSetCometInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2135c60575ab504d, []int{3}
}
func (m *MsgSetCometInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetCometInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetCometInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetCometInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetCometInfoResponse.Merge(m, src)
}
func (m *MsgSetCometInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetCometInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetCometInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetCometInfoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "cosmos.consensus.v1.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "cosmos.consensus.v1.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgSetCometInfo)(nil), "cosmos.consensus.v1.MsgSetCometInfo")
	proto.RegisterType((*MsgSetCometInfoResponse)(nil), "cosmos.consensus.v1.MsgSetCometInfoResponse")
}

func init() { proto.RegisterFile("cosmos/consensus/v1/tx.proto", fileDescriptor_2135c60575ab504d) }

var fileDescriptor_2135c60575ab504d = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x4f, 0xd4, 0x40,
	0x18, 0xa6, 0x7c, 0xba, 0xc3, 0xc6, 0xd5, 0x41, 0x43, 0xd9, 0x60, 0xb3, 0xae, 0x46, 0x91, 0x48,
	0xcb, 0x22, 0xa2, 0x90, 0x90, 0xc8, 0x12, 0x0d, 0x1c, 0x88, 0xa6, 0x04, 0x0f, 0x5e, 0x36, 0xd3,
	0x76, 0xd8, 0x36, 0xd0, 0x4e, 0xd3, 0x19, 0x2a, 0x7b, 0x33, 0x1e, 0x3d, 0xf9, 0x47, 0x4c, 0x38,
	0xf0, 0x23, 0x8c, 0x17, 0x89, 0x27, 0xe3, 0xc9, 0xec, 0x1e, 0x36, 0xfe, 0x0b, 0xd3, 0x99, 0x7e,
	0xec, 0x47, 0x37, 0x31, 0x5e, 0x9a, 0x74, 0xde, 0xe7, 0x79, 0xde, 0x77, 0x9e, 0xf7, 0x69, 0xc1,
	0xa2, 0x49, 0xa8, 0x4b, 0xa8, 0x66, 0x12, 0x8f, 0x62, 0x8f, 0x9e, 0x51, 0x2d, 0xac, 0x69, 0xec,
	0x5c, 0xf5, 0x03, 0xc2, 0x08, 0x9c, 0x13, 0x55, 0x35, 0xad, 0xaa, 0x61, 0xad, 0x7c, 0x13, 0xb9,
	0x8e, 0x47, 0x34, 0xfe, 0x14, 0xb8, 0xf2, 0x82, 0xc0, 0x35, 0xf8, 0x9b, 0x16, 0x93, 0x44, 0x69,
	0x3e, 0x6e, 0xe0, 0xd2, 0x66, 0x24, 0xed, 0xd2, 0x66, 0x5c, 0x50, 0x4c, 0xe2, 0x62, 0x66, 0x1c,
	0x33, 0x8d, 0xb5, 0x7c, 0xcc, 0xfb, 0xfa, 0x28, 0x40, 0x6e, 0x42, 0x5c, 0x4c, 0xeb, 0xc8, 0x30,
	0x1d, 0x3e, 0x56, 0x84, 0x13, 0xd5, 0xea, 0x97, 0x49, 0x50, 0x3a, 0xa0, 0xcd, 0x23, 0xdf, 0x42,
	0x0c, 0xbf, 0xe1, 0x3c, 0xb8, 0x01, 0x0a, 0xe8, 0x8c, 0xd9, 0x24, 0x70, 0x58, 0x4b, 0x96, 0x2a,
	0xd2, 0x52, 0xa1, 0x2e, 0xff, 0xb8, 0x5c, 0xb9, 0x15, 0xcf, 0xb3, 0x63, 0x59, 0x01, 0xa6, 0xf4,
	0x90, 0x05, 0x8e, 0xd7, 0xd4, 0x33, 0x28, 0x5c, 0x07, 0x53, 0xc6, 0x29, 0x31, 0x4f, 0xe4, 0xf1,
	0x8a, 0xb4, 0x34, 0xbb, 0xa6, 0xa8, 0x49, 0x67, 0x55, 0x74, 0x0c, 0x6b, 0x6a, 0x3d, 0xaa, 0x8b,
	0x36, 0xba, 0x00, 0xc3, 0x6d, 0x70, 0x0d, 0x87, 0x8e, 0x85, 0x3d, 0x13, 0xcb, 0x13, 0x9c, 0x78,
	0x37, 0x87, 0xf8, 0x32, 0x86, 0xc4, 0xdc, 0x94, 0x02, 0x5f, 0x80, 0x42, 0x88, 0x4e, 0x1d, 0x0b,
	0x31, 0x12, 0xc8, 0x93, 0x9c, 0x5f, 0xcd, 0xe1, 0xbf, 0x4d, 0x30, 0xb1, 0x40, 0x46, 0x82, 0x7b,
	0x60, 0x32, 0x72, 0x46, 0x9e, 0xe2, 0xe4, 0x3b, 0x39, 0xe4, 0x9d, 0xfa, 0xee, 0xbe, 0xe0, 0xd5,
	0x6f, 0xff, 0xba, 0x5c, 0x29, 0x09, 0x23, 0x56, 0xa8, 0x75, 0x52, 0x59, 0x55, 0x9f, 0xae, 0xca,
	0x92, 0xce, 0x15, 0xe0, 0x11, 0x28, 0xd0, 0x96, 0x67, 0xda, 0x01, 0xf1, 0x5a, 0xf2, 0xf4, 0xc8,
	0x59, 0x0e, 0x13, 0x4c, 0xac, 0x39, 0x37, 0xac, 0x59, 0xd3, 0x33, 0x25, 0xf8, 0x1a, 0xcc, 0x1c,
	0x63, 0xc4, 0xce, 0x02, 0x2c, 0xcf, 0x70, 0xd1, 0x4a, 0x8e, 0xe8, 0x2b, 0x81, 0x18, 0x2d, 0xb9,
	0xa6, 0x27, 0x2a, 0x5b, 0x9b, 0x1f, 0xbb, 0x17, 0xcb, 0xd9, 0xe2, 0x3e, 0x75, 0x2f, 0x96, 0x1f,
	0x64, 0x60, 0xed, 0xbc, 0x27, 0xc5, 0x03, 0xd9, 0xa8, 0x2e, 0x80, 0xf9, 0x81, 0x23, 0x1d, 0x53,
	0x3f, 0x82, 0x57, 0xbf, 0x8f, 0xf3, 0x28, 0x1d, 0x62, 0xb6, 0x1b, 0x4d, 0xb7, 0xef, 0x1d, 0x93,
	0xff, 0x8e, 0xd2, 0x66, 0x4f, 0x28, 0xc6, 0x2b, 0x13, 0xfd, 0x7b, 0x89, 0xbc, 0x8e, 0xae, 0x7c,
	0xe0, 0x50, 0x03, 0xdb, 0x28, 0x74, 0x48, 0xd0, 0x13, 0x88, 0x87, 0xa0, 0x94, 0xee, 0x96, 0x36,
	0x6c, 0x44, 0x6d, 0x1e, 0xab, 0xa2, 0x7e, 0x3d, 0x3b, 0xde, 0x43, 0xd4, 0x86, 0x8f, 0xc0, 0x0d,
	0x3f, 0x20, 0x3e, 0xa1, 0x38, 0x68, 0x20, 0x31, 0x08, 0x0f, 0x50, 0x51, 0x2f, 0x25, 0xe7, 0xf1,
	0x7c, 0x70, 0x1b, 0xcc, 0x9e, 0x22, 0xca, 0x1a, 0x26, 0x71, 0x5d, 0x87, 0xc5, 0x49, 0x59, 0x1c,
	0x9e, 0x68, 0x97, 0xd7, 0xa3, 0x9b, 0xeb, 0x20, 0x22, 0x88, 0xf7, 0xad, 0x8d, 0x61, 0xbf, 0xef,
	0x8d, 0xf6, 0x3b, 0x75, 0x2f, 0x36, 0xbb, 0xd7, 0xd0, 0xc4, 0xec, 0xb5, 0x3f, 0x12, 0x98, 0x38,
	0xa0, 0x4d, 0xf8, 0x1e, 0x14, 0xfb, 0xbe, 0xdd, 0xfb, 0x6a, 0xce, 0xaf, 0x46, 0x1d, 0x58, 0x59,
	0xf9, 0xf1, 0xbf, 0xa0, 0xd2, 0xc5, 0xce, 0x7d, 0x1b, 0x0c, 0xd3, 0xfa, 0x33, 0x68, 0x80, 0x62,
	0xdf, 0xa6, 0x47, 0x36, 0xee, 0x45, 0x8d, 0x6e, 0x9c, 0x77, 0xc9, 0xf2, 0xd4, 0x87, 0xee, 0xc5,
	0xb2, 0x54, 0x7f, 0xfe, 0xb5, 0xad, 0x48, 0x57, 0x6d, 0x45, 0xfa, 0xdd, 0x56, 0xa4, 0xcf, 0x1d,
	0x65, 0xec, 0xaa, 0xa3, 0x8c, 0xfd, 0xec, 0x28, 0x63, 0xef, 0x14, 0xa1, 0x46, 0xad, 0x13, 0xd5,
	0x21, 0x7d, 0x3e, 0xf2, 0xef, 0xc2, 0x98, 0xe6, 0x3f, 0xb9, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x65, 0xd1, 0xd8, 0x21, 0x9e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a governance operation for updating the x/consensus module parameters.
	// The authority is defined in the keeper.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// SetCometInfo defines how to set the comet info for the x/consensus module.
	SetCometInfo(ctx context.Context, in *MsgSetCometInfo, opts ...grpc.CallOption) (*MsgSetCometInfoResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.consensus.v1.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetCometInfo(ctx context.Context, in *MsgSetCometInfo, opts ...grpc.CallOption) (*MsgSetCometInfoResponse, error) {
	out := new(MsgSetCometInfoResponse)
	err := c.cc.Invoke(ctx, "/cosmos.consensus.v1.Msg/SetCometInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a governance operation for updating the x/consensus module parameters.
	// The authority is defined in the keeper.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// SetCometInfo defines how to set the comet info for the x/consensus module.
	SetCometInfo(context.Context, *MsgSetCometInfo) (*MsgSetCometInfoResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) SetCometInfo(ctx context.Context, req *MsgSetCometInfo) (*MsgSetCometInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCometInfo not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.consensus.v1.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetCometInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetCometInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetCometInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.consensus.v1.Msg/SetCometInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetCometInfo(ctx, req.(*MsgSetCometInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.consensus.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "SetCometInfo",
			Handler:    _Msg_SetCometInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/consensus/v1/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Feature != nil {
		{
			size, err := m.Feature.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Synchrony != nil {
		{
			size, err := m.Synchrony.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Abci != nil {
		{
			size, err := m.Abci.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Validator != nil {
		{
			size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Evidence != nil {
		{
			size, err := m.Evidence.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetCometInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetCometInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetCometInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastCommit != nil {
		{
			size, err := m.LastCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ProposerAddress) > 0 {
		i -= len(m.ProposerAddress)
		copy(dAtA[i:], m.ProposerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ProposerAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ValidatorsHash) > 0 {
		i -= len(m.ValidatorsHash)
		copy(dAtA[i:], m.ValidatorsHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorsHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Evidence) > 0 {
		for iNdEx := len(m.Evidence) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Evidence[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetCometInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetCometInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetCometInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Evidence != nil {
		l = m.Evidence.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Validator != nil {
		l = m.Validator.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Abci != nil {
		l = m.Abci.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Synchrony != nil {
		l = m.Synchrony.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Feature != nil {
		l = m.Feature.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetCometInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Evidence) > 0 {
		for _, e := range m.Evidence {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.ValidatorsHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ProposerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.LastCommit != nil {
		l = m.LastCommit.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetCometInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &v1.BlockParams{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evidence == nil {
				m.Evidence = &v1.EvidenceParams{}
			}
			if err := m.Evidence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validator == nil {
				m.Validator = &v1.ValidatorParams{}
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abci", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Abci == nil {
				m.Abci = &v1.ABCIParams{}
			}
			if err := m.Abci.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Synchrony", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Synchrony == nil {
				m.Synchrony = &v1.SynchronyParams{}
			}
			if err := m.Synchrony.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feature", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Feature == nil {
				m.Feature = &v1.FeatureParams{}
			}
			if err := m.Feature.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetCometInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetCometInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetCometInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evidence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Evidence = append(m.Evidence, &v11.Misbehavior{})
			if err := m.Evidence[len(m.Evidence)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorsHash = append(m.ValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorsHash == nil {
				m.ValidatorsHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerAddress = append(m.ProposerAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ProposerAddress == nil {
				m.ProposerAddress = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastCommit == nil {
				m.LastCommit = &v11.CommitInfo{}
			}
			if err := m.LastCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetCometInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetCometInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetCometInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
