// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: rumexchange.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RumMsgType int32

const (
	RumMsgType_IF_CONN    RumMsgType = 0
	RumMsgType_CONN_RESP  RumMsgType = 1
	RumMsgType_CHAIN_DATA RumMsgType = 2
	RumMsgType_RELAY_REQ  RumMsgType = 3
	RumMsgType_RELAY_RESP RumMsgType = 4
)

// Enum value maps for RumMsgType.
var (
	RumMsgType_name = map[int32]string{
		0: "IF_CONN",
		1: "CONN_RESP",
		2: "CHAIN_DATA",
		3: "RELAY_REQ",
		4: "RELAY_RESP",
	}
	RumMsgType_value = map[string]int32{
		"IF_CONN":    0,
		"CONN_RESP":  1,
		"CHAIN_DATA": 2,
		"RELAY_REQ":  3,
		"RELAY_RESP": 4,
	}
)

func (x RumMsgType) Enum() *RumMsgType {
	p := new(RumMsgType)
	*p = x
	return p
}

func (x RumMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RumMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_rumexchange_proto_enumTypes[0].Descriptor()
}

func (RumMsgType) Type() protoreflect.EnumType {
	return &file_rumexchange_proto_enumTypes[0]
}

func (x RumMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RumMsgType.Descriptor instead.
func (RumMsgType) EnumDescriptor() ([]byte, []int) {
	return file_rumexchange_proto_rawDescGZIP(), []int{0}
}

type RumMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType     RumMsgType       `protobuf:"varint,1,opt,name=MsgType,proto3,enum=quorum.pb.RumMsgType" json:"MsgType,omitempty"`
	IfConn      *SessionIfConn   `protobuf:"bytes,2,opt,name=IfConn,proto3,oneof" json:"IfConn,omitempty"`
	ConnResp    *SessionConnResp `protobuf:"bytes,3,opt,name=ConnResp,proto3,oneof" json:"ConnResp,omitempty"`
	DataPackage *Package         `protobuf:"bytes,4,opt,name=DataPackage,proto3,oneof" json:"DataPackage,omitempty"`
	RelayReq    *RelayReq        `protobuf:"bytes,5,opt,name=RelayReq,proto3,oneof" json:"RelayReq,omitempty"`
	RelayResp   *RelayResp       `protobuf:"bytes,6,opt,name=RelayResp,proto3,oneof" json:"RelayResp,omitempty"`
}

func (x *RumMsg) Reset() {
	*x = RumMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rumexchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RumMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RumMsg) ProtoMessage() {}

