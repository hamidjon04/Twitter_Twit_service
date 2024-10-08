// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: twit.proto

package twit

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	TwitServiceClient_CreateTwit_FullMethodName      = "/twit.TwitServiceClient/CreateTwit"
	TwitServiceClient_UpdateTwit_FullMethodName      = "/twit.TwitServiceClient/UpdateTwit"
	TwitServiceClient_DeleteTwit_FullMethodName      = "/twit.TwitServiceClient/DeleteTwit"
	TwitServiceClient_GetTwits_FullMethodName        = "/twit.TwitServiceClient/GetTwits"
	TwitServiceClient_AddLike_FullMethodName         = "/twit.TwitServiceClient/AddLike"
	TwitServiceClient_RemoveLike_FullMethodName      = "/twit.TwitServiceClient/RemoveLike"
	TwitServiceClient_AddComment_FullMethodName      = "/twit.TwitServiceClient/AddComment"
	TwitServiceClient_RemoveComment_FullMethodName   = "/twit.TwitServiceClient/RemoveComment"
	TwitServiceClient_GetComment_FullMethodName      = "/twit.TwitServiceClient/GetComment"
	TwitServiceClient_GetCommentById_FullMethodName  = "/twit.TwitServiceClient/GetCommentById"
	TwitServiceClient_UpdateComment_FullMethodName   = "/twit.TwitServiceClient/UpdateComment"
	TwitServiceClient_GetFollowerTwit_FullMethodName = "/twit.TwitServiceClient/GetFollowerTwit"
)

// TwitServiceClientClient is the client API for TwitServiceClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwitServiceClientClient interface {
	CreateTwit(ctx context.Context, in *CreateTwitReq, opts ...grpc.CallOption) (*CreateTwitResp, error)
	UpdateTwit(ctx context.Context, in *UpadateReq, opts ...grpc.CallOption) (*UpdateTwitResp, error)
	DeleteTwit(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Message, error)
	GetTwits(ctx context.Context, in *GetTwitsReq, opts ...grpc.CallOption) (*GetTwitsResp, error)
	AddLike(ctx context.Context, in *AddLikeReq, opts ...grpc.CallOption) (*AddLikeResp, error)
	RemoveLike(ctx context.Context, in *DeleteLikeReq, opts ...grpc.CallOption) (*Message, error)
	AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error)
	RemoveComment(ctx context.Context, in *RemoveCommentReq, opts ...grpc.CallOption) (*Message, error)
	GetComment(ctx context.Context, in *GetCommentReq, opts ...grpc.CallOption) (*GetCommentResp, error)
	GetCommentById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Comment, error)
	UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error)
	GetFollowerTwit(ctx context.Context, in *GetTwitsReq, opts ...grpc.CallOption) (*GetTwitsResp, error)
}

type twitServiceClientClient struct {
	cc grpc.ClientConnInterface
}

func NewTwitServiceClientClient(cc grpc.ClientConnInterface) TwitServiceClientClient {
	return &twitServiceClientClient{cc}
}

func (c *twitServiceClientClient) CreateTwit(ctx context.Context, in *CreateTwitReq, opts ...grpc.CallOption) (*CreateTwitResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTwitResp)
	err := c.cc.Invoke(ctx, TwitServiceClient_CreateTwit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) UpdateTwit(ctx context.Context, in *UpadateReq, opts ...grpc.CallOption) (*UpdateTwitResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTwitResp)
	err := c.cc.Invoke(ctx, TwitServiceClient_UpdateTwit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) DeleteTwit(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, TwitServiceClient_DeleteTwit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) GetTwits(ctx context.Context, in *GetTwitsReq, opts ...grpc.CallOption) (*GetTwitsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTwitsResp)
	err := c.cc.Invoke(ctx, TwitServiceClient_GetTwits_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) AddLike(ctx context.Context, in *AddLikeReq, opts ...grpc.CallOption) (*AddLikeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddLikeResp)
	err := c.cc.Invoke(ctx, TwitServiceClient_AddLike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) RemoveLike(ctx context.Context, in *DeleteLikeReq, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, TwitServiceClient_RemoveLike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) AddComment(ctx context.Context, in *AddCommentReq, opts ...grpc.CallOption) (*AddCommentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCommentResp)
	err := c.cc.Invoke(ctx, TwitServiceClient_AddComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) RemoveComment(ctx context.Context, in *RemoveCommentReq, opts ...grpc.CallOption) (*Message, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Message)
	err := c.cc.Invoke(ctx, TwitServiceClient_RemoveComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) GetComment(ctx context.Context, in *GetCommentReq, opts ...grpc.CallOption) (*GetCommentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentResp)
	err := c.cc.Invoke(ctx, TwitServiceClient_GetComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) GetCommentById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Comment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comment)
	err := c.cc.Invoke(ctx, TwitServiceClient_GetCommentById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) UpdateComment(ctx context.Context, in *UpdateCommentReq, opts ...grpc.CallOption) (*UpdateCommentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCommentResp)
	err := c.cc.Invoke(ctx, TwitServiceClient_UpdateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitServiceClientClient) GetFollowerTwit(ctx context.Context, in *GetTwitsReq, opts ...grpc.CallOption) (*GetTwitsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTwitsResp)
	err := c.cc.Invoke(ctx, TwitServiceClient_GetFollowerTwit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwitServiceClientServer is the server API for TwitServiceClient service.
// All implementations must embed UnimplementedTwitServiceClientServer
// for forward compatibility
type TwitServiceClientServer interface {
	CreateTwit(context.Context, *CreateTwitReq) (*CreateTwitResp, error)
	UpdateTwit(context.Context, *UpadateReq) (*UpdateTwitResp, error)
	DeleteTwit(context.Context, *Id) (*Message, error)
	GetTwits(context.Context, *GetTwitsReq) (*GetTwitsResp, error)
	AddLike(context.Context, *AddLikeReq) (*AddLikeResp, error)
	RemoveLike(context.Context, *DeleteLikeReq) (*Message, error)
	AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error)
	RemoveComment(context.Context, *RemoveCommentReq) (*Message, error)
	GetComment(context.Context, *GetCommentReq) (*GetCommentResp, error)
	GetCommentById(context.Context, *Id) (*Comment, error)
	UpdateComment(context.Context, *UpdateCommentReq) (*UpdateCommentResp, error)
	GetFollowerTwit(context.Context, *GetTwitsReq) (*GetTwitsResp, error)
	mustEmbedUnimplementedTwitServiceClientServer()
}

