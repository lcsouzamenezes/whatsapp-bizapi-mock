// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: contacts.proto

package model

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Contact_StatusEnum int32

const (
	Contact_processing Contact_StatusEnum = 0
	Contact_valid      Contact_StatusEnum = 1
	Contact_invalid    Contact_StatusEnum = 2
	Contact_failed     Contact_StatusEnum = 3
)

var Contact_StatusEnum_name = map[int32]string{
	0: "processing",
	1: "valid",
	2: "invalid",
	3: "failed",
}

var Contact_StatusEnum_value = map[string]int32{
	"processing": 0,
	"valid":      1,
	"invalid":    2,
	"failed":     3,
}

func (x Contact_StatusEnum) String() string {
	return proto.EnumName(Contact_StatusEnum_name, int32(x))
}

func (Contact_StatusEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{0, 0}
}

type ContactRequest_BlockingEnum int32

const (
	ContactRequest_wait    ContactRequest_BlockingEnum = 0
	ContactRequest_no_wait ContactRequest_BlockingEnum = 1
)

var ContactRequest_BlockingEnum_name = map[int32]string{
	0: "wait",
	1: "no_wait",
}

var ContactRequest_BlockingEnum_value = map[string]int32{
	"wait":    0,
	"no_wait": 1,
}

func (x ContactRequest_BlockingEnum) String() string {
	return proto.EnumName(ContactRequest_BlockingEnum_name, int32(x))
}

func (ContactRequest_BlockingEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{1, 0}
}

type Contact struct {
	WaId                 string             `protobuf:"bytes,1,opt,name=wa_id,json=waId,proto3" json:"wa_id,omitempty"`
	Input                string             `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	Status               Contact_StatusEnum `protobuf:"varint,3,opt,name=status,proto3,enum=whatsapp.Contact_StatusEnum" json:"status,omitempty"`
	Profile              *Contact_Profile   `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}
func (*Contact) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{0}
}
func (m *Contact) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contact.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact.Merge(m, src)
}
func (m *Contact) XXX_Size() int {
	return m.Size()
}
func (m *Contact) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact.DiscardUnknown(m)
}

var xxx_messageInfo_Contact proto.InternalMessageInfo

func (m *Contact) GetWaId() string {
	if m != nil {
		return m.WaId
	}
	return ""
}

func (m *Contact) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *Contact) GetStatus() Contact_StatusEnum {
	if m != nil {
		return m.Status
	}
	return Contact_processing
}

