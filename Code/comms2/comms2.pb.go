// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comms2.proto

package comms2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request_Log struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_Log) Reset()         { *m = Request_Log{} }
func (m *Request_Log) String() string { return proto.CompactTextString(m) }
func (*Request_Log) ProtoMessage()    {}
func (*Request_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdeb7a6dcd152929, []int{0}
}

func (m *Request_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Log.Unmarshal(m, b)
}
func (m *Request_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Log.Marshal(b, m, deterministic)
}
func (m *Request_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Log.Merge(m, src)
}
func (m *Request_Log) XXX_Size() int {
	return xxx_messageInfo_Request_Log.Size(m)
}
func (m *Request_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Log proto.InternalMessageInfo

type Response_Log struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Log) Reset()         { *m = Response_Log{} }
func (m *Response_Log) String() string { return proto.CompactTextString(m) }
func (*Response_Log) ProtoMessage()    {}
func (*Response_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdeb7a6dcd152929, []int{1}
}

func (m *Response_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Log.Unmarshal(m, b)
}
func (m *Response_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Log.Marshal(b, m, deterministic)
}
func (m *Response_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Log.Merge(m, src)
}
func (m *Response_Log) XXX_Size() int {
	return xxx_messageInfo_Response_Log.Size(m)
}
func (m *Response_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Log proto.InternalMessageInfo

type Request_Propuesta struct {
	Propuesta            string   `protobuf:"bytes,1,opt,name=propuesta,proto3" json:"propuesta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_Propuesta) Reset()         { *m = Request_Propuesta{} }
func (m *Request_Propuesta) String() string { return proto.CompactTextString(m) }
func (*Request_Propuesta) ProtoMessage()    {}
func (*Request_Propuesta) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdeb7a6dcd152929, []int{2}
}

func (m *Request_Propuesta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Propuesta.Unmarshal(m, b)
}
func (m *Request_Propuesta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Propuesta.Marshal(b, m, deterministic)
}
func (m *Request_Propuesta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Propuesta.Merge(m, src)
}
func (m *Request_Propuesta) XXX_Size() int {
	return xxx_messageInfo_Request_Propuesta.Size(m)
}
func (m *Request_Propuesta) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Propuesta.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Propuesta proto.InternalMessageInfo

func (m *Request_Propuesta) GetPropuesta() string {
	if m != nil {
		return m.Propuesta
	}
	return ""
}

type Response_Propuesta struct {
	Estado               int32    `protobuf:"varint,1,opt,name=estado,proto3" json:"estado,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Propuesta) Reset()         { *m = Response_Propuesta{} }
func (m *Response_Propuesta) String() string { return proto.CompactTextString(m) }
func (*Response_Propuesta) ProtoMessage()    {}
func (*Response_Propuesta) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdeb7a6dcd152929, []int{3}
}

func (m *Response_Propuesta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Propuesta.Unmarshal(m, b)
}
func (m *Response_Propuesta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Propuesta.Marshal(b, m, deterministic)
}
func (m *Response_Propuesta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Propuesta.Merge(m, src)
}
func (m *Response_Propuesta) XXX_Size() int {
	return xxx_messageInfo_Response_Propuesta.Size(m)
}
func (m *Response_Propuesta) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Propuesta.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Propuesta proto.InternalMessageInfo

func (m *Response_Propuesta) GetEstado() int32 {
	if m != nil {
		return m.Estado
	}
	return 0
}

type Request_Catlogo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_Catlogo) Reset()         { *m = Request_Catlogo{} }
func (m *Request_Catlogo) String() string { return proto.CompactTextString(m) }
func (*Request_Catlogo) ProtoMessage()    {}
func (*Request_Catlogo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdeb7a6dcd152929, []int{4}
}

func (m *Request_Catlogo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Catlogo.Unmarshal(m, b)
}
func (m *Request_Catlogo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Catlogo.Marshal(b, m, deterministic)
}
func (m *Request_Catlogo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Catlogo.Merge(m, src)
}
func (m *Request_Catlogo) XXX_Size() int {
	return xxx_messageInfo_Request_Catlogo.Size(m)
}
func (m *Request_Catlogo) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Catlogo.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Catlogo proto.InternalMessageInfo

