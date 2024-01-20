// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TinySearchBase struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TinySearchBase) Reset()         { *m = TinySearchBase{} }
func (m *TinySearchBase) String() string { return proto.CompactTextString(m) }
func (*TinySearchBase) ProtoMessage()    {}
func (*TinySearchBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

func (m *TinySearchBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TinySearchBase.Unmarshal(m, b)
}
func (m *TinySearchBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TinySearchBase.Marshal(b, m, deterministic)
}
func (m *TinySearchBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TinySearchBase.Merge(m, src)
}
func (m *TinySearchBase) XXX_Size() int {
	return xxx_messageInfo_TinySearchBase.Size(m)
}
func (m *TinySearchBase) XXX_DiscardUnknown() {
	xxx_messageInfo_TinySearchBase.DiscardUnknown(m)
}

var xxx_messageInfo_TinySearchBase proto.InternalMessageInfo

func (m *TinySearchBase) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// message for doc sync request
type DocSyncReq struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	DocId                string   `protobuf:"bytes,2,opt,name=docId,proto3" json:"docId,omitempty"`
	Json                 []byte   `protobuf:"bytes,3,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocSyncReq) Reset()         { *m = DocSyncReq{} }
func (m *DocSyncReq) String() string { return proto.CompactTextString(m) }
func (*DocSyncReq) ProtoMessage()    {}
func (*DocSyncReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{1}
}

func (m *DocSyncReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocSyncReq.Unmarshal(m, b)
}
func (m *DocSyncReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocSyncReq.Marshal(b, m, deterministic)
}
func (m *DocSyncReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocSyncReq.Merge(m, src)
}
func (m *DocSyncReq) XXX_Size() int {
	return xxx_messageInfo_DocSyncReq.Size(m)
}
func (m *DocSyncReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DocSyncReq.DiscardUnknown(m)
}

var xxx_messageInfo_DocSyncReq proto.InternalMessageInfo

func (m *DocSyncReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *DocSyncReq) GetDocId() string {
	if m != nil {
		return m.DocId
	}
	return ""
}

func (m *DocSyncReq) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

// message for doc remove request
type DocRemoveReq struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	DocId                []string `protobuf:"bytes,2,rep,name=docId,proto3" json:"docId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocRemoveReq) Reset()         { *m = DocRemoveReq{} }
func (m *DocRemoveReq) String() string { return proto.CompactTextString(m) }
func (*DocRemoveReq) ProtoMessage()    {}
func (*DocRemoveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{2}
}

func (m *DocRemoveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocRemoveReq.Unmarshal(m, b)
}
func (m *DocRemoveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocRemoveReq.Marshal(b, m, deterministic)
}
func (m *DocRemoveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocRemoveReq.Merge(m, src)
}
func (m *DocRemoveReq) XXX_Size() int {
	return xxx_messageInfo_DocRemoveReq.Size(m)
}
func (m *DocRemoveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DocRemoveReq.DiscardUnknown(m)
}

var xxx_messageInfo_DocRemoveReq proto.InternalMessageInfo

func (m *DocRemoveReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *DocRemoveReq) GetDocId() []string {
	if m != nil {
		return m.DocId
	}
	return nil
}

// message for doc sync response
type DocSyncResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocSyncResp) Reset()         { *m = DocSyncResp{} }
func (m *DocSyncResp) String() string { return proto.CompactTextString(m) }
func (*DocSyncResp) ProtoMessage()    {}
func (*DocSyncResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{3}
}

func (m *DocSyncResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocSyncResp.Unmarshal(m, b)
}
func (m *DocSyncResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocSyncResp.Marshal(b, m, deterministic)
}
func (m *DocSyncResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocSyncResp.Merge(m, src)
}
func (m *DocSyncResp) XXX_Size() int {
	return xxx_messageInfo_DocSyncResp.Size(m)
}
func (m *DocSyncResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DocSyncResp.DiscardUnknown(m)
}

var xxx_messageInfo_DocSyncResp proto.InternalMessageInfo

