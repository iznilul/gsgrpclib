// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//edition = "2023";

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: proto/business/business-rpc.proto

package business_rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdentityLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserName     string   `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	TaskName     string   `protobuf:"bytes,4,opt,name=taskName,proto3" json:"taskName,omitempty"`
	Comment      []string `protobuf:"bytes,5,rep,name=comment,proto3" json:"comment,omitempty"`
	FinishedTime string   `protobuf:"bytes,6,opt,name=finishedTime,proto3" json:"finishedTime,omitempty"`
	IsPass       bool     `protobuf:"varint,7,opt,name=isPass,proto3" json:"isPass,omitempty"`
}

func (x *IdentityLink) Reset() {
	*x = IdentityLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_business_business_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityLink) ProtoMessage() {}

func (x *IdentityLink) ProtoReflect() protoreflect.Message {
	mi := &file_proto_business_business_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityLink.ProtoReflect.Descriptor instead.
func (*IdentityLink) Descriptor() ([]byte, []int) {
	return file_proto_business_business_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityLink) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *IdentityLink) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *IdentityLink) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *IdentityLink) GetComment() []string {
	if x != nil {
		return x.Comment
	}
	return nil
}

func (x *IdentityLink) GetFinishedTime() string {
	if x != nil {
		return x.FinishedTime
	}
	return ""
}

func (x *IdentityLink) GetIsPass() bool {
	if x != nil {
		return x.IsPass
	}
	return false
}

type ProcInst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SerialNumber string   `protobuf:"bytes,2,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	ProcDefName  string   `protobuf:"bytes,3,opt,name=procDefName,proto3" json:"procDefName,omitempty"`
	Title        string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	StartTime    string   `protobuf:"bytes,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime      string   `protobuf:"bytes,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
	GlobalVar    []string `protobuf:"bytes,7,rep,name=globalVar,proto3" json:"globalVar,omitempty"`
	// map<string, google.protobuf.Any> globalVar = 7;
	Variable []string `protobuf:"bytes,8,rep,name=variable,proto3" json:"variable,omitempty"`
	// map<string, google.protobuf.Any> variable = 8;
	FirstOperator  string `protobuf:"bytes,10,opt,name=firstOperator,proto3" json:"firstOperator,omitempty"`
	SecondOperator string `protobuf:"bytes,11,opt,name=secondOperator,proto3" json:"secondOperator,omitempty"`
	SaleUserName   string `protobuf:"bytes,12,opt,name=saleUserName,proto3" json:"saleUserName,omitempty"`
	IsFinished     bool   `protobuf:"varint,13,opt,name=isFinished,proto3" json:"isFinished,omitempty"`
}

func (x *ProcInst) Reset() {
	*x = ProcInst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_business_business_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcInst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcInst) ProtoMessage() {}

func (x *ProcInst) ProtoReflect() protoreflect.Message {
	mi := &file_proto_business_business_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcInst.ProtoReflect.Descriptor instead.
func (*ProcInst) Descriptor() ([]byte, []int) {
	return file_proto_business_business_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ProcInst) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProcInst) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *ProcInst) GetProcDefName() string {
	if x != nil {
		return x.ProcDefName
	}
	return ""
}

func (x *ProcInst) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProcInst) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ProcInst) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *ProcInst) GetGlobalVar() []string {
	if x != nil {
		return x.GlobalVar
	}
	return nil
}

func (x *ProcInst) GetVariable() []string {
	if x != nil {
		return x.Variable
	}
	return nil
}

func (x *ProcInst) GetFirstOperator() string {
	if x != nil {
		return x.FirstOperator
	}
	return ""
}

func (x *ProcInst) GetSecondOperator() string {
	if x != nil {
		return x.SecondOperator
	}
	return ""
}

func (x *ProcInst) GetSaleUserName() string {
	if x != nil {
		return x.SaleUserName
	}
	return ""
}

func (x *ProcInst) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

type NowTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TaskName     string   `protobuf:"bytes,3,opt,name=taskName,proto3" json:"taskName,omitempty"`
	Form         []string `protobuf:"bytes,4,rep,name=form,proto3" json:"form,omitempty"`
	AssigneeName string   `protobuf:"bytes,7,opt,name=assigneeName,proto3" json:"assigneeName,omitempty"`
	CreateTime   string   `protobuf:"bytes,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
}

func (x *NowTask) Reset() {
	*x = NowTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_business_business_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NowTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NowTask) ProtoMessage() {}

func (x *NowTask) ProtoReflect() protoreflect.Message {
	mi := &file_proto_business_business_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NowTask.ProtoReflect.Descriptor instead.
func (*NowTask) Descriptor() ([]byte, []int) {
	return file_proto_business_business_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *NowTask) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NowTask) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *NowTask) GetForm() []string {
	if x != nil {
		return x.Form
	}
	return nil
}

func (x *NowTask) GetAssigneeName() string {
	if x != nil {
		return x.AssigneeName
	}
	return ""
}

func (x *NowTask) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type ProcInstVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcInst         *ProcInst       `protobuf:"bytes,1,opt,name=procInst,proto3" json:"procInst,omitempty"`
	IdentityLinkList []*IdentityLink `protobuf:"bytes,2,rep,name=identityLinkList,proto3" json:"identityLinkList,omitempty"`
	NowTask          *NowTask        `protobuf:"bytes,3,opt,name=nowTask,proto3,oneof" json:"nowTask,omitempty"`
}

func (x *ProcInstVO) Reset() {
	*x = ProcInstVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_business_business_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcInstVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcInstVO) ProtoMessage() {}

func (x *ProcInstVO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_business_business_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcInstVO.ProtoReflect.Descriptor instead.
func (*ProcInstVO) Descriptor() ([]byte, []int) {
	return file_proto_business_business_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *ProcInstVO) GetProcInst() *ProcInst {
	if x != nil {
		return x.ProcInst
	}
	return nil
}

func (x *ProcInstVO) GetIdentityLinkList() []*IdentityLink {
	if x != nil {
		return x.IdentityLinkList
	}
	return nil
}

func (x *ProcInstVO) GetNowTask() *NowTask {
	if x != nil {
		return x.NowTask
	}
	return nil
}

type RequestAO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     *anypb.Any   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	DataList []*anypb.Any `protobuf:"bytes,2,rep,name=dataList,proto3" json:"dataList,omitempty"`
	Map      *anypb.Any   `protobuf:"bytes,3,opt,name=map,proto3" json:"map,omitempty"`
	MapList  []*anypb.Any `protobuf:"bytes,4,rep,name=mapList,proto3" json:"mapList,omitempty"`
}

func (x *RequestAO) Reset() {
	*x = RequestAO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_business_business_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAO) ProtoMessage() {}

func (x *RequestAO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_business_business_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAO.ProtoReflect.Descriptor instead.
func (*RequestAO) Descriptor() ([]byte, []int) {
	return file_proto_business_business_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *RequestAO) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RequestAO) GetDataList() []*anypb.Any {
	if x != nil {
		return x.DataList
	}
	return nil
}

func (x *RequestAO) GetMap() *anypb.Any {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *RequestAO) GetMapList() []*anypb.Any {
	if x != nil {
		return x.MapList
	}
	return nil
}

type ResponseVO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// google.protobuf.Any data = 3;
	Data           *anypb.Any    `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	DataList       []*anypb.Any  `protobuf:"bytes,4,rep,name=dataList,proto3" json:"dataList,omitempty"`
	Map            *anypb.Any    `protobuf:"bytes,5,opt,name=map,proto3" json:"map,omitempty"`
	MapList        []*anypb.Any  `protobuf:"bytes,6,rep,name=mapList,proto3" json:"mapList,omitempty"`
	ProcInstVOList []*ProcInstVO `protobuf:"bytes,7,rep,name=procInstVOList,proto3" json:"procInstVOList,omitempty"`
}

func (x *ResponseVO) Reset() {
	*x = ResponseVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_business_business_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseVO) ProtoMessage() {}

func (x *ResponseVO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_business_business_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseVO.ProtoReflect.Descriptor instead.
func (*ResponseVO) Descriptor() ([]byte, []int) {
	return file_proto_business_business_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseVO) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResponseVO) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ResponseVO) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResponseVO) GetDataList() []*anypb.Any {
	if x != nil {
		return x.DataList
	}
	return nil
}

func (x *ResponseVO) GetMap() *anypb.Any {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *ResponseVO) GetMapList() []*anypb.Any {
	if x != nil {
		return x.MapList
	}
	return nil
}