type Response_Catalogo struct {
	Libros               string   `protobuf:"bytes,1,opt,name=libros,proto3" json:"libros,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Catalogo) Reset()         { *m = Response_Catalogo{} }
func (m *Response_Catalogo) String() string { return proto.CompactTextString(m) }
func (*Response_Catalogo) ProtoMessage()    {}
func (*Response_Catalogo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdeb7a6dcd152929, []int{5}
}

func (m *Response_Catalogo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Catalogo.Unmarshal(m, b)
}
func (m *Response_Catalogo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Catalogo.Marshal(b, m, deterministic)
}
func (m *Response_Catalogo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Catalogo.Merge(m, src)
}
func (m *Response_Catalogo) XXX_Size() int {
	return xxx_messageInfo_Response_Catalogo.Size(m)
}
func (m *Response_Catalogo) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Catalogo.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Catalogo proto.InternalMessageInfo

func (m *Response_Catalogo) GetLibros() string {
	if m != nil {
		return m.Libros
	}
	return ""
}

type Request_Libro struct {
	Numero               int32    `protobuf:"varint,1,opt,name=numero,proto3" json:"numero,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_Libro) Reset()         { *m = Request_Libro{} }
func (m *Request_Libro) String() string { return proto.CompactTextString(m) }
func (*Request_Libro) ProtoMessage()    {}
func (*Request_Libro) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdeb7a6dcd152929, []int{6}
}

func (m *Request_Libro) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_Libro.Unmarshal(m, b)
}
func (m *Request_Libro) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_Libro.Marshal(b, m, deterministic)
}
func (m *Request_Libro) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_Libro.Merge(m, src)
}
func (m *Request_Libro) XXX_Size() int {
	return xxx_messageInfo_Request_Libro.Size(m)
}
func (m *Request_Libro) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_Libro.DiscardUnknown(m)
}

var xxx_messageInfo_Request_Libro proto.InternalMessageInfo

func (m *Request_Libro) GetNumero() int32 {
	if m != nil {
		return m.Numero
	}
	return 0
}

