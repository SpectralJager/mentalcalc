// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: pkg/proto/gateway.proto

package gateway

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Scores   string `protobuf:"bytes,2,opt,name=scores,proto3" json:"scores,omitempty"`
	Lvl      int64  `protobuf:"varint,3,opt,name=lvl,proto3" json:"lvl,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Result) GetScores() string {
	if x != nil {
		return x.Scores
	}
	return ""
}

func (x *Result) GetLvl() int64 {
	if x != nil {
		return x.Lvl
	}
	return 0
}

type UserLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserLoginRequest) Reset() {
	*x = UserLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRequest) ProtoMessage() {}

func (x *UserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRequest.ProtoReflect.Descriptor instead.
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *UserLoginRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserLoginResponse) Reset() {
	*x = UserLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResponse) ProtoMessage() {}

func (x *UserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResponse.ProtoReflect.Descriptor instead.
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserRegisterRequest) Reset() {
	*x = UserRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRequest) ProtoMessage() {}

func (x *UserRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRequest.ProtoReflect.Descriptor instead.
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *UserRegisterRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User  *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserRegisterResponse) Reset() {
	*x = UserRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterResponse) ProtoMessage() {}

func (x *UserRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterResponse.ProtoReflect.Descriptor instead.
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *UserRegisterResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserRegisterResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewUser *User  `protobuf:"bytes,1,opt,name=newUser,proto3" json:"newUser,omitempty"`
	OldUser *User  `protobuf:"bytes,2,opt,name=oldUser,proto3" json:"oldUser,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserUpdateRequest) Reset() {
	*x = UserUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRequest) ProtoMessage() {}

func (x *UserUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *UserUpdateRequest) GetNewUser() *User {
	if x != nil {
		return x.NewUser
	}
	return nil
}

func (x *UserUpdateRequest) GetOldUser() *User {
	if x != nil {
		return x.OldUser
	}
	return nil
}

func (x *UserUpdateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User  *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserUpdateResponse) Reset() {
	*x = UserUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateResponse) ProtoMessage() {}

