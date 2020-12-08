package ws

import (
	"ginn/package/websocket"
	"github.com/gin-gonic/gin"
)

func GetConnClient(ctx *gin.Context)  {
	websocket.WebsocketManager.WsClient(ctx)
}