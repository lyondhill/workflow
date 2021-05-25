// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/common.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EncodingType int32

const (
	ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	ENCODING_TYPE_PROTO3      EncodingType = 1
	ENCODING_TYPE_JSON        EncodingType = 2
)

var EncodingType_name = map[int32]string{
	0: "Unspecified",
	1: "Proto3",
	2: "Json",
}

var EncodingType_value = map[string]int32{
	"Unspecified": 0,
	"Proto3":      1,
	"Json":        2,
}

func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768283dde248a0f8, []int{0}
}

type IndexedValueType int32

const (
	INDEXED_VALUE_TYPE_UNSPECIFIED IndexedValueType = 0
	INDEXED_VALUE_TYPE_STRING      IndexedValueType = 1
	INDEXED_VALUE_TYPE_KEYWORD     IndexedValueType = 2
	INDEXED_VALUE_TYPE_INT         IndexedValueType = 3
	INDEXED_VALUE_TYPE_DOUBLE      IndexedValueType = 4
	INDEXED_VALUE_TYPE_BOOL        IndexedValueType = 5
	INDEXED_VALUE_TYPE_DATETIME    IndexedValueType = 6
)

var IndexedValueType_name = map[int32]string{
	0: "Unspecified",
	1: "String",
	2: "Keyword",
	3: "Int",
	4: "Double",
	5: "Bool",
	6: "Datetime",
}

var IndexedValueType_value = map[string]int32{
	"Unspecified": 0,
	"String":      1,
	"Keyword":     2,
	"Int":         3,
	"Double":      4,
	"Bool":        5,
	"Datetime":    6,
}

func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768283dde248a0f8, []int{1}
}

type Severity int32

const (
	SEVERITY_UNSPECIFIED Severity = 0
	SEVERITY_HIGH        Severity = 1
	SEVERITY_MEDIUM      Severity = 2
	SEVERITY_LOW         Severity = 3
)

var Severity_name = map[int32]string{
	0: "Unspecified",
	1: "High",
	2: "Medium",
	3: "Low",
}

var Severity_value = map[string]int32{
	"Unspecified": 0,
	"High":        1,
	"Medium":      2,
	"Low":         3,
}

