// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: blockchain/v1/wallet.proto

package blockchainv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	GetAddress(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAddressResponse, error)
	GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*Asset, error)
	GetSignature(ctx context.Context, in *GetSignatureRequest, opts ...grpc.CallOption) (*GetSignatureResponse, error)
	SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) GetAddress(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, "/blockchain.v1.WalletService/GetAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*Asset, error) {
	out := new(Asset)
	err := c.cc.Invoke(ctx, "/blockchain.v1.WalletService/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) GetSignature(ctx context.Context, in *GetSignatureRequest, opts ...grpc.CallOption) (*GetSignatureResponse, error) {
	out := new(GetSignatureResponse)
	err := c.cc.Invoke(ctx, "/blockchain.v1.WalletService/GetSignature", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) SendTransaction(ctx context.Context, in *SendTransactionRequest, opts ...grpc.CallOption) (*SendTransactionResponse, error) {
	out := new(SendTransactionResponse)
	err := c.cc.Invoke(ctx, "/blockchain.v1.WalletService/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations must embed UnimplementedWalletServiceServer
// for forward compatibility
type WalletServiceServer interface {
	GetAddress(context.Context, *emptypb.Empty) (*GetAddressResponse, error)
	GetMetadata(context.Context, *GetMetadataRequest) (*Asset, error)
	GetSignature(context.Context, *GetSignatureRequest) (*GetSignatureResponse, error)
	SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionResponse, error)
	mustEmbedUnimplementedWalletServiceServer()
}

// UnimplementedWalletServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWalletServiceServer struct {
}

func (UnimplementedWalletServiceServer) GetAddress(context.Context, *emptypb.Empty) (*GetAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedWalletServiceServer) GetMetadata(context.Context, *GetMetadataRequest) (*Asset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedWalletServiceServer) GetSignature(context.Context, *GetSignatureRequest) (*GetSignatureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignature not implemented")
}
func (UnimplementedWalletServiceServer) SendTransaction(context.Context, *SendTransactionRequest) (*SendTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (UnimplementedWalletServiceServer) mustEmbedUnimplementedWalletServiceServer() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.v1.WalletService/GetAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetAddress(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.v1.WalletService/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetMetadata(ctx, req.(*GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_GetSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.v1.WalletService/GetSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetSignature(ctx, req.(*GetSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockchain.v1.WalletService/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).SendTransaction(ctx, req.(*SendTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blockchain.v1.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAddress",
			Handler:    _WalletService_GetAddress_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _WalletService_GetMetadata_Handler,
		},
		{
			MethodName: "GetSignature",
			Handler:    _WalletService_GetSignature_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _WalletService_SendTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain/v1/wallet.proto",
}
