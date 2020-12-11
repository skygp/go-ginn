package mysql

import (
	"errors"
	"ginn/api/models"
	errMsg "ginn/utils/code"
)

func CreateUser(user *models.RequestRegisterUser) (err error) {
	var obj models.DbUser
	obj.UserName = user.UserName
	obj.PassWord = user.PassWord
	obj.Email = user.Email
	obj.Role = 1
	err = db.Create(&obj).Error
	return
}

func FindUserIsExist(name string) (isExist bool) {
	resultUser := new(models.DbUser)
	tx := db.Where("userName = ?", name).First(resultUser)
	if tx.Error != nil || resultUser.Uid == 0 || resultUser.UserName == "" {
		return false
	}
	return true
}

func FindUserName(username string) (resultUser *models.DbUser, err error) {
	resultUser = new(models.DbUser)
	tx := db.Where("userName = ?", username).First(resultUser)
	if tx.Error != nil {
		return nil, errors.New(errMsg.ReturnCodeMsg(errMsg.CodeUserNotExist))
	}
	return
}
