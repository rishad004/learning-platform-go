// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: institute-service/proto/institute.proto

package institute_service

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_institute_service_proto_institute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_institute_service_proto_institute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_institute_service_proto_institute_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_institute_service_proto_institute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_institute_service_proto_institute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_institute_service_proto_institute_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BlockUnblockUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlockUnblockUserRequest) Reset() {
	*x = BlockUnblockUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_institute_service_proto_institute_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockUnblockUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockUnblockUserRequest) ProtoMessage() {}

func (x *BlockUnblockUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_institute_service_proto_institute_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockUnblockUserRequest.ProtoReflect.Descriptor instead.
func (*BlockUnblockUserRequest) Descriptor() ([]byte, []int) {
	return file_institute_service_proto_institute_proto_rawDescGZIP(), []int{2}
}

type BlockUnblockUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BlockUnblockUserResponse) Reset() {
	*x = BlockUnblockUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_institute_service_proto_institute_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockUnblockUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockUnblockUserResponse) ProtoMessage() {}

func (x *BlockUnblockUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_institute_service_proto_institute_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockUnblockUserResponse.ProtoReflect.Descriptor instead.
func (*BlockUnblockUserResponse) Descriptor() ([]byte, []int) {
	return file_institute_service_proto_institute_proto_rawDescGZIP(), []int{3}
}

func (x *BlockUnblockUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AddCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course string `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
	Price  int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *AddCourseRequest) Reset() {
	*x = AddCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_institute_service_proto_institute_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCourseRequest) ProtoMessage() {}

func (x *AddCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_institute_service_proto_institute_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCourseRequest.ProtoReflect.Descriptor instead.
func (*AddCourseRequest) Descriptor() ([]byte, []int) {
	return file_institute_service_proto_institute_proto_rawDescGZIP(), []int{4}
}

func (x *AddCourseRequest) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *AddCourseRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type AddCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddCourseResponse) Reset() {
	*x = AddCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_institute_service_proto_institute_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCourseResponse) ProtoMessage() {}

func (x *AddCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_institute_service_proto_institute_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCourseResponse.ProtoReflect.Descriptor instead.
func (*AddCourseResponse) Descriptor() ([]byte, []int) {
	return file_institute_service_proto_institute_proto_rawDescGZIP(), []int{5}
}

func (x *AddCourseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCourseRequest) Reset() {
	*x = DeleteCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_institute_service_proto_institute_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseRequest) ProtoMessage() {}

func (x *DeleteCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_institute_service_proto_institute_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseRequest.ProtoReflect.Descriptor instead.
func (*DeleteCourseRequest) Descriptor() ([]byte, []int) {
	return file_institute_service_proto_institute_proto_rawDescGZIP(), []int{6}
}

type DeleteCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteCourseResponse) Reset() {
	*x = DeleteCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_institute_service_proto_institute_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseResponse) ProtoMessage() {}

func (x *DeleteCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_institute_service_proto_institute_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseResponse.ProtoReflect.Descriptor instead.
func (*DeleteCourseResponse) Descriptor() ([]byte, []int) {
	return file_institute_service_proto_institute_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCourseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_institute_service_proto_institute_proto protoreflect.FileDescriptor

var file_institute_service_proto_institute_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74,
	0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x65, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x0d,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x30, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xc0, 0x02, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x69,
	0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5b, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x2e,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74,
	0x75, 0x74, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x2e, 0x69, 0x6e, 0x73, 0x74,
	0x69, 0x74, 0x75, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x73, 0x68, 0x61, 0x64, 0x30, 0x30, 0x34, 0x2f, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2d,
	0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_institute_service_proto_institute_proto_rawDescOnce sync.Once
	file_institute_service_proto_institute_proto_rawDescData = file_institute_service_proto_institute_proto_rawDesc
)

func file_institute_service_proto_institute_proto_rawDescGZIP() []byte {
	file_institute_service_proto_institute_proto_rawDescOnce.Do(func() {
		file_institute_service_proto_institute_proto_rawDescData = protoimpl.X.CompressGZIP(file_institute_service_proto_institute_proto_rawDescData)
	})
	return file_institute_service_proto_institute_proto_rawDescData
}

var file_institute_service_proto_institute_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_institute_service_proto_institute_proto_goTypes = []any{
	(*LoginRequest)(nil),             // 0: institute.LoginRequest
	(*LoginResponse)(nil),            // 1: institute.LoginResponse
	(*BlockUnblockUserRequest)(nil),  // 2: institute.BlockUnblockUserRequest
	(*BlockUnblockUserResponse)(nil), // 3: institute.BlockUnblockUserResponse
	(*AddCourseRequest)(nil),         // 4: institute.AddCourseRequest
	(*AddCourseResponse)(nil),        // 5: institute.AddCourseResponse
	(*DeleteCourseRequest)(nil),      // 6: institute.DeleteCourseRequest
	(*DeleteCourseResponse)(nil),     // 7: institute.DeleteCourseResponse
}
var file_institute_service_proto_institute_proto_depIdxs = []int32{
	0, // 0: institute.AdminService.Login:input_type -> institute.LoginRequest
	2, // 1: institute.AdminService.BlockUnblockUser:input_type -> institute.BlockUnblockUserRequest
	4, // 2: institute.AdminService.AddCourse:input_type -> institute.AddCourseRequest
	6, // 3: institute.AdminService.DeleteCourse:input_type -> institute.DeleteCourseRequest
	1, // 4: institute.AdminService.Login:output_type -> institute.LoginResponse
	3, // 5: institute.AdminService.BlockUnblockUser:output_type -> institute.BlockUnblockUserResponse
	5, // 6: institute.AdminService.AddCourse:output_type -> institute.AddCourseResponse
	7, // 7: institute.AdminService.DeleteCourse:output_type -> institute.DeleteCourseResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_institute_service_proto_institute_proto_init() }
func file_institute_service_proto_institute_proto_init() {
	if File_institute_service_proto_institute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_institute_service_proto_institute_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LoginRequest); i {
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
		file_institute_service_proto_institute_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResponse); i {
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
		file_institute_service_proto_institute_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BlockUnblockUserRequest); i {
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
		file_institute_service_proto_institute_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BlockUnblockUserResponse); i {
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
		file_institute_service_proto_institute_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddCourseRequest); i {
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
		file_institute_service_proto_institute_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AddCourseResponse); i {
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
		file_institute_service_proto_institute_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCourseRequest); i {
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
		file_institute_service_proto_institute_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteCourseResponse); i {
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
			RawDescriptor: file_institute_service_proto_institute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_institute_service_proto_institute_proto_goTypes,
		DependencyIndexes: file_institute_service_proto_institute_proto_depIdxs,
		MessageInfos:      file_institute_service_proto_institute_proto_msgTypes,
	}.Build()
	File_institute_service_proto_institute_proto = out.File
	file_institute_service_proto_institute_proto_rawDesc = nil
	file_institute_service_proto_institute_proto_goTypes = nil
	file_institute_service_proto_institute_proto_depIdxs = nil
}
