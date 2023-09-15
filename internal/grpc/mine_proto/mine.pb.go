// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: api/grpc/proto/mine.proto

package mine_proto

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

type GetBlockToMineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetBlockToMineRequest) Reset() {
	*x = GetBlockToMineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockToMineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockToMineRequest) ProtoMessage() {}

func (x *GetBlockToMineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockToMineRequest.ProtoReflect.Descriptor instead.
func (*GetBlockToMineRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlockToMineRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetBlockToMineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool           `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  *DataBlockMine `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Code  int32          `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Type  int32          `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Msg   string         `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetBlockToMineResponse) Reset() {
	*x = GetBlockToMineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockToMineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockToMineResponse) ProtoMessage() {}

func (x *GetBlockToMineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockToMineResponse.ProtoReflect.Descriptor instead.
func (*GetBlockToMineResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockToMineResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *GetBlockToMineResponse) GetData() *DataBlockMine {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetBlockToMineResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetBlockToMineResponse) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *GetBlockToMineResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DataBlockMine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data       []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp  string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	PrevHash   []byte `protobuf:"bytes,4,opt,name=prev_hash,json=prevHash,proto3" json:"prev_hash,omitempty"`
	Difficulty int32  `protobuf:"varint,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
}

func (x *DataBlockMine) Reset() {
	*x = DataBlockMine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBlockMine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBlockMine) ProtoMessage() {}

func (x *DataBlockMine) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBlockMine.ProtoReflect.Descriptor instead.
func (*DataBlockMine) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{2}
}

func (x *DataBlockMine) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DataBlockMine) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DataBlockMine) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *DataBlockMine) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *DataBlockMine) GetDifficulty() int32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

type RequestMineBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Hash       string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Nonce      int64  `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Difficulty int32  `protobuf:"varint,4,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	MinerId    string `protobuf:"bytes,5,opt,name=miner_id,json=minerId,proto3" json:"miner_id,omitempty"`
}

func (x *RequestMineBlock) Reset() {
	*x = RequestMineBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMineBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMineBlock) ProtoMessage() {}

func (x *RequestMineBlock) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMineBlock.ProtoReflect.Descriptor instead.
func (*RequestMineBlock) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{3}
}

func (x *RequestMineBlock) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RequestMineBlock) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *RequestMineBlock) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *RequestMineBlock) GetDifficulty() int32 {
	if x != nil {
		return x.Difficulty
	}
	return 0
}

func (x *RequestMineBlock) GetMinerId() string {
	if x != nil {
		return x.MinerId
	}
	return ""
}

type MineBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  bool   `protobuf:"varint,2,opt,name=data,proto3" json:"data,omitempty"`
	Code  int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Type  int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Msg   string `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MineBlockResponse) Reset() {
	*x = MineBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MineBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MineBlockResponse) ProtoMessage() {}

func (x *MineBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MineBlockResponse.ProtoReflect.Descriptor instead.
func (*MineBlockResponse) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{4}
}

func (x *MineBlockResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *MineBlockResponse) GetData() bool {
	if x != nil {
		return x.Data
	}
	return false
}

func (x *MineBlockResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *MineBlockResponse) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MineBlockResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RequestGenerateGenesis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyGenesis     string  `protobuf:"bytes,1,opt,name=key_genesis,json=keyGenesis,proto3" json:"key_genesis,omitempty"`
	UserId         string  `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	IdentityNumber string  `protobuf:"bytes,3,opt,name=identity_number,json=identityNumber,proto3" json:"identity_number,omitempty"`
	TokensEmmit    float64 `protobuf:"fixed64,4,opt,name=tokens_emmit,json=tokensEmmit,proto3" json:"tokens_emmit,omitempty"`
	WalletsEmmit   int32   `protobuf:"varint,5,opt,name=wallets_emmit,json=walletsEmmit,proto3" json:"wallets_emmit,omitempty"`
}

