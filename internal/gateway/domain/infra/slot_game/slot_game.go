package slot_game

import (
	"slot-framework/internal/gateway/domain/infra/slot_game/game1"
	"slot-framework/internal/gateway/domain/repository"
)

type gameManager struct {
	gameMap map[string]func() repository.InfGameRTPService
}

func (g *gameManager) NewGameModule(gameCode string) repository.InfGameRTPService {
	return g.gameMap[gameCode]()
}

func NewGameManager() repository.InfGameManagerService {
	return &gameManager{
		gameMap: map[string]func() repository.InfGameRTPService{
			"1": game1.NewRTP,
		},
	}
}
