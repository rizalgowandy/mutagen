// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/havoc-io/mutagen/pkg/filesystem/watch.proto

package filesystem

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

// WatchMode specifies the mode for filesystem watching.
type WatchMode int32

const (
	// Default represents an unspecified watch mode. It should be converted to
	// one of the following values based on the desired default behavior.
	WatchMode_Default WatchMode = 0
	// RecursiveHome specifies that native recursive watching should be used to
	// monitor paths on systems that support it if those paths fall under the
	// home directory. In these cases, a watch on the entire home directory is
	// established and filtered for events pertaining to the specified path. On
	// all other systems and for all other paths, poll-based watching is used.
	WatchMode_RecursiveHome WatchMode = 1
	// Poll specifies that only poll-based watching should be performed.
	WatchMode_Poll WatchMode = 2
)

var WatchMode_name = map[int32]string{
	0: "Default",
	1: "RecursiveHome",
	2: "Poll",
}
var WatchMode_value = map[string]int32{
	"Default":       0,
	"RecursiveHome": 1,
	"Poll":          2,
}

func (x WatchMode) String() string {
	return proto.EnumName(WatchMode_name, int32(x))
}
func (WatchMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_watch_8ae0081c18c36963, []int{0}
}

func init() {
	proto.RegisterEnum("filesystem.WatchMode", WatchMode_name, WatchMode_value)
}

func init() {
	proto.RegisterFile("github.com/havoc-io/mutagen/pkg/filesystem/watch.proto", fileDescriptor_watch_8ae0081c18c36963)
}

var fileDescriptor_watch_8ae0081c18c36963 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x48, 0x2c, 0xcb, 0x4f, 0xd6, 0xcd, 0xcc, 0xd7,
	0xcf, 0x2d, 0x2d, 0x49, 0x4c, 0x4f, 0xcd, 0xd3, 0x2f, 0xc8, 0x4e, 0xd7, 0x4f, 0xcb, 0xcc, 0x49,
	0x2d, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0xd5, 0x2f, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88, 0x6b, 0x99, 0x72, 0x71, 0x86, 0x83, 0xa4, 0x7c, 0xf3, 0x53,
	0x52, 0x85, 0xb8, 0xb9, 0xd8, 0x5d, 0x52, 0xd3, 0x12, 0x4b, 0x73, 0x4a, 0x04, 0x18, 0x84, 0x04,
	0xb9, 0x78, 0x83, 0x52, 0x93, 0x4b, 0x8b, 0x8a, 0x33, 0xcb, 0x52, 0x3d, 0xf2, 0x73, 0x53, 0x05,
	0x18, 0x85, 0x38, 0xb8, 0x58, 0x02, 0xf2, 0x73, 0x72, 0x04, 0x98, 0x92, 0xd8, 0xc0, 0x26, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x8b, 0x02, 0x60, 0x83, 0x00, 0x00, 0x00,
}
