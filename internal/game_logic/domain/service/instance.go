package service

import (
	"context"
	"slot-framework/internal/game_logic/domain/entity"
	"slot-framework/pkg/logger"
	protoGamePlayer "slot-framework/pkg/protobuf/proto_gen/player"
)

type service struct {
	gameMap      map[string]entity.GameEntity
	rpcPlayerSrv protoGamePlayer.PlayerService
	log          logger.Logger
}

func (s *service) GetGame(gameID string) (entity.GameEntity, error) {
	resp, err := s.rpcPlayerSrv.GetPlayerInfo(context.Background(), &protoGamePlayer.GetPlayerInfoRequest{
		PlayerId: gameID,
	})
	if err != nil {
		s.log.Error(err)
		return entity.GameEntity{}, err
	}

	s.log.Info(resp.GetPlayerId())

	return s.gameMap[gameID], nil
}

func NewGameService(log logger.Logger, rpcPlayerSrv protoGamePlayer.PlayerService) InfGameService {
	return &service{
		log:          log,
		rpcPlayerSrv: rpcPlayerSrv,
		gameMap: map[string]entity.GameEntity{
			"1": {
				ID:   "1",
				Name: "game1",
			},
			"2": {
				ID:   "2",
				Name: "game2",
			},
		},
	}
}
