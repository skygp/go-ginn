package user

import (
	"ginn/api/models"
	"ginn/api/service"
	errMsg "ginn/utils/code"
	"github.com/gin-gonic/gin"
)

func PostLogin(c *gin.Context) {
	var (
		user  models.RequestLoginUser
		token string
	)
	err := c.ShouldBind(&user)
	if nil != err {
		errMsg.ResponseMsgWithData(c, errMsg.CodeInvalidPassword, err.Error())
		return
	}
	token, err = service.LoginUser(&user)
	if nil != err {
		errMsg.ResponseMsgWithData(c, errMsg.CodeInvalidPassword, err.Error())
		return
	}
	errMsg.ResponseSuccess(c, errMsg.CodeSuccess, gin.H{
		"token": token,
	})
}

func PostRegister(c *gin.Context) {
	var (
		user models.RequestRegisterUser
		err  error
	)
	err = c.ShouldBind(&user)
	if err != nil {
		errMsg.ResponseMsgWithData(c, errMsg.CodeInvalidParam, err.Error())
		return
	}
	err = service.RegisterUser(&user)
	if err != nil {
		errMsg.ResponseMsgWithData(c, errMsg.CodeInvalidParam, err.Error())
		return
	}
	errMsg.ResponseSuccess(c, errMsg.CodeSuccess, nil)
}
