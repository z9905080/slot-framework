package repository

import "slot-framework/internal/gateway/domain/entity"

type InfGameRTPService interface {
	GetSpinResult(gameInput entity.SlotGameInput) ([]entity.SlotWinLine, error)
}

type InfGameManagerService interface {
	NewGameModule(gameCode string) InfGameRTPService
}
