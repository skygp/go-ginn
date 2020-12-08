package router

import (
	"ginn/api/v1/ws"
	"github.com/gin-gonic/gin"
)

func NewRegisterWS(r *gin.RouterGroup) {
	wsGroup := r.Group("/ws")
	{
		wsGroup.GET("/:channel", ws.GetConnClient)
	}
}

