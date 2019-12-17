// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
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

type AuthRequest_ResponseType int32

const (
	AuthRequest_UNKNOWN AuthRequest_ResponseType = 0
	AuthRequest_CODE    AuthRequest_ResponseType = 1
	AuthRequest_TOKEN   AuthRequest_ResponseType = 2
)

var AuthRequest_ResponseType_name = map[int32]string{
	0: "UNKNOWN",
	1: "CODE",
	2: "TOKEN",
}

var AuthRequest_ResponseType_value = map[string]int32{
	"UNKNOWN": 0,
	"CODE":    1,
	"TOKEN":   2,
}

func (x AuthRequest_ResponseType) String() string {
	return proto.EnumName(AuthRequest_ResponseType_name, int32(x))
}

func (AuthRequest_ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0, 0}
}

type Session_Stage int32

const (
	// A request to authenticate someone has been received, but upstream has not
	// authenticated the user.
	Session_REQUESTED Session_Stage = 0
	// Code flow was requested, and a code has been issued.
	Session_CODE Session_Stage = 1
	// An access token has been issued to the user, but the session is not for
	// offline access (aka no refresh token)
	Session_ACCESS_TOKEN_ISSUED Session_Stage = 2
	// An access token has been issued, along with a refresh token.
	Session_REFRESHABLE Session_Stage = 3
)

var Session_Stage_name = map[int32]string{
	0: "REQUESTED",
	1: "CODE",
	2: "ACCESS_TOKEN_ISSUED",
	3: "REFRESHABLE",
}

var Session_Stage_value = map[string]int32{
	"REQUESTED":           0,
	"CODE":                1,
	"ACCESS_TOKEN_ISSUED": 2,
	"REFRESHABLE":         3,
}

func (x Session_Stage) String() string {
	return proto.EnumName(Session_Stage_name, int32(x))
}

func (Session_Stage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2, 0}
}

// AuthRequest represents the information that the caller requested
// authorization with.
type AuthRequest struct {
	RedirectUri          string                   `protobuf:"bytes,1,opt,name=redirect_uri,proto3" json:"redirect_uri,omitempty"`
	State                string                   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Scopes               []string                 `protobuf:"bytes,3,rep,name=scopes,proto3" json:"scopes,omitempty"`
	Nonce                string                   `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ResponseType         AuthRequest_ResponseType `protobuf:"varint,5,opt,name=response_type,proto3,enum=oidc.core.v1beta1.AuthRequest_ResponseType" json:"response_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{0}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *AuthRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *AuthRequest) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *AuthRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *AuthRequest) GetResponseType() AuthRequest_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return AuthRequest_UNKNOWN
}

// Authorization represents the information that the authentication process
// authorized the user for.
type Authorization struct {
	Scopes               []string             `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
	Acr                  string               `protobuf:"bytes,2,opt,name=acr,proto3" json:"acr,omitempty"`
	Amr                  []string             `protobuf:"bytes,3,rep,name=amr,proto3" json:"amr,omitempty"`
	AuthorizedAt         *timestamp.Timestamp `protobuf:"bytes,4,opt,name=authorized_at,proto3" json:"authorized_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Authorization) Reset()         { *m = Authorization{} }
func (m *Authorization) String() string { return proto.CompactTextString(m) }
func (*Authorization) ProtoMessage()    {}
func (*Authorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{1}
}

func (m *Authorization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authorization.Unmarshal(m, b)
}
func (m *Authorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authorization.Marshal(b, m, deterministic)
}
func (m *Authorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authorization.Merge(m, src)
}
func (m *Authorization) XXX_Size() int {
	return xxx_messageInfo_Authorization.Size(m)
}
func (m *Authorization) XXX_DiscardUnknown() {
	xxx_messageInfo_Authorization.DiscardUnknown(m)
}

var xxx_messageInfo_Authorization proto.InternalMessageInfo

func (m *Authorization) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *Authorization) GetAcr() string {
	if m != nil {
		return m.Acr
	}
	return ""
}

func (m *Authorization) GetAmr() []string {
	if m != nil {
		return m.Amr
	}
	return nil
}

func (m *Authorization) GetAuthorizedAt() *timestamp.Timestamp {
	if m != nil {
		return m.AuthorizedAt
	}
	return nil
}

