package service

import (
	"slot-framework/internal/player/domain/entity"
	"slot-framework/pkg/logger"
)

type service struct {
	log logger.Logger
}

func (s *service) GetPlayer(playerID string) (entity.PlayerEntity, error) {
	s.log.InfoF("GetPlayer: %v", playerID)
	return entity.PlayerEntity{
		ID:   playerID,
		Name: "test",
	}, nil
}

func NewPlayerService(log logger.Logger) InfPlayerService {
	return &service{log: log}
}
