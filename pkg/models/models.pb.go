// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/models/models.proto

package models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Task struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Done                 bool     `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Id                   uint64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b9750c941d452a7, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Task) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Task) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Task) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type TaskList struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b9750c941d452a7, []int{1}
}

func (m *TaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskList.Unmarshal(m, b)
}
func (m *TaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskList.Marshal(b, m, deterministic)
}
func (m *TaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskList.Merge(m, src)
}
func (m *TaskList) XXX_Size() int {
	return xxx_messageInfo_TaskList.Size(m)
}
func (m *TaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskList proto.InternalMessageInfo

func (m *TaskList) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Text struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Desk                 string   `protobuf:"bytes,2,opt,name=desk,proto3" json:"desk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b9750c941d452a7, []int{2}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Text) GetDesk() string {
	if m != nil {
		return m.Desk
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b9750c941d452a7, []int{3}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Task)(nil), "models.Task")
	proto.RegisterType((*TaskList)(nil), "models.TaskList")
	proto.RegisterType((*Text)(nil), "models.Text")
	proto.RegisterType((*Void)(nil), "models.Void")
}

func init() { proto.RegisterFile("pkg/models/models.proto", fileDescriptor_2b9750c941d452a7) }

var fileDescriptor_2b9750c941d452a7 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x50, 0x4d, 0x4f, 0x03, 0x21,
	0x14, 0x2c, 0xbb, 0x74, 0x53, 0x9f, 0xc6, 0x18, 0x2e, 0x92, 0x9e, 0x36, 0x78, 0xe1, 0x84, 0x49,
	0xfd, 0x05, 0xde, 0x3d, 0x91, 0xa6, 0x77, 0xec, 0x23, 0x86, 0xa0, 0xd2, 0xf8, 0xde, 0xa1, 0x3f,
	0xdf, 0xc0, 0xe2, 0x47, 0x4f, 0x0c, 0x33, 0x93, 0x99, 0xc9, 0x83, 0xfb, 0x53, 0x7e, 0x7b, 0xfc,
	0x28, 0x18, 0xdf, 0xa9, 0x3f, 0xee, 0xf4, 0x55, 0xb8, 0xa8, 0x69, 0xf9, 0x19, 0x0f, 0x72, 0x1f,
	0x28, 0x2b, 0x05, 0x12, 0x03, 0x07, 0x2d, 0x66, 0x61, 0xaf, 0x7c, 0xc3, 0x8d, 0x2b, 0x9f, 0x51,
	0x0f, 0xb3, 0xb0, 0x1b, 0xdf, 0x70, 0xe3, 0x22, 0x1d, 0xf5, 0xd8, 0x7d, 0x91, 0x8e, 0xea, 0x16,
	0x86, 0x84, 0x5a, 0xce, 0xc2, 0x4a, 0x3f, 0x24, 0x34, 0x0e, 0x36, 0x35, 0xf3, 0x25, 0x11, 0x2b,
	0x03, 0x6b, 0x0e, 0x94, 0x49, 0x8b, 0x79, 0xb4, 0xd7, 0xbb, 0x1b, 0xd7, 0x57, 0x54, 0x83, 0x5f,
	0x24, 0xe3, 0x40, 0xee, 0xe3, 0x99, 0x6b, 0x36, 0xc7, 0x33, 0xff, 0x6c, 0xe0, 0xce, 0x61, 0xa4,
	0xdc, 0x36, 0x2c, 0x7d, 0xd9, 0x4c, 0x20, 0x0f, 0x25, 0xe1, 0xee, 0x00, 0xeb, 0x1a, 0x43, 0xca,
	0x82, 0x6c, 0x65, 0xbf, 0xe9, 0x55, 0xde, 0xde, 0xfd, 0xef, 0xaa, 0xba, 0x59, 0xa9, 0x07, 0x18,
	0x9f, 0x11, 0xff, 0x8c, 0xb5, 0x77, 0x7b, 0x31, 0xca, 0xac, 0x5e, 0xa7, 0x76, 0xa2, 0xa7, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x71, 0xf3, 0xa2, 0x3d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TasksClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error)
	Add(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Task, error)
}

type tasksClient struct {
	cc *grpc.ClientConn
}

func NewTasksClient(cc *grpc.ClientConn) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error) {
	out := new(TaskList)
	err := c.cc.Invoke(ctx, "/models.Tasks/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) Add(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/models.Tasks/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
type TasksServer interface {
	List(context.Context, *Void) (*TaskList, error)
	Add(context.Context, *Text) (*Task, error)
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Tasks/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Tasks/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).Add(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Tasks_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Tasks_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/models/models.proto",
}