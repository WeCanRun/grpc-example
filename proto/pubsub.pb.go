// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: proto/pubsub.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Publish string `protobuf:"bytes,1,opt,name=publish,proto3" json:"publish,omitempty"`
}

func (x *PubRequest) Reset() {
	*x = PubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubRequest) ProtoMessage() {}

func (x *PubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubRequest.ProtoReflect.Descriptor instead.
func (*PubRequest) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *PubRequest) GetPublish() string {
	if x != nil {
		return x.Publish
	}
	return ""
}

type PubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PubResponse) Reset() {
	*x = PubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pubsub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubResponse) ProtoMessage() {}

func (x *PubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubResponse.ProtoReflect.Descriptor instead.
func (*PubResponse) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{1}
}

func (x *PubResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type SubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscribe string `protobuf:"bytes,1,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
}

func (x *SubRequest) Reset() {
	*x = SubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pubsub_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubRequest) ProtoMessage() {}

func (x *SubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubRequest.ProtoReflect.Descriptor instead.
func (*SubRequest) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{2}
}

func (x *SubRequest) GetSubscribe() string {
	if x != nil {
		return x.Subscribe
	}
	return ""
}

type SubResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SubResponse) Reset() {
	*x = SubResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pubsub_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubResponse) ProtoMessage() {}

func (x *SubResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pubsub_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubResponse.ProtoReflect.Descriptor instead.
func (*SubResponse) Descriptor() ([]byte, []int) {
	return file_proto_pubsub_proto_rawDescGZIP(), []int{3}
}

func (x *SubResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_proto_pubsub_proto protoreflect.FileDescriptor

var file_proto_pubsub_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a,
	0x0a, 0x50, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x22, 0x29, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2a, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x22, 0x29, 0x0a, 0x0b,
	0x53, 0x75, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xab, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x53,
	0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x3a, 0x01, 0x2a, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pubsub_proto_rawDescOnce sync.Once
	file_proto_pubsub_proto_rawDescData = file_proto_pubsub_proto_rawDesc
)

func file_proto_pubsub_proto_rawDescGZIP() []byte {
	file_proto_pubsub_proto_rawDescOnce.Do(func() {
		file_proto_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pubsub_proto_rawDescData)
	})
	return file_proto_pubsub_proto_rawDescData
}

var file_proto_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_pubsub_proto_goTypes = []interface{}{
	(*PubRequest)(nil),  // 0: proto.PubRequest
	(*PubResponse)(nil), // 1: proto.PubResponse
	(*SubRequest)(nil),  // 2: proto.SubRequest
	(*SubResponse)(nil), // 3: proto.SubResponse
	(*Response)(nil),    // 4: proto.Response
}
var file_proto_pubsub_proto_depIdxs = []int32{
	0, // 0: proto.PubSubService.Publish:input_type -> proto.PubRequest
	2, // 1: proto.PubSubService.Subscribe:input_type -> proto.SubRequest
	4, // 2: proto.PubSubService.Publish:output_type -> proto.Response
	4, // 3: proto.PubSubService.Subscribe:output_type -> proto.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_pubsub_proto_init() }
func file_proto_pubsub_proto_init() {
	if File_proto_pubsub_proto != nil {
		return
	}
	file_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubRequest); i {
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
		file_proto_pubsub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubResponse); i {
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
		file_proto_pubsub_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubRequest); i {
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
		file_proto_pubsub_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubResponse); i {
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
			RawDescriptor: file_proto_pubsub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pubsub_proto_goTypes,
		DependencyIndexes: file_proto_pubsub_proto_depIdxs,
		MessageInfos:      file_proto_pubsub_proto_msgTypes,
	}.Build()
	File_proto_pubsub_proto = out.File
	file_proto_pubsub_proto_rawDesc = nil
	file_proto_pubsub_proto_goTypes = nil
	file_proto_pubsub_proto_depIdxs = nil
}
