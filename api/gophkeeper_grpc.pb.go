// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/gophkeeper.proto

package api

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
	GophKeeper_Register_FullMethodName     = "/gophkeeper.GophKeeper/Register"
	GophKeeper_Login_FullMethodName        = "/gophkeeper.GophKeeper/Login"
	GophKeeper_SaveData_FullMethodName     = "/gophkeeper.GophKeeper/SaveData"
	GophKeeper_SaveLogin_FullMethodName    = "/gophkeeper.GophKeeper/SaveLogin"
	GophKeeper_SaveText_FullMethodName     = "/gophkeeper.GophKeeper/SaveText"
	GophKeeper_SaveBankCard_FullMethodName = "/gophkeeper.GophKeeper/SaveBankCard"
	GophKeeper_ListData_FullMethodName     = "/gophkeeper.GophKeeper/ListData"
	GophKeeper_GetData_FullMethodName      = "/gophkeeper.GophKeeper/GetData"
	GophKeeper_GetLogin_FullMethodName     = "/gophkeeper.GophKeeper/GetLogin"
	GophKeeper_GetText_FullMethodName      = "/gophkeeper.GophKeeper/GetText"
)

// GophKeeperClient is the client API for GophKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophKeeperClient interface {
	Register(ctx context.Context, in *RegisterLoginReq, opts ...grpc.CallOption) (*RegisterLoginResp, error)
	Login(ctx context.Context, in *RegisterLoginReq, opts ...grpc.CallOption) (*RegisterLoginResp, error)
	SaveData(ctx context.Context, in *SaveDataReq, opts ...grpc.CallOption) (*SaveDataResp, error)
	SaveLogin(ctx context.Context, in *SaveLoginReq, opts ...grpc.CallOption) (*SaveDataResp, error)
	SaveText(ctx context.Context, in *SaveTextReq, opts ...grpc.CallOption) (*SaveDataResp, error)
	SaveBankCard(ctx context.Context, in *SaveBankCardReq, opts ...grpc.CallOption) (*SaveDataResp, error)
	ListData(ctx context.Context, in *ListDataReq, opts ...grpc.CallOption) (*ListDataResp, error)
	GetData(ctx context.Context, in *GetDataReq, opts ...grpc.CallOption) (*GetDataResp, error)
	GetLogin(ctx context.Context, in *GetDataReq, opts ...grpc.CallOption) (*GetLoginResp, error)
	GetText(ctx context.Context, in *GetDataReq, opts ...grpc.CallOption) (*GetTextResp, error)
}

type gophKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophKeeperClient(cc grpc.ClientConnInterface) GophKeeperClient {
	return &gophKeeperClient{cc}
}

