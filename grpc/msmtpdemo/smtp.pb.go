// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: smtp.proto

package msmtpdemo

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

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int64  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage string `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smtp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_smtp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_smtp_proto_rawDescGZIP(), []int{0}
}

func (x *Resp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Resp) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

type SendSmtpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendSmtpRequest) Reset() {
	*x = SendSmtpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smtp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmtpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmtpRequest) ProtoMessage() {}

func (x *SendSmtpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smtp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmtpRequest.ProtoReflect.Descriptor instead.
func (*SendSmtpRequest) Descriptor() ([]byte, []int) {
	return file_smtp_proto_rawDescGZIP(), []int{1}
}

func (x *SendSmtpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SendSmtpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp *Resp `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
}

func (x *SendSmtpResponse) Reset() {
	*x = SendSmtpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smtp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmtpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmtpResponse) ProtoMessage() {}

func (x *SendSmtpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smtp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmtpResponse.ProtoReflect.Descriptor instead.
func (*SendSmtpResponse) Descriptor() ([]byte, []int) {
	return file_smtp_proto_rawDescGZIP(), []int{2}
}

func (x *SendSmtpResponse) GetResp() *Resp {
	if x != nil {
		return x.Resp
	}
	return nil
}

type QueryVerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *QueryVerifyRequest) Reset() {
	*x = QueryVerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smtp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryVerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryVerifyRequest) ProtoMessage() {}

func (x *QueryVerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smtp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryVerifyRequest.ProtoReflect.Descriptor instead.
func (*QueryVerifyRequest) Descriptor() ([]byte, []int) {
	return file_smtp_proto_rawDescGZIP(), []int{3}
}

func (x *QueryVerifyRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type QueryVerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp   *Resp  `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	Verify string `protobuf:"bytes,2,opt,name=verify,proto3" json:"verify,omitempty"`
}

func (x *QueryVerifyResponse) Reset() {
	*x = QueryVerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smtp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryVerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryVerifyResponse) ProtoMessage() {}

func (x *QueryVerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smtp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryVerifyResponse.ProtoReflect.Descriptor instead.
func (*QueryVerifyResponse) Descriptor() ([]byte, []int) {
	return file_smtp_proto_rawDescGZIP(), []int{4}
}

func (x *QueryVerifyResponse) GetResp() *Resp {
	if x != nil {
		return x.Resp
	}
	return nil
}

func (x *QueryVerifyResponse) GetVerify() string {
	if x != nil {
		return x.Verify
	}
	return ""
}

var File_smtp_proto protoreflect.FileDescriptor

var file_smtp_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x6d, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x73,
	0x6d, 0x74, 0x70, 0x22, 0x4c, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x27, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x74, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x33, 0x0a, 0x10, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x6d, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d,
	0x73, 0x6d, 0x74, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x72, 0x65, 0x73, 0x70, 0x22,
	0x2a, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4e, 0x0a, 0x13, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x6d, 0x73, 0x6d, 0x74, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x04, 0x72,
	0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x32, 0x94, 0x01, 0x0a, 0x0b,
	0x53, 0x6d, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x6d, 0x74, 0x70, 0x12, 0x16, 0x2e, 0x6d, 0x73, 0x6d, 0x74, 0x70, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x6d, 0x73, 0x6d, 0x74, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x74, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x19, 0x2e, 0x6d, 0x73, 0x6d, 0x74,
	0x70, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x73, 0x6d, 0x74, 0x70, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x6d, 0x73, 0x6d, 0x74, 0x70, 0x64, 0x65, 0x6d,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_smtp_proto_rawDescOnce sync.Once
	file_smtp_proto_rawDescData = file_smtp_proto_rawDesc
)

func file_smtp_proto_rawDescGZIP() []byte {
	file_smtp_proto_rawDescOnce.Do(func() {
		file_smtp_proto_rawDescData = protoimpl.X.CompressGZIP(file_smtp_proto_rawDescData)
	})
	return file_smtp_proto_rawDescData
}

var file_smtp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_smtp_proto_goTypes = []interface{}{
	(*Resp)(nil),                // 0: msmtp.Resp
	(*SendSmtpRequest)(nil),     // 1: msmtp.SendSmtpRequest
	(*SendSmtpResponse)(nil),    // 2: msmtp.SendSmtpResponse
	(*QueryVerifyRequest)(nil),  // 3: msmtp.QueryVerifyRequest
	(*QueryVerifyResponse)(nil), // 4: msmtp.QueryVerifyResponse
}
var file_smtp_proto_depIdxs = []int32{
	0, // 0: msmtp.SendSmtpResponse.resp:type_name -> msmtp.Resp
	0, // 1: msmtp.QueryVerifyResponse.resp:type_name -> msmtp.Resp
	1, // 2: msmtp.SmtpService.SendSmtp:input_type -> msmtp.SendSmtpRequest
	3, // 3: msmtp.SmtpService.QueryVerify:input_type -> msmtp.QueryVerifyRequest
	2, // 4: msmtp.SmtpService.SendSmtp:output_type -> msmtp.SendSmtpResponse
	4, // 5: msmtp.SmtpService.QueryVerify:output_type -> msmtp.QueryVerifyResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_smtp_proto_init() }
func file_smtp_proto_init() {
	if File_smtp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smtp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
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
		file_smtp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmtpRequest); i {
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
		file_smtp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmtpResponse); i {
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
		file_smtp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryVerifyRequest); i {
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
		file_smtp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryVerifyResponse); i {
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
			RawDescriptor: file_smtp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smtp_proto_goTypes,
		DependencyIndexes: file_smtp_proto_depIdxs,
		MessageInfos:      file_smtp_proto_msgTypes,
	}.Build()
	File_smtp_proto = out.File
	file_smtp_proto_rawDesc = nil
	file_smtp_proto_goTypes = nil
	file_smtp_proto_depIdxs = nil
}
