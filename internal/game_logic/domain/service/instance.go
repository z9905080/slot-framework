package service

import "slot-framework/internal/game_logic/domain/entity"

type service struct {
	gameMap map[string]entity.GameEntity
}

func (s *service) GetGame(gameID string) (entity.GameEntity, error) {
	return s.gameMap[gameID], nil
}

func NewGameService() InfGameService {
	return &service{
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
