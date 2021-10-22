// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gopb

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

// SignServiceClient is the client API for SignService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SignServiceClient interface {
	//新增开立
	OrgAccountAdd(ctx context.Context, in *OrgAccountAddReq, opts ...grpc.CallOption) (*OrgAccountAddResp, error)
	UserAccountAdd(ctx context.Context, in *UserAccountAddReq, opts ...grpc.CallOption) (*UserAccountAddResp, error)
	TemplateCreate(ctx context.Context, in *CreateByTemplateReq, opts ...grpc.CallOption) (*CreateByTemplateResp, error)
	TemplateGet(ctx context.Context, in *TemplateGetReq, opts ...grpc.CallOption) (*TemplateGetResp, error)
	FlowOneStepCreate(ctx context.Context, in *FlowOneStepCreateReq, opts ...grpc.CallOption) (*FlowOneStepCreateResp, error)
	//开立合同
	VoucherIssue(ctx context.Context, in *VoucherIssueReq, opts ...grpc.CallOption) (*VoucherIssueResp, error)
	//签收完成后通知
	Notice(ctx context.Context, in *NoticeReq, opts ...grpc.CallOption) (*NoticeResp, error)
}

type signServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSignServiceClient(cc grpc.ClientConnInterface) SignServiceClient {
	return &signServiceClient{cc}
}

