// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.6
// source: gorunner.proto

package gogen

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

type RunGoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *RunGoRequest) Reset() {
	*x = RunGoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorunner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunGoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunGoRequest) ProtoMessage() {}

func (x *RunGoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gorunner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunGoRequest.ProtoReflect.Descriptor instead.
func (*RunGoRequest) Descriptor() ([]byte, []int) {
	return file_gorunner_proto_rawDescGZIP(), []int{0}
}

func (x *RunGoRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RunResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error  string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *RunResponse) Reset() {
	*x = RunResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gorunner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunResponse) ProtoMessage() {}

func (x *RunResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gorunner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunResponse.ProtoReflect.Descriptor instead.
func (*RunResponse) Descriptor() ([]byte, []int) {
	return file_gorunner_proto_rawDescGZIP(), []int{1}
}

func (x *RunResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *RunResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_gorunner_proto protoreflect.FileDescriptor

var file_gorunner_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x6f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x22, 0x0a,
	0x0c, 0x52, 0x75, 0x6e, 0x47, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x3b, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x48,
	0x0a, 0x08, 0x47, 0x6f, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x05, 0x52, 0x75,
	0x6e, 0x47, 0x6f, 0x12, 0x19, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x72, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x2e, 0x52, 0x75, 0x6e, 0x47, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x2e, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x67, 0x65, 0x6e,
	0x3b, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gorunner_proto_rawDescOnce sync.Once
	file_gorunner_proto_rawDescData = file_gorunner_proto_rawDesc
)

func file_gorunner_proto_rawDescGZIP() []byte {
	file_gorunner_proto_rawDescOnce.Do(func() {
		file_gorunner_proto_rawDescData = protoimpl.X.CompressGZIP(file_gorunner_proto_rawDescData)
	})
	return file_gorunner_proto_rawDescData
}

var file_gorunner_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gorunner_proto_goTypes = []interface{}{
	(*RunGoRequest)(nil), // 0: multirunner.RunGoRequest
	(*RunResponse)(nil),  // 1: multirunner.RunResponse
}
var file_gorunner_proto_depIdxs = []int32{
	0, // 0: multirunner.GoRunner.RunGo:input_type -> multirunner.RunGoRequest
	1, // 1: multirunner.GoRunner.RunGo:output_type -> multirunner.RunResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gorunner_proto_init() }
func file_gorunner_proto_init() {
	if File_gorunner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gorunner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunGoRequest); i {
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
		file_gorunner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunResponse); i {
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
			RawDescriptor: file_gorunner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gorunner_proto_goTypes,
		DependencyIndexes: file_gorunner_proto_depIdxs,
		MessageInfos:      file_gorunner_proto_msgTypes,
	}.Build()
	File_gorunner_proto = out.File
	file_gorunner_proto_rawDesc = nil
	file_gorunner_proto_goTypes = nil
	file_gorunner_proto_depIdxs = nil
}
