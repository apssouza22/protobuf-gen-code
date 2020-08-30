// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: book/v1/books.proto

package bookv1

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

// Delete book response
type DeleteBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBookResponse) Reset() {
	*x = DeleteBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookResponse) ProtoMessage() {}

func (x *DeleteBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookResponse.ProtoReflect.Descriptor instead.
func (*DeleteBookResponse) Descriptor() ([]byte, []int) {
	return file_book_v1_books_proto_rawDescGZIP(), []int{0}
}

// A book resource.
type CreateBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique book id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An author of the book.
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// A book title.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateBookResponse) Reset() {
	*x = CreateBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookResponse) ProtoMessage() {}

func (x *CreateBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookResponse.ProtoReflect.Descriptor instead.
func (*CreateBookResponse) Descriptor() ([]byte, []int) {
	return file_book_v1_books_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateBookResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateBookResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// A book resource.
type GetBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique book id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// An author of the book.
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// A book title.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *GetBookResponse) Reset() {
	*x = GetBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResponse) ProtoMessage() {}

func (x *GetBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResponse.ProtoReflect.Descriptor instead.
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return file_book_v1_books_proto_rawDescGZIP(), []int{2}
}

func (x *GetBookResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBookResponse) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetBookResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// Request message for ListBooks method.
type ListBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the shelf which books to list.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
}

func (x *ListBooksRequest) Reset() {
	*x = ListBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksRequest) ProtoMessage() {}

func (x *ListBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksRequest.ProtoReflect.Descriptor instead.
func (*ListBooksRequest) Descriptor() ([]byte, []int) {
	return file_book_v1_books_proto_rawDescGZIP(), []int{3}
}

func (x *ListBooksRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

// Response message to ListBooks method.
type ListBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The books on the shelf.
	Books []*CreateBookResponse `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *ListBooksResponse) Reset() {
	*x = ListBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBooksResponse) ProtoMessage() {}

func (x *ListBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBooksResponse.ProtoReflect.Descriptor instead.
func (*ListBooksResponse) Descriptor() ([]byte, []int) {
	return file_book_v1_books_proto_rawDescGZIP(), []int{4}
}

func (x *ListBooksResponse) GetBooks() []*CreateBookResponse {
	if x != nil {
		return x.Books
	}
	return nil
}

// Request message for CreateBook method.
type CreateBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf on which to create a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// A book resource to create on the shelf.
	Book *CreateBookResponse `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *CreateBookRequest) Reset() {
	*x = CreateBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_books_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookRequest) ProtoMessage() {}

func (x *CreateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_books_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBookRequest.ProtoReflect.Descriptor instead.
func (*CreateBookRequest) Descriptor() ([]byte, []int) {
	return file_book_v1_books_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *CreateBookRequest) GetBook() *CreateBookResponse {
	if x != nil {
		return x.Book
	}
	return nil
}

// Request message for GetBook method.
type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf from which to retrieve a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The ID of the book to retrieve.
	Book int64 `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_books_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_books_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_book_v1_books_proto_rawDescGZIP(), []int{6}
}

func (x *GetBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *GetBookRequest) GetBook() int64 {
	if x != nil {
		return x.Book
	}
	return 0
}

// Request message for DeleteBook method.
type DeleteBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the shelf from which to delete a book.
	Shelf int64 `protobuf:"varint,1,opt,name=shelf,proto3" json:"shelf,omitempty"`
	// The ID of the book to delete.
	Book int64 `protobuf:"varint,2,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *DeleteBookRequest) Reset() {
	*x = DeleteBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_v1_books_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookRequest) ProtoMessage() {}

func (x *DeleteBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_v1_books_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookRequest) Descriptor() ([]byte, []int) {
	return file_book_v1_books_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBookRequest) GetShelf() int64 {
	if x != nil {
		return x.Shelf
	}
	return 0
}

func (x *DeleteBookRequest) GetBook() int64 {
	if x != nil {
		return x.Book
	}
	return 0
}

var File_book_v1_books_proto protoreflect.FileDescriptor

var file_book_v1_books_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x28, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68,
	0x65, 0x6c, 0x66, 0x22, 0x46, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x5a, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x2f, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65,
	0x6c, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x22, 0x3d, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x6c,
	0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x65, 0x6c, 0x66, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x62, 0x6f,
	0x6f, 0x6b, 0x32, 0xaa, 0x02, 0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x19, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x75, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x73, 0x73, 0x6f, 0x75, 0x7a, 0x61, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0e,
	0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x6f, 0x6f, 0x6b, 0x76, 0x31, 0x92, 0x41,
	0x35, 0x12, 0x1e, 0x0a, 0x17, 0x42, 0x6f, 0x6f, 0x6b, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x03, 0x31, 0x2e,
	0x30, 0x1a, 0x0f, 0x6d, 0x79, 0x2e, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x67,
	0x6f, 0x76, 0x2a, 0x02, 0x02, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_v1_books_proto_rawDescOnce sync.Once
	file_book_v1_books_proto_rawDescData = file_book_v1_books_proto_rawDesc
)

func file_book_v1_books_proto_rawDescGZIP() []byte {
	file_book_v1_books_proto_rawDescOnce.Do(func() {
		file_book_v1_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_v1_books_proto_rawDescData)
	})
	return file_book_v1_books_proto_rawDescData
}