func (x *RumMsg) ProtoReflect() protoreflect.Message {
	mi := &file_rumexchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RumMsg.ProtoReflect.Descriptor instead.
func (*RumMsg) Descriptor() ([]byte, []int) {
	return file_rumexchange_proto_rawDescGZIP(), []int{0}
}

func (x *RumMsg) GetMsgType() RumMsgType {
	if x != nil {
		return x.MsgType
	}
	return RumMsgType_IF_CONN
}

func (x *RumMsg) GetIfConn() *SessionIfConn {
	if x != nil {
		return x.IfConn
	}
	return nil
}

func (x *RumMsg) GetConnResp() *SessionConnResp {
	if x != nil {
		return x.ConnResp
	}
	return nil
}

func (x *RumMsg) GetDataPackage() *Package {
	if x != nil {
		return x.DataPackage
	}
	return nil
}

func (x *RumMsg) GetRelayReq() *RelayReq {
	if x != nil {
		return x.RelayReq
	}
	return nil
}

func (x *RumMsg) GetRelayResp() *RelayResp {
	if x != nil {
		return x.RelayResp
	}
	return nil
}

type SessionIfConn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestPeerID   []byte     `protobuf:"bytes,1,opt,name=DestPeerID,proto3" json:"DestPeerID,omitempty"`
	SrcPeerID    []byte     `protobuf:"bytes,2,opt,name=SrcPeerID,proto3" json:"SrcPeerID,omitempty"`
	SessionToken []byte     `protobuf:"bytes,3,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	Signature    []byte     `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"` //sign by the srcPeer
	ChannelId    string     `protobuf:"bytes,5,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
	Peersroutes  []*PeerSig `protobuf:"bytes,6,rep,name=Peersroutes,proto3" json:"Peersroutes,omitempty"`
}

func (x *SessionIfConn) Reset() {
	*x = SessionIfConn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rumexchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionIfConn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionIfConn) ProtoMessage() {}

func (x *SessionIfConn) ProtoReflect() protoreflect.Message {
	mi := &file_rumexchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionIfConn.ProtoReflect.Descriptor instead.
func (*SessionIfConn) Descriptor() ([]byte, []int) {
	return file_rumexchange_proto_rawDescGZIP(), []int{1}
}

func (x *SessionIfConn) GetDestPeerID() []byte {
	if x != nil {
		return x.DestPeerID
	}
	return nil
}

func (x *SessionIfConn) GetSrcPeerID() []byte {
	if x != nil {
		return x.SrcPeerID
	}
	return nil
}

func (x *SessionIfConn) GetSessionToken() []byte {
	if x != nil {
		return x.SessionToken
	}
	return nil
}

func (x *SessionIfConn) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SessionIfConn) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SessionIfConn) GetPeersroutes() []*PeerSig {
	if x != nil {
		return x.Peersroutes
	}
	return nil
}

type SessionConnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DestPeerID   []byte     `protobuf:"bytes,1,opt,name=DestPeerID,proto3" json:"DestPeerID,omitempty"`
	SrcPeerID    []byte     `protobuf:"bytes,2,opt,name=SrcPeerID,proto3" json:"SrcPeerID,omitempty"`
	SessionToken []byte     `protobuf:"bytes,3,opt,name=SessionToken,proto3" json:"SessionToken,omitempty"`
	Signature    []byte     `protobuf:"bytes,4,opt,name=Signature,proto3" json:"Signature,omitempty"` //sign by the destPeer
	ChannelId    string     `protobuf:"bytes,5,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
	Peersroutes  []*PeerSig `protobuf:"bytes,6,rep,name=Peersroutes,proto3" json:"Peersroutes,omitempty"`
}

func (x *SessionConnResp) Reset() {
	*x = SessionConnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rumexchange_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionConnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionConnResp) ProtoMessage() {}

func (x *SessionConnResp) ProtoReflect() protoreflect.Message {
	mi := &file_rumexchange_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionConnResp.ProtoReflect.Descriptor instead.
func (*SessionConnResp) Descriptor() ([]byte, []int) {
	return file_rumexchange_proto_rawDescGZIP(), []int{2}
}

func (x *SessionConnResp) GetDestPeerID() []byte {
	if x != nil {
		return x.DestPeerID
	}
	return nil
}

func (x *SessionConnResp) GetSrcPeerID() []byte {
	if x != nil {
		return x.SrcPeerID
	}
	return nil
}

func (x *SessionConnResp) GetSessionToken() []byte {
	if x != nil {
		return x.SessionToken
	}
	return nil
}

func (x *SessionConnResp) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SessionConnResp) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SessionConnResp) GetPeersroutes() []*PeerSig {
	if x != nil {
		return x.Peersroutes
	}
	return nil
}

type PeerSig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId     []byte `protobuf:"bytes,1,opt,name=PeerId,proto3" json:"PeerId,omitempty"`
	SessionSig []byte `protobuf:"bytes,2,opt,name=SessionSig,proto3" json:"SessionSig,omitempty"`
}

func (x *PeerSig) Reset() {
	*x = PeerSig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rumexchange_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerSig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerSig) ProtoMessage() {}

