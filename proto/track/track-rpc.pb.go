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
// source: proto/track/track-rpc.proto

package track_rpc

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

type NoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoParam) Reset() {
	*x = NoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_track_track_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoParam) ProtoMessage() {}

func (x *NoParam) ProtoReflect() protoreflect.Message {
	mi := &file_proto_track_track_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoParam.ProtoReflect.Descriptor instead.
func (*NoParam) Descriptor() ([]byte, []int) {
	return file_proto_track_track_rpc_proto_rawDescGZIP(), []int{0}
}

type ContentNoAO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentNo string `protobuf:"bytes,1,opt,name=contentNo,proto3" json:"contentNo,omitempty"`
}

func (x *ContentNoAO) Reset() {
	*x = ContentNoAO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_track_track_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentNoAO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentNoAO) ProtoMessage() {}

func (x *ContentNoAO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_track_track_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentNoAO.ProtoReflect.Descriptor instead.
func (*ContentNoAO) Descriptor() ([]byte, []int) {
	return file_proto_track_track_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ContentNoAO) GetContentNo() string {
	if x != nil {
		return x.ContentNo
	}
	return ""
}

type SyncAO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentNoList    []string `protobuf:"bytes,1,rep,name=contentNoList,proto3" json:"contentNoList,omitempty"`
	StartStation     string   `protobuf:"bytes,2,opt,name=startStation,proto3" json:"startStation,omitempty"`
	StartStationName string   `protobuf:"bytes,3,opt,name=startStationName,proto3" json:"startStationName,omitempty"`
	EndStation       string   `protobuf:"bytes,4,opt,name=endStation,proto3" json:"endStation,omitempty"`
	EndStationName   string   `protobuf:"bytes,5,opt,name=endStationName,proto3" json:"endStationName,omitempty"`
	SerialNumber     string   `protobuf:"bytes,6,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
}

func (x *SyncAO) Reset() {
	*x = SyncAO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_track_track_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncAO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncAO) ProtoMessage() {}

func (x *SyncAO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_track_track_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncAO.ProtoReflect.Descriptor instead.
func (*SyncAO) Descriptor() ([]byte, []int) {
	return file_proto_track_track_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *SyncAO) GetContentNoList() []string {
	if x != nil {
		return x.ContentNoList
	}
	return nil
}

func (x *SyncAO) GetStartStation() string {
	if x != nil {
		return x.StartStation
	}
	return ""
}

func (x *SyncAO) GetStartStationName() string {
	if x != nil {
		return x.StartStationName
	}
	return ""
}

func (x *SyncAO) GetEndStation() string {
	if x != nil {
		return x.EndStation
	}
	return ""
}

func (x *SyncAO) GetEndStationName() string {
	if x != nil {
		return x.EndStationName
	}
	return ""
}

func (x *SyncAO) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

type QueryAO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table string `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Field string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *QueryAO) Reset() {
	*x = QueryAO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_track_track_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryAO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryAO) ProtoMessage() {}

