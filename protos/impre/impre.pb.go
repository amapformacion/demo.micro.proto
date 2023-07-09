// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: protos/impre/impre.proto

// package : package.Service/rpcMethod
// Ejemplo: impre.Emisiones/TestPdf

package impre

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

// The request message containing the user's name.
type TestPdfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TestPdfRequest) Reset() {
	*x = TestPdfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_impre_impre_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestPdfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPdfRequest) ProtoMessage() {}

func (x *TestPdfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_impre_impre_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPdfRequest.ProtoReflect.Descriptor instead.
func (*TestPdfRequest) Descriptor() ([]byte, []int) {
	return file_protos_impre_impre_proto_rawDescGZIP(), []int{0}
}

func (x *TestPdfRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings.
type TestPdfReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pdf     string `protobuf:"bytes,1,opt,name=pdf,proto3" json:"pdf,omitempty"`
	Tamanio int64  `protobuf:"varint,2,opt,name=tamanio,proto3" json:"tamanio,omitempty"`
}

func (x *TestPdfReply) Reset() {
	*x = TestPdfReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_impre_impre_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestPdfReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPdfReply) ProtoMessage() {}

func (x *TestPdfReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_impre_impre_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPdfReply.ProtoReflect.Descriptor instead.
func (*TestPdfReply) Descriptor() ([]byte, []int) {
	return file_protos_impre_impre_proto_rawDescGZIP(), []int{1}
}

func (x *TestPdfReply) GetPdf() string {
	if x != nil {
		return x.Pdf
	}
	return ""
}

func (x *TestPdfReply) GetTamanio() int64 {
	if x != nil {
		return x.Tamanio
	}
	return 0
}

var File_protos_impre_impre_proto protoreflect.FileDescriptor

var file_protos_impre_impre_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x2f, 0x69,
	0x6d, 0x70, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x6d, 0x70, 0x72,
	0x65, 0x22, 0x24, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x64, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x64, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x64, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x64, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x6d,
	0x61, 0x6e, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x61, 0x6d, 0x61,
	0x6e, 0x69, 0x6f, 0x32, 0x42, 0x0a, 0x09, 0x45, 0x6d, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x73,
	0x12, 0x35, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x50, 0x64, 0x66, 0x12, 0x15, 0x2e, 0x69, 0x6d,
	0x70, 0x72, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x64, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x64, 0x66, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x69, 0x6d, 0x70,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_impre_impre_proto_rawDescOnce sync.Once
	file_protos_impre_impre_proto_rawDescData = file_protos_impre_impre_proto_rawDesc
)

func file_protos_impre_impre_proto_rawDescGZIP() []byte {
	file_protos_impre_impre_proto_rawDescOnce.Do(func() {
		file_protos_impre_impre_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_impre_impre_proto_rawDescData)
	})
	return file_protos_impre_impre_proto_rawDescData
}

var file_protos_impre_impre_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_impre_impre_proto_goTypes = []interface{}{
	(*TestPdfRequest)(nil), // 0: impre.TestPdfRequest
	(*TestPdfReply)(nil),   // 1: impre.TestPdfReply
}
var file_protos_impre_impre_proto_depIdxs = []int32{
	0, // 0: impre.Emisiones.TestPdf:input_type -> impre.TestPdfRequest
	1, // 1: impre.Emisiones.TestPdf:output_type -> impre.TestPdfReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_impre_impre_proto_init() }
func file_protos_impre_impre_proto_init() {
	if File_protos_impre_impre_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_impre_impre_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestPdfRequest); i {
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
		file_protos_impre_impre_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestPdfReply); i {
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
			RawDescriptor: file_protos_impre_impre_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_impre_impre_proto_goTypes,
		DependencyIndexes: file_protos_impre_impre_proto_depIdxs,
		MessageInfos:      file_protos_impre_impre_proto_msgTypes,
	}.Build()
	File_protos_impre_impre_proto = out.File
	file_protos_impre_impre_proto_rawDesc = nil
	file_protos_impre_impre_proto_goTypes = nil
	file_protos_impre_impre_proto_depIdxs = nil
}
