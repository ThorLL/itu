package main

import (
	"context"
	"google.golang.org/grpc"
)

type CourseClient interface {
	// Sends a greeting
	updateCourse(ctx context.Context, in *course, opts ...grpc.CallOption)
}

type courseClient struct {
	cc grpc.ClientConnInterface
}

func (c *courseClient) updateCourse(ctx context.Context, in *course, opts ...grpc.CallOption) {
	out := "updated"
	c.cc.Invoke(ctx, "/updateCourse", in, out, opts...)
}

func NewCourseClient(cc grpc.ClientConnInterface) CourseClient {
	return &courseClient{cc}
}