func (m *Contact) GetProfile() *Contact_Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type Contact_Profile struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contact_Profile) Reset()         { *m = Contact_Profile{} }
func (m *Contact_Profile) String() string { return proto.CompactTextString(m) }
func (*Contact_Profile) ProtoMessage()    {}
func (*Contact_Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{0, 0}
}
func (m *Contact_Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contact_Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contact_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contact_Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contact_Profile.Merge(m, src)
}
func (m *Contact_Profile) XXX_Size() int {
	return m.Size()
}
func (m *Contact_Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Contact_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Contact_Profile proto.InternalMessageInfo

func (m *Contact_Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ContactRequest struct {
	Blocking             ContactRequest_BlockingEnum `protobuf:"varint,1,opt,name=blocking,proto3,enum=whatsapp.ContactRequest_BlockingEnum" json:"blocking,omitempty"`
	Contacts             []string                    `protobuf:"bytes,2,rep,name=contacts,proto3" json:"contacts,omitempty"`
	ForceCheck           bool                        `protobuf:"varint,3,opt,name=force_check,json=forceCheck,proto3" json:"force_check,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ContactRequest) Reset()         { *m = ContactRequest{} }
func (m *ContactRequest) String() string { return proto.CompactTextString(m) }
func (*ContactRequest) ProtoMessage()    {}
func (*ContactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{1}
}
func (m *ContactRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContactRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactRequest.Merge(m, src)
}
func (m *ContactRequest) XXX_Size() int {
	return m.Size()
}
func (m *ContactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactRequest proto.InternalMessageInfo

func (m *ContactRequest) GetBlocking() ContactRequest_BlockingEnum {
	if m != nil {
		return m.Blocking
	}
	return ContactRequest_wait
}

func (m *ContactRequest) GetContacts() []string {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *ContactRequest) GetForceCheck() bool {
	if m != nil {
		return m.ForceCheck
	}
	return false
}

type ContactResponse struct {
	Contacts             []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ContactResponse) Reset()         { *m = ContactResponse{} }
func (m *ContactResponse) String() string { return proto.CompactTextString(m) }
func (*ContactResponse) ProtoMessage()    {}
func (*ContactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{2}
}
func (m *ContactResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContactResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactResponse.Merge(m, src)
}
func (m *ContactResponse) XXX_Size() int {
	return m.Size()
}
func (m *ContactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactResponse proto.InternalMessageInfo

func (m *ContactResponse) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func init() {
	proto.RegisterEnum("whatsapp.Contact_StatusEnum", Contact_StatusEnum_name, Contact_StatusEnum_value)
	proto.RegisterEnum("whatsapp.ContactRequest_BlockingEnum", ContactRequest_BlockingEnum_name, ContactRequest_BlockingEnum_value)
	proto.RegisterType((*Contact)(nil), "whatsapp.Contact")
	proto.RegisterType((*Contact_Profile)(nil), "whatsapp.Contact.Profile")
	proto.RegisterType((*ContactRequest)(nil), "whatsapp.ContactRequest")
	proto.RegisterType((*ContactResponse)(nil), "whatsapp.ContactResponse")
}

func init() { proto.RegisterFile("contacts.proto", fileDescriptor_72e48e1bf84d56b8) }

var fileDescriptor_72e48e1bf84d56b8 = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x1d, 0xf7, 0x99, 0xde, 0xa2, 0x50, 0x2e, 0xb3, 0x08, 0x15, 0x94, 0x28, 0xd2, 0x48, 0xd9,
	0x34, 0x95, 0x3a, 0x7c, 0xc0, 0x3c, 0xc4, 0x82, 0x1d, 0x0a, 0x3b, 0x36, 0x95, 0xeb, 0xb8, 0xad,
	0x35, 0xa9, 0x6d, 0x62, 0x67, 0xca, 0xec, 0xf8, 0x2b, 0x7e, 0x81, 0x25, 0x9f, 0x80, 0xfa, 0x25,
	0xa8, 0x4e, 0x9a, 0x19, 0x29, 0xbb, 0x7b, 0x4e, 0xce, 0x3d, 0xe7, 0xdc, 0x24, 0xe0, 0x33, 0x25,
	0x2d, 0x65, 0xd6, 0x24, 0xba, 0x50, 0x56, 0xa1, 0x77, 0xd8, 0x51, 0x6b, 0xa8, 0xd6, 0xd3, 0xdb,
	0xad, 0xb0, 0xbb, 0x72, 0x9d, 0x30, 0xb5, 0x5f, 0x70, 0xf9, 0xa8, 0x9e, 0x74, 0xa1, 0x7e, 0x3e,
	0x2d, 0x9c, 0x8c, 0xcd, 0xb7, 0x5c, 0xce, 0x1f, 0x69, 0x2e, 0x32, 0x6a, 0xf9, 0xa2, 0x35, 0x54,
	0x66, 0xd1, 0xaf, 0x0e, 0x0c, 0xef, 0x2b, 0x7f, 0x7c, 0x0b, 0xfd, 0x03, 0x5d, 0x89, 0x2c, 0x20,
	0x21, 0x89, 0x47, 0x69, 0xef, 0x40, 0xbf, 0x64, 0x78, 0x09, 0x7d, 0x21, 0x75, 0x69, 0x83, 0x8e,
	0x23, 0x2b, 0x80, 0x9f, 0x60, 0x60, 0x2c, 0xb5, 0xa5, 0x09, 0xba, 0x21, 0x89, 0xfd, 0xe5, 0xfb,
	0xe4, 0x5c, 0x2a, 0xa9, 0xdd, 0x92, 0x6f, 0xee, 0xf9, 0x67, 0x59, 0xee, 0xd3, 0x5a, 0x8b, 0xd7,
	0x30, 0xd4, 0x85, 0xda, 0x88, 0x9c, 0x07, 0xbd, 0x90, 0xc4, 0xe3, 0xe5, 0xbb, 0xf6, 0xda, 0xd7,
	0x4a, 0x90, 0x9e, 0x95, 0xd3, 0x0f, 0x30, 0xac, 0x39, 0x44, 0xe8, 0x49, 0xba, 0xe7, 0xe7, 0x7e,
	0xa7, 0x39, 0xba, 0x01, 0x78, 0x4e, 0x42, 0x1f, 0x40, 0x17, 0x8a, 0x71, 0x63, 0x84, 0xdc, 0x4e,
	0x2e, 0x70, 0x04, 0x7d, 0x77, 0xf0, 0x84, 0xe0, 0x18, 0x86, 0x42, 0x56, 0xa0, 0x83, 0x00, 0x83,
	0x0d, 0x15, 0x39, 0xcf, 0x26, 0xdd, 0xe8, 0x37, 0x01, 0xbf, 0x4e, 0x4f, 0xf9, 0x8f, 0x92, 0x1b,
	0x8b, 0xb7, 0xe0, 0xad, 0x73, 0xc5, 0x1e, 0x84, 0xdc, 0xba, 0x30, 0x7f, 0x79, 0xd5, 0x6a, 0x5a,
	0x6b, 0x93, 0xbb, 0x5a, 0xe8, 0x2e, 0x6d, 0xd6, 0x70, 0x0a, 0xde, 0xf9, 0xbb, 0x05, 0x9d, 0xb0,
	0x1b, 0x8f, 0xd2, 0x06, 0xe3, 0x47, 0x18, 0x6f, 0x54, 0xc1, 0xf8, 0x8a, 0xed, 0x38, 0x7b, 0x70,
	0xaf, 0xd0, 0x4b, 0xc1, 0x51, 0xf7, 0x27, 0x26, 0xba, 0x82, 0x57, 0x2f, 0x6d, 0xd1, 0x83, 0xde,
	0x81, 0x0a, 0x3b, 0xb9, 0x38, 0x5d, 0x21, 0xd5, 0xca, 0x01, 0x12, 0xdd, 0xc0, 0xeb, 0xa6, 0x8c,
	0xd1, 0x4a, 0x1a, 0x8e, 0xf3, 0x17, 0xb1, 0x24, 0xec, 0xc6, 0xe3, 0xe5, 0x9b, 0x76, 0xf3, 0x46,
	0x72, 0x77, 0xf9, 0xe7, 0x38, 0x23, 0x7f, 0x8f, 0x33, 0xf2, 0xef, 0x38, 0x23, 0xdf, 0x07, 0x8b,
	0xbd, 0xca, 0x78, 0xbe, 0x1e, 0xb8, 0x7f, 0xe3, 0xfa, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb,
	0xc2, 0xb3, 0xc2, 0x7a, 0x02, 0x00, 0x00,
}

func (m *Contact) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contact) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contact) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Profile != nil {
		{
			size, err := m.Profile.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintContacts(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Status != 0 {
		i = encodeVarintContacts(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Input) > 0 {
		i -= len(m.Input)
		copy(dAtA[i:], m.Input)
		i = encodeVarintContacts(dAtA, i, uint64(len(m.Input)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.WaId) > 0 {
		i -= len(m.WaId)
		copy(dAtA[i:], m.WaId)
		i = encodeVarintContacts(dAtA, i, uint64(len(m.WaId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Contact_Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contact_Profile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contact_Profile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintContacts(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContactRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContactRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ForceCheck {
		i--
		if m.ForceCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Contacts) > 0 {
		for iNdEx := len(m.Contacts) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contacts[iNdEx])
			copy(dAtA[i:], m.Contacts[iNdEx])
			i = encodeVarintContacts(dAtA, i, uint64(len(m.Contacts[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Blocking != 0 {
		i = encodeVarintContacts(dAtA, i, uint64(m.Blocking))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ContactResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContactResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Contacts) > 0 {
		for iNdEx := len(m.Contacts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Contacts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintContacts(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintContacts(dAtA []byte, offset int, v uint64) int {
	offset -= sovContacts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Contact) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WaId)
	if l > 0 {
		n += 1 + l + sovContacts(uint64(l))
	}
	l = len(m.Input)
	if l > 0 {
		n += 1 + l + sovContacts(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovContacts(uint64(m.Status))
	}
	if m.Profile != nil {
		l = m.Profile.Size()
		n += 1 + l + sovContacts(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Contact_Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovContacts(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContactRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Blocking != 0 {
		n += 1 + sovContacts(uint64(m.Blocking))
	}
	if len(m.Contacts) > 0 {
		for _, s := range m.Contacts {
			l = len(s)
			n += 1 + l + sovContacts(uint64(l))
		}
	}
	if m.ForceCheck {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ContactResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Contacts) > 0 {
		for _, e := range m.Contacts {
			l = e.Size()
			n += 1 + l + sovContacts(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovContacts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContacts(x uint64) (n int) {
	return sovContacts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contact) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContacts
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Contact: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contact: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WaId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContacts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContacts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WaId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Input", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContacts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContacts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Input = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Contact_StatusEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContacts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContacts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Profile == nil {
				m.Profile = &Contact_Profile{}
			}
			if err := m.Profile.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContacts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContacts
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Contact_Profile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContacts
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContacts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContacts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContacts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContacts
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContactRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContacts
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContactRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocking", wireType)
			}
			m.Blocking = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Blocking |= ContactRequest_BlockingEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contacts", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContacts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContacts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contacts = append(m.Contacts, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForceCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ForceCheck = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipContacts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContacts
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContactResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContacts
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContactResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contacts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContacts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContacts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contacts = append(m.Contacts, &Contact{})
			if err := m.Contacts[len(m.Contacts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContacts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContacts
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipContacts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContacts
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowContacts
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthContacts
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContacts
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContacts
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContacts        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContacts          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContacts = fmt.Errorf("proto: unexpected end of group")
)
