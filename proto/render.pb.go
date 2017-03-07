// Code generated by protoc-gen-go.
// source: proto/render.proto
// DO NOT EDIT!

package renderdemo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RenderRequest struct {
	// GCS path to write output image into.
	GcsOutputBase string `protobuf:"bytes,1,opt,name=gcs_output_base,json=gcsOutputBase" json:"gcs_output_base,omitempty"`
	// Product logo image to transform, stored locally.
	// TODO(shadams): pull image from elsewhere, making render
	// more reusable.
	ImgPath string `protobuf:"bytes,2,opt,name=img_path,json=imgPath" json:"img_path,omitempty"`
	// Frame number (in range [0,36))
	Frame int64 `protobuf:"varint,3,opt,name=frame" json:"frame,omitempty"`
}

func (m *RenderRequest) Reset()                    { *m = RenderRequest{} }
func (m *RenderRequest) String() string            { return proto.CompactTextString(m) }
func (*RenderRequest) ProtoMessage()               {}
func (*RenderRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *RenderRequest) GetGcsOutputBase() string {
	if m != nil {
		return m.GcsOutputBase
	}
	return ""
}

func (m *RenderRequest) GetImgPath() string {
	if m != nil {
		return m.ImgPath
	}
	return ""
}

func (m *RenderRequest) GetFrame() int64 {
	if m != nil {
		return m.Frame
	}
	return 0
}

type RenderResponse struct {
	// GCS path image was written to.
	GcsOutput string `protobuf:"bytes,1,opt,name=gcs_output,json=gcsOutput" json:"gcs_output,omitempty"`
}

func (m *RenderResponse) Reset()                    { *m = RenderResponse{} }
func (m *RenderResponse) String() string            { return proto.CompactTextString(m) }
func (*RenderResponse) ProtoMessage()               {}
func (*RenderResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *RenderResponse) GetGcsOutput() string {
	if m != nil {
		return m.GcsOutput
	}
	return ""
}

func init() {
	proto.RegisterType((*RenderRequest)(nil), "renderdemo.RenderRequest")
	proto.RegisterType((*RenderResponse)(nil), "renderdemo.RenderResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Render service

type RenderClient interface {
	RenderFrame(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*RenderResponse, error)
}

type renderClient struct {
	cc *grpc.ClientConn
}

func NewRenderClient(cc *grpc.ClientConn) RenderClient {
	return &renderClient{cc}
}

func (c *renderClient) RenderFrame(ctx context.Context, in *RenderRequest, opts ...grpc.CallOption) (*RenderResponse, error) {
	out := new(RenderResponse)
	err := grpc.Invoke(ctx, "/renderdemo.Render/RenderFrame", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Render service

type RenderServer interface {
	RenderFrame(context.Context, *RenderRequest) (*RenderResponse, error)
}

func RegisterRenderServer(s *grpc.Server, srv RenderServer) {
	s.RegisterService(&_Render_serviceDesc, srv)
}

func _Render_RenderFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RenderServer).RenderFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/renderdemo.Render/RenderFrame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RenderServer).RenderFrame(ctx, req.(*RenderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Render_serviceDesc = grpc.ServiceDesc{
	ServiceName: "renderdemo.Render",
	HandlerType: (*RenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RenderFrame",
			Handler:    _Render_RenderFrame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/render.proto",
}

func init() { proto.RegisterFile("proto/render.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4a, 0xcd, 0x4b, 0x49, 0x2d, 0xd2, 0x03, 0x73, 0x84, 0xb8, 0x20, 0xbc, 0x94,
	0xd4, 0xdc, 0x7c, 0xa5, 0x0c, 0x2e, 0xde, 0x20, 0x30, 0x2f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8,
	0x44, 0x48, 0x8d, 0x8b, 0x3f, 0x3d, 0xb9, 0x38, 0x3e, 0xbf, 0xb4, 0xa4, 0xa0, 0xb4, 0x24, 0x3e,
	0x29, 0xb1, 0x38, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x37, 0x3d, 0xb9, 0xd8, 0x1f,
	0x2c, 0xea, 0x94, 0x58, 0x9c, 0x2a, 0x24, 0xc9, 0xc5, 0x91, 0x99, 0x9b, 0x1e, 0x5f, 0x90, 0x58,
	0x92, 0x21, 0xc1, 0x04, 0x56, 0xc0, 0x9e, 0x99, 0x9b, 0x1e, 0x90, 0x58, 0x92, 0x21, 0x24, 0xc2,
	0xc5, 0x9a, 0x56, 0x94, 0x98, 0x9b, 0x2a, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1c, 0x04, 0xe1, 0x28,
	0xe9, 0x73, 0xf1, 0xc1, 0x6c, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe5, 0xe2, 0x42,
	0x58, 0x05, 0xb5, 0x85, 0x13, 0x6e, 0x8b, 0x91, 0x1f, 0x17, 0x1b, 0x44, 0x83, 0x90, 0x0b, 0x17,
	0x37, 0x84, 0xe5, 0x06, 0x32, 0x49, 0x48, 0x52, 0x0f, 0xe1, 0x01, 0x3d, 0x14, 0xd7, 0x4b, 0x49,
	0x61, 0x93, 0x82, 0x58, 0xe7, 0xc4, 0x13, 0x85, 0xe4, 0xf1, 0x24, 0x36, 0x70, 0x58, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x4a, 0xc6, 0x5d, 0x21, 0x01, 0x00, 0x00,
}
