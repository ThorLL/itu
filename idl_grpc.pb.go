// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package itu

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

// CourseMethodsClient is the client API for CourseMethods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseMethodsClient interface {
	UpdateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Empty, error)
	CreateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Empty, error)
	DeleteCourse(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	GetCourseByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Course, error)
	GetCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Courses, error)
}

type courseMethodsClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseMethodsClient(cc grpc.ClientConnInterface) CourseMethodsClient {
	return &courseMethodsClient{cc}
}

func (c *courseMethodsClient) UpdateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/itu.CourseMethods/updateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseMethodsClient) CreateCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/itu.CourseMethods/CreateCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseMethodsClient) DeleteCourse(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/itu.CourseMethods/deleteCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseMethodsClient) GetCourseByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/itu.CourseMethods/getCourseByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseMethodsClient) GetCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := c.cc.Invoke(ctx, "/itu.CourseMethods/getCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseMethodsServer is the server API for CourseMethods service.
// All implementations must embed UnimplementedCourseMethodsServer
// for forward compatibility
type CourseMethodsServer interface {
	UpdateCourse(context.Context, *Course) (*Empty, error)
	CreateCourse(context.Context, *Course) (*Empty, error)
	DeleteCourse(context.Context, *ID) (*Empty, error)
	GetCourseByID(context.Context, *ID) (*Course, error)
	GetCourses(context.Context, *Empty) (*Courses, error)
	mustEmbedUnimplementedCourseMethodsServer()
}

// UnimplementedCourseMethodsServer must be embedded to have forward compatible implementations.
type UnimplementedCourseMethodsServer struct {
}

func (UnimplementedCourseMethodsServer) UpdateCourse(context.Context, *Course) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourse not implemented")
}
func (UnimplementedCourseMethodsServer) CreateCourse(context.Context, *Course) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseMethodsServer) DeleteCourse(context.Context, *ID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourse not implemented")
}
func (UnimplementedCourseMethodsServer) GetCourseByID(context.Context, *ID) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseByID not implemented")
}
func (UnimplementedCourseMethodsServer) GetCourses(context.Context, *Empty) (*Courses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedCourseMethodsServer) mustEmbedUnimplementedCourseMethodsServer() {}

// UnsafeCourseMethodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseMethodsServer will
// result in compilation errors.
type UnsafeCourseMethodsServer interface {
	mustEmbedUnimplementedCourseMethodsServer()
}

func RegisterCourseMethodsServer(s grpc.ServiceRegistrar, srv CourseMethodsServer) {
	s.RegisterService(&CourseMethods_ServiceDesc, srv)
}

func _CourseMethods_UpdateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseMethodsServer).UpdateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itu.CourseMethods/updateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseMethodsServer).UpdateCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseMethods_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseMethodsServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itu.CourseMethods/CreateCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseMethodsServer).CreateCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseMethods_DeleteCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseMethodsServer).DeleteCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itu.CourseMethods/deleteCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseMethodsServer).DeleteCourse(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseMethods_GetCourseByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseMethodsServer).GetCourseByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itu.CourseMethods/getCourseByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseMethodsServer).GetCourseByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseMethods_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseMethodsServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/itu.CourseMethods/getCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseMethodsServer).GetCourses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseMethods_ServiceDesc is the grpc.ServiceDesc for CourseMethods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseMethods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "itu.CourseMethods",
	HandlerType: (*CourseMethodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "updateCourse",
			Handler:    _CourseMethods_UpdateCourse_Handler,
		},
		{
			MethodName: "CreateCourse",
			Handler:    _CourseMethods_CreateCourse_Handler,
		},
		{
			MethodName: "deleteCourse",
			Handler:    _CourseMethods_DeleteCourse_Handler,
		},
		{
			MethodName: "getCourseByID",
			Handler:    _CourseMethods_GetCourseByID_Handler,
		},
		{
			MethodName: "getCourses",
			Handler:    _CourseMethods_GetCourses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl.proto",
}
