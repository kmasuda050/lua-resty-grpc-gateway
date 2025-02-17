// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ex                   []*ComplexMsg `protobuf:"bytes,2,rep,name=ex,proto3" json:"ex,omitempty"`
	Jobs                 []string      `protobuf:"bytes,3,rep,name=jobs,proto3" json:"jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetEx() []*ComplexMsg {
	if m != nil {
		return m.Ex
	}
	return nil
}

func (m *HelloRequest) GetJobs() []string {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type ComplexMsg struct {
	Displayname          string               `protobuf:"bytes,1,opt,name=displayname,proto3" json:"displayname,omitempty"`
	Foo                  *YetAnotherNestedMsg `protobuf:"bytes,2,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ComplexMsg) Reset()         { *m = ComplexMsg{} }
func (m *ComplexMsg) String() string { return proto.CompactTextString(m) }
func (*ComplexMsg) ProtoMessage()    {}
func (*ComplexMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *ComplexMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexMsg.Unmarshal(m, b)
}
func (m *ComplexMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexMsg.Marshal(b, m, deterministic)
}
func (m *ComplexMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexMsg.Merge(m, src)
}
func (m *ComplexMsg) XXX_Size() int {
	return xxx_messageInfo_ComplexMsg.Size(m)
}
func (m *ComplexMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexMsg proto.InternalMessageInfo

func (m *ComplexMsg) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *ComplexMsg) GetFoo() *YetAnotherNestedMsg {
	if m != nil {
		return m.Foo
	}
	return nil
}

type YetAnotherNestedMsg struct {
	Grades               []int32  `protobuf:"varint,1,rep,packed,name=grades,proto3" json:"grades,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YetAnotherNestedMsg) Reset()         { *m = YetAnotherNestedMsg{} }
func (m *YetAnotherNestedMsg) String() string { return proto.CompactTextString(m) }
func (*YetAnotherNestedMsg) ProtoMessage()    {}
func (*YetAnotherNestedMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *YetAnotherNestedMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YetAnotherNestedMsg.Unmarshal(m, b)
}
func (m *YetAnotherNestedMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YetAnotherNestedMsg.Marshal(b, m, deterministic)
}
func (m *YetAnotherNestedMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YetAnotherNestedMsg.Merge(m, src)
}
func (m *YetAnotherNestedMsg) XXX_Size() int {
	return xxx_messageInfo_YetAnotherNestedMsg.Size(m)
}
func (m *YetAnotherNestedMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_YetAnotherNestedMsg.DiscardUnknown(m)
}

var xxx_messageInfo_YetAnotherNestedMsg proto.InternalMessageInfo

func (m *YetAnotherNestedMsg) GetGrades() []int32 {
	if m != nil {
		return m.Grades
	}
	return nil
}

type HelloReply struct {
	Message              string               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ReplyAt              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=reply_at,json=replyAt,proto3" json:"reply_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{3}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetReplyAt() *timestamp.Timestamp {
	if m != nil {
		return m.ReplyAt
	}
	return nil
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*ComplexMsg)(nil), "helloworld.ComplexMsg")
	proto.RegisterType((*YetAnotherNestedMsg)(nil), "helloworld.YetAnotherNestedMsg")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4f, 0xf3, 0x30,
	0x10, 0xc5, 0xbf, 0x24, 0x1f, 0x6d, 0xb9, 0x32, 0x20, 0x23, 0x55, 0x56, 0x97, 0x5a, 0x19, 0x50,
	0x16, 0x52, 0x51, 0xc4, 0x8a, 0x54, 0x31, 0x00, 0x03, 0x0c, 0x81, 0x85, 0x01, 0x90, 0xa3, 0x5c,
	0xd3, 0x22, 0xa7, 0x36, 0xb6, 0x2b, 0x92, 0xff, 0x1e, 0x39, 0x4d, 0x54, 0x0f, 0xdd, 0xee, 0xee,
	0xbd, 0xd3, 0xef, 0xee, 0xc1, 0xf9, 0x1a, 0x85, 0x90, 0xbf, 0x52, 0x8b, 0x22, 0x55, 0x5a, 0x5a,
	0x49, 0xe0, 0x30, 0x99, 0xce, 0x4a, 0x29, 0x4b, 0x81, 0xf3, 0x56, 0xc9, 0x77, 0xab, 0xb9, 0xdd,
	0x54, 0x68, 0x2c, 0xaf, 0xd4, 0xde, 0x1c, 0x7f, 0xc2, 0xd9, 0xa3, 0xb3, 0x67, 0xf8, 0xb3, 0x43,
	0x63, 0x09, 0x81, 0xff, 0x5b, 0x5e, 0x21, 0x0d, 0x58, 0x90, 0x9c, 0x66, 0x6d, 0x4d, 0x2e, 0x21,
	0xc4, 0x9a, 0x86, 0x2c, 0x4a, 0xc6, 0x8b, 0x49, 0xea, 0xf1, 0xee, 0x65, 0xa5, 0x04, 0xd6, 0xcf,
	0xa6, 0xcc, 0x42, 0xac, 0xdd, 0xee, 0xb7, 0xcc, 0x0d, 0x8d, 0x58, 0xe4, 0x76, 0x5d, 0x1d, 0x73,
	0x80, 0x83, 0x8b, 0x30, 0x18, 0x17, 0x1b, 0xa3, 0x04, 0x6f, 0x3c, 0x88, 0x3f, 0x22, 0xd7, 0x10,
	0xad, 0xa4, 0xa4, 0x21, 0x0b, 0x92, 0xf1, 0x62, 0xe6, 0xc3, 0xde, 0xd1, 0x2e, 0xb7, 0xd2, 0xae,
	0x51, 0xbf, 0xa0, 0xb1, 0x58, 0x38, 0xaa, 0xf3, 0xc6, 0x57, 0x70, 0x71, 0x44, 0x23, 0x13, 0x18,
	0x94, 0x9a, 0x17, 0x68, 0x68, 0xc0, 0xa2, 0xe4, 0x24, 0xeb, 0xba, 0xf8, 0x03, 0xa0, 0xfb, 0x58,
	0x89, 0x86, 0x50, 0x18, 0x56, 0x68, 0x0c, 0x2f, 0xfb, 0x6b, 0xfa, 0x96, 0xdc, 0xc2, 0x48, 0x3b,
	0xcb, 0x17, 0xb7, 0xdd, 0x39, 0xd3, 0x74, 0x9f, 0x66, 0xda, 0xa7, 0x99, 0xbe, 0xf5, 0x69, 0x66,
	0xc3, 0xd6, 0xbb, 0xb4, 0x8b, 0x27, 0x18, 0x3e, 0x68, 0x44, 0x8b, 0x9a, 0xdc, 0xc1, 0xe8, 0x95,
	0x37, 0x2d, 0x8c, 0x50, 0xff, 0x15, 0x3f, 0xf1, 0xe9, 0xe4, 0x88, 0xa2, 0x44, 0x13, 0xff, 0xcb,
	0x07, 0x2d, 0xe7, 0xe6, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x11, 0x05, 0x27, 0xe3, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
