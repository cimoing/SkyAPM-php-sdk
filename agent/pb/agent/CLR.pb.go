// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/CLR.proto

package agent

import (
	common "agent/agent/pb/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CLRMetric struct {
	Time                 int64       `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu                  *common.CPU `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Gc                   *ClrGC      `protobuf:"bytes,3,opt,name=gc,proto3" json:"gc,omitempty"`
	Thread               *ClrThread  `protobuf:"bytes,4,opt,name=thread,proto3" json:"thread,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CLRMetric) Reset()         { *m = CLRMetric{} }
func (m *CLRMetric) String() string { return proto.CompactTextString(m) }
func (*CLRMetric) ProtoMessage()    {}
func (*CLRMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{0}
}

func (m *CLRMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CLRMetric.Unmarshal(m, b)
}
func (m *CLRMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CLRMetric.Marshal(b, m, deterministic)
}
func (m *CLRMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CLRMetric.Merge(m, src)
}
func (m *CLRMetric) XXX_Size() int {
	return xxx_messageInfo_CLRMetric.Size(m)
}
func (m *CLRMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_CLRMetric.DiscardUnknown(m)
}

var xxx_messageInfo_CLRMetric proto.InternalMessageInfo

func (m *CLRMetric) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *CLRMetric) GetCpu() *common.CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *CLRMetric) GetGc() *ClrGC {
	if m != nil {
		return m.Gc
	}
	return nil
}

func (m *CLRMetric) GetThread() *ClrThread {
	if m != nil {
		return m.Thread
	}
	return nil
}

type ClrGC struct {
	Gen0CollectCount     int64    `protobuf:"varint,1,opt,name=Gen0CollectCount,proto3" json:"Gen0CollectCount,omitempty"`
	Gen1CollectCount     int64    `protobuf:"varint,2,opt,name=Gen1CollectCount,proto3" json:"Gen1CollectCount,omitempty"`
	Gen2CollectCount     int64    `protobuf:"varint,3,opt,name=Gen2CollectCount,proto3" json:"Gen2CollectCount,omitempty"`
	HeapMemory           int64    `protobuf:"varint,4,opt,name=HeapMemory,proto3" json:"HeapMemory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClrGC) Reset()         { *m = ClrGC{} }
func (m *ClrGC) String() string { return proto.CompactTextString(m) }
func (*ClrGC) ProtoMessage()    {}
func (*ClrGC) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{1}
}

func (m *ClrGC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClrGC.Unmarshal(m, b)
}
func (m *ClrGC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClrGC.Marshal(b, m, deterministic)
}
func (m *ClrGC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClrGC.Merge(m, src)
}
func (m *ClrGC) XXX_Size() int {
	return xxx_messageInfo_ClrGC.Size(m)
}
func (m *ClrGC) XXX_DiscardUnknown() {
	xxx_messageInfo_ClrGC.DiscardUnknown(m)
}

var xxx_messageInfo_ClrGC proto.InternalMessageInfo

func (m *ClrGC) GetGen0CollectCount() int64 {
	if m != nil {
		return m.Gen0CollectCount
	}
	return 0
}

func (m *ClrGC) GetGen1CollectCount() int64 {
	if m != nil {
		return m.Gen1CollectCount
	}
	return 0
}

func (m *ClrGC) GetGen2CollectCount() int64 {
	if m != nil {
		return m.Gen2CollectCount
	}
	return 0
}

func (m *ClrGC) GetHeapMemory() int64 {
	if m != nil {
		return m.HeapMemory
	}
	return 0
}

