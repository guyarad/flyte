// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/task.proto

package admin

import (
	fmt "fmt"
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Represents a request structure to create a revision of a task.
type TaskCreateRequest struct {
	// id represents the unique identifier of the task.
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Represents the specification for task.
	Spec                 *TaskSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TaskCreateRequest) Reset()         { *m = TaskCreateRequest{} }
func (m *TaskCreateRequest) String() string { return proto.CompactTextString(m) }
func (*TaskCreateRequest) ProtoMessage()    {}
func (*TaskCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9204120d588b2162, []int{0}
}

func (m *TaskCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateRequest.Unmarshal(m, b)
}
func (m *TaskCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateRequest.Marshal(b, m, deterministic)
}
func (m *TaskCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateRequest.Merge(m, src)
}
func (m *TaskCreateRequest) XXX_Size() int {
	return xxx_messageInfo_TaskCreateRequest.Size(m)
}
func (m *TaskCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateRequest proto.InternalMessageInfo

func (m *TaskCreateRequest) GetId() *core.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *TaskCreateRequest) GetSpec() *TaskSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// Represents a response structure if task creation succeeds.
type TaskCreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskCreateResponse) Reset()         { *m = TaskCreateResponse{} }
func (m *TaskCreateResponse) String() string { return proto.CompactTextString(m) }
func (*TaskCreateResponse) ProtoMessage()    {}
func (*TaskCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9204120d588b2162, []int{1}
}

