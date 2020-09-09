// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: consensus.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ParticipationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GlobalSlotNumber             uint64  `protobuf:"varint,1,opt,name=global_slot_number,json=globalSlotNumber,proto3" json:"global_slot_number,omitempty"`
	EpochSlotNumber              uint64  `protobuf:"varint,2,opt,name=epoch_slot_number,json=epochSlotNumber,proto3" json:"epoch_slot_number,omitempty"`
	ExpectedParticipationBalance uint64  `protobuf:"varint,3,opt,name=expected_participation_balance,json=expectedParticipationBalance,proto3" json:"expected_participation_balance,omitempty"`
	CurrentParticipationBalance  uint64  `protobuf:"varint,4,opt,name=current_participation_balance,json=currentParticipationBalance,proto3" json:"current_participation_balance,omitempty"`
	ParticipationPercentage      float32 `protobuf:"fixed32,5,opt,name=participation_percentage,json=participationPercentage,proto3" json:"participation_percentage,omitempty"`
}

func (x *ParticipationStatus) Reset() {
	*x = ParticipationStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipationStatus) ProtoMessage() {}

func (x *ParticipationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipationStatus.ProtoReflect.Descriptor instead.
func (*ParticipationStatus) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{0}
}

func (x *ParticipationStatus) GetGlobalSlotNumber() uint64 {
	if x != nil {
		return x.GlobalSlotNumber
	}
	return 0
}

func (x *ParticipationStatus) GetEpochSlotNumber() uint64 {
	if x != nil {
		return x.EpochSlotNumber
	}
	return 0
}

func (x *ParticipationStatus) GetExpectedParticipationBalance() uint64 {
	if x != nil {
		return x.ExpectedParticipationBalance
	}
	return 0
}

func (x *ParticipationStatus) GetCurrentParticipationBalance() uint64 {
	if x != nil {
		return x.CurrentParticipationBalance
	}
	return 0
}

func (x *ParticipationStatus) GetParticipationPercentage() float32 {
	if x != nil {
		return x.ParticipationPercentage
	}
	return 0
}

type ComitteeInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index    uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Proposer bool   `protobuf:"varint,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	Voted    bool   `protobuf:"varint,3,opt,name=voted,proto3" json:"voted,omitempty"`
}

func (x *ComitteeInformation) Reset() {
	*x = ComitteeInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComitteeInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComitteeInformation) ProtoMessage() {}

func (x *ComitteeInformation) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComitteeInformation.ProtoReflect.Descriptor instead.
func (*ComitteeInformation) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{1}
}

func (x *ComitteeInformation) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ComitteeInformation) GetProposer() bool {
	if x != nil {
		return x.Proposer
	}
	return false
}

func (x *ComitteeInformation) GetVoted() bool {
	if x != nil {
		return x.Voted
	}
	return false
}

type SlotInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotNumber        uint64                 `protobuf:"varint,1,opt,name=slot_number,json=slotNumber,proto3" json:"slot_number,omitempty"`
	Passed            bool                   `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	RequiredComitters uint64                 `protobuf:"varint,3,opt,name=required_comitters,json=requiredComitters,proto3" json:"required_comitters,omitempty"`
	ComitteesIndexes  []*ComitteeInformation `protobuf:"bytes,4,rep,name=comittees_indexes,json=comitteesIndexes,proto3" json:"comittees_indexes,omitempty"`
}

func (x *SlotInfo) Reset() {
	*x = SlotInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlotInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlotInfo) ProtoMessage() {}

func (x *SlotInfo) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlotInfo.ProtoReflect.Descriptor instead.
func (*SlotInfo) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{2}
}

func (x *SlotInfo) GetSlotNumber() uint64 {
	if x != nil {
		return x.SlotNumber
	}
	return 0
}

func (x *SlotInfo) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

func (x *SlotInfo) GetRequiredComitters() uint64 {
	if x != nil {
		return x.RequiredComitters
	}
	return 0
}

func (x *SlotInfo) GetComitteesIndexes() []*ComitteeInformation {
	if x != nil {
		return x.ComitteesIndexes
	}
	return nil
}

type EpochInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpochNumber uint64      `protobuf:"varint,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	Passed      bool        `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	Slots       []*SlotInfo `protobuf:"bytes,4,rep,name=slots,proto3" json:"slots,omitempty"`
}

func (x *EpochInfo) Reset() {
	*x = EpochInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochInfo) ProtoMessage() {}

func (x *EpochInfo) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochInfo.ProtoReflect.Descriptor instead.
func (*EpochInfo) Descriptor() ([]byte, []int) {
	return file_consensus_proto_rawDescGZIP(), []int{3}
}

func (x *EpochInfo) GetEpochNumber() uint64 {
	if x != nil {
		return x.EpochNumber
	}
	return 0
}

func (x *EpochInfo) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

func (x *EpochInfo) GetSlots() []*SlotInfo {
	if x != nil {
		return x.Slots
	}
	return nil
}

var File_consensus_proto protoreflect.FileDescriptor

var file_consensus_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02,
	0x0a, 0x13, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x73, 0x6c, 0x6f,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x44, 0x0a, 0x1e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1c, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x1d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1b, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x18, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x17, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x22, 0x5d, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x6f,
	0x74, 0x65, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x08, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x43,
	0x6f, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x6f, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x65, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x22, 0x67, 0x0a, 0x09, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x65, 0x70, 0x6f, 0x63, 0x68, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73,
	0x6c, 0x6f, 0x74, 0x73, 0x32, 0x80, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73,
	0x75, 0x73, 0x12, 0x58, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69,
	0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x12, 0x18, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x07, 0x2e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x1a, 0x09, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x73, 0x6c, 0x6f, 0x74, 0x69, 0x6e, 0x66, 0x6f, 0x2f,
	0x7b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x7d, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x07, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x1a, 0x0a, 0x2e, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x28, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2f, 0x67, 0x65, 0x74, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x7b,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x7d, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consensus_proto_rawDescOnce sync.Once
	file_consensus_proto_rawDescData = file_consensus_proto_rawDesc
)

