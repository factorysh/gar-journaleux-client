// Code generated by protoc-gen-go. DO NOT EDIT.
// source: journaleux.proto

package rpc

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

type Predicate struct {
	Project              string            `protobuf:"bytes,1,opt,name=Project" json:"Project,omitempty"`
	Lines                int32             `protobuf:"varint,2,opt,name=Lines" json:"Lines,omitempty"`
	Follow               bool              `protobuf:"varint,3,opt,name=Follow" json:"Follow,omitempty"`
	Since                uint32            `protobuf:"varint,4,opt,name=Since" json:"Since,omitempty"`
	Until                uint32            `protobuf:"varint,5,opt,name=Until" json:"Until,omitempty"`
	Priority             uint32            `protobuf:"varint,6,opt,name=Priority" json:"Priority,omitempty"`
	Regexp               string            `protobuf:"bytes,7,opt,name=Regexp" json:"Regexp,omitempty"`
	Fields               map[string]string `protobuf:"bytes,8,rep,name=Fields" json:"Fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Predicate) Reset()         { *m = Predicate{} }
func (m *Predicate) String() string { return proto.CompactTextString(m) }
func (*Predicate) ProtoMessage()    {}
func (*Predicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_journaleux_91e203cf4eb20d4a, []int{0}
}
func (m *Predicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Predicate.Unmarshal(m, b)
}
func (m *Predicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Predicate.Marshal(b, m, deterministic)
}
func (dst *Predicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Predicate.Merge(dst, src)
}
func (m *Predicate) XXX_Size() int {
	return xxx_messageInfo_Predicate.Size(m)
}
func (m *Predicate) XXX_DiscardUnknown() {
	xxx_messageInfo_Predicate.DiscardUnknown(m)
}

var xxx_messageInfo_Predicate proto.InternalMessageInfo

func (m *Predicate) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Predicate) GetLines() int32 {
	if m != nil {
		return m.Lines
	}
	return 0
}

func (m *Predicate) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

func (m *Predicate) GetSince() uint32 {
	if m != nil {
		return m.Since
	}
	return 0
}

func (m *Predicate) GetUntil() uint32 {
	if m != nil {
		return m.Until
	}
	return 0
}

func (m *Predicate) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Predicate) GetRegexp() string {
	if m != nil {
		return m.Regexp
	}
	return ""
}

func (m *Predicate) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type Event struct {
	Monotonic            uint64            `protobuf:"varint,1,opt,name=Monotonic" json:"Monotonic,omitempty"`
	Realtime             uint64            `protobuf:"varint,2,opt,name=Realtime" json:"Realtime,omitempty"`
	Message              string            `protobuf:"bytes,3,opt,name=Message" json:"Message,omitempty"`
	Priority             uint32            `protobuf:"varint,4,opt,name=Priority" json:"Priority,omitempty"`
	Fields               map[string]string `protobuf:"bytes,5,rep,name=Fields" json:"Fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_journaleux_91e203cf4eb20d4a, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetMonotonic() uint64 {
	if m != nil {
		return m.Monotonic
	}
	return 0
}

func (m *Event) GetRealtime() uint64 {
	if m != nil {
		return m.Realtime
	}
	return 0
}

func (m *Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Event) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Event) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func init() {
	proto.RegisterType((*Predicate)(nil), "rpc.Predicate")
	proto.RegisterMapType((map[string]string)(nil), "rpc.Predicate.FieldsEntry")
	proto.RegisterType((*Event)(nil), "rpc.Event")
	proto.RegisterMapType((map[string]string)(nil), "rpc.Event.FieldsEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JournaleuxClient is the client API for Journaleux service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JournaleuxClient interface {
	Tail(ctx context.Context, in *Predicate, opts ...grpc.CallOption) (Journaleux_TailClient, error)
}

type journaleuxClient struct {
	cc *grpc.ClientConn
}

func NewJournaleuxClient(cc *grpc.ClientConn) JournaleuxClient {
	return &journaleuxClient{cc}
}

func (c *journaleuxClient) Tail(ctx context.Context, in *Predicate, opts ...grpc.CallOption) (Journaleux_TailClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Journaleux_serviceDesc.Streams[0], "/rpc.Journaleux/Tail", opts...)
	if err != nil {
		return nil, err
	}
	x := &journaleuxTailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Journaleux_TailClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type journaleuxTailClient struct {
	grpc.ClientStream
}

func (x *journaleuxTailClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JournaleuxServer is the server API for Journaleux service.
type JournaleuxServer interface {
	Tail(*Predicate, Journaleux_TailServer) error
}

func RegisterJournaleuxServer(s *grpc.Server, srv JournaleuxServer) {
	s.RegisterService(&_Journaleux_serviceDesc, srv)
}

func _Journaleux_Tail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Predicate)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JournaleuxServer).Tail(m, &journaleuxTailServer{stream})
}

type Journaleux_TailServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type journaleuxTailServer struct {
	grpc.ServerStream
}

