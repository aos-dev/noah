// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: job.proto

package proto

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

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type    uint32 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Job) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Job) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type JobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status  uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JobReply) Reset() {
	*x = JobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobReply) ProtoMessage() {}

func (x *JobReply) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobReply.ProtoReflect.Descriptor instead.
func (*JobReply) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{1}
}

func (x *JobReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobReply) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *JobReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CopyDir struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src     uint32 `protobuf:"varint,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst     uint32 `protobuf:"varint,2,opt,name=dst,proto3" json:"dst,omitempty"`
	SrcPath string `protobuf:"bytes,3,opt,name=src_path,json=srcPath,proto3" json:"src_path,omitempty"`
	DstPath string `protobuf:"bytes,4,opt,name=dst_path,json=dstPath,proto3" json:"dst_path,omitempty"`
}

func (x *CopyDir) Reset() {
	*x = CopyDir{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyDir) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyDir) ProtoMessage() {}

func (x *CopyDir) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyDir.ProtoReflect.Descriptor instead.
func (*CopyDir) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{2}
}

func (x *CopyDir) GetSrc() uint32 {
	if x != nil {
		return x.Src
	}
	return 0
}

func (x *CopyDir) GetDst() uint32 {
	if x != nil {
		return x.Dst
	}
	return 0
}

func (x *CopyDir) GetSrcPath() string {
	if x != nil {
		return x.SrcPath
	}
	return ""
}

func (x *CopyDir) GetDstPath() string {
	if x != nil {
		return x.DstPath
	}
	return ""
}

type CopyFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src     uint32 `protobuf:"varint,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst     uint32 `protobuf:"varint,2,opt,name=dst,proto3" json:"dst,omitempty"`
	SrcPath string `protobuf:"bytes,3,opt,name=src_path,json=srcPath,proto3" json:"src_path,omitempty"`
	DstPath string `protobuf:"bytes,4,opt,name=dst_path,json=dstPath,proto3" json:"dst_path,omitempty"`
}

func (x *CopyFile) Reset() {
	*x = CopyFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyFile) ProtoMessage() {}

func (x *CopyFile) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyFile.ProtoReflect.Descriptor instead.
func (*CopyFile) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{3}
}

func (x *CopyFile) GetSrc() uint32 {
	if x != nil {
		return x.Src
	}
	return 0
}

func (x *CopyFile) GetDst() uint32 {
	if x != nil {
		return x.Dst
	}
	return 0
}

func (x *CopyFile) GetSrcPath() string {
	if x != nil {
		return x.SrcPath
	}
	return ""
}

func (x *CopyFile) GetDstPath() string {
	if x != nil {
		return x.DstPath
	}
	return ""
}

type CopySingleFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src     uint32 `protobuf:"varint,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst     uint32 `protobuf:"varint,2,opt,name=dst,proto3" json:"dst,omitempty"`
	SrcPath string `protobuf:"bytes,3,opt,name=src_path,json=srcPath,proto3" json:"src_path,omitempty"`
	DstPath string `protobuf:"bytes,4,opt,name=dst_path,json=dstPath,proto3" json:"dst_path,omitempty"`
}

func (x *CopySingleFile) Reset() {
	*x = CopySingleFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopySingleFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopySingleFile) ProtoMessage() {}

func (x *CopySingleFile) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopySingleFile.ProtoReflect.Descriptor instead.
func (*CopySingleFile) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{4}
}

func (x *CopySingleFile) GetSrc() uint32 {
	if x != nil {
		return x.Src
	}
	return 0
}

func (x *CopySingleFile) GetDst() uint32 {
	if x != nil {
		return x.Dst
	}
	return 0
}

func (x *CopySingleFile) GetSrcPath() string {
	if x != nil {
		return x.SrcPath
	}
	return ""
}

func (x *CopySingleFile) GetDstPath() string {
	if x != nil {
		return x.DstPath
	}
	return ""
}

type CopyMultipartFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src     uint32 `protobuf:"varint,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst     uint32 `protobuf:"varint,2,opt,name=dst,proto3" json:"dst,omitempty"`
	SrcPath string `protobuf:"bytes,3,opt,name=src_path,json=srcPath,proto3" json:"src_path,omitempty"`
	DstPath string `protobuf:"bytes,4,opt,name=dst_path,json=dstPath,proto3" json:"dst_path,omitempty"`
}

func (x *CopyMultipartFile) Reset() {
	*x = CopyMultipartFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyMultipartFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyMultipartFile) ProtoMessage() {}

func (x *CopyMultipartFile) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyMultipartFile.ProtoReflect.Descriptor instead.
func (*CopyMultipartFile) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{5}
}

