// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Operation int32

const (
	Operation_sum            Operation = 0
	Operation_subtraction    Operation = 1
	Operation_multiplication Operation = 2
	Operation_division       Operation = 3
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0: "sum",
		1: "subtraction",
		2: "multiplication",
		3: "division",
	}
	Operation_value = map[string]int32{
		"sum":            0,
		"subtraction":    1,
		"multiplication": 2,
		"division":       3,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_calculatorpb_calculator_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_calculatorpb_calculator_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Number) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *Number) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   *Number   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Operation Operation `protobuf:"varint,2,opt,name=operation,proto3,enum=calculatorpb.Operation" json:"operation,omitempty"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorRequest) GetContent() *Number {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *CalculatorRequest) GetOperation() Operation {
	if x != nil {
		return x.Operation
	}
	return Operation_sum
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *CalculatorResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *AverageRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x22, 0x50, 0x0a,
	0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x7a, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x70, 0x62, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x12, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x28, 0x0a, 0x0e, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x2a, 0x47, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x07, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x73, 0x75, 0x62,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x32, 0xb4, 0x01, 0x0a,
	0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x41,
	0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_calculatorpb_calculator_proto_rawDescData = file_calculatorpb_calculator_proto_rawDesc
)

func file_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculatorpb_calculator_proto_rawDescData)
	})
	return file_calculatorpb_calculator_proto_rawDescData
}

var file_calculatorpb_calculator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calculatorpb_calculator_proto_goTypes = []interface{}{
	(Operation)(0),             // 0: calculatorpb.Operation
	(*Number)(nil),             // 1: calculatorpb.Number
	(*CalculatorRequest)(nil),  // 2: calculatorpb.CalculatorRequest
	(*CalculatorResponse)(nil), // 3: calculatorpb.CalculatorResponse
	(*AverageRequest)(nil),     // 4: calculatorpb.AverageRequest
}
var file_calculatorpb_calculator_proto_depIdxs = []int32{
	1, // 0: calculatorpb.CalculatorRequest.content:type_name -> calculatorpb.Number
	0, // 1: calculatorpb.CalculatorRequest.operation:type_name -> calculatorpb.Operation
	2, // 2: calculatorpb.CalculatorService.Calculate:input_type -> calculatorpb.CalculatorRequest
	4, // 3: calculatorpb.CalculatorService.Average:input_type -> calculatorpb.AverageRequest
	3, // 4: calculatorpb.CalculatorService.Calculate:output_type -> calculatorpb.CalculatorResponse
	3, // 5: calculatorpb.CalculatorService.Average:output_type -> calculatorpb.CalculatorResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_calculatorpb_calculator_proto_init() }
func file_calculatorpb_calculator_proto_init() {
	if File_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
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
		file_calculatorpb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
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
		file_calculatorpb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
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
			RawDescriptor: file_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_calculatorpb_calculator_proto_depIdxs,
		EnumInfos:         file_calculatorpb_calculator_proto_enumTypes,
		MessageInfos:      file_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_calculatorpb_calculator_proto = out.File
	file_calculatorpb_calculator_proto_rawDesc = nil
	file_calculatorpb_calculator_proto_goTypes = nil
	file_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	//Unary
	Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	//Streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Calculate(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculatorpb.CalculatorService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculatorpb.CalculatorService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*CalculatorResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*CalculatorResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalculatorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	//Unary
	Calculate(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	//Streaming
	Average(CalculatorService_AverageServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Calculate(context.Context, *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (*UnimplementedCalculatorServiceServer) Average(CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculatorpb.CalculatorService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Calculate(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*CalculatorResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *CalculatorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculatorpb.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculatorService_Calculate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculatorpb/calculator.proto",
}