var file_book_v1_books_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_book_v1_books_proto_goTypes = []interface{}{
	(*DeleteBookResponse)(nil), // 0: book.v1.DeleteBookResponse
	(*CreateBookResponse)(nil), // 1: book.v1.CreateBookResponse
	(*GetBookResponse)(nil),    // 2: book.v1.GetBookResponse
	(*ListBooksRequest)(nil),   // 3: book.v1.ListBooksRequest
	(*ListBooksResponse)(nil),  // 4: book.v1.ListBooksResponse
	(*CreateBookRequest)(nil),  // 5: book.v1.CreateBookRequest
	(*GetBookRequest)(nil),     // 6: book.v1.GetBookRequest
	(*DeleteBookRequest)(nil),  // 7: book.v1.DeleteBookRequest
}
var file_book_v1_books_proto_depIdxs = []int32{
	1, // 0: book.v1.ListBooksResponse.books:type_name -> book.v1.CreateBookResponse
	1, // 1: book.v1.CreateBookRequest.book:type_name -> book.v1.CreateBookResponse
	3, // 2: book.v1.BookstoreService.ListBooks:input_type -> book.v1.ListBooksRequest
	5, // 3: book.v1.BookstoreService.CreateBook:input_type -> book.v1.CreateBookRequest
	6, // 4: book.v1.BookstoreService.GetBook:input_type -> book.v1.GetBookRequest
	7, // 5: book.v1.BookstoreService.DeleteBook:input_type -> book.v1.DeleteBookRequest
	4, // 6: book.v1.BookstoreService.ListBooks:output_type -> book.v1.ListBooksResponse
	1, // 7: book.v1.BookstoreService.CreateBook:output_type -> book.v1.CreateBookResponse
	2, // 8: book.v1.BookstoreService.GetBook:output_type -> book.v1.GetBookResponse
	0, // 9: book.v1.BookstoreService.DeleteBook:output_type -> book.v1.DeleteBookResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_book_v1_books_proto_init() }
func file_book_v1_books_proto_init() {
	if File_book_v1_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_v1_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookResponse); i {
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
		file_book_v1_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookResponse); i {
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
		file_book_v1_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookResponse); i {
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
		file_book_v1_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksRequest); i {
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
		file_book_v1_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBooksResponse); i {
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
		file_book_v1_books_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBookRequest); i {
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
		file_book_v1_books_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRequest); i {
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
		file_book_v1_books_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookRequest); i {
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
			RawDescriptor: file_book_v1_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_v1_books_proto_goTypes,
		DependencyIndexes: file_book_v1_books_proto_depIdxs,
		MessageInfos:      file_book_v1_books_proto_msgTypes,
	}.Build()
	File_book_v1_books_proto = out.File
	file_book_v1_books_proto_rawDesc = nil
	file_book_v1_books_proto_goTypes = nil
	file_book_v1_books_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookstoreServiceClient is the client API for BookstoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookstoreServiceClient interface {
	// Returns a list of books on a shelf.
	ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error)
	// Creates a new book.
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error)
	// Returns a specific book.
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error)
	// Deletes a book from a shelf.
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error)
}

type bookstoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookstoreServiceClient(cc grpc.ClientConnInterface) BookstoreServiceClient {
	return &bookstoreServiceClient{cc}
}

func (c *bookstoreServiceClient) ListBooks(ctx context.Context, in *ListBooksRequest, opts ...grpc.CallOption) (*ListBooksResponse, error) {
	out := new(ListBooksResponse)
	err := c.cc.Invoke(ctx, "/book.v1.BookstoreService/ListBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*CreateBookResponse, error) {
	out := new(CreateBookResponse)
	err := c.cc.Invoke(ctx, "/book.v1.BookstoreService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*GetBookResponse, error) {
	out := new(GetBookResponse)
	err := c.cc.Invoke(ctx, "/book.v1.BookstoreService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookstoreServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*DeleteBookResponse, error) {
	out := new(DeleteBookResponse)
	err := c.cc.Invoke(ctx, "/book.v1.BookstoreService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookstoreServiceServer is the server API for BookstoreService service.
type BookstoreServiceServer interface {
	// Returns a list of books on a shelf.
	ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error)
	// Creates a new book.
	CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error)
	// Returns a specific book.
	GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error)
	// Deletes a book from a shelf.
	DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error)
}

// UnimplementedBookstoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookstoreServiceServer struct {
}

func (*UnimplementedBookstoreServiceServer) ListBooks(context.Context, *ListBooksRequest) (*ListBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (*UnimplementedBookstoreServiceServer) CreateBook(context.Context, *CreateBookRequest) (*CreateBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (*UnimplementedBookstoreServiceServer) GetBook(context.Context, *GetBookRequest) (*GetBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (*UnimplementedBookstoreServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*DeleteBookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}

func RegisterBookstoreServiceServer(s *grpc.Server, srv BookstoreServiceServer) {
	s.RegisterService(&_BookstoreService_serviceDesc, srv)
}

func _BookstoreService_ListBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBooksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServiceServer).ListBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.v1.BookstoreService/ListBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServiceServer).ListBooks(ctx, req.(*ListBooksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookstoreService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.v1.BookstoreService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookstoreService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.v1.BookstoreService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookstoreService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookstoreServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.v1.BookstoreService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookstoreServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookstoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.v1.BookstoreService",
	HandlerType: (*BookstoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBooks",
			Handler:    _BookstoreService_ListBooks_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookstoreService_CreateBook_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookstoreService_GetBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookstoreService_DeleteBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book/v1/books.proto",
}
