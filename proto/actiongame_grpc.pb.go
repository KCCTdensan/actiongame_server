// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: proto/actiongame.proto

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

const (
	Game_Hello_FullMethodName   = "/proto.Game/Hello"
	Game_GetUser_FullMethodName = "/proto.Game/GetUser"
	Game_Join_FullMethodName    = "/proto.Game/Join"
	Game_Move_FullMethodName    = "/proto.Game/Move"
)

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameClient interface {
	// basic
	Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error)
	// global
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRes, error)
	// game
	Join(ctx context.Context, in *JoinReq, opts ...grpc.CallOption) (Game_JoinClient, error)
	Move(ctx context.Context, opts ...grpc.CallOption) (Game_MoveClient, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Hello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error) {
	out := new(HelloRes)
	err := c.cc.Invoke(ctx, Game_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserRes, error) {
	out := new(GetUserRes)
	err := c.cc.Invoke(ctx, Game_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Join(ctx context.Context, in *JoinReq, opts ...grpc.CallOption) (Game_JoinClient, error) {
	stream, err := c.cc.NewStream(ctx, &Game_ServiceDesc.Streams[0], Game_Join_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gameJoinClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Game_JoinClient interface {
	Recv() (*UserConn, error)
	grpc.ClientStream
}

type gameJoinClient struct {
	grpc.ClientStream
}

func (x *gameJoinClient) Recv() (*UserConn, error) {
	m := new(UserConn)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameClient) Move(ctx context.Context, opts ...grpc.CallOption) (Game_MoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &Game_ServiceDesc.Streams[1], Game_Move_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gameMoveClient{stream}
	return x, nil
}

type Game_MoveClient interface {
	Send(*Pos) error
	Recv() (*UserMove, error)
	grpc.ClientStream
}

type gameMoveClient struct {
	grpc.ClientStream
}

func (x *gameMoveClient) Send(m *Pos) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameMoveClient) Recv() (*UserMove, error) {
	m := new(UserMove)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServer is the server API for Game service.
// All implementations must embed UnimplementedGameServer
// for forward compatibility
type GameServer interface {
	// basic
	Hello(context.Context, *HelloReq) (*HelloRes, error)
	// global
	GetUser(context.Context, *GetUserReq) (*GetUserRes, error)
	// game
	Join(*JoinReq, Game_JoinServer) error
	Move(Game_MoveServer) error
	mustEmbedUnimplementedGameServer()
}

// UnimplementedGameServer must be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (UnimplementedGameServer) Hello(context.Context, *HelloReq) (*HelloRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGameServer) GetUser(context.Context, *GetUserReq) (*GetUserRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedGameServer) Join(*JoinReq, Game_JoinServer) error {
	return status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedGameServer) Move(Game_MoveServer) error {
	return status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedGameServer) mustEmbedUnimplementedGameServer() {}

// UnsafeGameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServer will
// result in compilation errors.
type UnsafeGameServer interface {
	mustEmbedUnimplementedGameServer()
}

func RegisterGameServer(s grpc.ServiceRegistrar, srv GameServer) {
	s.RegisterService(&Game_ServiceDesc, srv)
}

func _Game_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Hello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Game_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Join_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServer).Join(m, &gameJoinServer{stream})
}

type Game_JoinServer interface {
	Send(*UserConn) error
	grpc.ServerStream
}

type gameJoinServer struct {
	grpc.ServerStream
}

func (x *gameJoinServer) Send(m *UserConn) error {
	return x.ServerStream.SendMsg(m)
}

func _Game_Move_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).Move(&gameMoveServer{stream})
}

type Game_MoveServer interface {
	Send(*UserMove) error
	Recv() (*Pos, error)
	grpc.ServerStream
}

type gameMoveServer struct {
	grpc.ServerStream
}

func (x *gameMoveServer) Send(m *UserMove) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameMoveServer) Recv() (*Pos, error) {
	m := new(Pos)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Game_ServiceDesc is the grpc.ServiceDesc for Game service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Game_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Game_Hello_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Game_GetUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Join",
			Handler:       _Game_Join_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Move",
			Handler:       _Game_Move_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/actiongame.proto",
}