type ClrThread struct {
	AvailableCompletionPortThreads int32    `protobuf:"varint,1,opt,name=AvailableCompletionPortThreads,proto3" json:"AvailableCompletionPortThreads,omitempty"`
	AvailableWorkerThreads         int32    `protobuf:"varint,2,opt,name=AvailableWorkerThreads,proto3" json:"AvailableWorkerThreads,omitempty"`
	MaxCompletionPortThreads       int32    `protobuf:"varint,3,opt,name=MaxCompletionPortThreads,proto3" json:"MaxCompletionPortThreads,omitempty"`
	MaxWorkerThreads               int32    `protobuf:"varint,4,opt,name=MaxWorkerThreads,proto3" json:"MaxWorkerThreads,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *ClrThread) Reset()         { *m = ClrThread{} }
func (m *ClrThread) String() string { return proto.CompactTextString(m) }
func (*ClrThread) ProtoMessage()    {}
func (*ClrThread) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10d56830892247a, []int{2}
}

func (m *ClrThread) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClrThread.Unmarshal(m, b)
}
func (m *ClrThread) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClrThread.Marshal(b, m, deterministic)
}
func (m *ClrThread) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClrThread.Merge(m, src)
}
func (m *ClrThread) XXX_Size() int {
	return xxx_messageInfo_ClrThread.Size(m)
}
func (m *ClrThread) XXX_DiscardUnknown() {
	xxx_messageInfo_ClrThread.DiscardUnknown(m)
}

var xxx_messageInfo_ClrThread proto.InternalMessageInfo

func (m *ClrThread) GetAvailableCompletionPortThreads() int32 {
	if m != nil {
		return m.AvailableCompletionPortThreads
	}
	return 0
}

func (m *ClrThread) GetAvailableWorkerThreads() int32 {
	if m != nil {
		return m.AvailableWorkerThreads
	}
	return 0
}

func (m *ClrThread) GetMaxCompletionPortThreads() int32 {
	if m != nil {
		return m.MaxCompletionPortThreads
	}
	return 0
}

func (m *ClrThread) GetMaxWorkerThreads() int32 {
	if m != nil {
		return m.MaxWorkerThreads
	}
	return 0
}

func init() {
	proto.RegisterType((*CLRMetric)(nil), "CLRMetric")
	proto.RegisterType((*ClrGC)(nil), "ClrGC")
	proto.RegisterType((*ClrThread)(nil), "ClrThread")
}

func init() { proto.RegisterFile("common/CLR.proto", fileDescriptor_a10d56830892247a) }

var fileDescriptor_a10d56830892247a = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0xeb, 0xd3, 0x30,
	0x18, 0xc7, 0xe9, 0x9f, 0xdf, 0x60, 0x8f, 0x97, 0x11, 0x65, 0x94, 0x1d, 0xc6, 0xe8, 0x69, 0x78,
	0xc8, 0xfe, 0x08, 0x1e, 0xbc, 0x69, 0xc0, 0x79, 0x58, 0xa5, 0x44, 0x65, 0xe0, 0x2d, 0x8b, 0xa1,
	0x2b, 0x4d, 0x93, 0x92, 0x65, 0x6e, 0x7b, 0x4b, 0xfa, 0xf2, 0x7c, 0x03, 0xd2, 0xb4, 0x16, 0xcb,
	0x36, 0x7e, 0x97, 0xf6, 0xe1, 0xf3, 0xfd, 0x7c, 0x79, 0xd2, 0x12, 0x18, 0x71, 0x5d, 0x96, 0x5a,
	0x2d, 0xc8, 0x96, 0xe2, 0xca, 0x68, 0xab, 0x27, 0x2f, 0x5b, 0xd2, 0xbc, 0x1a, 0x18, 0x1f, 0x61,
	0x48, 0xb6, 0x34, 0x11, 0xd6, 0xe4, 0x1c, 0x21, 0x08, 0x6d, 0x5e, 0x8a, 0xc8, 0x9b, 0x79, 0xf3,
	0x80, 0xba, 0x19, 0x8d, 0x21, 0xe0, 0xd5, 0x29, 0xf2, 0x67, 0xde, 0xfc, 0xc5, 0x3a, 0xc4, 0x24,
	0xfd, 0x46, 0x6b, 0x80, 0xc6, 0xe0, 0x67, 0x3c, 0x0a, 0x1c, 0x1e, 0x60, 0x22, 0xcd, 0x86, 0x50,
	0x3f, 0xe3, 0x28, 0x86, 0x81, 0x3d, 0x18, 0xc1, 0x7e, 0x44, 0xa1, 0xcb, 0xa0, 0xce, 0xbe, 0x3a,
	0x42, 0xdb, 0x24, 0xfe, 0xed, 0xc1, 0x93, 0x6b, 0xa0, 0xd7, 0x30, 0xda, 0x08, 0xb5, 0x24, 0x5a,
	0x4a, 0xc1, 0x2d, 0xd1, 0x27, 0x65, 0xdb, 0xed, 0x37, 0xbc, 0x75, 0x57, 0x3d, 0xd7, 0xef, 0xdc,
	0xd5, 0x1d, 0x77, 0xdd, 0x73, 0x83, 0xce, 0xed, 0x71, 0x34, 0x05, 0xf8, 0x24, 0x58, 0x95, 0x88,
	0x52, 0x9b, 0xab, 0x3b, 0x75, 0x40, 0xff, 0x23, 0xf1, 0x1f, 0x0f, 0x86, 0xdd, 0x37, 0xa0, 0x8f,
	0x30, 0x7d, 0xff, 0x93, 0xe5, 0x92, 0xed, 0xa5, 0x20, 0xba, 0xac, 0xa4, 0xb0, 0xb9, 0x56, 0xa9,
	0x36, 0xb6, 0x11, 0x8e, 0xee, 0xfc, 0x4f, 0xf4, 0x19, 0x0b, 0xbd, 0x85, 0x71, 0x67, 0xec, 0xb4,
	0x29, 0x84, 0xf9, 0xd7, 0xf7, 0x5d, 0xff, 0x41, 0x8a, 0xde, 0x41, 0x94, 0xb0, 0xcb, 0xfd, 0xcd,
	0x81, 0x6b, 0x3e, 0xcc, 0xeb, 0xbf, 0x92, 0xb0, 0x4b, 0x7f, 0x5b, 0xe8, 0x3a, 0x37, 0xfc, 0x43,
	0x06, 0x4b, 0x6d, 0x32, 0xcc, 0x2a, 0xc6, 0x0f, 0x02, 0x1f, 0x8b, 0xeb, 0x99, 0xc9, 0x22, 0x57,
	0x35, 0x29, 0xb1, 0x12, 0xf6, 0xac, 0x4d, 0x81, 0x25, 0x53, 0xd9, 0x89, 0x65, 0x02, 0xb3, 0x4c,
	0x28, 0x9b, 0x7a, 0xdf, 0x5f, 0xb9, 0x61, 0xd1, 0x3c, 0xab, 0x7d, 0x33, 0xfc, 0xf2, 0x27, 0x5f,
	0x8a, 0xeb, 0xae, 0xed, 0x7f, 0x6e, 0xba, 0x69, 0x7d, 0xfd, 0xb8, 0x96, 0xfb, 0x81, 0xbb, 0x88,
	0x6f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x81, 0x82, 0x4c, 0xb1, 0x02, 0x00, 0x00,
}