func (m *TaskCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskCreateResponse.Unmarshal(m, b)
}
func (m *TaskCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskCreateResponse.Marshal(b, m, deterministic)
}
func (m *TaskCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskCreateResponse.Merge(m, src)
}
func (m *TaskCreateResponse) XXX_Size() int {
	return xxx_messageInfo_TaskCreateResponse.Size(m)
}
func (m *TaskCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskCreateResponse proto.InternalMessageInfo

// Flyte workflows are composed of many ordered tasks. That is small, reusable, self-contained logical blocks
// arranged to process workflow inputs and produce a deterministic set of outputs.
// Tasks can come in many varieties tuned for specialized behavior.
type Task struct {
	// id represents the unique identifier of the task.
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// closure encapsulates all the fields that maps to a compiled version of the task.
	Closure              *TaskClosure `protobuf:"bytes,2,opt,name=closure,proto3" json:"closure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_9204120d588b2162, []int{2}
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

func (m *Task) GetId() *core.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Task) GetClosure() *TaskClosure {
	if m != nil {
		return m.Closure
	}
	return nil
}

// Represents a list of tasks returned from the admin.
type TaskList struct {
	// A list of tasks returned based on the request.
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9204120d588b2162, []int{3}
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

func (m *TaskList) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Represents a structure that encapsulates the user-configured specification of the task.
type TaskSpec struct {
	// Template of the task that encapsulates all the metadata of the task.
	Template             *core.TaskTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TaskSpec) Reset()         { *m = TaskSpec{} }
func (m *TaskSpec) String() string { return proto.CompactTextString(m) }
func (*TaskSpec) ProtoMessage()    {}
func (*TaskSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_9204120d588b2162, []int{4}
}

func (m *TaskSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskSpec.Unmarshal(m, b)
}
func (m *TaskSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskSpec.Marshal(b, m, deterministic)
}
func (m *TaskSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskSpec.Merge(m, src)
}
func (m *TaskSpec) XXX_Size() int {
	return xxx_messageInfo_TaskSpec.Size(m)
}
func (m *TaskSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TaskSpec proto.InternalMessageInfo

func (m *TaskSpec) GetTemplate() *core.TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

// Compute task attributes which include values derived from the TaskSpec, as well as plugin-specific data
// and task metadata.
type TaskClosure struct {
	// Represents the compiled representation of the task from the specification provided.
	CompiledTask *core.CompiledTask `protobuf:"bytes,1,opt,name=compiled_task,json=compiledTask,proto3" json:"compiled_task,omitempty"`
	// Time at which the task was created.
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TaskClosure) Reset()         { *m = TaskClosure{} }
func (m *TaskClosure) String() string { return proto.CompactTextString(m) }
func (*TaskClosure) ProtoMessage()    {}
func (*TaskClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_9204120d588b2162, []int{5}
}

func (m *TaskClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskClosure.Unmarshal(m, b)
}
func (m *TaskClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskClosure.Marshal(b, m, deterministic)
}
func (m *TaskClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskClosure.Merge(m, src)
}
func (m *TaskClosure) XXX_Size() int {
	return xxx_messageInfo_TaskClosure.Size(m)
}
func (m *TaskClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskClosure.DiscardUnknown(m)
}

var xxx_messageInfo_TaskClosure proto.InternalMessageInfo

func (m *TaskClosure) GetCompiledTask() *core.CompiledTask {
	if m != nil {
		return m.CompiledTask
	}
	return nil
}

func (m *TaskClosure) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskCreateRequest)(nil), "flyteidl.admin.TaskCreateRequest")
	proto.RegisterType((*TaskCreateResponse)(nil), "flyteidl.admin.TaskCreateResponse")
	proto.RegisterType((*Task)(nil), "flyteidl.admin.Task")
	proto.RegisterType((*TaskList)(nil), "flyteidl.admin.TaskList")
	proto.RegisterType((*TaskSpec)(nil), "flyteidl.admin.TaskSpec")
	proto.RegisterType((*TaskClosure)(nil), "flyteidl.admin.TaskClosure")
}

func init() { proto.RegisterFile("flyteidl/admin/task.proto", fileDescriptor_9204120d588b2162) }

var fileDescriptor_9204120d588b2162 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0xef, 0xd3, 0x30,
	0x10, 0xc5, 0x95, 0xf2, 0x2f, 0xb4, 0x2e, 0x20, 0x61, 0x75, 0x48, 0x5b, 0x04, 0x55, 0xa6, 0x82,
	0xc0, 0x96, 0x8a, 0xaa, 0x8a, 0x0d, 0xc8, 0x84, 0xd4, 0xc9, 0x74, 0x62, 0xa9, 0x12, 0xe7, 0x9a,
	0x5a, 0x4d, 0x62, 0x13, 0x3b, 0x03, 0x5f, 0x81, 0x4f, 0x8d, 0xec, 0x38, 0x51, 0x83, 0xca, 0xf0,
	0x1f, 0xcf, 0xf7, 0xcb, 0x7b, 0x77, 0x97, 0x87, 0x16, 0xe7, 0xe2, 0xb7, 0x01, 0x91, 0x15, 0x34,
	0xc9, 0x4a, 0x51, 0x51, 0x93, 0xe8, 0x2b, 0x51, 0xb5, 0x34, 0x12, 0xbf, 0xec, 0x5a, 0xc4, 0xb5,
	0x96, 0x6f, 0x7a, 0x94, 0xcb, 0x1a, 0xa8, 0xc8, 0xa0, 0x32, 0xe2, 0x2c, 0xa0, 0x6e, 0xf9, 0xe5,
	0x62, 0xd8, 0xb7, 0x4a, 0xda, 0xb7, 0x5e, 0x0f, 0x5b, 0x5c, 0x96, 0x4a, 0x14, 0xfd, 0x87, 0x6f,
	0x73, 0x29, 0xf3, 0x02, 0xa8, 0xab, 0xd2, 0xe6, 0x4c, 0x8d, 0x28, 0x41, 0x9b, 0xa4, 0x54, 0x2d,
	0x10, 0x15, 0xe8, 0xd5, 0x31, 0xd1, 0xd7, 0xb8, 0x86, 0xc4, 0x00, 0x83, 0x5f, 0x0d, 0x68, 0x83,
	0xdf, 0xa1, 0x91, 0xc8, 0xc2, 0x60, 0x1d, 0x6c, 0x66, 0xdb, 0x05, 0xe9, 0x67, 0xb5, 0x06, 0xe4,
	0x7b, 0x3f, 0x1b, 0x1b, 0x89, 0x0c, 0x7f, 0x40, 0x0f, 0x5a, 0x01, 0x0f, 0x47, 0x0e, 0x0e, 0xc9,
	0x70, 0x31, 0x62, 0xb5, 0x7f, 0x28, 0xe0, 0xcc, 0x51, 0xd1, 0x1c, 0xe1, 0x5b, 0x37, 0xad, 0x64,
	0xa5, 0x21, 0xba, 0xa0, 0x07, 0xfb, 0xfa, 0x18, 0xdb, 0x1d, 0x7a, 0xc6, 0x0b, 0xa9, 0x9b, 0x1a,
	0xbc, 0xf3, 0xea, 0x9e, 0x73, 0xdc, 0x22, 0xac, 0x63, 0xa3, 0x03, 0x9a, 0xd8, 0xf7, 0x83, 0xd0,
	0x06, 0xbf, 0x47, 0x63, 0x77, 0xc7, 0x30, 0x58, 0x3f, 0xd9, 0xcc, 0xb6, 0xf3, 0x7b, 0x02, 0xac,
	0x45, 0xf0, 0x1c, 0x8d, 0x8d, 0xbc, 0x42, 0xe5, 0xcc, 0xa6, 0xac, 0x2d, 0xa2, 0xb8, 0x55, 0xb3,
	0xfb, 0xe1, 0x3d, 0x9a, 0x18, 0x28, 0x55, 0x91, 0x18, 0xf0, 0x1b, 0xac, 0xfe, 0xd9, 0xc0, 0xa2,
	0x47, 0x8f, 0xb0, 0x1e, 0x8e, 0xfe, 0x04, 0x68, 0x76, 0x33, 0x2b, 0xfe, 0x82, 0x5e, 0xf8, 0x7f,
	0x98, 0x9d, 0xac, 0xf9, 0x7f, 0xd4, 0x62, 0xcf, 0xb8, 0x29, 0x9f, 0xf3, 0x9b, 0x0a, 0x7f, 0x46,
	0x88, 0xbb, 0x03, 0x67, 0xa7, 0xc4, 0xf8, 0xf3, 0x2c, 0x49, 0x1b, 0x04, 0xd2, 0x05, 0x81, 0x1c,
	0xbb, 0x20, 0xb0, 0xa9, 0xa7, 0xbf, 0x9a, 0x6f, 0xfb, 0x9f, 0xbb, 0x5c, 0x98, 0x4b, 0x93, 0x12,
	0x2e, 0x4b, 0xea, 0x1c, 0x65, 0x9d, 0xd3, 0x3e, 0x62, 0x39, 0x54, 0x54, 0xa5, 0x1f, 0x73, 0x49,
	0x87, 0xd9, 0x4e, 0x9f, 0x3a, 0xdd, 0x4f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x38, 0xe5, 0xba,
	0x67, 0xf4, 0x02, 0x00, 0x00,
}
