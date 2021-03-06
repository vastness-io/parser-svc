// Code generated by protoc-gen-go. DO NOT EDIT.
// source: parser.proto

package parser

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

type ParserRequest struct {
	RemoteUrl string      `protobuf:"bytes,1,opt,name=remote_url,json=remoteUrl" json:"remote_url,omitempty"`
	Version   string      `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	FileInfo  []*FileInfo `protobuf:"bytes,3,rep,name=file_info,json=fileInfo" json:"file_info,omitempty"`
	Type      string      `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
}

func (m *ParserRequest) Reset()                    { *m = ParserRequest{} }
func (m *ParserRequest) String() string            { return proto.CompactTextString(m) }
func (*ParserRequest) ProtoMessage()               {}
func (*ParserRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ParserRequest) GetRemoteUrl() string {
	if m != nil {
		return m.RemoteUrl
	}
	return ""
}

func (m *ParserRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ParserRequest) GetFileInfo() []*FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

func (m *ParserRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ParserResponse struct {
	MavenResponse []*MavenResponse `protobuf:"bytes,1,rep,name=maven_response,json=mavenResponse" json:"maven_response,omitempty"`
}

func (m *ParserResponse) Reset()                    { *m = ParserResponse{} }
func (m *ParserResponse) String() string            { return proto.CompactTextString(m) }
func (*ParserResponse) ProtoMessage()               {}
func (*ParserResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ParserResponse) GetMavenResponse() []*MavenResponse {
	if m != nil {
		return m.MavenResponse
	}
	return nil
}

type FileInfo struct {
	Language  string   `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	FileNames []string `protobuf:"bytes,2,rep,name=file_names,json=fileNames" json:"file_names,omitempty"`
}

func (m *FileInfo) Reset()                    { *m = FileInfo{} }
func (m *FileInfo) String() string            { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()               {}
func (*FileInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *FileInfo) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *FileInfo) GetFileNames() []string {
	if m != nil {
		return m.FileNames
	}
	return nil
}

func init() {
	proto.RegisterType((*ParserRequest)(nil), "parser.ParserRequest")
	proto.RegisterType((*ParserResponse)(nil), "parser.ParserResponse")
	proto.RegisterType((*FileInfo)(nil), "parser.FileInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Parser service

type ParserClient interface {
	Analyse(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error)
}

type parserClient struct {
	cc *grpc.ClientConn
}

func NewParserClient(cc *grpc.ClientConn) ParserClient {
	return &parserClient{cc}
}

func (c *parserClient) Analyse(ctx context.Context, in *ParserRequest, opts ...grpc.CallOption) (*ParserResponse, error) {
	out := new(ParserResponse)
	err := grpc.Invoke(ctx, "/parser.Parser/Analyse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Parser service

type ParserServer interface {
	Analyse(context.Context, *ParserRequest) (*ParserResponse, error)
}

func RegisterParserServer(s *grpc.Server, srv ParserServer) {
	s.RegisterService(&_Parser_serviceDesc, srv)
}

func _Parser_Analyse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServer).Analyse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.Parser/Analyse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServer).Analyse(ctx, req.(*ParserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Parser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "parser.Parser",
	HandlerType: (*ParserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Analyse",
			Handler:    _Parser_Analyse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser.proto",
}

func init() { proto.RegisterFile("parser.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0x3a, 0xba, 0xf6, 0xcd, 0x0d, 0x79, 0xa0, 0x84, 0x82, 0x50, 0x7a, 0xda, 0xc5,
	0x1d, 0xe6, 0xc5, 0x83, 0x17, 0x05, 0x05, 0x0f, 0x0e, 0x29, 0x78, 0x2e, 0x11, 0x5e, 0x47, 0x21,
	0x4d, 0x6a, 0xd2, 0x0e, 0xf6, 0x19, 0xfc, 0xd2, 0x92, 0xa4, 0x19, 0xea, 0xed, 0xfd, 0x7f, 0x2f,
	0xbc, 0xfc, 0xde, 0x83, 0x8b, 0x9e, 0x6b, 0x43, 0x7a, 0xdb, 0x6b, 0x35, 0x28, 0x4c, 0x7c, 0xca,
	0x97, 0x1d, 0x3f, 0x92, 0xf4, 0xb0, 0xfc, 0x8e, 0x60, 0xf5, 0xee, 0x78, 0x45, 0x5f, 0x23, 0x99,
	0x01, 0x6f, 0x00, 0x34, 0x75, 0x6a, 0xa0, 0x7a, 0xd4, 0x82, 0x45, 0x45, 0xb4, 0xc9, 0xaa, 0xcc,
	0x93, 0x0f, 0x2d, 0x90, 0xc1, 0xe2, 0x48, 0xda, 0xb4, 0x4a, 0xb2, 0x99, 0xeb, 0x85, 0x88, 0xb7,
	0x90, 0x35, 0xad, 0xa0, 0xba, 0x95, 0x8d, 0x62, 0x71, 0x11, 0x6f, 0x96, 0xbb, 0xcb, 0xed, 0x64,
	0xf0, 0xd2, 0x0a, 0x7a, 0x95, 0x8d, 0xaa, 0xd2, 0x66, 0xaa, 0x10, 0x61, 0x3e, 0x9c, 0x7a, 0x62,
	0x73, 0x37, 0xc5, 0xd5, 0xe5, 0x1e, 0xd6, 0x41, 0xc6, 0xf4, 0x4a, 0x1a, 0xc2, 0x07, 0x58, 0x3b,
	0xdd, 0x5a, 0x4f, 0x84, 0x45, 0x6e, 0xf2, 0x55, 0x98, 0xfc, 0x66, 0xbb, 0xe1, 0x79, 0xb5, 0xea,
	0x7e, 0xc7, 0xf2, 0x19, 0xd2, 0xf0, 0x33, 0xe6, 0x90, 0x0a, 0x2e, 0x0f, 0x23, 0x3f, 0xd0, 0xb4,
	0xd5, 0x39, 0xdb, 0x9d, 0x9d, 0xba, 0xe4, 0x1d, 0x19, 0x36, 0x2b, 0x62, 0xbb, 0xb3, 0x25, 0x7b,
	0x0b, 0x76, 0x4f, 0x90, 0x78, 0x2d, 0xbc, 0x87, 0xc5, 0xa3, 0xe4, 0xe2, 0x64, 0x08, 0xcf, 0x06,
	0x7f, 0xce, 0x97, 0x5f, 0xff, 0xc7, 0x5e, 0xe5, 0x33, 0x71, 0xf7, 0xbe, 0xfb, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xaf, 0xfd, 0x25, 0x20, 0x94, 0x01, 0x00, 0x00,
}