// Session represents an authenticated user from the time they are issued a
// code, until their last refresh/access token expires.
type Session struct {
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// stage represents where in the overall lifecycle this session is.
	Stage Session_Stage `protobuf:"varint,4,opt,name=stage,proto3,enum=oidc.core.v1beta1.Session_Stage" json:"stage,omitempty"`
	// request stores information about the original request we received.
	Request *AuthRequest `protobuf:"bytes,5,opt,name=request,proto3" json:"request,omitempty"`
	// tracks the details this session was actually authorized for
	Authorization *Authorization `protobuf:"bytes,6,opt,name=authorization,proto3" json:"authorization,omitempty"`
	// the client ID this session is bound to.
	ClientId string `protobuf:"bytes,7,opt,name=client_id,proto3" json:"client_id,omitempty"`
	// the scopes that have been granted for this session
	Scopes []string `protobuf:"bytes,8,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// The authorization code that was issued for the code flow.
	AuthCode *StoredToken `protobuf:"bytes,9,opt,name=auth_code,proto3" json:"auth_code,omitempty"`
	// if the auth code has been previously redeemed. If we get a subsequent
	// redemption, we should drop the whole session
	//
	// https://tools.ietf.org/html/rfc6819#section-4.4.1.1
	AuthCodeRedeemed bool `protobuf:"varint,10,opt,name=auth_code_redeemed,proto3" json:"auth_code_redeemed,omitempty"`
	// The current access token, if one has been issued. It's expiration time
	// should always be checked.
	AccessToken *StoredToken `protobuf:"bytes,11,opt,name=access_token,proto3" json:"access_token,omitempty"`
	// The currently valid refresh token for this session. I
	RefreshToken *StoredToken `protobuf:"bytes,12,opt,name=refresh_token,proto3" json:"refresh_token,omitempty"`
	// The time the whole session should be expired at. It should be garbage
	// collected at this time.
	ExpiresAt            *timestamp.Timestamp `protobuf:"bytes,13,opt,name=expires_at,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{2}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Session) GetStage() Session_Stage {
	if m != nil {
		return m.Stage
	}
	return Session_REQUESTED
}

func (m *Session) GetRequest() *AuthRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Session) GetAuthorization() *Authorization {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func (m *Session) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Session) GetScopes() []string {
	if m != nil {
		return m.Scopes
	}
	return nil
}

func (m *Session) GetAuthCode() *StoredToken {
	if m != nil {
		return m.AuthCode
	}
	return nil
}

func (m *Session) GetAuthCodeRedeemed() bool {
	if m != nil {
		return m.AuthCodeRedeemed
	}
	return false
}

func (m *Session) GetAccessToken() *StoredToken {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

func (m *Session) GetRefreshToken() *StoredToken {
	if m != nil {
		return m.RefreshToken
	}
	return nil
}

func (m *Session) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

// StoredToken represents the session-persisted state of a token
// we issued to a user
type StoredToken struct {
	// bcrypted version of the token that was issued to the user
	Bcrypted []byte `protobuf:"bytes,1,opt,name=bcrypted,proto3" json:"bcrypted,omitempty"`
	// when this token expires
	ExpiresAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expires_at,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StoredToken) Reset()         { *m = StoredToken{} }
func (m *StoredToken) String() string { return proto.CompactTextString(m) }
func (*StoredToken) ProtoMessage()    {}
func (*StoredToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{3}
}

func (m *StoredToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoredToken.Unmarshal(m, b)
}
func (m *StoredToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoredToken.Marshal(b, m, deterministic)
}
func (m *StoredToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredToken.Merge(m, src)
}
func (m *StoredToken) XXX_Size() int {
	return xxx_messageInfo_StoredToken.Size(m)
}
func (m *StoredToken) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredToken.DiscardUnknown(m)
}

var xxx_messageInfo_StoredToken proto.InternalMessageInfo

func (m *StoredToken) GetBcrypted() []byte {
	if m != nil {
		return m.Bcrypted
	}
	return nil
}

func (m *StoredToken) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

// UserToken is the value we issue directly to users. The message is serialized,
// then base64 encoded to make up the issued version.
type UserToken struct {
	// the ID of the session this token corresponds to
	SessionId string `protobuf:"bytes,1,opt,name=session_id,proto3" json:"session_id,omitempty"`
	// the token itself, to be compared to the bcrypt version on the backend
	Token                []byte   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserToken) Reset()         { *m = UserToken{} }
func (m *UserToken) String() string { return proto.CompactTextString(m) }
func (*UserToken) ProtoMessage()    {}
func (*UserToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d2c4ccf1453ffdb, []int{4}
}

func (m *UserToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserToken.Unmarshal(m, b)
}
func (m *UserToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserToken.Marshal(b, m, deterministic)
}
func (m *UserToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserToken.Merge(m, src)
}
func (m *UserToken) XXX_Size() int {
	return xxx_messageInfo_UserToken.Size(m)
}
func (m *UserToken) XXX_DiscardUnknown() {
	xxx_messageInfo_UserToken.DiscardUnknown(m)
}

var xxx_messageInfo_UserToken proto.InternalMessageInfo

