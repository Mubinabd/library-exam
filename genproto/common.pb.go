// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: protos/common.proto

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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{0}
}

type ById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ById) Reset() {
	*x = ById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ById) ProtoMessage() {}

func (x *ById) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ById.ProtoReflect.Descriptor instead.
func (*ById) Descriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{1}
}

func (x *ById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *ByName) Reset() {
	*x = ByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByName) ProtoMessage() {}

func (x *ByName) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByName.ProtoReflect.Descriptor instead.
func (*ByName) Descriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{2}
}

func (x *ByName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NameFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *NameFilter) Reset() {
	*x = NameFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameFilter) ProtoMessage() {}

func (x *NameFilter) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameFilter.ProtoReflect.Descriptor instead.
func (*NameFilter) Descriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{3}
}

func (x *NameFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_protos_common_proto protoreflect.FileDescriptor

var file_protos_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x06, 0x0a,
	0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x04, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x1c, 0x0a,
	0x06, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x4e,
	0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x5a,
	0x09, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_common_proto_rawDescOnce sync.Once
	file_protos_common_proto_rawDescData = file_protos_common_proto_rawDesc
)

func file_protos_common_proto_rawDescGZIP() []byte {
	file_protos_common_proto_rawDescOnce.Do(func() {
		file_protos_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_proto_rawDescData)
	})
	return file_protos_common_proto_rawDescData
}

var file_protos_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_common_proto_goTypes = []interface{}{
	(*Void)(nil),       // 0: protos.Void
	(*ById)(nil),       // 1: protos.ById
	(*ByName)(nil),     // 2: protos.ByName
	(*NameFilter)(nil), // 3: protos.NameFilter
}
var file_protos_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_common_proto_init() }
func file_protos_common_proto_init() {
	if File_protos_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_protos_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ById); i {
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
		file_protos_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByName); i {
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
		file_protos_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameFilter); i {
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
			RawDescriptor: file_protos_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_common_proto_goTypes,
		DependencyIndexes: file_protos_common_proto_depIdxs,
		MessageInfos:      file_protos_common_proto_msgTypes,
	}.Build()
	File_protos_common_proto = out.File
	file_protos_common_proto_rawDesc = nil
	file_protos_common_proto_goTypes = nil
	file_protos_common_proto_depIdxs = nil
}
