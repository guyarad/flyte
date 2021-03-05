// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/matchable_resource.proto

package admin

import (
	fmt "fmt"
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
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

// Defines a resource that can be configured by customizable Project-, ProjectDomain- or WorkflowAttributes
// based on matching tags.
type MatchableResource int32

const (
	// Applies to customizable task resource requests and limits.
	MatchableResource_TASK_RESOURCE MatchableResource = 0
	// Applies to configuring templated kubernetes cluster resources.
	MatchableResource_CLUSTER_RESOURCE MatchableResource = 1
	// Configures task and dynamic task execution queue assignment.
	MatchableResource_EXECUTION_QUEUE MatchableResource = 2
	// Configures the K8s cluster label to be used for execution to be run
	MatchableResource_EXECUTION_CLUSTER_LABEL MatchableResource = 3
	// Configures default quality of service when undefined in an execution spec.
	MatchableResource_QUALITY_OF_SERVICE_SPECIFICATION MatchableResource = 4
	// Selects configurable plugin implementation behavior for a given task type.
	MatchableResource_PLUGIN_OVERRIDE MatchableResource = 5
)

var MatchableResource_name = map[int32]string{
	0: "TASK_RESOURCE",
	1: "CLUSTER_RESOURCE",
	2: "EXECUTION_QUEUE",
	3: "EXECUTION_CLUSTER_LABEL",
	4: "QUALITY_OF_SERVICE_SPECIFICATION",
	5: "PLUGIN_OVERRIDE",
}

var MatchableResource_value = map[string]int32{
	"TASK_RESOURCE":                    0,
	"CLUSTER_RESOURCE":                 1,
	"EXECUTION_QUEUE":                  2,
	"EXECUTION_CLUSTER_LABEL":          3,
	"QUALITY_OF_SERVICE_SPECIFICATION": 4,
	"PLUGIN_OVERRIDE":                  5,
}

func (x MatchableResource) String() string {
	return proto.EnumName(MatchableResource_name, int32(x))
}

func (MatchableResource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{0}
}

type PluginOverride_MissingPluginBehavior int32

const (
	PluginOverride_FAIL PluginOverride_MissingPluginBehavior = 0
	// Uses the system-configured default implementation.
	PluginOverride_USE_DEFAULT PluginOverride_MissingPluginBehavior = 1
)

var PluginOverride_MissingPluginBehavior_name = map[int32]string{
	0: "FAIL",
	1: "USE_DEFAULT",
}

var PluginOverride_MissingPluginBehavior_value = map[string]int32{
	"FAIL":        0,
	"USE_DEFAULT": 1,
}

func (x PluginOverride_MissingPluginBehavior) String() string {
	return proto.EnumName(PluginOverride_MissingPluginBehavior_name, int32(x))
}

func (PluginOverride_MissingPluginBehavior) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{5, 0}
}

