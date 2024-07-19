// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: kademlia_grpc.proto

package knp

import (
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

type PingArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *PingArgs) Reset() {
	*x = PingArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingArgs) ProtoMessage() {}

func (x *PingArgs) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingArgs.ProtoReflect.Descriptor instead.
func (*PingArgs) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *PingArgs) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type PingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source   string     `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Contacts []*Contact `protobuf:"bytes,2,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *PingReply) Reset() {
	*x = PingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingReply) ProtoMessage() {}

func (x *PingReply) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingReply.ProtoReflect.Descriptor instead.
func (*PingReply) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *PingReply) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *PingReply) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type StoreArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Val    []byte `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *StoreArgs) Reset() {
	*x = StoreArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreArgs) ProtoMessage() {}

func (x *StoreArgs) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreArgs.ProtoReflect.Descriptor instead.
func (*StoreArgs) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *StoreArgs) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *StoreArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StoreArgs) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type StoreReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreReply) Reset() {
	*x = StoreReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreReply) ProtoMessage() {}

func (x *StoreReply) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreReply.ProtoReflect.Descriptor instead.
func (*StoreReply) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{3}
}

type FindValueArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *FindValueArgs) Reset() {
	*x = FindValueArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindValueArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindValueArgs) ProtoMessage() {}

func (x *FindValueArgs) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindValueArgs.ProtoReflect.Descriptor instead.
func (*FindValueArgs) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *FindValueArgs) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *FindValueArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type FindValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val      []byte     `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
	Contacts []*Contact `protobuf:"bytes,2,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *FindValueReply) Reset() {
	*x = FindValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindValueReply) ProtoMessage() {}

func (x *FindValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindValueReply.ProtoReflect.Descriptor instead.
func (*FindValueReply) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *FindValueReply) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *FindValueReply) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type FindNodeArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *FindNodeArgs) Reset() {
	*x = FindNodeArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNodeArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNodeArgs) ProtoMessage() {}

func (x *FindNodeArgs) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNodeArgs.ProtoReflect.Descriptor instead.
func (*FindNodeArgs) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *FindNodeArgs) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *FindNodeArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type FindNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contacts []*Contact `protobuf:"bytes,1,rep,name=contacts,proto3" json:"contacts,omitempty"`
}

func (x *FindNodeReply) Reset() {
	*x = FindNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNodeReply) ProtoMessage() {}

func (x *FindNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNodeReply.ProtoReflect.Descriptor instead.
func (*FindNodeReply) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *FindNodeReply) GetContacts() []*Contact {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source  string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kademlia_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_kademlia_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_kademlia_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *Contact) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Contact) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_kademlia_grpc_proto protoreflect.FileDescriptor

var file_kademlia_grpc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6b, 0x6e, 0x70, 0x22, 0x22, 0x0a, 0x08, 0x50, 0x69,
	0x6e, 0x67, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x4d,
	0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x6e, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22, 0x47, 0x0a,
	0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x39, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x4c, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x6e, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22, 0x38, 0x0a,
	0x0c, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x39, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6b, 0x6e, 0x70,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x73, 0x22, 0x3b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32,
	0xcb, 0x01, 0x0a, 0x0f, 0x4b, 0x61, 0x64, 0x65, 0x6d, 0x6c, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x6b, 0x6e,
	0x70, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x0e, 0x2e, 0x6b, 0x6e, 0x70,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x0e, 0x2e, 0x6b, 0x6e, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41,
	0x72, 0x67, 0x73, 0x1a, 0x0f, 0x2e, 0x6b, 0x6e, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x12, 0x2e, 0x6b, 0x6e, 0x70, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x13, 0x2e, 0x6b, 0x6e, 0x70, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x08, 0x46, 0x69,
	0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x2e, 0x6b, 0x6e, 0x70, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x12, 0x2e, 0x6b, 0x6e, 0x70, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x07, 0x5a,
	0x05, 0x2f, 0x3b, 0x6b, 0x6e, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kademlia_grpc_proto_rawDescOnce sync.Once
	file_kademlia_grpc_proto_rawDescData = file_kademlia_grpc_proto_rawDesc
)

func file_kademlia_grpc_proto_rawDescGZIP() []byte {
	file_kademlia_grpc_proto_rawDescOnce.Do(func() {
		file_kademlia_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_kademlia_grpc_proto_rawDescData)
	})
	return file_kademlia_grpc_proto_rawDescData
}

var file_kademlia_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_kademlia_grpc_proto_goTypes = []any{
	(*PingArgs)(nil),       // 0: knp.PingArgs
	(*PingReply)(nil),      // 1: knp.PingReply
	(*StoreArgs)(nil),      // 2: knp.StoreArgs
	(*StoreReply)(nil),     // 3: knp.StoreReply
	(*FindValueArgs)(nil),  // 4: knp.FindValueArgs
	(*FindValueReply)(nil), // 5: knp.FindValueReply
	(*FindNodeArgs)(nil),   // 6: knp.FindNodeArgs
	(*FindNodeReply)(nil),  // 7: knp.FindNodeReply
	(*Contact)(nil),        // 8: knp.Contact
}
var file_kademlia_grpc_proto_depIdxs = []int32{
	8, // 0: knp.PingReply.contacts:type_name -> knp.Contact
	8, // 1: knp.FindValueReply.contacts:type_name -> knp.Contact
	8, // 2: knp.FindNodeReply.contacts:type_name -> knp.Contact
	0, // 3: knp.KademliaService.Ping:input_type -> knp.PingArgs
	2, // 4: knp.KademliaService.Store:input_type -> knp.StoreArgs
	4, // 5: knp.KademliaService.FindValue:input_type -> knp.FindValueArgs
	6, // 6: knp.KademliaService.FindNode:input_type -> knp.FindNodeArgs
	1, // 7: knp.KademliaService.Ping:output_type -> knp.PingReply
	3, // 8: knp.KademliaService.Store:output_type -> knp.StoreReply
	5, // 9: knp.KademliaService.FindValue:output_type -> knp.FindValueReply
	7, // 10: knp.KademliaService.FindNode:output_type -> knp.FindNodeReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_kademlia_grpc_proto_init() }
func file_kademlia_grpc_proto_init() {
	if File_kademlia_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kademlia_grpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PingArgs); i {
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
		file_kademlia_grpc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PingReply); i {
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
		file_kademlia_grpc_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*StoreArgs); i {
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
		file_kademlia_grpc_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StoreReply); i {
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
		file_kademlia_grpc_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FindValueArgs); i {
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
		file_kademlia_grpc_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*FindValueReply); i {
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
		file_kademlia_grpc_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*FindNodeArgs); i {
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
		file_kademlia_grpc_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*FindNodeReply); i {
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
		file_kademlia_grpc_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Contact); i {
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
			RawDescriptor: file_kademlia_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kademlia_grpc_proto_goTypes,
		DependencyIndexes: file_kademlia_grpc_proto_depIdxs,
		MessageInfos:      file_kademlia_grpc_proto_msgTypes,
	}.Build()
	File_kademlia_grpc_proto = out.File
	file_kademlia_grpc_proto_rawDesc = nil
	file_kademlia_grpc_proto_goTypes = nil
	file_kademlia_grpc_proto_depIdxs = nil
}
