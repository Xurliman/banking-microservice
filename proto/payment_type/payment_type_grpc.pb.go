// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: payment_type.proto

package proto_bank

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

// PaymentTypeServiceClient is the client API for PaymentTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentTypeServiceClient interface {
	Create(ctx context.Context, in *CreatePaymentTypeRequest, opts ...grpc.CallOption) (*PaymentTypeProfileResponse, error)
	Read(ctx context.Context, in *SinglePaymentTypeRequest, opts ...grpc.CallOption) (*PaymentTypeProfileResponse, error)
	Update(ctx context.Context, in *UpdatePaymentTypeRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SinglePaymentTypeRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type paymentTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentTypeServiceClient(cc grpc.ClientConnInterface) PaymentTypeServiceClient {
	return &paymentTypeServiceClient{cc}
}

func (c *paymentTypeServiceClient) Create(ctx context.Context, in *CreatePaymentTypeRequest, opts ...grpc.CallOption) (*PaymentTypeProfileResponse, error) {
	out := new(PaymentTypeProfileResponse)
	err := c.cc.Invoke(ctx, "/PaymentTypeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentTypeServiceClient) Read(ctx context.Context, in *SinglePaymentTypeRequest, opts ...grpc.CallOption) (*PaymentTypeProfileResponse, error) {
	out := new(PaymentTypeProfileResponse)
	err := c.cc.Invoke(ctx, "/PaymentTypeService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentTypeServiceClient) Update(ctx context.Context, in *UpdatePaymentTypeRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/PaymentTypeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentTypeServiceClient) Delete(ctx context.Context, in *SinglePaymentTypeRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/PaymentTypeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentTypeServiceServer is the server API for PaymentTypeService service.
// All implementations must embed UnimplementedPaymentTypeServiceServer
// for forward compatibility
type PaymentTypeServiceServer interface {
	Create(context.Context, *CreatePaymentTypeRequest) (*PaymentTypeProfileResponse, error)
	Read(context.Context, *SinglePaymentTypeRequest) (*PaymentTypeProfileResponse, error)
	Update(context.Context, *UpdatePaymentTypeRequest) (*SuccessResponse, error)
	Delete(context.Context, *SinglePaymentTypeRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedPaymentTypeServiceServer()
}

// UnimplementedPaymentTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentTypeServiceServer struct {
}

func (UnimplementedPaymentTypeServiceServer) Create(context.Context, *CreatePaymentTypeRequest) (*PaymentTypeProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPaymentTypeServiceServer) Read(context.Context, *SinglePaymentTypeRequest) (*PaymentTypeProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedPaymentTypeServiceServer) Update(context.Context, *UpdatePaymentTypeRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPaymentTypeServiceServer) Delete(context.Context, *SinglePaymentTypeRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPaymentTypeServiceServer) mustEmbedUnimplementedPaymentTypeServiceServer() {}

// UnsafePaymentTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentTypeServiceServer will
// result in compilation errors.
type UnsafePaymentTypeServiceServer interface {
	mustEmbedUnimplementedPaymentTypeServiceServer()
}

func RegisterPaymentTypeServiceServer(s grpc.ServiceRegistrar, srv PaymentTypeServiceServer) {
	s.RegisterService(&PaymentTypeService_ServiceDesc, srv)
}

func _PaymentTypeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentTypeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentTypeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentTypeServiceServer).Create(ctx, req.(*CreatePaymentTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentTypeService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePaymentTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentTypeServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentTypeService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentTypeServiceServer).Read(ctx, req.(*SinglePaymentTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentTypeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentTypeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentTypeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentTypeServiceServer).Update(ctx, req.(*UpdatePaymentTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentTypeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SinglePaymentTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentTypeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PaymentTypeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentTypeServiceServer).Delete(ctx, req.(*SinglePaymentTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentTypeService_ServiceDesc is the grpc.ServiceDesc for PaymentTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PaymentTypeService",
	HandlerType: (*PaymentTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PaymentTypeService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _PaymentTypeService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PaymentTypeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PaymentTypeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment_type.proto",
}
