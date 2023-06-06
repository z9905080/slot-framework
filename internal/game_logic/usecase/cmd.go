package usecase

import "errors"

type CmdGameInit struct {
	GameID string
}

func (c *CmdGameInit) Validate() error {
	if c.GameID == "" {
		return errors.New("game_id is empty")
	}

	return nil
}