func (x *CopyMultipartFile) GetSrc() uint32 {
	if x != nil {
		return x.Src
	}
	return 0
}

func (x *CopyMultipartFile) GetDst() uint32 {
	if x != nil {
		return x.Dst
	}
	return 0
}

func (x *CopyMultipartFile) GetSrcPath() string {
	if x != nil {
		return x.SrcPath
	}
	return ""
}

func (x *CopyMultipartFile) GetDstPath() string {
	if x != nil {
		return x.DstPath
	}
	return ""
}

type CopyMultipart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Src         uint32 `protobuf:"varint,1,opt,name=src,proto3" json:"src,omitempty"`
	Dst         uint32 `protobuf:"varint,2,opt,name=dst,proto3" json:"dst,omitempty"`
	SrcPath     string `protobuf:"bytes,3,opt,name=src_path,json=srcPath,proto3" json:"src_path,omitempty"`
	DstPath     string `protobuf:"bytes,4,opt,name=dst_path,json=dstPath,proto3" json:"dst_path,omitempty"`
	MultipartId string `protobuf:"bytes,5,opt,name=multipart_id,json=multipartId,proto3" json:"multipart_id,omitempty"`
	Size        int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Index       uint32 `protobuf:"varint,7,opt,name=index,proto3" json:"index,omitempty"`
	Offset      int64  `protobuf:"varint,8,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *CopyMultipart) Reset() {
	*x = CopyMultipart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyMultipart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyMultipart) ProtoMessage() {}

func (x *CopyMultipart) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyMultipart.ProtoReflect.Descriptor instead.
func (*CopyMultipart) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{6}
}

func (x *CopyMultipart) GetSrc() uint32 {
	if x != nil {
		return x.Src
	}
	return 0
}

func (x *CopyMultipart) GetDst() uint32 {
	if x != nil {
		return x.Dst
	}
	return 0
}

func (x *CopyMultipart) GetSrcPath() string {
	if x != nil {
		return x.SrcPath
	}
	return ""
}

func (x *CopyMultipart) GetDstPath() string {
	if x != nil {
		return x.DstPath
	}
	return ""
}

func (x *CopyMultipart) GetMultipartId() string {
	if x != nil {
		return x.MultipartId
	}
	return ""
}

func (x *CopyMultipart) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *CopyMultipart) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CopyMultipart) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_job_proto protoreflect.FileDescriptor

var file_job_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62,
	0x22, 0x43, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x63, 0x0a, 0x07, 0x43, 0x6f, 0x70, 0x79, 0x44, 0x69, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x72, 0x63,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x64,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a,
	0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x64, 0x0a, 0x08, 0x43, 0x6f, 0x70, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x64, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x6a,
	0x0a, 0x0e, 0x43, 0x6f, 0x70, 0x79, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73,
	0x72, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x64, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x19, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x6d, 0x0a, 0x11, 0x43, 0x6f,
	0x70, 0x79, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x72,
	0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x64, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19,
	0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0xce, 0x01, 0x0a, 0x0d, 0x43, 0x6f,
	0x70, 0x79, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x72, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x64, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x73,
	0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6f, 0x73, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x6e, 0x6f, 0x61, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_job_proto_rawDescOnce sync.Once
	file_job_proto_rawDescData = file_job_proto_rawDesc
)

func file_job_proto_rawDescGZIP() []byte {
	file_job_proto_rawDescOnce.Do(func() {
		file_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_proto_rawDescData)
	})
	return file_job_proto_rawDescData
}

var file_job_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_job_proto_goTypes = []interface{}{
	(*Job)(nil),               // 0: job.Job
	(*JobReply)(nil),          // 1: job.JobReply
	(*CopyDir)(nil),           // 2: job.CopyDir
	(*CopyFile)(nil),          // 3: job.CopyFile
	(*CopySingleFile)(nil),    // 4: job.CopySingleFile
	(*CopyMultipartFile)(nil), // 5: job.CopyMultipartFile
	(*CopyMultipart)(nil),     // 6: job.CopyMultipart
}
var file_job_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_job_proto_init() }
func file_job_proto_init() {
	if File_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
		file_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobReply); i {
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
		file_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyDir); i {
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
		file_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyFile); i {
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
		file_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopySingleFile); i {
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
		file_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyMultipartFile); i {
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
		file_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyMultipart); i {
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
			RawDescriptor: file_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_job_proto_goTypes,
		DependencyIndexes: file_job_proto_depIdxs,
		MessageInfos:      file_job_proto_msgTypes,
	}.Build()
	File_job_proto = out.File
	file_job_proto_rawDesc = nil
	file_job_proto_goTypes = nil
	file_job_proto_depIdxs = nil
}
