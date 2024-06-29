// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: gRPC/proto/auth.proto

package auth

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

type ResponseType int32

const (
	ResponseType_SUCCESS      ResponseType = 0
	ResponseType_FAILED       ResponseType = 1
	ResponseType_EXPIRED_DATE ResponseType = 2
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
		2: "EXPIRED_DATE",
	}
	ResponseType_value = map[string]int32{
		"SUCCESS":      0,
		"FAILED":       1,
		"EXPIRED_DATE": 2,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_gRPC_proto_auth_proto_enumTypes[0].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_gRPC_proto_auth_proto_enumTypes[0]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{0}
}

type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	CreateDate int64  `protobuf:"varint,3,opt,name=createDate,proto3" json:"createDate,omitempty"`
	ExpireDate int64  `protobuf:"varint,4,opt,name=expireDate,proto3" json:"expireDate,omitempty"`
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthData) GetCreateDate() int64 {
	if x != nil {
		return x.CreateDate
	}
	return 0
}

func (x *AuthData) GetExpireDate() int64 {
	if x != nil {
		return x.ExpireDate
	}
	return 0
}

type Verify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ResponseType `protobuf:"varint,1,opt,name=status,proto3,enum=ResponseType" json:"status,omitempty"`
	Auth   *AuthData    `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *Verify) Reset() {
	*x = Verify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Verify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verify) ProtoMessage() {}

func (x *Verify) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verify.ProtoReflect.Descriptor instead.
func (*Verify) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *Verify) GetStatus() ResponseType {
	if x != nil {
		return x.Status
	}
	return ResponseType_SUCCESS
}

func (x *Verify) GetAuth() *AuthData {
	if x != nil {
		return x.Auth
	}
	return nil
}

type CreateTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *AuthData `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *CreateTokenReq) Reset() {
	*x = CreateTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenReq) ProtoMessage() {}

func (x *CreateTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenReq.ProtoReflect.Descriptor instead.
func (*CreateTokenReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTokenReq) GetAuth() *AuthData {
	if x != nil {
		return x.Auth
	}
	return nil
}

type CreateTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *AuthData `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *CreateTokenRes) Reset() {
	*x = CreateTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRes) ProtoMessage() {}

func (x *CreateTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRes.ProtoReflect.Descriptor instead.
func (*CreateTokenRes) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTokenRes) GetAuth() *AuthData {
	if x != nil {
		return x.Auth
	}
	return nil
}

type VerifyTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyTokenReq) Reset() {
	*x = VerifyTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenReq) ProtoMessage() {}

func (x *VerifyTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenReq.ProtoReflect.Descriptor instead.
func (*VerifyTokenReq) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V *Verify `protobuf:"bytes,1,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *VerifyTokenRes) Reset() {
	*x = VerifyTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRes) ProtoMessage() {}

func (x *VerifyTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRes.ProtoReflect.Descriptor instead.
func (*VerifyTokenRes) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyTokenRes) GetV() *Verify {
	if x != nil {
		return x.V
	}
	return nil
}

var File_gRPC_proto_auth_proto protoreflect.FileDescriptor

var file_gRPC_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x4e, 0x0a,
	0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x2f, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x1d, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x2f,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x1d, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22,
	0x26, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x27, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x01, 0x76, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x01, 0x76,
	0x2a, 0x39, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0x6d, 0x0a, 0x0b, 0x41,
	0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_proto_auth_proto_rawDescOnce sync.Once
	file_gRPC_proto_auth_proto_rawDescData = file_gRPC_proto_auth_proto_rawDesc
)

func file_gRPC_proto_auth_proto_rawDescGZIP() []byte {
	file_gRPC_proto_auth_proto_rawDescOnce.Do(func() {
		file_gRPC_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_proto_auth_proto_rawDescData)
	})
	return file_gRPC_proto_auth_proto_rawDescData
}

var file_gRPC_proto_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gRPC_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gRPC_proto_auth_proto_goTypes = []interface{}{
	(ResponseType)(0),      // 0: ResponseType
	(*AuthData)(nil),       // 1: AuthData
	(*Verify)(nil),         // 2: Verify
	(*CreateTokenReq)(nil), // 3: CreateTokenReq
	(*CreateTokenRes)(nil), // 4: CreateTokenRes
	(*VerifyTokenReq)(nil), // 5: VerifyTokenReq
	(*VerifyTokenRes)(nil), // 6: VerifyTokenRes
}
var file_gRPC_proto_auth_proto_depIdxs = []int32{
	0, // 0: Verify.status:type_name -> ResponseType
	1, // 1: Verify.auth:type_name -> AuthData
	1, // 2: CreateTokenReq.auth:type_name -> AuthData
	1, // 3: CreateTokenRes.auth:type_name -> AuthData
	2, // 4: VerifyTokenRes.v:type_name -> Verify
	3, // 5: AuthService.CreateAuth:input_type -> CreateTokenReq
	5, // 6: AuthService.VerifyAuth:input_type -> VerifyTokenReq
	4, // 7: AuthService.CreateAuth:output_type -> CreateTokenRes
	6, // 8: AuthService.VerifyAuth:output_type -> VerifyTokenRes
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_gRPC_proto_auth_proto_init() }
func file_gRPC_proto_auth_proto_init() {
	if File_gRPC_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPC_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
		file_gRPC_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Verify); i {
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
		file_gRPC_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenReq); i {
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
		file_gRPC_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRes); i {
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
		file_gRPC_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenReq); i {
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
		file_gRPC_proto_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenRes); i {
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
			RawDescriptor: file_gRPC_proto_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_proto_auth_proto_goTypes,
		DependencyIndexes: file_gRPC_proto_auth_proto_depIdxs,
		EnumInfos:         file_gRPC_proto_auth_proto_enumTypes,
		MessageInfos:      file_gRPC_proto_auth_proto_msgTypes,
	}.Build()
	File_gRPC_proto_auth_proto = out.File
	file_gRPC_proto_auth_proto_rawDesc = nil
	file_gRPC_proto_auth_proto_goTypes = nil
	file_gRPC_proto_auth_proto_depIdxs = nil
}
