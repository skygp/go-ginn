package file

import (
	"ginn/package/qiniuyun"
	errMsg "ginn/utils/code"
	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := qiniuyun.UploadFile(file, fileSize)
	if code == errMsg.CodeError {
		errMsg.ResponseError(c, errMsg.CodeUploadFileFailed)
	}
	errMsg.ResponseSuccess(c, errMsg.CodeSuccess, url)
}
