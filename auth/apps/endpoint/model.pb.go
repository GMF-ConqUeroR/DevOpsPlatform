// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.17.3
// source: auth/apps/endpoint/pb/model.proto

package endpoint

import (
	resource "github.com/infraboard/mcube/pb/resource"
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

type EndpointSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务名称
	// @gotags: json:"service" bson:"service"
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service" bson:"service"`
	// 功能名称
	// @gotags: json:"action" bson:"action"
	Action string `protobuf:"bytes,2,opt,name=action,proto3" json:"action" bson:"action"`
	// 功能描述
	// @gotags: json:"descrption" bson:"descrption"
	Descrption string `protobuf:"bytes,3,opt,name=descrption,proto3" json:"descrption" bson:"descrption"`
}

func (x *EndpointSpec) Reset() {
	*x = EndpointSpec{}
	mi := &file_auth_apps_endpoint_pb_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSpec) ProtoMessage() {}

func (x *EndpointSpec) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_endpoint_pb_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSpec.ProtoReflect.Descriptor instead.
func (*EndpointSpec) Descriptor() ([]byte, []int) {
	return file_auth_apps_endpoint_pb_model_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointSpec) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *EndpointSpec) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *EndpointSpec) GetDescrption() string {
	if x != nil {
		return x.Descrption
	}
	return ""
}

type EndpointSpecSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务功能列表
	// @gotags: json:"items" bson:"items"
	Items []*EndpointSpec `protobuf:"bytes,1,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *EndpointSpecSet) Reset() {
	*x = EndpointSpecSet{}
	mi := &file_auth_apps_endpoint_pb_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointSpecSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSpecSet) ProtoMessage() {}

func (x *EndpointSpecSet) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_endpoint_pb_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSpecSet.ProtoReflect.Descriptor instead.
func (*EndpointSpecSet) Descriptor() ([]byte, []int) {
	return file_auth_apps_endpoint_pb_model_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointSpecSet) GetItems() []*EndpointSpec {
	if x != nil {
		return x.Items
	}
	return nil
}

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元数据信息
	// @gotags: json:"meta" bson:",inline"
	Meta *resource.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// Endpint定义
	// @gotags: json:"spec" bson:",inline"
	Spec *EndpointSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	mi := &file_auth_apps_endpoint_pb_model_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_endpoint_pb_model_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_auth_apps_endpoint_pb_model_proto_rawDescGZIP(), []int{2}
}

func (x *Endpoint) GetMeta() *resource.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Endpoint) GetSpec() *EndpointSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type EndpointSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务功能列表
	// @gotags: json:"items" bson:"items"
	Items []*Endpoint `protobuf:"bytes,1,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *EndpointSet) Reset() {
	*x = EndpointSet{}
	mi := &file_auth_apps_endpoint_pb_model_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointSet) ProtoMessage() {}

func (x *EndpointSet) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_endpoint_pb_model_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointSet.ProtoReflect.Descriptor instead.
func (*EndpointSet) Descriptor() ([]byte, []int) {
	return file_auth_apps_endpoint_pb_model_proto_rawDescGZIP(), []int{3}
}

func (x *EndpointSet) GetItems() []*Endpoint {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_auth_apps_endpoint_pb_model_proto protoreflect.FileDescriptor

var file_auth_apps_endpoint_pb_model_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x1c, 0x6d,
	0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x0c, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a,
	0x0f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x53, 0x65, 0x74,
	0x12, 0x3c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x7b,
	0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x3a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x47, 0x0a, 0x0b, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x6f,
	0x70, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x42, 0x14, 0x5a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_auth_apps_endpoint_pb_model_proto_rawDescOnce sync.Once
	file_auth_apps_endpoint_pb_model_proto_rawDescData = file_auth_apps_endpoint_pb_model_proto_rawDesc
)

func file_auth_apps_endpoint_pb_model_proto_rawDescGZIP() []byte {
	file_auth_apps_endpoint_pb_model_proto_rawDescOnce.Do(func() {
		file_auth_apps_endpoint_pb_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_apps_endpoint_pb_model_proto_rawDescData)
	})
	return file_auth_apps_endpoint_pb_model_proto_rawDescData
}

var file_auth_apps_endpoint_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_apps_endpoint_pb_model_proto_goTypes = []any{
	(*EndpointSpec)(nil),    // 0: devops_platform.endpoint.EndpointSpec
	(*EndpointSpecSet)(nil), // 1: devops_platform.endpoint.EndpointSpecSet
	(*Endpoint)(nil),        // 2: devops_platform.endpoint.Endpoint
	(*EndpointSet)(nil),     // 3: devops_platform.endpoint.EndpointSet
	(*resource.Meta)(nil),   // 4: infraboard.mcube.resource.Meta
}
var file_auth_apps_endpoint_pb_model_proto_depIdxs = []int32{
	0, // 0: devops_platform.endpoint.EndpointSpecSet.items:type_name -> devops_platform.endpoint.EndpointSpec
	4, // 1: devops_platform.endpoint.Endpoint.meta:type_name -> infraboard.mcube.resource.Meta
	0, // 2: devops_platform.endpoint.Endpoint.spec:type_name -> devops_platform.endpoint.EndpointSpec
	2, // 3: devops_platform.endpoint.EndpointSet.items:type_name -> devops_platform.endpoint.Endpoint
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_auth_apps_endpoint_pb_model_proto_init() }
func file_auth_apps_endpoint_pb_model_proto_init() {
	if File_auth_apps_endpoint_pb_model_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_apps_endpoint_pb_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_apps_endpoint_pb_model_proto_goTypes,
		DependencyIndexes: file_auth_apps_endpoint_pb_model_proto_depIdxs,
		MessageInfos:      file_auth_apps_endpoint_pb_model_proto_msgTypes,
	}.Build()
	File_auth_apps_endpoint_pb_model_proto = out.File
	file_auth_apps_endpoint_pb_model_proto_rawDesc = nil
	file_auth_apps_endpoint_pb_model_proto_goTypes = nil
	file_auth_apps_endpoint_pb_model_proto_depIdxs = nil
}
