// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: proto/comunicacion.proto

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

type EstadoPersona struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre   string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,2,opt,name=apellido,proto3" json:"apellido,omitempty"`
	Estado   string `protobuf:"bytes,3,opt,name=estado,proto3" json:"estado,omitempty"`
}

func (x *EstadoPersona) Reset() {
	*x = EstadoPersona{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comunicacion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstadoPersona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstadoPersona) ProtoMessage() {}

func (x *EstadoPersona) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comunicacion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstadoPersona.ProtoReflect.Descriptor instead.
func (*EstadoPersona) Descriptor() ([]byte, []int) {
	return file_proto_comunicacion_proto_rawDescGZIP(), []int{0}
}

func (x *EstadoPersona) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *EstadoPersona) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

func (x *EstadoPersona) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

type IdPersona struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nombre   string `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Apellido string `protobuf:"bytes,3,opt,name=apellido,proto3" json:"apellido,omitempty"`
}

func (x *IdPersona) Reset() {
	*x = IdPersona{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comunicacion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdPersona) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdPersona) ProtoMessage() {}

func (x *IdPersona) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comunicacion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdPersona.ProtoReflect.Descriptor instead.
func (*IdPersona) Descriptor() ([]byte, []int) {
	return file_proto_comunicacion_proto_rawDescGZIP(), []int{1}
}

func (x *IdPersona) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IdPersona) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *IdPersona) GetApellido() string {
	if x != nil {
		return x.Apellido
	}
	return ""
}

var File_proto_comunicacion_proto protoreflect.FileDescriptor

var file_proto_comunicacion_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x63, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63,
	0x22, 0x5b, 0x0a, 0x0d, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65,
	0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65,
	0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x4f, 0x0a,
	0x09, 0x49, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x69, 0x64, 0x6f, 0x32, 0x6f,
	0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x12, 0x33, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x4f, 0x4d, 0x53, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x73, 0x74, 0x61,
	0x64, 0x6f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x12, 0x2f,
	0x0a, 0x0b, 0x4f, 0x4d, 0x53, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x1a, 0x0f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_comunicacion_proto_rawDescOnce sync.Once
	file_proto_comunicacion_proto_rawDescData = file_proto_comunicacion_proto_rawDesc
)

func file_proto_comunicacion_proto_rawDescGZIP() []byte {
	file_proto_comunicacion_proto_rawDescOnce.Do(func() {
		file_proto_comunicacion_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_comunicacion_proto_rawDescData)
	})
	return file_proto_comunicacion_proto_rawDescData
}

var file_proto_comunicacion_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_comunicacion_proto_goTypes = []interface{}{
	(*EstadoPersona)(nil), // 0: grpc.EstadoPersona
	(*IdPersona)(nil),     // 1: grpc.IdPersona
}
var file_proto_comunicacion_proto_depIdxs = []int32{
	0, // 0: grpc.ComServ.ContOMS:input_type -> grpc.EstadoPersona
	1, // 1: grpc.ComServ.OMSDataNode:input_type -> grpc.IdPersona
	0, // 2: grpc.ComServ.ContOMS:output_type -> grpc.EstadoPersona
	1, // 3: grpc.ComServ.OMSDataNode:output_type -> grpc.IdPersona
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_comunicacion_proto_init() }
func file_proto_comunicacion_proto_init() {
	if File_proto_comunicacion_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_comunicacion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstadoPersona); i {
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
		file_proto_comunicacion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdPersona); i {
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
			RawDescriptor: file_proto_comunicacion_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_comunicacion_proto_goTypes,
		DependencyIndexes: file_proto_comunicacion_proto_depIdxs,
		MessageInfos:      file_proto_comunicacion_proto_msgTypes,
	}.Build()
	File_proto_comunicacion_proto = out.File
	file_proto_comunicacion_proto_rawDesc = nil
	file_proto_comunicacion_proto_goTypes = nil
	file_proto_comunicacion_proto_depIdxs = nil
}