func (c *signServiceClient) OrgAccountAdd(ctx context.Context, in *OrgAccountAddReq, opts ...grpc.CallOption) (*OrgAccountAddResp, error) {
	out := new(OrgAccountAddResp)
	err := c.cc.Invoke(ctx, "/SignService/OrgAccountAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signServiceClient) UserAccountAdd(ctx context.Context, in *UserAccountAddReq, opts ...grpc.CallOption) (*UserAccountAddResp, error) {
	out := new(UserAccountAddResp)
	err := c.cc.Invoke(ctx, "/SignService/UserAccountAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signServiceClient) TemplateCreate(ctx context.Context, in *CreateByTemplateReq, opts ...grpc.CallOption) (*CreateByTemplateResp, error) {
	out := new(CreateByTemplateResp)
	err := c.cc.Invoke(ctx, "/SignService/TemplateCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signServiceClient) TemplateGet(ctx context.Context, in *TemplateGetReq, opts ...grpc.CallOption) (*TemplateGetResp, error) {
	out := new(TemplateGetResp)
	err := c.cc.Invoke(ctx, "/SignService/TemplateGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signServiceClient) FlowOneStepCreate(ctx context.Context, in *FlowOneStepCreateReq, opts ...grpc.CallOption) (*FlowOneStepCreateResp, error) {
	out := new(FlowOneStepCreateResp)
	err := c.cc.Invoke(ctx, "/SignService/FlowOneStepCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signServiceClient) VoucherIssue(ctx context.Context, in *VoucherIssueReq, opts ...grpc.CallOption) (*VoucherIssueResp, error) {
	out := new(VoucherIssueResp)
	err := c.cc.Invoke(ctx, "/SignService/VoucherIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signServiceClient) Notice(ctx context.Context, in *NoticeReq, opts ...grpc.CallOption) (*NoticeResp, error) {
	out := new(NoticeResp)
	err := c.cc.Invoke(ctx, "/SignService/Notice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SignServiceServer is the server API for SignService service.
// All implementations must embed UnimplementedSignServiceServer
// for forward compatibility
type SignServiceServer interface {
	//新增开立
	OrgAccountAdd(context.Context, *OrgAccountAddReq) (*OrgAccountAddResp, error)
	UserAccountAdd(context.Context, *UserAccountAddReq) (*UserAccountAddResp, error)
	TemplateCreate(context.Context, *CreateByTemplateReq) (*CreateByTemplateResp, error)
	TemplateGet(context.Context, *TemplateGetReq) (*TemplateGetResp, error)
	FlowOneStepCreate(context.Context, *FlowOneStepCreateReq) (*FlowOneStepCreateResp, error)
	//开立合同
	VoucherIssue(context.Context, *VoucherIssueReq) (*VoucherIssueResp, error)
	//签收完成后通知
	Notice(context.Context, *NoticeReq) (*NoticeResp, error)
	mustEmbedUnimplementedSignServiceServer()
}

// UnimplementedSignServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSignServiceServer struct {
}

func (UnimplementedSignServiceServer) OrgAccountAdd(context.Context, *OrgAccountAddReq) (*OrgAccountAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrgAccountAdd not implemented")
}
func (UnimplementedSignServiceServer) UserAccountAdd(context.Context, *UserAccountAddReq) (*UserAccountAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAccountAdd not implemented")
}
func (UnimplementedSignServiceServer) TemplateCreate(context.Context, *CreateByTemplateReq) (*CreateByTemplateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemplateCreate not implemented")
}
func (UnimplementedSignServiceServer) TemplateGet(context.Context, *TemplateGetReq) (*TemplateGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemplateGet not implemented")
}
func (UnimplementedSignServiceServer) FlowOneStepCreate(context.Context, *FlowOneStepCreateReq) (*FlowOneStepCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FlowOneStepCreate not implemented")
}
func (UnimplementedSignServiceServer) VoucherIssue(context.Context, *VoucherIssueReq) (*VoucherIssueResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoucherIssue not implemented")
}
func (UnimplementedSignServiceServer) Notice(context.Context, *NoticeReq) (*NoticeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notice not implemented")
}
func (UnimplementedSignServiceServer) mustEmbedUnimplementedSignServiceServer() {}

// UnsafeSignServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SignServiceServer will
// result in compilation errors.
type UnsafeSignServiceServer interface {
	mustEmbedUnimplementedSignServiceServer()
}

func RegisterSignServiceServer(s grpc.ServiceRegistrar, srv SignServiceServer) {
	s.RegisterService(&SignService_ServiceDesc, srv)
}

func _SignService_OrgAccountAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrgAccountAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).OrgAccountAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignService/OrgAccountAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).OrgAccountAdd(ctx, req.(*OrgAccountAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignService_UserAccountAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccountAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).UserAccountAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignService/UserAccountAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).UserAccountAdd(ctx, req.(*UserAccountAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignService_TemplateCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateByTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).TemplateCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignService/TemplateCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).TemplateCreate(ctx, req.(*CreateByTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignService_TemplateGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).TemplateGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignService/TemplateGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).TemplateGet(ctx, req.(*TemplateGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignService_FlowOneStepCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlowOneStepCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).FlowOneStepCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignService/FlowOneStepCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).FlowOneStepCreate(ctx, req.(*FlowOneStepCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignService_VoucherIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoucherIssueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).VoucherIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignService/VoucherIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).VoucherIssue(ctx, req.(*VoucherIssueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SignService_Notice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SignServiceServer).Notice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SignService/Notice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SignServiceServer).Notice(ctx, req.(*NoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SignService_ServiceDesc is the grpc.ServiceDesc for SignService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SignService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SignService",
	HandlerType: (*SignServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrgAccountAdd",
			Handler:    _SignService_OrgAccountAdd_Handler,
		},
		{
			MethodName: "UserAccountAdd",
			Handler:    _SignService_UserAccountAdd_Handler,
		},
		{
			MethodName: "TemplateCreate",
			Handler:    _SignService_TemplateCreate_Handler,
		},
		{
			MethodName: "TemplateGet",
			Handler:    _SignService_TemplateGet_Handler,
		},
		{
			MethodName: "FlowOneStepCreate",
			Handler:    _SignService_FlowOneStepCreate_Handler,
		},
		{
			MethodName: "VoucherIssue",
			Handler:    _SignService_VoucherIssue_Handler,
		},
		{
			MethodName: "Notice",
			Handler:    _SignService_Notice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sign_service.proto",
}