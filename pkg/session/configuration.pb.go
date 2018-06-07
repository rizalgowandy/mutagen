// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/havoc-io/mutagen/pkg/session/configuration.proto

package session

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import filesystem "github.com/havoc-io/mutagen/pkg/filesystem"
import sync "github.com/havoc-io/mutagen/pkg/sync"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configuration encodes session configuration parameters. Each session has two
// configuration components: the configuration due to parameters provided at
// creation time and the snapshot of the global configuration at creation time.
type Configuration struct {
	// Ignores specifies the ignore patterns that should be used in
	// synchronization.
	Ignores []string `protobuf:"bytes,1,rep,name=ignores" json:"ignores,omitempty"`
	// SymlinkMode specifies the symlink mode that should be used in
	// synchronization.
	SymlinkMode sync.SymlinkMode `protobuf:"varint,11,opt,name=symlinkMode,enum=sync.SymlinkMode" json:"symlinkMode,omitempty"`
	// WatchMode specifies the filesystem watching mode.
	WatchMode filesystem.WatchMode `protobuf:"varint,21,opt,name=watchMode,enum=filesystem.WatchMode" json:"watchMode,omitempty"`
	// WatchPollingInterval specifies the interval (in seconds) for poll-based
	// file monitoring. A value of 0 specifies that the default interval should
	// be used.
	WatchPollingInterval uint32   `protobuf:"varint,22,opt,name=watchPollingInterval" json:"watchPollingInterval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_configuration_f81253d65f8d0ff7, []int{0}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetIgnores() []string {
	if m != nil {
		return m.Ignores
	}
	return nil
}

func (m *Configuration) GetSymlinkMode() sync.SymlinkMode {
	if m != nil {
		return m.SymlinkMode
	}
	return sync.SymlinkMode_Default
}

func (m *Configuration) GetWatchMode() filesystem.WatchMode {
	if m != nil {
		return m.WatchMode
	}
	return filesystem.WatchMode_Default
}

func (m *Configuration) GetWatchPollingInterval() uint32 {
	if m != nil {
		return m.WatchPollingInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*Configuration)(nil), "session.Configuration")
}

func init() {
	proto.RegisterFile("github.com/havoc-io/mutagen/pkg/session/configuration.proto", fileDescriptor_configuration_f81253d65f8d0ff7)
}

var fileDescriptor_configuration_f81253d65f8d0ff7 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0x31, 0x85, 0x86, 0x28, 0xa4, 0x50, 0xd1, 0x14, 0x93, 0xc9, 0x74, 0xf2, 0x52, 0x09,
	0x6c, 0xe8, 0xd2, 0xb1, 0x53, 0x87, 0x42, 0x71, 0x87, 0xce, 0x8a, 0xaa, 0xc8, 0x47, 0xe4, 0xbb,
	0x20, 0xc9, 0x29, 0x7e, 0xc2, 0xbe, 0x56, 0xa9, 0x92, 0xe0, 0x0c, 0x05, 0x8f, 0xc7, 0x77, 0xdf,
	0x7f, 0xf7, 0xb3, 0x67, 0x0b, 0xb1, 0xed, 0x37, 0x42, 0x53, 0x27, 0x5b, 0x75, 0x20, 0xfd, 0x08,
	0x24, 0xbb, 0x3e, 0x2a, 0x6b, 0x50, 0xee, 0x77, 0x56, 0x06, 0x13, 0x02, 0x10, 0x4a, 0x4d, 0xb8,
	0x05, 0xdb, 0x7b, 0x15, 0x81, 0x50, 0xec, 0x3d, 0x45, 0xe2, 0xb3, 0x13, 0x5c, 0x3f, 0x4d, 0xa5,
	0x6c, 0xc1, 0x99, 0x30, 0x84, 0x68, 0x3a, 0xf9, 0xad, 0xa2, 0x6e, 0x8f, 0x01, 0xeb, 0x6a, 0xf2,
	0xfa, 0x80, 0x5a, 0x86, 0xa1, 0x73, 0x80, 0xbb, 0xa3, 0xf3, 0xf0, 0x93, 0xb1, 0xe5, 0xcb, 0xe5,
	0x33, 0x3c, 0x67, 0x33, 0xb0, 0x48, 0xde, 0x84, 0x3c, 0x2b, 0xae, 0xca, 0x79, 0x73, 0x1e, 0x79,
	0xcd, 0x16, 0x27, 0xf9, 0x8d, 0xbe, 0x4c, 0xbe, 0x28, 0xb2, 0xf2, 0xa6, 0xba, 0x15, 0x7f, 0xa9,
	0xe2, 0x63, 0x04, 0xcd, 0xe5, 0x16, 0xaf, 0xd9, 0x3c, 0xfd, 0x98, 0x94, 0x55, 0x52, 0x56, 0x62,
	0x2c, 0x20, 0x3e, 0xcf, 0xb0, 0x19, 0xf7, 0x78, 0xc5, 0xee, 0xd2, 0xf0, 0x4e, 0xce, 0x01, 0xda,
	0x57, 0x8c, 0xc6, 0x1f, 0x94, 0xcb, 0xef, 0x8b, 0xac, 0x5c, 0x36, 0xff, 0xb2, 0xcd, 0x75, 0x2a,
	0x54, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x9b, 0x9b, 0xec, 0x84, 0x01, 0x00, 0x00,
}
