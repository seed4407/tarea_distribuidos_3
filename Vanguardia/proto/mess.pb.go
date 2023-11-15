// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mess.proto

package Tarea_distribuidos_3

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

type ParametroSectorBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombreSector string `protobuf:"bytes,1,opt,name=nombre_sector,json=nombreSector,proto3" json:"nombre_sector,omitempty"`
	NombreBase   string `protobuf:"bytes,2,opt,name=nombre_base,json=nombreBase,proto3" json:"nombre_base,omitempty"`
}

func (x *ParametroSectorBase) Reset() {
	*x = ParametroSectorBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParametroSectorBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParametroSectorBase) ProtoMessage() {}

func (x *ParametroSectorBase) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParametroSectorBase.ProtoReflect.Descriptor instead.
func (*ParametroSectorBase) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{0}
}

func (x *ParametroSectorBase) GetNombreSector() string {
	if x != nil {
		return x.NombreSector
	}
	return ""
}

func (x *ParametroSectorBase) GetNombreBase() string {
	if x != nil {
		return x.NombreBase
	}
	return ""
}

type RetornoCantSoldOscuridadReloj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CantSoldadosOscuridad int32   `protobuf:"varint,1,opt,name=cant_soldados_oscuridad,json=cantSoldadosOscuridad,proto3" json:"cant_soldados_oscuridad,omitempty"`
	Reloj                 []int32 `protobuf:"varint,2,rep,packed,name=reloj,proto3" json:"reloj,omitempty"`
}

func (x *RetornoCantSoldOscuridadReloj) Reset() {
	*x = RetornoCantSoldOscuridadReloj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetornoCantSoldOscuridadReloj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetornoCantSoldOscuridadReloj) ProtoMessage() {}

func (x *RetornoCantSoldOscuridadReloj) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetornoCantSoldOscuridadReloj.ProtoReflect.Descriptor instead.
func (*RetornoCantSoldOscuridadReloj) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{1}
}

func (x *RetornoCantSoldOscuridadReloj) GetCantSoldadosOscuridad() int32 {
	if x != nil {
		return x.CantSoldadosOscuridad
	}
	return 0
}

func (x *RetornoCantSoldOscuridadReloj) GetReloj() []int32 {
	if x != nil {
		return x.Reloj
	}
	return nil
}

type ServidorAConsultar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ServidorAConsultar) Reset() {
	*x = ServidorAConsultar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServidorAConsultar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServidorAConsultar) ProtoMessage() {}

func (x *ServidorAConsultar) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServidorAConsultar.ProtoReflect.Descriptor instead.
func (*ServidorAConsultar) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{2}
}

func (x *ServidorAConsultar) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type ParamamtroVacio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ParamamtroVacio) Reset() {
	*x = ParamamtroVacio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParamamtroVacio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParamamtroVacio) ProtoMessage() {}

func (x *ParamamtroVacio) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParamamtroVacio.ProtoReflect.Descriptor instead.
func (*ParamamtroVacio) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{3}
}

