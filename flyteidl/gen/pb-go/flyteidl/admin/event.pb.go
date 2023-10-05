// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/event.proto

package admin

import (
	fmt "fmt"
	math "math"

	event "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/event"
	proto "github.com/golang/protobuf/proto"
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

// Indicates that a sent event was not used to update execution state due to
// the referenced execution already being terminated (and therefore ineligible
// for further state transitions).
type EventErrorAlreadyInTerminalState struct {
	// +required
	CurrentPhase         string   `protobuf:"bytes,1,opt,name=current_phase,json=currentPhase,proto3" json:"current_phase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventErrorAlreadyInTerminalState) Reset()         { *m = EventErrorAlreadyInTerminalState{} }
func (m *EventErrorAlreadyInTerminalState) String() string { return proto.CompactTextString(m) }
func (*EventErrorAlreadyInTerminalState) ProtoMessage()    {}
func (*EventErrorAlreadyInTerminalState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{0}
}

func (m *EventErrorAlreadyInTerminalState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventErrorAlreadyInTerminalState.Unmarshal(m, b)
}
func (m *EventErrorAlreadyInTerminalState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventErrorAlreadyInTerminalState.Marshal(b, m, deterministic)
}
func (m *EventErrorAlreadyInTerminalState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventErrorAlreadyInTerminalState.Merge(m, src)
}
func (m *EventErrorAlreadyInTerminalState) XXX_Size() int {
	return xxx_messageInfo_EventErrorAlreadyInTerminalState.Size(m)
}
func (m *EventErrorAlreadyInTerminalState) XXX_DiscardUnknown() {
	xxx_messageInfo_EventErrorAlreadyInTerminalState.DiscardUnknown(m)
}

var xxx_messageInfo_EventErrorAlreadyInTerminalState proto.InternalMessageInfo

func (m *EventErrorAlreadyInTerminalState) GetCurrentPhase() string {
	if m != nil {
		return m.CurrentPhase
	}
	return ""
}

// Indicates an event was rejected because it came from a different cluster than
// is on record as running the execution.
type EventErrorIncompatibleCluster struct {
	// The cluster which has been recorded as processing the execution.
	// +required
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventErrorIncompatibleCluster) Reset()         { *m = EventErrorIncompatibleCluster{} }
func (m *EventErrorIncompatibleCluster) String() string { return proto.CompactTextString(m) }
func (*EventErrorIncompatibleCluster) ProtoMessage()    {}
func (*EventErrorIncompatibleCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{1}
}

func (m *EventErrorIncompatibleCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventErrorIncompatibleCluster.Unmarshal(m, b)
}
func (m *EventErrorIncompatibleCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventErrorIncompatibleCluster.Marshal(b, m, deterministic)
}
func (m *EventErrorIncompatibleCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventErrorIncompatibleCluster.Merge(m, src)
}
func (m *EventErrorIncompatibleCluster) XXX_Size() int {
	return xxx_messageInfo_EventErrorIncompatibleCluster.Size(m)
}
func (m *EventErrorIncompatibleCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_EventErrorIncompatibleCluster.DiscardUnknown(m)
}

var xxx_messageInfo_EventErrorIncompatibleCluster proto.InternalMessageInfo

func (m *EventErrorIncompatibleCluster) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

// Indicates why a sent event was not used to update execution.
type EventFailureReason struct {
	// +required
	//
	// Types that are valid to be assigned to Reason:
	//	*EventFailureReason_AlreadyInTerminalState
	//	*EventFailureReason_IncompatibleCluster
	Reason               isEventFailureReason_Reason `protobuf_oneof:"reason"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *EventFailureReason) Reset()         { *m = EventFailureReason{} }
func (m *EventFailureReason) String() string { return proto.CompactTextString(m) }
func (*EventFailureReason) ProtoMessage()    {}
func (*EventFailureReason) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{2}
}

func (m *EventFailureReason) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventFailureReason.Unmarshal(m, b)
}
func (m *EventFailureReason) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventFailureReason.Marshal(b, m, deterministic)
}
func (m *EventFailureReason) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFailureReason.Merge(m, src)
}
func (m *EventFailureReason) XXX_Size() int {
	return xxx_messageInfo_EventFailureReason.Size(m)
}
func (m *EventFailureReason) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFailureReason.DiscardUnknown(m)
}

