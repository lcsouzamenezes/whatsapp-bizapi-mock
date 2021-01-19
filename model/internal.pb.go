// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: protobuf/internal.proto

package model

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

type InternalContact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *InternalContact) Reset() {
	*x = InternalContact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalContact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalContact) ProtoMessage() {}

func (x *InternalContact) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalContact.ProtoReflect.Descriptor instead.
func (*InternalContact) Descriptor() ([]byte, []int) {
	return file_protobuf_internal_proto_rawDescGZIP(), []int{0}
}

func (x *InternalContact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InternalContact) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type InternalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      string             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Status       string             `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Contacts     []*InternalContact `protobuf:"bytes,3,rep,name=contacts,proto3" json:"contacts,omitempty"`
	UploadDir    string             `protobuf:"bytes,4,opt,name=uploadDir,proto3" json:"uploadDir,omitempty"`
	WebhookUrl   string             `protobuf:"bytes,5,opt,name=webhookUrl,proto3" json:"webhookUrl,omitempty"`
	Users        map[string]string  `protobuf:"bytes,6,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InboundMedia map[string]string  `protobuf:"bytes,7,rep,name=inboundMedia,proto3" json:"inboundMedia,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InternalConfig) Reset() {
	*x = InternalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalConfig) ProtoMessage() {}

func (x *InternalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalConfig.ProtoReflect.Descriptor instead.
func (*InternalConfig) Descriptor() ([]byte, []int) {
	return file_protobuf_internal_proto_rawDescGZIP(), []int{1}
}

func (x *InternalConfig) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *InternalConfig) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InternalConfig) GetContacts() []*InternalContact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

func (x *InternalConfig) GetUploadDir() string {
	if x != nil {
		return x.UploadDir
	}
	return ""
}

func (x *InternalConfig) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

func (x *InternalConfig) GetUsers() map[string]string {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *InternalConfig) GetInboundMedia() map[string]string {
	if x != nil {
		return x.InboundMedia
	}
	return nil
}

var File_protobuf_internal_proto protoreflect.FileDescriptor

var file_protobuf_internal_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0x35, 0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xb4, 0x03, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x08,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x36,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x0c, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x1a, 0x38, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a,
	0x11, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2c,
	0x0a, 0x18, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6b, 0x6f, 0x6d, 0x2e, 0x77, 0x68, 0x61,
	0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x6d, 0x6f, 0x63, 0x6b, 0x42, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5a, 0x07, 0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_internal_proto_rawDescOnce sync.Once
	file_protobuf_internal_proto_rawDescData = file_protobuf_internal_proto_rawDesc
)

func file_protobuf_internal_proto_rawDescGZIP() []byte {
	file_protobuf_internal_proto_rawDescOnce.Do(func() {
		file_protobuf_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_internal_proto_rawDescData)
	})
	return file_protobuf_internal_proto_rawDescData
}

var file_protobuf_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuf_internal_proto_goTypes = []interface{}{
	(*InternalContact)(nil), // 0: model.InternalContact
	(*InternalConfig)(nil),  // 1: model.InternalConfig
	nil,                     // 2: model.InternalConfig.UsersEntry
	nil,                     // 3: model.InternalConfig.InboundMediaEntry
}
var file_protobuf_internal_proto_depIdxs = []int32{
	0, // 0: model.InternalConfig.contacts:type_name -> model.InternalContact
	2, // 1: model.InternalConfig.users:type_name -> model.InternalConfig.UsersEntry
	3, // 2: model.InternalConfig.inboundMedia:type_name -> model.InternalConfig.InboundMediaEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protobuf_internal_proto_init() }
func file_protobuf_internal_proto_init() {
	if File_protobuf_internal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalContact); i {
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
		file_protobuf_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalConfig); i {
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
			RawDescriptor: file_protobuf_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_internal_proto_goTypes,
		DependencyIndexes: file_protobuf_internal_proto_depIdxs,
		MessageInfos:      file_protobuf_internal_proto_msgTypes,
	}.Build()
	File_protobuf_internal_proto = out.File
	file_protobuf_internal_proto_rawDesc = nil
	file_protobuf_internal_proto_goTypes = nil
	file_protobuf_internal_proto_depIdxs = nil
}