// UnimplementedTwitServiceClientServer must be embedded to have forward compatible implementations.
type UnimplementedTwitServiceClientServer struct {
}

func (UnimplementedTwitServiceClientServer) CreateTwit(context.Context, *CreateTwitReq) (*CreateTwitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTwit not implemented")
}
func (UnimplementedTwitServiceClientServer) UpdateTwit(context.Context, *UpadateReq) (*UpdateTwitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTwit not implemented")
}
func (UnimplementedTwitServiceClientServer) DeleteTwit(context.Context, *Id) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTwit not implemented")
}
func (UnimplementedTwitServiceClientServer) GetTwits(context.Context, *GetTwitsReq) (*GetTwitsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTwits not implemented")
}
func (UnimplementedTwitServiceClientServer) AddLike(context.Context, *AddLikeReq) (*AddLikeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLike not implemented")
}
func (UnimplementedTwitServiceClientServer) RemoveLike(context.Context, *DeleteLikeReq) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLike not implemented")
}
func (UnimplementedTwitServiceClientServer) AddComment(context.Context, *AddCommentReq) (*AddCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedTwitServiceClientServer) RemoveComment(context.Context, *RemoveCommentReq) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveComment not implemented")
}
func (UnimplementedTwitServiceClientServer) GetComment(context.Context, *GetCommentReq) (*GetCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedTwitServiceClientServer) GetCommentById(context.Context, *Id) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentById not implemented")
}
func (UnimplementedTwitServiceClientServer) UpdateComment(context.Context, *UpdateCommentReq) (*UpdateCommentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedTwitServiceClientServer) GetFollowerTwit(context.Context, *GetTwitsReq) (*GetTwitsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowerTwit not implemented")
}
func (UnimplementedTwitServiceClientServer) mustEmbedUnimplementedTwitServiceClientServer() {}

// UnsafeTwitServiceClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwitServiceClientServer will
// result in compilation errors.
type UnsafeTwitServiceClientServer interface {
	mustEmbedUnimplementedTwitServiceClientServer()
}

func RegisterTwitServiceClientServer(s grpc.ServiceRegistrar, srv TwitServiceClientServer) {
	s.RegisterService(&TwitServiceClient_ServiceDesc, srv)
}

func _TwitServiceClient_CreateTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTwitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).CreateTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_CreateTwit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).CreateTwit(ctx, req.(*CreateTwitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_UpdateTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpadateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).UpdateTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_UpdateTwit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).UpdateTwit(ctx, req.(*UpadateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_DeleteTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).DeleteTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_DeleteTwit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).DeleteTwit(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_GetTwits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTwitsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).GetTwits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_GetTwits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).GetTwits(ctx, req.(*GetTwitsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_AddLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).AddLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_AddLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).AddLike(ctx, req.(*AddLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_RemoveLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).RemoveLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_RemoveLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).RemoveLike(ctx, req.(*DeleteLikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).AddComment(ctx, req.(*AddCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_RemoveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).RemoveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_RemoveComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).RemoveComment(ctx, req.(*RemoveCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_GetComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).GetComment(ctx, req.(*GetCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_GetCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).GetCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_GetCommentById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).GetCommentById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).UpdateComment(ctx, req.(*UpdateCommentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitServiceClient_GetFollowerTwit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTwitsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitServiceClientServer).GetFollowerTwit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwitServiceClient_GetFollowerTwit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitServiceClientServer).GetFollowerTwit(ctx, req.(*GetTwitsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TwitServiceClient_ServiceDesc is the grpc.ServiceDesc for TwitServiceClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TwitServiceClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "twit.TwitServiceClient",
	HandlerType: (*TwitServiceClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTwit",
			Handler:    _TwitServiceClient_CreateTwit_Handler,
		},
		{
			MethodName: "UpdateTwit",
			Handler:    _TwitServiceClient_UpdateTwit_Handler,
		},
		{
			MethodName: "DeleteTwit",
			Handler:    _TwitServiceClient_DeleteTwit_Handler,
		},
		{
			MethodName: "GetTwits",
			Handler:    _TwitServiceClient_GetTwits_Handler,
		},
		{
			MethodName: "AddLike",
			Handler:    _TwitServiceClient_AddLike_Handler,
		},
		{
			MethodName: "RemoveLike",
			Handler:    _TwitServiceClient_RemoveLike_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _TwitServiceClient_AddComment_Handler,
		},
		{
			MethodName: "RemoveComment",
			Handler:    _TwitServiceClient_RemoveComment_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _TwitServiceClient_GetComment_Handler,
		},
		{
			MethodName: "GetCommentById",
			Handler:    _TwitServiceClient_GetCommentById_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _TwitServiceClient_UpdateComment_Handler,
		},
		{
			MethodName: "GetFollowerTwit",
			Handler:    _TwitServiceClient_GetFollowerTwit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "twit.proto",
}