type Response_Libro struct {
	Ubicaciones          string   `protobuf:"bytes,1,opt,name=ubicaciones,proto3" json:"ubicaciones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Libro) Reset()         { *m = Response_Libro{} }
func (m *Response_Libro) String() string { return proto.CompactTextString(m) }
func (*Response_Libro) ProtoMessage()    {}
func (*Response_Libro) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdeb7a6dcd152929, []int{7}
}

func (m *Response_Libro) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Libro.Unmarshal(m, b)
}
func (m *Response_Libro) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Libro.Marshal(b, m, deterministic)
}
func (m *Response_Libro) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Libro.Merge(m, src)
}
func (m *Response_Libro) XXX_Size() int {
	return xxx_messageInfo_Response_Libro.Size(m)
}
func (m *Response_Libro) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Libro.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Libro proto.InternalMessageInfo

func (m *Response_Libro) GetUbicaciones() string {
	if m != nil {
		return m.Ubicaciones
	}
	return ""
}

func init() {
	proto.RegisterType((*Request_Log)(nil), "comms2.Request_Log")
	proto.RegisterType((*Response_Log)(nil), "comms2.Response_Log")
	proto.RegisterType((*Request_Propuesta)(nil), "comms2.Request_Propuesta")
	proto.RegisterType((*Response_Propuesta)(nil), "comms2.Response_Propuesta")
	proto.RegisterType((*Request_Catlogo)(nil), "comms2.Request_Catlogo")
	proto.RegisterType((*Response_Catalogo)(nil), "comms2.Response_Catalogo")
	proto.RegisterType((*Request_Libro)(nil), "comms2.Request_Libro")
	proto.RegisterType((*Response_Libro)(nil), "comms2.Response_Libro")
}

func init() {
	proto.RegisterFile("comms2.proto", fileDescriptor_fdeb7a6dcd152929)
}

var fileDescriptor_fdeb7a6dcd152929 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x4f, 0x83, 0x40,
	0x14, 0x6c, 0x35, 0x12, 0x79, 0xb4, 0x35, 0x5d, 0xb5, 0xda, 0x8d, 0x87, 0x66, 0x2f, 0x9a, 0x68,
	0x9a, 0x88, 0x77, 0x63, 0x52, 0x8f, 0x1e, 0x1a, 0xfe, 0x80, 0x01, 0xba, 0x69, 0x48, 0x0a, 0x0f,
	0x77, 0xe1, 0x1f, 0xf8, 0xc3, 0xcb, 0x6e, 0xf7, 0x83, 0x94, 0x13, 0xcc, 0xbc, 0x37, 0xf3, 0x98,
	0x09, 0x30, 0xc9, 0xb1, 0x2c, 0x65, 0xbc, 0xae, 0x05, 0x36, 0x48, 0x82, 0x13, 0x62, 0x53, 0x88,
	0x12, 0xfe, 0xd7, 0x72, 0xd9, 0xfc, 0xfe, 0xe0, 0x9e, 0xcd, 0x60, 0x92, 0x70, 0x59, 0x63, 0x25,
	0xb9, 0xc6, 0xef, 0x30, 0xb7, 0xe3, 0xad, 0xc0, 0x5a, 0xbd, 0xa4, 0xe4, 0x09, 0xc2, 0xda, 0x82,
	0xc7, 0xf1, 0x6a, 0xfc, 0x12, 0x26, 0x9e, 0x60, 0x6f, 0x40, 0x9c, 0x85, 0xd7, 0x2c, 0x20, 0x50,
	0xcf, 0x1d, 0x6a, 0xc1, 0x55, 0x62, 0x10, 0x9b, 0xc3, 0x8d, 0x3d, 0xb0, 0x49, 0x9b, 0x03, 0xee,
	0x91, 0xbd, 0xaa, 0x9b, 0xc6, 0xa0, 0xe3, 0x52, 0x45, 0x2a, 0xfd, 0xa1, 0xc8, 0x04, 0x4a, 0x73,
	0xd0, 0x20, 0xf6, 0x0c, 0x53, 0xf7, 0xfd, 0x8a, 0x51, 0x8b, 0x55, 0x5b, 0x72, 0xe1, 0x0e, 0x9d,
	0x10, 0x8b, 0x61, 0xe6, 0x93, 0xe9, 0xcd, 0x15, 0x44, 0x6d, 0x56, 0xe4, 0x69, 0x5e, 0x60, 0xc5,
	0xad, 0x6f, 0x9f, 0x8a, 0xff, 0x2f, 0x20, 0xd8, 0xe8, 0x9e, 0x48, 0x0c, 0x97, 0x5d, 0x1f, 0xe4,
	0x76, 0x6d, 0x5a, 0xec, 0x95, 0x46, 0xef, 0x3c, 0xd9, 0xab, 0x6e, 0x44, 0xbe, 0x21, 0xf4, 0x05,
	0x2c, 0xcf, 0x95, 0x6e, 0x44, 0xe9, 0x40, 0xef, 0x66, 0x9d, 0xcb, 0x17, 0x5c, 0xbb, 0x16, 0x1e,
	0xce, 0x4d, 0x4c, 0x67, 0x74, 0x39, 0xb0, 0xb0, 0x9a, 0xce, 0xe1, 0x13, 0xa2, 0x2d, 0xdf, 0x15,
	0xc2, 0xe4, 0xbe, 0x1f, 0x64, 0x50, 0x34, 0x5d, 0x0c, 0x53, 0x28, 0x9e, 0x8d, 0xb2, 0x40, 0xff,
	0x32, 0x1f, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x19, 0xf0, 0xb1, 0x42, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Comms2Client is the client API for Comms2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Comms2Client interface {
	Log(ctx context.Context, in *Request_Log, opts ...grpc.CallOption) (*Response_Log, error)
	Propuesta(ctx context.Context, in *Request_Propuesta, opts ...grpc.CallOption) (*Response_Propuesta, error)
	Catalogo(ctx context.Context, in *Request_Catlogo, opts ...grpc.CallOption) (*Response_Catalogo, error)
	Pedir_Libro(ctx context.Context, in *Request_Libro, opts ...grpc.CallOption) (*Response_Libro, error)
}

type comms2Client struct {
	cc grpc.ClientConnInterface
}

func NewComms2Client(cc grpc.ClientConnInterface) Comms2Client {
	return &comms2Client{cc}
}

func (c *comms2Client) Log(ctx context.Context, in *Request_Log, opts ...grpc.CallOption) (*Response_Log, error) {
	out := new(Response_Log)
	err := c.cc.Invoke(ctx, "/comms2.Comms2/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comms2Client) Propuesta(ctx context.Context, in *Request_Propuesta, opts ...grpc.CallOption) (*Response_Propuesta, error) {
	out := new(Response_Propuesta)
	err := c.cc.Invoke(ctx, "/comms2.Comms2/Propuesta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comms2Client) Catalogo(ctx context.Context, in *Request_Catlogo, opts ...grpc.CallOption) (*Response_Catalogo, error) {
	out := new(Response_Catalogo)
	err := c.cc.Invoke(ctx, "/comms2.Comms2/Catalogo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comms2Client) Pedir_Libro(ctx context.Context, in *Request_Libro, opts ...grpc.CallOption) (*Response_Libro, error) {
	out := new(Response_Libro)
	err := c.cc.Invoke(ctx, "/comms2.Comms2/Pedir_Libro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Comms2Server is the server API for Comms2 service.
type Comms2Server interface {
	Log(context.Context, *Request_Log) (*Response_Log, error)
	Propuesta(context.Context, *Request_Propuesta) (*Response_Propuesta, error)
	Catalogo(context.Context, *Request_Catlogo) (*Response_Catalogo, error)
	Pedir_Libro(context.Context, *Request_Libro) (*Response_Libro, error)
}

// UnimplementedComms2Server can be embedded to have forward compatible implementations.
type UnimplementedComms2Server struct {
}

func (*UnimplementedComms2Server) Log(ctx context.Context, req *Request_Log) (*Response_Log, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (*UnimplementedComms2Server) Propuesta(ctx context.Context, req *Request_Propuesta) (*Response_Propuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Propuesta not implemented")
}
func (*UnimplementedComms2Server) Catalogo(ctx context.Context, req *Request_Catlogo) (*Response_Catalogo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Catalogo not implemented")
}
func (*UnimplementedComms2Server) Pedir_Libro(ctx context.Context, req *Request_Libro) (*Response_Libro, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pedir_Libro not implemented")
}

func RegisterComms2Server(s *grpc.Server, srv Comms2Server) {
	s.RegisterService(&_Comms2_serviceDesc, srv)
}

func _Comms2_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Log)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Comms2Server).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms2.Comms2/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Comms2Server).Log(ctx, req.(*Request_Log))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms2_Propuesta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Propuesta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Comms2Server).Propuesta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms2.Comms2/Propuesta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Comms2Server).Propuesta(ctx, req.(*Request_Propuesta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms2_Catalogo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Catlogo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Comms2Server).Catalogo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms2.Comms2/Catalogo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Comms2Server).Catalogo(ctx, req.(*Request_Catlogo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Comms2_Pedir_Libro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request_Libro)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Comms2Server).Pedir_Libro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comms2.Comms2/Pedir_Libro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Comms2Server).Pedir_Libro(ctx, req.(*Request_Libro))
	}
	return interceptor(ctx, in, info, handler)
}

var _Comms2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comms2.Comms2",
	HandlerType: (*Comms2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _Comms2_Log_Handler,
		},
		{
			MethodName: "Propuesta",
			Handler:    _Comms2_Propuesta_Handler,
		},
		{
			MethodName: "Catalogo",
			Handler:    _Comms2_Catalogo_Handler,
		},
		{
			MethodName: "Pedir_Libro",
			Handler:    _Comms2_Pedir_Libro_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comms2.proto",
}
