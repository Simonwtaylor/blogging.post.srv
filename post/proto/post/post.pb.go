// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/post/post.proto

package go_micro_srv_post

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UserInput struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Active               bool     `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInput) Reset()         { *m = UserInput{} }
func (m *UserInput) String() string { return proto.CompactTextString(m) }
func (*UserInput) ProtoMessage()    {}
func (*UserInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{2}
}

func (m *UserInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInput.Unmarshal(m, b)
}
func (m *UserInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInput.Marshal(b, m, deterministic)
}
func (m *UserInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInput.Merge(m, src)
}
func (m *UserInput) XXX_Size() int {
	return xxx_messageInfo_UserInput.Size(m)
}
func (m *UserInput) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInput.DiscardUnknown(m)
}

var xxx_messageInfo_UserInput proto.InternalMessageInfo

func (m *UserInput) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserInput) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInput) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInput) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type User struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Active               bool     `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	CreatedAt            int32    `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int32    `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *User) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *User) GetUpdatedAt() int32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type Post struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt            int32    `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            int32    `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{4}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Post) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Post) GetUpdatedAt() int32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

type Meta struct {
	Cursor               int32    `protobuf:"varint,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	TotalItems           int32    `protobuf:"varint,2,opt,name=totalItems,proto3" json:"totalItems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{5}
}