func (m *DocSyncResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DocSyncResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

// message for doc get
type DocGetReq struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	DocIds               []string `protobuf:"bytes,2,rep,name=docIds,proto3" json:"docIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocGetReq) Reset()         { *m = DocGetReq{} }
func (m *DocGetReq) String() string { return proto.CompactTextString(m) }
func (*DocGetReq) ProtoMessage()    {}
func (*DocGetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{4}
}

func (m *DocGetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocGetReq.Unmarshal(m, b)
}
func (m *DocGetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocGetReq.Marshal(b, m, deterministic)
}
func (m *DocGetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocGetReq.Merge(m, src)
}
func (m *DocGetReq) XXX_Size() int {
	return xxx_messageInfo_DocGetReq.Size(m)
}
func (m *DocGetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DocGetReq.DiscardUnknown(m)
}

var xxx_messageInfo_DocGetReq proto.InternalMessageInfo

func (m *DocGetReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *DocGetReq) GetDocIds() []string {
	if m != nil {
		return m.DocIds
	}
	return nil
}

// message for doc get response
type DocGetResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	JsonByte             [][]byte `protobuf:"bytes,3,rep,name=jsonByte,proto3" json:"jsonByte,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocGetResp) Reset()         { *m = DocGetResp{} }
func (m *DocGetResp) String() string { return proto.CompactTextString(m) }
func (*DocGetResp) ProtoMessage()    {}
func (*DocGetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{5}
}

func (m *DocGetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocGetResp.Unmarshal(m, b)
}
func (m *DocGetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocGetResp.Marshal(b, m, deterministic)
}
func (m *DocGetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocGetResp.Merge(m, src)
}
func (m *DocGetResp) XXX_Size() int {
	return xxx_messageInfo_DocGetResp.Size(m)
}
func (m *DocGetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DocGetResp.DiscardUnknown(m)
}

var xxx_messageInfo_DocGetResp proto.InternalMessageInfo

func (m *DocGetResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DocGetResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *DocGetResp) GetJsonByte() [][]byte {
	if m != nil {
		return m.JsonByte
	}
	return nil
}

// message for doc query request
type DocQueryReq struct {
	Kind                 int32    `protobuf:"varint,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Json                 []byte   `protobuf:"bytes,3,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocQueryReq) Reset()         { *m = DocQueryReq{} }
func (m *DocQueryReq) String() string { return proto.CompactTextString(m) }
func (*DocQueryReq) ProtoMessage()    {}
func (*DocQueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{6}
}

func (m *DocQueryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocQueryReq.Unmarshal(m, b)
}
func (m *DocQueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocQueryReq.Marshal(b, m, deterministic)
}
func (m *DocQueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocQueryReq.Merge(m, src)
}
func (m *DocQueryReq) XXX_Size() int {
	return xxx_messageInfo_DocQueryReq.Size(m)
}
func (m *DocQueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DocQueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_DocQueryReq proto.InternalMessageInfo

func (m *DocQueryReq) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *DocQueryReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *DocQueryReq) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

// message for doc query response
type DocQueryResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	JsonByte             []byte   `protobuf:"bytes,3,opt,name=jsonByte,proto3" json:"jsonByte,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocQueryResp) Reset()         { *m = DocQueryResp{} }
func (m *DocQueryResp) String() string { return proto.CompactTextString(m) }
func (*DocQueryResp) ProtoMessage()    {}
func (*DocQueryResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{7}
}

func (m *DocQueryResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocQueryResp.Unmarshal(m, b)
}
func (m *DocQueryResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocQueryResp.Marshal(b, m, deterministic)
}
func (m *DocQueryResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocQueryResp.Merge(m, src)
}
func (m *DocQueryResp) XXX_Size() int {
	return xxx_messageInfo_DocQueryResp.Size(m)
}
func (m *DocQueryResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DocQueryResp.DiscardUnknown(m)
}

var xxx_messageInfo_DocQueryResp proto.InternalMessageInfo

func (m *DocQueryResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DocQueryResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *DocQueryResp) GetJsonByte() []byte {
	if m != nil {
		return m.JsonByte
	}
	return nil
}

// message for index create
type IndexCreateReq struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexCreateReq) Reset()         { *m = IndexCreateReq{} }
func (m *IndexCreateReq) String() string { return proto.CompactTextString(m) }
func (*IndexCreateReq) ProtoMessage()    {}
func (*IndexCreateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{8}
}

func (m *IndexCreateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexCreateReq.Unmarshal(m, b)
}
func (m *IndexCreateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexCreateReq.Marshal(b, m, deterministic)
}
func (m *IndexCreateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexCreateReq.Merge(m, src)
}
func (m *IndexCreateReq) XXX_Size() int {
	return xxx_messageInfo_IndexCreateReq.Size(m)
}
func (m *IndexCreateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexCreateReq.DiscardUnknown(m)
}

var xxx_messageInfo_IndexCreateReq proto.InternalMessageInfo

func (m *IndexCreateReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

// message for index create response
type IndexCreateResp struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexCreateResp) Reset()         { *m = IndexCreateResp{} }
func (m *IndexCreateResp) String() string { return proto.CompactTextString(m) }
func (*IndexCreateResp) ProtoMessage()    {}
func (*IndexCreateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{9}
}

func (m *IndexCreateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexCreateResp.Unmarshal(m, b)
}
func (m *IndexCreateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexCreateResp.Marshal(b, m, deterministic)
}
func (m *IndexCreateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexCreateResp.Merge(m, src)
}
func (m *IndexCreateResp) XXX_Size() int {
	return xxx_messageInfo_IndexCreateResp.Size(m)
}
func (m *IndexCreateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexCreateResp.DiscardUnknown(m)
}

var xxx_messageInfo_IndexCreateResp proto.InternalMessageInfo

func (m *IndexCreateResp) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *IndexCreateResp) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*TinySearchBase)(nil), "search.TinySearchBase")
	proto.RegisterType((*DocSyncReq)(nil), "search.DocSyncReq")
	proto.RegisterType((*DocRemoveReq)(nil), "search.DocRemoveReq")
	proto.RegisterType((*DocSyncResp)(nil), "search.DocSyncResp")
	proto.RegisterType((*DocGetReq)(nil), "search.DocGetReq")
	proto.RegisterType((*DocGetResp)(nil), "search.DocGetResp")
	proto.RegisterType((*DocQueryReq)(nil), "search.DocQueryReq")
	proto.RegisterType((*DocQueryResp)(nil), "search.DocQueryResp")
	proto.RegisterType((*IndexCreateReq)(nil), "search.IndexCreateReq")
	proto.RegisterType((*IndexCreateResp)(nil), "search.IndexCreateResp")
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_453745cff914010e) }

var fileDescriptor_453745cff914010e = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0xaf, 0xd2, 0x40,
	0x10, 0x4e, 0xdb, 0xf7, 0xfa, 0x60, 0xa8, 0x28, 0x23, 0x81, 0xa6, 0x27, 0xb2, 0x07, 0xc3, 0xa9,
	0x26, 0x18, 0x38, 0x6a, 0x02, 0x24, 0xc8, 0xc1, 0x83, 0xc5, 0x13, 0xf1, 0x52, 0xb6, 0x1b, 0x44,
	0x42, 0xb7, 0x76, 0x17, 0x62, 0xff, 0x1e, 0xff, 0x51, 0xd3, 0xed, 0xb6, 0x16, 0xa9, 0x31, 0x21,
	0xef, 0x36, 0xdf, 0xec, 0xfc, 0xf8, 0x66, 0xe6, 0x5b, 0x70, 0x04, 0x0b, 0x53, 0xfa, 0xcd, 0x4f,
	0x52, 0x2e, 0x39, 0xda, 0x05, 0x22, 0x6f, 0xa0, 0xfb, 0xe5, 0x10, 0x67, 0x1b, 0x85, 0xe6, 0xa1,
	0x60, 0xd8, 0x87, 0x47, 0xc9, 0x8f, 0x2c, 0x76, 0x8d, 0x91, 0x31, 0x6e, 0x07, 0x05, 0x20, 0x1f,
	0x01, 0x96, 0x9c, 0x6e, 0xb2, 0x98, 0x06, 0xec, 0x07, 0xbe, 0x02, 0x4b, 0x86, 0x7b, 0x1d, 0x91,
	0x9b, 0x79, 0x56, 0xc4, 0xe9, 0x3a, 0x72, 0xcd, 0x22, 0x4b, 0x01, 0x44, 0x78, 0xf8, 0x2e, 0x78,
	0xec, 0x5a, 0x23, 0x63, 0xec, 0x04, 0xca, 0x26, 0x33, 0x70, 0x96, 0x9c, 0x06, 0xec, 0xc4, 0x2f,
	0xec, 0xbf, 0xb5, 0xac, 0xaa, 0x16, 0xf9, 0x00, 0x9d, 0x8a, 0x81, 0x48, 0xd0, 0x85, 0x27, 0x71,
	0xa6, 0x94, 0x09, 0xa1, 0x52, 0x5b, 0x41, 0x09, 0x71, 0x00, 0x36, 0x4b, 0xd3, 0x4f, 0x62, 0xaf,
	0xb9, 0x68, 0x44, 0xa6, 0xd0, 0x5e, 0x72, 0xba, 0x62, 0xb2, 0xb9, 0xeb, 0x00, 0x6c, 0xd5, 0x48,
	0xe8, 0xb6, 0x1a, 0x91, 0xad, 0x9a, 0x5c, 0xa5, 0xdd, 0xd3, 0x16, 0x3d, 0x68, 0xe5, 0x73, 0xcf,
	0x33, 0xc9, 0x5c, 0x6b, 0x64, 0x8d, 0x9d, 0xa0, 0xc2, 0x64, 0xa5, 0x66, 0xfa, 0x7c, 0x66, 0x69,
	0x96, 0x93, 0x42, 0x78, 0x38, 0x1e, 0xe2, 0x48, 0x55, 0x7e, 0x0c, 0x94, 0x5d, 0x12, 0x35, 0xff,
	0x10, 0x6d, 0x5a, 0xea, 0x57, 0xb5, 0x54, 0x5d, 0xe8, 0x19, 0x68, 0x1a, 0x57, 0x34, 0x09, 0x74,
	0xd7, 0x71, 0xc4, 0x7e, 0x2e, 0x52, 0x16, 0xca, 0xe6, 0xa3, 0x91, 0x05, 0xbc, 0xbc, 0x8a, 0xb9,
	0x87, 0xc4, 0xe4, 0x97, 0x09, 0x2f, 0x0a, 0x29, 0x6e, 0x58, 0x7a, 0x39, 0x50, 0x86, 0x53, 0x68,
	0x95, 0x83, 0xe1, 0x6b, 0x5f, 0x4b, 0xb8, 0xb6, 0x33, 0xaf, 0x7f, 0xeb, 0x14, 0x09, 0xbe, 0x05,
	0xbb, 0x38, 0x1a, 0xf6, 0x6a, 0xef, 0xc5, 0xed, 0x3d, 0xfc, 0xdb, 0x25, 0x12, 0x9c, 0x29, 0x71,
	0x14, 0xaa, 0xc4, 0x7a, 0xcd, 0x4a, 0xa8, 0x5e, 0xbd, 0x7d, 0x25, 0xc3, 0x09, 0x3c, 0x69, 0x88,
	0x78, 0xf3, 0xfe, 0x8f, 0x9c, 0xf7, 0xd0, 0xa9, 0xad, 0x0a, 0x07, 0x65, 0xcc, 0xf5, 0x8e, 0xbd,
	0x61, 0xa3, 0x5f, 0x24, 0xf3, 0x21, 0xf4, 0x28, 0x3f, 0xf9, 0x92, 0xfa, 0xb2, 0xfa, 0xba, 0x5b,
	0x33, 0xd9, 0xed, 0x6c, 0xf5, 0xb7, 0xdf, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x60, 0xe7, 0x13,
	0xf5, 0xeb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	// doc query, include general query, agg, suggest, etc.
	DocQuery(ctx context.Context, in *DocQueryReq, opts ...grpc.CallOption) (*DocQueryResp, error)
	// doc get
	DocGet(ctx context.Context, in *DocGetReq, opts ...grpc.CallOption) (*DocGetResp, error)
	// doc remove
	DocRemove(ctx context.Context, in *DocRemoveReq, opts ...grpc.CallOption) (*DocSyncResp, error)
	// doc sync
	DocSync(ctx context.Context, in *DocSyncReq, opts ...grpc.CallOption) (*DocSyncResp, error)
	// index create
	IndexCreate(ctx context.Context, in *IndexCreateReq, opts ...grpc.CallOption) (*IndexCreateResp, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) DocQuery(ctx context.Context, in *DocQueryReq, opts ...grpc.CallOption) (*DocQueryResp, error) {
	out := new(DocQueryResp)
	err := c.cc.Invoke(ctx, "/search.SearchService/DocQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) DocGet(ctx context.Context, in *DocGetReq, opts ...grpc.CallOption) (*DocGetResp, error) {
	out := new(DocGetResp)
	err := c.cc.Invoke(ctx, "/search.SearchService/DocGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) DocRemove(ctx context.Context, in *DocRemoveReq, opts ...grpc.CallOption) (*DocSyncResp, error) {
	out := new(DocSyncResp)
	err := c.cc.Invoke(ctx, "/search.SearchService/DocRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) DocSync(ctx context.Context, in *DocSyncReq, opts ...grpc.CallOption) (*DocSyncResp, error) {
	out := new(DocSyncResp)
	err := c.cc.Invoke(ctx, "/search.SearchService/DocSync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) IndexCreate(ctx context.Context, in *IndexCreateReq, opts ...grpc.CallOption) (*IndexCreateResp, error) {
	out := new(IndexCreateResp)
	err := c.cc.Invoke(ctx, "/search.SearchService/IndexCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	// doc query, include general query, agg, suggest, etc.
	DocQuery(context.Context, *DocQueryReq) (*DocQueryResp, error)
	// doc get
	DocGet(context.Context, *DocGetReq) (*DocGetResp, error)
	// doc remove
	DocRemove(context.Context, *DocRemoveReq) (*DocSyncResp, error)
	// doc sync
	DocSync(context.Context, *DocSyncReq) (*DocSyncResp, error)
	// index create
	IndexCreate(context.Context, *IndexCreateReq) (*IndexCreateResp, error)
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) DocQuery(ctx context.Context, req *DocQueryReq) (*DocQueryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocQuery not implemented")
}
func (*UnimplementedSearchServiceServer) DocGet(ctx context.Context, req *DocGetReq) (*DocGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocGet not implemented")
}
func (*UnimplementedSearchServiceServer) DocRemove(ctx context.Context, req *DocRemoveReq) (*DocSyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocRemove not implemented")
}
func (*UnimplementedSearchServiceServer) DocSync(ctx context.Context, req *DocSyncReq) (*DocSyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocSync not implemented")
}
func (*UnimplementedSearchServiceServer) IndexCreate(ctx context.Context, req *IndexCreateReq) (*IndexCreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IndexCreate not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_DocQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocQueryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).DocQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/DocQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).DocQuery(ctx, req.(*DocQueryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_DocGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).DocGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/DocGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).DocGet(ctx, req.(*DocGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_DocRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocRemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).DocRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/DocRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).DocRemove(ctx, req.(*DocRemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_DocSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocSyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).DocSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/DocSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).DocSync(ctx, req.(*DocSyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_IndexCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).IndexCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/IndexCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).IndexCreate(ctx, req.(*IndexCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DocQuery",
			Handler:    _SearchService_DocQuery_Handler,
		},
		{
			MethodName: "DocGet",
			Handler:    _SearchService_DocGet_Handler,
		},
		{
			MethodName: "DocRemove",
			Handler:    _SearchService_DocRemove_Handler,
		},
		{
			MethodName: "DocSync",
			Handler:    _SearchService_DocSync_Handler,
		},
		{
			MethodName: "IndexCreate",
			Handler:    _SearchService_IndexCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}
