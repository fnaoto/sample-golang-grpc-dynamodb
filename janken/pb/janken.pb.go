// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/janken.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Hands int32

const (
	Hands_UNKNOWN_HANDS Hands = 0
	Hands_CHOKI         Hands = 1
	Hands_PA            Hands = 2
	Hands_GU            Hands = 3
)

// Enum value maps for Hands.
var (
	Hands_name = map[int32]string{
		0: "UNKNOWN_HANDS",
		1: "CHOKI",
		2: "PA",
		3: "GU",
	}
	Hands_value = map[string]int32{
		"UNKNOWN_HANDS": 0,
		"CHOKI":         1,
		"PA":            2,
		"GU":            3,
	}
)

func (x Hands) Enum() *Hands {
	p := new(Hands)
	*p = x
	return p
}

func (x Hands) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Hands) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_janken_proto_enumTypes[0].Descriptor()
}

func (Hands) Type() protoreflect.EnumType {
	return &file_proto_janken_proto_enumTypes[0]
}

func (x Hands) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Hands.Descriptor instead.
func (Hands) EnumDescriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{0}
}

type Result int32

const (
	Result_UNKNOWN_RESULT Result = 0
	Result_WIN            Result = 1
	Result_LOSE           Result = 2
	Result_DRAW           Result = 3
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "UNKNOWN_RESULT",
		1: "WIN",
		2: "LOSE",
		3: "DRAW",
	}
	Result_value = map[string]int32{
		"UNKNOWN_RESULT": 0,
		"WIN":            1,
		"LOSE":           2,
		"DRAW":           3,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_janken_proto_enumTypes[1].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_proto_janken_proto_enumTypes[1]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{1}
}

type JankenResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyHand    Hands                  `protobuf:"varint,1,opt,name=myHand,proto3,enum=janken.Hands" json:"myHand,omitempty"`
	YourHand  Hands                  `protobuf:"varint,2,opt,name=yourHand,proto3,enum=janken.Hands" json:"yourHand,omitempty"`
	Result    Result                 `protobuf:"varint,3,opt,name=result,proto3,enum=janken.Result" json:"result,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *JankenResult) Reset() {
	*x = JankenResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_janken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JankenResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JankenResult) ProtoMessage() {}

func (x *JankenResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_janken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JankenResult.ProtoReflect.Descriptor instead.
func (*JankenResult) Descriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{0}
}

func (x *JankenResult) GetMyHand() Hands {
	if x != nil {
		return x.MyHand
	}
	return Hands_UNKNOWN_HANDS
}

func (x *JankenResult) GetYourHand() Hands {
	if x != nil {
		return x.YourHand
	}
	return Hands_UNKNOWN_HANDS
}

func (x *JankenResult) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_UNKNOWN_RESULT
}

func (x *JankenResult) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TryGames      int32           `protobuf:"varint,1,opt,name=tryGames,proto3" json:"tryGames,omitempty"`
	Wins          int32           `protobuf:"varint,2,opt,name=wins,proto3" json:"wins,omitempty"`
	JankenResults []*JankenResult `protobuf:"bytes,3,rep,name=jankenResults,proto3" json:"jankenResults,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_janken_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_proto_janken_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{1}
}

func (x *Report) GetTryGames() int32 {
	if x != nil {
		return x.TryGames
	}
	return 0
}

func (x *Report) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *Report) GetJankenResults() []*JankenResult {
	if x != nil {
		return x.JankenResults
	}
	return nil
}

type PlayJankenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hands Hands `protobuf:"varint,1,opt,name=hands,proto3,enum=janken.Hands" json:"hands,omitempty"`
}

func (x *PlayJankenRequest) Reset() {
	*x = PlayJankenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_janken_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayJankenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayJankenRequest) ProtoMessage() {}

func (x *PlayJankenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_janken_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayJankenRequest.ProtoReflect.Descriptor instead.
func (*PlayJankenRequest) Descriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{2}
}

func (x *PlayJankenRequest) GetHands() Hands {
	if x != nil {
		return x.Hands
	}
	return Hands_UNKNOWN_HANDS
}

type PlayJankenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JankenResult *JankenResult `protobuf:"bytes,1,opt,name=jankenResult,proto3" json:"jankenResult,omitempty"`
}

func (x *PlayJankenResponse) Reset() {
	*x = PlayJankenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_janken_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayJankenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayJankenResponse) ProtoMessage() {}

func (x *PlayJankenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_janken_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayJankenResponse.ProtoReflect.Descriptor instead.
func (*PlayJankenResponse) Descriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{3}
}

func (x *PlayJankenResponse) GetJankenResult() *JankenResult {
	if x != nil {
		return x.JankenResult
	}
	return nil
}

type PlayResultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PlayResultRequest) Reset() {
	*x = PlayResultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_janken_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResultRequest) ProtoMessage() {}

func (x *PlayResultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_janken_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResultRequest.ProtoReflect.Descriptor instead.
func (*PlayResultRequest) Descriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{4}
}

type PlayResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *PlayResultResponse) Reset() {
	*x = PlayResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_janken_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResultResponse) ProtoMessage() {}

