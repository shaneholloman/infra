// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: orchestrator.proto

package orchestrator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SandboxConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data required for creating a new sandbox.
	TemplateId         string            `protobuf:"bytes,1,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	BuildId            string            `protobuf:"bytes,2,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	KernelVersion      string            `protobuf:"bytes,3,opt,name=kernel_version,json=kernelVersion,proto3" json:"kernel_version,omitempty"`
	FirecrackerVersion string            `protobuf:"bytes,4,opt,name=firecracker_version,json=firecrackerVersion,proto3" json:"firecracker_version,omitempty"`
	HugePages          bool              `protobuf:"varint,5,opt,name=huge_pages,json=hugePages,proto3" json:"huge_pages,omitempty"`
	SandboxId          string            `protobuf:"bytes,6,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
	EnvVars            map[string]string `protobuf:"bytes,7,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata about the sandbox.
	Metadata        map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Alias           *string           `protobuf:"bytes,9,opt,name=alias,proto3,oneof" json:"alias,omitempty"`
	EnvdVersion     string            `protobuf:"bytes,10,opt,name=envd_version,json=envdVersion,proto3" json:"envd_version,omitempty"`
	Vcpu            int64             `protobuf:"varint,11,opt,name=vcpu,proto3" json:"vcpu,omitempty"`
	RamMb           int64             `protobuf:"varint,12,opt,name=ram_mb,json=ramMb,proto3" json:"ram_mb,omitempty"`
	TotalDiskSizeMb int64             `protobuf:"varint,13,opt,name=total_disk_size_mb,json=totalDiskSizeMb,proto3" json:"total_disk_size_mb,omitempty"`
	TeamId          string            `protobuf:"bytes,14,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// Maximum length of the sandbox in Hours.
	MaxSandboxLength int64 `protobuf:"varint,15,opt,name=max_sandbox_length,json=maxSandboxLength,proto3" json:"max_sandbox_length,omitempty"`
}

func (x *SandboxConfig) Reset() {
	*x = SandboxConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxConfig) ProtoMessage() {}

func (x *SandboxConfig) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxConfig.ProtoReflect.Descriptor instead.
func (*SandboxConfig) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{0}
}

func (x *SandboxConfig) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SandboxConfig) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

func (x *SandboxConfig) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *SandboxConfig) GetFirecrackerVersion() string {
	if x != nil {
		return x.FirecrackerVersion
	}
	return ""
}

func (x *SandboxConfig) GetHugePages() bool {
	if x != nil {
		return x.HugePages
	}
	return false
}

func (x *SandboxConfig) GetSandboxId() string {
	if x != nil {
		return x.SandboxId
	}
	return ""
}

func (x *SandboxConfig) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

func (x *SandboxConfig) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SandboxConfig) GetAlias() string {
	if x != nil && x.Alias != nil {
		return *x.Alias
	}
	return ""
}

func (x *SandboxConfig) GetEnvdVersion() string {
	if x != nil {
		return x.EnvdVersion
	}
	return ""
}

func (x *SandboxConfig) GetVcpu() int64 {
	if x != nil {
		return x.Vcpu
	}
	return 0
}

func (x *SandboxConfig) GetRamMb() int64 {
	if x != nil {
		return x.RamMb
	}
	return 0
}

func (x *SandboxConfig) GetTotalDiskSizeMb() int64 {
	if x != nil {
		return x.TotalDiskSizeMb
	}
	return 0
}

func (x *SandboxConfig) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *SandboxConfig) GetMaxSandboxLength() int64 {
	if x != nil {
		return x.MaxSandboxLength
	}
	return 0
}

type SandboxCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sandbox   *SandboxConfig         `protobuf:"bytes,1,opt,name=sandbox,proto3" json:"sandbox,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *SandboxCreateRequest) Reset() {
	*x = SandboxCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxCreateRequest) ProtoMessage() {}

