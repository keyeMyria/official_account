// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

package igrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SaleOrderTemplateMessage struct {
	SaleOrderId     int64 `protobuf:"varint,1,opt,name=SaleOrderId" json:"SaleOrderId,omitempty"`
	SaleOrderItemId int64 `protobuf:"varint,2,opt,name=SaleOrderItemId" json:"SaleOrderItemId,omitempty"`
	Type            int64 `protobuf:"varint,3,opt,name=Type" json:"Type,omitempty"`
}

func (m *SaleOrderTemplateMessage) Reset()                    { *m = SaleOrderTemplateMessage{} }
func (m *SaleOrderTemplateMessage) String() string            { return proto.CompactTextString(m) }
func (*SaleOrderTemplateMessage) ProtoMessage()               {}
func (*SaleOrderTemplateMessage) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *SaleOrderTemplateMessage) GetSaleOrderId() int64 {
	if m != nil {
		return m.SaleOrderId
	}
	return 0
}

func (m *SaleOrderTemplateMessage) GetSaleOrderItemId() int64 {
	if m != nil {
		return m.SaleOrderItemId
	}
	return 0
}

func (m *SaleOrderTemplateMessage) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

type SmsMessage struct {
	Mobile  string `protobuf:"bytes,1,opt,name=Mobile" json:"Mobile,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content" json:"Content,omitempty"`
}

func (m *SmsMessage) Reset()                    { *m = SmsMessage{} }
func (m *SmsMessage) String() string            { return proto.CompactTextString(m) }
func (*SmsMessage) ProtoMessage()               {}
func (*SmsMessage) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *SmsMessage) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SmsMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type EmailMessage struct {
	Email    string `protobuf:"bytes,1,opt,name=Email" json:"Email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=Title" json:"Title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=Content" json:"Content,omitempty"`
}

func (m *EmailMessage) Reset()                    { *m = EmailMessage{} }
func (m *EmailMessage) String() string            { return proto.CompactTextString(m) }
func (*EmailMessage) ProtoMessage()               {}
func (*EmailMessage) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *EmailMessage) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailMessage) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *EmailMessage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EmailMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Message struct {
	Bts   []byte `protobuf:"bytes,1,opt,name=Bts,proto3" json:"Bts,omitempty"`
	Topic string `protobuf:"bytes,2,opt,name=Topic" json:"Topic,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *Message) GetBts() []byte {
	if m != nil {
		return m.Bts
	}
	return nil
}

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*SaleOrderTemplateMessage)(nil), "igrpc.SaleOrderTemplateMessage")
	proto.RegisterType((*SmsMessage)(nil), "igrpc.SmsMessage")
	proto.RegisterType((*EmailMessage)(nil), "igrpc.EmailMessage")
	proto.RegisterType((*Message)(nil), "igrpc.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0xa9, 0xdb, 0x3f, 0xee, 0x58, 0x51, 0x82, 0x48, 0xf0, 0x54, 0xf6, 0xd4, 0x93, 0x20,
	0xde, 0x3d, 0x28, 0x1e, 0x7a, 0x28, 0x4a, 0xba, 0x2f, 0x90, 0x36, 0x43, 0x09, 0x24, 0x4d, 0x48,
	0x02, 0xa2, 0x4f, 0x2f, 0x9d, 0x24, 0x75, 0xed, 0x6d, 0xbe, 0x99, 0xc9, 0xef, 0x0b, 0x03, 0xd7,
	0x16, 0x63, 0x94, 0x7b, 0x7c, 0xf4, 0xc1, 0x25, 0xc7, 0x26, 0x7a, 0x1f, 0xfc, 0xae, 0xfb, 0x01,
	0xbe, 0x91, 0x06, 0x3f, 0x82, 0xc2, 0xd0, 0xa3, 0xf5, 0x46, 0x26, 0x5c, 0xe7, 0x45, 0xb6, 0x80,
	0xab, 0xd3, 0x6c, 0xa5, 0xf8, 0x68, 0x31, 0x5a, 0x36, 0x62, 0xd8, 0x62, 0x4b, 0xb8, 0xf9, 0xc3,
	0x84, 0x76, 0xa5, 0xf8, 0x05, 0x6d, 0x9d, 0xb7, 0x19, 0x83, 0x71, 0xff, 0xed, 0x91, 0x37, 0x34,
	0xa6, 0xba, 0x7b, 0x01, 0xd8, 0xd8, 0x58, 0x6d, 0xf7, 0x30, 0x5d, 0xbb, 0xad, 0x36, 0x48, 0xa2,
	0x56, 0x14, 0x62, 0x1c, 0x66, 0x6f, 0xee, 0x90, 0xf0, 0x90, 0x28, 0xbb, 0x15, 0x15, 0x3b, 0x0f,
	0xf3, 0x77, 0x2b, 0xb5, 0xa9, 0x09, 0x77, 0x30, 0x21, 0x2e, 0x01, 0x19, 0xd8, 0x03, 0x5c, 0x7e,
	0xca, 0x18, 0xbf, 0x5c, 0x50, 0x25, 0xe0, 0xc4, 0xc7, 0x17, 0xbd, 0x4e, 0x26, 0x7f, 0xab, 0x15,
	0x19, 0x86, 0xc6, 0xf1, 0x7f, 0xe3, 0x13, 0xcc, 0xaa, 0xec, 0x16, 0x9a, 0xd7, 0x14, 0x49, 0x35,
	0x17, 0xc7, 0x92, 0xc2, 0x9c, 0xd7, 0xbb, 0x62, 0xc9, 0xb0, 0x9d, 0xd2, 0xb9, 0x9f, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xe0, 0xcc, 0x63, 0x9f, 0x7f, 0x01, 0x00, 0x00,
}
