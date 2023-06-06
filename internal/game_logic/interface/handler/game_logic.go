package handler

import (
	"context"
	"slot-framework/internal/game_logic/usecase"
	"slot-framework/pkg/logger"
	protoGameLogic "slot-framework/pkg/protobuf/proto_gen/game_logic"
)

type gameLogicHandler struct {
	gameLogicUsecase usecase.InfGameLogic
	log              logger.Logger
}

func (g *gameLogicHandler) GameInit(ctx context.Context, request *protoGameLogic.GameInitRequest, response *protoGameLogic.GameInitResponse) error {

	cmd := usecase.CmdGameInit{
		GameID: request.GetGameId(),
	}

	init, err := g.gameLogicUsecase.GameInit(cmd)
	if err != nil {
		return err
	}

	response.GameId = init.GameID
	response.GameName = init.GameName
	response.GameDesc = init.GameDescription

	return nil
}

func NewGameLogicHandler(gameLogicUsecase usecase.InfGameLogic, logger logger.Logger) protoGameLogic.GameLogicServiceHandler {
	return &gameLogicHandler{
		gameLogicUsecase: gameLogicUsecase,
		log:              logger,
	}
}
