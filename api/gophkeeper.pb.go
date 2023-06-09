// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: api/gophkeeper.proto

package api

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

// The request message for register and login
type RegisterLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterLoginReq) Reset() {
	*x = RegisterLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterLoginReq) ProtoMessage() {}

func (x *RegisterLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterLoginReq.ProtoReflect.Descriptor instead.
func (*RegisterLoginReq) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterLoginReq) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *RegisterLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// The response message containing the token
type RegisterLoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Success bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RegisterLoginResp) Reset() {
	*x = RegisterLoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterLoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterLoginResp) ProtoMessage() {}

func (x *RegisterLoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterLoginResp.ProtoReflect.Descriptor instead.
func (*RegisterLoginResp) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterLoginResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterLoginResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// The request message for save data
type SaveDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SaveDataReq) Reset() {
	*x = SaveDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveDataReq) ProtoMessage() {}

func (x *SaveDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveDataReq.ProtoReflect.Descriptor instead.
func (*SaveDataReq) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{2}
}

func (x *SaveDataReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SaveDataReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SaveDataReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// The request message for save login
type SaveLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Login string `protobuf:"bytes,3,opt,name=login,proto3" json:"login,omitempty"`
	Pass  string `protobuf:"bytes,4,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *SaveLoginReq) Reset() {
	*x = SaveLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveLoginReq) ProtoMessage() {}

func (x *SaveLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveLoginReq.ProtoReflect.Descriptor instead.
func (*SaveLoginReq) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{3}
}

func (x *SaveLoginReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SaveLoginReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SaveLoginReq) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *SaveLoginReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

// The request message for save text
type SaveTextReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text  string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SaveTextReq) Reset() {
	*x = SaveTextReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveTextReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveTextReq) ProtoMessage() {}

func (x *SaveTextReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveTextReq.ProtoReflect.Descriptor instead.
func (*SaveTextReq) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{4}
}

func (x *SaveTextReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SaveTextReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SaveTextReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// The request message for save bank card
type SaveBankCardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CardNo string `protobuf:"bytes,3,opt,name=cardNo,proto3" json:"cardNo,omitempty"`
	Valid  string `protobuf:"bytes,4,opt,name=valid,proto3" json:"valid,omitempty"`
	Cvv    string `protobuf:"bytes,5,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *SaveBankCardReq) Reset() {
	*x = SaveBankCardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBankCardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBankCardReq) ProtoMessage() {}

func (x *SaveBankCardReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBankCardReq.ProtoReflect.Descriptor instead.
func (*SaveBankCardReq) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{5}
}

func (x *SaveBankCardReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SaveBankCardReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SaveBankCardReq) GetCardNo() string {
	if x != nil {
		return x.CardNo
	}
	return ""
}

func (x *SaveBankCardReq) GetValid() string {
	if x != nil {
		return x.Valid
	}
	return ""
}

func (x *SaveBankCardReq) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

// The response message containing data id
type SaveDataResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveDataResp) Reset() {
	*x = SaveDataResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveDataResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveDataResp) ProtoMessage() {}

func (x *SaveDataResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveDataResp.ProtoReflect.Descriptor instead.
func (*SaveDataResp) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{6}
}

func (x *SaveDataResp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The request message for list data
type ListDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListDataReq) Reset() {
	*x = ListDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDataReq) ProtoMessage() {}

func (x *ListDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDataReq.ProtoReflect.Descriptor instead.
func (*ListDataReq) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{7}
}

