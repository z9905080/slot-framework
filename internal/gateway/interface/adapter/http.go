package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/z9905080/melody"
	"net/http"
	"slot-framework/environment"
	"slot-framework/internal/gateway/interface/handler"
	"slot-framework/internal/gateway/usecase"
	"slot-framework/pkg/define"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	ginRequestID "github.com/gin-contrib/requestid"
)

// NewHTTPServer
func NewHTTPServer(config environment.Config, gatewayUsecase usecase.InfGateway) http.Handler {
	app := gin.Default()
	m := melody.New(
		melody.DialChannelBufferSize(500),
		melody.DialEnableCompress(false),
	)
	// 調整緩衝區大小
	m.Config.MaxMessageSize = 0

	h := handler.Handler{
		App:          app,
		WebsocketApp: m,
		GatewaySrv:   gatewayUsecase,
	}

	app.Use(ginRequestID.New())
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:    []string{""},
	}))

	//app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// gzip 壓縮response
	app.Use(gzip.Gzip(gzip.BestCompression))

	app.GET("/ws", func(c *gin.Context) {
		sessionDataMap := make(map[string]interface{}, 0)
		sessionDataMap[define.SESSION_KEY_USER_CLIENT_IP] = c.ClientIP()
		m.HandleRequestWithKeys(c.Writer, c.Request, sessionDataMap)
		return
	})

	// set websocket router
	h.SetWebsocketRouter()

	// set router
	//h.SetRouter()

	return h.App
}