func (m *Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meta.Unmarshal(m, b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return xxx_messageInfo_Meta.Size(m)
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

func (m *Meta) GetCursor() int32 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *Meta) GetTotalItems() int32 {
	if m != nil {
		return m.TotalItems
	}
	return 0
}

type PaginatedQuery struct {
	Cursor               int32    `protobuf:"varint,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaginatedQuery) Reset()         { *m = PaginatedQuery{} }
func (m *PaginatedQuery) String() string { return proto.CompactTextString(m) }
func (*PaginatedQuery) ProtoMessage()    {}
func (*PaginatedQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{6}
}

func (m *PaginatedQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginatedQuery.Unmarshal(m, b)
}
func (m *PaginatedQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginatedQuery.Marshal(b, m, deterministic)
}
func (m *PaginatedQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginatedQuery.Merge(m, src)
}
func (m *PaginatedQuery) XXX_Size() int {
	return xxx_messageInfo_PaginatedQuery.Size(m)
}
func (m *PaginatedQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginatedQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PaginatedQuery proto.InternalMessageInfo

func (m *PaginatedQuery) GetCursor() int32 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func (m *PaginatedQuery) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type CreateUserRequest struct {
	User                 *UserInput `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{7}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *UserInput {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{8}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetAllUsersRequest struct {
	Query                *PaginatedQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAllUsersRequest) Reset()         { *m = GetAllUsersRequest{} }
func (m *GetAllUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllUsersRequest) ProtoMessage()    {}
func (*GetAllUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{9}
}

func (m *GetAllUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUsersRequest.Unmarshal(m, b)
}
func (m *GetAllUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetAllUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUsersRequest.Merge(m, src)
}
func (m *GetAllUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllUsersRequest.Size(m)
}
func (m *GetAllUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUsersRequest proto.InternalMessageInfo

func (m *GetAllUsersRequest) GetQuery() *PaginatedQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type GetAllUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Meta                 *Meta    `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllUsersResponse) Reset()         { *m = GetAllUsersResponse{} }
func (m *GetAllUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllUsersResponse) ProtoMessage()    {}
func (*GetAllUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{10}
}

func (m *GetAllUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUsersResponse.Unmarshal(m, b)
}
func (m *GetAllUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetAllUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUsersResponse.Merge(m, src)
}
func (m *GetAllUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllUsersResponse.Size(m)
}
func (m *GetAllUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUsersResponse proto.InternalMessageInfo

func (m *GetAllUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *GetAllUsersResponse) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type CreatePostRequest struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRequest) Reset()         { *m = CreatePostRequest{} }
func (m *CreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePostRequest) ProtoMessage()    {}
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{11}
}

func (m *CreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRequest.Unmarshal(m, b)
}
func (m *CreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRequest.Marshal(b, m, deterministic)
}
func (m *CreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRequest.Merge(m, src)
}
func (m *CreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePostRequest.Size(m)
}
func (m *CreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRequest proto.InternalMessageInfo

func (m *CreatePostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreatePostRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreatePostResponse struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostResponse) Reset()         { *m = CreatePostResponse{} }
func (m *CreatePostResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePostResponse) ProtoMessage()    {}
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{12}
}

func (m *CreatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostResponse.Unmarshal(m, b)
}
func (m *CreatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostResponse.Marshal(b, m, deterministic)
}
func (m *CreatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostResponse.Merge(m, src)
}
func (m *CreatePostResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePostResponse.Size(m)
}
func (m *CreatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostResponse proto.InternalMessageInfo

func (m *CreatePostResponse) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type GetAllPostsRequest struct {
	Query                *PaginatedQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAllPostsRequest) Reset()         { *m = GetAllPostsRequest{} }
func (m *GetAllPostsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllPostsRequest) ProtoMessage()    {}
func (*GetAllPostsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{13}
}

func (m *GetAllPostsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPostsRequest.Unmarshal(m, b)
}
func (m *GetAllPostsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPostsRequest.Marshal(b, m, deterministic)
}
func (m *GetAllPostsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPostsRequest.Merge(m, src)
}
func (m *GetAllPostsRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllPostsRequest.Size(m)
}
func (m *GetAllPostsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPostsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPostsRequest proto.InternalMessageInfo

func (m *GetAllPostsRequest) GetQuery() *PaginatedQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type GetAllPostsResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	Meta                 *Meta    `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllPostsResponse) Reset()         { *m = GetAllPostsResponse{} }
func (m *GetAllPostsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllPostsResponse) ProtoMessage()    {}
func (*GetAllPostsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_de323a772ac9b83f, []int{14}
}

func (m *GetAllPostsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPostsResponse.Unmarshal(m, b)
}
func (m *GetAllPostsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPostsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllPostsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPostsResponse.Merge(m, src)
}
func (m *GetAllPostsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllPostsResponse.Size(m)
}
func (m *GetAllPostsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPostsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPostsResponse proto.InternalMessageInfo

func (m *GetAllPostsResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *GetAllPostsResponse) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.post.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.post.Response")
	proto.RegisterType((*UserInput)(nil), "go.micro.srv.post.UserInput")
	proto.RegisterType((*User)(nil), "go.micro.srv.post.User")
	proto.RegisterType((*Post)(nil), "go.micro.srv.post.Post")
	proto.RegisterType((*Meta)(nil), "go.micro.srv.post.Meta")
	proto.RegisterType((*PaginatedQuery)(nil), "go.micro.srv.post.PaginatedQuery")
	proto.RegisterType((*CreateUserRequest)(nil), "go.micro.srv.post.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "go.micro.srv.post.CreateUserResponse")
	proto.RegisterType((*GetAllUsersRequest)(nil), "go.micro.srv.post.GetAllUsersRequest")
	proto.RegisterType((*GetAllUsersResponse)(nil), "go.micro.srv.post.GetAllUsersResponse")
	proto.RegisterType((*CreatePostRequest)(nil), "go.micro.srv.post.CreatePostRequest")
	proto.RegisterType((*CreatePostResponse)(nil), "go.micro.srv.post.CreatePostResponse")
	proto.RegisterType((*GetAllPostsRequest)(nil), "go.micro.srv.post.GetAllPostsRequest")
	proto.RegisterType((*GetAllPostsResponse)(nil), "go.micro.srv.post.GetAllPostsResponse")
}

func init() {
	proto.RegisterFile("proto/post/post.proto", fileDescriptor_de323a772ac9b83f)
}

var fileDescriptor_de323a772ac9b83f = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5f, 0x8b, 0xd3, 0x4e,
	0x14, 0xdd, 0xb4, 0x49, 0xff, 0xdc, 0xc2, 0xf2, 0xdb, 0xf9, 0xe9, 0x1a, 0x6a, 0x95, 0x3a, 0xa8,
	0x14, 0x16, 0xa3, 0xac, 0x0f, 0xbe, 0x09, 0x65, 0x2b, 0xd2, 0x87, 0x85, 0x35, 0x8b, 0x4f, 0x82,
	0x10, 0xd3, 0xa1, 0x04, 0x92, 0x4c, 0x3a, 0x33, 0xa9, 0xe8, 0x37, 0xf1, 0x9b, 0xf8, 0x19, 0xfc,
	0x54, 0x32, 0x37, 0xd3, 0x26, 0xb1, 0x69, 0x16, 0xc1, 0x97, 0x25, 0x77, 0xee, 0x99, 0x73, 0xcf,
	0x3d, 0x73, 0x76, 0x17, 0xee, 0x67, 0x82, 0x2b, 0xfe, 0x32, 0xe3, 0x52, 0xe1, 0x0f, 0x0f, 0x6b,
	0x72, 0xb6, 0xe6, 0x5e, 0x12, 0x85, 0x82, 0x7b, 0x52, 0x6c, 0x3d, 0xdd, 0xa0, 0x8f, 0xa0, 0xef,
	0xb3, 0x4d, 0xce, 0xa4, 0x22, 0x04, 0xec, 0x34, 0x48, 0x98, 0x6b, 0x4d, 0xad, 0xd9, 0xd0, 0xc7,
	0x6f, 0x3a, 0x81, 0x81, 0xcf, 0x64, 0xc6, 0x53, 0xc9, 0xc8, 0x7f, 0xd0, 0x4d, 0xe4, 0xda, 0xb4,
	0xf5, 0x27, 0xdd, 0xc0, 0xf0, 0xa3, 0x64, 0x62, 0x99, 0x66, 0xb9, 0x22, 0x63, 0x18, 0xe4, 0x92,
	0x89, 0x0a, 0xc5, 0xbe, 0xd6, 0xbd, 0x2c, 0x90, 0xf2, 0x2b, 0x17, 0x2b, 0xb7, 0x53, 0xf4, 0x76,
	0x35, 0xb9, 0x07, 0x0e, 0x4b, 0x82, 0x28, 0x76, 0xbb, 0xd8, 0x28, 0x0a, 0x72, 0x0e, 0xbd, 0x20,
	0x54, 0xd1, 0x96, 0xb9, 0xf6, 0xd4, 0x9a, 0x0d, 0x7c, 0x53, 0xd1, 0x9f, 0x16, 0xd8, 0x7a, 0x26,
	0x39, 0x85, 0xce, 0x72, 0x61, 0x06, 0x75, 0x96, 0x8b, 0xda, 0xf8, 0x4e, 0xcb, 0xf8, 0xee, 0xb1,
	0xf1, 0x76, 0xf3, 0x78, 0xa7, 0x3a, 0x9e, 0x4c, 0x60, 0x18, 0x0a, 0x16, 0x28, 0xb6, 0x9a, 0x2b,
	0xb7, 0x37, 0xb5, 0x66, 0x8e, 0x5f, 0x1e, 0xe8, 0x6e, 0x9e, 0xad, 0x4c, 0xb7, 0x5f, 0x74, 0xf7,
	0x07, 0xf4, 0x87, 0x05, 0xf6, 0x0d, 0x97, 0xea, 0x40, 0xba, 0x0b, 0xfd, 0x90, 0xa7, 0x8a, 0xa5,
	0xca, 0x28, 0xdf, 0x95, 0xe4, 0x02, 0x6c, 0xbd, 0x04, 0x8a, 0x1e, 0x5d, 0x3e, 0xf0, 0x0e, 0xde,
	0xcf, 0xd3, 0x5e, 0xf8, 0x08, 0xaa, 0x6b, 0xb3, 0x5b, 0xb5, 0x39, 0x7f, 0x6a, 0x7b, 0x0b, 0xf6,
	0x35, 0x53, 0x81, 0xde, 0x3b, 0xcc, 0x85, 0xe4, 0x02, 0xe5, 0x39, 0xbe, 0xa9, 0xc8, 0x63, 0x00,
	0xc5, 0x55, 0x10, 0x2f, 0x15, 0x4b, 0x24, 0xaa, 0x74, 0xfc, 0xca, 0x09, 0x5d, 0xc0, 0xe9, 0x4d,
	0xb0, 0x8e, 0x52, 0x4d, 0xf7, 0x21, 0x67, 0xe2, 0xdb, 0x51, 0x26, 0x7c, 0x8b, 0x35, 0xbb, 0x8d,
	0xbe, 0x33, 0xc3, 0xb3, 0xaf, 0xe9, 0x3b, 0x38, 0xbb, 0x42, 0xc1, 0xb8, 0x95, 0x89, 0xe5, 0x2b,
	0xe3, 0x81, 0x85, 0x1e, 0x4c, 0x8e, 0x78, 0x80, 0x19, 0x2c, 0x8c, 0xa0, 0x73, 0x20, 0x55, 0x1a,
	0x13, 0xdf, 0x8b, 0x1a, 0x4f, 0xbb, 0x97, 0xf4, 0x1a, 0xc8, 0x7b, 0xa6, 0xe6, 0x71, 0xac, 0xcf,
	0xe4, 0x4e, 0xca, 0x1b, 0x70, 0x36, 0x7a, 0x39, 0xc3, 0xf1, 0xa4, 0x81, 0xa3, 0xee, 0x82, 0x5f,
	0xe0, 0xe9, 0x06, 0xfe, 0xaf, 0xd1, 0x19, 0x49, 0x2f, 0xc0, 0xd1, 0xd3, 0xa4, 0x6b, 0x4d, 0xbb,
	0x6d, 0x9a, 0x0a, 0x94, 0xde, 0x20, 0x61, 0x2a, 0x40, 0xdb, 0x9a, 0xd1, 0xfa, 0x0d, 0x7d, 0x04,
	0x95, 0x5e, 0xea, 0xc8, 0xed, 0x16, 0xa8, 0x24, 0xcd, 0xaa, 0x27, 0xed, 0x1c, 0x7a, 0x7a, 0xc8,
	0x72, 0x65, 0x1e, 0xc5, 0x54, 0xa5, 0x97, 0x05, 0x4d, 0xe9, 0xa5, 0x9e, 0xd7, 0xe2, 0x25, 0xc2,
	0x11, 0x54, 0x7a, 0xa9, 0xcf, 0xfe, 0xa1, 0x97, 0x86, 0xae, 0xf4, 0x52, 0x5f, 0x6a, 0xf3, 0x12,
	0x35, 0x15, 0xa8, 0xbf, 0xf2, 0xf2, 0xf2, 0x57, 0x17, 0x46, 0xfa, 0xf2, 0x2d, 0x13, 0xdb, 0x28,
	0x64, 0x64, 0x0e, 0xf6, 0x55, 0x10, 0xc7, 0x64, 0xdc, 0x70, 0xcd, 0xec, 0x37, 0x7e, 0xd8, 0xd8,
	0x2b, 0xc4, 0xd2, 0x13, 0xf2, 0x09, 0xa0, 0xcc, 0x28, 0x79, 0xda, 0x00, 0x3e, 0xf8, 0x4d, 0x18,
	0x3f, 0xbb, 0x03, 0xb5, 0x27, 0xff, 0x0c, 0xa3, 0x4a, 0xdc, 0x48, 0xd3, 0xbd, 0xc3, 0x74, 0x8f,
	0x9f, 0xdf, 0x05, 0x3b, 0x14, 0x8f, 0x7f, 0xce, 0x8e, 0x8b, 0xaf, 0x44, 0xaf, 0x45, 0x7c, 0x35,
	0x59, 0x55, 0xf1, 0xf8, 0xbe, 0x2d, 0xe2, 0xab, 0x71, 0x6a, 0x11, 0x5f, 0x8b, 0x09, 0x3d, 0xf9,
	0xd2, 0xc3, 0xff, 0x85, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x6d, 0x38, 0x4a, 0x24,
	0x07, 0x00, 0x00,
}
