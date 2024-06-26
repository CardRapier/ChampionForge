// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: tournament.proto

package proto_tournament

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

// Get All tournaments based on the dates
type GetTournamentsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTime int64 `protobuf:"varint,1,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	FinalTime int64 `protobuf:"varint,2,opt,name=FinalTime,proto3" json:"FinalTime,omitempty"`
}

func (x *GetTournamentsReq) Reset() {
	*x = GetTournamentsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tournament_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTournamentsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTournamentsReq) ProtoMessage() {}

func (x *GetTournamentsReq) ProtoReflect() protoreflect.Message {
	mi := &file_tournament_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTournamentsReq.ProtoReflect.Descriptor instead.
func (*GetTournamentsReq) Descriptor() ([]byte, []int) {
	return file_tournament_proto_rawDescGZIP(), []int{0}
}

func (x *GetTournamentsReq) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetTournamentsReq) GetFinalTime() int64 {
	if x != nil {
		return x.FinalTime
	}
	return 0
}

type GetTournamentsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tournaments []*Tournament `protobuf:"bytes,1,rep,name=Tournaments,proto3" json:"Tournaments,omitempty"`
}

func (x *GetTournamentsRes) Reset() {
	*x = GetTournamentsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tournament_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTournamentsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTournamentsRes) ProtoMessage() {}

func (x *GetTournamentsRes) ProtoReflect() protoreflect.Message {
	mi := &file_tournament_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTournamentsRes.ProtoReflect.Descriptor instead.
func (*GetTournamentsRes) Descriptor() ([]byte, []int) {
	return file_tournament_proto_rawDescGZIP(), []int{1}
}

func (x *GetTournamentsRes) GetTournaments() []*Tournament {
	if x != nil {
		return x.Tournaments
	}
	return nil
}

// Get a tournament
type GetTournamentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TournamentId int64 `protobuf:"varint,1,opt,name=TournamentId,proto3" json:"TournamentId,omitempty"`
}

func (x *GetTournamentReq) Reset() {
	*x = GetTournamentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tournament_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTournamentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTournamentReq) ProtoMessage() {}

func (x *GetTournamentReq) ProtoReflect() protoreflect.Message {
	mi := &file_tournament_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTournamentReq.ProtoReflect.Descriptor instead.
func (*GetTournamentReq) Descriptor() ([]byte, []int) {
	return file_tournament_proto_rawDescGZIP(), []int{2}
}

func (x *GetTournamentReq) GetTournamentId() int64 {
	if x != nil {
		return x.TournamentId
	}
	return 0
}

type GetTournamentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tournament *Tournament `protobuf:"bytes,1,opt,name=Tournament,proto3" json:"Tournament,omitempty"`
}

func (x *GetTournamentRes) Reset() {
	*x = GetTournamentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tournament_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTournamentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTournamentRes) ProtoMessage() {}

func (x *GetTournamentRes) ProtoReflect() protoreflect.Message {
	mi := &file_tournament_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTournamentRes.ProtoReflect.Descriptor instead.
func (*GetTournamentRes) Descriptor() ([]byte, []int) {
	return file_tournament_proto_rawDescGZIP(), []int{3}
}

func (x *GetTournamentRes) GetTournament() *Tournament {
	if x != nil {
		return x.Tournament
	}
	return nil
}

// Create Tournaments
type CreateTournamentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTournaments []*CreateTournamentDTO `protobuf:"bytes,1,rep,name=CreateTournaments,proto3" json:"CreateTournaments,omitempty"`
}

func (x *CreateTournamentReq) Reset() {
	*x = CreateTournamentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tournament_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTournamentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTournamentReq) ProtoMessage() {}

func (x *CreateTournamentReq) ProtoReflect() protoreflect.Message {
	mi := &file_tournament_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTournamentReq.ProtoReflect.Descriptor instead.
func (*CreateTournamentReq) Descriptor() ([]byte, []int) {
	return file_tournament_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTournamentReq) GetCreateTournaments() []*CreateTournamentDTO {
	if x != nil {
		return x.CreateTournaments
	}
	return nil
}

type CreateTournamentRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tournaments []*Tournament `protobuf:"bytes,1,rep,name=Tournaments,proto3" json:"Tournaments,omitempty"`
}

func (x *CreateTournamentRes) Reset() {
	*x = CreateTournamentRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tournament_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTournamentRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTournamentRes) ProtoMessage() {}

func (x *CreateTournamentRes) ProtoReflect() protoreflect.Message {
	mi := &file_tournament_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTournamentRes.ProtoReflect.Descriptor instead.
func (*CreateTournamentRes) Descriptor() ([]byte, []int) {
	return file_tournament_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTournamentRes) GetTournaments() []*Tournament {
	if x != nil {
		return x.Tournaments
	}
	return nil
}

// Types
type Tournament struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	EntryFee  int64  `protobuf:"varint,3,opt,name=EntryFee,proto3" json:"EntryFee,omitempty"`
	StartTime int64  `protobuf:"varint,4,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
}

