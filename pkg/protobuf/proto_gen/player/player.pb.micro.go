// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: player.proto

package proto_game_player

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Api Endpoints for PlayerService service

func NewPlayerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PlayerService service

type PlayerService interface {
	GetPlayerInfo(ctx context.Context, in *GetPlayerInfoRequest, opts ...client.CallOption) (*GetPlayerInfoResponse, error)
}

type playerService struct {
	c    client.Client
	name string
}

func NewPlayerService(name string, c client.Client) PlayerService {
	return &playerService{
		c:    c,
		name: name,
	}
}

func (c *playerService) GetPlayerInfo(ctx context.Context, in *GetPlayerInfoRequest, opts ...client.CallOption) (*GetPlayerInfoResponse, error) {
	req := c.c.NewRequest(c.name, "PlayerService.GetPlayerInfo", in)
	out := new(GetPlayerInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlayerService service

type PlayerServiceHandler interface {
	GetPlayerInfo(context.Context, *GetPlayerInfoRequest, *GetPlayerInfoResponse) error
}

func RegisterPlayerServiceHandler(s server.Server, hdlr PlayerServiceHandler, opts ...server.HandlerOption) error {
	type playerService interface {
		GetPlayerInfo(ctx context.Context, in *GetPlayerInfoRequest, out *GetPlayerInfoResponse) error
	}
	type PlayerService struct {
		playerService
	}
	h := &playerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PlayerService{h}, opts...))
}

type playerServiceHandler struct {
	PlayerServiceHandler
}

func (h *playerServiceHandler) GetPlayerInfo(ctx context.Context, in *GetPlayerInfoRequest, out *GetPlayerInfoResponse) error {
	return h.PlayerServiceHandler.GetPlayerInfo(ctx, in, out)
}