func (x *PeerSig) ProtoReflect() protoreflect.Message {
	mi := &file_rumexchange_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerSig.ProtoReflect.Descriptor instead.
func (*PeerSig) Descriptor() ([]byte, []int) {
	return file_rumexchange_proto_rawDescGZIP(), []int{3}
}

func (x *PeerSig) GetPeerId() []byte {
	if x != nil {
		return x.PeerId
	}
	return nil
}

func (x *PeerSig) GetSessionSig() []byte {
	if x != nil {
		return x.SessionSig
	}
	return nil
}

type GroupRelayItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelayId     string `protobuf:"bytes,1,opt,name=RelayId,proto3" json:"RelayId,omitempty"`
	GroupId     string `protobuf:"bytes,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	UserPubkey  string `protobuf:"bytes,3,opt,name=UserPubkey,proto3" json:"UserPubkey,omitempty"`
	Duration    int64  `protobuf:"varint,4,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Type        string `protobuf:"bytes,5,opt,name=Type,proto3" json:"Type,omitempty"`
	SenderSign  string `protobuf:"bytes,6,opt,name=SenderSign,proto3" json:"SenderSign,omitempty"`
	Memo        string `protobuf:"bytes,7,opt,name=Memo,proto3" json:"Memo,omitempty"`
	ApproveTime int64  `protobuf:"varint,8,opt,name=ApproveTime,proto3" json:"ApproveTime,omitempty"`
	ReqPeerId   string `protobuf:"bytes,9,opt,name=ReqPeerId,proto3" json:"ReqPeerId,omitempty"`
	RelayPeerId string `protobuf:"bytes,10,opt,name=RelayPeerId,proto3" json:"RelayPeerId,omitempty"`
}

func (x *GroupRelayItem) Reset() {
	*x = GroupRelayItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rumexchange_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRelayItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRelayItem) ProtoMessage() {}

func (x *GroupRelayItem) ProtoReflect() protoreflect.Message {
	mi := &file_rumexchange_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRelayItem.ProtoReflect.Descriptor instead.
func (*GroupRelayItem) Descriptor() ([]byte, []int) {
	return file_rumexchange_proto_rawDescGZIP(), []int{4}
}

func (x *GroupRelayItem) GetRelayId() string {
	if x != nil {
		return x.RelayId
	}
	return ""
}

func (x *GroupRelayItem) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GroupRelayItem) GetUserPubkey() string {
	if x != nil {
		return x.UserPubkey
	}
	return ""
}

func (x *GroupRelayItem) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *GroupRelayItem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GroupRelayItem) GetSenderSign() string {
	if x != nil {
		return x.SenderSign
	}
	return ""
}

func (x *GroupRelayItem) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *GroupRelayItem) GetApproveTime() int64 {
	if x != nil {
		return x.ApproveTime
	}
	return 0
}

func (x *GroupRelayItem) GetReqPeerId() string {
	if x != nil {
		return x.ReqPeerId
	}
	return ""
}

func (x *GroupRelayItem) GetRelayPeerId() string {
	if x != nil {
		return x.RelayPeerId
	}
	return ""
}

type RelayReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    string `protobuf:"bytes,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	UserPubkey string `protobuf:"bytes,2,opt,name=UserPubkey,proto3" json:"UserPubkey,omitempty"`
	Duration   int64  `protobuf:"varint,3,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Type       string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	SenderSign string `protobuf:"bytes,5,opt,name=SenderSign,proto3" json:"SenderSign,omitempty"`
	Memo       string `protobuf:"bytes,6,opt,name=Memo,proto3" json:"Memo,omitempty"`
}

func (x *RelayReq) Reset() {
	*x = RelayReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rumexchange_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayReq) ProtoMessage() {}

