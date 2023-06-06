package usecase

import "github.com/z9905080/melody"

type EventOperation struct {
	server  *melody.Melody
	session *melody.Session
	data    string
}
