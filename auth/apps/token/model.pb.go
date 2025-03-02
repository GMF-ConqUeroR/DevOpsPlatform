// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.17.3
// source: auth/apps/token/pb/model.proto

package token

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

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 访问令牌
	// @gotags: json:"access_token" bson:"access_token"
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token" bson:"access_token"`
	// 刷新Token
	// @gotags: json:"refresh_token" bson:"refresh_token"
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token" bson:"refresh_token"`
	// 令牌颁发时间
	// @gotags: json:"issue_at" bson:"issue_at"
	IssueAt int64 `protobuf:"varint,3,opt,name=issue_at,json=issueAt,proto3" json:"issue_at" bson:"issue_at"`
	// 多少秒后过期
	// @gotags: json:"access_token_expired_second" bson:"access_token_expired_second"
	AccessTokenExpiredSecond int64 `protobuf:"varint,4,opt,name=access_token_expired_second,json=accessTokenExpiredSecond,proto3" json:"access_token_expired_second" bson:"access_token_expired_second"`
	// 多少秒后过期
	// @gotags: json:"refresh_token_expired_second" bson:"refresh_token_expired_second"
	RefreshTokenExpiredSecond int64 `protobuf:"varint,6,opt,name=refresh_token_expired_second,json=refreshTokenExpiredSecond,proto3" json:"refresh_token_expired_second" bson:"refresh_token_expired_second"`
	// 颁发的用户
	// @gotags: json:"username" bson:"username"
	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username" bson:"username"`
}

func (x *Token) Reset() {
	*x = Token{}
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_auth_apps_token_pb_model_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Token) GetIssueAt() int64 {
	if x != nil {
		return x.IssueAt
	}
	return 0
}

func (x *Token) GetAccessTokenExpiredSecond() int64 {
	if x != nil {
		return x.AccessTokenExpiredSecond
	}
	return 0
}

func (x *Token) GetRefreshTokenExpiredSecond() int64 {
	if x != nil {
		return x.RefreshTokenExpiredSecond
	}
	return 0
}

func (x *Token) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type IssueTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户名称
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// 密码
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *IssueTokenRequest) Reset() {
	*x = IssueTokenRequest{}
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IssueTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueTokenRequest) ProtoMessage() {}

func (x *IssueTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueTokenRequest.ProtoReflect.Descriptor instead.
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_apps_token_pb_model_proto_rawDescGZIP(), []int{1}
}

func (x *IssueTokenRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IssueTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RevolkTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要被撤销的Token
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// 使用Refresh Token做撤销验证
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RevolkTokenRequest) Reset() {
	*x = RevolkTokenRequest{}
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevolkTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevolkTokenRequest) ProtoMessage() {}

func (x *RevolkTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevolkTokenRequest.ProtoReflect.Descriptor instead.
func (*RevolkTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_apps_token_pb_model_proto_rawDescGZIP(), []int{2}
}

func (x *RevolkTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RevolkTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type ValidateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要验证的Token
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_apps_token_pb_model_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type DescribeTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 需要查看的Token
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *DescribeTokenRequest) Reset() {
	*x = DescribeTokenRequest{}
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeTokenRequest) ProtoMessage() {}

func (x *DescribeTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_apps_token_pb_model_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeTokenRequest.ProtoReflect.Descriptor instead.
func (*DescribeTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_apps_token_pb_model_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_auth_apps_token_pb_model_proto protoreflect.FileDescriptor

var file_auth_apps_token_pb_model_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x86, 0x02, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x18, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x12, 0x3f, 0x0a, 0x1c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x19, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x4b, 0x0a, 0x11, 0x49, 0x73, 0x73, 0x75, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5c, 0x0a,
	0x12, 0x52, 0x65, 0x76, 0x6f, 0x6c, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x14, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x39, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x11, 0x5a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_apps_token_pb_model_proto_rawDescOnce sync.Once
	file_auth_apps_token_pb_model_proto_rawDescData = file_auth_apps_token_pb_model_proto_rawDesc
)

func file_auth_apps_token_pb_model_proto_rawDescGZIP() []byte {
	file_auth_apps_token_pb_model_proto_rawDescOnce.Do(func() {
		file_auth_apps_token_pb_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_apps_token_pb_model_proto_rawDescData)
	})
	return file_auth_apps_token_pb_model_proto_rawDescData
}

var file_auth_apps_token_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_auth_apps_token_pb_model_proto_goTypes = []any{
	(*Token)(nil),                // 0: devops_platform.token.Token
	(*IssueTokenRequest)(nil),    // 1: devops_platform.token.IssueTokenRequest
	(*RevolkTokenRequest)(nil),   // 2: devops_platform.token.RevolkTokenRequest
	(*ValidateTokenRequest)(nil), // 3: devops_platform.token.ValidateTokenRequest
	(*DescribeTokenRequest)(nil), // 4: devops_platform.token.DescribeTokenRequest
}
var file_auth_apps_token_pb_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_apps_token_pb_model_proto_init() }
func file_auth_apps_token_pb_model_proto_init() {
	if File_auth_apps_token_pb_model_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_apps_token_pb_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_apps_token_pb_model_proto_goTypes,
		DependencyIndexes: file_auth_apps_token_pb_model_proto_depIdxs,
		MessageInfos:      file_auth_apps_token_pb_model_proto_msgTypes,
	}.Build()
	File_auth_apps_token_pb_model_proto = out.File
	file_auth_apps_token_pb_model_proto_rawDesc = nil
	file_auth_apps_token_pb_model_proto_goTypes = nil
	file_auth_apps_token_pb_model_proto_depIdxs = nil
}
