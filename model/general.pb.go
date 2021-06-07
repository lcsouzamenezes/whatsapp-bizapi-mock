// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: general.proto

package model

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Details string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Href    string `protobuf:"bytes,4,opt,name=href,proto3" json:"href,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_general_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_general_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Error) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *Error) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_general_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_general_proto_rawDescGZIP(), []int{1}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Id `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Media    []*Id `protobuf:"bytes,2,rep,name=media,proto3" json:"media,omitempty"`
}

func (x *IdResponse) Reset() {
	*x = IdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdResponse) ProtoMessage() {}

func (x *IdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_general_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdResponse.ProtoReflect.Descriptor instead.
func (*IdResponse) Descriptor() ([]byte, []int) {
	return file_general_proto_rawDescGZIP(), []int{2}
}

func (x *IdResponse) GetMessages() []*Id {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *IdResponse) GetMedia() []*Id {
	if x != nil {
		return x.Media
	}
	return nil
}

type MetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *MetaResponse) Reset() {
	*x = MetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaResponse) ProtoMessage() {}

func (x *MetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_general_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaResponse.ProtoReflect.Descriptor instead.
func (*MetaResponse) Descriptor() ([]byte, []int) {
	return file_general_proto_rawDescGZIP(), []int{3}
}

func (x *MetaResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *Meta    `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Errors []*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_general_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_general_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_general_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ErrorResponse) GetErrors() []*Error {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_general_proto protoreflect.FileDescriptor

var file_general_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5a, 0x0a, 0x0a, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70,
	0x70, 0x2e, 0x49, 0x64, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x2e, 0x0a, 0x0c, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x58, 0x0a, 0x0d, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77,
	0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_general_proto_rawDescOnce sync.Once
	file_general_proto_rawDescData = file_general_proto_rawDesc
)

func file_general_proto_rawDescGZIP() []byte {
	file_general_proto_rawDescOnce.Do(func() {
		file_general_proto_rawDescData = protoimpl.X.CompressGZIP(file_general_proto_rawDescData)
	})
	return file_general_proto_rawDescData
}

var file_general_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_general_proto_goTypes = []interface{}{
	(*Error)(nil),         // 0: whatsapp.Error
	(*Id)(nil),            // 1: whatsapp.Id
	(*IdResponse)(nil),    // 2: whatsapp.IdResponse
	(*MetaResponse)(nil),  // 3: whatsapp.MetaResponse
	(*ErrorResponse)(nil), // 4: whatsapp.ErrorResponse
	(*Meta)(nil),          // 5: meta.Meta
}
var file_general_proto_depIdxs = []int32{
	1, // 0: whatsapp.IdResponse.messages:type_name -> whatsapp.Id
	1, // 1: whatsapp.IdResponse.media:type_name -> whatsapp.Id
	5, // 2: whatsapp.MetaResponse.meta:type_name -> meta.Meta
	5, // 3: whatsapp.ErrorResponse.meta:type_name -> meta.Meta
	0, // 4: whatsapp.ErrorResponse.errors:type_name -> whatsapp.Error
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_general_proto_init() }
func file_general_proto_init() {
	if File_general_proto != nil {
		return
	}
	file_meta_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_general_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_general_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_general_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdResponse); i {
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
		file_general_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaResponse); i {
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
		file_general_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
			RawDescriptor: file_general_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_general_proto_goTypes,
		DependencyIndexes: file_general_proto_depIdxs,
		MessageInfos:      file_general_proto_msgTypes,
	}.Build()
	File_general_proto = out.File
	file_general_proto_rawDesc = nil
	file_general_proto_goTypes = nil
	file_general_proto_depIdxs = nil
}