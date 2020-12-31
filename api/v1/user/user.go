package user

import (
	"ginn/api/models"
	"ginn/api/service"
	"ginn/package/logger"
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
	errMsg.ResponseSuccess(c, 200, gin.H{
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

func PostUserInfo(c *gin.Context) {
	var (
		role string
		username string
		err error
	)
	tokenStruct := new(models.RequestGetToken)
	err = c.ShouldBind(tokenStruct)
	if nil != err {
		errMsg.ResponseMsgWithData(c, errMsg.CodeInvalidToken, err.Error())
		return
	}
	role,username,err = service.GetUserInfo(tokenStruct.Token)
	if nil != err{
		errMsg.ResponseMsgWithData(c, errMsg.CodeInvalidToken, err.Error())
		return
	}
	errMsg.ResponseSuccess(c, 200, gin.H{
		"permissions": append([]string{}, role),
		"username": username,
		"avatar": "https://i.gtimg.cn/club/item/face/img/2/15922_100.gif",
	})
}

func PostLogout(c *gin.Context) {
	var (
		token string
	)
	err := c.ShouldBind(&token)
	logger.Logger.Info("token: ", token)
	if nil != err {
		errMsg.ResponseMsgWithData(c, errMsg.CodeInvalidParam, err.Error())
		return
	}
	err = service.Logout(token)
	if nil != err{
		errMsg.ResponseMsgWithData(c, errMsg.CodeInvalidToken, err.Error())
		return
	}
	errMsg.ResponseSuccess(c, 200, gin.H{
	})
}