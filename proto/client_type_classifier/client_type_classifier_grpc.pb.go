// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: client_type_classifier.proto

package proto_client_type_classifier

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

// ClientTypeClassifierServiceClient is the client API for ClientTypeClassifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientTypeClassifierServiceClient interface {
	Create(ctx context.Context, in *CreateClientTypeClassifierRequest, opts ...grpc.CallOption) (*ClientTypeClassifierProfileResponse, error)
	Read(ctx context.Context, in *SingleClientTypeClassifierRequest, opts ...grpc.CallOption) (*ClientTypeClassifierProfileResponse, error)
	Update(ctx context.Context, in *UpdateClientTypeClassifierRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SingleClientTypeClassifierRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type clientTypeClassifierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientTypeClassifierServiceClient(cc grpc.ClientConnInterface) ClientTypeClassifierServiceClient {
	return &clientTypeClassifierServiceClient{cc}
}

func (c *clientTypeClassifierServiceClient) Create(ctx context.Context, in *CreateClientTypeClassifierRequest, opts ...grpc.CallOption) (*ClientTypeClassifierProfileResponse, error) {
	out := new(ClientTypeClassifierProfileResponse)
	err := c.cc.Invoke(ctx, "/ClientTypeClassifierService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTypeClassifierServiceClient) Read(ctx context.Context, in *SingleClientTypeClassifierRequest, opts ...grpc.CallOption) (*ClientTypeClassifierProfileResponse, error) {
	out := new(ClientTypeClassifierProfileResponse)
	err := c.cc.Invoke(ctx, "/ClientTypeClassifierService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTypeClassifierServiceClient) Update(ctx context.Context, in *UpdateClientTypeClassifierRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/ClientTypeClassifierService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientTypeClassifierServiceClient) Delete(ctx context.Context, in *SingleClientTypeClassifierRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/ClientTypeClassifierService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientTypeClassifierServiceServer is the server API for ClientTypeClassifierService service.
// All implementations must embed UnimplementedClientTypeClassifierServiceServer
// for forward compatibility
type ClientTypeClassifierServiceServer interface {
	Create(context.Context, *CreateClientTypeClassifierRequest) (*ClientTypeClassifierProfileResponse, error)
	Read(context.Context, *SingleClientTypeClassifierRequest) (*ClientTypeClassifierProfileResponse, error)
	Update(context.Context, *UpdateClientTypeClassifierRequest) (*SuccessResponse, error)
	Delete(context.Context, *SingleClientTypeClassifierRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedClientTypeClassifierServiceServer()
}

// UnimplementedClientTypeClassifierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientTypeClassifierServiceServer struct {
}

func (UnimplementedClientTypeClassifierServiceServer) Create(context.Context, *CreateClientTypeClassifierRequest) (*ClientTypeClassifierProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClientTypeClassifierServiceServer) Read(context.Context, *SingleClientTypeClassifierRequest) (*ClientTypeClassifierProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedClientTypeClassifierServiceServer) Update(context.Context, *UpdateClientTypeClassifierRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedClientTypeClassifierServiceServer) Delete(context.Context, *SingleClientTypeClassifierRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClientTypeClassifierServiceServer) mustEmbedUnimplementedClientTypeClassifierServiceServer() {
}

// UnsafeClientTypeClassifierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientTypeClassifierServiceServer will
// result in compilation errors.
type UnsafeClientTypeClassifierServiceServer interface {
	mustEmbedUnimplementedClientTypeClassifierServiceServer()
}

func RegisterClientTypeClassifierServiceServer(s grpc.ServiceRegistrar, srv ClientTypeClassifierServiceServer) {
	s.RegisterService(&ClientTypeClassifierService_ServiceDesc, srv)
}

func _ClientTypeClassifierService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClientTypeClassifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeClassifierServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientTypeClassifierService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeClassifierServiceServer).Create(ctx, req.(*CreateClientTypeClassifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTypeClassifierService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleClientTypeClassifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeClassifierServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientTypeClassifierService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeClassifierServiceServer).Read(ctx, req.(*SingleClientTypeClassifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTypeClassifierService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClientTypeClassifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeClassifierServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientTypeClassifierService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeClassifierServiceServer).Update(ctx, req.(*UpdateClientTypeClassifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientTypeClassifierService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleClientTypeClassifierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientTypeClassifierServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ClientTypeClassifierService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientTypeClassifierServiceServer).Delete(ctx, req.(*SingleClientTypeClassifierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientTypeClassifierService_ServiceDesc is the grpc.ServiceDesc for ClientTypeClassifierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientTypeClassifierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ClientTypeClassifierService",
	HandlerType: (*ClientTypeClassifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientTypeClassifierService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ClientTypeClassifierService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClientTypeClassifierService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClientTypeClassifierService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client_type_classifier.proto",
}
