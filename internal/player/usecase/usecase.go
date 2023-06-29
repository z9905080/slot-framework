package usecase

import "slot-framework/internal/player/domain/service"

func NewPlayerUsecase(dSrv service.InfPlayerService) InfPlayer {
	return &usecase{
		dSrv: dSrv,
	}
}

type usecase struct {
	dSrv service.InfPlayerService
}

func (u *usecase) GetPlayer(cmd CmdPlayerGet) (EventPlayerGot, error) {
	player, err := u.dSrv.GetPlayer(cmd.PlayerID)
	if err != nil {
		return EventPlayerGot{}, err
	}

	return EventPlayerGot{
		PlayerID: player.ID,
	}, nil
}