func (x *RelayReq) ProtoReflect() protoreflect.Message {
	mi := &file_rumexchange_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayReq.ProtoReflect.Descriptor instead.
func (*RelayReq) Descriptor() ([]byte, []int) {
	return file_rumexchange_proto_rawDescGZIP(), []int{5}
}

func (x *RelayReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *RelayReq) GetUserPubkey() string {
	if x != nil {
		return x.UserPubkey
	}
	return ""
}

func (x *RelayReq) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *RelayReq) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RelayReq) GetSenderSign() string {
	if x != nil {
		return x.SenderSign
	}
	return ""
}

func (x *RelayReq) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type RelayResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelayId     string `protobuf:"bytes,1,opt,name=RelayId,proto3" json:"RelayId,omitempty"`
	GroupId     string `protobuf:"bytes,2,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	UserPubkey  string `protobuf:"bytes,3,opt,name=UserPubkey,proto3" json:"UserPubkey,omitempty"`
	Duration    int64  `protobuf:"varint,4,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Type        string `protobuf:"bytes,5,opt,name=Type,proto3" json:"Type,omitempty"`
	SenderSign  string `protobuf:"bytes,6,opt,name=SenderSign,proto3" json:"SenderSign,omitempty"`
	Memo        string `protobuf:"bytes,7,opt,name=Memo,proto3" json:"Memo,omitempty"`
	ApproveTime int64  `protobuf:"varint,8,opt,name=ApproveTime,proto3" json:"ApproveTime,omitempty"`
	RelayPeerId []byte `protobuf:"bytes,9,opt,name=RelayPeerId,proto3" json:"RelayPeerId,omitempty"`
}

func (x *RelayResp) Reset() {
	*x = RelayResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rumexchange_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayResp) ProtoMessage() {}

func (x *RelayResp) ProtoReflect() protoreflect.Message {
	mi := &file_rumexchange_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayResp.ProtoReflect.Descriptor instead.
func (*RelayResp) Descriptor() ([]byte, []int) {
	return file_rumexchange_proto_rawDescGZIP(), []int{6}
}

func (x *RelayResp) GetRelayId() string {
	if x != nil {
		return x.RelayId
	}
	return ""
}

func (x *RelayResp) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *RelayResp) GetUserPubkey() string {
	if x != nil {
		return x.UserPubkey
	}
	return ""
}

func (x *RelayResp) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *RelayResp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RelayResp) GetSenderSign() string {
	if x != nil {
		return x.SenderSign
	}
	return ""
}

func (x *RelayResp) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *RelayResp) GetApproveTime() int64 {
	if x != nil {
		return x.ApproveTime
	}
	return 0
}

func (x *RelayResp) GetRelayPeerId() []byte {
	if x != nil {
		return x.RelayPeerId
	}
	return nil
}

var File_rumexchange_proto protoreflect.FileDescriptor

var file_rumexchange_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x75, 0x6d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x62, 0x1a, 0x0b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x03, 0x0a, 0x06,
	0x52, 0x75, 0x6d, 0x4d, 0x73, 0x67, 0x12, 0x2f, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6d, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07,
	0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x49, 0x66, 0x43, 0x6f, 0x6e,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x66, 0x43, 0x6f, 0x6e,
	0x6e, 0x48, 0x00, 0x52, 0x06, 0x49, 0x66, 0x43, 0x6f, 0x6e, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3b,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x48, 0x01, 0x52, 0x08,
	0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0b, 0x44,
	0x61, 0x74, 0x61, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x48, 0x02, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x71, 0x75, 0x6f, 0x72, 0x75,
	0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x48, 0x03, 0x52,
	0x08, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x09,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x48, 0x04, 0x52, 0x09, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x49, 0x66, 0x43, 0x6f, 0x6e, 0x6e,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x44, 0x61, 0x74, 0x61, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0xe3, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x66, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65,
	0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x44, 0x65, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x72,
	0x63, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53,
	0x72, 0x63, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x69,
	0x67, 0x52, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x73, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0xe5,
	0x01, 0x0a, 0x0f, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x44, 0x65, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x72, 0x63, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x72, 0x63, 0x50, 0x65, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x73, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x53, 0x69, 0x67, 0x52, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x07, 0x50, 0x65, 0x65, 0x72, 0x53, 0x69,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x22, 0xaa, 0x02, 0x0a, 0x0e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x50, 0x65, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x50, 0x65,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x50, 0x65, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa8, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x79,
	0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x65, 0x6d,
	0x6f, 0x22, 0x87, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x69, 0x67,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x6c,
	0x61, 0x79, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x2a, 0x57, 0x0a, 0x0a, 0x52,
	0x75, 0x6d, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x46, 0x5f,
	0x43, 0x4f, 0x4e, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x4c, 0x41, 0x59, 0x5f, 0x52,
	0x45, 0x51, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x4c, 0x41, 0x59, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x10, 0x04, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x75, 0x6d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x72, 0x75, 0x6d,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rumexchange_proto_rawDescOnce sync.Once
	file_rumexchange_proto_rawDescData = file_rumexchange_proto_rawDesc
)