type TaskResourceSpec struct {
	Cpu                  string   `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gpu                  string   `protobuf:"bytes,2,opt,name=gpu,proto3" json:"gpu,omitempty"`
	Memory               string   `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage              string   `protobuf:"bytes,4,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResourceSpec) Reset()         { *m = TaskResourceSpec{} }
func (m *TaskResourceSpec) String() string { return proto.CompactTextString(m) }
func (*TaskResourceSpec) ProtoMessage()    {}
func (*TaskResourceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{0}
}

func (m *TaskResourceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResourceSpec.Unmarshal(m, b)
}
func (m *TaskResourceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResourceSpec.Marshal(b, m, deterministic)
}
func (m *TaskResourceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResourceSpec.Merge(m, src)
}
func (m *TaskResourceSpec) XXX_Size() int {
	return xxx_messageInfo_TaskResourceSpec.Size(m)
}
func (m *TaskResourceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResourceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResourceSpec proto.InternalMessageInfo

func (m *TaskResourceSpec) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func (m *TaskResourceSpec) GetGpu() string {
	if m != nil {
		return m.Gpu
	}
	return ""
}

func (m *TaskResourceSpec) GetMemory() string {
	if m != nil {
		return m.Memory
	}
	return ""
}

func (m *TaskResourceSpec) GetStorage() string {
	if m != nil {
		return m.Storage
	}
	return ""
}

type TaskResourceAttributes struct {
	Defaults             *TaskResourceSpec `protobuf:"bytes,1,opt,name=defaults,proto3" json:"defaults,omitempty"`
	Limits               *TaskResourceSpec `protobuf:"bytes,2,opt,name=limits,proto3" json:"limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskResourceAttributes) Reset()         { *m = TaskResourceAttributes{} }
func (m *TaskResourceAttributes) String() string { return proto.CompactTextString(m) }
func (*TaskResourceAttributes) ProtoMessage()    {}
func (*TaskResourceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{1}
}

func (m *TaskResourceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResourceAttributes.Unmarshal(m, b)
}
func (m *TaskResourceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResourceAttributes.Marshal(b, m, deterministic)
}
func (m *TaskResourceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResourceAttributes.Merge(m, src)
}
func (m *TaskResourceAttributes) XXX_Size() int {
	return xxx_messageInfo_TaskResourceAttributes.Size(m)
}
func (m *TaskResourceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResourceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResourceAttributes proto.InternalMessageInfo

func (m *TaskResourceAttributes) GetDefaults() *TaskResourceSpec {
	if m != nil {
		return m.Defaults
	}
	return nil
}

func (m *TaskResourceAttributes) GetLimits() *TaskResourceSpec {
	if m != nil {
		return m.Limits
	}
	return nil
}

type ClusterResourceAttributes struct {
	// Custom resource attributes which will be applied in cluster resource creation (e.g. quotas).
	// Map keys are the *case-sensitive* names of variables in templatized resource files.
	// Map values should be the custom values which get substituted during resource creation.
	Attributes           map[string]string `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClusterResourceAttributes) Reset()         { *m = ClusterResourceAttributes{} }
func (m *ClusterResourceAttributes) String() string { return proto.CompactTextString(m) }
func (*ClusterResourceAttributes) ProtoMessage()    {}
func (*ClusterResourceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{2}
}

func (m *ClusterResourceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterResourceAttributes.Unmarshal(m, b)
}
func (m *ClusterResourceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterResourceAttributes.Marshal(b, m, deterministic)
}
func (m *ClusterResourceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterResourceAttributes.Merge(m, src)
}
func (m *ClusterResourceAttributes) XXX_Size() int {
	return xxx_messageInfo_ClusterResourceAttributes.Size(m)
}
func (m *ClusterResourceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterResourceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterResourceAttributes proto.InternalMessageInfo

func (m *ClusterResourceAttributes) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type ExecutionQueueAttributes struct {
	// Tags used for assigning execution queues for tasks defined within this project.
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionQueueAttributes) Reset()         { *m = ExecutionQueueAttributes{} }
func (m *ExecutionQueueAttributes) String() string { return proto.CompactTextString(m) }
func (*ExecutionQueueAttributes) ProtoMessage()    {}
func (*ExecutionQueueAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{3}
}

func (m *ExecutionQueueAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionQueueAttributes.Unmarshal(m, b)
}
func (m *ExecutionQueueAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionQueueAttributes.Marshal(b, m, deterministic)
}
func (m *ExecutionQueueAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionQueueAttributes.Merge(m, src)
}
func (m *ExecutionQueueAttributes) XXX_Size() int {
	return xxx_messageInfo_ExecutionQueueAttributes.Size(m)
}
func (m *ExecutionQueueAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionQueueAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionQueueAttributes proto.InternalMessageInfo

func (m *ExecutionQueueAttributes) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ExecutionClusterLabel struct {
	// Label value to determine where the execution will be run
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionClusterLabel) Reset()         { *m = ExecutionClusterLabel{} }
func (m *ExecutionClusterLabel) String() string { return proto.CompactTextString(m) }
func (*ExecutionClusterLabel) ProtoMessage()    {}
func (*ExecutionClusterLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{4}
}

func (m *ExecutionClusterLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionClusterLabel.Unmarshal(m, b)
}
func (m *ExecutionClusterLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionClusterLabel.Marshal(b, m, deterministic)
}
func (m *ExecutionClusterLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionClusterLabel.Merge(m, src)
}
func (m *ExecutionClusterLabel) XXX_Size() int {
	return xxx_messageInfo_ExecutionClusterLabel.Size(m)
}
func (m *ExecutionClusterLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionClusterLabel.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionClusterLabel proto.InternalMessageInfo

func (m *ExecutionClusterLabel) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// This MatchableAttribute configures selecting alternate plugin implementations for a given task type.
// In addition to an override implementation a selection of fallbacks can be provided or other modes
// for handling cases where the desired plugin override is not enabled in a given Flyte deployment.
type PluginOverride struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// A set of plugin ids which should handle tasks of this type instead of the default registered plugin. The list will be tried in order until a plugin is found with that id.
	PluginId []string `protobuf:"bytes,2,rep,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	// Defines the behavior when no plugin from the plugin_id list is not found.
	MissingPluginBehavior PluginOverride_MissingPluginBehavior `protobuf:"varint,4,opt,name=missing_plugin_behavior,json=missingPluginBehavior,proto3,enum=flyteidl.admin.PluginOverride_MissingPluginBehavior" json:"missing_plugin_behavior,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                             `json:"-"`
	XXX_unrecognized      []byte                               `json:"-"`
	XXX_sizecache         int32                                `json:"-"`
}

func (m *PluginOverride) Reset()         { *m = PluginOverride{} }
func (m *PluginOverride) String() string { return proto.CompactTextString(m) }
func (*PluginOverride) ProtoMessage()    {}
func (*PluginOverride) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{5}
}

func (m *PluginOverride) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginOverride.Unmarshal(m, b)
}
func (m *PluginOverride) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginOverride.Marshal(b, m, deterministic)
}
func (m *PluginOverride) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginOverride.Merge(m, src)
}
func (m *PluginOverride) XXX_Size() int {
	return xxx_messageInfo_PluginOverride.Size(m)
}
func (m *PluginOverride) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginOverride.DiscardUnknown(m)
}

var xxx_messageInfo_PluginOverride proto.InternalMessageInfo

func (m *PluginOverride) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *PluginOverride) GetPluginId() []string {
	if m != nil {
		return m.PluginId
	}
	return nil
}

func (m *PluginOverride) GetMissingPluginBehavior() PluginOverride_MissingPluginBehavior {
	if m != nil {
		return m.MissingPluginBehavior
	}
	return PluginOverride_FAIL
}

type PluginOverrides struct {
	Overrides            []*PluginOverride `protobuf:"bytes,1,rep,name=overrides,proto3" json:"overrides,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PluginOverrides) Reset()         { *m = PluginOverrides{} }
func (m *PluginOverrides) String() string { return proto.CompactTextString(m) }
func (*PluginOverrides) ProtoMessage()    {}
func (*PluginOverrides) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{6}
}

func (m *PluginOverrides) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginOverrides.Unmarshal(m, b)
}
func (m *PluginOverrides) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginOverrides.Marshal(b, m, deterministic)
}
func (m *PluginOverrides) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginOverrides.Merge(m, src)
}
func (m *PluginOverrides) XXX_Size() int {
	return xxx_messageInfo_PluginOverrides.Size(m)
}
func (m *PluginOverrides) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginOverrides.DiscardUnknown(m)
}

