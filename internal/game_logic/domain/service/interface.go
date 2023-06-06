package service

import "slot-framework/internal/game_logic/domain/entity"

type InfGameService interface {
	GetGame(gameID string) (entity.GameEntity, error)
}