func (x *RequestGenerateGenesis) Reset() {
	*x = RequestGenerateGenesis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGenerateGenesis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGenerateGenesis) ProtoMessage() {}

func (x *RequestGenerateGenesis) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGenerateGenesis.ProtoReflect.Descriptor instead.
func (*RequestGenerateGenesis) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{5}
}

func (x *RequestGenerateGenesis) GetKeyGenesis() string {
	if x != nil {
		return x.KeyGenesis
	}
	return ""
}

func (x *RequestGenerateGenesis) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RequestGenerateGenesis) GetIdentityNumber() string {
	if x != nil {
		return x.IdentityNumber
	}
	return ""
}

func (x *RequestGenerateGenesis) GetTokensEmmit() float64 {
	if x != nil {
		return x.TokensEmmit
	}
	return 0
}

func (x *RequestGenerateGenesis) GetWalletsEmmit() int32 {
	if x != nil {
		return x.WalletsEmmit
	}
	return 0
}

type ResponseGenerateGenesis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool  `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  *Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Code  int32 `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	Type  int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	Msg   string `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ResponseGenerateGenesis) Reset() {
	*x = ResponseGenerateGenesis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseGenerateGenesis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseGenerateGenesis) ProtoMessage() {}

func (x *ResponseGenerateGenesis) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseGenerateGenesis.ProtoReflect.Descriptor instead.
func (*ResponseGenerateGenesis) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseGenerateGenesis) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *ResponseGenerateGenesis) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResponseGenerateGenesis) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseGenerateGenesis) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ResponseGenerateGenesis) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string        `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	WalletsMain []*WalletMain `protobuf:"bytes,2,rep,name=wallets_main,json=walletsMain,proto3" json:"wallets_main,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{7}
}

func (x *Data) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Data) GetWalletsMain() []*WalletMain {
	if x != nil {
		return x.WalletsMain
	}
	return nil
}

type WalletMain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mnemonic string `protobuf:"bytes,2,opt,name=mnemonic,proto3" json:"mnemonic,omitempty"`
}

func (x *WalletMain) Reset() {
	*x = WalletMain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_proto_mine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletMain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletMain) ProtoMessage() {}

func (x *WalletMain) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_proto_mine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletMain.ProtoReflect.Descriptor instead.
func (*WalletMain) Descriptor() ([]byte, []int) {
	return file_api_grpc_proto_mine_proto_rawDescGZIP(), []int{8}
}

func (x *WalletMain) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WalletMain) GetMnemonic() string {
	if x != nil {
		return x.Mnemonic
	}
	return ""
}

var File_api_grpc_proto_mine_proto protoreflect.FileDescriptor

var file_api_grpc_proto_mine_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x69, 0x6e,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x4d, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x4d, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x69,
	0x6e, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x22, 0x8e, 0x01, 0x0a, 0x0d, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4d, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x69, 0x6e, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x77, 0x0a,
	0x11, 0x4d, 0x69, 0x6e, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xc2, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f, 0x65, 0x6d,
	0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x45, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x73, 0x5f, 0x65, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x45, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x17,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x69,
	0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x5a, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x52, 0x0b, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x22, 0x38, 0x0a, 0x0a, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f,
	0x6e, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6e, 0x65, 0x6d, 0x6f,
	0x6e, 0x69, 0x63, 0x32, 0xa3, 0x02, 0x0a, 0x17, 0x6d, 0x69, 0x6e, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12,
	0x59, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x4d, 0x69, 0x6e,
	0x65, 0x12, 0x21, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x4d, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x6f, 0x4d, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x4d, 0x69,
	0x6e, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x69, 0x6e, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x1d, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x69, 0x6e, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x12, 0x22,
	0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x73,
	0x69, 0x73, 0x1a, 0x23, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x6d,
	0x69, 0x6e, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_grpc_proto_mine_proto_rawDescOnce sync.Once
	file_api_grpc_proto_mine_proto_rawDescData = file_api_grpc_proto_mine_proto_rawDesc
)

func file_api_grpc_proto_mine_proto_rawDescGZIP() []byte {
	file_api_grpc_proto_mine_proto_rawDescOnce.Do(func() {
		file_api_grpc_proto_mine_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_proto_mine_proto_rawDescData)
	})
	return file_api_grpc_proto_mine_proto_rawDescData
}

var file_api_grpc_proto_mine_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_grpc_proto_mine_proto_goTypes = []interface{}{
	(*GetBlockToMineRequest)(nil),   // 0: mine_proto.GetBlockToMineRequest
	(*GetBlockToMineResponse)(nil),  // 1: mine_proto.GetBlockToMineResponse
	(*DataBlockMine)(nil),           // 2: mine_proto.DataBlockMine
	(*RequestMineBlock)(nil),        // 3: mine_proto.RequestMineBlock
	(*MineBlockResponse)(nil),       // 4: mine_proto.MineBlockResponse
	(*RequestGenerateGenesis)(nil),  // 5: mine_proto.RequestGenerateGenesis
	(*ResponseGenerateGenesis)(nil), // 6: mine_proto.ResponseGenerateGenesis
	(*Data)(nil),                    // 7: mine_proto.Data
	(*WalletMain)(nil),              // 8: mine_proto.WalletMain
}
var file_api_grpc_proto_mine_proto_depIdxs = []int32{
	2, // 0: mine_proto.GetBlockToMineResponse.data:type_name -> mine_proto.DataBlockMine
	7, // 1: mine_proto.ResponseGenerateGenesis.data:type_name -> mine_proto.Data
	8, // 2: mine_proto.Data.wallets_main:type_name -> mine_proto.WalletMain
	0, // 3: mine_proto.mineBlockServicesBlocks.GetBlockToMine:input_type -> mine_proto.GetBlockToMineRequest
	3, // 4: mine_proto.mineBlockServicesBlocks.MineBlock:input_type -> mine_proto.RequestMineBlock
	5, // 5: mine_proto.mineBlockServicesBlocks.GenerateBlockGenesis:input_type -> mine_proto.RequestGenerateGenesis
	1, // 6: mine_proto.mineBlockServicesBlocks.GetBlockToMine:output_type -> mine_proto.GetBlockToMineResponse
	4, // 7: mine_proto.mineBlockServicesBlocks.MineBlock:output_type -> mine_proto.MineBlockResponse
	6, // 8: mine_proto.mineBlockServicesBlocks.GenerateBlockGenesis:output_type -> mine_proto.ResponseGenerateGenesis
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_grpc_proto_mine_proto_init() }
func file_api_grpc_proto_mine_proto_init() {
	if File_api_grpc_proto_mine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_proto_mine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockToMineRequest); i {
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
		file_api_grpc_proto_mine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockToMineResponse); i {
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
		file_api_grpc_proto_mine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBlockMine); i {
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
		file_api_grpc_proto_mine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMineBlock); i {
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
		file_api_grpc_proto_mine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MineBlockResponse); i {
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
		file_api_grpc_proto_mine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGenerateGenesis); i {
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
		file_api_grpc_proto_mine_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseGenerateGenesis); i {
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
		file_api_grpc_proto_mine_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_api_grpc_proto_mine_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletMain); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_grpc_proto_mine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_proto_mine_proto_goTypes,
		DependencyIndexes: file_api_grpc_proto_mine_proto_depIdxs,
		MessageInfos:      file_api_grpc_proto_mine_proto_msgTypes,
	}.Build()
	File_api_grpc_proto_mine_proto = out.File
	file_api_grpc_proto_mine_proto_rawDesc = nil
	file_api_grpc_proto_mine_proto_goTypes = nil
	file_api_grpc_proto_mine_proto_depIdxs = nil
}