var xxx_messageInfo_PluginOverrides proto.InternalMessageInfo

func (m *PluginOverrides) GetOverrides() []*PluginOverride {
	if m != nil {
		return m.Overrides
	}
	return nil
}

// Generic container for encapsulating all types of the above attributes messages.
type MatchingAttributes struct {
	// Types that are valid to be assigned to Target:
	//	*MatchingAttributes_TaskResourceAttributes
	//	*MatchingAttributes_ClusterResourceAttributes
	//	*MatchingAttributes_ExecutionQueueAttributes
	//	*MatchingAttributes_ExecutionClusterLabel
	//	*MatchingAttributes_QualityOfService
	//	*MatchingAttributes_PluginOverrides
	Target               isMatchingAttributes_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *MatchingAttributes) Reset()         { *m = MatchingAttributes{} }
func (m *MatchingAttributes) String() string { return proto.CompactTextString(m) }
func (*MatchingAttributes) ProtoMessage()    {}
func (*MatchingAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{7}
}

func (m *MatchingAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchingAttributes.Unmarshal(m, b)
}
func (m *MatchingAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchingAttributes.Marshal(b, m, deterministic)
}
func (m *MatchingAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchingAttributes.Merge(m, src)
}
func (m *MatchingAttributes) XXX_Size() int {
	return xxx_messageInfo_MatchingAttributes.Size(m)
}
func (m *MatchingAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchingAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_MatchingAttributes proto.InternalMessageInfo

type isMatchingAttributes_Target interface {
	isMatchingAttributes_Target()
}

type MatchingAttributes_TaskResourceAttributes struct {
	TaskResourceAttributes *TaskResourceAttributes `protobuf:"bytes,1,opt,name=task_resource_attributes,json=taskResourceAttributes,proto3,oneof"`
}

type MatchingAttributes_ClusterResourceAttributes struct {
	ClusterResourceAttributes *ClusterResourceAttributes `protobuf:"bytes,2,opt,name=cluster_resource_attributes,json=clusterResourceAttributes,proto3,oneof"`
}

type MatchingAttributes_ExecutionQueueAttributes struct {
	ExecutionQueueAttributes *ExecutionQueueAttributes `protobuf:"bytes,3,opt,name=execution_queue_attributes,json=executionQueueAttributes,proto3,oneof"`
}

type MatchingAttributes_ExecutionClusterLabel struct {
	ExecutionClusterLabel *ExecutionClusterLabel `protobuf:"bytes,4,opt,name=execution_cluster_label,json=executionClusterLabel,proto3,oneof"`
}

type MatchingAttributes_QualityOfService struct {
	QualityOfService *core.QualityOfService `protobuf:"bytes,5,opt,name=quality_of_service,json=qualityOfService,proto3,oneof"`
}

type MatchingAttributes_PluginOverrides struct {
	PluginOverrides *PluginOverrides `protobuf:"bytes,6,opt,name=plugin_overrides,json=pluginOverrides,proto3,oneof"`
}

func (*MatchingAttributes_TaskResourceAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ClusterResourceAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ExecutionQueueAttributes) isMatchingAttributes_Target() {}

func (*MatchingAttributes_ExecutionClusterLabel) isMatchingAttributes_Target() {}

func (*MatchingAttributes_QualityOfService) isMatchingAttributes_Target() {}

func (*MatchingAttributes_PluginOverrides) isMatchingAttributes_Target() {}

func (m *MatchingAttributes) GetTarget() isMatchingAttributes_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *MatchingAttributes) GetTaskResourceAttributes() *TaskResourceAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_TaskResourceAttributes); ok {
		return x.TaskResourceAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetClusterResourceAttributes() *ClusterResourceAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_ClusterResourceAttributes); ok {
		return x.ClusterResourceAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetExecutionQueueAttributes() *ExecutionQueueAttributes {
	if x, ok := m.GetTarget().(*MatchingAttributes_ExecutionQueueAttributes); ok {
		return x.ExecutionQueueAttributes
	}
	return nil
}

func (m *MatchingAttributes) GetExecutionClusterLabel() *ExecutionClusterLabel {
	if x, ok := m.GetTarget().(*MatchingAttributes_ExecutionClusterLabel); ok {
		return x.ExecutionClusterLabel
	}
	return nil
}

func (m *MatchingAttributes) GetQualityOfService() *core.QualityOfService {
	if x, ok := m.GetTarget().(*MatchingAttributes_QualityOfService); ok {
		return x.QualityOfService
	}
	return nil
}

func (m *MatchingAttributes) GetPluginOverrides() *PluginOverrides {
	if x, ok := m.GetTarget().(*MatchingAttributes_PluginOverrides); ok {
		return x.PluginOverrides
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MatchingAttributes) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MatchingAttributes_TaskResourceAttributes)(nil),
		(*MatchingAttributes_ClusterResourceAttributes)(nil),
		(*MatchingAttributes_ExecutionQueueAttributes)(nil),
		(*MatchingAttributes_ExecutionClusterLabel)(nil),
		(*MatchingAttributes_QualityOfService)(nil),
		(*MatchingAttributes_PluginOverrides)(nil),
	}
}

