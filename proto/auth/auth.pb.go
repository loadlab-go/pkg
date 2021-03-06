// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: proto/auth/auth.proto

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

type GenerateJWTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GenerateJWTRequest) Reset() {
	*x = GenerateJWTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateJWTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateJWTRequest) ProtoMessage() {}

func (x *GenerateJWTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateJWTRequest.ProtoReflect.Descriptor instead.
func (*GenerateJWTRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateJWTRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GenerateJWTRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GenerateJWTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GenerateJWTResponse) Reset() {
	*x = GenerateJWTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateJWTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateJWTResponse) ProtoMessage() {}

func (x *GenerateJWTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateJWTResponse.ProtoReflect.Descriptor instead.
func (*GenerateJWTResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateJWTResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateJWTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ValidateJWTRequest) Reset() {
	*x = ValidateJWTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateJWTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateJWTRequest) ProtoMessage() {}

func (x *ValidateJWTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateJWTRequest.ProtoReflect.Descriptor instead.
func (*ValidateJWTRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateJWTRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateJWTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aud []string `protobuf:"bytes,1,rep,name=aud,proto3" json:"aud,omitempty"`
	Exp int64    `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
	Jti string   `protobuf:"bytes,3,opt,name=jti,proto3" json:"jti,omitempty"`
	Iat int64    `protobuf:"varint,4,opt,name=iat,proto3" json:"iat,omitempty"`
	Iss string   `protobuf:"bytes,5,opt,name=iss,proto3" json:"iss,omitempty"`
	Nbf int64    `protobuf:"varint,6,opt,name=nbf,proto3" json:"nbf,omitempty"`
	Sub string   `protobuf:"bytes,7,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *ValidateJWTResponse) Reset() {
	*x = ValidateJWTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateJWTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateJWTResponse) ProtoMessage() {}

func (x *ValidateJWTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateJWTResponse.ProtoReflect.Descriptor instead.
func (*ValidateJWTResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateJWTResponse) GetAud() []string {
	if x != nil {
		return x.Aud
	}
	return nil
}

func (x *ValidateJWTResponse) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *ValidateJWTResponse) GetJti() string {
	if x != nil {
		return x.Jti
	}
	return ""
}

func (x *ValidateJWTResponse) GetIat() int64 {
	if x != nil {
		return x.Iat
	}
	return 0
}

func (x *ValidateJWTResponse) GetIss() string {
	if x != nil {
		return x.Iss
	}
	return ""
}

func (x *ValidateJWTResponse) GetNbf() int64 {
	if x != nil {
		return x.Nbf
	}
	return 0
}

func (x *ValidateJWTResponse) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

var File_proto_auth_auth_proto protoreflect.FileDescriptor

var file_proto_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x40, 0x0a,
	0x12, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x2b, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x12,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x13, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x75, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x61,
	0x75, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x65, 0x78, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x74, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6a, 0x74, 0x69, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x62,
	0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6e, 0x62, 0x66, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x75, 0x62, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x75, 0x62, 0x32, 0x91,
	0x01, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4a, 0x57, 0x54, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4a,
	0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x12, 0x18, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x6c, 0x61, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_auth_auth_proto_rawDescOnce sync.Once
	file_proto_auth_auth_proto_rawDescData = file_proto_auth_auth_proto_rawDesc
)

func file_proto_auth_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_auth_auth_proto_rawDescData)
	})
	return file_proto_auth_auth_proto_rawDescData
}

var file_proto_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_auth_auth_proto_goTypes = []interface{}{
	(*GenerateJWTRequest)(nil),  // 0: auth.GenerateJWTRequest
	(*GenerateJWTResponse)(nil), // 1: auth.GenerateJWTResponse
	(*ValidateJWTRequest)(nil),  // 2: auth.ValidateJWTRequest
	(*ValidateJWTResponse)(nil), // 3: auth.ValidateJWTResponse
}
var file_proto_auth_auth_proto_depIdxs = []int32{
	0, // 0: auth.JWT.GenerateJWT:input_type -> auth.GenerateJWTRequest
	2, // 1: auth.JWT.ValidateJWT:input_type -> auth.ValidateJWTRequest
	1, // 2: auth.JWT.GenerateJWT:output_type -> auth.GenerateJWTResponse
	3, // 3: auth.JWT.ValidateJWT:output_type -> auth.ValidateJWTResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_auth_auth_proto_init() }
func file_proto_auth_auth_proto_init() {
	if File_proto_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateJWTRequest); i {
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
		file_proto_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateJWTResponse); i {
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
		file_proto_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateJWTRequest); i {
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
		file_proto_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateJWTResponse); i {
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
			RawDescriptor: file_proto_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_auth_proto = out.File
	file_proto_auth_auth_proto_rawDesc = nil
	file_proto_auth_auth_proto_goTypes = nil
	file_proto_auth_auth_proto_depIdxs = nil
}