func file_consensus_proto_rawDescGZIP() []byte {
	file_consensus_proto_rawDescOnce.Do(func() {
		file_consensus_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_proto_rawDescData)
	})
	return file_consensus_proto_rawDescData
}

var file_consensus_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_consensus_proto_goTypes = []interface{}{
	(*ParticipationStatus)(nil), // 0: ParticipationStatus
	(*ComitteeInformation)(nil), // 1: ComitteeInformation
	(*SlotInfo)(nil),            // 2: SlotInfo
	(*EpochInfo)(nil),           // 3: EpochInfo
	(*Empty)(nil),               // 4: Empty
	(*Number)(nil),              // 5: Number
}
var file_consensus_proto_depIdxs = []int32{
	1, // 0: SlotInfo.comittees_indexes:type_name -> ComitteeInformation
	2, // 1: EpochInfo.slots:type_name -> SlotInfo
	4, // 2: Consensus.GetParticipationStatus:input_type -> Empty
	5, // 3: Consensus.GetSlotInfo:input_type -> Number
	5, // 4: Consensus.GetEpochInfo:input_type -> Number
	0, // 5: Consensus.GetParticipationStatus:output_type -> ParticipationStatus
	2, // 6: Consensus.GetSlotInfo:output_type -> SlotInfo
	3, // 7: Consensus.GetEpochInfo:output_type -> EpochInfo
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_consensus_proto_init() }
func file_consensus_proto_init() {
	if File_consensus_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_consensus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipationStatus); i {
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
		file_consensus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComitteeInformation); i {
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
		file_consensus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlotInfo); i {
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
		file_consensus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochInfo); i {
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
			RawDescriptor: file_consensus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_consensus_proto_goTypes,
		DependencyIndexes: file_consensus_proto_depIdxs,
		MessageInfos:      file_consensus_proto_msgTypes,
	}.Build()
	File_consensus_proto = out.File
	file_consensus_proto_rawDesc = nil
	file_consensus_proto_goTypes = nil
	file_consensus_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConsensusClient is the client API for Consensus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsensusClient interface {
	GetParticipationStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ParticipationStatus, error)
	GetSlotInfo(ctx context.Context, in *Number, opts ...grpc.CallOption) (*SlotInfo, error)
	GetEpochInfo(ctx context.Context, in *Number, opts ...grpc.CallOption) (*EpochInfo, error)
}

type consensusClient struct {
	cc grpc.ClientConnInterface
}

func NewConsensusClient(cc grpc.ClientConnInterface) ConsensusClient {
	return &consensusClient{cc}
}

func (c *consensusClient) GetParticipationStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ParticipationStatus, error) {
	out := new(ParticipationStatus)
	err := c.cc.Invoke(ctx, "/Consensus/GetParticipationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusClient) GetSlotInfo(ctx context.Context, in *Number, opts ...grpc.CallOption) (*SlotInfo, error) {
	out := new(SlotInfo)
	err := c.cc.Invoke(ctx, "/Consensus/GetSlotInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusClient) GetEpochInfo(ctx context.Context, in *Number, opts ...grpc.CallOption) (*EpochInfo, error) {
	out := new(EpochInfo)
	err := c.cc.Invoke(ctx, "/Consensus/GetEpochInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsensusServer is the server API for Consensus service.
type ConsensusServer interface {
	GetParticipationStatus(context.Context, *Empty) (*ParticipationStatus, error)
	GetSlotInfo(context.Context, *Number) (*SlotInfo, error)
	GetEpochInfo(context.Context, *Number) (*EpochInfo, error)
}

// UnimplementedConsensusServer can be embedded to have forward compatible implementations.
type UnimplementedConsensusServer struct {
}

func (*UnimplementedConsensusServer) GetParticipationStatus(context.Context, *Empty) (*ParticipationStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParticipationStatus not implemented")
}
func (*UnimplementedConsensusServer) GetSlotInfo(context.Context, *Number) (*SlotInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSlotInfo not implemented")
}
func (*UnimplementedConsensusServer) GetEpochInfo(context.Context, *Number) (*EpochInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEpochInfo not implemented")
}

func RegisterConsensusServer(s *grpc.Server, srv ConsensusServer) {
	s.RegisterService(&_Consensus_serviceDesc, srv)
}

func _Consensus_GetParticipationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServer).GetParticipationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Consensus/GetParticipationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServer).GetParticipationStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consensus_GetSlotInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServer).GetSlotInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Consensus/GetSlotInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServer).GetSlotInfo(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consensus_GetEpochInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Number)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServer).GetEpochInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Consensus/GetEpochInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServer).GetEpochInfo(ctx, req.(*Number))
	}
	return interceptor(ctx, in, info, handler)
}

var _Consensus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Consensus",
	HandlerType: (*ConsensusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParticipationStatus",
			Handler:    _Consensus_GetParticipationStatus_Handler,
		},
		{
			MethodName: "GetSlotInfo",
			Handler:    _Consensus_GetSlotInfo_Handler,
		},
		{
			MethodName: "GetEpochInfo",
			Handler:    _Consensus_GetEpochInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consensus.proto",
}
