package models

import (
	"ginn/package/snowflake"
	"ginn/utils"
	"gorm.io/gorm"
)

var roles = map[int]string{
	1:"admin",
	0:"ordinary",
}

func GetRole(code int) string {
	return roles[code]
}

type DbUser struct {
	gorm.Model
	Uid      int64  `gorm:"primaryKey;column:uid"  json:"-"`
	UserName string `gorm:"size:64;not null;unique;column:userName" json:"userName" form:"userName" binding:"required"`
	PassWord string `gorm:"size:128;not null;unique;column:passWord" json:"passWord" form:"passWord" binding:"required"`
	Email    string `gorm:"size:64;not null;unique;column:email" json:"email" form:"email" binding:"required"`
	Role     int    `gorm:"size:32;not null;default:1;column:role" json:"role" form:"role" binding:"-"`
}

func (*DbUser) TableName() string {
	return "user"
}

func (d *DbUser) BeforeCreate(tx *gorm.DB) (err error) {
	d.Uid = snowflake.GenID()
	d.PassWord, err = utils.GenerateFromPassword(d.PassWord)
	return
}
