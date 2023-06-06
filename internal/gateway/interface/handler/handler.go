package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/z9905080/melody"
	usecase2 "slot-framework/internal/gateway/usecase"
)

type Handler struct {
	App          *gin.Engine
	WebsocketApp *melody.Melody
	GatewaySrv   usecase2.InfGateway
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
		fn(usecase2.CmdOperation{
			Server:  h.WebsocketApp,
			Session: s,
			Data:    string(msg),
		})
	}
}

func (h *Handler) getServerFuncMap() map[string]func(flow usecase2.CmdOperation) (usecase2.EventOperation, error) {
	return map[string]func(usecase2.CmdOperation) (usecase2.EventOperation, error){
		"game_flow": h.GatewaySrv.GameFlow,
	}
}
