// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: proto/echo.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/jexia/semaphore/api"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Enum int32

const (
	Enum_UNKNOWN Enum = 0
	Enum_ON      Enum = 1
	Enum_OFF     Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "UNKNOWN",
		1: "ON",
		2: "OFF",
	}
	Enum_value = map[string]int32{
		"UNKNOWN": 0,
		"ON":      1,
		"OFF":     2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_echo_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_proto_echo_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_proto_echo_proto_rawDescGZIP(), []int{0}
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enum    Enum    `protobuf:"varint,1,opt,name=enum,proto3,enum=semaphore.typetest.Enum" json:"enum,omitempty"`
	String_ string  `protobuf:"bytes,2,opt,name=string,proto3" json:"string,omitempty"`
	Integer int64   `protobuf:"varint,3,opt,name=integer,proto3" json:"integer,omitempty"`
	Double  float64 `protobuf:"fixed64,4,opt,name=double,proto3" json:"double,omitempty"`
	Numbers []int64 `protobuf:"varint,5,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_proto_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_proto_echo_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetEnum() Enum {
	if x != nil {
		return x.Enum
	}
	return Enum_UNKNOWN
}

func (x *Data) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Data) GetInteger() int64 {
	if x != nil {
		return x.Integer
	}
	return 0
}

func (x *Data) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

func (x *Data) GetNumbers() []int64 {
	if x != nil {
		return x.Numbers
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_echo_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo *Data `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_echo_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetEcho() *Data {
	if x != nil {
		return x.Echo
	}
	return nil
}

var File_proto_echo_proto protoreflect.FileDescriptor

var file_proto_echo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04,
	0x65, 0x6e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x37, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a,
	0x04, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65,
	0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x2a, 0x24, 0x0a, 0x04, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10,
	0x02, 0x32, 0x4e, 0x0a, 0x08, 0x54, 0x79, 0x70, 0x65, 0x74, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a,
	0x03, 0x52, 0x75, 0x6e, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x32, 0x5b, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3c, 0x0a,
	0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x18, 0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x1a, 0x11, 0xe2, 0xb5, 0x18,
	0x0d, 0x22, 0x04, 0x67, 0x72, 0x70, 0x63, 0x2a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x65, 0x78,
	0x69, 0x61, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x65, 0x32, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_echo_proto_rawDescOnce sync.Once
	file_proto_echo_proto_rawDescData = file_proto_echo_proto_rawDesc
)

func file_proto_echo_proto_rawDescGZIP() []byte {
	file_proto_echo_proto_rawDescOnce.Do(func() {
		file_proto_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_echo_proto_rawDescData)
	})
	return file_proto_echo_proto_rawDescData
}

var file_proto_echo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_echo_proto_goTypes = []interface{}{
	(Enum)(0),        // 0: semaphore.typetest.Enum
	(*Data)(nil),     // 1: semaphore.typetest.Data
	(*Request)(nil),  // 2: semaphore.typetest.Request
	(*Response)(nil), // 3: semaphore.typetest.Response
}
var file_proto_echo_proto_depIdxs = []int32{
	0, // 0: semaphore.typetest.Data.enum:type_name -> semaphore.typetest.Enum
	1, // 1: semaphore.typetest.Request.data:type_name -> semaphore.typetest.Data
	1, // 2: semaphore.typetest.Response.echo:type_name -> semaphore.typetest.Data
	2, // 3: semaphore.typetest.Typetest.Run:input_type -> semaphore.typetest.Request
	1, // 4: semaphore.typetest.External.Post:input_type -> semaphore.typetest.Data
	3, // 5: semaphore.typetest.Typetest.Run:output_type -> semaphore.typetest.Response
	1, // 6: semaphore.typetest.External.Post:output_type -> semaphore.typetest.Data
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_echo_proto_init() }
func file_proto_echo_proto_init() {
	if File_proto_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_proto_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_echo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_echo_proto_goTypes,
		DependencyIndexes: file_proto_echo_proto_depIdxs,
		EnumInfos:         file_proto_echo_proto_enumTypes,
		MessageInfos:      file_proto_echo_proto_msgTypes,
	}.Build()
	File_proto_echo_proto = out.File
	file_proto_echo_proto_rawDesc = nil
	file_proto_echo_proto_goTypes = nil
	file_proto_echo_proto_depIdxs = nil
}
