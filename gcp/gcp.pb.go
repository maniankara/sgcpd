// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gcp.proto

package gcp

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CopyFile struct {
	// Types that are valid to be assigned to Chunktype:
	//	*CopyFile_FilePath
	//	*CopyFile_FileData
	Chunktype            isCopyFile_Chunktype `protobuf_oneof:"chunktype"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CopyFile) Reset()         { *m = CopyFile{} }
func (m *CopyFile) String() string { return proto.CompactTextString(m) }
func (*CopyFile) ProtoMessage()    {}
func (*CopyFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_gcp_21d47f9b6369817b, []int{0}
}
func (m *CopyFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CopyFile.Unmarshal(m, b)
}
func (m *CopyFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CopyFile.Marshal(b, m, deterministic)
}
func (dst *CopyFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CopyFile.Merge(dst, src)
}
func (m *CopyFile) XXX_Size() int {
	return xxx_messageInfo_CopyFile.Size(m)
}
func (m *CopyFile) XXX_DiscardUnknown() {
	xxx_messageInfo_CopyFile.DiscardUnknown(m)
}

var xxx_messageInfo_CopyFile proto.InternalMessageInfo

type isCopyFile_Chunktype interface {
	isCopyFile_Chunktype()
}

type CopyFile_FilePath struct {
	FilePath string `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3,oneof"`
}
type CopyFile_FileData struct {
	FileData []byte `protobuf:"bytes,2,opt,name=file_data,json=fileData,proto3,oneof"`
}

func (*CopyFile_FilePath) isCopyFile_Chunktype() {}
func (*CopyFile_FileData) isCopyFile_Chunktype() {}

func (m *CopyFile) GetChunktype() isCopyFile_Chunktype {
	if m != nil {
		return m.Chunktype
	}
	return nil
}

func (m *CopyFile) GetFilePath() string {
	if x, ok := m.GetChunktype().(*CopyFile_FilePath); ok {
		return x.FilePath
	}
	return ""
}

func (m *CopyFile) GetFileData() []byte {
	if x, ok := m.GetChunktype().(*CopyFile_FileData); ok {
		return x.FileData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CopyFile) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CopyFile_OneofMarshaler, _CopyFile_OneofUnmarshaler, _CopyFile_OneofSizer, []interface{}{
		(*CopyFile_FilePath)(nil),
		(*CopyFile_FileData)(nil),
	}
}

func _CopyFile_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CopyFile)
	// chunktype
	switch x := m.Chunktype.(type) {
	case *CopyFile_FilePath:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.FilePath)
	case *CopyFile_FileData:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.FileData)
	case nil:
	default:
		return fmt.Errorf("CopyFile.Chunktype has unexpected type %T", x)
	}
	return nil
}

func _CopyFile_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CopyFile)
	switch tag {
	case 1: // chunktype.file_path
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Chunktype = &CopyFile_FilePath{x}
		return true, err
	case 2: // chunktype.file_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Chunktype = &CopyFile_FileData{x}
		return true, err
	default:
		return false, nil
	}
}

func _CopyFile_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CopyFile)
	// chunktype
	switch x := m.Chunktype.(type) {
	case *CopyFile_FilePath:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.FilePath)))
		n += len(x.FilePath)
	case *CopyFile_FileData:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.FileData)))
		n += len(x.FileData)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TransferPath struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferPath) Reset()         { *m = TransferPath{} }
func (m *TransferPath) String() string { return proto.CompactTextString(m) }
func (*TransferPath) ProtoMessage()    {}
func (*TransferPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_gcp_21d47f9b6369817b, []int{1}
}
func (m *TransferPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferPath.Unmarshal(m, b)
}
func (m *TransferPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferPath.Marshal(b, m, deterministic)
}
func (dst *TransferPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferPath.Merge(dst, src)
}
func (m *TransferPath) XXX_Size() int {
	return xxx_messageInfo_TransferPath.Size(m)
}
func (m *TransferPath) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferPath.DiscardUnknown(m)
}

var xxx_messageInfo_TransferPath proto.InternalMessageInfo

func (m *TransferPath) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*CopyFile)(nil), "gcp.CopyFile")
	proto.RegisterType((*TransferPath)(nil), "gcp.TransferPath")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GcpClient is the client API for Gcp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GcpClient interface {
	// transfer file from local (client) to remote (server)
	CopyTo(ctx context.Context, opts ...grpc.CallOption) (Gcp_CopyToClient, error)
	// transfer file from remote (server) to local (client)
	CopyFrom(ctx context.Context, in *TransferPath, opts ...grpc.CallOption) (Gcp_CopyFromClient, error)
}

type gcpClient struct {
	cc *grpc.ClientConn
}

func NewGcpClient(cc *grpc.ClientConn) GcpClient {
	return &gcpClient{cc}
}

func (c *gcpClient) CopyTo(ctx context.Context, opts ...grpc.CallOption) (Gcp_CopyToClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Gcp_serviceDesc.Streams[0], "/gcp.Gcp/CopyTo", opts...)
	if err != nil {
		return nil, err
	}
	x := &gcpCopyToClient{stream}
	return x, nil
}