// Represents a custom set of attributes applied for either a domain; a domain and project; or
// domain, project and workflow name.
type MatchableAttributesConfiguration struct {
	Attributes           *MatchingAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Domain               string              `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Project              string              `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	Workflow             string              `protobuf:"bytes,4,opt,name=workflow,proto3" json:"workflow,omitempty"`
	LaunchPlan           string              `protobuf:"bytes,5,opt,name=launch_plan,json=launchPlan,proto3" json:"launch_plan,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MatchableAttributesConfiguration) Reset()         { *m = MatchableAttributesConfiguration{} }
func (m *MatchableAttributesConfiguration) String() string { return proto.CompactTextString(m) }
func (*MatchableAttributesConfiguration) ProtoMessage()    {}
func (*MatchableAttributesConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{8}
}

func (m *MatchableAttributesConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchableAttributesConfiguration.Unmarshal(m, b)
}
func (m *MatchableAttributesConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchableAttributesConfiguration.Marshal(b, m, deterministic)
}
func (m *MatchableAttributesConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchableAttributesConfiguration.Merge(m, src)
}
func (m *MatchableAttributesConfiguration) XXX_Size() int {
	return xxx_messageInfo_MatchableAttributesConfiguration.Size(m)
}
func (m *MatchableAttributesConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchableAttributesConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_MatchableAttributesConfiguration proto.InternalMessageInfo

func (m *MatchableAttributesConfiguration) GetAttributes() *MatchingAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *MatchableAttributesConfiguration) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func (m *MatchableAttributesConfiguration) GetLaunchPlan() string {
	if m != nil {
		return m.LaunchPlan
	}
	return ""
}

// Request all matching resource attributes.
type ListMatchableAttributesRequest struct {
	ResourceType         MatchableResource `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListMatchableAttributesRequest) Reset()         { *m = ListMatchableAttributesRequest{} }
