// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/auth/v1beta1/query.proto

package authv1beta1

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
	Query_Accounts_FullMethodName             = "/cosmos.auth.v1beta1.Query/Accounts"
	Query_Account_FullMethodName              = "/cosmos.auth.v1beta1.Query/Account"
	Query_AccountAddressByID_FullMethodName   = "/cosmos.auth.v1beta1.Query/AccountAddressByID"
	Query_Params_FullMethodName               = "/cosmos.auth.v1beta1.Query/Params"
	Query_ModuleAccounts_FullMethodName       = "/cosmos.auth.v1beta1.Query/ModuleAccounts"
	Query_ModuleAccountByName_FullMethodName  = "/cosmos.auth.v1beta1.Query/ModuleAccountByName"
	Query_Bech32Prefix_FullMethodName         = "/cosmos.auth.v1beta1.Query/Bech32Prefix"
	Query_AddressBytesToString_FullMethodName = "/cosmos.auth.v1beta1.Query/AddressBytesToString"
	Query_AddressStringToBytes_FullMethodName = "/cosmos.auth.v1beta1.Query/AddressStringToBytes"
	Query_AccountInfo_FullMethodName          = "/cosmos.auth.v1beta1.Query/AccountInfo"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Accounts returns all the existing accounts.
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	Accounts(ctx context.Context, in *QueryAccountsRequest, opts ...grpc.CallOption) (*QueryAccountsResponse, error)
	// Account returns account details based on address.
	Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*QueryAccountResponse, error)
	// AccountAddressByID returns account address based on account number.
	AccountAddressByID(ctx context.Context, in *QueryAccountAddressByIDRequest, opts ...grpc.CallOption) (*QueryAccountAddressByIDResponse, error)
	// Params queries all parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// ModuleAccounts returns all the existing module accounts.
	ModuleAccounts(ctx context.Context, in *QueryModuleAccountsRequest, opts ...grpc.CallOption) (*QueryModuleAccountsResponse, error)
	// ModuleAccountByName returns the module account info by module name
	ModuleAccountByName(ctx context.Context, in *QueryModuleAccountByNameRequest, opts ...grpc.CallOption) (*QueryModuleAccountByNameResponse, error)
	// Bech32Prefix queries bech32Prefix
	Bech32Prefix(ctx context.Context, in *Bech32PrefixRequest, opts ...grpc.CallOption) (*Bech32PrefixResponse, error)
	// AddressBytesToString converts Account Address bytes to string
	AddressBytesToString(ctx context.Context, in *AddressBytesToStringRequest, opts ...grpc.CallOption) (*AddressBytesToStringResponse, error)
	// AddressStringToBytes converts Address string to bytes
	AddressStringToBytes(ctx context.Context, in *AddressStringToBytesRequest, opts ...grpc.CallOption) (*AddressStringToBytesResponse, error)
	// AccountInfo queries account info which is common to all account types.
	AccountInfo(ctx context.Context, in *QueryAccountInfoRequest, opts ...grpc.CallOption) (*QueryAccountInfoResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Accounts(ctx context.Context, in *QueryAccountsRequest, opts ...grpc.CallOption) (*QueryAccountsResponse, error) {
	out := new(QueryAccountsResponse)
	err := c.cc.Invoke(ctx, Query_Accounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Account(ctx context.Context, in *QueryAccountRequest, opts ...grpc.CallOption) (*QueryAccountResponse, error) {
	out := new(QueryAccountResponse)
	err := c.cc.Invoke(ctx, Query_Account_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AccountAddressByID(ctx context.Context, in *QueryAccountAddressByIDRequest, opts ...grpc.CallOption) (*QueryAccountAddressByIDResponse, error) {
	out := new(QueryAccountAddressByIDResponse)
	err := c.cc.Invoke(ctx, Query_AccountAddressByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ModuleAccounts(ctx context.Context, in *QueryModuleAccountsRequest, opts ...grpc.CallOption) (*QueryModuleAccountsResponse, error) {
	out := new(QueryModuleAccountsResponse)
	err := c.cc.Invoke(ctx, Query_ModuleAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ModuleAccountByName(ctx context.Context, in *QueryModuleAccountByNameRequest, opts ...grpc.CallOption) (*QueryModuleAccountByNameResponse, error) {
	out := new(QueryModuleAccountByNameResponse)
	err := c.cc.Invoke(ctx, Query_ModuleAccountByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Bech32Prefix(ctx context.Context, in *Bech32PrefixRequest, opts ...grpc.CallOption) (*Bech32PrefixResponse, error) {
	out := new(Bech32PrefixResponse)
	err := c.cc.Invoke(ctx, Query_Bech32Prefix_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AddressBytesToString(ctx context.Context, in *AddressBytesToStringRequest, opts ...grpc.CallOption) (*AddressBytesToStringResponse, error) {
	out := new(AddressBytesToStringResponse)
	err := c.cc.Invoke(ctx, Query_AddressBytesToString_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AddressStringToBytes(ctx context.Context, in *AddressStringToBytesRequest, opts ...grpc.CallOption) (*AddressStringToBytesResponse, error) {
	out := new(AddressStringToBytesResponse)
	err := c.cc.Invoke(ctx, Query_AddressStringToBytes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AccountInfo(ctx context.Context, in *QueryAccountInfoRequest, opts ...grpc.CallOption) (*QueryAccountInfoResponse, error) {
	out := new(QueryAccountInfoResponse)
	err := c.cc.Invoke(ctx, Query_AccountInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Accounts returns all the existing accounts.
	// When called from another module, this query might consume a high amount of
	// gas if the pagination field is incorrectly set.
	Accounts(context.Context, *QueryAccountsRequest) (*QueryAccountsResponse, error)
	// Account returns account details based on address.
	Account(context.Context, *QueryAccountRequest) (*QueryAccountResponse, error)
	// AccountAddressByID returns account address based on account number.
	AccountAddressByID(context.Context, *QueryAccountAddressByIDRequest) (*QueryAccountAddressByIDResponse, error)
	// Params queries all parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// ModuleAccounts returns all the existing module accounts.
	ModuleAccounts(context.Context, *QueryModuleAccountsRequest) (*QueryModuleAccountsResponse, error)
	// ModuleAccountByName returns the module account info by module name
	ModuleAccountByName(context.Context, *QueryModuleAccountByNameRequest) (*QueryModuleAccountByNameResponse, error)
	// Bech32Prefix queries bech32Prefix
	Bech32Prefix(context.Context, *Bech32PrefixRequest) (*Bech32PrefixResponse, error)
	// AddressBytesToString converts Account Address bytes to string
	AddressBytesToString(context.Context, *AddressBytesToStringRequest) (*AddressBytesToStringResponse, error)
	// AddressStringToBytes converts Address string to bytes
	AddressStringToBytes(context.Context, *AddressStringToBytesRequest) (*AddressStringToBytesResponse, error)
	// AccountInfo queries account info which is common to all account types.
	AccountInfo(context.Context, *QueryAccountInfoRequest) (*QueryAccountInfoResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Accounts(context.Context, *QueryAccountsRequest) (*QueryAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accounts not implemented")
}
func (UnimplementedQueryServer) Account(context.Context, *QueryAccountRequest) (*QueryAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Account not implemented")
}
func (UnimplementedQueryServer) AccountAddressByID(context.Context, *QueryAccountAddressByIDRequest) (*QueryAccountAddressByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountAddressByID not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) ModuleAccounts(context.Context, *QueryModuleAccountsRequest) (*QueryModuleAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModuleAccounts not implemented")
}
func (UnimplementedQueryServer) ModuleAccountByName(context.Context, *QueryModuleAccountByNameRequest) (*QueryModuleAccountByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModuleAccountByName not implemented")
}
func (UnimplementedQueryServer) Bech32Prefix(context.Context, *Bech32PrefixRequest) (*Bech32PrefixResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bech32Prefix not implemented")
}
func (UnimplementedQueryServer) AddressBytesToString(context.Context, *AddressBytesToStringRequest) (*AddressBytesToStringResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressBytesToString not implemented")
}
func (UnimplementedQueryServer) AddressStringToBytes(context.Context, *AddressStringToBytesRequest) (*AddressStringToBytesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressStringToBytes not implemented")
}
func (UnimplementedQueryServer) AccountInfo(context.Context, *QueryAccountInfoRequest) (*QueryAccountInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountInfo not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Accounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Accounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Accounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Accounts(ctx, req.(*QueryAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Account_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Account(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Account_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Account(ctx, req.(*QueryAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AccountAddressByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountAddressByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AccountAddressByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AccountAddressByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AccountAddressByID(ctx, req.(*QueryAccountAddressByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ModuleAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryModuleAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ModuleAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ModuleAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ModuleAccounts(ctx, req.(*QueryModuleAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ModuleAccountByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryModuleAccountByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ModuleAccountByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ModuleAccountByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ModuleAccountByName(ctx, req.(*QueryModuleAccountByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Bech32Prefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bech32PrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Bech32Prefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Bech32Prefix_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Bech32Prefix(ctx, req.(*Bech32PrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AddressBytesToString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressBytesToStringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AddressBytesToString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AddressBytesToString_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AddressBytesToString(ctx, req.(*AddressBytesToStringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AddressStringToBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressStringToBytesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AddressStringToBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AddressStringToBytes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AddressStringToBytes(ctx, req.(*AddressStringToBytesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_AccountInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AccountInfo(ctx, req.(*QueryAccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.auth.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Accounts",
			Handler:    _Query_Accounts_Handler,
		},
		{
			MethodName: "Account",
			Handler:    _Query_Account_Handler,
		},
		{
			MethodName: "AccountAddressByID",
			Handler:    _Query_AccountAddressByID_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ModuleAccounts",
			Handler:    _Query_ModuleAccounts_Handler,
		},
		{
			MethodName: "ModuleAccountByName",
			Handler:    _Query_ModuleAccountByName_Handler,
		},
		{
			MethodName: "Bech32Prefix",
			Handler:    _Query_Bech32Prefix_Handler,
		},
		{
			MethodName: "AddressBytesToString",
			Handler:    _Query_AddressBytesToString_Handler,
		},
		{
			MethodName: "AddressStringToBytes",
			Handler:    _Query_AddressStringToBytes_Handler,
		},
		{
			MethodName: "AccountInfo",
			Handler:    _Query_AccountInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/auth/v1beta1/query.proto",
}