func file_rumexchange_proto_rawDescGZIP() []byte {
	file_rumexchange_proto_rawDescOnce.Do(func() {
		file_rumexchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_rumexchange_proto_rawDescData)
	})
	return file_rumexchange_proto_rawDescData
}

var file_rumexchange_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rumexchange_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rumexchange_proto_goTypes = []interface{}{
	(RumMsgType)(0),         // 0: quorum.pb.RumMsgType
	(*RumMsg)(nil),          // 1: quorum.pb.RumMsg
	(*SessionIfConn)(nil),   // 2: quorum.pb.SessionIfConn
	(*SessionConnResp)(nil), // 3: quorum.pb.SessionConnResp
	(*PeerSig)(nil),         // 4: quorum.pb.PeerSig
	(*GroupRelayItem)(nil),  // 5: quorum.pb.GroupRelayItem
	(*RelayReq)(nil),        // 6: quorum.pb.RelayReq
	(*RelayResp)(nil),       // 7: quorum.pb.RelayResp
	(*Package)(nil),         // 8: quorum.pb.Package
}
var file_rumexchange_proto_depIdxs = []int32{
	0, // 0: quorum.pb.RumMsg.MsgType:type_name -> quorum.pb.RumMsgType
	2, // 1: quorum.pb.RumMsg.IfConn:type_name -> quorum.pb.SessionIfConn
	3, // 2: quorum.pb.RumMsg.ConnResp:type_name -> quorum.pb.SessionConnResp
	8, // 3: quorum.pb.RumMsg.DataPackage:type_name -> quorum.pb.Package
	6, // 4: quorum.pb.RumMsg.RelayReq:type_name -> quorum.pb.RelayReq
	7, // 5: quorum.pb.RumMsg.RelayResp:type_name -> quorum.pb.RelayResp
	4, // 6: quorum.pb.SessionIfConn.Peersroutes:type_name -> quorum.pb.PeerSig
	4, // 7: quorum.pb.SessionConnResp.Peersroutes:type_name -> quorum.pb.PeerSig
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_rumexchange_proto_init() }
func file_rumexchange_proto_init() {
	if File_rumexchange_proto != nil {
		return
	}
	file_chain_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rumexchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RumMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rumexchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionIfConn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rumexchange_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionConnResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rumexchange_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerSig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rumexchange_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRelayItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rumexchange_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rumexchange_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_rumexchange_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rumexchange_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rumexchange_proto_goTypes,
		DependencyIndexes: file_rumexchange_proto_depIdxs,
		EnumInfos:         file_rumexchange_proto_enumTypes,
		MessageInfos:      file_rumexchange_proto_msgTypes,
	}.Build()
	File_rumexchange_proto = out.File
	file_rumexchange_proto_rawDesc = nil
	file_rumexchange_proto_goTypes = nil
	file_rumexchange_proto_depIdxs = nil
}
