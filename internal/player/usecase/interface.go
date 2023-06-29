package usecase

type InfPlayer interface {
	GetPlayer(initCmd CmdPlayerGet) (EventPlayerGot, error)
}
