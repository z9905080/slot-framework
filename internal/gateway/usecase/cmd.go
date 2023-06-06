package usecase

import "github.com/z9905080/melody"

type CmdOperation struct {
	Server  *melody.Melody
	Session *melody.Session
	Data    string
}
