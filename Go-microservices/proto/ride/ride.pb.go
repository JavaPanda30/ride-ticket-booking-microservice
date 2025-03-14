// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.20.3
// source: proto/ride.proto

package ride

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RideRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	PassengerId    string                 `protobuf:"bytes,1,opt,name=passenger_id,json=passengerId,proto3" json:"passenger_id,omitempty"`
	PickupLocation string                 `protobuf:"bytes,2,opt,name=pickup_location,json=pickupLocation,proto3" json:"pickup_location,omitempty"`
	DropLocation   string                 `protobuf:"bytes,3,opt,name=drop_location,json=dropLocation,proto3" json:"drop_location,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *RideRequest) Reset() {
	*x = RideRequest{}
	mi := &file_proto_ride_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideRequest) ProtoMessage() {}

func (x *RideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ride_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideRequest.ProtoReflect.Descriptor instead.
func (*RideRequest) Descriptor() ([]byte, []int) {
	return file_proto_ride_proto_rawDescGZIP(), []int{0}
}

func (x *RideRequest) GetPassengerId() string {
	if x != nil {
		return x.PassengerId
	}
	return ""
}

func (x *RideRequest) GetPickupLocation() string {
	if x != nil {
		return x.PickupLocation
	}
	return ""
}

func (x *RideRequest) GetDropLocation() string {
	if x != nil {
		return x.DropLocation
	}
	return ""
}

type RideResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RideId        string                 `protobuf:"bytes,1,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	DriverId      string                 `protobuf:"bytes,3,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	Success       bool                   `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RideResponse) Reset() {
	*x = RideResponse{}
	mi := &file_proto_ride_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideResponse) ProtoMessage() {}

func (x *RideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ride_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideResponse.ProtoReflect.Descriptor instead.
func (*RideResponse) Descriptor() ([]byte, []int) {
	return file_proto_ride_proto_rawDescGZIP(), []int{1}
}

func (x *RideResponse) GetRideId() string {
	if x != nil {
		return x.RideId
	}
	return ""
}

func (x *RideResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RideResponse) GetDriverId() string {
	if x != nil {
		return x.DriverId
	}
	return ""
}

func (x *RideResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RideStatusUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RideId        string                 `protobuf:"bytes,1,opt,name=ride_id,json=rideId,proto3" json:"ride_id,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RideStatusUpdate) Reset() {
	*x = RideStatusUpdate{}
	mi := &file_proto_ride_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RideStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RideStatusUpdate) ProtoMessage() {}

func (x *RideStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ride_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RideStatusUpdate.ProtoReflect.Descriptor instead.
func (*RideStatusUpdate) Descriptor() ([]byte, []int) {
	return file_proto_ride_proto_rawDescGZIP(), []int{2}
}

func (x *RideStatusUpdate) GetRideId() string {
	if x != nil {
		return x.RideId
	}
	return ""
}

func (x *RideStatusUpdate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_ride_proto protoreflect.FileDescriptor

var file_proto_ride_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x72, 0x69, 0x64, 0x65, 0x22, 0x7e, 0x0a, 0x0b, 0x52, 0x69, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x69,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x72, 0x6f, 0x70,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x0c, 0x52, 0x69, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x69, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x69, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x43, 0x0a, 0x10, 0x52, 0x69, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x69, 0x64, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x82, 0x01, 0x0a, 0x0b, 0x52, 0x69, 0x64, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x69, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x52, 0x69,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x69, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x2e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x52, 0x69, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x12, 0x2e, 0x72, 0x69, 0x64, 0x65, 0x2e, 0x52, 0x69,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f,
	0x47, 0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x69, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_proto_ride_proto_rawDescOnce sync.Once
	file_proto_ride_proto_rawDescData []byte
)

func file_proto_ride_proto_rawDescGZIP() []byte {
	file_proto_ride_proto_rawDescOnce.Do(func() {
		file_proto_ride_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_ride_proto_rawDesc), len(file_proto_ride_proto_rawDesc)))
	})
	return file_proto_ride_proto_rawDescData
}

var file_proto_ride_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_ride_proto_goTypes = []any{
	(*RideRequest)(nil),      // 0: ride.RideRequest
	(*RideResponse)(nil),     // 1: ride.RideResponse
	(*RideStatusUpdate)(nil), // 2: ride.RideStatusUpdate
}
var file_proto_ride_proto_depIdxs = []int32{
	0, // 0: ride.RideService.CreateRide:input_type -> ride.RideRequest
	2, // 1: ride.RideService.UpdateRideStatus:input_type -> ride.RideStatusUpdate
	1, // 2: ride.RideService.CreateRide:output_type -> ride.RideResponse
	1, // 3: ride.RideService.UpdateRideStatus:output_type -> ride.RideResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ride_proto_init() }
func file_proto_ride_proto_init() {
	if File_proto_ride_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_ride_proto_rawDesc), len(file_proto_ride_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ride_proto_goTypes,
		DependencyIndexes: file_proto_ride_proto_depIdxs,
		MessageInfos:      file_proto_ride_proto_msgTypes,
	}.Build()
	File_proto_ride_proto = out.File
	file_proto_ride_proto_goTypes = nil
	file_proto_ride_proto_depIdxs = nil
}
