// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capability.proto

package gobgpapi

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

type AddPathMode int32

const (
	AddPathMode_MODE_NONE    AddPathMode = 0
	AddPathMode_MODE_RECEIVE AddPathMode = 1
	AddPathMode_MODE_SEND    AddPathMode = 2
	AddPathMode_MODE_BOTH    AddPathMode = 3
)

var AddPathMode_name = map[int32]string{
	0: "MODE_NONE",
	1: "MODE_RECEIVE",
	2: "MODE_SEND",
	3: "MODE_BOTH",
}
var AddPathMode_value = map[string]int32{
	"MODE_NONE":    0,
	"MODE_RECEIVE": 1,
	"MODE_SEND":    2,
	"MODE_BOTH":    3,
}

func (x AddPathMode) String() string {
	return proto.EnumName(AddPathMode_name, int32(x))
}
func (AddPathMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{0}
}

type MultiProtocolCapability struct {
	Family               *Family  `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiProtocolCapability) Reset()         { *m = MultiProtocolCapability{} }
func (m *MultiProtocolCapability) String() string { return proto.CompactTextString(m) }
func (*MultiProtocolCapability) ProtoMessage()    {}
func (*MultiProtocolCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{0}
}
func (m *MultiProtocolCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiProtocolCapability.Unmarshal(m, b)
}
func (m *MultiProtocolCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiProtocolCapability.Marshal(b, m, deterministic)
}
func (dst *MultiProtocolCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiProtocolCapability.Merge(dst, src)
}
func (m *MultiProtocolCapability) XXX_Size() int {
	return xxx_messageInfo_MultiProtocolCapability.Size(m)
}
func (m *MultiProtocolCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiProtocolCapability.DiscardUnknown(m)
}

var xxx_messageInfo_MultiProtocolCapability proto.InternalMessageInfo

func (m *MultiProtocolCapability) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

type RouteRefreshCapability struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteRefreshCapability) Reset()         { *m = RouteRefreshCapability{} }
func (m *RouteRefreshCapability) String() string { return proto.CompactTextString(m) }
func (*RouteRefreshCapability) ProtoMessage()    {}
func (*RouteRefreshCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{1}
}
func (m *RouteRefreshCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteRefreshCapability.Unmarshal(m, b)
}
func (m *RouteRefreshCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteRefreshCapability.Marshal(b, m, deterministic)
}
func (dst *RouteRefreshCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteRefreshCapability.Merge(dst, src)
}
func (m *RouteRefreshCapability) XXX_Size() int {
	return xxx_messageInfo_RouteRefreshCapability.Size(m)
}
func (m *RouteRefreshCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteRefreshCapability.DiscardUnknown(m)
}

var xxx_messageInfo_RouteRefreshCapability proto.InternalMessageInfo

type CarryingLabelInfoCapability struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CarryingLabelInfoCapability) Reset()         { *m = CarryingLabelInfoCapability{} }
func (m *CarryingLabelInfoCapability) String() string { return proto.CompactTextString(m) }
func (*CarryingLabelInfoCapability) ProtoMessage()    {}
func (*CarryingLabelInfoCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{2}
}
func (m *CarryingLabelInfoCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CarryingLabelInfoCapability.Unmarshal(m, b)
}
func (m *CarryingLabelInfoCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CarryingLabelInfoCapability.Marshal(b, m, deterministic)
}
func (dst *CarryingLabelInfoCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CarryingLabelInfoCapability.Merge(dst, src)
}
func (m *CarryingLabelInfoCapability) XXX_Size() int {
	return xxx_messageInfo_CarryingLabelInfoCapability.Size(m)
}
func (m *CarryingLabelInfoCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_CarryingLabelInfoCapability.DiscardUnknown(m)
}

var xxx_messageInfo_CarryingLabelInfoCapability proto.InternalMessageInfo

type ExtendedNexthopCapabilityTuple struct {
	NlriFamily *Family `protobuf:"bytes,1,opt,name=nlri_family,json=nlriFamily,proto3" json:"nlri_family,omitempty"`
	// Nexthop AFI must be either
	// gobgp.IPv4 or
	// gobgp.IPv6.
	NexthopFamily        *Family  `protobuf:"bytes,2,opt,name=nexthop_family,json=nexthopFamily,proto3" json:"nexthop_family,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtendedNexthopCapabilityTuple) Reset()         { *m = ExtendedNexthopCapabilityTuple{} }
func (m *ExtendedNexthopCapabilityTuple) String() string { return proto.CompactTextString(m) }
func (*ExtendedNexthopCapabilityTuple) ProtoMessage()    {}
func (*ExtendedNexthopCapabilityTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{3}
}
func (m *ExtendedNexthopCapabilityTuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtendedNexthopCapabilityTuple.Unmarshal(m, b)
}
func (m *ExtendedNexthopCapabilityTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtendedNexthopCapabilityTuple.Marshal(b, m, deterministic)
}
func (dst *ExtendedNexthopCapabilityTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedNexthopCapabilityTuple.Merge(dst, src)
}
func (m *ExtendedNexthopCapabilityTuple) XXX_Size() int {
	return xxx_messageInfo_ExtendedNexthopCapabilityTuple.Size(m)
}
func (m *ExtendedNexthopCapabilityTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedNexthopCapabilityTuple.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedNexthopCapabilityTuple proto.InternalMessageInfo

func (m *ExtendedNexthopCapabilityTuple) GetNlriFamily() *Family {
	if m != nil {
		return m.NlriFamily
	}
	return nil
}

func (m *ExtendedNexthopCapabilityTuple) GetNexthopFamily() *Family {
	if m != nil {
		return m.NexthopFamily
	}
	return nil
}

type ExtendedNexthopCapability struct {
	Tuples               []*ExtendedNexthopCapabilityTuple `protobuf:"bytes,1,rep,name=tuples,proto3" json:"tuples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ExtendedNexthopCapability) Reset()         { *m = ExtendedNexthopCapability{} }
func (m *ExtendedNexthopCapability) String() string { return proto.CompactTextString(m) }
func (*ExtendedNexthopCapability) ProtoMessage()    {}
func (*ExtendedNexthopCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{4}
}
func (m *ExtendedNexthopCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtendedNexthopCapability.Unmarshal(m, b)
}
func (m *ExtendedNexthopCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtendedNexthopCapability.Marshal(b, m, deterministic)
}
func (dst *ExtendedNexthopCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedNexthopCapability.Merge(dst, src)
}
func (m *ExtendedNexthopCapability) XXX_Size() int {
	return xxx_messageInfo_ExtendedNexthopCapability.Size(m)
}
func (m *ExtendedNexthopCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedNexthopCapability.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedNexthopCapability proto.InternalMessageInfo

func (m *ExtendedNexthopCapability) GetTuples() []*ExtendedNexthopCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type GracefulRestartCapabilityTuple struct {
	Family               *Family  `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Flags                uint32   `protobuf:"varint,2,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GracefulRestartCapabilityTuple) Reset()         { *m = GracefulRestartCapabilityTuple{} }
func (m *GracefulRestartCapabilityTuple) String() string { return proto.CompactTextString(m) }
func (*GracefulRestartCapabilityTuple) ProtoMessage()    {}
func (*GracefulRestartCapabilityTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{5}
}
func (m *GracefulRestartCapabilityTuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GracefulRestartCapabilityTuple.Unmarshal(m, b)
}
func (m *GracefulRestartCapabilityTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GracefulRestartCapabilityTuple.Marshal(b, m, deterministic)
}
func (dst *GracefulRestartCapabilityTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GracefulRestartCapabilityTuple.Merge(dst, src)
}
func (m *GracefulRestartCapabilityTuple) XXX_Size() int {
	return xxx_messageInfo_GracefulRestartCapabilityTuple.Size(m)
}
func (m *GracefulRestartCapabilityTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_GracefulRestartCapabilityTuple.DiscardUnknown(m)
}

var xxx_messageInfo_GracefulRestartCapabilityTuple proto.InternalMessageInfo

func (m *GracefulRestartCapabilityTuple) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *GracefulRestartCapabilityTuple) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type GracefulRestartCapability struct {
	Flags                uint32                            `protobuf:"varint,1,opt,name=flags,proto3" json:"flags,omitempty"`
	Time                 uint32                            `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Tuples               []*GracefulRestartCapabilityTuple `protobuf:"bytes,3,rep,name=tuples,proto3" json:"tuples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *GracefulRestartCapability) Reset()         { *m = GracefulRestartCapability{} }
func (m *GracefulRestartCapability) String() string { return proto.CompactTextString(m) }
func (*GracefulRestartCapability) ProtoMessage()    {}
func (*GracefulRestartCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{6}
}
func (m *GracefulRestartCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GracefulRestartCapability.Unmarshal(m, b)
}
func (m *GracefulRestartCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GracefulRestartCapability.Marshal(b, m, deterministic)
}
func (dst *GracefulRestartCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GracefulRestartCapability.Merge(dst, src)
}
func (m *GracefulRestartCapability) XXX_Size() int {
	return xxx_messageInfo_GracefulRestartCapability.Size(m)
}
func (m *GracefulRestartCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_GracefulRestartCapability.DiscardUnknown(m)
}

var xxx_messageInfo_GracefulRestartCapability proto.InternalMessageInfo

func (m *GracefulRestartCapability) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *GracefulRestartCapability) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *GracefulRestartCapability) GetTuples() []*GracefulRestartCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type FourOctetASNumberCapability struct {
	As                   uint32   `protobuf:"varint,1,opt,name=as,proto3" json:"as,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FourOctetASNumberCapability) Reset()         { *m = FourOctetASNumberCapability{} }
func (m *FourOctetASNumberCapability) String() string { return proto.CompactTextString(m) }
func (*FourOctetASNumberCapability) ProtoMessage()    {}
func (*FourOctetASNumberCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{7}
}
func (m *FourOctetASNumberCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FourOctetASNumberCapability.Unmarshal(m, b)
}
func (m *FourOctetASNumberCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FourOctetASNumberCapability.Marshal(b, m, deterministic)
}
func (dst *FourOctetASNumberCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FourOctetASNumberCapability.Merge(dst, src)
}
func (m *FourOctetASNumberCapability) XXX_Size() int {
	return xxx_messageInfo_FourOctetASNumberCapability.Size(m)
}
func (m *FourOctetASNumberCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_FourOctetASNumberCapability.DiscardUnknown(m)
}

var xxx_messageInfo_FourOctetASNumberCapability proto.InternalMessageInfo

func (m *FourOctetASNumberCapability) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

type AddPathCapabilityTuple struct {
	Family               *Family     `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Mode                 AddPathMode `protobuf:"varint,2,opt,name=mode,proto3,enum=gobgpapi.AddPathMode" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddPathCapabilityTuple) Reset()         { *m = AddPathCapabilityTuple{} }
func (m *AddPathCapabilityTuple) String() string { return proto.CompactTextString(m) }
func (*AddPathCapabilityTuple) ProtoMessage()    {}
func (*AddPathCapabilityTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{8}
}
func (m *AddPathCapabilityTuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPathCapabilityTuple.Unmarshal(m, b)
}
func (m *AddPathCapabilityTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPathCapabilityTuple.Marshal(b, m, deterministic)
}
func (dst *AddPathCapabilityTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPathCapabilityTuple.Merge(dst, src)
}
func (m *AddPathCapabilityTuple) XXX_Size() int {
	return xxx_messageInfo_AddPathCapabilityTuple.Size(m)
}
func (m *AddPathCapabilityTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPathCapabilityTuple.DiscardUnknown(m)
}

var xxx_messageInfo_AddPathCapabilityTuple proto.InternalMessageInfo

func (m *AddPathCapabilityTuple) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *AddPathCapabilityTuple) GetMode() AddPathMode {
	if m != nil {
		return m.Mode
	}
	return AddPathMode_MODE_NONE
}

type AddPathCapability struct {
	Tuples               []*AddPathCapabilityTuple `protobuf:"bytes,1,rep,name=tuples,proto3" json:"tuples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AddPathCapability) Reset()         { *m = AddPathCapability{} }
func (m *AddPathCapability) String() string { return proto.CompactTextString(m) }
func (*AddPathCapability) ProtoMessage()    {}
func (*AddPathCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{9}
}
func (m *AddPathCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPathCapability.Unmarshal(m, b)
}
func (m *AddPathCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPathCapability.Marshal(b, m, deterministic)
}
func (dst *AddPathCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPathCapability.Merge(dst, src)
}
func (m *AddPathCapability) XXX_Size() int {
	return xxx_messageInfo_AddPathCapability.Size(m)
}
func (m *AddPathCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPathCapability.DiscardUnknown(m)
}

var xxx_messageInfo_AddPathCapability proto.InternalMessageInfo

func (m *AddPathCapability) GetTuples() []*AddPathCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type EnhancedRouteRefreshCapability struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnhancedRouteRefreshCapability) Reset()         { *m = EnhancedRouteRefreshCapability{} }
func (m *EnhancedRouteRefreshCapability) String() string { return proto.CompactTextString(m) }
func (*EnhancedRouteRefreshCapability) ProtoMessage()    {}
func (*EnhancedRouteRefreshCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{10}
}
func (m *EnhancedRouteRefreshCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnhancedRouteRefreshCapability.Unmarshal(m, b)
}
func (m *EnhancedRouteRefreshCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnhancedRouteRefreshCapability.Marshal(b, m, deterministic)
}
func (dst *EnhancedRouteRefreshCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnhancedRouteRefreshCapability.Merge(dst, src)
}
func (m *EnhancedRouteRefreshCapability) XXX_Size() int {
	return xxx_messageInfo_EnhancedRouteRefreshCapability.Size(m)
}
func (m *EnhancedRouteRefreshCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_EnhancedRouteRefreshCapability.DiscardUnknown(m)
}

var xxx_messageInfo_EnhancedRouteRefreshCapability proto.InternalMessageInfo

type LongLivedGracefulRestartCapabilityTuple struct {
	Family               *Family  `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Flags                uint32   `protobuf:"varint,2,opt,name=flags,proto3" json:"flags,omitempty"`
	Time                 uint32   `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongLivedGracefulRestartCapabilityTuple) Reset() {
	*m = LongLivedGracefulRestartCapabilityTuple{}
}
func (m *LongLivedGracefulRestartCapabilityTuple) String() string { return proto.CompactTextString(m) }
func (*LongLivedGracefulRestartCapabilityTuple) ProtoMessage()    {}
func (*LongLivedGracefulRestartCapabilityTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{11}
}
func (m *LongLivedGracefulRestartCapabilityTuple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongLivedGracefulRestartCapabilityTuple.Unmarshal(m, b)
}
func (m *LongLivedGracefulRestartCapabilityTuple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongLivedGracefulRestartCapabilityTuple.Marshal(b, m, deterministic)
}
func (dst *LongLivedGracefulRestartCapabilityTuple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongLivedGracefulRestartCapabilityTuple.Merge(dst, src)
}
func (m *LongLivedGracefulRestartCapabilityTuple) XXX_Size() int {
	return xxx_messageInfo_LongLivedGracefulRestartCapabilityTuple.Size(m)
}
func (m *LongLivedGracefulRestartCapabilityTuple) XXX_DiscardUnknown() {
	xxx_messageInfo_LongLivedGracefulRestartCapabilityTuple.DiscardUnknown(m)
}

var xxx_messageInfo_LongLivedGracefulRestartCapabilityTuple proto.InternalMessageInfo

func (m *LongLivedGracefulRestartCapabilityTuple) GetFamily() *Family {
	if m != nil {
		return m.Family
	}
	return nil
}

func (m *LongLivedGracefulRestartCapabilityTuple) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LongLivedGracefulRestartCapabilityTuple) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type LongLivedGracefulRestartCapability struct {
	Tuples               []*LongLivedGracefulRestartCapabilityTuple `protobuf:"bytes,1,rep,name=tuples,proto3" json:"tuples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *LongLivedGracefulRestartCapability) Reset()         { *m = LongLivedGracefulRestartCapability{} }
func (m *LongLivedGracefulRestartCapability) String() string { return proto.CompactTextString(m) }
func (*LongLivedGracefulRestartCapability) ProtoMessage()    {}
func (*LongLivedGracefulRestartCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{12}
}
func (m *LongLivedGracefulRestartCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongLivedGracefulRestartCapability.Unmarshal(m, b)
}
func (m *LongLivedGracefulRestartCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongLivedGracefulRestartCapability.Marshal(b, m, deterministic)
}
func (dst *LongLivedGracefulRestartCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongLivedGracefulRestartCapability.Merge(dst, src)
}
func (m *LongLivedGracefulRestartCapability) XXX_Size() int {
	return xxx_messageInfo_LongLivedGracefulRestartCapability.Size(m)
}
func (m *LongLivedGracefulRestartCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_LongLivedGracefulRestartCapability.DiscardUnknown(m)
}

var xxx_messageInfo_LongLivedGracefulRestartCapability proto.InternalMessageInfo

func (m *LongLivedGracefulRestartCapability) GetTuples() []*LongLivedGracefulRestartCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type RouteRefreshCiscoCapability struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteRefreshCiscoCapability) Reset()         { *m = RouteRefreshCiscoCapability{} }
func (m *RouteRefreshCiscoCapability) String() string { return proto.CompactTextString(m) }
func (*RouteRefreshCiscoCapability) ProtoMessage()    {}
func (*RouteRefreshCiscoCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{13}
}
func (m *RouteRefreshCiscoCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteRefreshCiscoCapability.Unmarshal(m, b)
}
func (m *RouteRefreshCiscoCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteRefreshCiscoCapability.Marshal(b, m, deterministic)
}
func (dst *RouteRefreshCiscoCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteRefreshCiscoCapability.Merge(dst, src)
}
func (m *RouteRefreshCiscoCapability) XXX_Size() int {
	return xxx_messageInfo_RouteRefreshCiscoCapability.Size(m)
}
func (m *RouteRefreshCiscoCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteRefreshCiscoCapability.DiscardUnknown(m)
}

var xxx_messageInfo_RouteRefreshCiscoCapability proto.InternalMessageInfo

type UnknownCapability struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnknownCapability) Reset()         { *m = UnknownCapability{} }
func (m *UnknownCapability) String() string { return proto.CompactTextString(m) }
func (*UnknownCapability) ProtoMessage()    {}
func (*UnknownCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_capability_5eb0efa1b48300f5, []int{14}
}
func (m *UnknownCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnknownCapability.Unmarshal(m, b)
}
func (m *UnknownCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnknownCapability.Marshal(b, m, deterministic)
}
func (dst *UnknownCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnknownCapability.Merge(dst, src)
}
func (m *UnknownCapability) XXX_Size() int {
	return xxx_messageInfo_UnknownCapability.Size(m)
}
func (m *UnknownCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_UnknownCapability.DiscardUnknown(m)
}

var xxx_messageInfo_UnknownCapability proto.InternalMessageInfo

func (m *UnknownCapability) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UnknownCapability) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*MultiProtocolCapability)(nil), "gobgpapi.MultiProtocolCapability")
	proto.RegisterType((*RouteRefreshCapability)(nil), "gobgpapi.RouteRefreshCapability")
	proto.RegisterType((*CarryingLabelInfoCapability)(nil), "gobgpapi.CarryingLabelInfoCapability")
	proto.RegisterType((*ExtendedNexthopCapabilityTuple)(nil), "gobgpapi.ExtendedNexthopCapabilityTuple")
	proto.RegisterType((*ExtendedNexthopCapability)(nil), "gobgpapi.ExtendedNexthopCapability")
	proto.RegisterType((*GracefulRestartCapabilityTuple)(nil), "gobgpapi.GracefulRestartCapabilityTuple")
	proto.RegisterType((*GracefulRestartCapability)(nil), "gobgpapi.GracefulRestartCapability")
	proto.RegisterType((*FourOctetASNumberCapability)(nil), "gobgpapi.FourOctetASNumberCapability")
	proto.RegisterType((*AddPathCapabilityTuple)(nil), "gobgpapi.AddPathCapabilityTuple")
	proto.RegisterType((*AddPathCapability)(nil), "gobgpapi.AddPathCapability")
	proto.RegisterType((*EnhancedRouteRefreshCapability)(nil), "gobgpapi.EnhancedRouteRefreshCapability")
	proto.RegisterType((*LongLivedGracefulRestartCapabilityTuple)(nil), "gobgpapi.LongLivedGracefulRestartCapabilityTuple")
	proto.RegisterType((*LongLivedGracefulRestartCapability)(nil), "gobgpapi.LongLivedGracefulRestartCapability")
	proto.RegisterType((*RouteRefreshCiscoCapability)(nil), "gobgpapi.RouteRefreshCiscoCapability")
	proto.RegisterType((*UnknownCapability)(nil), "gobgpapi.UnknownCapability")
	proto.RegisterEnum("gobgpapi.AddPathMode", AddPathMode_name, AddPathMode_value)
}

