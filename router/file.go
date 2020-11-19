package router

import (
	"ginn/api/v1/file"
	"github.com/gin-gonic/gin"
)

func NewRegisterFile(v *gin.RouterGroup) {
	userGroup := v.Group("/file")
	{
		userGroup.POST("/upload", file.UploadFile)
	}
}
