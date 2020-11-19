package models

type BaseUser struct {
	UserName string `json:"userName" form:"userName" binding:"required"`
	PassWord string `json:"passWord" form:"passWord" binding:"required"`
}

type RequestLoginUser struct {
	*BaseUser
}

type RequestRegisterUser struct {
	Email string `json:"email" form:"email" binding:"required"`
	*BaseUser
}