func (x *SandboxCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxCreateRequest.ProtoReflect.Descriptor instead.
func (*SandboxCreateRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{1}
}

func (x *SandboxCreateRequest) GetSandbox() *SandboxConfig {
	if x != nil {
		return x.Sandbox
	}
	return nil
}

func (x *SandboxCreateRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *SandboxCreateRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type SandboxCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *SandboxCreateResponse) Reset() {
	*x = SandboxCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxCreateResponse) ProtoMessage() {}

func (x *SandboxCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxCreateResponse.ProtoReflect.Descriptor instead.
func (*SandboxCreateResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{2}
}

func (x *SandboxCreateResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type SandboxUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxId string                 `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *SandboxUpdateRequest) Reset() {
	*x = SandboxUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxUpdateRequest) ProtoMessage() {}

func (x *SandboxUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxUpdateRequest.ProtoReflect.Descriptor instead.
func (*SandboxUpdateRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{3}
}

func (x *SandboxUpdateRequest) GetSandboxId() string {
	if x != nil {
		return x.SandboxId
	}
	return ""
}

func (x *SandboxUpdateRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type SandboxDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxId string `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
}

func (x *SandboxDeleteRequest) Reset() {
	*x = SandboxDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxDeleteRequest) ProtoMessage() {}

func (x *SandboxDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxDeleteRequest.ProtoReflect.Descriptor instead.
func (*SandboxDeleteRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{4}
}

func (x *SandboxDeleteRequest) GetSandboxId() string {
	if x != nil {
		return x.SandboxId
	}
	return ""
}

type SandboxPauseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxId  string `protobuf:"bytes,1,opt,name=sandbox_id,json=sandboxId,proto3" json:"sandbox_id,omitempty"`
	TemplateId string `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	BuildId    string `protobuf:"bytes,3,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
}

func (x *SandboxPauseRequest) Reset() {
	*x = SandboxPauseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxPauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxPauseRequest) ProtoMessage() {}

func (x *SandboxPauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxPauseRequest.ProtoReflect.Descriptor instead.
func (*SandboxPauseRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{5}
}

func (x *SandboxPauseRequest) GetSandboxId() string {
	if x != nil {
		return x.SandboxId
	}
	return ""
}

func (x *SandboxPauseRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *SandboxPauseRequest) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type RunningSandbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config    *SandboxConfig         `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	ClientId  string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *RunningSandbox) Reset() {
	*x = RunningSandbox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunningSandbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunningSandbox) ProtoMessage() {}

func (x *RunningSandbox) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunningSandbox.ProtoReflect.Descriptor instead.
func (*RunningSandbox) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{6}
}

func (x *RunningSandbox) GetConfig() *SandboxConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RunningSandbox) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *RunningSandbox) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *RunningSandbox) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type SandboxListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sandboxes []*RunningSandbox `protobuf:"bytes,1,rep,name=sandboxes,proto3" json:"sandboxes,omitempty"`
}

func (x *SandboxListResponse) Reset() {
	*x = SandboxListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxListResponse) ProtoMessage() {}

func (x *SandboxListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxListResponse.ProtoReflect.Descriptor instead.
func (*SandboxListResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{7}
}

func (x *SandboxListResponse) GetSandboxes() []*RunningSandbox {
	if x != nil {
		return x.Sandboxes
	}
	return nil
}

var File_orchestrator_proto protoreflect.FileDescriptor

var file_orchestrator_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb3, 0x05, 0x0a, 0x0d, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x65, 0x63,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x69, 0x72, 0x65, 0x63, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x75, 0x67, 0x65,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x75,
	0x67, 0x65, 0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x5f, 0x76, 0x61,
	0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62,
	0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x12, 0x38,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x64, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x64, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x63, 0x70, 0x75, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x76, 0x63, 0x70, 0x75, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x61,
	0x6d, 0x5f, 0x6d, 0x62, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x61, 0x6d, 0x4d,
	0x62, 0x12, 0x2b, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x62, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x4d, 0x62, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x1a, 0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x14, 0x53, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x07, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x39, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x34, 0x0a,
	0x15, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x6c, 0x0a, 0x14, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x35, 0x0a, 0x14, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x13, 0x53, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x78, 0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x0e, 0x52,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x26, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x13, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52,
	0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x32, 0xa8, 0x02, 0x0a, 0x0e, 0x53,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x34, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x14, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35,
	0x0a, 0x05, 0x50, 0x61, 0x75, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x50, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2f, 0x5a, 0x2d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x32, 0x62, 0x2d,
	0x64, 0x65, 0x76, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orchestrator_proto_rawDescOnce sync.Once
	file_orchestrator_proto_rawDescData = file_orchestrator_proto_rawDesc
)

func file_orchestrator_proto_rawDescGZIP() []byte {
	file_orchestrator_proto_rawDescOnce.Do(func() {
		file_orchestrator_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestrator_proto_rawDescData)
	})
	return file_orchestrator_proto_rawDescData
}

var file_orchestrator_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_orchestrator_proto_goTypes = []any{
	(*SandboxConfig)(nil),         // 0: SandboxConfig
	(*SandboxCreateRequest)(nil),  // 1: SandboxCreateRequest
	(*SandboxCreateResponse)(nil), // 2: SandboxCreateResponse
	(*SandboxUpdateRequest)(nil),  // 3: SandboxUpdateRequest
	(*SandboxDeleteRequest)(nil),  // 4: SandboxDeleteRequest
	(*SandboxPauseRequest)(nil),   // 5: SandboxPauseRequest
	(*RunningSandbox)(nil),        // 6: RunningSandbox
	(*SandboxListResponse)(nil),   // 7: SandboxListResponse
	nil,                           // 8: SandboxConfig.EnvVarsEntry
	nil,                           // 9: SandboxConfig.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_orchestrator_proto_depIdxs = []int32{
	8,  // 0: SandboxConfig.env_vars:type_name -> SandboxConfig.EnvVarsEntry
	9,  // 1: SandboxConfig.metadata:type_name -> SandboxConfig.MetadataEntry
	0,  // 2: SandboxCreateRequest.sandbox:type_name -> SandboxConfig
	10, // 3: SandboxCreateRequest.start_time:type_name -> google.protobuf.Timestamp
	10, // 4: SandboxCreateRequest.end_time:type_name -> google.protobuf.Timestamp
	10, // 5: SandboxUpdateRequest.end_time:type_name -> google.protobuf.Timestamp
	0,  // 6: RunningSandbox.config:type_name -> SandboxConfig
	10, // 7: RunningSandbox.start_time:type_name -> google.protobuf.Timestamp
	10, // 8: RunningSandbox.end_time:type_name -> google.protobuf.Timestamp
	6,  // 9: SandboxListResponse.sandboxes:type_name -> RunningSandbox
	1,  // 10: SandboxService.Create:input_type -> SandboxCreateRequest
	3,  // 11: SandboxService.Update:input_type -> SandboxUpdateRequest
	11, // 12: SandboxService.List:input_type -> google.protobuf.Empty
	4,  // 13: SandboxService.Delete:input_type -> SandboxDeleteRequest
	5,  // 14: SandboxService.Pause:input_type -> SandboxPauseRequest
	2,  // 15: SandboxService.Create:output_type -> SandboxCreateResponse
	11, // 16: SandboxService.Update:output_type -> google.protobuf.Empty
	7,  // 17: SandboxService.List:output_type -> SandboxListResponse
	11, // 18: SandboxService.Delete:output_type -> google.protobuf.Empty
	11, // 19: SandboxService.Pause:output_type -> google.protobuf.Empty
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_orchestrator_proto_init() }
func file_orchestrator_proto_init() {
	if File_orchestrator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orchestrator_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SandboxConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orchestrator_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SandboxCreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orchestrator_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SandboxCreateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orchestrator_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SandboxUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orchestrator_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SandboxDeleteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orchestrator_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SandboxPauseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orchestrator_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RunningSandbox); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_orchestrator_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SandboxListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_orchestrator_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orchestrator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestrator_proto_goTypes,
		DependencyIndexes: file_orchestrator_proto_depIdxs,
		MessageInfos:      file_orchestrator_proto_msgTypes,
	}.Build()
	File_orchestrator_proto = out.File
	file_orchestrator_proto_rawDesc = nil
	file_orchestrator_proto_goTypes = nil
	file_orchestrator_proto_depIdxs = nil
}