type Gcp_CopyToClient interface {
	Send(*CopyFile) error
	CloseAndRecv() (*TransferPath, error)
	grpc.ClientStream
}

type gcpCopyToClient struct {
	grpc.ClientStream
}

func (x *gcpCopyToClient) Send(m *CopyFile) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gcpCopyToClient) CloseAndRecv() (*TransferPath, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TransferPath)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gcpClient) CopyFrom(ctx context.Context, in *TransferPath, opts ...grpc.CallOption) (Gcp_CopyFromClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Gcp_serviceDesc.Streams[1], "/gcp.Gcp/CopyFrom", opts...)
	if err != nil {
		return nil, err
	}
	x := &gcpCopyFromClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Gcp_CopyFromClient interface {
	Recv() (*CopyFile, error)
	grpc.ClientStream
}

type gcpCopyFromClient struct {
	grpc.ClientStream
}

func (x *gcpCopyFromClient) Recv() (*CopyFile, error) {
	m := new(CopyFile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GcpServer is the server API for Gcp service.
type GcpServer interface {
	// transfer file from local (client) to remote (server)
	CopyTo(Gcp_CopyToServer) error
	// transfer file from remote (server) to local (client)
	CopyFrom(*TransferPath, Gcp_CopyFromServer) error
}

func RegisterGcpServer(s *grpc.Server, srv GcpServer) {
	s.RegisterService(&_Gcp_serviceDesc, srv)
}

func _Gcp_CopyTo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GcpServer).CopyTo(&gcpCopyToServer{stream})
}

type Gcp_CopyToServer interface {
	SendAndClose(*TransferPath) error
	Recv() (*CopyFile, error)
	grpc.ServerStream
}

type gcpCopyToServer struct {
	grpc.ServerStream
}

func (x *gcpCopyToServer) SendAndClose(m *TransferPath) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gcpCopyToServer) Recv() (*CopyFile, error) {
	m := new(CopyFile)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Gcp_CopyFrom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransferPath)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GcpServer).CopyFrom(m, &gcpCopyFromServer{stream})
}

type Gcp_CopyFromServer interface {
	Send(*CopyFile) error
	grpc.ServerStream
}

type gcpCopyFromServer struct {
	grpc.ServerStream
}

func (x *gcpCopyFromServer) Send(m *CopyFile) error {
	return x.ServerStream.SendMsg(m)
}

var _Gcp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gcp.Gcp",
	HandlerType: (*GcpServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CopyTo",
			Handler:       _Gcp_CopyTo_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CopyFrom",
			Handler:       _Gcp_CopyFrom_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gcp.proto",
}

func init() { proto.RegisterFile("gcp.proto", fileDescriptor_gcp_21d47f9b6369817b) }

var fileDescriptor_gcp_21d47f9b6369817b = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4f, 0x2e, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x4f, 0x2e, 0x50, 0x0a, 0xe5, 0xe2, 0x70, 0xce,
	0x2f, 0xa8, 0x74, 0xcb, 0xcc, 0x49, 0x15, 0x92, 0xe5, 0xe2, 0x4c, 0xcb, 0xcc, 0x49, 0x8d, 0x2f,
	0x48, 0x2c, 0xc9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0xf4, 0x60, 0x08, 0xe2, 0x00, 0x09, 0x05,
	0x24, 0x96, 0x64, 0xc0, 0xa5, 0x53, 0x12, 0x4b, 0x12, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x60,
	0xd2, 0x2e, 0x89, 0x25, 0x89, 0x4e, 0xdc, 0x5c, 0x9c, 0xc9, 0x19, 0xa5, 0x79, 0xd9, 0x25, 0x95,
	0x05, 0xa9, 0x4a, 0x4a, 0x5c, 0x3c, 0x21, 0x45, 0x89, 0x79, 0xc5, 0x69, 0xa9, 0x45, 0x60, 0xbd,
	0x42, 0x5c, 0x2c, 0x08, 0x53, 0x83, 0xc0, 0x6c, 0xa3, 0x74, 0x2e, 0x66, 0xf7, 0xe4, 0x02, 0x21,
	0x3d, 0x2e, 0x36, 0x90, 0x0b, 0x42, 0xf2, 0x85, 0x78, 0xf5, 0x40, 0x8e, 0x83, 0x39, 0x47, 0x4a,
	0x10, 0xcc, 0x45, 0x36, 0x46, 0x89, 0x41, 0x83, 0x51, 0xc8, 0x00, 0xea, 0xe2, 0xa2, 0xfc, 0x5c,
	0x21, 0x4c, 0x25, 0x52, 0xa8, 0x86, 0x28, 0x31, 0x18, 0x30, 0x26, 0xb1, 0x81, 0xfd, 0x6b, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x3c, 0x9e, 0x06, 0xfc, 0x00, 0x00, 0x00,
}