var xxx_messageInfo_EventFailureReason proto.InternalMessageInfo

type isEventFailureReason_Reason interface {
	isEventFailureReason_Reason()
}

type EventFailureReason_AlreadyInTerminalState struct {
	AlreadyInTerminalState *EventErrorAlreadyInTerminalState `protobuf:"bytes,1,opt,name=already_in_terminal_state,json=alreadyInTerminalState,proto3,oneof"`
}

type EventFailureReason_IncompatibleCluster struct {
	IncompatibleCluster *EventErrorIncompatibleCluster `protobuf:"bytes,2,opt,name=incompatible_cluster,json=incompatibleCluster,proto3,oneof"`
}

func (*EventFailureReason_AlreadyInTerminalState) isEventFailureReason_Reason() {}

func (*EventFailureReason_IncompatibleCluster) isEventFailureReason_Reason() {}

func (m *EventFailureReason) GetReason() isEventFailureReason_Reason {
	if m != nil {
		return m.Reason
	}
	return nil
}

func (m *EventFailureReason) GetAlreadyInTerminalState() *EventErrorAlreadyInTerminalState {
	if x, ok := m.GetReason().(*EventFailureReason_AlreadyInTerminalState); ok {
		return x.AlreadyInTerminalState
	}
	return nil
}

func (m *EventFailureReason) GetIncompatibleCluster() *EventErrorIncompatibleCluster {
	if x, ok := m.GetReason().(*EventFailureReason_IncompatibleCluster); ok {
		return x.IncompatibleCluster
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EventFailureReason) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EventFailureReason_AlreadyInTerminalState)(nil),
		(*EventFailureReason_IncompatibleCluster)(nil),
	}
}

// Request to send a notification that a workflow execution event has occurred.
type WorkflowExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Details about the event that occurred.
	Event                *event.WorkflowExecutionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *WorkflowExecutionEventRequest) Reset()         { *m = WorkflowExecutionEventRequest{} }
func (m *WorkflowExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEventRequest) ProtoMessage()    {}
func (*WorkflowExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{3}
}

func (m *WorkflowExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Unmarshal(m, b)
}
func (m *WorkflowExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Marshal(b, m, deterministic)
}
func (m *WorkflowExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEventRequest.Merge(m, src)
}
func (m *WorkflowExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEventRequest.Size(m)
}
func (m *WorkflowExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEventRequest proto.InternalMessageInfo

func (m *WorkflowExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *WorkflowExecutionEventRequest) GetEvent() *event.WorkflowExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type WorkflowExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecutionEventResponse) Reset()         { *m = WorkflowExecutionEventResponse{} }
func (m *WorkflowExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEventResponse) ProtoMessage()    {}
func (*WorkflowExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{4}
}

func (m *WorkflowExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Unmarshal(m, b)
}
func (m *WorkflowExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Marshal(b, m, deterministic)
}
func (m *WorkflowExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEventResponse.Merge(m, src)
}
func (m *WorkflowExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEventResponse.Size(m)
}
func (m *WorkflowExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEventResponse proto.InternalMessageInfo

// Request to send a notification that a node execution event has occurred.
type NodeExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Details about the event that occurred.
	Event                *event.NodeExecutionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *NodeExecutionEventRequest) Reset()         { *m = NodeExecutionEventRequest{} }
func (m *NodeExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEventRequest) ProtoMessage()    {}
func (*NodeExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{5}
}

func (m *NodeExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEventRequest.Unmarshal(m, b)
}
func (m *NodeExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEventRequest.Marshal(b, m, deterministic)
}
func (m *NodeExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEventRequest.Merge(m, src)
}
func (m *NodeExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEventRequest.Size(m)
}
func (m *NodeExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEventRequest proto.InternalMessageInfo

func (m *NodeExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *NodeExecutionEventRequest) GetEvent() *event.NodeExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type NodeExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionEventResponse) Reset()         { *m = NodeExecutionEventResponse{} }
func (m *NodeExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEventResponse) ProtoMessage()    {}
func (*NodeExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{6}
}

