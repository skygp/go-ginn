package errmsg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code CodeType    `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CheckMsgWithPanic(err error) {
	if nil != err {
		panic(err)
	}
}

func ResponseMsgWithData(c *gin.Context, code CodeType, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  codeAndMsgMap[code],
		Data: data,
	})
}

func ResponseSuccess(c *gin.Context, code CodeType, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  codeAndMsgMap[code],
		Data: data,
	})
}

func ResponseError(c *gin.Context, code CodeType) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  codeAndMsgMap[code],
		Data: nil,
	})
}

func ReturnCodeMsg(code CodeType) (msg string) {
	return codeAndMsgMap[code]
}

func ResponseRedirect(c *gin.Context, location string) {
	c.Redirect(http.StatusMovedPermanently, location)
}