func (x *QueryAO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_track_track_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryAO.ProtoReflect.Descriptor instead.
func (*QueryAO) Descriptor() ([]byte, []int) {
	return file_proto_track_track_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *QueryAO) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *QueryAO) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *QueryAO) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type FindStationAO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value       string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	CurrentPage int32  `protobuf:"varint,2,opt,name=currentPage,proto3" json:"currentPage,omitempty"`
	PageSize    int32  `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *FindStationAO) Reset() {
	*x = FindStationAO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_track_track_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStationAO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStationAO) ProtoMessage() {}

func (x *FindStationAO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_track_track_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStationAO.ProtoReflect.Descriptor instead.
func (*FindStationAO) Descriptor() ([]byte, []int) {
	return file_proto_track_track_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *FindStationAO) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *FindStationAO) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *FindStationAO) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
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
		mi := &file_proto_track_track_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAO) ProtoMessage() {}

func (x *RequestAO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_track_track_rpc_proto_msgTypes[5]
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
	return file_proto_track_track_rpc_proto_rawDescGZIP(), []int{5}
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
	Data     *anypb.Any   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	DataList []*anypb.Any `protobuf:"bytes,4,rep,name=dataList,proto3" json:"dataList,omitempty"`
	Map      *anypb.Any   `protobuf:"bytes,5,opt,name=map,proto3" json:"map,omitempty"`
	MapList  []*anypb.Any `protobuf:"bytes,6,rep,name=mapList,proto3" json:"mapList,omitempty"`
}

func (x *ResponseVO) Reset() {
	*x = ResponseVO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_track_track_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseVO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseVO) ProtoMessage() {}

func (x *ResponseVO) ProtoReflect() protoreflect.Message {
	mi := &file_proto_track_track_rpc_proto_msgTypes[6]
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
	return file_proto_track_track_rpc_proto_rawDescGZIP(), []int{6}
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
		mi := &file_proto_track_track_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_proto_track_track_rpc_proto_msgTypes[7]
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
	return file_proto_track_track_rpc_proto_rawDescGZIP(), []int{7}
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

var File_proto_track_track_rpc_proto protoreflect.FileDescriptor

var file_proto_track_track_rpc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2f, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x2d, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74,
	0x72, 0x70, 0x63, 0x1a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22,
	0x2b, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x41, 0x4f, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x22, 0xea, 0x01, 0x0a,
	0x06, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x4f, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x4e, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x0e, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x07, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x41, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x63, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x4f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x09,
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
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xe6, 0x01,
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
	0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x32, 0xdd, 0x03, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x50, 0x43, 0x12, 0x35, 0x0a,
	0x0e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x0f, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f,
	0x1a, 0x10, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x4f, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x79, 0x6e, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0f,
	0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x0f, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f,
	0x1a, 0x10, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x4f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x67, 0x6f,
	0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x12, 0x0f, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41,
	0x4f, 0x1a, 0x10, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x56, 0x4f, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10,
	0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12,
	0x0f, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f,
	0x1a, 0x10, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x56, 0x4f, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x23, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x53, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x74, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x4f, 0x1a, 0x10, 0x2e, 0x74,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x4f, 0x22, 0x00,
	0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x7a, 0x6e, 0x69, 0x6c, 0x75, 0x6c, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_track_track_rpc_proto_rawDescOnce sync.Once
	file_proto_track_track_rpc_proto_rawDescData = file_proto_track_track_rpc_proto_rawDesc
)

func file_proto_track_track_rpc_proto_rawDescGZIP() []byte {
	file_proto_track_track_rpc_proto_rawDescOnce.Do(func() {
		file_proto_track_track_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_track_track_rpc_proto_rawDescData)
	})
	return file_proto_track_track_rpc_proto_rawDescData
}

var file_proto_track_track_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_track_track_rpc_proto_goTypes = []any{
	(*NoParam)(nil),       // 0: trpc.NoParam
	(*ContentNoAO)(nil),   // 1: trpc.ContentNoAO
	(*SyncAO)(nil),        // 2: trpc.SyncAO
	(*QueryAO)(nil),       // 3: trpc.QueryAO
	(*FindStationAO)(nil), // 4: trpc.FindStationAO
	(*RequestAO)(nil),     // 5: trpc.RequestAO
	(*ResponseVO)(nil),    // 6: trpc.ResponseVO
	(*KeyValue)(nil),      // 7: trpc.KeyValue
	(*anypb.Any)(nil),     // 8: google.protobuf.Any
}
var file_proto_track_track_rpc_proto_depIdxs = []int32{
	8,  // 0: trpc.RequestAO.data:type_name -> google.protobuf.Any
	8,  // 1: trpc.RequestAO.dataList:type_name -> google.protobuf.Any
	8,  // 2: trpc.RequestAO.map:type_name -> google.protobuf.Any
	8,  // 3: trpc.RequestAO.mapList:type_name -> google.protobuf.Any
	8,  // 4: trpc.ResponseVO.data:type_name -> google.protobuf.Any
	8,  // 5: trpc.ResponseVO.dataList:type_name -> google.protobuf.Any
	8,  // 6: trpc.ResponseVO.map:type_name -> google.protobuf.Any
	8,  // 7: trpc.ResponseVO.mapList:type_name -> google.protobuf.Any
	8,  // 8: trpc.KeyValue.value:type_name -> google.protobuf.Any
	5,  // 9: trpc.TrackRPC.QueryFieldList:input_type -> trpc.RequestAO
	5,  // 10: trpc.TrackRPC.SendSyncRequest:input_type -> trpc.RequestAO
	5,  // 11: trpc.TrackRPC.FindStationList:input_type -> trpc.RequestAO
	5,  // 12: trpc.TrackRPC.FindOngoingTrackList:input_type -> trpc.RequestAO
	5,  // 13: trpc.TrackRPC.FinishTrack:input_type -> trpc.RequestAO
	5,  // 14: trpc.TrackRPC.FindHistoryTrackList:input_type -> trpc.RequestAO
	5,  // 15: trpc.TrackRPC.SyncTrack:input_type -> trpc.RequestAO
	5,  // 16: trpc.TrackRPC.FindContentNoListBySerialNumberList:input_type -> trpc.RequestAO
	6,  // 17: trpc.TrackRPC.QueryFieldList:output_type -> trpc.ResponseVO
	6,  // 18: trpc.TrackRPC.SendSyncRequest:output_type -> trpc.ResponseVO
	6,  // 19: trpc.TrackRPC.FindStationList:output_type -> trpc.ResponseVO
	6,  // 20: trpc.TrackRPC.FindOngoingTrackList:output_type -> trpc.ResponseVO
	6,  // 21: trpc.TrackRPC.FinishTrack:output_type -> trpc.ResponseVO
	6,  // 22: trpc.TrackRPC.FindHistoryTrackList:output_type -> trpc.ResponseVO
	6,  // 23: trpc.TrackRPC.SyncTrack:output_type -> trpc.ResponseVO
	6,  // 24: trpc.TrackRPC.FindContentNoListBySerialNumberList:output_type -> trpc.ResponseVO
	17, // [17:25] is the sub-list for method output_type
	9,  // [9:17] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_track_track_rpc_proto_init() }
func file_proto_track_track_rpc_proto_init() {
	if File_proto_track_track_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_track_track_rpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NoParam); i {
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
		file_proto_track_track_rpc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ContentNoAO); i {
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
		file_proto_track_track_rpc_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SyncAO); i {
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
		file_proto_track_track_rpc_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*QueryAO); i {
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
		file_proto_track_track_rpc_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FindStationAO); i {
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
		file_proto_track_track_rpc_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
		file_proto_track_track_rpc_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
		file_proto_track_track_rpc_proto_msgTypes[7].Exporter = func(v any, i int) any {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_track_track_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_track_track_rpc_proto_goTypes,
		DependencyIndexes: file_proto_track_track_rpc_proto_depIdxs,
		MessageInfos:      file_proto_track_track_rpc_proto_msgTypes,
	}.Build()
	File_proto_track_track_rpc_proto = out.File
	file_proto_track_track_rpc_proto_rawDesc = nil
	file_proto_track_track_rpc_proto_goTypes = nil
	file_proto_track_track_rpc_proto_depIdxs = nil
}
