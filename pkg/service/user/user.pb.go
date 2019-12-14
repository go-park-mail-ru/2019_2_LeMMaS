// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

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

type UserData struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UserData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserData.Unmarshal(m, b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return xxx_messageInfo_UserData.Size(m)
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetAllParams struct {
	Nothing              bool     `protobuf:"varint,1,opt,name=nothing,proto3" json:"nothing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllParams) Reset()         { *m = GetAllParams{} }
func (m *GetAllParams) String() string { return proto.CompactTextString(m) }
func (*GetAllParams) ProtoMessage()    {}
func (*GetAllParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *GetAllParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllParams.Unmarshal(m, b)
}
func (m *GetAllParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllParams.Marshal(b, m, deterministic)
}
func (m *GetAllParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllParams.Merge(m, src)
}
func (m *GetAllParams) XXX_Size() int {
	return xxx_messageInfo_GetAllParams.Size(m)
}
func (m *GetAllParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllParams proto.InternalMessageInfo

func (m *GetAllParams) GetNothing() bool {
	if m != nil {
		return m.Nothing
	}
	return false
}

type GetAllResult struct {
	Users                []*UserData `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Error                string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetAllResult) Reset()         { *m = GetAllResult{} }
func (m *GetAllResult) String() string { return proto.CompactTextString(m) }
func (*GetAllResult) ProtoMessage()    {}
func (*GetAllResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *GetAllResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResult.Unmarshal(m, b)
}
func (m *GetAllResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResult.Marshal(b, m, deterministic)
}
func (m *GetAllResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResult.Merge(m, src)
}
func (m *GetAllResult) XXX_Size() int {
	return xxx_messageInfo_GetAllResult.Size(m)
}
func (m *GetAllResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResult proto.InternalMessageInfo

func (m *GetAllResult) GetUsers() []*UserData {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *GetAllResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetByIDParams struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIDParams) Reset()         { *m = GetByIDParams{} }
func (m *GetByIDParams) String() string { return proto.CompactTextString(m) }
func (*GetByIDParams) ProtoMessage()    {}
func (*GetByIDParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *GetByIDParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIDParams.Unmarshal(m, b)
}
func (m *GetByIDParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIDParams.Marshal(b, m, deterministic)
}
func (m *GetByIDParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIDParams.Merge(m, src)
}
func (m *GetByIDParams) XXX_Size() int {
	return xxx_messageInfo_GetByIDParams.Size(m)
}
func (m *GetByIDParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIDParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIDParams proto.InternalMessageInfo

func (m *GetByIDParams) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetByIDResult struct {
	User                 *UserData `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Error                string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetByIDResult) Reset()         { *m = GetByIDResult{} }
func (m *GetByIDResult) String() string { return proto.CompactTextString(m) }
func (*GetByIDResult) ProtoMessage()    {}
func (*GetByIDResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *GetByIDResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIDResult.Unmarshal(m, b)
}
func (m *GetByIDResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIDResult.Marshal(b, m, deterministic)
}
func (m *GetByIDResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIDResult.Merge(m, src)
}
func (m *GetByIDResult) XXX_Size() int {
	return xxx_messageInfo_GetByIDResult.Size(m)
}
func (m *GetByIDResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIDResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIDResult proto.InternalMessageInfo

func (m *GetByIDResult) GetUser() *UserData {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetByIDResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetBySessionParams struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBySessionParams) Reset()         { *m = GetBySessionParams{} }
func (m *GetBySessionParams) String() string { return proto.CompactTextString(m) }
func (*GetBySessionParams) ProtoMessage()    {}
func (*GetBySessionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *GetBySessionParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBySessionParams.Unmarshal(m, b)
}
func (m *GetBySessionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBySessionParams.Marshal(b, m, deterministic)
}
func (m *GetBySessionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBySessionParams.Merge(m, src)
}
func (m *GetBySessionParams) XXX_Size() int {
	return xxx_messageInfo_GetBySessionParams.Size(m)
}
func (m *GetBySessionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBySessionParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetBySessionParams proto.InternalMessageInfo

func (m *GetBySessionParams) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type GetBySessionResult struct {
	User                 *UserData `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Error                string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetBySessionResult) Reset()         { *m = GetBySessionResult{} }
func (m *GetBySessionResult) String() string { return proto.CompactTextString(m) }
func (*GetBySessionResult) ProtoMessage()    {}
func (*GetBySessionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *GetBySessionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBySessionResult.Unmarshal(m, b)
}
func (m *GetBySessionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBySessionResult.Marshal(b, m, deterministic)
}
func (m *GetBySessionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBySessionResult.Merge(m, src)
}
func (m *GetBySessionResult) XXX_Size() int {
	return xxx_messageInfo_GetBySessionResult.Size(m)
}
func (m *GetBySessionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBySessionResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetBySessionResult proto.InternalMessageInfo

func (m *GetBySessionResult) GetUser() *UserData {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *GetBySessionResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UpdateParams struct {
	UserId               int32    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateParams) Reset()         { *m = UpdateParams{} }
func (m *UpdateParams) String() string { return proto.CompactTextString(m) }
func (*UpdateParams) ProtoMessage()    {}
func (*UpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *UpdateParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateParams.Unmarshal(m, b)
}
func (m *UpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateParams.Marshal(b, m, deterministic)
}
func (m *UpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateParams.Merge(m, src)
}
func (m *UpdateParams) XXX_Size() int {
	return xxx_messageInfo_UpdateParams.Size(m)
}
func (m *UpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateParams proto.InternalMessageInfo

func (m *UpdateParams) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateParams) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdateParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateResult struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResult) Reset()         { *m = UpdateResult{} }
func (m *UpdateResult) String() string { return proto.CompactTextString(m) }
func (*UpdateResult) ProtoMessage()    {}
func (*UpdateResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UpdateResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResult.Unmarshal(m, b)
}
func (m *UpdateResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResult.Marshal(b, m, deterministic)
}
func (m *UpdateResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResult.Merge(m, src)
}
func (m *UpdateResult) XXX_Size() int {
	return xxx_messageInfo_UpdateResult.Size(m)
}
func (m *UpdateResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResult.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResult proto.InternalMessageInfo

func (m *UpdateResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UpdateAvatarParams struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AvatarPath           string   `protobuf:"bytes,2,opt,name=avatar_path,json=avatarPath,proto3" json:"avatar_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAvatarParams) Reset()         { *m = UpdateAvatarParams{} }
func (m *UpdateAvatarParams) String() string { return proto.CompactTextString(m) }
func (*UpdateAvatarParams) ProtoMessage()    {}
func (*UpdateAvatarParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UpdateAvatarParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAvatarParams.Unmarshal(m, b)
}
func (m *UpdateAvatarParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAvatarParams.Marshal(b, m, deterministic)
}
func (m *UpdateAvatarParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAvatarParams.Merge(m, src)
}
func (m *UpdateAvatarParams) XXX_Size() int {
	return xxx_messageInfo_UpdateAvatarParams.Size(m)
}
func (m *UpdateAvatarParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAvatarParams.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAvatarParams proto.InternalMessageInfo

func (m *UpdateAvatarParams) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateAvatarParams) GetAvatarPath() string {
	if m != nil {
		return m.AvatarPath
	}
	return ""
}

type UpdateAvatarResult struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAvatarResult) Reset()         { *m = UpdateAvatarResult{} }
func (m *UpdateAvatarResult) String() string { return proto.CompactTextString(m) }
func (*UpdateAvatarResult) ProtoMessage()    {}
func (*UpdateAvatarResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *UpdateAvatarResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAvatarResult.Unmarshal(m, b)
}
func (m *UpdateAvatarResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAvatarResult.Marshal(b, m, deterministic)
}
func (m *UpdateAvatarResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAvatarResult.Merge(m, src)
}
func (m *UpdateAvatarResult) XXX_Size() int {
	return xxx_messageInfo_UpdateAvatarResult.Size(m)
}
func (m *UpdateAvatarResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAvatarResult.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAvatarResult proto.InternalMessageInfo

func (m *UpdateAvatarResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type GetSpecialAvatarParams struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSpecialAvatarParams) Reset()         { *m = GetSpecialAvatarParams{} }
func (m *GetSpecialAvatarParams) String() string { return proto.CompactTextString(m) }
func (*GetSpecialAvatarParams) ProtoMessage()    {}
func (*GetSpecialAvatarParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *GetSpecialAvatarParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSpecialAvatarParams.Unmarshal(m, b)
}
func (m *GetSpecialAvatarParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSpecialAvatarParams.Marshal(b, m, deterministic)
}
func (m *GetSpecialAvatarParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpecialAvatarParams.Merge(m, src)
}
func (m *GetSpecialAvatarParams) XXX_Size() int {
	return xxx_messageInfo_GetSpecialAvatarParams.Size(m)
}
func (m *GetSpecialAvatarParams) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpecialAvatarParams.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpecialAvatarParams proto.InternalMessageInfo

func (m *GetSpecialAvatarParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetSpecialAvatarResult struct {
	AvatarUrl            string   `protobuf:"bytes,1,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSpecialAvatarResult) Reset()         { *m = GetSpecialAvatarResult{} }
func (m *GetSpecialAvatarResult) String() string { return proto.CompactTextString(m) }
func (*GetSpecialAvatarResult) ProtoMessage()    {}
func (*GetSpecialAvatarResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{12}
}

func (m *GetSpecialAvatarResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSpecialAvatarResult.Unmarshal(m, b)
}
func (m *GetSpecialAvatarResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSpecialAvatarResult.Marshal(b, m, deterministic)
}
func (m *GetSpecialAvatarResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpecialAvatarResult.Merge(m, src)
}
func (m *GetSpecialAvatarResult) XXX_Size() int {
	return xxx_messageInfo_GetSpecialAvatarResult.Size(m)
}
func (m *GetSpecialAvatarResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpecialAvatarResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpecialAvatarResult proto.InternalMessageInfo

func (m *GetSpecialAvatarResult) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*UserData)(nil), "user.UserData")
	proto.RegisterType((*GetAllParams)(nil), "user.GetAllParams")
	proto.RegisterType((*GetAllResult)(nil), "user.GetAllResult")
	proto.RegisterType((*GetByIDParams)(nil), "user.GetByIDParams")
	proto.RegisterType((*GetByIDResult)(nil), "user.GetByIDResult")
	proto.RegisterType((*GetBySessionParams)(nil), "user.GetBySessionParams")
	proto.RegisterType((*GetBySessionResult)(nil), "user.GetBySessionResult")
	proto.RegisterType((*UpdateParams)(nil), "user.UpdateParams")
	proto.RegisterType((*UpdateResult)(nil), "user.UpdateResult")
	proto.RegisterType((*UpdateAvatarParams)(nil), "user.UpdateAvatarParams")
	proto.RegisterType((*UpdateAvatarResult)(nil), "user.UpdateAvatarResult")
	proto.RegisterType((*GetSpecialAvatarParams)(nil), "user.GetSpecialAvatarParams")
	proto.RegisterType((*GetSpecialAvatarResult)(nil), "user.GetSpecialAvatarResult")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0xfd, 0x48, 0x9b, 0xbb, 0xeb, 0x22, 0x57, 0xd1, 0x10, 0x5c, 0x76, 0x19, 0xf6, 0xa1,
	0x88, 0xec, 0x43, 0x16, 0xf1, 0x79, 0xa5, 0x50, 0xe2, 0x43, 0x29, 0x29, 0xc5, 0xc7, 0x30, 0x9a,
	0xc1, 0x06, 0xd2, 0x24, 0xcc, 0x4c, 0x15, 0x9f, 0xfd, 0xe3, 0x92, 0xf9, 0x32, 0xb1, 0xa9, 0x20,
	0xbe, 0xe5, 0xde, 0x7b, 0xe6, 0x9c, 0x33, 0x73, 0x4f, 0x0b, 0x70, 0x14, 0x8c, 0xdf, 0x37, 0xbc,
	0x96, 0x35, 0x4e, 0xda, 0x6f, 0xb2, 0x81, 0xf9, 0x4e, 0x30, 0xbe, 0xa4, 0x92, 0xe2, 0x0b, 0x98,
	0xb2, 0x03, 0x2d, 0xca, 0xd0, 0xbb, 0xf5, 0x16, 0x41, 0xaa, 0x0b, 0x8c, 0x60, 0xde, 0x50, 0x21,
	0xbe, 0xd7, 0x3c, 0x0f, 0x47, 0x6a, 0xe0, 0x6a, 0x44, 0x98, 0x54, 0xf4, 0xc0, 0xc2, 0xb1, 0xea,
	0xab, 0x6f, 0xb2, 0x80, 0xcb, 0x15, 0x93, 0x8f, 0x65, 0xb9, 0xa1, 0x9c, 0x1e, 0x04, 0x86, 0x30,
	0xab, 0x6a, 0xb9, 0x2f, 0xaa, 0xaf, 0x8a, 0x77, 0x9e, 0xda, 0x92, 0x7c, 0xb4, 0xc8, 0x94, 0x89,
	0x63, 0x29, 0xf1, 0x0e, 0xa6, 0xad, 0x27, 0x11, 0x7a, 0xb7, 0xe3, 0xc5, 0x45, 0x7c, 0x75, 0xaf,
	0xdc, 0x5a, 0x7b, 0xa9, 0x1e, 0x2a, 0x97, 0x9c, 0xd7, 0xdc, 0x98, 0xd1, 0x05, 0xb9, 0x81, 0xa7,
	0x2b, 0x26, 0x3f, 0xfc, 0x48, 0x96, 0x46, 0xf6, 0x0a, 0x46, 0x45, 0xae, 0x14, 0xa7, 0xe9, 0xa8,
	0xc8, 0x49, 0xe2, 0x00, 0x46, 0x8d, 0x80, 0x7a, 0x01, 0x05, 0x39, 0x15, 0x53, 0xb3, 0x33, 0x5a,
	0x0f, 0x80, 0x8a, 0x6a, 0xcb, 0x84, 0x28, 0xea, 0xca, 0x08, 0x5e, 0x03, 0x08, 0xdd, 0xc8, 0x8c,
	0x70, 0x90, 0x06, 0xa6, 0x93, 0xe4, 0x64, 0xdd, 0x3f, 0xf4, 0xdf, 0x26, 0x3e, 0xc1, 0xe5, 0xae,
	0xc9, 0xa9, 0x64, 0x46, 0xfe, 0x15, 0xcc, 0x5a, 0x74, 0xe6, 0x2e, 0xed, 0xb7, 0x65, 0x92, 0xff,
	0xf3, 0xfe, 0xee, 0x2c, 0xb1, 0xb1, 0xe8, 0xe4, 0xbd, 0xae, 0xfc, 0x1a, 0x50, 0xa3, 0x1e, 0xbf,
	0x51, 0x49, 0xf9, 0xb0, 0x89, 0xc0, 0x99, 0xb8, 0x81, 0x0b, 0xaa, 0x80, 0x59, 0x43, 0xe5, 0xde,
	0xf8, 0x00, 0x6a, 0xce, 0xca, 0x3d, 0x79, 0xd3, 0xe7, 0xfb, 0xab, 0xf6, 0x5b, 0x78, 0xb9, 0x62,
	0x72, 0xdb, 0xb0, 0x2f, 0x05, 0x2d, 0x7b, 0xfa, 0xf6, 0x3e, 0x5e, 0xe7, 0x3e, 0xef, 0x4f, 0xd1,
	0x86, 0xfd, 0x1a, 0x8c, 0x83, 0xec, 0xc8, 0x6d, 0xe8, 0x03, 0xdd, 0xd9, 0xf1, 0x32, 0xfe, 0x39,
	0x86, 0x49, 0xbb, 0x0a, 0x8c, 0xc1, 0xd7, 0x39, 0x45, 0xd4, 0x0b, 0xea, 0xe6, 0x3b, 0xea, 0xf5,
	0x34, 0x33, 0x79, 0x82, 0xef, 0x60, 0x66, 0xe2, 0x86, 0xcf, 0x1d, 0xe0, 0x77, 0x3c, 0xa3, 0x7e,
	0xd3, 0x1d, 0x5b, 0xaa, 0x9f, 0x84, 0x4b, 0x09, 0x86, 0x1d, 0x58, 0x2f, 0x6e, 0xd1, 0xc0, 0xc4,
	0xb1, 0xc4, 0xe0, 0xeb, 0xc7, 0xb4, 0x86, 0xbb, 0x49, 0x89, 0x7a, 0xbd, 0xae, 0x72, 0x77, 0x01,
	0x56, 0xf9, 0x74, 0xc9, 0xd1, 0xc0, 0xc4, 0xb1, 0x6c, 0xe0, 0xd9, 0x9f, 0x8f, 0x8d, 0xaf, 0x9d,
	0xd3, 0x81, 0x95, 0x45, 0x67, 0xa6, 0x96, 0xf1, 0xb3, 0xaf, 0xfe, 0xad, 0x1e, 0x7e, 0x05, 0x00,
	0x00, 0xff, 0xff, 0xca, 0xc7, 0x48, 0xde, 0xbb, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	GetAll(ctx context.Context, in *GetAllParams, opts ...grpc.CallOption) (*GetAllResult, error)
	GetByID(ctx context.Context, in *GetByIDParams, opts ...grpc.CallOption) (*GetByIDResult, error)
	GetBySession(ctx context.Context, in *GetBySessionParams, opts ...grpc.CallOption) (*GetBySessionResult, error)
	Update(ctx context.Context, in *UpdateParams, opts ...grpc.CallOption) (*UpdateResult, error)
	UpdateAvatar(ctx context.Context, in *UpdateAvatarParams, opts ...grpc.CallOption) (*UpdateAvatarResult, error)
	GetSpecialAvatar(ctx context.Context, in *GetSpecialAvatarParams, opts ...grpc.CallOption) (*GetSpecialAvatarResult, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetAll(ctx context.Context, in *GetAllParams, opts ...grpc.CallOption) (*GetAllResult, error) {
	out := new(GetAllResult)
	err := c.cc.Invoke(ctx, "/user.User/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetByID(ctx context.Context, in *GetByIDParams, opts ...grpc.CallOption) (*GetByIDResult, error) {
	out := new(GetByIDResult)
	err := c.cc.Invoke(ctx, "/user.User/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetBySession(ctx context.Context, in *GetBySessionParams, opts ...grpc.CallOption) (*GetBySessionResult, error) {
	out := new(GetBySessionResult)
	err := c.cc.Invoke(ctx, "/user.User/GetBySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Update(ctx context.Context, in *UpdateParams, opts ...grpc.CallOption) (*UpdateResult, error) {
	out := new(UpdateResult)
	err := c.cc.Invoke(ctx, "/user.User/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateAvatar(ctx context.Context, in *UpdateAvatarParams, opts ...grpc.CallOption) (*UpdateAvatarResult, error) {
	out := new(UpdateAvatarResult)
	err := c.cc.Invoke(ctx, "/user.User/UpdateAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetSpecialAvatar(ctx context.Context, in *GetSpecialAvatarParams, opts ...grpc.CallOption) (*GetSpecialAvatarResult, error) {
	out := new(GetSpecialAvatarResult)
	err := c.cc.Invoke(ctx, "/user.User/GetSpecialAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	GetAll(context.Context, *GetAllParams) (*GetAllResult, error)
	GetByID(context.Context, *GetByIDParams) (*GetByIDResult, error)
	GetBySession(context.Context, *GetBySessionParams) (*GetBySessionResult, error)
	Update(context.Context, *UpdateParams) (*UpdateResult, error)
	UpdateAvatar(context.Context, *UpdateAvatarParams) (*UpdateAvatarResult, error)
	GetSpecialAvatar(context.Context, *GetSpecialAvatarParams) (*GetSpecialAvatarResult, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetAll(ctx context.Context, req *GetAllParams) (*GetAllResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedUserServer) GetByID(ctx context.Context, req *GetByIDParams) (*GetByIDResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (*UnimplementedUserServer) GetBySession(ctx context.Context, req *GetBySessionParams) (*GetBySessionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBySession not implemented")
}
func (*UnimplementedUserServer) Update(ctx context.Context, req *UpdateParams) (*UpdateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedUserServer) UpdateAvatar(ctx context.Context, req *UpdateAvatarParams) (*UpdateAvatarResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvatar not implemented")
}
func (*UnimplementedUserServer) GetSpecialAvatar(ctx context.Context, req *GetSpecialAvatarParams) (*GetSpecialAvatarResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpecialAvatar not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAll(ctx, req.(*GetAllParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetByID(ctx, req.(*GetByIDParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetBySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBySessionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetBySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetBySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetBySession(ctx, req.(*GetBySessionParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Update(ctx, req.(*UpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvatarParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/UpdateAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateAvatar(ctx, req.(*UpdateAvatarParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetSpecialAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecialAvatarParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetSpecialAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetSpecialAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetSpecialAvatar(ctx, req.(*GetSpecialAvatarParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _User_GetAll_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _User_GetByID_Handler,
		},
		{
			MethodName: "GetBySession",
			Handler:    _User_GetBySession_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _User_Update_Handler,
		},
		{
			MethodName: "UpdateAvatar",
			Handler:    _User_UpdateAvatar_Handler,
		},
		{
			MethodName: "GetSpecialAvatar",
			Handler:    _User_GetSpecialAvatar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
