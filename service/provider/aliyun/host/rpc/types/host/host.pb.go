// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: host/rpc/host.proto

package host

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 创建结构体
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKeyId     string `protobuf:"bytes,1,opt,name=AccessKeyId,proto3" json:"AccessKeyId,omitempty"`
	AccessKeySecret string `protobuf:"bytes,2,opt,name=AccessKeySecret,proto3" json:"AccessKeySecret,omitempty"`
	Vendor          string `protobuf:"bytes,3,opt,name=Vendor,proto3" json:"Vendor,omitempty"`
	Region          string `protobuf:"bytes,6,opt,name=Region,proto3" json:"Region,omitempty"`
	TaskType        string `protobuf:"bytes,4,opt,name=TaskType,proto3" json:"TaskType,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_rpc_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_rpc_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_host_rpc_host_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *CreateRequest) GetAccessKeySecret() string {
	if x != nil {
		return x.AccessKeySecret
	}
	return ""
}

func (x *CreateRequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *CreateRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateRequest) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId              string `protobuf:"bytes,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	Regionid                string `protobuf:"bytes,2,opt,name=Regionid,proto3" json:"Regionid,omitempty"`
	InstanceName            string `protobuf:"bytes,3,opt,name=InstanceName,proto3" json:"InstanceName,omitempty"`
	ExpiredTime             string `protobuf:"bytes,4,opt,name=ExpiredTime,proto3" json:"ExpiredTime,omitempty"`
	CreationTime            string `protobuf:"bytes,5,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	KeyPairName             string `protobuf:"bytes,6,opt,name=KeyPairName,proto3" json:"KeyPairName,omitempty"`
	Description             string `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`
	OsName                  string `protobuf:"bytes,8,opt,name=OsName,proto3" json:"OsName,omitempty"`
	ImageId                 int64  `protobuf:"varint,9,opt,name=ImageId,proto3" json:"ImageId,omitempty"`
	GpuAmount               string `protobuf:"bytes,10,opt,name=GpuAmount,proto3" json:"GpuAmount,omitempty"`
	Cpu                     string `protobuf:"bytes,11,opt,name=Cpu,proto3" json:"Cpu,omitempty"`
	Memory                  string `protobuf:"bytes,12,opt,name=Memory,proto3" json:"Memory,omitempty"`
	OsType                  string `protobuf:"bytes,13,opt,name=OsType,proto3" json:"OsType,omitempty"`
	InstanceType            string `protobuf:"bytes,14,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	InstanceChargeType      string `protobuf:"bytes,15,opt,name=InstanceChargeType,proto3" json:"InstanceChargeType,omitempty"`
	InternetMaxBandwidthOut int64  `protobuf:"varint,16,opt,name=InternetMaxBandwidth_out,json=InternetMaxBandwidthOut,proto3" json:"InternetMaxBandwidth_out,omitempty"`
	InternetMaxBandwidthIn  int64  `protobuf:"varint,17,opt,name=InternetMaxBandwidth_in,json=InternetMaxBandwidthIn,proto3" json:"InternetMaxBandwidth_in,omitempty"`
	Primaryip               string `protobuf:"bytes,18,opt,name=Primaryip,proto3" json:"Primaryip,omitempty"`
	Publicip                string `protobuf:"bytes,19,opt,name=Publicip,proto3" json:"Publicip,omitempty"`
	EipAddresses            string `protobuf:"bytes,20,opt,name=EipAddresses,proto3" json:"EipAddresses,omitempty"`
	SecurityGroupId         string `protobuf:"bytes,21,opt,name=security_group_id,json=securityGroupId,proto3" json:"security_group_id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_rpc_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_host_rpc_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_host_rpc_host_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *CreateResponse) GetRegionid() string {
	if x != nil {
		return x.Regionid
	}
	return ""
}

func (x *CreateResponse) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *CreateResponse) GetExpiredTime() string {
	if x != nil {
		return x.ExpiredTime
	}
	return ""
}

func (x *CreateResponse) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *CreateResponse) GetKeyPairName() string {
	if x != nil {
		return x.KeyPairName
	}
	return ""
}

func (x *CreateResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateResponse) GetOsName() string {
	if x != nil {
		return x.OsName
	}
	return ""
}

func (x *CreateResponse) GetImageId() int64 {
	if x != nil {
		return x.ImageId
	}
	return 0
}

func (x *CreateResponse) GetGpuAmount() string {
	if x != nil {
		return x.GpuAmount
	}
	return ""
}

func (x *CreateResponse) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *CreateResponse) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *CreateResponse) GetOsType() string {
	if x != nil {
		return x.OsType
	}
	return ""
}

func (x *CreateResponse) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *CreateResponse) GetInstanceChargeType() string {
	if x != nil {
		return x.InstanceChargeType
	}
	return ""
}

func (x *CreateResponse) GetInternetMaxBandwidthOut() int64 {
	if x != nil {
		return x.InternetMaxBandwidthOut
	}
	return 0
}

func (x *CreateResponse) GetInternetMaxBandwidthIn() int64 {
	if x != nil {
		return x.InternetMaxBandwidthIn
	}
	return 0
}

func (x *CreateResponse) GetPrimaryip() string {
	if x != nil {
		return x.Primaryip
	}
	return ""
}

func (x *CreateResponse) GetPublicip() string {
	if x != nil {
		return x.Publicip
	}
	return ""
}

func (x *CreateResponse) GetEipAddresses() string {
	if x != nil {
		return x.EipAddresses
	}
	return ""
}

func (x *CreateResponse) GetSecurityGroupId() string {
	if x != nil {
		return x.SecurityGroupId
	}
	return ""
}

// 删除结构体
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId string `protobuf:"bytes,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_rpc_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_rpc_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_host_rpc_host_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId              string `protobuf:"bytes,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	Regionid                string `protobuf:"bytes,2,opt,name=Regionid,proto3" json:"Regionid,omitempty"`
	InstanceName            string `protobuf:"bytes,3,opt,name=InstanceName,proto3" json:"InstanceName,omitempty"`
	ExpiredTime             string `protobuf:"bytes,4,opt,name=ExpiredTime,proto3" json:"ExpiredTime,omitempty"`
	CreationTime            string `protobuf:"bytes,5,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	KeyPairName             string `protobuf:"bytes,6,opt,name=KeyPairName,proto3" json:"KeyPairName,omitempty"`
	Description             string `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`
	OsName                  string `protobuf:"bytes,8,opt,name=OsName,proto3" json:"OsName,omitempty"`
	ImageId                 string `protobuf:"bytes,9,opt,name=ImageId,proto3" json:"ImageId,omitempty"`
	GpuAmount               int64  `protobuf:"varint,10,opt,name=GpuAmount,proto3" json:"GpuAmount,omitempty"`
	Cpu                     int64  `protobuf:"varint,11,opt,name=Cpu,proto3" json:"Cpu,omitempty"`
	Memory                  int64  `protobuf:"varint,12,opt,name=Memory,proto3" json:"Memory,omitempty"`
	OsType                  string `protobuf:"bytes,13,opt,name=OsType,proto3" json:"OsType,omitempty"`
	InstanceType            string `protobuf:"bytes,14,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	InstanceChargeType      string `protobuf:"bytes,15,opt,name=InstanceChargeType,proto3" json:"InstanceChargeType,omitempty"`
	InternetMaxBandwidthOut int64  `protobuf:"varint,16,opt,name=InternetMaxBandwidth_out,json=InternetMaxBandwidthOut,proto3" json:"InternetMaxBandwidth_out,omitempty"`
	InternetMaxBandwidthIn  int64  `protobuf:"varint,17,opt,name=InternetMaxBandwidth_in,json=InternetMaxBandwidthIn,proto3" json:"InternetMaxBandwidth_in,omitempty"`
	Primaryip               string `protobuf:"bytes,18,opt,name=Primaryip,proto3" json:"Primaryip,omitempty"`
	Publicip                string `protobuf:"bytes,19,opt,name=Publicip,proto3" json:"Publicip,omitempty"`
	EipAddresses            string `protobuf:"bytes,20,opt,name=EipAddresses,proto3" json:"EipAddresses,omitempty"`
	SecurityGroupId         string `protobuf:"bytes,21,opt,name=security_group_id,json=securityGroupId,proto3" json:"security_group_id,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_rpc_host_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_host_rpc_host_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_host_rpc_host_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteResponse) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *DeleteResponse) GetRegionid() string {
	if x != nil {
		return x.Regionid
	}
	return ""
}

func (x *DeleteResponse) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *DeleteResponse) GetExpiredTime() string {
	if x != nil {
		return x.ExpiredTime
	}
	return ""
}

func (x *DeleteResponse) GetCreationTime() string {
	if x != nil {
		return x.CreationTime
	}
	return ""
}

func (x *DeleteResponse) GetKeyPairName() string {
	if x != nil {
		return x.KeyPairName
	}
	return ""
}

func (x *DeleteResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DeleteResponse) GetOsName() string {
	if x != nil {
		return x.OsName
	}
	return ""
}

func (x *DeleteResponse) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *DeleteResponse) GetGpuAmount() int64 {
	if x != nil {
		return x.GpuAmount
	}
	return 0
}

func (x *DeleteResponse) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *DeleteResponse) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *DeleteResponse) GetOsType() string {
	if x != nil {
		return x.OsType
	}
	return ""
}

func (x *DeleteResponse) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *DeleteResponse) GetInstanceChargeType() string {
	if x != nil {
		return x.InstanceChargeType
	}
	return ""
}

func (x *DeleteResponse) GetInternetMaxBandwidthOut() int64 {
	if x != nil {
		return x.InternetMaxBandwidthOut
	}
	return 0
}

func (x *DeleteResponse) GetInternetMaxBandwidthIn() int64 {
	if x != nil {
		return x.InternetMaxBandwidthIn
	}
	return 0
}

func (x *DeleteResponse) GetPrimaryip() string {
	if x != nil {
		return x.Primaryip
	}
	return ""
}

func (x *DeleteResponse) GetPublicip() string {
	if x != nil {
		return x.Publicip
	}
	return ""
}

func (x *DeleteResponse) GetEipAddresses() string {
	if x != nil {
		return x.EipAddresses
	}
	return ""
}

func (x *DeleteResponse) GetSecurityGroupId() string {
	if x != nil {
		return x.SecurityGroupId
	}
	return ""
}

// 分页查询
type GetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_rpc_host_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_rpc_host_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_host_rpc_host_proto_rawDescGZIP(), []int{4}
}

func (x *GetListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DeleteResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_rpc_host_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_host_rpc_host_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_host_rpc_host_proto_rawDescGZIP(), []int{5}
}

func (x *GetListResponse) GetData() []*DeleteResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

// ID查询
type GetIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceId string `protobuf:"bytes,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
}

