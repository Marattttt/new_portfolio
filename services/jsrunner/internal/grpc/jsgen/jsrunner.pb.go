// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.6
// source: jsrunner.proto

package jsgen

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

type JsRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *JsRunRequest) Reset() {
	*x = JsRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsrunner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsRunRequest) ProtoMessage() {}

func (x *JsRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jsrunner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsRunRequest.ProtoReflect.Descriptor instead.
func (*JsRunRequest) Descriptor() ([]byte, []int) {
	return file_jsrunner_proto_rawDescGZIP(), []int{0}
}

func (x *JsRunRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type JsRunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *JsRunResponse) Reset() {
	*x = JsRunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsrunner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsRunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsRunResponse) ProtoMessage() {}

func (x *JsRunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jsrunner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsRunResponse.ProtoReflect.Descriptor instead.
func (*JsRunResponse) Descriptor() ([]byte, []int) {
	return file_jsrunner_proto_rawDescGZIP(), []int{1}
}

func (x *JsRunResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *JsRunResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_jsrunner_proto protoreflect.FileDescriptor

var file_jsrunner_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6a, 0x73, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6a, 0x73, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0c, 0x4a, 0x73,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3d,
	0x0a, 0x0d, 0x4a, 0x73, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x44, 0x0a,
	0x08, 0x4a, 0x73, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x05, 0x52, 0x75, 0x6e,
	0x4a, 0x73, 0x12, 0x16, 0x2e, 0x6a, 0x73, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x4a, 0x73,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6a, 0x73, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x4a, 0x73, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6a, 0x73, 0x67, 0x65, 0x6e, 0x3b, 0x6a, 0x73, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jsrunner_proto_rawDescOnce sync.Once
	file_jsrunner_proto_rawDescData = file_jsrunner_proto_rawDesc
)

func file_jsrunner_proto_rawDescGZIP() []byte {
	file_jsrunner_proto_rawDescOnce.Do(func() {
		file_jsrunner_proto_rawDescData = protoimpl.X.CompressGZIP(file_jsrunner_proto_rawDescData)
	})
	return file_jsrunner_proto_rawDescData
}

var file_jsrunner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_jsrunner_proto_goTypes = []interface{}{
	(*JsRunRequest)(nil),  // 0: jsrunner.JsRunRequest
	(*JsRunResponse)(nil), // 1: jsrunner.JsRunResponse
}
var file_jsrunner_proto_depIdxs = []int32{
	0, // 0: jsrunner.JsRunner.RunJs:input_type -> jsrunner.JsRunRequest
	1, // 1: jsrunner.JsRunner.RunJs:output_type -> jsrunner.JsRunResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jsrunner_proto_init() }
func file_jsrunner_proto_init() {
	if File_jsrunner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jsrunner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsRunRequest); i {
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
		file_jsrunner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsRunResponse); i {
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
			RawDescriptor: file_jsrunner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jsrunner_proto_goTypes,
		DependencyIndexes: file_jsrunner_proto_depIdxs,
		MessageInfos:      file_jsrunner_proto_msgTypes,
	}.Build()
	File_jsrunner_proto = out.File
	file_jsrunner_proto_rawDesc = nil
	file_jsrunner_proto_goTypes = nil
	file_jsrunner_proto_depIdxs = nil
}