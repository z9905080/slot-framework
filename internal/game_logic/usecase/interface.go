package usecase

type InfGameLogic interface {
	GameInit(initCmd CmdGameInit) (EventGameInit, error)
}