type Recepcion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok string `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Recepcion) Reset() {
	*x = Recepcion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mess_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recepcion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recepcion) ProtoMessage() {}

func (x *Recepcion) ProtoReflect() protoreflect.Message {
	mi := &file_mess_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recepcion.ProtoReflect.Descriptor instead.
func (*Recepcion) Descriptor() ([]byte, []int) {
	return file_mess_proto_rawDescGZIP(), []int{4}
}

func (x *Recepcion) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

var File_mess_proto protoreflect.FileDescriptor

var file_mess_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x22, 0x5d, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f,
	0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6e,
	0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x42, 0x61, 0x73,
	0x65, 0x22, 0x71, 0x0a, 0x21, 0x72, 0x65, 0x74, 0x6f, 0x72, 0x6e, 0x6f, 0x5f, 0x63, 0x61, 0x6e,
	0x74, 0x5f, 0x73, 0x6f, 0x6c, 0x64, 0x5f, 0x6f, 0x73, 0x63, 0x75, 0x72, 0x69, 0x64, 0x61, 0x64,
	0x5f, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x61, 0x6e, 0x74, 0x5f, 0x73,
	0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x5f, 0x6f, 0x73, 0x63, 0x75, 0x72, 0x69, 0x64, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x63, 0x61, 0x6e, 0x74, 0x53, 0x6f, 0x6c,
	0x64, 0x61, 0x64, 0x6f, 0x73, 0x4f, 0x73, 0x63, 0x75, 0x72, 0x69, 0x64, 0x61, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x72,
	0x65, 0x6c, 0x6f, 0x6a, 0x22, 0x26, 0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72,
	0x5f, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x12, 0x0a, 0x10,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x61, 0x6d, 0x74, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x63, 0x69, 0x6f,
	0x22, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x63, 0x0a,
	0x0a, 0x56, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x12, 0x55, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x53, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x72,
	0x65, 0x74, 0x6f, 0x72, 0x6e, 0x6f, 0x5f, 0x63, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x6f, 0x6c, 0x64,
	0x5f, 0x6f, 0x73, 0x63, 0x75, 0x72, 0x69, 0x64, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x6c, 0x6f, 0x6a,
	0x22, 0x00, 0x32, 0x6b, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65,
	0x73, 0x12, 0x5c, 0x0a, 0x24, 0x4f, 0x62, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x64, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x5f,
	0x61, 0x6c, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x69, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x61, 0x6d, 0x74, 0x72, 0x6f, 0x5f, 0x76, 0x61, 0x63, 0x69,
	0x6f, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x64, 0x6f,
	0x72, 0x5f, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x72, 0x22, 0x00, 0x42,
	0x5a, 0x0a, 0x1b, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0f,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65,
	0x65, 0x64, 0x34, 0x34, 0x30, 0x37, 0x2f, 0x54, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x69, 0x64, 0x6f, 0x73, 0x5f, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mess_proto_rawDescOnce sync.Once
	file_mess_proto_rawDescData = file_mess_proto_rawDesc
)

func file_mess_proto_rawDescGZIP() []byte {
	file_mess_proto_rawDescOnce.Do(func() {
		file_mess_proto_rawDescData = protoimpl.X.CompressGZIP(file_mess_proto_rawDescData)
	})
	return file_mess_proto_rawDescData
}

var file_mess_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mess_proto_goTypes = []interface{}{
	(*ParametroSectorBase)(nil),           // 0: grpc.parametro_sector_base
	(*RetornoCantSoldOscuridadReloj)(nil), // 1: grpc.retorno_cant_sold_oscuridad_reloj
	(*ServidorAConsultar)(nil),            // 2: grpc.servidor_a_consultar
	(*ParamamtroVacio)(nil),               // 3: grpc.paramamtro_vacio
	(*Recepcion)(nil),                     // 4: grpc.recepcion
}
var file_mess_proto_depIdxs = []int32{
	0, // 0: grpc.Vanguardia.GetSoldados:input_type -> grpc.parametro_sector_base
	3, // 1: grpc.Informantes.Obtener_servidor_consultar_aleatorio:input_type -> grpc.paramamtro_vacio
	1, // 2: grpc.Vanguardia.GetSoldados:output_type -> grpc.retorno_cant_sold_oscuridad_reloj
	2, // 3: grpc.Informantes.Obtener_servidor_consultar_aleatorio:output_type -> grpc.servidor_a_consultar
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mess_proto_init() }
func file_mess_proto_init() {
	if File_mess_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mess_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParametroSectorBase); i {
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
		file_mess_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetornoCantSoldOscuridadReloj); i {
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
		file_mess_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServidorAConsultar); i {
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
		file_mess_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParamamtroVacio); i {
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
		file_mess_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recepcion); i {
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
			RawDescriptor: file_mess_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_mess_proto_goTypes,
		DependencyIndexes: file_mess_proto_depIdxs,
		MessageInfos:      file_mess_proto_msgTypes,
	}.Build()
	File_mess_proto = out.File
	file_mess_proto_rawDesc = nil
	file_mess_proto_goTypes = nil
	file_mess_proto_depIdxs = nil
}