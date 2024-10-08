// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.21.3
// source: template/store.proto

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

type StoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data          []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Category      string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Type          string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Name          string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Extension     string `protobuf:"bytes,5,opt,name=extension,proto3" json:"extension,omitempty"`
	BaseDir       string `protobuf:"bytes,6,opt,name=baseDir,proto3" json:"baseDir,omitempty"`
	ExpiresInSecs int64  `protobuf:"varint,7,opt,name=expiresInSecs,proto3" json:"expiresInSecs,omitempty"`
	AccountId     uint64 `protobuf:"varint,8,opt,name=accountId,proto3" json:"accountId,omitempty"`
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_template_store_proto_rawDescGZIP(), []int{0}
}

func (x *StoreRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StoreRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *StoreRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoreRequest) GetExtension() string {
	if x != nil {
		return x.Extension
	}
	return ""
}

func (x *StoreRequest) GetBaseDir() string {
	if x != nil {
		return x.BaseDir
	}
	return ""
}

func (x *StoreRequest) GetExpiresInSecs() int64 {
	if x != nil {
		return x.ExpiresInSecs
	}
	return 0
}

func (x *StoreRequest) GetAccountId() uint64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type StoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *File  `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Url  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_template_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResponse.ProtoReflect.Descriptor instead.
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return file_template_store_proto_rawDescGZIP(), []int{1}
}

func (x *StoreResponse) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *StoreResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_template_store_proto protoreflect.FileDescriptor

var file_template_store_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x13, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65,
	0x44, 0x69, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x44,
	0x69, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x53,
	0x65, 0x63, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_store_proto_rawDescOnce sync.Once
	file_template_store_proto_rawDescData = file_template_store_proto_rawDesc
)

func file_template_store_proto_rawDescGZIP() []byte {
	file_template_store_proto_rawDescOnce.Do(func() {
		file_template_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_store_proto_rawDescData)
	})
	return file_template_store_proto_rawDescData
}

var file_template_store_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_template_store_proto_goTypes = []interface{}{
	(*StoreRequest)(nil),  // 0: file.StoreRequest
	(*StoreResponse)(nil), // 1: file.StoreResponse
	(*File)(nil),          // 2: file.File
}
var file_template_store_proto_depIdxs = []int32{
	2, // 0: file.StoreResponse.file:type_name -> file.File
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_template_store_proto_init() }
func file_template_store_proto_init() {
	if File_template_store_proto != nil {
		return
	}
	file_template_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_template_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreRequest); i {
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
		file_template_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreResponse); i {
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
			RawDescriptor: file_template_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_store_proto_goTypes,
		DependencyIndexes: file_template_store_proto_depIdxs,
		MessageInfos:      file_template_store_proto_msgTypes,
	}.Build()
	File_template_store_proto = out.File
	file_template_store_proto_rawDesc = nil
	file_template_store_proto_goTypes = nil
	file_template_store_proto_depIdxs = nil
}
