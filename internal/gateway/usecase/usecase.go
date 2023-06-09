package usecase

import (
	"context"
	"slot-framework/internal/gateway/domain/repository"
	"slot-framework/pkg/logger"
	protoGameLogic "slot-framework/pkg/protobuf/proto_gen/game_logic"
)

type usecase struct {
	log          logger.Logger
	gameLogicSrv protoGameLogic.GameLogicService
	gameManager  repository.InfGameManagerService
}

func (u *usecase) GameFlow(flow CmdOperation) (EventOperation, error) {
	u.log.InfoF("GameFlow: %v", flow)

	//u.gameManager.NewGameModule(flow.Data)

	resp, err := u.gameLogicSrv.GameInit(context.Background(), &protoGameLogic.GameInitRequest{
		GameId: flow.Data,
	})
	if err != nil {
		return EventOperation{}, err
	}

	u.log.DebugF("GameInitResponse: %v", resp)

	flow.Session.Write([]byte(resp.GetGameName()))

	return EventOperation{
		server:  flow.Server,
		session: flow.Session,
		data:    resp.GetGameId(),
	}, nil

}

func NewUsecase(logger logger.Logger, service protoGameLogic.GameLogicService, gameManager repository.InfGameManagerService) InfGateway {
	return &usecase{
		log:          logger,
		gameLogicSrv: service,
		gameManager:  gameManager,
	}
}
