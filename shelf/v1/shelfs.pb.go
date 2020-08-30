// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: apssouza/shelf/v1/shelfs.proto

package shelfv1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// DeleteShelf response
type DeleteShelfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteShelfResponse) Reset() {
	*x = DeleteShelfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShelfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelfResponse) ProtoMessage() {}

func (x *DeleteShelfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShelfResponse.ProtoReflect.Descriptor instead.
func (*DeleteShelfResponse) Descriptor() ([]byte, []int) {
	return file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP(), []int{0}
}

// ListShelves request
type ListShelvesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListShelvesRequest) Reset() {
	*x = ListShelvesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShelvesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesRequest) ProtoMessage() {}

func (x *ListShelvesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShelvesRequest.ProtoReflect.Descriptor instead.
func (*ListShelvesRequest) Descriptor() ([]byte, []int) {
	return file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP(), []int{1}
}

// A shelf resource.
type GetShelfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique shelf id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// A theme of the shelf (fiction, poetry, etc).
	Theme string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
}

func (x *GetShelfResponse) Reset() {
	*x = GetShelfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelfResponse) ProtoMessage() {}

func (x *GetShelfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelfResponse.ProtoReflect.Descriptor instead.
func (*GetShelfResponse) Descriptor() ([]byte, []int) {
	return file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP(), []int{2}
}

func (x *GetShelfResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetShelfResponse) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

// A shelf resource.
type CreateShelfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique shelf id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// A theme of the shelf (fiction, poetry, etc).
	Theme string `protobuf:"bytes,2,opt,name=theme,proto3" json:"theme,omitempty"`
}

func (x *CreateShelfResponse) Reset() {
	*x = CreateShelfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfResponse) ProtoMessage() {}

func (x *CreateShelfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShelfResponse.ProtoReflect.Descriptor instead.
func (*CreateShelfResponse) Descriptor() ([]byte, []int) {
	return file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP(), []int{3}
}

func (x *CreateShelfResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateShelfResponse) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

// Response to ListShelves call.
type ListShelvesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Shelves in the bookstore.
	Shelves []*CreateShelfResponse `protobuf:"bytes,1,rep,name=shelves,proto3" json:"shelves,omitempty"`
}

func (x *ListShelvesResponse) Reset() {
	*x = ListShelvesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShelvesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShelvesResponse) ProtoMessage() {}

