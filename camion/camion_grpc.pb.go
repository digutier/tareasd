// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package camion

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CamionServiceClient is the client API for CamionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CamionServiceClient interface {
	RecibirPaquete(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Paquete, error)
	EntregaPaquete(ctx context.Context, in *Paquete, opts ...grpc.CallOption) (*Message, error)
	EnviarEstado(ctx context.Context, in *Paquete, opts ...grpc.CallOption) (*Message, error)
}

type camionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCamionServiceClient(cc grpc.ClientConnInterface) CamionServiceClient {
	return &camionServiceClient{cc}
}

func (c *camionServiceClient) RecibirPaquete(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Paquete, error) {
	out := new(Paquete)
	err := c.cc.Invoke(ctx, "/camion.CamionService/RecibirPaquete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *camionServiceClient) EntregaPaquete(ctx context.Context, in *Paquete, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/camion.CamionService/EntregaPaquete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *camionServiceClient) EnviarEstado(ctx context.Context, in *Paquete, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/camion.CamionService/EnviarEstado", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CamionServiceServer is the server API for CamionService service.
// All implementations must embed UnimplementedCamionServiceServer
// for forward compatibility
type CamionServiceServer interface {
	RecibirPaquete(context.Context, *Message) (*Paquete, error)
	EntregaPaquete(context.Context, *Paquete) (*Message, error)
	EnviarEstado(context.Context, *Paquete) (*Message, error)
	mustEmbedUnimplementedCamionServiceServer()
}

// UnimplementedCamionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCamionServiceServer struct {
}

func (UnimplementedCamionServiceServer) RecibirPaquete(context.Context, *Message) (*Paquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecibirPaquete not implemented")
}
func (UnimplementedCamionServiceServer) EntregaPaquete(context.Context, *Paquete) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntregaPaquete not implemented")
}
func (UnimplementedCamionServiceServer) EnviarEstado(context.Context, *Paquete) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarEstado not implemented")
}
func (UnimplementedCamionServiceServer) mustEmbedUnimplementedCamionServiceServer() {}

// UnsafeCamionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CamionServiceServer will
// result in compilation errors.
type UnsafeCamionServiceServer interface {
	mustEmbedUnimplementedCamionServiceServer()
}

func RegisterCamionServiceServer(s *grpc.Server, srv CamionServiceServer) {
	s.RegisterService(&_CamionService_serviceDesc, srv)
}

func _CamionService_RecibirPaquete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionServiceServer).RecibirPaquete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/camion.CamionService/RecibirPaquete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionServiceServer).RecibirPaquete(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _CamionService_EntregaPaquete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paquete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionServiceServer).EntregaPaquete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/camion.CamionService/EntregaPaquete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionServiceServer).EntregaPaquete(ctx, req.(*Paquete))
	}
	return interceptor(ctx, in, info, handler)
}

func _CamionService_EnviarEstado_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paquete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionServiceServer).EnviarEstado(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/camion.CamionService/EnviarEstado",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionServiceServer).EnviarEstado(ctx, req.(*Paquete))
	}
	return interceptor(ctx, in, info, handler)
}

var _CamionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "camion.CamionService",
	HandlerType: (*CamionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecibirPaquete",
			Handler:    _CamionService_RecibirPaquete_Handler,
		},
		{
			MethodName: "EntregaPaquete",
			Handler:    _CamionService_EntregaPaquete_Handler,
		},
		{
			MethodName: "EnviarEstado",
			Handler:    _CamionService_EnviarEstado_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "camion.proto",
}
