// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/server/api/enums/v1/workflow.proto

package enums

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

type WorkflowExecutionState int32

const (
	WORKFLOW_EXECUTION_STATE_UNSPECIFIED WorkflowExecutionState = 0
	WORKFLOW_EXECUTION_STATE_CREATED     WorkflowExecutionState = 1
	WORKFLOW_EXECUTION_STATE_RUNNING     WorkflowExecutionState = 2
	WORKFLOW_EXECUTION_STATE_COMPLETED   WorkflowExecutionState = 3
	WORKFLOW_EXECUTION_STATE_ZOMBIE      WorkflowExecutionState = 4
	WORKFLOW_EXECUTION_STATE_VOID        WorkflowExecutionState = 5
	WORKFLOW_EXECUTION_STATE_CORRUPTED   WorkflowExecutionState = 6
)

var WorkflowExecutionState_name = map[int32]string{
	0: "Unspecified",
	1: "Created",
	2: "Running",
	3: "Completed",
	4: "Zombie",
	5: "Void",
	6: "Corrupted",
}

var WorkflowExecutionState_value = map[string]int32{
	"Unspecified": 0,
	"Created":     1,
	"Running":     2,
	"Completed":   3,
	"Zombie":      4,
	"Void":        5,
	"Corrupted":   6,
}

func (WorkflowExecutionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_004b7fefe981a755, []int{0}
}

type WorkflowBackoffType int32

const (
	WORKFLOW_BACKOFF_TYPE_UNSPECIFIED WorkflowBackoffType = 0
	WORKFLOW_BACKOFF_TYPE_RETRY       WorkflowBackoffType = 1
	WORKFLOW_BACKOFF_TYPE_CRON        WorkflowBackoffType = 2
)

var WorkflowBackoffType_name = map[int32]string{
	0: "Unspecified",
	1: "Retry",
	2: "Cron",
}

var WorkflowBackoffType_value = map[string]int32{
	"Unspecified": 0,
	"Retry":       1,
	"Cron":        2,
}

func (WorkflowBackoffType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_004b7fefe981a755, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.server.api.enums.v1.WorkflowExecutionState", WorkflowExecutionState_name, WorkflowExecutionState_value)
	proto.RegisterEnum("temporal.server.api.enums.v1.WorkflowBackoffType", WorkflowBackoffType_name, WorkflowBackoffType_value)
}

func init() {
	proto.RegisterFile("temporal/server/api/enums/v1/workflow.proto", fileDescriptor_004b7fefe981a755)
}

var fileDescriptor_004b7fefe981a755 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x49, 0xcd, 0x2d,
	0xc8, 0x2f, 0x4a, 0xcc, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x4f, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcd, 0x2b, 0xcd, 0x2d, 0xd6, 0x2f, 0x33, 0xd4, 0x2f, 0xcf, 0x2f, 0xca, 0x4e, 0xcb, 0xc9,
	0x2f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x81, 0x29, 0xd6, 0x83, 0x28, 0xd6, 0x4b,
	0x2c, 0xc8, 0xd4, 0x03, 0x2b, 0xd6, 0x2b, 0x33, 0xd4, 0x5a, 0xc6, 0xc4, 0x25, 0x16, 0x0e, 0xd5,
	0xe0, 0x5a, 0x91, 0x9a, 0x5c, 0x5a, 0x92, 0x99, 0x9f, 0x17, 0x5c, 0x92, 0x58, 0x92, 0x2a, 0xa4,
	0xc1, 0xa5, 0x12, 0xee, 0x1f, 0xe4, 0xed, 0xe6, 0xe3, 0x1f, 0x1e, 0xef, 0x1a, 0xe1, 0xea, 0x1c,
	0x1a, 0xe2, 0xe9, 0xef, 0x17, 0x1f, 0x1c, 0xe2, 0x18, 0xe2, 0x1a, 0x1f, 0xea, 0x17, 0x1c, 0xe0,
	0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0xa4, 0xc2, 0xa5, 0x80, 0x53, 0xa5, 0x73,
	0x90, 0xab, 0x63, 0x88, 0xab, 0x8b, 0x00, 0x23, 0x5e, 0x55, 0x41, 0xa1, 0x7e, 0x7e, 0x9e, 0x7e,
	0xee, 0x02, 0x4c, 0x42, 0x6a, 0x5c, 0x4a, 0xb8, 0xcd, 0xf2, 0xf7, 0x0d, 0xf0, 0x71, 0x05, 0x99,
	0xc6, 0x2c, 0xa4, 0xcc, 0x25, 0x8f, 0x53, 0x5d, 0x94, 0xbf, 0xaf, 0x93, 0xa7, 0xab, 0x00, 0x8b,
	0x90, 0x22, 0x97, 0x2c, 0x4e, 0x45, 0x61, 0xfe, 0x9e, 0x2e, 0x02, 0xac, 0x04, 0xec, 0x0b, 0x0a,
	0x0a, 0x0d, 0x00, 0xd9, 0xc7, 0xa6, 0x55, 0xcb, 0x25, 0x0c, 0x0b, 0x27, 0xa7, 0xc4, 0xe4, 0xec,
	0xfc, 0xb4, 0xb4, 0x90, 0xca, 0x82, 0x54, 0x21, 0x55, 0x2e, 0x45, 0xb8, 0x76, 0x27, 0x47, 0x67,
	0x6f, 0x7f, 0x37, 0xb7, 0xf8, 0x90, 0xc8, 0x00, 0xf4, 0x10, 0x92, 0xe7, 0x92, 0xc6, 0xae, 0x2c,
	0xc8, 0x35, 0x24, 0x28, 0x52, 0x80, 0x51, 0x48, 0x8e, 0x4b, 0x0a, 0xbb, 0x02, 0xe7, 0x20, 0x7f,
	0x3f, 0x01, 0x26, 0xa7, 0xf4, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50,
	0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x86, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xb0, 0x64, 0x90, 0x99, 0x0f, 0x67, 0xa2,
	0xa4, 0x1b, 0x6b, 0x30, 0x23, 0x89, 0x0d, 0x9c, 0x6a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x9e, 0xde, 0x89, 0x71, 0x64, 0x02, 0x00, 0x00,
}

func (x WorkflowExecutionState) String() string {
	s, ok := WorkflowExecutionState_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x WorkflowBackoffType) String() string {
	s, ok := WorkflowBackoffType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}