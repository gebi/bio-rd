// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: github.com/bio-routing/bio-rd/net/api/net.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IP_Version int32

const (
	IP_IPv4 IP_Version = 0
	IP_IPv6 IP_Version = 1
)

// Enum value maps for IP_Version.
var (
	IP_Version_name = map[int32]string{
		0: "IPv4",
		1: "IPv6",
	}
	IP_Version_value = map[string]int32{
		"IPv4": 0,
		"IPv6": 1,
	}
)

func (x IP_Version) Enum() *IP_Version {
	p := new(IP_Version)
	*p = x
	return p
}

func (x IP_Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IP_Version) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_bio_routing_bio_rd_net_api_net_proto_enumTypes[0].Descriptor()
}

func (IP_Version) Type() protoreflect.EnumType {
	return &file_github_com_bio_routing_bio_rd_net_api_net_proto_enumTypes[0]
}

func (x IP_Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IP_Version.Descriptor instead.
func (IP_Version) EnumDescriptor() ([]byte, []int) {
	return file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescGZIP(), []int{1, 0}
}

type Prefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *IP    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Pfxlen  uint32 `protobuf:"varint,2,opt,name=pfxlen,proto3" json:"pfxlen,omitempty"`
}

func (x *Prefix) Reset() {
	*x = Prefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_bio_routing_bio_rd_net_api_net_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prefix) ProtoMessage() {}

func (x *Prefix) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_bio_routing_bio_rd_net_api_net_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prefix.ProtoReflect.Descriptor instead.
func (*Prefix) Descriptor() ([]byte, []int) {
	return file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescGZIP(), []int{0}
}

func (x *Prefix) GetAddress() *IP {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Prefix) GetPfxlen() uint32 {
	if x != nil {
		return x.Pfxlen
	}
	return 0
}

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Higher  uint64     `protobuf:"varint,1,opt,name=higher,proto3" json:"higher,omitempty"`
	Lower   uint64     `protobuf:"varint,2,opt,name=lower,proto3" json:"lower,omitempty"`
	Version IP_Version `protobuf:"varint,3,opt,name=version,proto3,enum=bio.net.IP_Version" json:"version,omitempty"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_bio_routing_bio_rd_net_api_net_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_bio_routing_bio_rd_net_api_net_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescGZIP(), []int{1}
}

func (x *IP) GetHigher() uint64 {
	if x != nil {
		return x.Higher
	}
	return 0
}

func (x *IP) GetLower() uint64 {
	if x != nil {
		return x.Lower
	}
	return 0
}

func (x *IP) GetVersion() IP_Version {
	if x != nil {
		return x.Version
	}
	return IP_IPv4
}

var File_github_com_bio_routing_bio_rd_net_api_net_proto protoreflect.FileDescriptor

var file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6f,
	0x2d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x69, 0x6f, 0x2d, 0x72, 0x64, 0x2f,
	0x6e, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x22, 0x47, 0x0a, 0x06, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6f, 0x2e, 0x6e, 0x65, 0x74, 0x2e,
	0x49, 0x50, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x66, 0x78, 0x6c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x66, 0x78,
	0x6c, 0x65, 0x6e, 0x22, 0x80, 0x01, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69,
	0x67, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x69, 0x67, 0x68,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x62, 0x69, 0x6f, 0x2e,
	0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x1d, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x34, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x50, 0x76, 0x36, 0x10, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6f, 0x2d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x62, 0x69, 0x6f, 0x2d, 0x72, 0x64, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescOnce sync.Once
	file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescData = file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDesc
)

func file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescGZIP() []byte {
	file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescOnce.Do(func() {
		file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescData)
	})
	return file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDescData
}

var file_github_com_bio_routing_bio_rd_net_api_net_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_bio_routing_bio_rd_net_api_net_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_bio_routing_bio_rd_net_api_net_proto_goTypes = []interface{}{
	(IP_Version)(0), // 0: bio.net.IP.Version
	(*Prefix)(nil),  // 1: bio.net.Prefix
	(*IP)(nil),      // 2: bio.net.IP
}
var file_github_com_bio_routing_bio_rd_net_api_net_proto_depIdxs = []int32{
	2, // 0: bio.net.Prefix.address:type_name -> bio.net.IP
	0, // 1: bio.net.IP.version:type_name -> bio.net.IP.Version
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_bio_routing_bio_rd_net_api_net_proto_init() }
func file_github_com_bio_routing_bio_rd_net_api_net_proto_init() {
	if File_github_com_bio_routing_bio_rd_net_api_net_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_bio_routing_bio_rd_net_api_net_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prefix); i {
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
		file_github_com_bio_routing_bio_rd_net_api_net_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
			RawDescriptor: file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_bio_routing_bio_rd_net_api_net_proto_goTypes,
		DependencyIndexes: file_github_com_bio_routing_bio_rd_net_api_net_proto_depIdxs,
		EnumInfos:         file_github_com_bio_routing_bio_rd_net_api_net_proto_enumTypes,
		MessageInfos:      file_github_com_bio_routing_bio_rd_net_api_net_proto_msgTypes,
	}.Build()
	File_github_com_bio_routing_bio_rd_net_api_net_proto = out.File
	file_github_com_bio_routing_bio_rd_net_api_net_proto_rawDesc = nil
	file_github_com_bio_routing_bio_rd_net_api_net_proto_goTypes = nil
	file_github_com_bio_routing_bio_rd_net_api_net_proto_depIdxs = nil
}
