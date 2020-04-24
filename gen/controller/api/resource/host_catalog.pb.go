// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: controller/api/resource/v1/host_catalog.proto

package resource

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// HostCatalog contains fields which are settable and modifiable by the end user.  This message should be used in
// relevant requests.
type HostCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendlyName *wrappers.StringValue `protobuf:"bytes,1,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	// Marks the host_catalog as disabled.  Default is false.
	Disabled *wrappers.BoolValue `protobuf:"bytes,2,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Types that are assignable to TypeDetails:
	//	*HostCatalog_StaticHostCatalogDetails
	//	*HostCatalog_Ec2HostCatalogDetails
	TypeDetails isHostCatalog_TypeDetails `protobuf_oneof:"type_details"`
}

func (x *HostCatalog) Reset() {
	*x = HostCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCatalog) ProtoMessage() {}

func (x *HostCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCatalog.ProtoReflect.Descriptor instead.
func (*HostCatalog) Descriptor() ([]byte, []int) {
	return file_controller_api_resource_v1_host_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *HostCatalog) GetFriendlyName() *wrappers.StringValue {
	if x != nil {
		return x.FriendlyName
	}
	return nil
}

func (x *HostCatalog) GetDisabled() *wrappers.BoolValue {
	if x != nil {
		return x.Disabled
	}
	return nil
}

func (m *HostCatalog) GetTypeDetails() isHostCatalog_TypeDetails {
	if m != nil {
		return m.TypeDetails
	}
	return nil
}

func (x *HostCatalog) GetStaticHostCatalogDetails() *StaticHostCatalog {
	if x, ok := x.GetTypeDetails().(*HostCatalog_StaticHostCatalogDetails); ok {
		return x.StaticHostCatalogDetails
	}
	return nil
}

func (x *HostCatalog) GetEc2HostCatalogDetails() *Ec2HostCatalog {
	if x, ok := x.GetTypeDetails().(*HostCatalog_Ec2HostCatalogDetails); ok {
		return x.Ec2HostCatalogDetails
	}
	return nil
}

type isHostCatalog_TypeDetails interface {
	isHostCatalog_TypeDetails()
}

type HostCatalog_StaticHostCatalogDetails struct {
	StaticHostCatalogDetails *StaticHostCatalog `protobuf:"bytes,3,opt,name=static_host_catalog_details,json=staticHostCatalogDetails,proto3,oneof"`
}

type HostCatalog_Ec2HostCatalogDetails struct {
	Ec2HostCatalogDetails *Ec2HostCatalog `protobuf:"bytes,4,opt,name=ec2_host_catalog_details,json=ec2HostCatalogDetails,proto3,oneof"`
}

func (*HostCatalog_StaticHostCatalogDetails) isHostCatalog_TypeDetails() {}

func (*HostCatalog_Ec2HostCatalogDetails) isHostCatalog_TypeDetails() {}

// HostCatalogResult contains all fields related to a HostCatalog resource.  The result object should be used in
// responses but never in requests.
type HostCatalogResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must be in the format "org/<org_id>/project/<project_id>/host_catalogs/<catalog_id>"
	Uri          string                `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	FriendlyName *wrappers.StringValue `protobuf:"bytes,2,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	// The time this host_catalog was created.
	CreatedTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	// The time this host_catalog was last updated.
	UpdatedTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_time,json=updatedTime,proto3" json:"updated_time,omitempty"`
	// Marks the host_catalog as disabled.  Default is false.
	Disabled *wrappers.BoolValue `protobuf:"bytes,5,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Types that are assignable to TypeDetails:
	//	*HostCatalogResult_StaticHostCatalogDetails
	//	*HostCatalogResult_Ec2HostCatalogDetails
	TypeDetails isHostCatalogResult_TypeDetails `protobuf_oneof:"type_details"`
}

func (x *HostCatalogResult) Reset() {
	*x = HostCatalogResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCatalogResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCatalogResult) ProtoMessage() {}

func (x *HostCatalogResult) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCatalogResult.ProtoReflect.Descriptor instead.
func (*HostCatalogResult) Descriptor() ([]byte, []int) {
	return file_controller_api_resource_v1_host_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *HostCatalogResult) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *HostCatalogResult) GetFriendlyName() *wrappers.StringValue {
	if x != nil {
		return x.FriendlyName
	}
	return nil
}

func (x *HostCatalogResult) GetCreatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *HostCatalogResult) GetUpdatedTime() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *HostCatalogResult) GetDisabled() *wrappers.BoolValue {
	if x != nil {
		return x.Disabled
	}
	return nil
}

func (m *HostCatalogResult) GetTypeDetails() isHostCatalogResult_TypeDetails {
	if m != nil {
		return m.TypeDetails
	}
	return nil
}

func (x *HostCatalogResult) GetStaticHostCatalogDetails() *StaticHostCatalogResult {
	if x, ok := x.GetTypeDetails().(*HostCatalogResult_StaticHostCatalogDetails); ok {
		return x.StaticHostCatalogDetails
	}
	return nil
}

func (x *HostCatalogResult) GetEc2HostCatalogDetails() *Ec2HostCatalogResult {
	if x, ok := x.GetTypeDetails().(*HostCatalogResult_Ec2HostCatalogDetails); ok {
		return x.Ec2HostCatalogDetails
	}
	return nil
}

type isHostCatalogResult_TypeDetails interface {
	isHostCatalogResult_TypeDetails()
}

type HostCatalogResult_StaticHostCatalogDetails struct {
	StaticHostCatalogDetails *StaticHostCatalogResult `protobuf:"bytes,6,opt,name=static_host_catalog_details,json=staticHostCatalogDetails,proto3,oneof"`
}

type HostCatalogResult_Ec2HostCatalogDetails struct {
	Ec2HostCatalogDetails *Ec2HostCatalogResult `protobuf:"bytes,7,opt,name=ec2_host_catalog_details,json=ec2HostCatalogDetails,proto3,oneof"`
}

func (*HostCatalogResult_StaticHostCatalogDetails) isHostCatalogResult_TypeDetails() {}

func (*HostCatalogResult_Ec2HostCatalogDetails) isHostCatalogResult_TypeDetails() {}

type StaticHostCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StaticHostCatalog) Reset() {
	*x = StaticHostCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticHostCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticHostCatalog) ProtoMessage() {}

func (x *StaticHostCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticHostCatalog.ProtoReflect.Descriptor instead.
func (*StaticHostCatalog) Descriptor() ([]byte, []int) {
	return file_controller_api_resource_v1_host_catalog_proto_rawDescGZIP(), []int{2}
}

type StaticHostCatalogResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StaticHostCatalogResult) Reset() {
	*x = StaticHostCatalogResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticHostCatalogResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticHostCatalogResult) ProtoMessage() {}

func (x *StaticHostCatalogResult) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticHostCatalogResult.ProtoReflect.Descriptor instead.
func (*StaticHostCatalogResult) Descriptor() ([]byte, []int) {
	return file_controller_api_resource_v1_host_catalog_proto_rawDescGZIP(), []int{3}
}

type Ec2HostCatalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The AWS regions from which this catalog will retrieve the EC2 instances.
	Regions []string `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
	// The access key used for authenticating with AWS when retrieving EC2 instance details.
	AccessKey *wrappers.StringValue `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey *wrappers.StringValue `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	Rotate    *wrappers.BoolValue   `protobuf:"bytes,4,opt,name=rotate,proto3" json:"rotate,omitempty"`
}