func (x *ResponseVO) GetProcInstVOList() []*ProcInstVO {
	if x != nil {
		return x.ProcInstVOList
	}
	return nil
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string     `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *anypb.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_business_business_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_business_business_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_proto_business_business_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_proto_business_business_rpc_proto protoreflect.FileDescriptor

var file_proto_business_business_rpc_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2d, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x72, 0x70, 0x63, 0x1a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x0c, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x22, 0xfa, 0x02, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x63, 0x49, 0x6e, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65,
	0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x63, 0x44, 0x65, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x63, 0x44, 0x65, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6c,
	0x6f, 0x62, 0x61, 0x6c, 0x56, 0x61, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x56, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x61, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x61, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x4e, 0x6f, 0x77, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x49,
	0x6e, 0x73, 0x74, 0x56, 0x4f, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x49, 0x6e, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x49, 0x6e, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x49, 0x6e, 0x73,
	0x74, 0x12, 0x3e, 0x0a, 0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x6e,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x62, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x77, 0x54, 0x61, 0x73,
	0x6b, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x6e, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x22, 0xbf, 0x01, 0x0a, 0x09,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x2e, 0x0a,
	0x07, 0x6d, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa0, 0x02,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d,
	0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x49, 0x6e,
	0x73, 0x74, 0x56, 0x4f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x49, 0x6e, 0x73, 0x74, 0x56, 0x4f,
	0x52, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x49, 0x6e, 0x73, 0x74, 0x56, 0x4f, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x48, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x8a, 0x06, 0x0a, 0x0b, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x50, 0x43, 0x12, 0x34, 0x0a, 0x0d, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x62, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x62,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00,
	0x12, 0x34, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x56, 0x4f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x63, 0x49, 0x6e, 0x73, 0x74, 0x42, 0x79, 0x43, 0x68, 0x61, 0x74, 0x49, 0x44, 0x12, 0x0f,
	0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a,
	0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56,
	0x4f, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x62,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e,
	0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x69, 0x6e, 0x69, 0x41, 0x6e, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f,
	0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x4f, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x62, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x62,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x0f, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74,
	0x12, 0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x4f, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x56, 0x4f, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x16, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f,
	0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x4f, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x0f, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x62, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x7a, 0x6e, 0x69, 0x6c, 0x75, 0x6c, 0x2f, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_business_business_rpc_proto_rawDescOnce sync.Once
	file_proto_business_business_rpc_proto_rawDescData = file_proto_business_business_rpc_proto_rawDesc
)

func file_proto_business_business_rpc_proto_rawDescGZIP() []byte {
	file_proto_business_business_rpc_proto_rawDescOnce.Do(func() {
		file_proto_business_business_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_business_business_rpc_proto_rawDescData)
	})
	return file_proto_business_business_rpc_proto_rawDescData
}

var file_proto_business_business_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_business_business_rpc_proto_goTypes = []any{
	(*IdentityLink)(nil), // 0: brpc.IdentityLink
	(*ProcInst)(nil),     // 1: brpc.ProcInst
	(*NowTask)(nil),      // 2: brpc.NowTask
	(*ProcInstVO)(nil),   // 3: brpc.ProcInstVO
	(*RequestAO)(nil),    // 4: brpc.RequestAO
	(*ResponseVO)(nil),   // 5: brpc.ResponseVO
	(*KeyValue)(nil),     // 6: brpc.KeyValue
	(*anypb.Any)(nil),    // 7: google.protobuf.Any
}
var file_proto_business_business_rpc_proto_depIdxs = []int32{
	1,  // 0: brpc.ProcInstVO.procInst:type_name -> brpc.ProcInst
	0,  // 1: brpc.ProcInstVO.identityLinkList:type_name -> brpc.IdentityLink
	2,  // 2: brpc.ProcInstVO.nowTask:type_name -> brpc.NowTask
	7,  // 3: brpc.RequestAO.data:type_name -> google.protobuf.Any
	7,  // 4: brpc.RequestAO.dataList:type_name -> google.protobuf.Any
	7,  // 5: brpc.RequestAO.map:type_name -> google.protobuf.Any
	7,  // 6: brpc.RequestAO.mapList:type_name -> google.protobuf.Any
	7,  // 7: brpc.ResponseVO.data:type_name -> google.protobuf.Any
	7,  // 8: brpc.ResponseVO.dataList:type_name -> google.protobuf.Any
	7,  // 9: brpc.ResponseVO.map:type_name -> google.protobuf.Any
	7,  // 10: brpc.ResponseVO.mapList:type_name -> google.protobuf.Any
	3,  // 11: brpc.ResponseVO.procInstVOList:type_name -> brpc.ProcInstVO
	7,  // 12: brpc.KeyValue.value:type_name -> google.protobuf.Any
	4,  // 13: brpc.BusinessRPC.GenerateOrder:input_type -> brpc.RequestAO
	4,  // 14: brpc.BusinessRPC.FindOrderInfo:input_type -> brpc.RequestAO
	4,  // 15: brpc.BusinessRPC.FindProcInstByChatID:input_type -> brpc.RequestAO
	4,  // 16: brpc.BusinessRPC.UpdateTrack:input_type -> brpc.RequestAO
	4,  // 17: brpc.BusinessRPC.FindContentNoList:input_type -> brpc.RequestAO
	4,  // 18: brpc.BusinessRPC.UpdateNotifyMode:input_type -> brpc.RequestAO
	4,  // 19: brpc.BusinessRPC.UpdateCustomerRemark:input_type -> brpc.RequestAO
	4,  // 20: brpc.BusinessRPC.UpdateMiniAndAccount:input_type -> brpc.RequestAO
	4,  // 21: brpc.BusinessRPC.QueryIndicatorCount:input_type -> brpc.RequestAO
	4,  // 22: brpc.BusinessRPC.QueryIndicatorDetail:input_type -> brpc.RequestAO
	4,  // 23: brpc.BusinessRPC.SyncOrderProfit:input_type -> brpc.RequestAO
	4,  // 24: brpc.BusinessRPC.CalculateUserIndicator:input_type -> brpc.RequestAO
	4,  // 25: brpc.BusinessRPC.QueryIndicatorCountInBatch:input_type -> brpc.RequestAO
	5,  // 26: brpc.BusinessRPC.GenerateOrder:output_type -> brpc.ResponseVO
	5,  // 27: brpc.BusinessRPC.FindOrderInfo:output_type -> brpc.ResponseVO
	5,  // 28: brpc.BusinessRPC.FindProcInstByChatID:output_type -> brpc.ResponseVO
	5,  // 29: brpc.BusinessRPC.UpdateTrack:output_type -> brpc.ResponseVO
	5,  // 30: brpc.BusinessRPC.FindContentNoList:output_type -> brpc.ResponseVO
	5,  // 31: brpc.BusinessRPC.UpdateNotifyMode:output_type -> brpc.ResponseVO
	5,  // 32: brpc.BusinessRPC.UpdateCustomerRemark:output_type -> brpc.ResponseVO
	5,  // 33: brpc.BusinessRPC.UpdateMiniAndAccount:output_type -> brpc.ResponseVO
	5,  // 34: brpc.BusinessRPC.QueryIndicatorCount:output_type -> brpc.ResponseVO
	5,  // 35: brpc.BusinessRPC.QueryIndicatorDetail:output_type -> brpc.ResponseVO
	5,  // 36: brpc.BusinessRPC.SyncOrderProfit:output_type -> brpc.ResponseVO
	5,  // 37: brpc.BusinessRPC.CalculateUserIndicator:output_type -> brpc.ResponseVO
	5,  // 38: brpc.BusinessRPC.QueryIndicatorCountInBatch:output_type -> brpc.ResponseVO
	26, // [26:39] is the sub-list for method output_type
	13, // [13:26] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_business_business_rpc_proto_init() }
func file_proto_business_business_rpc_proto_init() {
	if File_proto_business_business_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_business_business_rpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IdentityLink); i {
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
		file_proto_business_business_rpc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProcInst); i {
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
		file_proto_business_business_rpc_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NowTask); i {
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
		file_proto_business_business_rpc_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ProcInstVO); i {
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
		file_proto_business_business_rpc_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RequestAO); i {
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
		file_proto_business_business_rpc_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ResponseVO); i {
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
		file_proto_business_business_rpc_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*KeyValue); i {
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
	file_proto_business_business_rpc_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_business_business_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_business_business_rpc_proto_goTypes,
		DependencyIndexes: file_proto_business_business_rpc_proto_depIdxs,
		MessageInfos:      file_proto_business_business_rpc_proto_msgTypes,
	}.Build()
	File_proto_business_business_rpc_proto = out.File
	file_proto_business_business_rpc_proto_rawDesc = nil
	file_proto_business_business_rpc_proto_goTypes = nil
	file_proto_business_business_rpc_proto_depIdxs = nil
}
