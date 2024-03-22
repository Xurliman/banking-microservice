// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: payment_type.proto

package proto_bank

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

type CreatePaymentTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ActivationDate   string `protobuf:"bytes,3,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,4,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	OldName          string `protobuf:"bytes,5,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	OldCode          int64  `protobuf:"varint,6,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
}

func (x *CreatePaymentTypeRequest) Reset() {
	*x = CreatePaymentTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentTypeRequest) ProtoMessage() {}

func (x *CreatePaymentTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentTypeRequest.ProtoReflect.Descriptor instead.
func (*CreatePaymentTypeRequest) Descriptor() ([]byte, []int) {
	return file_payment_type_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaymentTypeRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreatePaymentTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePaymentTypeRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreatePaymentTypeRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *CreatePaymentTypeRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *CreatePaymentTypeRequest) GetOldCode() int64 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

type SinglePaymentTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SinglePaymentTypeRequest) Reset() {
	*x = SinglePaymentTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SinglePaymentTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SinglePaymentTypeRequest) ProtoMessage() {}

func (x *SinglePaymentTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SinglePaymentTypeRequest.ProtoReflect.Descriptor instead.
func (*SinglePaymentTypeRequest) Descriptor() ([]byte, []int) {
	return file_payment_type_proto_rawDescGZIP(), []int{1}
}

func (x *SinglePaymentTypeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PaymentTypeProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ActivationDate   string `protobuf:"bytes,4,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,5,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	OldName          string `protobuf:"bytes,6,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	OldCode          int64  `protobuf:"varint,7,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
}

func (x *PaymentTypeProfileResponse) Reset() {
	*x = PaymentTypeProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentTypeProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentTypeProfileResponse) ProtoMessage() {}

func (x *PaymentTypeProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentTypeProfileResponse.ProtoReflect.Descriptor instead.
func (*PaymentTypeProfileResponse) Descriptor() ([]byte, []int) {
	return file_payment_type_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentTypeProfileResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentTypeProfileResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PaymentTypeProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PaymentTypeProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *PaymentTypeProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *PaymentTypeProfileResponse) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *PaymentTypeProfileResponse) GetOldCode() int64 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_payment_type_proto_rawDescGZIP(), []int{3}
}

func (x *SuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdatePaymentTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ActivationDate   string `protobuf:"bytes,4,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,5,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	OldName          string `protobuf:"bytes,6,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	OldCode          int64  `protobuf:"varint,7,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
}

func (x *UpdatePaymentTypeRequest) Reset() {
	*x = UpdatePaymentTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePaymentTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePaymentTypeRequest) ProtoMessage() {}

func (x *UpdatePaymentTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePaymentTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdatePaymentTypeRequest) Descriptor() ([]byte, []int) {
	return file_payment_type_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePaymentTypeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePaymentTypeRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdatePaymentTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePaymentTypeRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdatePaymentTypeRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *UpdatePaymentTypeRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *UpdatePaymentTypeRequest) GetOldCode() int64 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

var File_payment_type_proto protoreflect.FileDescriptor

var file_payment_type_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c,
	0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x6c,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xe0, 0x01, 0x0a, 0x1a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x6c, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xde, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x6c, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x84, 0x02, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x19,
	0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_payment_type_proto_rawDescOnce sync.Once
	file_payment_type_proto_rawDescData = file_payment_type_proto_rawDesc
)

func file_payment_type_proto_rawDescGZIP() []byte {
	file_payment_type_proto_rawDescOnce.Do(func() {
		file_payment_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_type_proto_rawDescData)
	})
	return file_payment_type_proto_rawDescData
}

var file_payment_type_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_payment_type_proto_goTypes = []interface{}{
	(*CreatePaymentTypeRequest)(nil),   // 0: CreatePaymentTypeRequest
	(*SinglePaymentTypeRequest)(nil),   // 1: SinglePaymentTypeRequest
	(*PaymentTypeProfileResponse)(nil), // 2: PaymentTypeProfileResponse
	(*SuccessResponse)(nil),            // 3: SuccessResponse
	(*UpdatePaymentTypeRequest)(nil),   // 4: UpdatePaymentTypeRequest
}
var file_payment_type_proto_depIdxs = []int32{
	0, // 0: PaymentTypeService.Create:input_type -> CreatePaymentTypeRequest
	1, // 1: PaymentTypeService.Read:input_type -> SinglePaymentTypeRequest
	4, // 2: PaymentTypeService.Update:input_type -> UpdatePaymentTypeRequest
	1, // 3: PaymentTypeService.Delete:input_type -> SinglePaymentTypeRequest
	2, // 4: PaymentTypeService.Create:output_type -> PaymentTypeProfileResponse
	2, // 5: PaymentTypeService.Read:output_type -> PaymentTypeProfileResponse
	3, // 6: PaymentTypeService.Update:output_type -> SuccessResponse
	3, // 7: PaymentTypeService.Delete:output_type -> SuccessResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_type_proto_init() }
func file_payment_type_proto_init() {
	if File_payment_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentTypeRequest); i {
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
		file_payment_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SinglePaymentTypeRequest); i {
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
		file_payment_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentTypeProfileResponse); i {
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
		file_payment_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_payment_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePaymentTypeRequest); i {
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
			RawDescriptor: file_payment_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_type_proto_goTypes,
		DependencyIndexes: file_payment_type_proto_depIdxs,
		MessageInfos:      file_payment_type_proto_msgTypes,
	}.Build()
	File_payment_type_proto = out.File
	file_payment_type_proto_rawDesc = nil
	file_payment_type_proto_goTypes = nil
	file_payment_type_proto_depIdxs = nil
}