func (x *Ec2HostCatalog) Reset() {
	*x = Ec2HostCatalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ec2HostCatalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ec2HostCatalog) ProtoMessage() {}

func (x *Ec2HostCatalog) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ec2HostCatalog.ProtoReflect.Descriptor instead.
func (*Ec2HostCatalog) Descriptor() ([]byte, []int) {
	return file_controller_api_resource_v1_host_catalog_proto_rawDescGZIP(), []int{4}
}

func (x *Ec2HostCatalog) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Ec2HostCatalog) GetAccessKey() *wrappers.StringValue {
	if x != nil {
		return x.AccessKey
	}
	return nil
}

func (x *Ec2HostCatalog) GetSecretKey() *wrappers.StringValue {
	if x != nil {
		return x.SecretKey
	}
	return nil
}

func (x *Ec2HostCatalog) GetRotate() *wrappers.BoolValue {
	if x != nil {
		return x.Rotate
	}
	return nil
}

type Ec2HostCatalogResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The AWS regions from which this catalog will retrieve the EC2 instances.
	Regions []string `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
	// The access key used for authenticating with AWS when retrieving EC2 instance details.
	AccessKey *wrappers.StringValue `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
}

func (x *Ec2HostCatalogResult) Reset() {
	*x = Ec2HostCatalogResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ec2HostCatalogResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ec2HostCatalogResult) ProtoMessage() {}