func (x *ListDataReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// The response message for list data
type ListDataResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ListData `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListDataResp) Reset() {
	*x = ListDataResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDataResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDataResp) ProtoMessage() {}

func (x *ListDataResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDataResp.ProtoReflect.Descriptor instead.
func (*ListDataResp) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{8}
}

func (x *ListDataResp) GetItems() []*ListData {
	if x != nil {
		return x.Items
	}
	return nil
}

// The item for ListDataResp
type ListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Type  string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ListData) Reset() {
	*x = ListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListData) ProtoMessage() {}

func (x *ListData) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListData.ProtoReflect.Descriptor instead.
func (*ListData) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{9}
}

func (x *ListData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

// The request message for GetData
type GetDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	DataId int32  `protobuf:"varint,2,opt,name=dataId,proto3" json:"dataId,omitempty"`
}

func (x *GetDataReq) Reset() {
	*x = GetDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataReq) ProtoMessage() {}

func (x *GetDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataReq.ProtoReflect.Descriptor instead.
func (*GetDataReq) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{10}
}

func (x *GetDataReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetDataReq) GetDataId() int32 {
	if x != nil {
		return x.DataId
	}
	return 0
}

// The response message for GetData
type GetDataResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetDataResp) Reset() {
	*x = GetDataResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataResp) ProtoMessage() {}

func (x *GetDataResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataResp.ProtoReflect.Descriptor instead.
func (*GetDataResp) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{11}
}

func (x *GetDataResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// The response message for GetLogin
type GetLoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Pass  string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
}

func (x *GetLoginResp) Reset() {
	*x = GetLoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginResp) ProtoMessage() {}

func (x *GetLoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginResp.ProtoReflect.Descriptor instead.
func (*GetLoginResp) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{12}
}

func (x *GetLoginResp) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetLoginResp) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

// The response message for GetText
type GetTextResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *GetTextResp) Reset() {
	*x = GetTextResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gophkeeper_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextResp) ProtoMessage() {}

func (x *GetTextResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_gophkeeper_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextResp.ProtoReflect.Descriptor instead.
func (*GetTextResp) Descriptor() ([]byte, []int) {
	return file_api_gophkeeper_proto_rawDescGZIP(), []int{13}
}

func (x *GetTextResp) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_api_gophkeeper_proto protoreflect.FileDescriptor

var file_api_gophkeeper_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x22, 0x44, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4d, 0x0a,
	0x0b, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x64, 0x0a, 0x0c,
	0x53, 0x61, 0x76, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x73, 0x73, 0x22, 0x4d, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x22, 0x7d, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x76,
	0x22, 0x1e, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x23, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x61, 0x74, 0x61, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x61, 0x74,
	0x61, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73,
	0x22, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x32, 0x96, 0x05, 0x0a, 0x0a, 0x47, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x12, 0x47, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x67,
	0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x44, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3f, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x45, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x67,
	0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x70, 0x68,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x3c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x23, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x69, 0x73, 0x6b,
	0x38, 0x34, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_gophkeeper_proto_rawDescOnce sync.Once
	file_api_gophkeeper_proto_rawDescData = file_api_gophkeeper_proto_rawDesc
)

func file_api_gophkeeper_proto_rawDescGZIP() []byte {
	file_api_gophkeeper_proto_rawDescOnce.Do(func() {
		file_api_gophkeeper_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gophkeeper_proto_rawDescData)
	})
	return file_api_gophkeeper_proto_rawDescData
}

var file_api_gophkeeper_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_gophkeeper_proto_goTypes = []interface{}{
	(*RegisterLoginReq)(nil),  // 0: gophkeeper.RegisterLoginReq
	(*RegisterLoginResp)(nil), // 1: gophkeeper.RegisterLoginResp
	(*SaveDataReq)(nil),       // 2: gophkeeper.SaveDataReq
	(*SaveLoginReq)(nil),      // 3: gophkeeper.SaveLoginReq
	(*SaveTextReq)(nil),       // 4: gophkeeper.SaveTextReq
	(*SaveBankCardReq)(nil),   // 5: gophkeeper.SaveBankCardReq
	(*SaveDataResp)(nil),      // 6: gophkeeper.SaveDataResp
	(*ListDataReq)(nil),       // 7: gophkeeper.ListDataReq
	(*ListDataResp)(nil),      // 8: gophkeeper.ListDataResp
	(*ListData)(nil),          // 9: gophkeeper.ListData
	(*GetDataReq)(nil),        // 10: gophkeeper.GetDataReq
	(*GetDataResp)(nil),       // 11: gophkeeper.GetDataResp
	(*GetLoginResp)(nil),      // 12: gophkeeper.GetLoginResp
	(*GetTextResp)(nil),       // 13: gophkeeper.GetTextResp
}
var file_api_gophkeeper_proto_depIdxs = []int32{
	9,  // 0: gophkeeper.ListDataResp.items:type_name -> gophkeeper.ListData
	0,  // 1: gophkeeper.GophKeeper.Register:input_type -> gophkeeper.RegisterLoginReq
	0,  // 2: gophkeeper.GophKeeper.Login:input_type -> gophkeeper.RegisterLoginReq
	2,  // 3: gophkeeper.GophKeeper.SaveData:input_type -> gophkeeper.SaveDataReq
	3,  // 4: gophkeeper.GophKeeper.SaveLogin:input_type -> gophkeeper.SaveLoginReq
	4,  // 5: gophkeeper.GophKeeper.SaveText:input_type -> gophkeeper.SaveTextReq
	5,  // 6: gophkeeper.GophKeeper.SaveBankCard:input_type -> gophkeeper.SaveBankCardReq
	7,  // 7: gophkeeper.GophKeeper.ListData:input_type -> gophkeeper.ListDataReq
	10, // 8: gophkeeper.GophKeeper.GetData:input_type -> gophkeeper.GetDataReq
	10, // 9: gophkeeper.GophKeeper.GetLogin:input_type -> gophkeeper.GetDataReq
	10, // 10: gophkeeper.GophKeeper.GetText:input_type -> gophkeeper.GetDataReq
	1,  // 11: gophkeeper.GophKeeper.Register:output_type -> gophkeeper.RegisterLoginResp
	1,  // 12: gophkeeper.GophKeeper.Login:output_type -> gophkeeper.RegisterLoginResp
	6,  // 13: gophkeeper.GophKeeper.SaveData:output_type -> gophkeeper.SaveDataResp
	6,  // 14: gophkeeper.GophKeeper.SaveLogin:output_type -> gophkeeper.SaveDataResp
	6,  // 15: gophkeeper.GophKeeper.SaveText:output_type -> gophkeeper.SaveDataResp
	6,  // 16: gophkeeper.GophKeeper.SaveBankCard:output_type -> gophkeeper.SaveDataResp
	8,  // 17: gophkeeper.GophKeeper.ListData:output_type -> gophkeeper.ListDataResp
	11, // 18: gophkeeper.GophKeeper.GetData:output_type -> gophkeeper.GetDataResp
	12, // 19: gophkeeper.GophKeeper.GetLogin:output_type -> gophkeeper.GetLoginResp
	13, // 20: gophkeeper.GophKeeper.GetText:output_type -> gophkeeper.GetTextResp
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_gophkeeper_proto_init() }
func file_api_gophkeeper_proto_init() {
	if File_api_gophkeeper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gophkeeper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterLoginReq); i {
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
		file_api_gophkeeper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterLoginResp); i {
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
		file_api_gophkeeper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveDataReq); i {
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
		file_api_gophkeeper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveLoginReq); i {
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
		file_api_gophkeeper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveTextReq); i {
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
		file_api_gophkeeper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveBankCardReq); i {
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
		file_api_gophkeeper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveDataResp); i {
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
		file_api_gophkeeper_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDataReq); i {
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
		file_api_gophkeeper_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDataResp); i {
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
		file_api_gophkeeper_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListData); i {
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
		file_api_gophkeeper_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataReq); i {
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
		file_api_gophkeeper_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataResp); i {
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
		file_api_gophkeeper_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginResp); i {
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
		file_api_gophkeeper_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextResp); i {
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
			RawDescriptor: file_api_gophkeeper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_gophkeeper_proto_goTypes,
		DependencyIndexes: file_api_gophkeeper_proto_depIdxs,
		MessageInfos:      file_api_gophkeeper_proto_msgTypes,
	}.Build()
	File_api_gophkeeper_proto = out.File
	file_api_gophkeeper_proto_rawDesc = nil
	file_api_gophkeeper_proto_goTypes = nil
	file_api_gophkeeper_proto_depIdxs = nil
}
