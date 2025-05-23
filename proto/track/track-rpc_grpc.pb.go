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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: proto/track/track-rpc.proto

package track_rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TrackRPC_QueryFieldList_FullMethodName                      = "/trpc.TrackRPC/QueryFieldList"
	TrackRPC_SendSyncRequest_FullMethodName                     = "/trpc.TrackRPC/SendSyncRequest"
	TrackRPC_FindStationList_FullMethodName                     = "/trpc.TrackRPC/FindStationList"
	TrackRPC_FindOngoingTrackList_FullMethodName                = "/trpc.TrackRPC/FindOngoingTrackList"
	TrackRPC_FinishTrack_FullMethodName                         = "/trpc.TrackRPC/FinishTrack"
	TrackRPC_FindHistoryTrackList_FullMethodName                = "/trpc.TrackRPC/FindHistoryTrackList"
	TrackRPC_SyncTrack_FullMethodName                           = "/trpc.TrackRPC/SyncTrack"
	TrackRPC_FindContentNoListBySerialNumberList_FullMethodName = "/trpc.TrackRPC/FindContentNoListBySerialNumberList"
)

// TrackRPCClient is the client API for TrackRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Interface exported by the server.
type TrackRPCClient interface {
	QueryFieldList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error)
	SendSyncRequest(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error)
	FindStationList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error)
	FindOngoingTrackList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error)
	FinishTrack(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error)
	FindHistoryTrackList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error)
	SyncTrack(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error)
	FindContentNoListBySerialNumberList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error)
}

type trackRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTrackRPCClient(cc grpc.ClientConnInterface) TrackRPCClient {
	return &trackRPCClient{cc}
}

