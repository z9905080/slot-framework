package handler

import (
	"context"
	"slot-framework/internal/player/usecase"
	"slot-framework/pkg/logger"
	protoGamePlayer "slot-framework/pkg/protobuf/proto_gen/player"
)

type handler struct {
	playerUsecase usecase.InfPlayer
	log           logger.Logger
}

func (g *handler) GetPlayerInfo(ctx context.Context, request *protoGamePlayer.GetPlayerInfoRequest, response *protoGamePlayer.GetPlayerInfoResponse) error {
	e, _ := g.playerUsecase.GetPlayer(usecase.CmdPlayerGet{
		PlayerID: request.GetPlayerId(),
	})
	response.PlayerId = e.PlayerID
	return nil
}

func NewPlayerHandler(u usecase.InfPlayer, logger logger.Logger) protoGamePlayer.PlayerServiceHandler {
	return &handler{
		playerUsecase: u,
		log:           logger,
	}
}
