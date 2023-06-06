package base

import (
	"slot-framework/internal/gateway/domain/entity"
	"slot-framework/internal/gateway/domain/repository"
)

type BaseRTP struct {
}

func (b *BaseRTP) GetSpinResult(gameInput entity.SlotGameInput) ([]entity.SlotWinLine, error) {
	return nil, nil
}

func NewBaseRTP() repository.InfGameRTPService {
	return &BaseRTP{}
}
