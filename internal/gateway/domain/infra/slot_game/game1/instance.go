package game1

import (
	"github.com/shopspring/decimal"
	"slot-framework/internal/gateway/domain/entity"
	"slot-framework/internal/gateway/domain/infra/slot_game/base"
	"slot-framework/internal/gateway/domain/repository"
)

type rtp struct {
	base.BaseRTP
}

func (r *rtp) GetSpinResult(gameInput entity.SlotGameInput) ([]entity.SlotWinLine, error) {
	return []entity.SlotWinLine{
		{
			WinType:       1,
			LineNo:        0,
			Credit:        0,
			Multiply:      0,
			SymbolID:      0,
			Count:         0,
			IsWinPosition: nil,
			PayComboID:    0,
			WinPoint:      decimal.Decimal{},
			WinLineCount:  0,
		},
	}, nil
}

func NewRTP() repository.InfGameRTPService {
	return &rtp{}
}