func (m *ListMatchableAttributesRequest) String() string { return proto.CompactTextString(m) }
func (*ListMatchableAttributesRequest) ProtoMessage()    {}
func (*ListMatchableAttributesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{9}
}

func (m *ListMatchableAttributesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatchableAttributesRequest.Unmarshal(m, b)
}
func (m *ListMatchableAttributesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatchableAttributesRequest.Marshal(b, m, deterministic)
}
func (m *ListMatchableAttributesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatchableAttributesRequest.Merge(m, src)
}
func (m *ListMatchableAttributesRequest) XXX_Size() int {
	return xxx_messageInfo_ListMatchableAttributesRequest.Size(m)
}
func (m *ListMatchableAttributesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatchableAttributesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatchableAttributesRequest proto.InternalMessageInfo

func (m *ListMatchableAttributesRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Response for a request for all matching resource attributes.
type ListMatchableAttributesResponse struct {
	Configurations       []*MatchableAttributesConfiguration `protobuf:"bytes,1,rep,name=configurations,proto3" json:"configurations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ListMatchableAttributesResponse) Reset()         { *m = ListMatchableAttributesResponse{} }
func (m *ListMatchableAttributesResponse) String() string { return proto.CompactTextString(m) }
func (*ListMatchableAttributesResponse) ProtoMessage()    {}
func (*ListMatchableAttributesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d15bcabb02640f4, []int{10}
}

func (m *ListMatchableAttributesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatchableAttributesResponse.Unmarshal(m, b)
}
func (m *ListMatchableAttributesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatchableAttributesResponse.Marshal(b, m, deterministic)
}
func (m *ListMatchableAttributesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatchableAttributesResponse.Merge(m, src)
}
func (m *ListMatchableAttributesResponse) XXX_Size() int {
	return xxx_messageInfo_ListMatchableAttributesResponse.Size(m)
}
func (m *ListMatchableAttributesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatchableAttributesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatchableAttributesResponse proto.InternalMessageInfo

func (m *ListMatchableAttributesResponse) GetConfigurations() []*MatchableAttributesConfiguration {
	if m != nil {
		return m.Configurations
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.admin.MatchableResource", MatchableResource_name, MatchableResource_value)
	proto.RegisterEnum("flyteidl.admin.PluginOverride_MissingPluginBehavior", PluginOverride_MissingPluginBehavior_name, PluginOverride_MissingPluginBehavior_value)
	proto.RegisterType((*TaskResourceSpec)(nil), "flyteidl.admin.TaskResourceSpec")
	proto.RegisterType((*TaskResourceAttributes)(nil), "flyteidl.admin.TaskResourceAttributes")
	proto.RegisterType((*ClusterResourceAttributes)(nil), "flyteidl.admin.ClusterResourceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.ClusterResourceAttributes.AttributesEntry")
	proto.RegisterType((*ExecutionQueueAttributes)(nil), "flyteidl.admin.ExecutionQueueAttributes")
	proto.RegisterType((*ExecutionClusterLabel)(nil), "flyteidl.admin.ExecutionClusterLabel")
	proto.RegisterType((*PluginOverride)(nil), "flyteidl.admin.PluginOverride")
	proto.RegisterType((*PluginOverrides)(nil), "flyteidl.admin.PluginOverrides")
	proto.RegisterType((*MatchingAttributes)(nil), "flyteidl.admin.MatchingAttributes")
	proto.RegisterType((*MatchableAttributesConfiguration)(nil), "flyteidl.admin.MatchableAttributesConfiguration")
	proto.RegisterType((*ListMatchableAttributesRequest)(nil), "flyteidl.admin.ListMatchableAttributesRequest")
	proto.RegisterType((*ListMatchableAttributesResponse)(nil), "flyteidl.admin.ListMatchableAttributesResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/matchable_resource.proto", fileDescriptor_1d15bcabb02640f4)
}

var fileDescriptor_1d15bcabb02640f4 = []byte{
	// 978 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x72, 0xda, 0x46,
	0x14, 0x46, 0x06, 0x53, 0x38, 0x34, 0xa0, 0x6c, 0x63, 0x5b, 0xc1, 0xd3, 0x98, 0x6a, 0xfa, 0xe3,
	0x76, 0x26, 0xd0, 0xa1, 0xed, 0x34, 0xed, 0xb4, 0x17, 0x40, 0x44, 0x61, 0x2a, 0x07, 0x7b, 0x81,
	0x4c, 0xd2, 0x1b, 0x8d, 0x10, 0x8b, 0x50, 0x91, 0xb4, 0xb2, 0xb4, 0x72, 0xca, 0xf4, 0x25, 0xfa,
	0x1a, 0x7d, 0x80, 0x3e, 0x4a, 0x2f, 0xfa, 0x0e, 0x7d, 0x88, 0x8c, 0xfe, 0x10, 0x28, 0x90, 0xf1,
	0x9d, 0xce, 0xd9, 0x73, 0xce, 0x77, 0xf6, 0xec, 0xb7, 0x9f, 0x16, 0xbe, 0x58, 0x98, 0x6b, 0x46,
	0x8c, 0xb9, 0xd9, 0x52, 0xe7, 0x96, 0x61, 0xb7, 0x2c, 0x95, 0x69, 0x4b, 0x75, 0x66, 0x12, 0xc5,
	0x25, 0x1e, 0xf5, 0x5d, 0x8d, 0x34, 0x1d, 0x97, 0x32, 0x8a, 0xaa, 0x49, 0x60, 0x33, 0x0c, 0xac,
	0x9f, 0x67, 0x12, 0x35, 0x6a, 0x59, 0xd4, 0x8e, 0x82, 0xeb, 0x1f, 0x6f, 0x16, 0x35, 0xea, 0x92,
	0x16, 0xf9, 0x83, 0x68, 0x3e, 0x33, 0x92, 0x65, 0x71, 0x09, 0xfc, 0x44, 0xf5, 0x56, 0x38, 0x46,
	0x18, 0x3b, 0x44, 0x43, 0x3c, 0xe4, 0x35, 0xc7, 0x17, 0xb8, 0x06, 0x77, 0x59, 0xc6, 0xc1, 0x67,
	0xe0, 0xd1, 0x1d, 0x5f, 0x38, 0x8a, 0x3c, 0xba, 0xe3, 0xa3, 0x53, 0x28, 0x5a, 0xc4, 0xa2, 0xee,
	0x5a, 0xc8, 0x87, 0xce, 0xd8, 0x42, 0x02, 0x7c, 0xe0, 0x31, 0xea, 0xaa, 0x3a, 0x11, 0x0a, 0xe1,
	0x42, 0x62, 0x8a, 0x7f, 0x71, 0x70, 0xba, 0x0d, 0xd5, 0x61, 0xcc, 0x35, 0x66, 0x3e, 0x23, 0x1e,
	0xfa, 0x09, 0x4a, 0x73, 0xb2, 0x50, 0x7d, 0x93, 0x79, 0x21, 0x6a, 0xa5, 0xdd, 0x68, 0xee, 0xee,
	0xb1, 0x99, 0x6d, 0x12, 0x6f, 0x32, 0xd0, 0x33, 0x28, 0x9a, 0x86, 0x65, 0x30, 0x2f, 0xec, 0xef,
	0x3e, 0xb9, 0x71, 0xbc, 0xf8, 0x0f, 0x07, 0x8f, 0x7b, 0xa6, 0xef, 0x31, 0xe2, 0xee, 0xe9, 0xea,
	0x35, 0x80, 0xba, 0xb1, 0x04, 0xae, 0x91, 0xbf, 0xac, 0xb4, 0x7f, 0xc8, 0xd6, 0x3e, 0x98, 0xde,
	0x4c, 0x3f, 0x25, 0x9b, 0xb9, 0x6b, 0xbc, 0x55, 0xac, 0xfe, 0x33, 0xd4, 0x32, 0xcb, 0xc1, 0x88,
	0x57, 0x64, 0x9d, 0x0c, 0x7d, 0x45, 0xd6, 0xe8, 0x11, 0x1c, 0xdf, 0xa9, 0xa6, 0x4f, 0xe2, 0xb1,
	0x47, 0xc6, 0x8f, 0x47, 0xcf, 0x38, 0xb1, 0x09, 0x82, 0x94, 0x9c, 0xe3, 0x8d, 0x4f, 0xfc, 0xed,
	0xae, 0x11, 0x14, 0x98, 0xaa, 0x47, 0xfd, 0x96, 0x71, 0xf8, 0x2d, 0x3e, 0x85, 0x93, 0x4d, 0x7c,
	0xdc, 0xb0, 0xac, 0xce, 0x88, 0x99, 0x42, 0x70, 0x5b, 0x10, 0xe2, 0xff, 0x1c, 0x54, 0xaf, 0x4d,
	0x5f, 0x37, 0xec, 0xd1, 0x1d, 0x71, 0x5d, 0x63, 0x4e, 0xd0, 0x39, 0x94, 0x99, 0xea, 0xad, 0x14,
	0xb6, 0x76, 0x92, 0xe0, 0x52, 0xe0, 0x98, 0xac, 0x9d, 0x70, 0xd1, 0x09, 0xc3, 0x15, 0x63, 0x2e,
	0x1c, 0x85, 0xb8, 0xa5, 0xc8, 0x31, 0x9c, 0x23, 0x13, 0xce, 0x2c, 0xc3, 0xf3, 0x0c, 0x5b, 0x57,
	0xe2, 0xa0, 0x19, 0x59, 0xaa, 0x77, 0x06, 0x75, 0x43, 0x82, 0x54, 0xdb, 0xdf, 0x66, 0x47, 0xba,
	0x0b, 0xdd, 0xbc, 0x8a, 0xb2, 0x23, 0x6f, 0x37, 0xce, 0xc5, 0x27, 0xd6, 0x3e, 0xb7, 0xd8, 0x86,
	0x93, 0xbd, 0xf1, 0xa8, 0x04, 0x85, 0x7e, 0x67, 0x28, 0xf3, 0x39, 0x54, 0x83, 0xca, 0x74, 0x2c,
	0x29, 0xcf, 0xa5, 0x7e, 0x67, 0x2a, 0x4f, 0x78, 0x4e, 0x1c, 0x41, 0x6d, 0x17, 0x32, 0x20, 0x64,
	0x99, 0x26, 0x46, 0x7c, 0xf2, 0x4f, 0xde, 0xdf, 0x26, 0x4e, 0x13, 0xc4, 0xff, 0x0a, 0x80, 0xae,
	0x82, 0xcb, 0x6b, 0xd8, 0xfa, 0xd6, 0xc9, 0xcc, 0x40, 0x08, 0x67, 0x98, 0xdc, 0x66, 0x65, 0x87,
	0x5d, 0x01, 0x73, 0x3f, 0x7f, 0x1f, 0x73, 0xd3, 0x4a, 0x83, 0x1c, 0x3e, 0x65, 0xfb, 0x6f, 0xd2,
	0x0a, 0xce, 0xb5, 0xe8, 0x80, 0xf7, 0xc2, 0x44, 0x17, 0xe4, 0xcb, 0x7b, 0x93, 0x78, 0x90, 0xc3,
	0x8f, 0xb5, 0x83, 0x17, 0x64, 0x09, 0xf5, 0x8d, 0x9c, 0x28, 0xb7, 0x01, 0x0f, 0xb7, 0xb1, 0xf2,
	0x21, 0xd6, 0x65, 0x16, 0xeb, 0x10, 0x71, 0x07, 0x39, 0x2c, 0x90, 0x43, 0xa4, 0x56, 0xe0, 0x2c,
	0x45, 0x4a, 0x36, 0x68, 0x06, 0x14, 0x0e, 0x49, 0x54, 0x69, 0x7f, 0x76, 0x10, 0x66, 0x9b, 0xef,
	0x83, 0x1c, 0x3e, 0x21, 0x7b, 0x2f, 0xc2, 0x08, 0xd0, 0xad, 0xaf, 0x9a, 0x06, 0x5b, 0x2b, 0x74,
	0xa1, 0x78, 0xc4, 0xbd, 0x33, 0x34, 0x22, 0x1c, 0x87, 0xb5, 0x2f, 0xd2, 0xda, 0x81, 0x84, 0x36,
	0x6f, 0xa2, 0xc0, 0xd1, 0x62, 0x1c, 0x85, 0x0d, 0x72, 0x98, 0xbf, 0xcd, 0xf8, 0x90, 0x0c, 0x7c,
	0x4c, 0xf7, 0x94, 0x48, 0xc5, 0x6c, 0xb9, 0x7d, 0x44, 0x0a, 0x06, 0x51, 0x73, 0x76, 0x5d, 0xdd,
	0x12, 0x14, 0x99, 0xea, 0xea, 0x84, 0x89, 0xff, 0x72, 0xd0, 0xb8, 0x4a, 0x7e, 0x0c, 0xe9, 0x84,
	0x7a, 0xd4, 0x5e, 0x18, 0xba, 0xef, 0xaa, 0xc1, 0xce, 0x50, 0x37, 0xa3, 0x5c, 0x01, 0xac, 0x98,
	0x85, 0x7d, 0x97, 0xa1, 0xdb, 0x12, 0x15, 0x08, 0xfc, 0x9c, 0x5a, 0xaa, 0x61, 0xc7, 0xf2, 0x13,
	0x5b, 0x81, 0xc0, 0x3b, 0x2e, 0xfd, 0x9d, 0x68, 0x2c, 0x56, 0xfe, 0xc4, 0x44, 0x75, 0x28, 0xbd,
	0xa1, 0xee, 0x6a, 0x61, 0xd2, 0x37, 0xb1, 0xf6, 0x6f, 0x6c, 0x74, 0x01, 0x15, 0x53, 0xf5, 0x6d,
	0x6d, 0xa9, 0x38, 0xa6, 0x6a, 0x87, 0x83, 0x2d, 0x63, 0x88, 0x5c, 0xd7, 0xa6, 0x6a, 0x8b, 0x4b,
	0x78, 0x22, 0x1b, 0x1e, 0xdb, 0xb3, 0x35, 0x4c, 0x6e, 0x7d, 0xe2, 0x31, 0xd4, 0x87, 0x07, 0x1b,
	0x4a, 0x6f, 0x64, 0xa8, 0xda, 0xfe, 0x64, 0xef, 0xbe, 0x82, 0x12, 0x09, 0x63, 0xf1, 0x87, 0x49,
	0x5e, 0xa0, 0x56, 0xe2, 0x9f, 0x70, 0x71, 0x10, 0xc9, 0x73, 0xa8, 0xed, 0x11, 0xf4, 0x0a, 0xaa,
	0xda, 0xf6, 0x40, 0x13, 0x0d, 0xf8, 0xfa, 0x20, 0xd6, 0x81, 0x93, 0xc0, 0x99, 0x3a, 0x5f, 0xfd,
	0xcd, 0xc1, 0xc3, 0x77, 0x1a, 0x44, 0x0f, 0xe1, 0xc1, 0xa4, 0x33, 0xfe, 0x55, 0xc1, 0xd2, 0x78,
	0x34, 0xc5, 0x3d, 0x89, 0xcf, 0xa1, 0x47, 0xc0, 0xf7, 0xe4, 0xe9, 0x78, 0x22, 0xe1, 0xd4, 0xcb,
	0xa1, 0x8f, 0xa0, 0x26, 0xbd, 0x92, 0x7a, 0xd3, 0xc9, 0x70, 0xf4, 0x42, 0xb9, 0x99, 0x4a, 0x53,
	0x89, 0x3f, 0x42, 0xe7, 0x70, 0x96, 0x3a, 0x93, 0x24, 0xb9, 0xd3, 0x95, 0x64, 0x3e, 0x8f, 0x3e,
	0x85, 0xc6, 0xcd, 0xb4, 0x23, 0x0f, 0x27, 0xaf, 0x95, 0x51, 0x5f, 0x19, 0x4b, 0xf8, 0xe5, 0xb0,
	0x27, 0x29, 0xe3, 0x6b, 0xa9, 0x37, 0xec, 0x0f, 0x7b, 0x9d, 0x20, 0x87, 0x2f, 0x04, 0x75, 0xaf,
	0xe5, 0xe9, 0x2f, 0xc3, 0x17, 0xca, 0xe8, 0xa5, 0x84, 0xf1, 0xf0, 0xb9, 0xc4, 0x1f, 0x77, 0xbf,
	0xff, 0xed, 0x3b, 0xdd, 0x60, 0x4b, 0x7f, 0xd6, 0xd4, 0xa8, 0xd5, 0x0a, 0x77, 0x4e, 0x5d, 0xbd,
	0xb5, 0x79, 0x4f, 0xe8, 0xc4, 0x6e, 0x39, 0xb3, 0xa7, 0x3a, 0x6d, 0xed, 0xbe, 0x3f, 0x66, 0xc5,
	0xf0, 0x69, 0xf1, 0xcd, 0xdb, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x33, 0x3c, 0xa1, 0xd1, 0x08,
	0x00, 0x00,
}