func (x *UserUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateResponse.ProtoReflect.Descriptor instead.
func (*UserUpdateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *UserUpdateResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserUpdateResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type ResultSaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameType string  `protobuf:"bytes,1,opt,name=gameType,proto3" json:"gameType,omitempty"`
	GameMode string  `protobuf:"bytes,2,opt,name=gameMode,proto3" json:"gameMode,omitempty"`
	Result   *Result `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultSaveRequest) Reset() {
	*x = ResultSaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultSaveRequest) ProtoMessage() {}

func (x *ResultSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultSaveRequest.ProtoReflect.Descriptor instead.
func (*ResultSaveRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{8}
}

func (x *ResultSaveRequest) GetGameType() string {
	if x != nil {
		return x.GameType
	}
	return ""
}

func (x *ResultSaveRequest) GetGameMode() string {
	if x != nil {
		return x.GameMode
	}
	return ""
}

func (x *ResultSaveRequest) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type ResultSaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position int64   `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Result   *Result `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *ResultSaveResponse) Reset() {
	*x = ResultSaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultSaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultSaveResponse) ProtoMessage() {}

func (x *ResultSaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultSaveResponse.ProtoReflect.Descriptor instead.
func (*ResultSaveResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{9}
}

func (x *ResultSaveResponse) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *ResultSaveResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type ResultGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameType string `protobuf:"bytes,1,opt,name=gameType,proto3" json:"gameType,omitempty"`
	GameMode string `protobuf:"bytes,2,opt,name=gameMode,proto3" json:"gameMode,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ResultGetRequest) Reset() {
	*x = ResultGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultGetRequest) ProtoMessage() {}

func (x *ResultGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultGetRequest.ProtoReflect.Descriptor instead.
func (*ResultGetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{10}
}

func (x *ResultGetRequest) GetGameType() string {
	if x != nil {
		return x.GameType
	}
	return ""
}

func (x *ResultGetRequest) GetGameMode() string {
	if x != nil {
		return x.GameMode
	}
	return ""
}

func (x *ResultGetRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ResultGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position int64   `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Result   *Result `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultGetResponse) Reset() {
	*x = ResultGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultGetResponse) ProtoMessage() {}

func (x *ResultGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultGetResponse.ProtoReflect.Descriptor instead.
func (*ResultGetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{11}
}

func (x *ResultGetResponse) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *ResultGetResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type ResultsGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameType string `protobuf:"bytes,1,opt,name=gameType,proto3" json:"gameType,omitempty"`
	GameMode string `protobuf:"bytes,2,opt,name=gameMode,proto3" json:"gameMode,omitempty"`
	Start    int64  `protobuf:"varint,3,opt,name=start,proto3" json:"start,omitempty"`
	Count    int64  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ResultsGetRequest) Reset() {
	*x = ResultsGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultsGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultsGetRequest) ProtoMessage() {}

func (x *ResultsGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultsGetRequest.ProtoReflect.Descriptor instead.
func (*ResultsGetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{12}
}

func (x *ResultsGetRequest) GetGameType() string {
	if x != nil {
		return x.GameType
	}
	return ""
}

func (x *ResultsGetRequest) GetGameMode() string {
	if x != nil {
		return x.GameMode
	}
	return ""
}

func (x *ResultsGetRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ResultsGetRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ResultsGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position int64   `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Result   *Result `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ResultsGetResponse) Reset() {
	*x = ResultsGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_gateway_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultsGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultsGetResponse) ProtoMessage() {}

func (x *ResultsGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_gateway_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultsGetResponse.ProtoReflect.Descriptor instead.
func (*ResultsGetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_gateway_proto_rawDescGZIP(), []int{13}
}

func (x *ResultsGetResponse) GetPosition() int64 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *ResultsGetResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_pkg_proto_gateway_proto protoreflect.FileDescriptor

var file_pkg_proto_gateway_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x22, 0x3e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x4e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x76, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6c,
	0x76, 0x6c, 0x22, 0x35, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x29, 0x0a, 0x11, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x4f,
	0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x7b, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x0a,
	0x07, 0x6f, 0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x6f,
	0x6c, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4d, 0x0a, 0x12,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x74, 0x0a, 0x11, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x59, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x66, 0x0a, 0x10,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x77,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xb0, 0x03, 0x0a, 0x02, 0x76, 0x31, 0x12, 0x42, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x61, 0x76, 0x65, 0x12,
	0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x47, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_proto_gateway_proto_rawDescOnce sync.Once
	file_pkg_proto_gateway_proto_rawDescData = file_pkg_proto_gateway_proto_rawDesc
)

func file_pkg_proto_gateway_proto_rawDescGZIP() []byte {
	file_pkg_proto_gateway_proto_rawDescOnce.Do(func() {
		file_pkg_proto_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_gateway_proto_rawDescData)
	})
	return file_pkg_proto_gateway_proto_rawDescData
}

var file_pkg_proto_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_pkg_proto_gateway_proto_goTypes = []interface{}{
	(*User)(nil),                 // 0: gateway.User
	(*Result)(nil),               // 1: gateway.Result
	(*UserLoginRequest)(nil),     // 2: gateway.UserLoginRequest
	(*UserLoginResponse)(nil),    // 3: gateway.UserLoginResponse
	(*UserRegisterRequest)(nil),  // 4: gateway.UserRegisterRequest
	(*UserRegisterResponse)(nil), // 5: gateway.UserRegisterResponse
	(*UserUpdateRequest)(nil),    // 6: gateway.UserUpdateRequest
	(*UserUpdateResponse)(nil),   // 7: gateway.UserUpdateResponse
	(*ResultSaveRequest)(nil),    // 8: gateway.ResultSaveRequest
	(*ResultSaveResponse)(nil),   // 9: gateway.ResultSaveResponse
	(*ResultGetRequest)(nil),     // 10: gateway.ResultGetRequest
	(*ResultGetResponse)(nil),    // 11: gateway.ResultGetResponse
	(*ResultsGetRequest)(nil),    // 12: gateway.ResultsGetRequest
	(*ResultsGetResponse)(nil),   // 13: gateway.ResultsGetResponse
}
var file_pkg_proto_gateway_proto_depIdxs = []int32{
	0,  // 0: gateway.UserLoginRequest.user:type_name -> gateway.User
	0,  // 1: gateway.UserRegisterRequest.user:type_name -> gateway.User
	0,  // 2: gateway.UserRegisterResponse.user:type_name -> gateway.User
	0,  // 3: gateway.UserUpdateRequest.newUser:type_name -> gateway.User
	0,  // 4: gateway.UserUpdateRequest.oldUser:type_name -> gateway.User
	0,  // 5: gateway.UserUpdateResponse.user:type_name -> gateway.User
	1,  // 6: gateway.ResultSaveRequest.result:type_name -> gateway.Result
	1,  // 7: gateway.ResultSaveResponse.Result:type_name -> gateway.Result
	1,  // 8: gateway.ResultGetResponse.result:type_name -> gateway.Result
	1,  // 9: gateway.ResultsGetResponse.result:type_name -> gateway.Result
	2,  // 10: gateway.v1.UserLogin:input_type -> gateway.UserLoginRequest
	4,  // 11: gateway.v1.UserRegister:input_type -> gateway.UserRegisterRequest
	6,  // 12: gateway.v1.UserUpdate:input_type -> gateway.UserUpdateRequest
	8,  // 13: gateway.v1.ResultSave:input_type -> gateway.ResultSaveRequest
	10, // 14: gateway.v1.ResultGet:input_type -> gateway.ResultGetRequest
	12, // 15: gateway.v1.ResultsGet:input_type -> gateway.ResultsGetRequest
	3,  // 16: gateway.v1.UserLogin:output_type -> gateway.UserLoginResponse
	5,  // 17: gateway.v1.UserRegister:output_type -> gateway.UserRegisterResponse
	7,  // 18: gateway.v1.UserUpdate:output_type -> gateway.UserUpdateResponse
	9,  // 19: gateway.v1.ResultSave:output_type -> gateway.ResultSaveResponse
	11, // 20: gateway.v1.ResultGet:output_type -> gateway.ResultGetResponse
	13, // 21: gateway.v1.ResultsGet:output_type -> gateway.ResultsGetResponse
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_proto_gateway_proto_init() }
func file_pkg_proto_gateway_proto_init() {
	if File_pkg_proto_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_pkg_proto_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_pkg_proto_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRequest); i {
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
		file_pkg_proto_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResponse); i {
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
		file_pkg_proto_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterRequest); i {
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
		file_pkg_proto_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterResponse); i {
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
		file_pkg_proto_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateRequest); i {
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
		file_pkg_proto_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateResponse); i {
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
		file_pkg_proto_gateway_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultSaveRequest); i {
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
		file_pkg_proto_gateway_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultSaveResponse); i {
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
		file_pkg_proto_gateway_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultGetRequest); i {
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
		file_pkg_proto_gateway_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultGetResponse); i {
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
		file_pkg_proto_gateway_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultsGetRequest); i {
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
		file_pkg_proto_gateway_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultsGetResponse); i {
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
			RawDescriptor: file_pkg_proto_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_gateway_proto_goTypes,
		DependencyIndexes: file_pkg_proto_gateway_proto_depIdxs,
		MessageInfos:      file_pkg_proto_gateway_proto_msgTypes,
	}.Build()
	File_pkg_proto_gateway_proto = out.File
	file_pkg_proto_gateway_proto_rawDesc = nil
	file_pkg_proto_gateway_proto_goTypes = nil
	file_pkg_proto_gateway_proto_depIdxs = nil
}