func (x *Tournament) Reset() {
	*x = Tournament{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tournament_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tournament) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tournament) ProtoMessage() {}

func (x *Tournament) ProtoReflect() protoreflect.Message {
	mi := &file_tournament_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tournament.ProtoReflect.Descriptor instead.
func (*Tournament) Descriptor() ([]byte, []int) {
	return file_tournament_proto_rawDescGZIP(), []int{6}
}

func (x *Tournament) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tournament) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tournament) GetEntryFee() int64 {
	if x != nil {
		return x.EntryFee
	}
	return 0
}

func (x *Tournament) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

// DTOs
type CreateTournamentDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	EntryFee  int64  `protobuf:"varint,3,opt,name=EntryFee,proto3" json:"EntryFee,omitempty"`
	StartTime int64  `protobuf:"varint,4,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
}

func (x *CreateTournamentDTO) Reset() {
	*x = CreateTournamentDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tournament_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTournamentDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTournamentDTO) ProtoMessage() {}

func (x *CreateTournamentDTO) ProtoReflect() protoreflect.Message {
	mi := &file_tournament_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTournamentDTO.ProtoReflect.Descriptor instead.
func (*CreateTournamentDTO) Descriptor() ([]byte, []int) {
	return file_tournament_proto_rawDescGZIP(), []int{7}
}

func (x *CreateTournamentDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTournamentDTO) GetEntryFee() int64 {
	if x != nil {
		return x.EntryFee
	}
	return 0
}

func (x *CreateTournamentDTO) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

var File_tournament_proto protoreflect.FileDescriptor

var file_tournament_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x3f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x59, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x42, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x44, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x2d, 0x0a, 0x0b, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x6a, 0x0a, 0x0a, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x46, 0x65, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x46, 0x65, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x63, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x54, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x46, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x46, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x32, 0xcb, 0x01, 0x0a, 0x11, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x2a, 0x5a, 0x28, 0x6b, 0x72, 0x6f, 0x77, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tournament_proto_rawDescOnce sync.Once
	file_tournament_proto_rawDescData = file_tournament_proto_rawDesc
)

func file_tournament_proto_rawDescGZIP() []byte {
	file_tournament_proto_rawDescOnce.Do(func() {
		file_tournament_proto_rawDescData = protoimpl.X.CompressGZIP(file_tournament_proto_rawDescData)
	})
	return file_tournament_proto_rawDescData
}

var file_tournament_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tournament_proto_goTypes = []interface{}{
	(*GetTournamentsReq)(nil),   // 0: GetTournamentsReq
	(*GetTournamentsRes)(nil),   // 1: GetTournamentsRes
	(*GetTournamentReq)(nil),    // 2: GetTournamentReq
	(*GetTournamentRes)(nil),    // 3: GetTournamentRes
	(*CreateTournamentReq)(nil), // 4: CreateTournamentReq
	(*CreateTournamentRes)(nil), // 5: CreateTournamentRes
	(*Tournament)(nil),          // 6: Tournament
	(*CreateTournamentDTO)(nil), // 7: CreateTournamentDTO
}
var file_tournament_proto_depIdxs = []int32{
	6, // 0: GetTournamentsRes.Tournaments:type_name -> Tournament
	6, // 1: GetTournamentRes.Tournament:type_name -> Tournament
	7, // 2: CreateTournamentReq.CreateTournaments:type_name -> CreateTournamentDTO
	6, // 3: CreateTournamentRes.Tournaments:type_name -> Tournament
	0, // 4: TournamentService.GetTournaments:input_type -> GetTournamentsReq
	2, // 5: TournamentService.GetTournament:input_type -> GetTournamentReq
	4, // 6: TournamentService.CreateTournaments:input_type -> CreateTournamentReq
	1, // 7: TournamentService.GetTournaments:output_type -> GetTournamentsRes
	3, // 8: TournamentService.GetTournament:output_type -> GetTournamentRes
	5, // 9: TournamentService.CreateTournaments:output_type -> CreateTournamentRes
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tournament_proto_init() }
func file_tournament_proto_init() {
	if File_tournament_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tournament_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTournamentsReq); i {
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
		file_tournament_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTournamentsRes); i {
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
		file_tournament_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTournamentReq); i {
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
		file_tournament_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTournamentRes); i {
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
		file_tournament_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTournamentReq); i {
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
		file_tournament_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTournamentRes); i {
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
		file_tournament_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tournament); i {
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
		file_tournament_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTournamentDTO); i {
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
			RawDescriptor: file_tournament_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tournament_proto_goTypes,
		DependencyIndexes: file_tournament_proto_depIdxs,
		MessageInfos:      file_tournament_proto_msgTypes,
	}.Build()
	File_tournament_proto = out.File
	file_tournament_proto_rawDesc = nil
	file_tournament_proto_goTypes = nil
	file_tournament_proto_depIdxs = nil
}
