// Code generated by protoc-gen-go.
// source: scout.proto
// DO NOT EDIT!

/*
Package scout is a generated protocol buffer package.

It is generated from these files:
	scout.proto

It has these top-level messages:
	ScoutReport
	FileInfo
*/
package scout

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ScoutReport contains information when a file is ready to be fetched
// Fetch contains a bool stating if file is ready to be fetched or not
// UUID is a unique ID for the FileEvent
// Timestamp is the time of FileEvent creation
type ScoutReport struct {
	Fetch     bool      `protobuf:"varint,1,opt,name=Fetch" json:"Fetch,omitempty"`
	Timestamp string    `protobuf:"bytes,2,opt,name=Timestamp" json:"Timestamp,omitempty"`
	UUID      string    `protobuf:"bytes,3,opt,name=UUID" json:"UUID,omitempty"`
	FileInfo  *FileInfo `protobuf:"bytes,4,opt,name=FileInfo" json:"FileInfo,omitempty"`
}

func (m *ScoutReport) Reset()                    { *m = ScoutReport{} }
func (m *ScoutReport) String() string            { return proto.CompactTextString(m) }
func (*ScoutReport) ProtoMessage()               {}
func (*ScoutReport) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ScoutReport) GetFetch() bool {
	if m != nil {
		return m.Fetch
	}
	return false
}

func (m *ScoutReport) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ScoutReport) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *ScoutReport) GetFileInfo() *FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

// FileInfo message contains information about file to fetch
// Path shows location to the file
// Filename is the name of the file
// Size is the file size in bytes
type FileInfo struct {
	Path     string `protobuf:"bytes,1,opt,name=Path" json:"Path,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=Filename" json:"Filename,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=Size" json:"Size,omitempty"`
}

func (m *FileInfo) Reset()                    { *m = FileInfo{} }
func (m *FileInfo) String() string            { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()               {}
func (*FileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FileInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FileInfo) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *FileInfo) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func init() {
	proto.RegisterType((*ScoutReport)(nil), "scout.ScoutReport")
	proto.RegisterType((*FileInfo)(nil), "scout.FileInfo")
}

func init() { proto.RegisterFile("scout.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x4e, 0xce, 0x2f,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x9a, 0x18, 0xb9, 0xb8,
	0x83, 0x41, 0xac, 0xa0, 0xd4, 0x82, 0xfc, 0xa2, 0x12, 0x21, 0x11, 0x2e, 0x56, 0xb7, 0xd4, 0x92,
	0xe4, 0x0c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x08, 0x47, 0x48, 0x86, 0x8b, 0x33, 0x24,
	0x33, 0x37, 0xb5, 0xb8, 0x24, 0x31, 0xb7, 0x40, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0x21,
	0x20, 0x24, 0xc4, 0xc5, 0x12, 0x1a, 0xea, 0xe9, 0x22, 0xc1, 0x0c, 0x96, 0x00, 0xb3, 0x85, 0xb4,
	0xb9, 0x38, 0xdc, 0x32, 0x73, 0x52, 0x3d, 0xf3, 0xd2, 0xf2, 0x25, 0x58, 0x14, 0x18, 0x35, 0xb8,
	0x8d, 0xf8, 0xf5, 0x20, 0xd6, 0xc3, 0x84, 0x83, 0xe0, 0x0a, 0x94, 0xfc, 0x10, 0x8a, 0x41, 0x86,
	0x05, 0x24, 0x96, 0x40, 0xec, 0xe7, 0x0c, 0x02, 0xb3, 0x85, 0xa4, 0x20, 0xf2, 0x79, 0x89, 0xb9,
	0xa9, 0x50, 0xdb, 0xe1, 0x7c, 0x90, 0xfa, 0xe0, 0xcc, 0xaa, 0x54, 0xb0, 0xe5, 0xcc, 0x41, 0x60,
	0x76, 0x12, 0x1b, 0xd8, 0x8b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x5e, 0x7e, 0x13,
	0xf1, 0x00, 0x00, 0x00,
}