func (x *GetIdRequest) Reset() {
	*x = GetIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_host_rpc_host_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdRequest) ProtoMessage() {}

func (x *GetIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_host_rpc_host_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdRequest.ProtoReflect.Descriptor instead.
func (*GetIdRequest) Descriptor() ([]byte, []int) {
	return file_host_rpc_host_proto_rawDescGZIP(), []int{6}
}

func (x *GetIdRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

var File_host_rpc_host_proto protoreflect.FileDescriptor

var file_host_rpc_host_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x22, 0xde, 0x05, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x45,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x4f, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x70, 0x75, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x70, 0x75, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x70, 0x75, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x43, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68,
	0x61, 0x72, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x18, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x5f, 0x6f, 0x75, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x4f, 0x75, 0x74, 0x12, 0x37, 0x0a, 0x17, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d,
	0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61,
	0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x49, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x69, 0x70, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x69, 0x70, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x69, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x69, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0xde, 0x05, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4f, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x70, 0x75, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x47, 0x70, 0x75, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x70, 0x75, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x43, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x18, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x4f, 0x75, 0x74, 0x12, 0x37, 0x0a, 0x17, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x69, 0x6e,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x4d, 0x61, 0x78, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x49, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x69, 0x70, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x70, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x69, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x69, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x45, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x2e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x32, 0xa0, 0x02, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x48, 0x6f,
	0x73, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x13, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x6f,
	0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x48,
	0x6f, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x68, 0x6f, 0x73, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x09, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x68, 0x6f, 0x73,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x68, 0x6f, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_host_rpc_host_proto_rawDescOnce sync.Once
	file_host_rpc_host_proto_rawDescData = file_host_rpc_host_proto_rawDesc
)

func file_host_rpc_host_proto_rawDescGZIP() []byte {
	file_host_rpc_host_proto_rawDescOnce.Do(func() {
		file_host_rpc_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_host_rpc_host_proto_rawDescData)
	})
	return file_host_rpc_host_proto_rawDescData
}

