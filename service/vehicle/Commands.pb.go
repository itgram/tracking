// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: vehicle/Commands.proto

package vehicle

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

type RegisterVehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId string `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	Model     string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *RegisterVehicle) Reset() {
	*x = RegisterVehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_Commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterVehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterVehicle) ProtoMessage() {}

func (x *RegisterVehicle) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_Commands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterVehicle.ProtoReflect.Descriptor instead.
func (*RegisterVehicle) Descriptor() ([]byte, []int) {
	return file_vehicle_Commands_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterVehicle) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

func (x *RegisterVehicle) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type AdjustMaxSpeedVehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId string `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	MaxSpeed  int32  `protobuf:"varint,2,opt,name=max_speed,json=maxSpeed,proto3" json:"max_speed,omitempty"`
}

func (x *AdjustMaxSpeedVehicle) Reset() {
	*x = AdjustMaxSpeedVehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_Commands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdjustMaxSpeedVehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdjustMaxSpeedVehicle) ProtoMessage() {}

func (x *AdjustMaxSpeedVehicle) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_Commands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdjustMaxSpeedVehicle.ProtoReflect.Descriptor instead.
func (*AdjustMaxSpeedVehicle) Descriptor() ([]byte, []int) {
	return file_vehicle_Commands_proto_rawDescGZIP(), []int{1}
}

func (x *AdjustMaxSpeedVehicle) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

func (x *AdjustMaxSpeedVehicle) GetMaxSpeed() int32 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

var File_vehicle_Commands_proto protoreflect.FileDescriptor

var file_vehicle_Commands_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x69, 0x74, 0x67, 0x72, 0x61, 0x6d,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x0f, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x22, 0x53, 0x0a, 0x15, 0x41, 0x64, 0x6a, 0x75, 0x73, 0x74, 0x4d, 0x61, 0x78, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61,
	0x78, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d,
	0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x67, 0x72, 0x61, 0x6d, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vehicle_Commands_proto_rawDescOnce sync.Once
	file_vehicle_Commands_proto_rawDescData = file_vehicle_Commands_proto_rawDesc
)

func file_vehicle_Commands_proto_rawDescGZIP() []byte {
	file_vehicle_Commands_proto_rawDescOnce.Do(func() {
		file_vehicle_Commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_vehicle_Commands_proto_rawDescData)
	})
	return file_vehicle_Commands_proto_rawDescData
}

var file_vehicle_Commands_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vehicle_Commands_proto_goTypes = []interface{}{
	(*RegisterVehicle)(nil),       // 0: itgram.tracking.commands.vehicle.RegisterVehicle
	(*AdjustMaxSpeedVehicle)(nil), // 1: itgram.tracking.commands.vehicle.AdjustMaxSpeedVehicle
}
var file_vehicle_Commands_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vehicle_Commands_proto_init() }
func file_vehicle_Commands_proto_init() {
	if File_vehicle_Commands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vehicle_Commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterVehicle); i {
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
		file_vehicle_Commands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdjustMaxSpeedVehicle); i {
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
			RawDescriptor: file_vehicle_Commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vehicle_Commands_proto_goTypes,
		DependencyIndexes: file_vehicle_Commands_proto_depIdxs,
		MessageInfos:      file_vehicle_Commands_proto_msgTypes,
	}.Build()
	File_vehicle_Commands_proto = out.File
	file_vehicle_Commands_proto_rawDesc = nil
	file_vehicle_Commands_proto_goTypes = nil
	file_vehicle_Commands_proto_depIdxs = nil
}
