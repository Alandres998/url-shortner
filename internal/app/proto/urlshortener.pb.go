// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.23.4
// source: urlshortener.proto

package proto

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

type CreateShortURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *CreateShortURLRequest) Reset() {
	*x = CreateShortURLRequest{}
	mi := &file_urlshortener_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShortURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLRequest) ProtoMessage() {}

func (x *CreateShortURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLRequest.ProtoReflect.Descriptor instead.
func (*CreateShortURLRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShortURLRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type CreateShortURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *CreateShortURLResponse) Reset() {
	*x = CreateShortURLResponse{}
	mi := &file_urlshortener_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateShortURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortURLResponse) ProtoMessage() {}

func (x *CreateShortURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortURLResponse.ProtoReflect.Descriptor instead.
func (*CreateShortURLResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortURLResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetOriginalURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
}

func (x *GetOriginalURLRequest) Reset() {
	*x = GetOriginalURLRequest{}
	mi := &file_urlshortener_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOriginalURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalURLRequest) ProtoMessage() {}

func (x *GetOriginalURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalURLRequest.ProtoReflect.Descriptor instead.
func (*GetOriginalURLRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{2}
}

func (x *GetOriginalURLRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

type GetOriginalURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginalUrl string `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *GetOriginalURLResponse) Reset() {
	*x = GetOriginalURLResponse{}
	mi := &file_urlshortener_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOriginalURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginalURLResponse) ProtoMessage() {}

func (x *GetOriginalURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginalURLResponse.ProtoReflect.Descriptor instead.
func (*GetOriginalURLResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{3}
}

func (x *GetOriginalURLResponse) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type GetUserURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserURLsRequest) Reset() {
	*x = GetUserURLsRequest{}
	mi := &file_urlshortener_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserURLsRequest) ProtoMessage() {}

func (x *GetUserURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserURLsRequest.ProtoReflect.Descriptor instead.
func (*GetUserURLsRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserURLsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*UserURL `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *GetUserURLsResponse) Reset() {
	*x = GetUserURLsResponse{}
	mi := &file_urlshortener_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserURLsResponse) ProtoMessage() {}

func (x *GetUserURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserURLsResponse.ProtoReflect.Descriptor instead.
func (*GetUserURLsResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserURLsResponse) GetUrls() []*UserURL {
	if x != nil {
		return x.Urls
	}
	return nil
}

type UserURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrl    string `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"`
	OriginalUrl string `protobuf:"bytes,2,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"`
}

func (x *UserURL) Reset() {
	*x = UserURL{}
	mi := &file_urlshortener_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserURL) ProtoMessage() {}

func (x *UserURL) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserURL.ProtoReflect.Descriptor instead.
func (*UserURL) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{6}
}

func (x *UserURL) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *UserURL) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

type DeleteUserURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortUrls []string `protobuf:"bytes,1,rep,name=short_urls,json=shortUrls,proto3" json:"short_urls,omitempty"`
}

func (x *DeleteUserURLsRequest) Reset() {
	*x = DeleteUserURLsRequest{}
	mi := &file_urlshortener_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserURLsRequest) ProtoMessage() {}

func (x *DeleteUserURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserURLsRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserURLsRequest) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserURLsRequest) GetShortUrls() []string {
	if x != nil {
		return x.ShortUrls
	}
	return nil
}

type DeleteUserURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteUserURLsResponse) Reset() {
	*x = DeleteUserURLsResponse{}
	mi := &file_urlshortener_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserURLsResponse) ProtoMessage() {}

func (x *DeleteUserURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_urlshortener_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserURLsResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserURLsResponse) Descriptor() ([]byte, []int) {
	return file_urlshortener_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserURLsResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_urlshortener_proto protoreflect.FileDescriptor

var file_urlshortener_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x22, 0x3a, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x35,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x34, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x3b, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x75,
	0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x52, 0x4c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x49, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x55, 0x52, 0x4c, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x72, 0x6c, 0x22, 0x36, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x32, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x80, 0x03, 0x0a, 0x13, 0x55, 0x52, 0x4c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x23, 0x2e, 0x75, 0x72, 0x6c,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x23, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75,
	0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c,
	0x73, 0x12, 0x20, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x23, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_urlshortener_proto_rawDescOnce sync.Once
	file_urlshortener_proto_rawDescData = file_urlshortener_proto_rawDesc
)

func file_urlshortener_proto_rawDescGZIP() []byte {
	file_urlshortener_proto_rawDescOnce.Do(func() {
		file_urlshortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_urlshortener_proto_rawDescData)
	})
	return file_urlshortener_proto_rawDescData
}

var file_urlshortener_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_urlshortener_proto_goTypes = []any{
	(*CreateShortURLRequest)(nil),  // 0: urlshortener.CreateShortURLRequest
	(*CreateShortURLResponse)(nil), // 1: urlshortener.CreateShortURLResponse
	(*GetOriginalURLRequest)(nil),  // 2: urlshortener.GetOriginalURLRequest
	(*GetOriginalURLResponse)(nil), // 3: urlshortener.GetOriginalURLResponse
	(*GetUserURLsRequest)(nil),     // 4: urlshortener.GetUserURLsRequest
	(*GetUserURLsResponse)(nil),    // 5: urlshortener.GetUserURLsResponse
	(*UserURL)(nil),                // 6: urlshortener.UserURL
	(*DeleteUserURLsRequest)(nil),  // 7: urlshortener.DeleteUserURLsRequest
	(*DeleteUserURLsResponse)(nil), // 8: urlshortener.DeleteUserURLsResponse
}
var file_urlshortener_proto_depIdxs = []int32{
	6, // 0: urlshortener.GetUserURLsResponse.urls:type_name -> urlshortener.UserURL
	0, // 1: urlshortener.URLShortenerService.CreateShortURL:input_type -> urlshortener.CreateShortURLRequest
	2, // 2: urlshortener.URLShortenerService.GetOriginalURL:input_type -> urlshortener.GetOriginalURLRequest
	4, // 3: urlshortener.URLShortenerService.GetUserURLs:input_type -> urlshortener.GetUserURLsRequest
	7, // 4: urlshortener.URLShortenerService.DeleteUserURLs:input_type -> urlshortener.DeleteUserURLsRequest
	1, // 5: urlshortener.URLShortenerService.CreateShortURL:output_type -> urlshortener.CreateShortURLResponse
	3, // 6: urlshortener.URLShortenerService.GetOriginalURL:output_type -> urlshortener.GetOriginalURLResponse
	5, // 7: urlshortener.URLShortenerService.GetUserURLs:output_type -> urlshortener.GetUserURLsResponse
	8, // 8: urlshortener.URLShortenerService.DeleteUserURLs:output_type -> urlshortener.DeleteUserURLsResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_urlshortener_proto_init() }
func file_urlshortener_proto_init() {
	if File_urlshortener_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_urlshortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_urlshortener_proto_goTypes,
		DependencyIndexes: file_urlshortener_proto_depIdxs,
		MessageInfos:      file_urlshortener_proto_msgTypes,
	}.Build()
	File_urlshortener_proto = out.File
	file_urlshortener_proto_rawDesc = nil
	file_urlshortener_proto_goTypes = nil
	file_urlshortener_proto_depIdxs = nil
}
