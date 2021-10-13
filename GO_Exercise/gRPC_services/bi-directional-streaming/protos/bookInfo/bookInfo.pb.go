// this generates the implementation of the interface for the class
// maps you how to handle the other methods that you define for your service to the
// server

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: protos/bookInfo.proto

package bookInfo

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

type Ratings int32

const (
	Ratings_R1 Ratings = 0
	Ratings_R2 Ratings = 1
	Ratings_R3 Ratings = 2
	Ratings_R4 Ratings = 3
	Ratings_R5 Ratings = 4
)

// Enum value maps for Ratings.
var (
	Ratings_name = map[int32]string{
		0: "R1",
		1: "R2",
		2: "R3",
		3: "R4",
		4: "R5",
	}
	Ratings_value = map[string]int32{
		"R1": 0,
		"R2": 1,
		"R3": 2,
		"R4": 3,
		"R5": 4,
	}
)

func (x Ratings) Enum() *Ratings {
	p := new(Ratings)
	*p = x
	return p
}

func (x Ratings) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ratings) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_bookInfo_proto_enumTypes[0].Descriptor()
}

func (Ratings) Type() protoreflect.EnumType {
	return &file_protos_bookInfo_proto_enumTypes[0]
}

func (x Ratings) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ratings.Descriptor instead.
func (Ratings) EnumDescriptor() ([]byte, []int) {
	return file_protos_bookInfo_proto_rawDescGZIP(), []int{0}
}

// RateRequest defines the request for a GetRate call
type RateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base is the base currency code for the rate
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Destination is the destination currency code for the rate
	Review string `protobuf:"bytes,2,opt,name=review,proto3" json:"review,omitempty"`
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bookInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bookInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_protos_bookInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RateRequest) GetReview() string {
	if x != nil {
		return x.Review
	}
	return ""
}

// RateResponse is the response from a GetRate call, it contains
// rate which is a floating point number and can be used to convert between the
// two currencies specified in the request.
type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating Ratings `protobuf:"varint,1,opt,name=rating,proto3,enum=Ratings" json:"rating,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bookInfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bookInfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_protos_bookInfo_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetRating() Ratings {
	if x != nil {
		return x.Rating
	}
	return Ratings_R1
}

var File_protos_bookInfo_proto protoreflect.FileDescriptor

var file_protos_bookInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x22, 0x30, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2a, 0x31, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x31, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x32, 0x10,
	0x01, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x33, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x34, 0x10,
	0x03, 0x12, 0x06, 0x0a, 0x02, 0x52, 0x35, 0x10, 0x04, 0x32, 0x64, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x0c, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x0d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x0c,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_bookInfo_proto_rawDescOnce sync.Once
	file_protos_bookInfo_proto_rawDescData = file_protos_bookInfo_proto_rawDesc
)

func file_protos_bookInfo_proto_rawDescGZIP() []byte {
	file_protos_bookInfo_proto_rawDescOnce.Do(func() {
		file_protos_bookInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_bookInfo_proto_rawDescData)
	})
	return file_protos_bookInfo_proto_rawDescData
}

var file_protos_bookInfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_bookInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_bookInfo_proto_goTypes = []interface{}{
	(Ratings)(0),         // 0: Ratings
	(*RateRequest)(nil),  // 1: RateRequest
	(*RateResponse)(nil), // 2: RateResponse
}
var file_protos_bookInfo_proto_depIdxs = []int32{
	0, // 0: RateResponse.rating:type_name -> Ratings
	1, // 1: bookInfo.GetRate:input_type -> RateRequest
	1, // 2: bookInfo.SubscribeRate:input_type -> RateRequest
	2, // 3: bookInfo.GetRate:output_type -> RateResponse
	2, // 4: bookInfo.SubscribeRate:output_type -> RateResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_bookInfo_proto_init() }
func file_protos_bookInfo_proto_init() {
	if File_protos_bookInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_bookInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRequest); i {
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
		file_protos_bookInfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
			RawDescriptor: file_protos_bookInfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_bookInfo_proto_goTypes,
		DependencyIndexes: file_protos_bookInfo_proto_depIdxs,
		EnumInfos:         file_protos_bookInfo_proto_enumTypes,
		MessageInfos:      file_protos_bookInfo_proto_msgTypes,
	}.Build()
	File_protos_bookInfo_proto = out.File
	file_protos_bookInfo_proto_rawDesc = nil
	file_protos_bookInfo_proto_goTypes = nil
	file_protos_bookInfo_proto_depIdxs = nil
}