func init() { proto.RegisterFile("capability.proto", fileDescriptor_capability_5eb0efa1b48300f5) }

var fileDescriptor_capability_5eb0efa1b48300f5 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x88, 0x60, 0xd2, 0x44, 0xee, 0x0a, 0x4a, 0x4a, 0xd4, 0x28, 0xda, 0x0b, 0x01,
	0x89, 0x48, 0x0d, 0x07, 0xb8, 0x20, 0x51, 0x52, 0x17, 0x22, 0xe5, 0xa3, 0x72, 0x0b, 0x37, 0x54,
	0x36, 0xf6, 0xc6, 0x59, 0xb1, 0xde, 0xb5, 0xec, 0x75, 0x69, 0x0e, 0x9c, 0xb9, 0xf0, 0xa3, 0x91,
	0x3f, 0x62, 0x9b, 0x54, 0x6e, 0x2b, 0xa4, 0xde, 0x66, 0xbc, 0x33, 0x6f, 0xde, 0x9b, 0xb7, 0x6b,
	0xd0, 0x2d, 0xe2, 0x91, 0x05, 0xe3, 0x4c, 0xad, 0x07, 0x9e, 0x2f, 0x95, 0x44, 0x8f, 0x1c, 0xb9,
	0x70, 0x3c, 0xe2, 0xb1, 0xe7, 0x8d, 0x38, 0x4a, 0x3e, 0xe3, 0x11, 0x3c, 0x9b, 0x86, 0x5c, 0xb1,
	0xd3, 0x28, 0xb3, 0x24, 0x1f, 0x65, 0x7d, 0xa8, 0x0f, 0xf5, 0x25, 0x71, 0x19, 0x5f, 0xb7, 0xb5,
	0x9e, 0xd6, 0x6f, 0x0c, 0xf5, 0xc1, 0x06, 0x62, 0x70, 0x12, 0x7f, 0x37, 0xd3, 0x73, 0xdc, 0x86,
	0x3d, 0x53, 0x86, 0x8a, 0x9a, 0x74, 0xe9, 0xd3, 0x60, 0x95, 0x63, 0xe0, 0x03, 0xe8, 0x8c, 0x88,
	0xef, 0xaf, 0x99, 0x70, 0x26, 0x64, 0x41, 0xf9, 0x58, 0x2c, 0x65, 0xe1, 0xf8, 0x8f, 0x06, 0x5d,
	0xe3, 0x4a, 0x51, 0x61, 0x53, 0x7b, 0x46, 0xaf, 0xd4, 0x4a, 0x7a, 0xf9, 0xe9, 0x79, 0xe8, 0x71,
	0x8a, 0x0e, 0xa1, 0x21, 0xb8, 0xcf, 0x2e, 0x6e, 0xa1, 0x02, 0x51, 0x51, 0x12, 0xa3, 0xb7, 0xd0,
	0x12, 0x09, 0xd8, 0xa6, 0xab, 0x52, 0xd2, 0xd5, 0x4c, 0xeb, 0x92, 0x14, 0x7f, 0x83, 0xfd, 0x52,
	0x36, 0xe8, 0x03, 0xd4, 0x55, 0xc4, 0x28, 0x68, 0x6b, 0xbd, 0x6a, 0xbf, 0x31, 0xec, 0xe7, 0x68,
	0x37, 0x4b, 0x30, 0xd3, 0x3e, 0xfc, 0x1d, 0xba, 0x9f, 0x7c, 0x62, 0xd1, 0x65, 0xc8, 0x4d, 0x1a,
	0x28, 0xe2, 0xab, 0x6d, 0xb1, 0x77, 0x5e, 0x39, 0x7a, 0x02, 0x0f, 0x97, 0x9c, 0x38, 0x41, 0x2c,
	0xad, 0x69, 0x26, 0x09, 0xfe, 0xad, 0xc1, 0x7e, 0xe9, 0x88, 0xbc, 0x47, 0x2b, 0xf4, 0x20, 0x04,
	0x35, 0xc5, 0x5c, 0x9a, 0x02, 0xc5, 0x71, 0x41, 0x6b, 0x75, 0x5b, 0xeb, 0xcd, 0x0a, 0x32, 0xad,
	0xaf, 0xa1, 0x73, 0x22, 0x43, 0x7f, 0x6e, 0x29, 0xaa, 0x8e, 0xce, 0x66, 0xa1, 0xbb, 0xa0, 0x7e,
	0x81, 0x4a, 0x0b, 0x2a, 0x64, 0xc3, 0xa3, 0x42, 0x02, 0xec, 0xc2, 0xde, 0x91, 0x6d, 0x9f, 0x12,
	0xb5, 0xfa, 0xff, 0x95, 0xbc, 0x84, 0x9a, 0x2b, 0xed, 0x44, 0x48, 0x6b, 0xf8, 0x34, 0xaf, 0x4b,
	0x91, 0xa7, 0xd2, 0xa6, 0x66, 0x5c, 0x82, 0xa7, 0xb0, 0x7b, 0x6d, 0x1c, 0x7a, 0xb7, 0x65, 0x70,
	0xef, 0x1a, 0x42, 0x99, 0xd8, 0x1e, 0x74, 0x0d, 0xb1, 0x22, 0xc2, 0xa2, 0x76, 0xc9, 0x3b, 0xf8,
	0x05, 0x2f, 0x26, 0x52, 0x38, 0x13, 0x76, 0x49, 0xed, 0xfb, 0xbd, 0x03, 0x99, 0x9f, 0xd5, 0xdc,
	0x4f, 0x2c, 0x01, 0xdf, 0x3e, 0x1e, 0x8d, 0xb7, 0x16, 0x70, 0x98, 0x4f, 0xbe, 0x23, 0xf9, 0x6c,
	0x23, 0x07, 0xd0, 0xf9, 0x67, 0x13, 0x2c, 0xb0, 0x8a, 0xef, 0xfe, 0x3d, 0xec, 0x7e, 0x11, 0x3f,
	0x84, 0xfc, 0x29, 0x0a, 0xe3, 0x11, 0xd4, 0xac, 0xc8, 0xbf, 0xe4, 0x56, 0xc4, 0x71, 0x24, 0xf1,
	0x92, 0xf0, 0x30, 0x31, 0x75, 0xc7, 0x4c, 0x92, 0x57, 0x13, 0x68, 0x14, 0x3c, 0x45, 0x4d, 0x78,
	0x3c, 0x9d, 0x1f, 0x1b, 0x17, 0xb3, 0xf9, 0xcc, 0xd0, 0x1f, 0x20, 0x1d, 0x76, 0xe2, 0xd4, 0x34,
	0x46, 0xc6, 0xf8, 0xab, 0xa1, 0x6b, 0x59, 0xc1, 0x99, 0x31, 0x3b, 0xd6, 0x2b, 0x59, 0xfa, 0x71,
	0x7e, 0xfe, 0x59, 0xaf, 0x2e, 0xea, 0xf1, 0x9f, 0xf0, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x3b, 0xec, 0xd3, 0x4d, 0x34, 0x05, 0x00, 0x00,
}
