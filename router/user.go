package router

import (
	"ginn/api/v1/user"
	"github.com/gin-gonic/gin"
)

func NewRegisterUser(v *gin.RouterGroup) {
	userGroup := v.Group("/user")
	{
		userGroup.POST("/login", user.PostLogin)
		userGroup.POST("/register", user.PostRegister)
		userGroup.POST("/userInfo", user.PostUserInfo)
		userGroup.POST("/logout", user.PostLogout)
	}
}