func (x *Ec2HostCatalogResult) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resource_v1_host_catalog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ec2HostCatalogResult.ProtoReflect.Descriptor instead.
func (*Ec2HostCatalogResult) Descriptor() ([]byte, []int) {
	return file_controller_api_resource_v1_host_catalog_proto_rawDescGZIP(), []int{5}
}

func (x *Ec2HostCatalogResult) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *Ec2HostCatalogResult) GetAccessKey() *wrappers.StringValue {
	if x != nil {
		return x.AccessKey
	}
	return nil
}

var File_controller_api_resource_v1_host_catalog_proto protoreflect.FileDescriptor

var file_controller_api_resource_v1_host_catalog_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x02, 0x0a,
	0x0b, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x41, 0x0a, 0x0d,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x36, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x6e, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x18, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x65, 0x0a, 0x18, 0x65, 0x63, 0x32, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x32, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x15, 0x65, 0x63, 0x32, 0x48, 0x6f, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x0e,
	0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x91,
	0x04, 0x0a, 0x11, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x41, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x66, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x74, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x18, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x6b, 0x0a, 0x18, 0x65, 0x63, 0x32, 0x5f, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x32, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x15, 0x65, 0x63, 0x32,
	0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x48, 0x6f, 0x73, 0x74,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x63, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0xd8, 0x01, 0x0a, 0x0e, 0x45, 0x63, 0x32, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x3b, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x0a,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x22, 0x6d, 0x0a,
	0x14, 0x45, 0x63, 0x32, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x3b, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x42, 0x46, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x74, 0x6f, 0x77, 0x65, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_resource_v1_host_catalog_proto_rawDescOnce sync.Once
	file_controller_api_resource_v1_host_catalog_proto_rawDescData = file_controller_api_resource_v1_host_catalog_proto_rawDesc
)

func file_controller_api_resource_v1_host_catalog_proto_rawDescGZIP() []byte {
	file_controller_api_resource_v1_host_catalog_proto_rawDescOnce.Do(func() {
		file_controller_api_resource_v1_host_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resource_v1_host_catalog_proto_rawDescData)
	})
	return file_controller_api_resource_v1_host_catalog_proto_rawDescData
}

