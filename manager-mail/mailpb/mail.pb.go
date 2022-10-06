// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: mailpb/mail.proto

package mailpb

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

type Mail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject  string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Sender   string `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Mail) Reset() {
	*x = Mail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailpb_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mail) ProtoMessage() {}

func (x *Mail) ProtoReflect() protoreflect.Message {
	mi := &file_mailpb_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mail.ProtoReflect.Descriptor instead.
func (*Mail) Descriptor() ([]byte, []int) {
	return file_mailpb_mail_proto_rawDescGZIP(), []int{0}
}

func (x *Mail) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Mail) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Mail) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *Mail) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreateMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail *Mail `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
}

func (x *CreateMailRequest) Reset() {
	*x = CreateMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailpb_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMailRequest) ProtoMessage() {}

func (x *CreateMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mailpb_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMailRequest.ProtoReflect.Descriptor instead.
func (*CreateMailRequest) Descriptor() ([]byte, []int) {
	return file_mailpb_mail_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMailRequest) GetMail() *Mail {
	if x != nil {
		return x.Mail
	}
	return nil
}

type CreateMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail *Mail `protobuf:"bytes,1,opt,name=mail,proto3" json:"mail,omitempty"`
}

func (x *CreateMailResponse) Reset() {
	*x = CreateMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailpb_mail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMailResponse) ProtoMessage() {}

func (x *CreateMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mailpb_mail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMailResponse.ProtoReflect.Descriptor instead.
func (*CreateMailResponse) Descriptor() ([]byte, []int) {
	return file_mailpb_mail_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMailResponse) GetMail() *Mail {
	if x != nil {
		return x.Mail
	}
	return nil
}

type GetAllMailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllMailRequest) Reset() {
	*x = GetAllMailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailpb_mail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMailRequest) ProtoMessage() {}

func (x *GetAllMailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mailpb_mail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMailRequest.ProtoReflect.Descriptor instead.
func (*GetAllMailRequest) Descriptor() ([]byte, []int) {
	return file_mailpb_mail_proto_rawDescGZIP(), []int{3}
}

type GetAllMailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mail []*Mail `protobuf:"bytes,1,rep,name=mail,proto3" json:"mail,omitempty"`
}

func (x *GetAllMailResponse) Reset() {
	*x = GetAllMailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailpb_mail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMailResponse) ProtoMessage() {}

func (x *GetAllMailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mailpb_mail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMailResponse.ProtoReflect.Descriptor instead.
func (*GetAllMailResponse) Descriptor() ([]byte, []int) {
	return file_mailpb_mail_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllMailResponse) GetMail() []*Mail {
	if x != nil {
		return x.Mail
	}
	return nil
}

var File_mailpb_mail_proto protoreflect.FileDescriptor

var file_mailpb_mail_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x6e, 0x0a, 0x04, 0x4d, 0x61, 0x69,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d,
	0x61, 0x69, 0x6c, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x34,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x04,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x6d, 0x61, 0x69, 0x6c, 0x32,
	0x94, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x69,
	0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x6d, 0x61, 0x69, 0x6c, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mailpb_mail_proto_rawDescOnce sync.Once
	file_mailpb_mail_proto_rawDescData = file_mailpb_mail_proto_rawDesc
)

func file_mailpb_mail_proto_rawDescGZIP() []byte {
	file_mailpb_mail_proto_rawDescOnce.Do(func() {
		file_mailpb_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mailpb_mail_proto_rawDescData)
	})
	return file_mailpb_mail_proto_rawDescData
}

var file_mailpb_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mailpb_mail_proto_goTypes = []interface{}{
	(*Mail)(nil),               // 0: mail.Mail
	(*CreateMailRequest)(nil),  // 1: mail.CreateMailRequest
	(*CreateMailResponse)(nil), // 2: mail.CreateMailResponse
	(*GetAllMailRequest)(nil),  // 3: mail.GetAllMailRequest
	(*GetAllMailResponse)(nil), // 4: mail.GetAllMailResponse
}
var file_mailpb_mail_proto_depIdxs = []int32{
	0, // 0: mail.CreateMailRequest.mail:type_name -> mail.Mail
	0, // 1: mail.CreateMailResponse.mail:type_name -> mail.Mail
	0, // 2: mail.GetAllMailResponse.mail:type_name -> mail.Mail
	1, // 3: mail.MailService.CreateMail:input_type -> mail.CreateMailRequest
	3, // 4: mail.MailService.GetAllMails:input_type -> mail.GetAllMailRequest
	2, // 5: mail.MailService.CreateMail:output_type -> mail.CreateMailResponse
	4, // 6: mail.MailService.GetAllMails:output_type -> mail.GetAllMailResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mailpb_mail_proto_init() }
func file_mailpb_mail_proto_init() {
	if File_mailpb_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mailpb_mail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mail); i {
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
		file_mailpb_mail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMailRequest); i {
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
		file_mailpb_mail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMailResponse); i {
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
		file_mailpb_mail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMailRequest); i {
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
		file_mailpb_mail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMailResponse); i {
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
			RawDescriptor: file_mailpb_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mailpb_mail_proto_goTypes,
		DependencyIndexes: file_mailpb_mail_proto_depIdxs,
		MessageInfos:      file_mailpb_mail_proto_msgTypes,
	}.Build()
	File_mailpb_mail_proto = out.File
	file_mailpb_mail_proto_rawDesc = nil
	file_mailpb_mail_proto_goTypes = nil
	file_mailpb_mail_proto_depIdxs = nil
}