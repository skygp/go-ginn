package core

import (
	"github.com/gin-gonic/gin"
)

var (
	Engine *gin.Engine
	//Logger *zap.SugaredLogger
)

func init() {
	Engine = NewEngine()
	//Logger = logger.NewLogger(config.Gconfig.Server.Mode)
}
