// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: machine/v1/machine.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type GetMachineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineId string `protobuf:"bytes,2,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"` // machine_id 不能为空字符串
}

func (x *GetMachineReq) Reset() {
	*x = GetMachineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_v1_machine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachineReq) ProtoMessage() {}

func (x *GetMachineReq) ProtoReflect() protoreflect.Message {
	mi := &file_machine_v1_machine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachineReq.ProtoReflect.Descriptor instead.
func (*GetMachineReq) Descriptor() ([]byte, []int) {
	return file_machine_v1_machine_proto_rawDescGZIP(), []int{0}
}

func (x *GetMachineReq) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

type GetMachineReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MachineName string `protobuf:"bytes,2,opt,name=machine_name,json=machineName,proto3" json:"machine_name,omitempty"`
}

func (x *GetMachineReply) Reset() {
	*x = GetMachineReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_machine_v1_machine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMachineReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMachineReply) ProtoMessage() {}

func (x *GetMachineReply) ProtoReflect() protoreflect.Message {
	mi := &file_machine_v1_machine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMachineReply.ProtoReflect.Descriptor instead.
func (*GetMachineReply) Descriptor() ([]byte, []int) {
	return file_machine_v1_machine_proto_rawDescGZIP(), []int{1}
}

func (x *GetMachineReply) GetMachineName() string {
	if x != nil {
		return x.MachineName
	}
	return ""
}

var File_machine_v1_machine_proto protoreflect.FileDescriptor

var file_machine_v1_machine_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6b, 0x74, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64,
	0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x0c, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x0b, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32,
	0x88, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x7a,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x23, 0x2e, 0x6b,
	0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x25, 0x2e, 0x6b, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2f, 0x7b, 0x6d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x32, 0x0a, 0x14, 0x6b, 0x74,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x18, 0x2f, 0x6b, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_machine_v1_machine_proto_rawDescOnce sync.Once
	file_machine_v1_machine_proto_rawDescData = file_machine_v1_machine_proto_rawDesc
)

func file_machine_v1_machine_proto_rawDescGZIP() []byte {
	file_machine_v1_machine_proto_rawDescOnce.Do(func() {
		file_machine_v1_machine_proto_rawDescData = protoimpl.X.CompressGZIP(file_machine_v1_machine_proto_rawDescData)
	})
	return file_machine_v1_machine_proto_rawDescData
}

var file_machine_v1_machine_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_machine_v1_machine_proto_goTypes = []interface{}{
	(*GetMachineReq)(nil),   // 0: kt.api.helloworld.v1.GetMachineReq
	(*GetMachineReply)(nil), // 1: kt.api.helloworld.v1.GetMachineReply
}
var file_machine_v1_machine_proto_depIdxs = []int32{
	0, // 0: kt.api.helloworld.v1.GetMachine.GetMachine:input_type -> kt.api.helloworld.v1.GetMachineReq
	1, // 1: kt.api.helloworld.v1.GetMachine.GetMachine:output_type -> kt.api.helloworld.v1.GetMachineReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_machine_v1_machine_proto_init() }
func file_machine_v1_machine_proto_init() {
	if File_machine_v1_machine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_machine_v1_machine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachineReq); i {
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
		file_machine_v1_machine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMachineReply); i {
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
			RawDescriptor: file_machine_v1_machine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_machine_v1_machine_proto_goTypes,
		DependencyIndexes: file_machine_v1_machine_proto_depIdxs,
		MessageInfos:      file_machine_v1_machine_proto_msgTypes,
	}.Build()
	File_machine_v1_machine_proto = out.File
	file_machine_v1_machine_proto_rawDesc = nil
	file_machine_v1_machine_proto_goTypes = nil
	file_machine_v1_machine_proto_depIdxs = nil
}