func (c *trackRPCClient) QueryFieldList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVO)
	err := c.cc.Invoke(ctx, TrackRPC_QueryFieldList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackRPCClient) SendSyncRequest(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVO)
	err := c.cc.Invoke(ctx, TrackRPC_SendSyncRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackRPCClient) FindStationList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVO)
	err := c.cc.Invoke(ctx, TrackRPC_FindStationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackRPCClient) FindOngoingTrackList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVO)
	err := c.cc.Invoke(ctx, TrackRPC_FindOngoingTrackList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackRPCClient) FinishTrack(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVO)
	err := c.cc.Invoke(ctx, TrackRPC_FinishTrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackRPCClient) FindHistoryTrackList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVO)
	err := c.cc.Invoke(ctx, TrackRPC_FindHistoryTrackList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackRPCClient) SyncTrack(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVO)
	err := c.cc.Invoke(ctx, TrackRPC_SyncTrack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackRPCClient) FindContentNoListBySerialNumberList(ctx context.Context, in *RequestAO, opts ...grpc.CallOption) (*ResponseVO, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResponseVO)
	err := c.cc.Invoke(ctx, TrackRPC_FindContentNoListBySerialNumberList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrackRPCServer is the server API for TrackRPC service.
// All implementations must embed UnimplementedTrackRPCServer
// for forward compatibility.
//
// Interface exported by the server.
type TrackRPCServer interface {
	QueryFieldList(context.Context, *RequestAO) (*ResponseVO, error)
	SendSyncRequest(context.Context, *RequestAO) (*ResponseVO, error)
	FindStationList(context.Context, *RequestAO) (*ResponseVO, error)
	FindOngoingTrackList(context.Context, *RequestAO) (*ResponseVO, error)
	FinishTrack(context.Context, *RequestAO) (*ResponseVO, error)
	FindHistoryTrackList(context.Context, *RequestAO) (*ResponseVO, error)
	SyncTrack(context.Context, *RequestAO) (*ResponseVO, error)
	FindContentNoListBySerialNumberList(context.Context, *RequestAO) (*ResponseVO, error)
	mustEmbedUnimplementedTrackRPCServer()
}

// UnimplementedTrackRPCServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTrackRPCServer struct{}

func (UnimplementedTrackRPCServer) QueryFieldList(context.Context, *RequestAO) (*ResponseVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFieldList not implemented")
}
func (UnimplementedTrackRPCServer) SendSyncRequest(context.Context, *RequestAO) (*ResponseVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSyncRequest not implemented")
}
func (UnimplementedTrackRPCServer) FindStationList(context.Context, *RequestAO) (*ResponseVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStationList not implemented")
}
func (UnimplementedTrackRPCServer) FindOngoingTrackList(context.Context, *RequestAO) (*ResponseVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOngoingTrackList not implemented")
}
func (UnimplementedTrackRPCServer) FinishTrack(context.Context, *RequestAO) (*ResponseVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishTrack not implemented")
}
func (UnimplementedTrackRPCServer) FindHistoryTrackList(context.Context, *RequestAO) (*ResponseVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindHistoryTrackList not implemented")
}
func (UnimplementedTrackRPCServer) SyncTrack(context.Context, *RequestAO) (*ResponseVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncTrack not implemented")
}
func (UnimplementedTrackRPCServer) FindContentNoListBySerialNumberList(context.Context, *RequestAO) (*ResponseVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindContentNoListBySerialNumberList not implemented")
}
func (UnimplementedTrackRPCServer) mustEmbedUnimplementedTrackRPCServer() {}
func (UnimplementedTrackRPCServer) testEmbeddedByValue()                  {}

// UnsafeTrackRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrackRPCServer will
// result in compilation errors.
type UnsafeTrackRPCServer interface {
	mustEmbedUnimplementedTrackRPCServer()
}

func RegisterTrackRPCServer(s grpc.ServiceRegistrar, srv TrackRPCServer) {
	// If the following call pancis, it indicates UnimplementedTrackRPCServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TrackRPC_ServiceDesc, srv)
}

func _TrackRPC_QueryFieldList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackRPCServer).QueryFieldList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackRPC_QueryFieldList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackRPCServer).QueryFieldList(ctx, req.(*RequestAO))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackRPC_SendSyncRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackRPCServer).SendSyncRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackRPC_SendSyncRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackRPCServer).SendSyncRequest(ctx, req.(*RequestAO))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackRPC_FindStationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackRPCServer).FindStationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackRPC_FindStationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackRPCServer).FindStationList(ctx, req.(*RequestAO))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackRPC_FindOngoingTrackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackRPCServer).FindOngoingTrackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackRPC_FindOngoingTrackList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackRPCServer).FindOngoingTrackList(ctx, req.(*RequestAO))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackRPC_FinishTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackRPCServer).FinishTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackRPC_FinishTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackRPCServer).FinishTrack(ctx, req.(*RequestAO))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackRPC_FindHistoryTrackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackRPCServer).FindHistoryTrackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackRPC_FindHistoryTrackList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackRPCServer).FindHistoryTrackList(ctx, req.(*RequestAO))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackRPC_SyncTrack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackRPCServer).SyncTrack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackRPC_SyncTrack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackRPCServer).SyncTrack(ctx, req.(*RequestAO))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrackRPC_FindContentNoListBySerialNumberList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackRPCServer).FindContentNoListBySerialNumberList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrackRPC_FindContentNoListBySerialNumberList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackRPCServer).FindContentNoListBySerialNumberList(ctx, req.(*RequestAO))
	}
	return interceptor(ctx, in, info, handler)
}

// TrackRPC_ServiceDesc is the grpc.ServiceDesc for TrackRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrackRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trpc.TrackRPC",
	HandlerType: (*TrackRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryFieldList",
			Handler:    _TrackRPC_QueryFieldList_Handler,
		},
		{
			MethodName: "SendSyncRequest",
			Handler:    _TrackRPC_SendSyncRequest_Handler,
		},
		{
			MethodName: "FindStationList",
			Handler:    _TrackRPC_FindStationList_Handler,
		},
		{
			MethodName: "FindOngoingTrackList",
			Handler:    _TrackRPC_FindOngoingTrackList_Handler,
		},
		{
			MethodName: "FinishTrack",
			Handler:    _TrackRPC_FinishTrack_Handler,
		},
		{
			MethodName: "FindHistoryTrackList",
			Handler:    _TrackRPC_FindHistoryTrackList_Handler,
		},
		{
			MethodName: "SyncTrack",
			Handler:    _TrackRPC_SyncTrack_Handler,
		},
		{
			MethodName: "FindContentNoListBySerialNumberList",
			Handler:    _TrackRPC_FindContentNoListBySerialNumberList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/track/track-rpc.proto",
}
