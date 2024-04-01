// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: proto/bridge.proto

package proto

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Person) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

type PersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *PersonRequest) Reset() {
	*x = PersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRequest) ProtoMessage() {}

func (x *PersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRequest.ProtoReflect.Descriptor instead.
func (*PersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *PersonRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type PersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hi string `protobuf:"bytes,1,opt,name=hi,proto3" json:"hi,omitempty"`
}

func (x *PersonResponse) Reset() {
	*x = PersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bridge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonResponse) ProtoMessage() {}

func (x *PersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bridge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonResponse.ProtoReflect.Descriptor instead.
func (*PersonResponse) Descriptor() ([]byte, []int) {
	return file_proto_bridge_proto_rawDescGZIP(), []int{2}
}

func (x *PersonResponse) GetHi() string {
	if x != nil {
		return x.Hi
	}
	return ""
}

var File_proto_bridge_proto protoreflect.FileDescriptor

var file_proto_bridge_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69,
	0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69,
	0x64, 0x6f, 0x22, 0x30, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x68, 0x69, 0x32, 0x40, 0x0a, 0x0d, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x6c, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bridge_proto_rawDescOnce sync.Once
	file_proto_bridge_proto_rawDescData = file_proto_bridge_proto_rawDesc
)

func file_proto_bridge_proto_rawDescGZIP() []byte {
	file_proto_bridge_proto_rawDescOnce.Do(func() {
		file_proto_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bridge_proto_rawDescData)
	})
	return file_proto_bridge_proto_rawDescData
}

var file_proto_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_bridge_proto_goTypes = []interface{}{
	(*Person)(nil),         // 0: Person
	(*PersonRequest)(nil),  // 1: PersonRequest
	(*PersonResponse)(nil), // 2: PersonResponse
}
var file_proto_bridge_proto_depIdxs = []int32{
	0, // 0: PersonRequest.person:type_name -> Person
	1, // 1: BridgeService.CallPerson:input_type -> PersonRequest
	2, // 2: BridgeService.CallPerson:output_type -> PersonResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_bridge_proto_init() }
func file_proto_bridge_proto_init() {
	if File_proto_bridge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_proto_bridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonRequest); i {
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
		file_proto_bridge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonResponse); i {
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
			RawDescriptor: file_proto_bridge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_bridge_proto_goTypes,
		DependencyIndexes: file_proto_bridge_proto_depIdxs,
		MessageInfos:      file_proto_bridge_proto_msgTypes,
	}.Build()
	File_proto_bridge_proto = out.File
	file_proto_bridge_proto_rawDesc = nil
	file_proto_bridge_proto_goTypes = nil
	file_proto_bridge_proto_depIdxs = nil
}
