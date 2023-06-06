package usecase

type InfValidator interface {
	Validate() error
}

var _ InfValidator = (*CmdGameInit)(nil)
