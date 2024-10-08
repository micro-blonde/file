// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.21.3
// source: template/properties.proto

package file

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

type PropertiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PropertiesRequest) Reset() {
	*x = PropertiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_properties_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertiesRequest) ProtoMessage() {}

func (x *PropertiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_properties_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertiesRequest.ProtoReflect.Descriptor instead.
func (*PropertiesRequest) Descriptor() ([]byte, []int) {
	return file_template_properties_proto_rawDescGZIP(), []int{0}
}

type Properties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyBaseUrl string `protobuf:"bytes,1,opt,name=keyBaseUrl,proto3" json:"keyBaseUrl,omitempty"`
}

func (x *Properties) Reset() {
	*x = Properties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_properties_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Properties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Properties) ProtoMessage() {}

func (x *Properties) ProtoReflect() protoreflect.Message {
	mi := &file_template_properties_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Properties.ProtoReflect.Descriptor instead.
func (*Properties) Descriptor() ([]byte, []int) {
	return file_template_properties_proto_rawDescGZIP(), []int{1}
}

func (x *Properties) GetKeyBaseUrl() string {
	if x != nil {
		return x.KeyBaseUrl
	}
	return ""
}

var File_template_properties_proto protoreflect.FileDescriptor

var file_template_properties_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x22, 0x13, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x42, 0x61, 0x73, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x65, 0x79, 0x42, 0x61, 0x73,
	0x65, 0x55, 0x72, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_properties_proto_rawDescOnce sync.Once
	file_template_properties_proto_rawDescData = file_template_properties_proto_rawDesc
)

func file_template_properties_proto_rawDescGZIP() []byte {
	file_template_properties_proto_rawDescOnce.Do(func() {
		file_template_properties_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_properties_proto_rawDescData)
	})
	return file_template_properties_proto_rawDescData
}

var file_template_properties_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_template_properties_proto_goTypes = []interface{}{
	(*PropertiesRequest)(nil), // 0: file.PropertiesRequest
	(*Properties)(nil),        // 1: file.Properties
}
var file_template_properties_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_template_properties_proto_init() }
func file_template_properties_proto_init() {
	if File_template_properties_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_properties_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertiesRequest); i {
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
		file_template_properties_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Properties); i {
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
			RawDescriptor: file_template_properties_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_properties_proto_goTypes,
		DependencyIndexes: file_template_properties_proto_depIdxs,
		MessageInfos:      file_template_properties_proto_msgTypes,
	}.Build()
	File_template_properties_proto = out.File
	file_template_properties_proto_rawDesc = nil
	file_template_properties_proto_goTypes = nil
	file_template_properties_proto_depIdxs = nil
}