func (x *journaleuxTailServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Journaleux_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Journaleux",
	HandlerType: (*JournaleuxServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tail",
			Handler:       _Journaleux_Tail_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "journaleux.proto",
}

func init() { proto.RegisterFile("journaleux.proto", fileDescriptor_journaleux_91e203cf4eb20d4a) }

var fileDescriptor_journaleux_91e203cf4eb20d4a = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x6e, 0xe2, 0x30,
	0x14, 0x85, 0xc7, 0xe4, 0x07, 0x72, 0xd1, 0x8c, 0x90, 0x35, 0x42, 0x56, 0x34, 0x8b, 0x88, 0xc5,
	0x28, 0xab, 0xa8, 0xa2, 0x5d, 0xb4, 0xdd, 0xc3, 0xa2, 0x2a, 0x12, 0x72, 0xdb, 0x07, 0x48, 0xc3,
	0x15, 0x32, 0x75, 0xed, 0xc8, 0x31, 0x14, 0x5e, 0xa7, 0x8f, 0xd6, 0x27, 0xa9, 0xe2, 0x84, 0x50,
	0xba, 0xee, 0x2e, 0xdf, 0xb5, 0x63, 0x9d, 0xef, 0xe8, 0xc2, 0x68, 0xa3, 0xb7, 0x46, 0xe5, 0x12,
	0xb7, 0xfb, 0xac, 0x34, 0xda, 0x6a, 0xea, 0x99, 0xb2, 0x98, 0xbc, 0xf7, 0x20, 0x5a, 0x1a, 0x5c,
	0x89, 0x22, 0xb7, 0x48, 0x19, 0xf4, 0x97, 0x46, 0x6f, 0xb0, 0xb0, 0x8c, 0x24, 0x24, 0x8d, 0xf8,
	0x11, 0xe9, 0x5f, 0x08, 0xee, 0x85, 0xc2, 0x8a, 0xf5, 0x12, 0x92, 0x06, 0xbc, 0x01, 0x3a, 0x86,
	0x70, 0xae, 0xa5, 0xd4, 0x6f, 0xcc, 0x4b, 0x48, 0x3a, 0xe0, 0x2d, 0xd5, 0xb7, 0x1f, 0x84, 0x2a,
	0x90, 0xf9, 0x09, 0x49, 0x7f, 0xf3, 0x06, 0xea, 0xe9, 0x93, 0xb2, 0x42, 0xb2, 0xa0, 0x99, 0x3a,
	0xa0, 0x31, 0x0c, 0x96, 0x46, 0x68, 0x23, 0xec, 0x81, 0x85, 0xee, 0xa0, 0xe3, 0xfa, 0x7d, 0x8e,
	0x6b, 0xdc, 0x97, 0xac, 0xef, 0xe2, 0xb4, 0x44, 0xa7, 0x10, 0xce, 0x05, 0xca, 0x55, 0xc5, 0x06,
	0x89, 0x97, 0x0e, 0xa7, 0x71, 0x66, 0xca, 0x22, 0xeb, 0x3c, 0xb2, 0xe6, 0x70, 0xa6, 0xac, 0x39,
	0xf0, 0xf6, 0x66, 0x7c, 0x03, 0xc3, 0x2f, 0x63, 0x3a, 0x02, 0xef, 0x05, 0x0f, 0xad, 0x66, 0xfd,
	0x59, 0xc7, 0xdb, 0xe5, 0x72, 0x8b, 0x4e, 0x31, 0xe2, 0x0d, 0xdc, 0xf6, 0xae, 0xc9, 0xe4, 0x83,
	0x40, 0x30, 0xdb, 0xa1, 0xb2, 0xf4, 0x1f, 0x44, 0x0b, 0xad, 0xb4, 0xd5, 0x4a, 0x14, 0xee, 0x5f,
	0x9f, 0x9f, 0x06, 0xb5, 0x0a, 0xc7, 0x5c, 0x5a, 0xf1, 0xda, 0x3c, 0xe2, 0xf3, 0x8e, 0xeb, 0x6a,
	0x17, 0x58, 0x55, 0xf9, 0x1a, 0x5d, 0x57, 0x11, 0x3f, 0xe2, 0x59, 0x01, 0xfe, 0xb7, 0x02, 0xb2,
	0x4e, 0x34, 0x70, 0xa2, 0x63, 0x27, 0xea, 0xb2, 0xfc, 0xb0, 0xe4, 0xf4, 0x0a, 0xe0, 0xae, 0x5b,
	0x11, 0xfa, 0x1f, 0xfc, 0xc7, 0x5c, 0x48, 0xfa, 0xe7, 0xbc, 0xd9, 0x18, 0x4e, 0x01, 0x26, 0xbf,
	0x2e, 0xc8, 0x73, 0xe8, 0x76, 0xe9, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x86, 0x45, 0x91,
	0x5f, 0x02, 0x00, 0x00,
}