func (x *ListShelvesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShelvesResponse.ProtoReflect.Descriptor instead.
func (*ListShelvesResponse) Descriptor() ([]byte, []int) {
	return file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP(), []int{4}
}

func (x *ListShelvesResponse) GetShelves() []*CreateShelfResponse {
	if x != nil {
		return x.Shelves
	}
	return nil
}

// Request message for CreateShelf method.
type CreateShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The shelf resource to create.
	Shelf *CreateShelfResponse `protobuf:"bytes,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *CreateShelfRequest) Reset() {
	*x = CreateShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShelfRequest) ProtoMessage() {}

func (x *CreateShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShelfRequest.ProtoReflect.Descriptor instead.
func (*CreateShelfRequest) Descriptor() ([]byte, []int) {
	return file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP(), []int{5}
}

func (x *CreateShelfRequest) GetShelf() *CreateShelfResponse {
	if x != nil {
		return x.Shelf
	}
	return nil
}

// Request message for GetShelf method.
type GetShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf resource to retrieve.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *GetShelfRequest) Reset() {
	*x = GetShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShelfRequest) ProtoMessage() {}

func (x *GetShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShelfRequest.ProtoReflect.Descriptor instead.
func (*GetShelfRequest) Descriptor() ([]byte, []int) {
	return file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP(), []int{6}
}

func (x *GetShelfRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

// Request message for DeleteShelf method.
type DeleteShelfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf to delete.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *DeleteShelfRequest) Reset() {
	*x = DeleteShelfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteShelfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteShelfRequest) ProtoMessage() {}

func (x *DeleteShelfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apssouza_shelf_v1_shelfs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteShelfRequest.ProtoReflect.Descriptor instead.
func (*DeleteShelfRequest) Descriptor() ([]byte, []int) {
	return file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteShelfRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

var File_apssouza_shelf_v1_shelfs_proto protoreflect.FileDescriptor

var file_apssouza_shelf_v1_shelfs_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a, 0x61, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65,
	0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x38, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x22, 0x57, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40,
	0x0a, 0x07, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x73, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73,
	0x22, 0x52, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a, 0x61,
	0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x73,
	0x68, 0x65, 0x6c, 0x66, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x22, 0x2a, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x32, 0x85, 0x03, 0x0a, 0x0c, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x73, 0x73,
	0x6f, 0x75, 0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x76, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x73, 0x73,
	0x6f, 0x75, 0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a,
	0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68,
	0x65, 0x6c, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x73,
	0x73, 0x6f, 0x75, 0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x12, 0x25, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75,
	0x7a, 0x61, 0x2e, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x79, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a,
	0x61, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x10, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x76, 0x31, 0x92, 0x41, 0x36, 0x12, 0x1f, 0x0a, 0x18, 0x53, 0x68, 0x65, 0x6c, 0x66, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x1a, 0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f,
	0x73, 0x74, 0x3a, 0x35, 0x30, 0x30, 0x35, 0x32, 0x2a, 0x02, 0x02, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apssouza_shelf_v1_shelfs_proto_rawDescOnce sync.Once
	file_apssouza_shelf_v1_shelfs_proto_rawDescData = file_apssouza_shelf_v1_shelfs_proto_rawDesc
)

func file_apssouza_shelf_v1_shelfs_proto_rawDescGZIP() []byte {
	file_apssouza_shelf_v1_shelfs_proto_rawDescOnce.Do(func() {
		file_apssouza_shelf_v1_shelfs_proto_rawDescData = protoimpl.X.CompressGZIP(file_apssouza_shelf_v1_shelfs_proto_rawDescData)
	})
	return file_apssouza_shelf_v1_shelfs_proto_rawDescData
}

var file_apssouza_shelf_v1_shelfs_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_apssouza_shelf_v1_shelfs_proto_goTypes = []interface{}{
	(*DeleteShelfResponse)(nil), // 0: apssouza.shelf.v1.DeleteShelfResponse
	(*ListShelvesRequest)(nil),  // 1: apssouza.shelf.v1.ListShelvesRequest
	(*GetShelfResponse)(nil),    // 2: apssouza.shelf.v1.GetShelfResponse
	(*CreateShelfResponse)(nil), // 3: apssouza.shelf.v1.CreateShelfResponse
	(*ListShelvesResponse)(nil), // 4: apssouza.shelf.v1.ListShelvesResponse
	(*CreateShelfRequest)(nil),  // 5: apssouza.shelf.v1.CreateShelfRequest
	(*GetShelfRequest)(nil),     // 6: apssouza.shelf.v1.GetShelfRequest
	(*DeleteShelfRequest)(nil),  // 7: apssouza.shelf.v1.DeleteShelfRequest
}
var file_apssouza_shelf_v1_shelfs_proto_depIdxs = []int32{
	3, // 0: apssouza.shelf.v1.ListShelvesResponse.shelves:type_name -> apssouza.shelf.v1.CreateShelfResponse
	3, // 1: apssouza.shelf.v1.CreateShelfRequest.shelf:type_name -> apssouza.shelf.v1.CreateShelfResponse
	1, // 2: apssouza.shelf.v1.ShelfService.ListShelves:input_type -> apssouza.shelf.v1.ListShelvesRequest
	5, // 3: apssouza.shelf.v1.ShelfService.CreateShelf:input_type -> apssouza.shelf.v1.CreateShelfRequest
	6, // 4: apssouza.shelf.v1.ShelfService.GetShelf:input_type -> apssouza.shelf.v1.GetShelfRequest
	7, // 5: apssouza.shelf.v1.ShelfService.DeleteShelf:input_type -> apssouza.shelf.v1.DeleteShelfRequest
	4, // 6: apssouza.shelf.v1.ShelfService.ListShelves:output_type -> apssouza.shelf.v1.ListShelvesResponse
	3, // 7: apssouza.shelf.v1.ShelfService.CreateShelf:output_type -> apssouza.shelf.v1.CreateShelfResponse
	2, // 8: apssouza.shelf.v1.ShelfService.GetShelf:output_type -> apssouza.shelf.v1.GetShelfResponse
	0, // 9: apssouza.shelf.v1.ShelfService.DeleteShelf:output_type -> apssouza.shelf.v1.DeleteShelfResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apssouza_shelf_v1_shelfs_proto_init() }
func file_apssouza_shelf_v1_shelfs_proto_init() {
	if File_apssouza_shelf_v1_shelfs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apssouza_shelf_v1_shelfs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShelfResponse); i {
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
		file_apssouza_shelf_v1_shelfs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShelvesRequest); i {
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
		file_apssouza_shelf_v1_shelfs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelfResponse); i {
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
		file_apssouza_shelf_v1_shelfs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShelfResponse); i {
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
		file_apssouza_shelf_v1_shelfs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShelvesResponse); i {
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
		file_apssouza_shelf_v1_shelfs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShelfRequest); i {
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
		file_apssouza_shelf_v1_shelfs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShelfRequest); i {
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
		file_apssouza_shelf_v1_shelfs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteShelfRequest); i {
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
			RawDescriptor: file_apssouza_shelf_v1_shelfs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apssouza_shelf_v1_shelfs_proto_goTypes,
		DependencyIndexes: file_apssouza_shelf_v1_shelfs_proto_depIdxs,
		MessageInfos:      file_apssouza_shelf_v1_shelfs_proto_msgTypes,
	}.Build()
	File_apssouza_shelf_v1_shelfs_proto = out.File
	file_apssouza_shelf_v1_shelfs_proto_rawDesc = nil
	file_apssouza_shelf_v1_shelfs_proto_goTypes = nil
	file_apssouza_shelf_v1_shelfs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShelfServiceClient is the client API for ShelfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShelfServiceClient interface {
	// Returns a list of all shelves in the bookstore.
	ListShelves(ctx context.Context, in *ListShelvesRequest, opts ...grpc.CallOption) (*ListShelvesResponse, error)
	// Creates a new shelf in the bookstore.
	CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*CreateShelfResponse, error)
	// Returns a specific bookstore shelf.
	GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*GetShelfResponse, error)
	// Deletes a shelf, including all books that are stored on the shelf.
	DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*DeleteShelfResponse, error)
}

type shelfServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShelfServiceClient(cc grpc.ClientConnInterface) ShelfServiceClient {
	return &shelfServiceClient{cc}
}

func (c *shelfServiceClient) ListShelves(ctx context.Context, in *ListShelvesRequest, opts ...grpc.CallOption) (*ListShelvesResponse, error) {
	out := new(ListShelvesResponse)
	err := c.cc.Invoke(ctx, "/apssouza.shelf.v1.ShelfService/ListShelves", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelfServiceClient) CreateShelf(ctx context.Context, in *CreateShelfRequest, opts ...grpc.CallOption) (*CreateShelfResponse, error) {
	out := new(CreateShelfResponse)
	err := c.cc.Invoke(ctx, "/apssouza.shelf.v1.ShelfService/CreateShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelfServiceClient) GetShelf(ctx context.Context, in *GetShelfRequest, opts ...grpc.CallOption) (*GetShelfResponse, error) {
	out := new(GetShelfResponse)
	err := c.cc.Invoke(ctx, "/apssouza.shelf.v1.ShelfService/GetShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shelfServiceClient) DeleteShelf(ctx context.Context, in *DeleteShelfRequest, opts ...grpc.CallOption) (*DeleteShelfResponse, error) {
	out := new(DeleteShelfResponse)
	err := c.cc.Invoke(ctx, "/apssouza.shelf.v1.ShelfService/DeleteShelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShelfServiceServer is the server API for ShelfService service.
type ShelfServiceServer interface {
	// Returns a list of all shelves in the bookstore.
	ListShelves(context.Context, *ListShelvesRequest) (*ListShelvesResponse, error)
	// Creates a new shelf in the bookstore.
	CreateShelf(context.Context, *CreateShelfRequest) (*CreateShelfResponse, error)
	// Returns a specific bookstore shelf.
	GetShelf(context.Context, *GetShelfRequest) (*GetShelfResponse, error)
	// Deletes a shelf, including all books that are stored on the shelf.
	DeleteShelf(context.Context, *DeleteShelfRequest) (*DeleteShelfResponse, error)
}

// UnimplementedShelfServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShelfServiceServer struct {
}

func (*UnimplementedShelfServiceServer) ListShelves(context.Context, *ListShelvesRequest) (*ListShelvesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListShelves not implemented")
}
func (*UnimplementedShelfServiceServer) CreateShelf(context.Context, *CreateShelfRequest) (*CreateShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShelf not implemented")
}
func (*UnimplementedShelfServiceServer) GetShelf(context.Context, *GetShelfRequest) (*GetShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShelf not implemented")
}
func (*UnimplementedShelfServiceServer) DeleteShelf(context.Context, *DeleteShelfRequest) (*DeleteShelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteShelf not implemented")
}

func RegisterShelfServiceServer(s *grpc.Server, srv ShelfServiceServer) {
	s.RegisterService(&_ShelfService_serviceDesc, srv)
}

func _ShelfService_ListShelves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListShelvesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).ListShelves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apssouza.shelf.v1.ShelfService/ListShelves",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).ListShelves(ctx, req.(*ListShelvesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelfService_CreateShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).CreateShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apssouza.shelf.v1.ShelfService/CreateShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).CreateShelf(ctx, req.(*CreateShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelfService_GetShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).GetShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apssouza.shelf.v1.ShelfService/GetShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).GetShelf(ctx, req.(*GetShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShelfService_DeleteShelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteShelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShelfServiceServer).DeleteShelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apssouza.shelf.v1.ShelfService/DeleteShelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShelfServiceServer).DeleteShelf(ctx, req.(*DeleteShelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShelfService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apssouza.shelf.v1.ShelfService",
	HandlerType: (*ShelfServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListShelves",
			Handler:    _ShelfService_ListShelves_Handler,
		},
		{
			MethodName: "CreateShelf",
			Handler:    _ShelfService_CreateShelf_Handler,
		},
		{
			MethodName: "GetShelf",
			Handler:    _ShelfService_GetShelf_Handler,
		},
		{
			MethodName: "DeleteShelf",
			Handler:    _ShelfService_DeleteShelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apssouza/shelf/v1/shelfs.proto",
}
