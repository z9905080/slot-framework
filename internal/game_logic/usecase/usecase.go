package usecase

import "slot-framework/internal/game_logic/domain/service"

type usecase struct {
	domainGameService service.InfGameService
}

func (u *usecase) GameInit(initCmd CmdGameInit) (EventGameInit, error) {
	game, err := u.domainGameService.GetGame(initCmd.GameID)
	if err != nil {
		return EventGameInit{}, err
	}

	return EventGameInit{
		GameID:          game.ID,
		GameName:        game.Name,
		GameDescription: "test",
	}, err
}

func NewGameLogicUsecase(domainGameService service.InfGameService) InfGameLogic {
	return &usecase{
		domainGameService: domainGameService,
	}
}
