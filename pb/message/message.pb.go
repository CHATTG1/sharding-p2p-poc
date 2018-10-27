// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package proto_message

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

type Response_Status int32

const (
	Response_SUCCESS Response_Status = 0
	Response_FAILURE Response_Status = 1
)

var Response_Status_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILURE",
}
var Response_Status_value = map[string]int32{
	"SUCCESS": 0,
	"FAILURE": 1,
}

func (x Response_Status) String() string {
	return proto.EnumName(Response_Status_name, int32(x))
}
func (Response_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{0, 0}
}

type Response struct {
	Status               Response_Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.message.Response_Status" json:"status,omitempty"`
	Message              string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{0}
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

func (m *Response) GetStatus() Response_Status {
	if m != nil {
		return m.Status
	}
	return Response_SUCCESS
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddPeerRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPeerRequest) Reset()         { *m = AddPeerRequest{} }
func (m *AddPeerRequest) String() string { return proto.CompactTextString(m) }
func (*AddPeerRequest) ProtoMessage()    {}
func (*AddPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{1}
}
func (m *AddPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPeerRequest.Unmarshal(m, b)
}
func (m *AddPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPeerRequest.Marshal(b, m, deterministic)
}
func (dst *AddPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPeerRequest.Merge(dst, src)
}
func (m *AddPeerRequest) XXX_Size() int {
	return xxx_messageInfo_AddPeerRequest.Size(m)
}
func (m *AddPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPeerRequest proto.InternalMessageInfo

func (m *AddPeerRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddPeerResponse struct {
	Response             *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddPeerResponse) Reset()         { *m = AddPeerResponse{} }
func (m *AddPeerResponse) String() string { return proto.CompactTextString(m) }
func (*AddPeerResponse) ProtoMessage()    {}
func (*AddPeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{2}
}
func (m *AddPeerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPeerResponse.Unmarshal(m, b)
}
func (m *AddPeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPeerResponse.Marshal(b, m, deterministic)
}
func (dst *AddPeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPeerResponse.Merge(dst, src)
}
func (m *AddPeerResponse) XXX_Size() int {
	return xxx_messageInfo_AddPeerResponse.Size(m)
}
func (m *AddPeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPeerResponse proto.InternalMessageInfo

func (m *AddPeerResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

type Collation struct {
	ShardID              int64    `protobuf:"varint,1,opt,name=shardID,proto3" json:"shardID,omitempty"`
	Period               int64    `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	Blobs                []byte   `protobuf:"bytes,3,opt,name=blobs,proto3" json:"blobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Collation) Reset()         { *m = Collation{} }
func (m *Collation) String() string { return proto.CompactTextString(m) }
func (*Collation) ProtoMessage()    {}
func (*Collation) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{3}
}
func (m *Collation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collation.Unmarshal(m, b)
}
func (m *Collation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collation.Marshal(b, m, deterministic)
}
func (dst *Collation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collation.Merge(dst, src)
}
func (m *Collation) XXX_Size() int {
	return xxx_messageInfo_Collation.Size(m)
}
func (m *Collation) XXX_DiscardUnknown() {
	xxx_messageInfo_Collation.DiscardUnknown(m)
}

var xxx_messageInfo_Collation proto.InternalMessageInfo

func (m *Collation) GetShardID() int64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *Collation) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *Collation) GetBlobs() []byte {
	if m != nil {
		return m.Blobs
	}
	return nil
}

type CollationRequest struct {
	ShardID              int64    `protobuf:"varint,1,opt,name=shardID,proto3" json:"shardID,omitempty"`
	Period               int64    `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	Hash                 string   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollationRequest) Reset()         { *m = CollationRequest{} }
func (m *CollationRequest) String() string { return proto.CompactTextString(m) }
func (*CollationRequest) ProtoMessage()    {}
func (*CollationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{4}
}
func (m *CollationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollationRequest.Unmarshal(m, b)
}
func (m *CollationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollationRequest.Marshal(b, m, deterministic)
}
func (dst *CollationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollationRequest.Merge(dst, src)
}
func (m *CollationRequest) XXX_Size() int {
	return xxx_messageInfo_CollationRequest.Size(m)
}
func (m *CollationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollationRequest proto.InternalMessageInfo

func (m *CollationRequest) GetShardID() int64 {
	if m != nil {
		return m.ShardID
	}
	return 0
}

func (m *CollationRequest) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *CollationRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type CollationResponse struct {
	Response             *Response  `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Collation            *Collation `protobuf:"bytes,2,opt,name=collation,proto3" json:"collation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CollationResponse) Reset()         { *m = CollationResponse{} }
func (m *CollationResponse) String() string { return proto.CompactTextString(m) }
func (*CollationResponse) ProtoMessage()    {}
func (*CollationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{5}
}
func (m *CollationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollationResponse.Unmarshal(m, b)
}
func (m *CollationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollationResponse.Marshal(b, m, deterministic)
}
func (dst *CollationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollationResponse.Merge(dst, src)
}
func (m *CollationResponse) XXX_Size() int {
	return xxx_messageInfo_CollationResponse.Size(m)
}
func (m *CollationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollationResponse proto.InternalMessageInfo

func (m *CollationResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *CollationResponse) GetCollation() *Collation {
	if m != nil {
		return m.Collation
	}
	return nil
}

type ShardPeerRequest struct {
	ShardIDs             []int64  `protobuf:"varint,1,rep,packed,name=shardIDs,proto3" json:"shardIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardPeerRequest) Reset()         { *m = ShardPeerRequest{} }
func (m *ShardPeerRequest) String() string { return proto.CompactTextString(m) }
func (*ShardPeerRequest) ProtoMessage()    {}
func (*ShardPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{6}
}
func (m *ShardPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardPeerRequest.Unmarshal(m, b)
}
func (m *ShardPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardPeerRequest.Marshal(b, m, deterministic)
}
func (dst *ShardPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardPeerRequest.Merge(dst, src)
}
func (m *ShardPeerRequest) XXX_Size() int {
	return xxx_messageInfo_ShardPeerRequest.Size(m)
}
func (m *ShardPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShardPeerRequest proto.InternalMessageInfo

func (m *ShardPeerRequest) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

type ShardPeerResponse struct {
	Response             *Response        `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	ShardPeers           map[int64]*Peers `protobuf:"bytes,2,rep,name=shardPeers,proto3" json:"shardPeers,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ShardPeerResponse) Reset()         { *m = ShardPeerResponse{} }
func (m *ShardPeerResponse) String() string { return proto.CompactTextString(m) }
func (*ShardPeerResponse) ProtoMessage()    {}
func (*ShardPeerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{7}
}
func (m *ShardPeerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardPeerResponse.Unmarshal(m, b)
}
func (m *ShardPeerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardPeerResponse.Marshal(b, m, deterministic)
}
func (dst *ShardPeerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardPeerResponse.Merge(dst, src)
}
func (m *ShardPeerResponse) XXX_Size() int {
	return xxx_messageInfo_ShardPeerResponse.Size(m)
}
func (m *ShardPeerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardPeerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShardPeerResponse proto.InternalMessageInfo

func (m *ShardPeerResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ShardPeerResponse) GetShardPeers() map[int64]*Peers {
	if m != nil {
		return m.ShardPeers
	}
	return nil
}

type NotifyShardsRequest struct {
	ShardIDs             []int64  `protobuf:"varint,1,rep,packed,name=shardIDs,proto3" json:"shardIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyShardsRequest) Reset()         { *m = NotifyShardsRequest{} }
func (m *NotifyShardsRequest) String() string { return proto.CompactTextString(m) }
func (*NotifyShardsRequest) ProtoMessage()    {}
func (*NotifyShardsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{8}
}
func (m *NotifyShardsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyShardsRequest.Unmarshal(m, b)
}
func (m *NotifyShardsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyShardsRequest.Marshal(b, m, deterministic)
}
func (dst *NotifyShardsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyShardsRequest.Merge(dst, src)
}
func (m *NotifyShardsRequest) XXX_Size() int {
	return xxx_messageInfo_NotifyShardsRequest.Size(m)
}
func (m *NotifyShardsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyShardsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyShardsRequest proto.InternalMessageInfo

func (m *NotifyShardsRequest) GetShardIDs() []int64 {
	if m != nil {
		return m.ShardIDs
	}
	return nil
}

type GeneralRequest struct {
	MsgType              int64    `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralRequest) Reset()         { *m = GeneralRequest{} }
func (m *GeneralRequest) String() string { return proto.CompactTextString(m) }
func (*GeneralRequest) ProtoMessage()    {}
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{9}
}
func (m *GeneralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralRequest.Unmarshal(m, b)
}
func (m *GeneralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralRequest.Marshal(b, m, deterministic)
}
func (dst *GeneralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralRequest.Merge(dst, src)
}
func (m *GeneralRequest) XXX_Size() int {
	return xxx_messageInfo_GeneralRequest.Size(m)
}
func (m *GeneralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralRequest proto.InternalMessageInfo

func (m *GeneralRequest) GetMsgType() int64 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *GeneralRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GeneralResponse struct {
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeneralResponse) Reset()         { *m = GeneralResponse{} }
func (m *GeneralResponse) String() string { return proto.CompactTextString(m) }
func (*GeneralResponse) ProtoMessage()    {}
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{10}
}
func (m *GeneralResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralResponse.Unmarshal(m, b)
}
func (m *GeneralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralResponse.Marshal(b, m, deterministic)
}
func (dst *GeneralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralResponse.Merge(dst, src)
}
func (m *GeneralResponse) XXX_Size() int {
	return xxx_messageInfo_GeneralResponse.Size(m)
}
func (m *GeneralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralResponse proto.InternalMessageInfo

func (m *GeneralResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type MessageWithType struct {
	MsgType              int64    `protobuf:"varint,1,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageWithType) Reset()         { *m = MessageWithType{} }
func (m *MessageWithType) String() string { return proto.CompactTextString(m) }
func (*MessageWithType) ProtoMessage()    {}
func (*MessageWithType) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{11}
}
func (m *MessageWithType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithType.Unmarshal(m, b)
}
func (m *MessageWithType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithType.Marshal(b, m, deterministic)
}
func (dst *MessageWithType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithType.Merge(dst, src)
}
func (m *MessageWithType) XXX_Size() int {
	return xxx_messageInfo_MessageWithType.Size(m)
}
func (m *MessageWithType) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithType.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithType proto.InternalMessageInfo

func (m *MessageWithType) GetMsgType() int64 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *MessageWithType) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Peers struct {
	Peers                []string `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peers) Reset()         { *m = Peers{} }
func (m *Peers) String() string { return proto.CompactTextString(m) }
func (*Peers) ProtoMessage()    {}
func (*Peers) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_9fa843598f47b987, []int{12}
}
func (m *Peers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peers.Unmarshal(m, b)
}
func (m *Peers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peers.Marshal(b, m, deterministic)
}
func (dst *Peers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peers.Merge(dst, src)
}
func (m *Peers) XXX_Size() int {
	return xxx_messageInfo_Peers.Size(m)
}
func (m *Peers) XXX_DiscardUnknown() {
	xxx_messageInfo_Peers.DiscardUnknown(m)
}

var xxx_messageInfo_Peers proto.InternalMessageInfo

func (m *Peers) GetPeers() []string {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*Response)(nil), "proto.message.Response")
	proto.RegisterType((*AddPeerRequest)(nil), "proto.message.AddPeerRequest")
	proto.RegisterType((*AddPeerResponse)(nil), "proto.message.AddPeerResponse")
	proto.RegisterType((*Collation)(nil), "proto.message.Collation")
	proto.RegisterType((*CollationRequest)(nil), "proto.message.CollationRequest")
	proto.RegisterType((*CollationResponse)(nil), "proto.message.CollationResponse")
	proto.RegisterType((*ShardPeerRequest)(nil), "proto.message.ShardPeerRequest")
	proto.RegisterType((*ShardPeerResponse)(nil), "proto.message.ShardPeerResponse")
	proto.RegisterMapType((map[int64]*Peers)(nil), "proto.message.ShardPeerResponse.ShardPeersEntry")
	proto.RegisterType((*NotifyShardsRequest)(nil), "proto.message.NotifyShardsRequest")
	proto.RegisterType((*GeneralRequest)(nil), "proto.message.GeneralRequest")
	proto.RegisterType((*GeneralResponse)(nil), "proto.message.GeneralResponse")
	proto.RegisterType((*MessageWithType)(nil), "proto.message.MessageWithType")
	proto.RegisterType((*Peers)(nil), "proto.message.Peers")
	proto.RegisterEnum("proto.message.Response_Status", Response_Status_name, Response_Status_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_9fa843598f47b987) }

var fileDescriptor_message_9fa843598f47b987 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x51, 0x6b, 0x13, 0x41,
	0x10, 0x76, 0x73, 0x26, 0xcd, 0x4d, 0xda, 0xe4, 0xba, 0x16, 0x3d, 0x0a, 0x4a, 0x58, 0x10, 0x8e,
	0x3e, 0x1c, 0x9a, 0x42, 0x11, 0x1f, 0x94, 0x12, 0x53, 0x29, 0xa8, 0x94, 0x3d, 0x8b, 0xbe, 0x6e,
	0xbc, 0xb5, 0x09, 0x9e, 0xd9, 0xf3, 0x76, 0x23, 0xe4, 0xcd, 0x7f, 0xec, 0x5f, 0x90, 0x9d, 0xdb,
	0xbb, 0x5c, 0x4f, 0x85, 0x96, 0x3e, 0xdd, 0x7c, 0x7b, 0xdf, 0x7c, 0x33, 0xdf, 0xcc, 0xc0, 0xde,
	0x77, 0xa9, 0xb5, 0xb8, 0x92, 0x71, 0x5e, 0x28, 0xa3, 0xe8, 0x1e, 0x7e, 0x62, 0xf7, 0xc8, 0x7e,
	0x11, 0xe8, 0x73, 0xa9, 0x73, 0xb5, 0xd2, 0x92, 0x9e, 0x40, 0x4f, 0x1b, 0x61, 0xd6, 0x3a, 0x24,
	0x63, 0x12, 0x0d, 0x27, 0x4f, 0xe2, 0x6b, 0xe4, 0xb8, 0x22, 0xc6, 0x09, 0xb2, 0xb8, 0x63, 0xd3,
	0x10, 0x76, 0x1c, 0x25, 0xec, 0x8c, 0x49, 0xe4, 0xf3, 0x0a, 0x32, 0x06, 0xbd, 0x92, 0x4b, 0x07,
	0xb0, 0x93, 0x5c, 0x4e, 0xa7, 0xb3, 0x24, 0x09, 0xee, 0x59, 0x70, 0x76, 0x7a, 0xfe, 0xee, 0x92,
	0xcf, 0x02, 0xc2, 0x8e, 0x60, 0x78, 0x9a, 0xa6, 0x17, 0x52, 0x16, 0x5c, 0xfe, 0x58, 0x4b, 0x6d,
	0x9a, 0x7a, 0xe4, 0xba, 0xde, 0x19, 0x8c, 0x6a, 0xae, 0x6b, 0xfa, 0x18, 0xfa, 0x85, 0x8b, 0x91,
	0x3d, 0x98, 0x3c, 0xfa, 0x4f, 0xdb, 0xbc, 0x26, 0xb2, 0x04, 0xfc, 0xa9, 0xca, 0x32, 0x61, 0x96,
	0x6a, 0x65, 0xcb, 0xe9, 0x85, 0x28, 0xd2, 0xf3, 0x37, 0x28, 0xe0, 0xf1, 0x0a, 0xd2, 0x87, 0xd0,
	0xcb, 0x65, 0xb1, 0x54, 0x29, 0xfa, 0xf2, 0xb8, 0x43, 0xf4, 0x00, 0xba, 0xf3, 0x4c, 0xcd, 0x75,
	0xe8, 0x8d, 0x49, 0xb4, 0xcb, 0x4b, 0xc0, 0x3e, 0x43, 0x50, 0x8b, 0x36, 0xac, 0xdc, 0x52, 0x9b,
	0xc2, 0xfd, 0x85, 0xd0, 0x0b, 0x94, 0xf6, 0x39, 0xc6, 0x76, 0x4b, 0xfb, 0x0d, 0xe9, 0x3b, 0x38,
	0xa7, 0x27, 0xe0, 0x7f, 0xa9, 0x94, 0xb0, 0xf2, 0x60, 0x12, 0xb6, 0xb2, 0xb6, 0x95, 0xb6, 0x54,
	0x16, 0x43, 0x90, 0xd8, 0xce, 0x9b, 0x7b, 0x3a, 0x84, 0xbe, 0x73, 0x63, 0x2f, 0xc6, 0x8b, 0x3c,
	0x5e, 0x63, 0xf6, 0x9b, 0xc0, 0x7e, 0x23, 0xe1, 0x2e, 0x2d, 0x5f, 0x00, 0xe8, 0x4a, 0x49, 0x87,
	0x9d, 0xb1, 0x17, 0x0d, 0x26, 0xcf, 0x5a, 0x69, 0x7f, 0x95, 0xda, 0xbe, 0xe8, 0xd9, 0xca, 0x14,
	0x1b, 0xde, 0xd0, 0x38, 0x4c, 0x60, 0xd4, 0xfa, 0x4d, 0x03, 0xf0, 0xbe, 0xc9, 0x8d, 0x5b, 0x92,
	0x0d, 0xe9, 0x11, 0x74, 0x7f, 0x8a, 0x6c, 0x2d, 0xdd, 0x94, 0x0e, 0x5a, 0x15, 0x31, 0x97, 0x97,
	0x94, 0x97, 0x9d, 0x17, 0x84, 0x3d, 0x87, 0x07, 0x1f, 0x94, 0x59, 0x7e, 0xdd, 0xa0, 0xb4, 0xbe,
	0xc9, 0x90, 0x5e, 0xc1, 0xf0, 0xad, 0x5c, 0xc9, 0x42, 0x64, 0xcd, 0xd3, 0xd7, 0x57, 0x1f, 0x37,
	0xb9, 0xac, 0xee, 0xc5, 0x41, 0x7b, 0x17, 0xa9, 0x30, 0x02, 0xbb, 0xd9, 0xe5, 0x18, 0xb3, 0xa7,
	0x30, 0xaa, 0xf3, 0xdd, 0xb0, 0xfe, 0x45, 0x7b, 0x0d, 0xa3, 0xf7, 0x65, 0xd7, 0x9f, 0x96, 0x66,
	0x81, 0x6a, 0xb7, 0xab, 0xf3, 0x18, 0xba, 0x68, 0xd7, 0x1e, 0x7e, 0x8e, 0x5b, 0xb0, 0x4e, 0x7c,
	0x5e, 0x82, 0x79, 0x0f, 0x27, 0x73, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x43, 0xd8, 0x15, 0x3e,
	0x6b, 0x04, 0x00, 0x00,
}