func (x *PlayResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_janken_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResultResponse.ProtoReflect.Descriptor instead.
func (*PlayResultResponse) Descriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{5}
}

func (x *PlayResultResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type Timestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_janken_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_janken_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_proto_janken_proto_rawDescGZIP(), []int{6}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_proto_janken_proto protoreflect.FileDescriptor

var file_proto_janken_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01,
	0x0a, 0x0c, 0x4a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25,
	0x0a, 0x06, 0x6d, 0x79, 0x48, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x06, 0x6d,
	0x79, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x79, 0x6f, 0x75, 0x72, 0x48, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e,
	0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x08, 0x79, 0x6f, 0x75, 0x72, 0x48, 0x61, 0x6e, 0x64,
	0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x74, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x72, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x74, 0x72, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x3a, 0x0a,
	0x0d, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x4a, 0x61,
	0x6e, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0d, 0x6a, 0x61, 0x6e, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x50, 0x6c, 0x61,
	0x79, 0x4a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x05, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x22, 0x4e, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x4a, 0x61, 0x6e, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0c, 0x6a, 0x61, 0x6e,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x4a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x3b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61,
	0x6e, 0x6f, 0x73, 0x2a, 0x35, 0x0a, 0x05, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x11, 0x0a, 0x0d,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x53, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x43, 0x48, 0x4f, 0x4b, 0x49, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x50, 0x41,
	0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x47, 0x55, 0x10, 0x03, 0x2a, 0x39, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x52, 0x41, 0x57, 0x10, 0x03, 0x32, 0xa4, 0x01, 0x0a, 0x0d, 0x4a, 0x61, 0x6e, 0x6b, 0x65, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x4a,
	0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x4a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x4a, 0x61,
	0x6e, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c,
	0x0a, 0x11, 0x50, 0x6c, 0x61, 0x79, 0x4a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6a, 0x61, 0x6e, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_janken_proto_rawDescOnce sync.Once
	file_proto_janken_proto_rawDescData = file_proto_janken_proto_rawDesc
)

func file_proto_janken_proto_rawDescGZIP() []byte {
	file_proto_janken_proto_rawDescOnce.Do(func() {
		file_proto_janken_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_janken_proto_rawDescData)
	})
	return file_proto_janken_proto_rawDescData
}

var file_proto_janken_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_janken_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_janken_proto_goTypes = []interface{}{
	(Hands)(0),                    // 0: janken.Hands
	(Result)(0),                   // 1: janken.Result
	(*JankenResult)(nil),          // 2: janken.JankenResult
	(*Report)(nil),                // 3: janken.Report
	(*PlayJankenRequest)(nil),     // 4: janken.PlayJankenRequest
	(*PlayJankenResponse)(nil),    // 5: janken.PlayJankenResponse
	(*PlayResultRequest)(nil),     // 6: janken.PlayResultRequest
	(*PlayResultResponse)(nil),    // 7: janken.PlayResultResponse
	(*Timestamp)(nil),             // 8: janken.Timestamp
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_proto_janken_proto_depIdxs = []int32{
	0,  // 0: janken.JankenResult.myHand:type_name -> janken.Hands
	0,  // 1: janken.JankenResult.yourHand:type_name -> janken.Hands
	1,  // 2: janken.JankenResult.result:type_name -> janken.Result
	9,  // 3: janken.JankenResult.created_at:type_name -> google.protobuf.Timestamp
	2,  // 4: janken.Report.jankenResults:type_name -> janken.JankenResult
	0,  // 5: janken.PlayJankenRequest.hands:type_name -> janken.Hands
	2,  // 6: janken.PlayJankenResponse.jankenResult:type_name -> janken.JankenResult
	3,  // 7: janken.PlayResultResponse.report:type_name -> janken.Report
	4,  // 8: janken.JankenService.PlayJanken:input_type -> janken.PlayJankenRequest
	6,  // 9: janken.JankenService.PlayJankenResults:input_type -> janken.PlayResultRequest
	5,  // 10: janken.JankenService.PlayJanken:output_type -> janken.PlayJankenResponse
	7,  // 11: janken.JankenService.PlayJankenResults:output_type -> janken.PlayResultResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_janken_proto_init() }
func file_proto_janken_proto_init() {
	if File_proto_janken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_janken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JankenResult); i {
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
		file_proto_janken_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_proto_janken_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayJankenRequest); i {
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
		file_proto_janken_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayJankenResponse); i {
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
		file_proto_janken_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResultRequest); i {
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
		file_proto_janken_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResultResponse); i {
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
		file_proto_janken_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
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
			RawDescriptor: file_proto_janken_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_janken_proto_goTypes,
		DependencyIndexes: file_proto_janken_proto_depIdxs,
		EnumInfos:         file_proto_janken_proto_enumTypes,
		MessageInfos:      file_proto_janken_proto_msgTypes,
	}.Build()
	File_proto_janken_proto = out.File
	file_proto_janken_proto_rawDesc = nil
	file_proto_janken_proto_goTypes = nil
	file_proto_janken_proto_depIdxs = nil
}