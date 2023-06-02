// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: summaryEmailService/proto/summaryEmailService.proto

package greeter

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for SummaryAccount service

func NewSummaryAccountEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SummaryAccount service

type SummaryAccountService interface {
	GetSummaryAccount(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type summaryAccountService struct {
	c    client.Client
	name string
}

func NewSummaryAccountService(name string, c client.Client) SummaryAccountService {
	return &summaryAccountService{
		c:    c,
		name: name,
	}
}

func (c *summaryAccountService) GetSummaryAccount(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SummaryAccount.GetSummaryAccount", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SummaryAccount service

type SummaryAccountHandler interface {
	GetSummaryAccount(context.Context, *Request, *Response) error
}

func RegisterSummaryAccountHandler(s server.Server, hdlr SummaryAccountHandler, opts ...server.HandlerOption) error {
	type summaryAccount interface {
		GetSummaryAccount(ctx context.Context, in *Request, out *Response) error
	}
	type SummaryAccount struct {
		summaryAccount
	}
	h := &summaryAccountHandler{hdlr}
	return s.Handle(s.NewHandler(&SummaryAccount{h}, opts...))
}

type summaryAccountHandler struct {
	SummaryAccountHandler
}

func (h *summaryAccountHandler) GetSummaryAccount(ctx context.Context, in *Request, out *Response) error {
	return h.SummaryAccountHandler.GetSummaryAccount(ctx, in, out)
}