func (m *NodeExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEventResponse.Unmarshal(m, b)
}
func (m *NodeExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEventResponse.Marshal(b, m, deterministic)
}
func (m *NodeExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEventResponse.Merge(m, src)
}
func (m *NodeExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEventResponse.Size(m)
}
func (m *NodeExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEventResponse proto.InternalMessageInfo

// Request to send a notification that a task execution event has occurred.
type TaskExecutionEventRequest struct {
	// Unique ID for this request that can be traced between services
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Details about the event that occurred.
	Event                *event.TaskExecutionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *TaskExecutionEventRequest) Reset()         { *m = TaskExecutionEventRequest{} }
func (m *TaskExecutionEventRequest) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEventRequest) ProtoMessage()    {}
func (*TaskExecutionEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{7}
}

func (m *TaskExecutionEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEventRequest.Unmarshal(m, b)
}
func (m *TaskExecutionEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEventRequest.Marshal(b, m, deterministic)
}
func (m *TaskExecutionEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEventRequest.Merge(m, src)
}
func (m *TaskExecutionEventRequest) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEventRequest.Size(m)
}
func (m *TaskExecutionEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEventRequest proto.InternalMessageInfo

func (m *TaskExecutionEventRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *TaskExecutionEventRequest) GetEvent() *event.TaskExecutionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

type TaskExecutionEventResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionEventResponse) Reset()         { *m = TaskExecutionEventResponse{} }
func (m *TaskExecutionEventResponse) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEventResponse) ProtoMessage()    {}
func (*TaskExecutionEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581752dea61f248, []int{8}
}

func (m *TaskExecutionEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEventResponse.Unmarshal(m, b)
}
func (m *TaskExecutionEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEventResponse.Marshal(b, m, deterministic)
}
func (m *TaskExecutionEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEventResponse.Merge(m, src)
}
func (m *TaskExecutionEventResponse) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEventResponse.Size(m)
}
func (m *TaskExecutionEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEventResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventErrorAlreadyInTerminalState)(nil), "flyteidl.admin.EventErrorAlreadyInTerminalState")
	proto.RegisterType((*EventErrorIncompatibleCluster)(nil), "flyteidl.admin.EventErrorIncompatibleCluster")
	proto.RegisterType((*EventFailureReason)(nil), "flyteidl.admin.EventFailureReason")
	proto.RegisterType((*WorkflowExecutionEventRequest)(nil), "flyteidl.admin.WorkflowExecutionEventRequest")
	proto.RegisterType((*WorkflowExecutionEventResponse)(nil), "flyteidl.admin.WorkflowExecutionEventResponse")
	proto.RegisterType((*NodeExecutionEventRequest)(nil), "flyteidl.admin.NodeExecutionEventRequest")
	proto.RegisterType((*NodeExecutionEventResponse)(nil), "flyteidl.admin.NodeExecutionEventResponse")
	proto.RegisterType((*TaskExecutionEventRequest)(nil), "flyteidl.admin.TaskExecutionEventRequest")
	proto.RegisterType((*TaskExecutionEventResponse)(nil), "flyteidl.admin.TaskExecutionEventResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/event.proto", fileDescriptor_4581752dea61f248) }

