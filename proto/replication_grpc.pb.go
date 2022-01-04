// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReplicationClient is the client API for Replication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicationClient interface {
	Increment(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReplicaReply, error)
}

type replicationClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationClient(cc grpc.ClientConnInterface) ReplicationClient {
	return &replicationClient{cc}
}

func (c *replicationClient) Increment(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ReplicaReply, error) {
	out := new(ReplicaReply)
	err := c.cc.Invoke(ctx, "/Replication/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServer is the server API for Replication service.
// All implementations must embed UnimplementedReplicationServer
// for forward compatibility
type ReplicationServer interface {
	Increment(context.Context, *Empty) (*ReplicaReply, error)
	mustEmbedUnimplementedReplicationServer()
}

// UnimplementedReplicationServer must be embedded to have forward compatible implementations.
type UnimplementedReplicationServer struct {
}

func (UnimplementedReplicationServer) Increment(context.Context, *Empty) (*ReplicaReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedReplicationServer) mustEmbedUnimplementedReplicationServer() {}

// UnsafeReplicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServer will
// result in compilation errors.
type UnsafeReplicationServer interface {
	mustEmbedUnimplementedReplicationServer()
}

func RegisterReplicationServer(s grpc.ServiceRegistrar, srv ReplicationServer) {
	s.RegisterService(&Replication_ServiceDesc, srv)
}

func _Replication_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Replication/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServer).Increment(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Replication_ServiceDesc is the grpc.ServiceDesc for Replication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Replication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Replication",
	HandlerType: (*ReplicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Increment",
			Handler:    _Replication_Increment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/replication.proto",
}

// ElectionClient is the client API for Election service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectionClient interface {
	Election(ctx context.Context, in *ElectionMessage, opts ...grpc.CallOption) (*ElectionMessage, error)
	Elected(ctx context.Context, in *ElectedMessage, opts ...grpc.CallOption) (*Empty, error)
	Heartbeat(ctx context.Context, in *HeartbeatMessage, opts ...grpc.CallOption) (*Empty, error)
}

type electionClient struct {
	cc grpc.ClientConnInterface
}

func NewElectionClient(cc grpc.ClientConnInterface) ElectionClient {
	return &electionClient{cc}
}

func (c *electionClient) Election(ctx context.Context, in *ElectionMessage, opts ...grpc.CallOption) (*ElectionMessage, error) {
	out := new(ElectionMessage)
	err := c.cc.Invoke(ctx, "/Election/Election", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionClient) Elected(ctx context.Context, in *ElectedMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Election/Elected", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionClient) Heartbeat(ctx context.Context, in *HeartbeatMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Election/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectionServer is the server API for Election service.
// All implementations must embed UnimplementedElectionServer
// for forward compatibility
type ElectionServer interface {
	Election(context.Context, *ElectionMessage) (*ElectionMessage, error)
	Elected(context.Context, *ElectedMessage) (*Empty, error)
	Heartbeat(context.Context, *HeartbeatMessage) (*Empty, error)
	mustEmbedUnimplementedElectionServer()
}

// UnimplementedElectionServer must be embedded to have forward compatible implementations.
type UnimplementedElectionServer struct {
}

func (UnimplementedElectionServer) Election(context.Context, *ElectionMessage) (*ElectionMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Election not implemented")
}
func (UnimplementedElectionServer) Elected(context.Context, *ElectedMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Elected not implemented")
}
func (UnimplementedElectionServer) Heartbeat(context.Context, *HeartbeatMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedElectionServer) mustEmbedUnimplementedElectionServer() {}

// UnsafeElectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectionServer will
// result in compilation errors.
type UnsafeElectionServer interface {
	mustEmbedUnimplementedElectionServer()
}

func RegisterElectionServer(s grpc.ServiceRegistrar, srv ElectionServer) {
	s.RegisterService(&Election_ServiceDesc, srv)
}

func _Election_Election_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).Election(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Election/Election",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).Election(ctx, req.(*ElectionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Election_Elected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).Elected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Election/Elected",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).Elected(ctx, req.(*ElectedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Election_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Election/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServer).Heartbeat(ctx, req.(*HeartbeatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Election_ServiceDesc is the grpc.ServiceDesc for Election service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Election_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Election",
	HandlerType: (*ElectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Election",
			Handler:    _Election_Election_Handler,
		},
		{
			MethodName: "Elected",
			Handler:    _Election_Elected_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Election_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/replication.proto",
}
