package utils

import (
	"errors"
	"ginn/config"
	errMsg "ginn/utils/code"
	"github.com/gin-gonic/gin"

	"strconv"
)

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(config.Gconfig.Auth.CoxFlag)
	if !ok {
		err = errors.New(errMsg.ReturnCodeMsg(errMsg.CodeUserNotLogin))
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = errors.New(errMsg.ReturnCodeMsg(errMsg.CodeUserNotLogin))
		return
	}
	return
}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	var (
		page int64
		size int64
		err  error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
