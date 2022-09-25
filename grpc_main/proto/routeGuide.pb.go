// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: routeGuide/routeGuide.proto

package routeGuide

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

type Detail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Detail) Reset() {
	*x = Detail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeGuide_routeGuide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detail) ProtoMessage() {}

func (x *Detail) ProtoReflect() protoreflect.Message {
	mi := &file_routeGuide_routeGuide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detail.ProtoReflect.Descriptor instead.
func (*Detail) Descriptor() ([]byte, []int) {
	return file_routeGuide_routeGuide_proto_rawDescGZIP(), []int{0}
}

func (x *Detail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Respond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res bool `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *Respond) Reset() {
	*x = Respond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeGuide_routeGuide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Respond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Respond) ProtoMessage() {}

func (x *Respond) ProtoReflect() protoreflect.Message {
	mi := &file_routeGuide_routeGuide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Respond.ProtoReflect.Descriptor instead.
func (*Respond) Descriptor() ([]byte, []int) {
	return file_routeGuide_routeGuide_proto_rawDescGZIP(), []int{1}
}

func (x *Respond) GetRes() bool {
	if x != nil {
		return x.Res
	}
	return false
}

type Temp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Temp) Reset() {
	*x = Temp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeGuide_routeGuide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Temp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Temp) ProtoMessage() {}

func (x *Temp) ProtoReflect() protoreflect.Message {
	mi := &file_routeGuide_routeGuide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Temp.ProtoReflect.Descriptor instead.
func (*Temp) Descriptor() ([]byte, []int) {
	return file_routeGuide_routeGuide_proto_rawDescGZIP(), []int{2}
}

type GetOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders string `protobuf:"bytes,1,opt,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrders) Reset() {
	*x = GetOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routeGuide_routeGuide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrders) ProtoMessage() {}

func (x *GetOrders) ProtoReflect() protoreflect.Message {
	mi := &file_routeGuide_routeGuide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrders.ProtoReflect.Descriptor instead.
func (*GetOrders) Descriptor() ([]byte, []int) {
	return file_routeGuide_routeGuide_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrders) GetOrders() string {
	if x != nil {
		return x.Orders
	}
	return ""
}

var File_routeGuide_routeGuide_proto protoreflect.FileDescriptor

var file_routeGuide_routeGuide_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x22, 0x1c, 0x0a, 0x06, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1b, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x72, 0x65, 0x73, 0x22, 0x06, 0x0a, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x22, 0x23, 0x0a, 0x09,
	0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x32, 0x70, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12,
	0x30, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x12, 0x12, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x13, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x47, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x1a, 0x15, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_routeGuide_routeGuide_proto_rawDescOnce sync.Once
	file_routeGuide_routeGuide_proto_rawDescData = file_routeGuide_routeGuide_proto_rawDesc
)

func file_routeGuide_routeGuide_proto_rawDescGZIP() []byte {
	file_routeGuide_routeGuide_proto_rawDescOnce.Do(func() {
		file_routeGuide_routeGuide_proto_rawDescData = protoimpl.X.CompressGZIP(file_routeGuide_routeGuide_proto_rawDescData)
	})
	return file_routeGuide_routeGuide_proto_rawDescData
}

var file_routeGuide_routeGuide_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_routeGuide_routeGuide_proto_goTypes = []interface{}{
	(*Detail)(nil),    // 0: routeGuide.detail
	(*Respond)(nil),   // 1: routeGuide.respond
	(*Temp)(nil),      // 2: routeGuide.temp
	(*GetOrders)(nil), // 3: routeGuide.getOrders
}
var file_routeGuide_routeGuide_proto_depIdxs = []int32{
	0, // 0: routeGuide.RouteGuide.add:input_type -> routeGuide.detail
	2, // 1: routeGuide.RouteGuide.get:input_type -> routeGuide.temp
	1, // 2: routeGuide.RouteGuide.add:output_type -> routeGuide.respond
	3, // 3: routeGuide.RouteGuide.get:output_type -> routeGuide.getOrders
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_routeGuide_routeGuide_proto_init() }
func file_routeGuide_routeGuide_proto_init() {
	if File_routeGuide_routeGuide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routeGuide_routeGuide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detail); i {
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
		file_routeGuide_routeGuide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Respond); i {
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
		file_routeGuide_routeGuide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Temp); i {
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
		file_routeGuide_routeGuide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrders); i {
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
			RawDescriptor: file_routeGuide_routeGuide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routeGuide_routeGuide_proto_goTypes,
		DependencyIndexes: file_routeGuide_routeGuide_proto_depIdxs,
		MessageInfos:      file_routeGuide_routeGuide_proto_msgTypes,
	}.Build()
	File_routeGuide_routeGuide_proto = out.File
	file_routeGuide_routeGuide_proto_rawDesc = nil
	file_routeGuide_routeGuide_proto_goTypes = nil
	file_routeGuide_routeGuide_proto_depIdxs = nil
}