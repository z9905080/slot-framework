package usecase

type InfGateway interface {
	GameFlow(flow CmdOperation) (EventOperation, error)
}
