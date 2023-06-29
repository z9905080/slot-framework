package service

import "slot-framework/internal/player/domain/entity"

type InfPlayerService interface {
	GetPlayer(playerID string) (entity.PlayerEntity, error)
}