var file_host_rpc_host_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_host_rpc_host_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),   // 0: host.CreateRequest
	(*CreateResponse)(nil),  // 1: host.CreateResponse
	(*DeleteRequest)(nil),   // 2: host.DeleteRequest
	(*DeleteResponse)(nil),  // 3: host.DeleteResponse
	(*GetListRequest)(nil),  // 4: host.GetListRequest
	(*GetListResponse)(nil), // 5: host.GetListResponse
	(*GetIdRequest)(nil),    // 6: host.GetIdRequest
}
var file_host_rpc_host_proto_depIdxs = []int32{
	3, // 0: host.GetListResponse.data:type_name -> host.DeleteResponse
	0, // 1: host.Host.HostSync:input_type -> host.CreateRequest
	2, // 2: host.Host.HostDelete:input_type -> host.DeleteRequest
	0, // 3: host.Host.HostUpdate:input_type -> host.CreateRequest
	4, // 4: host.Host.HostList:input_type -> host.GetListRequest
	6, // 5: host.Host.HostGetId:input_type -> host.GetIdRequest
	5, // 6: host.Host.HostSync:output_type -> host.GetListResponse
	3, // 7: host.Host.HostDelete:output_type -> host.DeleteResponse
	1, // 8: host.Host.HostUpdate:output_type -> host.CreateResponse
	5, // 9: host.Host.HostList:output_type -> host.GetListResponse
	3, // 10: host.Host.HostGetId:output_type -> host.DeleteResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_host_rpc_host_proto_init() }
func file_host_rpc_host_proto_init() {
	if File_host_rpc_host_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_host_rpc_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_host_rpc_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_host_rpc_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_host_rpc_host_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_host_rpc_host_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRequest); i {
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
		file_host_rpc_host_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListResponse); i {
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
		file_host_rpc_host_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_host_rpc_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_host_rpc_host_proto_goTypes,
		DependencyIndexes: file_host_rpc_host_proto_depIdxs,
		MessageInfos:      file_host_rpc_host_proto_msgTypes,
	}.Build()
	File_host_rpc_host_proto = out.File
	file_host_rpc_host_proto_rawDesc = nil
	file_host_rpc_host_proto_goTypes = nil
	file_host_rpc_host_proto_depIdxs = nil
}