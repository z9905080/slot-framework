package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/z9905080/melody"
	usecase "slot-framework/internal/gateway/usecase"
)

type Handler struct {
	App          *gin.Engine
	WebsocketApp *melody.Melody
	GatewaySrv   usecase.InfGateway
}

func (h *Handler) SetWebsocketRouter() {
	h.WebsocketApp.HandleMessage(func(s *melody.Session, msg []byte) {
		h.ClientRequestHandler(s, msg)
	})

	h.WebsocketApp.HandleMessageBinary(func(s *melody.Session, msg []byte) {
		h.ClientRequestHandler(s, msg)
	})

	h.WebsocketApp.HandleConnect(func(s *melody.Session) {
		s.Write([]byte("hello"))
	})
}

func (h *Handler) ClientRequestHandler(s *melody.Session, msg []byte) {
	funcMap := h.getServerFuncMap()
	if fn, isExist := funcMap["game_flow"]; isExist {
		fn(usecase.CmdOperation{
			Server:  h.WebsocketApp,
			Session: s,
			Data:    string(msg),
		})
	}
}

func (h *Handler) getServerFuncMap() map[string]func(flow usecase.CmdOperation) (usecase.EventOperation, error) {
	return map[string]func(usecase.CmdOperation) (usecase.EventOperation, error){
		"game_flow": h.GatewaySrv.GameFlow,
	}
}
