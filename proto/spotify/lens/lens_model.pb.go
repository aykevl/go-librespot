// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: spotify/lens/lens_model.proto

package lens

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

type Lens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *Lens) Reset() {
	*x = Lens{}
	mi := &file_spotify_lens_lens_model_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Lens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lens) ProtoMessage() {}

func (x *Lens) ProtoReflect() protoreflect.Message {
	mi := &file_spotify_lens_lens_model_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lens.ProtoReflect.Descriptor instead.
func (*Lens) Descriptor() ([]byte, []int) {
	return file_spotify_lens_lens_model_proto_rawDescGZIP(), []int{0}
}

func (x *Lens) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type LensState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Revision   []byte `protobuf:"bytes,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *LensState) Reset() {
	*x = LensState{}
	mi := &file_spotify_lens_lens_model_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LensState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LensState) ProtoMessage() {}

func (x *LensState) ProtoReflect() protoreflect.Message {
	mi := &file_spotify_lens_lens_model_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LensState.ProtoReflect.Descriptor instead.
func (*LensState) Descriptor() ([]byte, []int) {
	return file_spotify_lens_lens_model_proto_rawDescGZIP(), []int{1}
}

func (x *LensState) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *LensState) GetRevision() []byte {
	if x != nil {
		return x.Revision
	}
	return nil
}

var File_spotify_lens_lens_model_proto protoreflect.FileDescriptor

var file_spotify_lens_lens_model_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2f, 0x6c, 0x65, 0x6e, 0x73, 0x2f, 0x6c,
	0x65, 0x6e, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x6c, 0x65, 0x6e, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x26, 0x0a, 0x04, 0x4c, 0x65, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x09, 0x4c,
	0x65, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0xc8, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x70, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x6c, 0x65, 0x6e, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42,
	0x0e, 0x4c, 0x65, 0x6e, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65,
	0x76, 0x67, 0x69, 0x61, 0x6e, 0x6c, 0x75, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2f, 0x6c, 0x65, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x53, 0x4c, 0x4d, 0xaa, 0x02, 0x12,
	0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4c, 0x65, 0x6e, 0x73, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0xca, 0x02, 0x12, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5c, 0x4c, 0x65, 0x6e,
	0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0xe2, 0x02, 0x1e, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x5c, 0x4c, 0x65, 0x6e, 0x73, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x70, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x3a, 0x3a, 0x4c, 0x65, 0x6e, 0x73, 0x3a, 0x3a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spotify_lens_lens_model_proto_rawDescOnce sync.Once
	file_spotify_lens_lens_model_proto_rawDescData = file_spotify_lens_lens_model_proto_rawDesc
)

func file_spotify_lens_lens_model_proto_rawDescGZIP() []byte {
	file_spotify_lens_lens_model_proto_rawDescOnce.Do(func() {
		file_spotify_lens_lens_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_spotify_lens_lens_model_proto_rawDescData)
	})
	return file_spotify_lens_lens_model_proto_rawDescData
}

var file_spotify_lens_lens_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spotify_lens_lens_model_proto_goTypes = []any{
	(*Lens)(nil),      // 0: spotify.lens.model.Lens
	(*LensState)(nil), // 1: spotify.lens.model.LensState
}
var file_spotify_lens_lens_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spotify_lens_lens_model_proto_init() }
func file_spotify_lens_lens_model_proto_init() {
	if File_spotify_lens_lens_model_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spotify_lens_lens_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spotify_lens_lens_model_proto_goTypes,
		DependencyIndexes: file_spotify_lens_lens_model_proto_depIdxs,
		MessageInfos:      file_spotify_lens_lens_model_proto_msgTypes,
	}.Build()
	File_spotify_lens_lens_model_proto = out.File
	file_spotify_lens_lens_model_proto_rawDesc = nil
	file_spotify_lens_lens_model_proto_goTypes = nil
	file_spotify_lens_lens_model_proto_depIdxs = nil
}