func (m *UserToken) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *UserToken) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterEnum("oidc.core.v1beta1.AuthRequest_ResponseType", AuthRequest_ResponseType_name, AuthRequest_ResponseType_value)
	proto.RegisterEnum("oidc.core.v1beta1.Session_Stage", Session_Stage_name, Session_Stage_value)
	proto.RegisterType((*AuthRequest)(nil), "oidc.core.v1beta1.AuthRequest")
	proto.RegisterType((*Authorization)(nil), "oidc.core.v1beta1.Authorization")
	proto.RegisterType((*Session)(nil), "oidc.core.v1beta1.Session")
	proto.RegisterType((*StoredToken)(nil), "oidc.core.v1beta1.StoredToken")
	proto.RegisterType((*UserToken)(nil), "oidc.core.v1beta1.UserToken")
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor_0d2c4ccf1453ffdb) }

var fileDescriptor_0d2c4ccf1453ffdb = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0xd3, 0x4a,
	0x10, 0xbd, 0x76, 0x3e, 0x3d, 0x49, 0x7a, 0x7d, 0xf7, 0x5e, 0x5d, 0x4c, 0x84, 0x4a, 0xe4, 0xa7,
	0x48, 0x48, 0x2e, 0x0d, 0x12, 0x42, 0x88, 0x07, 0xd2, 0xc6, 0x15, 0xb4, 0x28, 0x55, 0xd7, 0x89,
	0x90, 0x78, 0xb1, 0x1c, 0x7b, 0x9a, 0x5a, 0x34, 0x5e, 0xb3, 0xbb, 0x01, 0xc2, 0x7f, 0xe0, 0x91,
	0x3f, 0xcb, 0x13, 0xf2, 0xda, 0x6d, 0xed, 0xd2, 0x42, 0x79, 0xf3, 0xec, 0x9c, 0x33, 0x33, 0xe7,
	0xec, 0x78, 0xa1, 0x27, 0x24, 0xe3, 0xc1, 0x12, 0x9d, 0x94, 0x33, 0xc9, 0xc8, 0x3f, 0x2c, 0x8e,
	0x42, 0x27, 0x64, 0x1c, 0x9d, 0x8f, 0xbb, 0x0b, 0x94, 0xc1, 0x6e, 0xff, 0xe1, 0x92, 0xb1, 0xe5,
	0x39, 0xee, 0x28, 0xc0, 0x62, 0x7d, 0xba, 0x23, 0xe3, 0x15, 0x0a, 0x19, 0xac, 0xd2, 0x9c, 0xd3,
	0xdf, 0xbe, 0x0e, 0xf8, 0xc4, 0x83, 0x34, 0x45, 0x2e, 0x8a, 0xfc, 0xfd, 0xeb, 0xf9, 0x20, 0xd9,
	0xe4, 0x29, 0xfb, 0xbb, 0x06, 0x9d, 0xf1, 0x5a, 0x9e, 0x51, 0xfc, 0xb0, 0x46, 0x21, 0x89, 0x0d,
	0x5d, 0x8e, 0x51, 0xcc, 0x31, 0x94, 0xfe, 0x9a, 0xc7, 0x96, 0x36, 0xd0, 0x86, 0x06, 0xad, 0x9c,
	0x91, 0xff, 0xa0, 0x21, 0x64, 0x20, 0xd1, 0xd2, 0x55, 0x32, 0x0f, 0xc8, 0xff, 0xd0, 0x14, 0x21,
	0x4b, 0x51, 0x58, 0xb5, 0x41, 0x6d, 0x68, 0xd0, 0x22, 0xca, 0xd0, 0x09, 0x4b, 0x42, 0xb4, 0xea,
	0x39, 0x5a, 0x05, 0xe4, 0x04, 0x7a, 0x1c, 0x45, 0xca, 0x12, 0x81, 0xbe, 0xdc, 0xa4, 0x68, 0x35,
	0x06, 0xda, 0x70, 0x6b, 0xf4, 0xc8, 0xf9, 0x49, 0xbe, 0x53, 0x1a, 0xcf, 0xa1, 0x05, 0x67, 0xb6,
	0x49, 0x91, 0x56, 0x2b, 0xd8, 0x8f, 0xa1, 0x5b, 0x4e, 0x93, 0x0e, 0xb4, 0xe6, 0xd3, 0xa3, 0xe9,
	0xf1, 0xdb, 0xa9, 0xf9, 0x17, 0x69, 0x43, 0x7d, 0xff, 0x78, 0xe2, 0x9a, 0x1a, 0x31, 0xa0, 0x31,
	0x3b, 0x3e, 0x72, 0xa7, 0xa6, 0x6e, 0x7f, 0xd5, 0xa0, 0x97, 0x55, 0x67, 0x3c, 0xfe, 0x12, 0xc8,
	0x98, 0x25, 0x25, 0x11, 0x5a, 0x45, 0x84, 0x09, 0xb5, 0x20, 0xe4, 0x85, 0xe0, 0xec, 0x53, 0x9d,
	0xac, 0x78, 0xa1, 0x35, 0xfb, 0x24, 0x2f, 0xa1, 0x17, 0x14, 0xc5, 0x30, 0xf2, 0x03, 0xa9, 0x04,
	0x77, 0x46, 0x7d, 0x27, 0x77, 0xdf, 0xb9, 0x70, 0xdf, 0x99, 0x5d, 0x5c, 0x1f, 0xad, 0x12, 0xec,
	0x6f, 0x0d, 0x68, 0x79, 0x28, 0x44, 0x36, 0xc9, 0x16, 0xe8, 0x71, 0x64, 0xd5, 0x54, 0x43, 0x3d,
	0x8e, 0xc8, 0x53, 0x65, 0xfa, 0x32, 0xb7, 0x71, 0x6b, 0x34, 0xb8, 0xc1, 0xa8, 0x82, 0xea, 0x78,
	0x19, 0x8e, 0xe6, 0x70, 0xf2, 0x0c, 0x5a, 0x3c, 0x37, 0x4f, 0x59, 0xdc, 0x19, 0x6d, 0xff, 0xda,
	0x62, 0x7a, 0x01, 0x27, 0x07, 0x57, 0x7a, 0x94, 0x39, 0x56, 0x53, 0xf1, 0x07, 0xb7, 0xf0, 0x2f,
	0x71, 0xb4, 0x4a, 0x23, 0x0f, 0xc0, 0x08, 0xcf, 0x63, 0x4c, 0xa4, 0x1f, 0x47, 0x56, 0x4b, 0x09,
	0xba, 0x3a, 0x28, 0x39, 0xde, 0xae, 0x38, 0xfe, 0x02, 0x8c, 0xac, 0x8c, 0x1f, 0xb2, 0x08, 0x2d,
	0xe3, 0xd6, 0xc9, 0x3d, 0xc9, 0x38, 0x46, 0x33, 0xf6, 0x1e, 0x13, 0x7a, 0x45, 0x20, 0x0e, 0x90,
	0xcb, 0xc0, 0xe7, 0x18, 0x21, 0xae, 0x30, 0xb2, 0x60, 0xa0, 0x0d, 0xdb, 0xf4, 0x86, 0x0c, 0xd9,
	0x83, 0x6e, 0x10, 0x86, 0x28, 0x84, 0x2f, 0xb3, 0x52, 0x56, 0xe7, 0x4e, 0x0d, 0x2b, 0x1c, 0x32,
	0xc9, 0x56, 0xfa, 0x94, 0xa3, 0x38, 0x2b, 0x8a, 0x74, 0xef, 0x54, 0xa4, 0x4a, 0x22, 0xcf, 0x01,
	0xf0, 0x73, 0x1a, 0x73, 0x14, 0xd9, 0x0a, 0xf5, 0x7e, 0xbb, 0x42, 0x25, 0xb4, 0x7d, 0x08, 0x0d,
	0x75, 0xf7, 0xa4, 0x07, 0x06, 0x75, 0x4f, 0xe6, 0xae, 0x37, 0x73, 0x27, 0x95, 0xe5, 0xbf, 0x07,
	0xff, 0x8e, 0xf7, 0xf7, 0x5d, 0xcf, 0xf3, 0xd5, 0x3f, 0xe0, 0xbf, 0xf6, 0xbc, 0xb9, 0x3b, 0x31,
	0x75, 0xf2, 0x37, 0x74, 0xa8, 0x7b, 0x40, 0x5d, 0xef, 0xd5, 0x78, 0xef, 0x8d, 0x6b, 0xd6, 0x0e,
	0xeb, 0x6d, 0xcd, 0xd4, 0x0f, 0xeb, 0x6d, 0xdd, 0xac, 0xd9, 0x08, 0x9d, 0xd2, 0xc4, 0xa4, 0x0f,
	0xed, 0x45, 0xc8, 0x37, 0xa9, 0xc4, 0x48, 0xbd, 0x0f, 0x5d, 0x7a, 0x19, 0x5f, 0x1b, 0x5f, 0xff,
	0xa3, 0xf1, 0xc7, 0x60, 0xcc, 0x05, 0xf2, 0xbc, 0xc9, 0x36, 0x80, 0xc8, 0xf7, 0x39, 0x5b, 0x9b,
	0xfc, 0x19, 0x2a, 0x9d, 0x64, 0xcf, 0x4a, 0xee, 0xb2, 0xae, 0x26, 0xc8, 0x83, 0xbd, 0xe6, 0xbb,
	0x7a, 0x66, 0xf4, 0xa2, 0xa9, 0x5a, 0x3d, 0xf9, 0x11, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x34, 0xbf,
	0x4d, 0x5d, 0x05, 0x00, 0x00,
}
