// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: calc.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_calc_proto protoreflect.FileDescriptor

var file_calc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x6c, 0x63, 0x1a, 0x09, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x70,
	0x72, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x76, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x73, 0x71, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa4, 0x02,
	0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x13, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x4e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x34, 0x0a,
	0x0a, 0x41, 0x76, 0x67, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x41, 0x76, 0x67, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x41, 0x76, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x12, 0x40, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2d, 0x0a, 0x04, 0x53, 0x71, 0x72, 0x74, 0x12, 0x11, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x53, 0x71, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x65, 0x79, 0x61, 0x61, 0x64, 0x65, 0x65, 0x7a, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_calc_proto_goTypes = []interface{}{
	(*SumRequest)(nil),      // 0: calc.SumRequest
	(*PrimeNRequest)(nil),   // 1: calc.PrimeNRequest
	(*AvgNRequest)(nil),     // 2: calc.AvgNRequest
	(*StreamNRequest)(nil),  // 3: calc.StreamNRequest
	(*SqrtRequest)(nil),     // 4: calc.SqrtRequest
	(*SumResponse)(nil),     // 5: calc.SumResponse
	(*PrimeNResponse)(nil),  // 6: calc.PrimeNResponse
	(*AvgResponse)(nil),     // 7: calc.AvgResponse
	(*StreamNResponse)(nil), // 8: calc.StreamNResponse
	(*SqrtResponse)(nil),    // 9: calc.SqrtResponse
}
var file_calc_proto_depIdxs = []int32{
	0, // 0: calc.CalcService.AddNumbers:input_type -> calc.SumRequest
	1, // 1: calc.CalcService.PrimeNumbers:input_type -> calc.PrimeNRequest
	2, // 2: calc.CalcService.AvgNumbers:input_type -> calc.AvgNRequest
	3, // 3: calc.CalcService.StreamNumbers:input_type -> calc.StreamNRequest
	4, // 4: calc.CalcService.Sqrt:input_type -> calc.SqrtRequest
	5, // 5: calc.CalcService.AddNumbers:output_type -> calc.SumResponse
	6, // 6: calc.CalcService.PrimeNumbers:output_type -> calc.PrimeNResponse
	7, // 7: calc.CalcService.AvgNumbers:output_type -> calc.AvgResponse
	8, // 8: calc.CalcService.StreamNumbers:output_type -> calc.StreamNResponse
	9, // 9: calc.CalcService.Sqrt:output_type -> calc.SqrtResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calc_proto_init() }
func file_calc_proto_init() {
	if File_calc_proto != nil {
		return
	}
	file_sum_proto_init()
	file_prime_proto_init()
	file_avg_proto_init()
	file_stre_proto_init()
	file_sqrt_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_proto_goTypes,
		DependencyIndexes: file_calc_proto_depIdxs,
	}.Build()
	File_calc_proto = out.File
	file_calc_proto_rawDesc = nil
	file_calc_proto_goTypes = nil
	file_calc_proto_depIdxs = nil
}