func (Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_768283dde248a0f8, []int{2}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.EncodingType", EncodingType_name, EncodingType_value)
	proto.RegisterEnum("temporal.api.enums.v1.IndexedValueType", IndexedValueType_name, IndexedValueType_value)
	proto.RegisterEnum("temporal.api.enums.v1.Severity", Severity_name, Severity_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/common.proto", fileDescriptor_768283dde248a0f8)
}

var fileDescriptor_768283dde248a0f8 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x6e, 0xd3, 0x40,
	0x18, 0xc7, 0x7d, 0x2e, 0x54, 0xe8, 0xa3, 0x88, 0xe3, 0xa0, 0x25, 0x6d, 0xd5, 0x43, 0x74, 0xcc,
	0xe0, 0x28, 0xea, 0x66, 0xa6, 0x24, 0xfe, 0x48, 0x0f, 0x12, 0x9f, 0x15, 0x5f, 0x5c, 0xc2, 0x40,
	0x14, 0x1a, 0xab, 0x3a, 0xa9, 0xf1, 0x59, 0x21, 0x8d, 0xe8, 0xc6, 0x23, 0xf0, 0x18, 0x88, 0x95,
	0x97, 0x60, 0xcc, 0xd8, 0x91, 0x38, 0x0b, 0x62, 0x40, 0x7d, 0x04, 0x14, 0x4b, 0x89, 0x08, 0x32,
	0xdb, 0xe9, 0xff, 0xfb, 0xdf, 0xf7, 0xe9, 0x4e, 0x3f, 0x38, 0x9e, 0xc4, 0xa3, 0xd4, 0x8c, 0x07,
	0x97, 0x95, 0x41, 0xaa, 0x2b, 0x71, 0x72, 0x35, 0xfa, 0x50, 0x99, 0x56, 0x2b, 0xe7, 0x66, 0x34,
	0x32, 0x89, 0x93, 0x8e, 0xcd, 0xc4, 0xb0, 0xdd, 0x55, 0xc7, 0x19, 0xa4, 0xda, 0xc9, 0x3b, 0xce,
	0xb4, 0x5a, 0xee, 0xc3, 0x0e, 0x26, 0xe7, 0x66, 0xa8, 0x93, 0x0b, 0x75, 0x9d, 0xc6, 0xec, 0x08,
	0xf6, 0xd1, 0x6f, 0x48, 0x4f, 0xf8, 0xcd, 0xbe, 0xea, 0x05, 0xd8, 0xef, 0xfa, 0x61, 0x80, 0x0d,
	0xf1, 0x52, 0xa0, 0x47, 0x2d, 0x56, 0x82, 0x27, 0x9b, 0x38, 0xe8, 0x48, 0x25, 0x4f, 0x28, 0x61,
	0x7b, 0xc0, 0x36, 0xc9, 0xab, 0x50, 0xfa, 0xd4, 0x2e, 0xff, 0x26, 0x40, 0x45, 0x32, 0x8c, 0x3f,
	0xc6, 0xc3, 0x68, 0x70, 0x79, 0x15, 0xe7, 0x5b, 0x8e, 0x81, 0x0b, 0xdf, 0xc3, 0x37, 0xe8, 0xf5,
	0xa3, 0x5a, 0xab, 0x8b, 0x45, 0xab, 0x8e, 0x60, 0xbf, 0xa0, 0x13, 0xaa, 0x8e, 0xf0, 0x9b, 0x94,
	0x30, 0x0e, 0x07, 0x05, 0xf8, 0x35, 0xf6, 0xce, 0x64, 0xc7, 0xa3, 0x36, 0x3b, 0x80, 0xbd, 0x02,
	0x2e, 0x7c, 0x45, 0xb7, 0xfe, 0x33, 0xda, 0x93, 0xdd, 0x7a, 0x0b, 0xe9, 0x1d, 0x76, 0x08, 0x4f,
	0x0b, 0x70, 0x5d, 0xca, 0x16, 0xbd, 0xcb, 0x9e, 0xc1, 0x61, 0xd1, 0xdd, 0x9a, 0x42, 0x25, 0xda,
	0x48, 0xb7, 0xcb, 0xef, 0xe0, 0x5e, 0x18, 0x4f, 0xe3, 0xb1, 0x9e, 0x5c, 0x2f, 0xbf, 0x2b, 0xc4,
	0x08, 0x3b, 0x42, 0xf5, 0xfe, 0x79, 0xdd, 0x23, 0x78, 0xb0, 0x26, 0xa7, 0xa2, 0x79, 0x4a, 0x09,
	0x7b, 0x0c, 0x0f, 0xd7, 0x51, 0x1b, 0x3d, 0xd1, 0x6d, 0x53, 0x9b, 0x51, 0xd8, 0x59, 0x87, 0x2d,
	0x79, 0x46, 0xb7, 0xea, 0xdf, 0xc8, 0x6c, 0xce, 0xad, 0x9b, 0x39, 0xb7, 0x6e, 0xe7, 0x9c, 0x7c,
	0xca, 0x38, 0xf9, 0x92, 0x71, 0xf2, 0x3d, 0xe3, 0x64, 0x96, 0x71, 0xf2, 0x23, 0xe3, 0xe4, 0x67,
	0xc6, 0xad, 0xdb, 0x8c, 0x93, 0xcf, 0x0b, 0x6e, 0xcd, 0x16, 0xdc, 0xba, 0x59, 0x70, 0x0b, 0x4a,
	0xda, 0x38, 0x85, 0x0a, 0xd4, 0xef, 0x37, 0x72, 0x4f, 0x82, 0xa5, 0x26, 0x01, 0x79, 0xfb, 0xfc,
	0xe2, 0xaf, 0xa2, 0x36, 0x1b, 0x4a, 0xbd, 0xc8, 0x0f, 0x5f, 0xed, 0x5d, 0xb5, 0x2a, 0xd4, 0x52,
	0xed, 0x60, 0x3e, 0x29, 0xaa, 0xfe, 0xb2, 0x4b, 0xab, 0xdc, 0x75, 0x6b, 0xa9, 0x76, 0xdd, 0x9c,
	0xb8, 0x6e, 0x54, 0x7d, 0xbf, 0x9d, 0x5b, 0x78, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x36,
	0x14, 0xcf, 0xab, 0x02, 0x00, 0x00,
}

func (x EncodingType) String() string {
	s, ok := EncodingType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x IndexedValueType) String() string {
	s, ok := IndexedValueType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x Severity) String() string {
	s, ok := Severity_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}