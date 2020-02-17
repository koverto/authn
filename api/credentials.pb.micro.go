// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: credentials.proto

package credentials

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/koverto/uuid"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Credentials service

type CredentialsService interface {
	Create(ctx context.Context, in *Credential, opts ...client.CallOption) (*CredentialResponse, error)
	Validate(ctx context.Context, in *Credential, opts ...client.CallOption) (*CredentialResponse, error)
	Update(ctx context.Context, in *CredentialUpdate, opts ...client.CallOption) (*CredentialResponse, error)
}

type credentialsService struct {
	c    client.Client
	name string
}

func NewCredentialsService(name string, c client.Client) CredentialsService {
	return &credentialsService{
		c:    c,
		name: name,
	}
}

func (c *credentialsService) Create(ctx context.Context, in *Credential, opts ...client.CallOption) (*CredentialResponse, error) {
	req := c.c.NewRequest(c.name, "Credentials.Create", in)
	out := new(CredentialResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsService) Validate(ctx context.Context, in *Credential, opts ...client.CallOption) (*CredentialResponse, error) {
	req := c.c.NewRequest(c.name, "Credentials.Validate", in)
	out := new(CredentialResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *credentialsService) Update(ctx context.Context, in *CredentialUpdate, opts ...client.CallOption) (*CredentialResponse, error) {
	req := c.c.NewRequest(c.name, "Credentials.Update", in)
	out := new(CredentialResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Credentials service

type CredentialsHandler interface {
	Create(context.Context, *Credential, *CredentialResponse) error
	Validate(context.Context, *Credential, *CredentialResponse) error
	Update(context.Context, *CredentialUpdate, *CredentialResponse) error
}

func RegisterCredentialsHandler(s server.Server, hdlr CredentialsHandler, opts ...server.HandlerOption) error {
	type credentials interface {
		Create(ctx context.Context, in *Credential, out *CredentialResponse) error
		Validate(ctx context.Context, in *Credential, out *CredentialResponse) error
		Update(ctx context.Context, in *CredentialUpdate, out *CredentialResponse) error
	}
	type Credentials struct {
		credentials
	}
	h := &credentialsHandler{hdlr}
	return s.Handle(s.NewHandler(&Credentials{h}, opts...))
}

type credentialsHandler struct {
	CredentialsHandler
}

func (h *credentialsHandler) Create(ctx context.Context, in *Credential, out *CredentialResponse) error {
	return h.CredentialsHandler.Create(ctx, in, out)
}

func (h *credentialsHandler) Validate(ctx context.Context, in *Credential, out *CredentialResponse) error {
	return h.CredentialsHandler.Validate(ctx, in, out)
}

func (h *credentialsHandler) Update(ctx context.Context, in *CredentialUpdate, out *CredentialResponse) error {
	return h.CredentialsHandler.Update(ctx, in, out)
}