var fileDescriptor_4581752dea61f248 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xdf, 0x6b, 0xdb, 0x30,
	0x10, 0xc7, 0x93, 0xc1, 0xb2, 0xe5, 0xf6, 0xe3, 0xc1, 0x1b, 0x23, 0x09, 0xcb, 0x08, 0x1e, 0x8c,
	0xbd, 0xc4, 0x1e, 0xdb, 0xcb, 0x4a, 0xdb, 0x87, 0xa6, 0xa4, 0x4d, 0x5e, 0x4a, 0x71, 0x03, 0x85,
	0xbe, 0x18, 0xd9, 0xbe, 0x38, 0x22, 0xb2, 0xe4, 0x4a, 0x72, 0xdb, 0x40, 0xff, 0xe9, 0xfe, 0x07,
	0x25, 0xb2, 0x53, 0x27, 0xad, 0x43, 0xa1, 0xd0, 0xb7, 0xd3, 0xe9, 0xee, 0x7b, 0x9f, 0xaf, 0xc4,
	0x41, 0x67, 0xca, 0x16, 0x1a, 0x69, 0xc4, 0x5c, 0x12, 0x25, 0x94, 0xbb, 0x78, 0x85, 0x5c, 0x3b,
	0xa9, 0x14, 0x5a, 0x58, 0x9f, 0x57, 0x77, 0x8e, 0xb9, 0xeb, 0x94, 0xb5, 0xa6, 0x6a, 0xbd, 0xd6,
	0x3e, 0x86, 0xde, 0x70, 0x79, 0x1c, 0x4a, 0x29, 0xe4, 0x01, 0x93, 0x48, 0xa2, 0xc5, 0x98, 0x4f,
	0x50, 0x26, 0x94, 0x13, 0x76, 0xa6, 0x89, 0x46, 0xeb, 0x27, 0x7c, 0x0a, 0x33, 0x29, 0x91, 0x6b,
	0x3f, 0x9d, 0x11, 0x85, 0xad, 0x7a, 0xaf, 0xfe, 0xbb, 0xe9, 0x7d, 0x2c, 0x92, 0xa7, 0xcb, 0x9c,
	0xbd, 0x03, 0xdd, 0x52, 0x68, 0xcc, 0x43, 0x91, 0xa4, 0x44, 0xd3, 0x80, 0xe1, 0x21, 0xcb, 0x94,
	0x46, 0x69, 0xb5, 0xe0, 0x5d, 0x98, 0x87, 0x45, 0xff, 0xea, 0x68, 0xdf, 0xd5, 0xc1, 0x32, 0xbd,
	0x47, 0x84, 0xb2, 0x4c, 0xa2, 0x87, 0x44, 0x09, 0x6e, 0x25, 0xd0, 0x26, 0x39, 0x90, 0x4f, 0xb9,
	0xaf, 0x0b, 0x24, 0x5f, 0x2d, 0x99, 0x8c, 0xc4, 0x87, 0xbf, 0x7f, 0x9c, 0x4d, 0xab, 0xce, 0x73,
	0x5e, 0x46, 0x35, 0xef, 0x1b, 0xa9, 0x76, 0x19, 0xc0, 0x57, 0xba, 0x86, 0xed, 0xaf, 0x60, 0xdf,
	0x98, 0x49, 0xfd, 0xed, 0x93, 0x2a, 0xcc, 0x8e, 0x6a, 0xde, 0x17, 0xfa, 0x34, 0x3d, 0x78, 0x0f,
	0x0d, 0x69, 0xcc, 0xd9, 0xb7, 0xd0, 0x3d, 0x17, 0x72, 0x3e, 0x65, 0xe2, 0x7a, 0x78, 0x83, 0x61,
	0xa6, 0xa9, 0xe0, 0x46, 0xd2, 0xc3, 0xcb, 0x0c, 0x95, 0xb6, 0xba, 0x00, 0x32, 0x0f, 0x7d, 0x1a,
	0x15, 0x2f, 0xd6, 0x2c, 0x32, 0xe3, 0xc8, 0xda, 0x83, 0xb7, 0xe6, 0x1b, 0x0b, 0xbc, 0x5f, 0x25,
	0x5e, 0xfe, 0xbb, 0x5b, 0xc4, 0xf3, 0x26, 0xbb, 0x07, 0x3f, 0xb6, 0x4d, 0x57, 0xa9, 0xe0, 0x0a,
	0x6d, 0x0d, 0xed, 0x13, 0x11, 0xe1, 0x8b, 0xd8, 0xfe, 0x6f, 0xb2, 0xd9, 0x8f, 0xd9, 0x2a, 0x84,
	0x0b, 0xae, 0xef, 0xd0, 0xa9, 0x9a, 0x5a, 0x32, 0x4d, 0x88, 0x9a, 0xbf, 0x0a, 0x53, 0x85, 0x70,
	0xc9, 0x54, 0x35, 0x35, 0x67, 0x1a, 0xec, 0x5f, 0xec, 0xc6, 0x54, 0xcf, 0xb2, 0xc0, 0x09, 0x45,
	0xe2, 0x1a, 0x51, 0x21, 0xe3, 0x3c, 0x70, 0x1f, 0xf6, 0x2e, 0x46, 0xee, 0xa6, 0x41, 0x3f, 0x16,
	0xee, 0xe6, 0xda, 0x06, 0x0d, 0xb3, 0x85, 0xff, 0xee, 0x03, 0x00, 0x00, 0xff, 0xff, 0xea, 0x29,
	0x98, 0x7c, 0xcf, 0x03, 0x00, 0x00,
}