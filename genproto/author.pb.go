// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/author.proto

package genproto

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

type AuthorCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Biography string `protobuf:"bytes,3,opt,name=Biography,proto3" json:"Biography,omitempty"`
}

func (x *AuthorCreate) Reset() {
	*x = AuthorCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorCreate) ProtoMessage() {}

func (x *AuthorCreate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorCreate.ProtoReflect.Descriptor instead.
func (*AuthorCreate) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorCreate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthorCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthorCreate) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Biography string `protobuf:"bytes,3,opt,name=Biography,proto3" json:"Biography,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{1}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

type Authors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*Author `protobuf:"bytes,1,rep,name=Authors,proto3" json:"Authors,omitempty"`
}

func (x *Authors) Reset() {
	*x = Authors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authors) ProtoMessage() {}

func (x *Authors) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authors.ProtoReflect.Descriptor instead.
func (*Authors) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{2}
}

func (x *Authors) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type UserBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=Books,proto3" json:"Books,omitempty"`
}

func (x *UserBook) Reset() {
	*x = UserBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBook) ProtoMessage() {}

func (x *UserBook) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBook.ProtoReflect.Descriptor instead.
func (*UserBook) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{3}
}

func (x *UserBook) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type AuthorID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId string `protobuf:"bytes,1,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
}

func (x *AuthorID) Reset() {
	*x = AuthorID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_author_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorID) ProtoMessage() {}

func (x *AuthorID) ProtoReflect() protoreflect.Message {
	mi := &file_protos_author_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorID.ProtoReflect.Descriptor instead.
func (*AuthorID) Descriptor() ([]byte, []int) {
	return file_protos_author_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorID) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

var File_protos_author_proto protoreflect.FileDescriptor

var file_protos_author_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x13, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x69, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x69,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x22, 0x4a, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x69, 0x6f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x79, 0x22, 0x33, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x28,
	0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52,
	0x07, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x22, 0x2e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x22, 0x0a, 0x05, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x05, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x26, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x32, 0xc6, 0x02, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x44, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_author_proto_rawDescOnce sync.Once
	file_protos_author_proto_rawDescData = file_protos_author_proto_rawDesc
)

func file_protos_author_proto_rawDescGZIP() []byte {
	file_protos_author_proto_rawDescOnce.Do(func() {
		file_protos_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_author_proto_rawDescData)
	})
	return file_protos_author_proto_rawDescData
}

var file_protos_author_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_author_proto_goTypes = []interface{}{
	(*AuthorCreate)(nil), // 0: protos.AuthorCreate
	(*Author)(nil),       // 1: protos.Author
	(*Authors)(nil),      // 2: protos.Authors
	(*UserBook)(nil),     // 3: protos.UserBook
	(*AuthorID)(nil),     // 4: protos.AuthorID
	(*Book)(nil),         // 5: protos.Book
	(*ById)(nil),         // 6: protos.ById
	(*NameFilter)(nil),   // 7: protos.NameFilter
	(*Void)(nil),         // 8: protos.Void
}
var file_protos_author_proto_depIdxs = []int32{
	1, // 0: protos.Authors.Authors:type_name -> protos.Author
	5, // 1: protos.UserBook.Books:type_name -> protos.Book
	0, // 2: protos.AuthorService.CreateAuthor:input_type -> protos.AuthorCreate
	6, // 3: protos.AuthorService.GetAuthor:input_type -> protos.ById
	0, // 4: protos.AuthorService.UpdateAuthor:input_type -> protos.AuthorCreate
	6, // 5: protos.AuthorService.DeleteAuthor:input_type -> protos.ById
	7, // 6: protos.AuthorService.GetAllAuthors:input_type -> protos.NameFilter
	4, // 7: protos.AuthorService.GetAuthorBooks:input_type -> protos.AuthorID
	1, // 8: protos.AuthorService.CreateAuthor:output_type -> protos.Author
	1, // 9: protos.AuthorService.GetAuthor:output_type -> protos.Author
	8, // 10: protos.AuthorService.UpdateAuthor:output_type -> protos.Void
	8, // 11: protos.AuthorService.DeleteAuthor:output_type -> protos.Void
	2, // 12: protos.AuthorService.GetAllAuthors:output_type -> protos.Authors
	3, // 13: protos.AuthorService.GetAuthorBooks:output_type -> protos.UserBook
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_author_proto_init() }
func file_protos_author_proto_init() {
	if File_protos_author_proto != nil {
		return
	}
	file_protos_common_proto_init()
	file_protos_book_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_author_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorCreate); i {
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
		file_protos_author_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
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
		file_protos_author_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authors); i {
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
		file_protos_author_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBook); i {
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
		file_protos_author_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorID); i {
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
			RawDescriptor: file_protos_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_author_proto_goTypes,
		DependencyIndexes: file_protos_author_proto_depIdxs,
		MessageInfos:      file_protos_author_proto_msgTypes,
	}.Build()
	File_protos_author_proto = out.File
	file_protos_author_proto_rawDesc = nil
	file_protos_author_proto_goTypes = nil
	file_protos_author_proto_depIdxs = nil
}
