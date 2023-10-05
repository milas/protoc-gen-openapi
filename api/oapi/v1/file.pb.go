// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: oapi/v1/file.proto

package oapiv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The default host for all services and methods defined in a file. This can
	// be overridden by the service or a method definition.
	//
	// Deprecated: Do not use.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// The default prefix for all services and methods in a file. This can be
	// overridden by the service or a method definition.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Security schemes defined in a file. All files will be consolidated into one
	// list.
	SecuritySchemes []*SecurityScheme `protobuf:"bytes,3,rep,name=security_schemes,json=securitySchemes,proto3" json:"security_schemes,omitempty"`
	// Security to be used on all services and routes by default in the file.
	Security []*Security `protobuf:"bytes,4,rep,name=security,proto3" json:"security,omitempty"`
	// The servers to add to the global. These are used on all services and routes
	// by default.
	Servers []*Server `protobuf:"bytes,5,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oapi_v1_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_oapi_v1_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
func (*FileOptions) Descriptor() ([]byte, []int) {
	return file_oapi_v1_file_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Do not use.
func (x *FileOptions) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *FileOptions) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *FileOptions) GetSecuritySchemes() []*SecurityScheme {
	if x != nil {
		return x.SecuritySchemes
	}
	return nil
}

func (x *FileOptions) GetSecurity() []*Security {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *FileOptions) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

var file_oapi_v1_file_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileOptions)(nil),
		Field:         5150,
		Name:          "oapi.v1.file",
		Tag:           "bytes,5150,opt,name=file",
		Filename:      "oapi/v1/file.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional oapi.v1.FileOptions file = 5150;
	E_File = &file_oapi_v1_file_proto_extTypes[0]
)

var File_oapi_v1_file_proto protoreflect.FileDescriptor

var file_oapi_v1_file_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01,
	0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x42, 0x0a,
	0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x52, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x73, 0x12, 0x2d, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x3a, 0x47, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x9e, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x42, 0x97, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65,
	0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x6a, 0x6f, 0x73, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f, 0x61, 0x70, 0x69,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4f, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13,
	0x4f, 0x61, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4f, 0x61, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oapi_v1_file_proto_rawDescOnce sync.Once
	file_oapi_v1_file_proto_rawDescData = file_oapi_v1_file_proto_rawDesc
)

func file_oapi_v1_file_proto_rawDescGZIP() []byte {
	file_oapi_v1_file_proto_rawDescOnce.Do(func() {
		file_oapi_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_oapi_v1_file_proto_rawDescData)
	})
	return file_oapi_v1_file_proto_rawDescData
}

var file_oapi_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_oapi_v1_file_proto_goTypes = []interface{}{
	(*FileOptions)(nil),              // 0: oapi.v1.FileOptions
	(*SecurityScheme)(nil),           // 1: oapi.v1.SecurityScheme
	(*Security)(nil),                 // 2: oapi.v1.Security
	(*Server)(nil),                   // 3: oapi.v1.Server
	(*descriptorpb.FileOptions)(nil), // 4: google.protobuf.FileOptions
}
var file_oapi_v1_file_proto_depIdxs = []int32{
	1, // 0: oapi.v1.FileOptions.security_schemes:type_name -> oapi.v1.SecurityScheme
	2, // 1: oapi.v1.FileOptions.security:type_name -> oapi.v1.Security
	3, // 2: oapi.v1.FileOptions.servers:type_name -> oapi.v1.Server
	4, // 3: oapi.v1.file:extendee -> google.protobuf.FileOptions
	0, // 4: oapi.v1.file:type_name -> oapi.v1.FileOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_oapi_v1_file_proto_init() }
func file_oapi_v1_file_proto_init() {
	if File_oapi_v1_file_proto != nil {
		return
	}
	file_oapi_v1_security_proto_init()
	file_oapi_v1_server_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_oapi_v1_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOptions); i {
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
			RawDescriptor: file_oapi_v1_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_oapi_v1_file_proto_goTypes,
		DependencyIndexes: file_oapi_v1_file_proto_depIdxs,
		MessageInfos:      file_oapi_v1_file_proto_msgTypes,
		ExtensionInfos:    file_oapi_v1_file_proto_extTypes,
	}.Build()
	File_oapi_v1_file_proto = out.File
	file_oapi_v1_file_proto_rawDesc = nil
	file_oapi_v1_file_proto_goTypes = nil
	file_oapi_v1_file_proto_depIdxs = nil
}