var file_controller_api_resource_v1_host_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_controller_api_resource_v1_host_catalog_proto_goTypes = []interface{}{
	(*HostCatalog)(nil),             // 0: controller.api.resource.v1.HostCatalog
	(*HostCatalogResult)(nil),       // 1: controller.api.resource.v1.HostCatalogResult
	(*StaticHostCatalog)(nil),       // 2: controller.api.resource.v1.StaticHostCatalog
	(*StaticHostCatalogResult)(nil), // 3: controller.api.resource.v1.StaticHostCatalogResult
	(*Ec2HostCatalog)(nil),          // 4: controller.api.resource.v1.Ec2HostCatalog
	(*Ec2HostCatalogResult)(nil),    // 5: controller.api.resource.v1.Ec2HostCatalogResult
	(*wrappers.StringValue)(nil),    // 6: google.protobuf.StringValue
	(*wrappers.BoolValue)(nil),      // 7: google.protobuf.BoolValue
	(*timestamp.Timestamp)(nil),     // 8: google.protobuf.Timestamp
}
var file_controller_api_resource_v1_host_catalog_proto_depIdxs = []int32{
	6,  // 0: controller.api.resource.v1.HostCatalog.friendly_name:type_name -> google.protobuf.StringValue
	7,  // 1: controller.api.resource.v1.HostCatalog.disabled:type_name -> google.protobuf.BoolValue
	2,  // 2: controller.api.resource.v1.HostCatalog.static_host_catalog_details:type_name -> controller.api.resource.v1.StaticHostCatalog
	4,  // 3: controller.api.resource.v1.HostCatalog.ec2_host_catalog_details:type_name -> controller.api.resource.v1.Ec2HostCatalog
	6,  // 4: controller.api.resource.v1.HostCatalogResult.friendly_name:type_name -> google.protobuf.StringValue
	8,  // 5: controller.api.resource.v1.HostCatalogResult.created_time:type_name -> google.protobuf.Timestamp
	8,  // 6: controller.api.resource.v1.HostCatalogResult.updated_time:type_name -> google.protobuf.Timestamp
	7,  // 7: controller.api.resource.v1.HostCatalogResult.disabled:type_name -> google.protobuf.BoolValue
	3,  // 8: controller.api.resource.v1.HostCatalogResult.static_host_catalog_details:type_name -> controller.api.resource.v1.StaticHostCatalogResult
	5,  // 9: controller.api.resource.v1.HostCatalogResult.ec2_host_catalog_details:type_name -> controller.api.resource.v1.Ec2HostCatalogResult
	6,  // 10: controller.api.resource.v1.Ec2HostCatalog.access_key:type_name -> google.protobuf.StringValue
	6,  // 11: controller.api.resource.v1.Ec2HostCatalog.secret_key:type_name -> google.protobuf.StringValue
	7,  // 12: controller.api.resource.v1.Ec2HostCatalog.rotate:type_name -> google.protobuf.BoolValue
	6,  // 13: controller.api.resource.v1.Ec2HostCatalogResult.access_key:type_name -> google.protobuf.StringValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_controller_api_resource_v1_host_catalog_proto_init() }
func file_controller_api_resource_v1_host_catalog_proto_init() {
	if File_controller_api_resource_v1_host_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resource_v1_host_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCatalog); i {
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
		file_controller_api_resource_v1_host_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostCatalogResult); i {
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
		file_controller_api_resource_v1_host_catalog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticHostCatalog); i {
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
		file_controller_api_resource_v1_host_catalog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticHostCatalogResult); i {
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
		file_controller_api_resource_v1_host_catalog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ec2HostCatalog); i {
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
		file_controller_api_resource_v1_host_catalog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ec2HostCatalogResult); i {
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
	file_controller_api_resource_v1_host_catalog_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HostCatalog_StaticHostCatalogDetails)(nil),
		(*HostCatalog_Ec2HostCatalogDetails)(nil),
	}
	file_controller_api_resource_v1_host_catalog_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HostCatalogResult_StaticHostCatalogDetails)(nil),
		(*HostCatalogResult_Ec2HostCatalogDetails)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_api_resource_v1_host_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resource_v1_host_catalog_proto_goTypes,
		DependencyIndexes: file_controller_api_resource_v1_host_catalog_proto_depIdxs,
		MessageInfos:      file_controller_api_resource_v1_host_catalog_proto_msgTypes,
	}.Build()
	File_controller_api_resource_v1_host_catalog_proto = out.File
	file_controller_api_resource_v1_host_catalog_proto_rawDesc = nil
	file_controller_api_resource_v1_host_catalog_proto_goTypes = nil
	file_controller_api_resource_v1_host_catalog_proto_depIdxs = nil
}