func (c *gophKeeperClient) Register(ctx context.Context, in *RegisterLoginReq, opts ...grpc.CallOption) (*RegisterLoginResp, error) {
	out := new(RegisterLoginResp)
	err := c.cc.Invoke(ctx, GophKeeper_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) Login(ctx context.Context, in *RegisterLoginReq, opts ...grpc.CallOption) (*RegisterLoginResp, error) {
	out := new(RegisterLoginResp)
	err := c.cc.Invoke(ctx, GophKeeper_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SaveData(ctx context.Context, in *SaveDataReq, opts ...grpc.CallOption) (*SaveDataResp, error) {
	out := new(SaveDataResp)
	err := c.cc.Invoke(ctx, GophKeeper_SaveData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SaveLogin(ctx context.Context, in *SaveLoginReq, opts ...grpc.CallOption) (*SaveDataResp, error) {
	out := new(SaveDataResp)
	err := c.cc.Invoke(ctx, GophKeeper_SaveLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SaveText(ctx context.Context, in *SaveTextReq, opts ...grpc.CallOption) (*SaveDataResp, error) {
	out := new(SaveDataResp)
	err := c.cc.Invoke(ctx, GophKeeper_SaveText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) SaveBankCard(ctx context.Context, in *SaveBankCardReq, opts ...grpc.CallOption) (*SaveDataResp, error) {
	out := new(SaveDataResp)
	err := c.cc.Invoke(ctx, GophKeeper_SaveBankCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) ListData(ctx context.Context, in *ListDataReq, opts ...grpc.CallOption) (*ListDataResp, error) {
	out := new(ListDataResp)
	err := c.cc.Invoke(ctx, GophKeeper_ListData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetData(ctx context.Context, in *GetDataReq, opts ...grpc.CallOption) (*GetDataResp, error) {
	out := new(GetDataResp)
	err := c.cc.Invoke(ctx, GophKeeper_GetData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetLogin(ctx context.Context, in *GetDataReq, opts ...grpc.CallOption) (*GetLoginResp, error) {
	out := new(GetLoginResp)
	err := c.cc.Invoke(ctx, GophKeeper_GetLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) GetText(ctx context.Context, in *GetDataReq, opts ...grpc.CallOption) (*GetTextResp, error) {
	out := new(GetTextResp)
	err := c.cc.Invoke(ctx, GophKeeper_GetText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophKeeperServer is the server API for GophKeeper service.
// All implementations must embed UnimplementedGophKeeperServer
// for forward compatibility
type GophKeeperServer interface {
	Register(context.Context, *RegisterLoginReq) (*RegisterLoginResp, error)
	Login(context.Context, *RegisterLoginReq) (*RegisterLoginResp, error)
	SaveData(context.Context, *SaveDataReq) (*SaveDataResp, error)
	SaveLogin(context.Context, *SaveLoginReq) (*SaveDataResp, error)
	SaveText(context.Context, *SaveTextReq) (*SaveDataResp, error)
	SaveBankCard(context.Context, *SaveBankCardReq) (*SaveDataResp, error)
	ListData(context.Context, *ListDataReq) (*ListDataResp, error)
	GetData(context.Context, *GetDataReq) (*GetDataResp, error)
	GetLogin(context.Context, *GetDataReq) (*GetLoginResp, error)
	GetText(context.Context, *GetDataReq) (*GetTextResp, error)
	mustEmbedUnimplementedGophKeeperServer()
}

// UnimplementedGophKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophKeeperServer struct {
}

func (UnimplementedGophKeeperServer) Register(context.Context, *RegisterLoginReq) (*RegisterLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGophKeeperServer) Login(context.Context, *RegisterLoginReq) (*RegisterLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGophKeeperServer) SaveData(context.Context, *SaveDataReq) (*SaveDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveData not implemented")
}
func (UnimplementedGophKeeperServer) SaveLogin(context.Context, *SaveLoginReq) (*SaveDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLogin not implemented")
}
func (UnimplementedGophKeeperServer) SaveText(context.Context, *SaveTextReq) (*SaveDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveText not implemented")
}
func (UnimplementedGophKeeperServer) SaveBankCard(context.Context, *SaveBankCardReq) (*SaveDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBankCard not implemented")
}
func (UnimplementedGophKeeperServer) ListData(context.Context, *ListDataReq) (*ListDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListData not implemented")
}
func (UnimplementedGophKeeperServer) GetData(context.Context, *GetDataReq) (*GetDataResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedGophKeeperServer) GetLogin(context.Context, *GetDataReq) (*GetLoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogin not implemented")
}
func (UnimplementedGophKeeperServer) GetText(context.Context, *GetDataReq) (*GetTextResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetText not implemented")
}
func (UnimplementedGophKeeperServer) mustEmbedUnimplementedGophKeeperServer() {}

// UnsafeGophKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophKeeperServer will
// result in compilation errors.
type UnsafeGophKeeperServer interface {
	mustEmbedUnimplementedGophKeeperServer()
}

func RegisterGophKeeperServer(s grpc.ServiceRegistrar, srv GophKeeperServer) {
	s.RegisterService(&GophKeeper_ServiceDesc, srv)
}

func _GophKeeper_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Register(ctx, req.(*RegisterLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Login(ctx, req.(*RegisterLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SaveData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SaveData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SaveData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SaveData(ctx, req.(*SaveDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SaveLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SaveLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SaveLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SaveLogin(ctx, req.(*SaveLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SaveText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTextReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SaveText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SaveText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SaveText(ctx, req.(*SaveTextReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_SaveBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBankCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).SaveBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_SaveBankCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).SaveBankCard(ctx, req.(*SaveBankCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_ListData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).ListData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_ListData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).ListData(ctx, req.(*ListDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetData(ctx, req.(*GetDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetLogin(ctx, req.(*GetDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_GetText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).GetText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeper_GetText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).GetText(ctx, req.(*GetDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GophKeeper_ServiceDesc is the grpc.ServiceDesc for GophKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.GophKeeper",
	HandlerType: (*GophKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _GophKeeper_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GophKeeper_Login_Handler,
		},
		{
			MethodName: "SaveData",
			Handler:    _GophKeeper_SaveData_Handler,
		},
		{
			MethodName: "SaveLogin",
			Handler:    _GophKeeper_SaveLogin_Handler,
		},
		{
			MethodName: "SaveText",
			Handler:    _GophKeeper_SaveText_Handler,
		},
		{
			MethodName: "SaveBankCard",
			Handler:    _GophKeeper_SaveBankCard_Handler,
		},
		{
			MethodName: "ListData",
			Handler:    _GophKeeper_ListData_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _GophKeeper_GetData_Handler,
		},
		{
			MethodName: "GetLogin",
			Handler:    _GophKeeper_GetLogin_Handler,
		},
		{
			MethodName: "GetText",
			Handler:    _GophKeeper_GetText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/gophkeeper.proto",
}