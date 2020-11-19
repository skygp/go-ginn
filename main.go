package main

import (
	"ginn/api/dao/mysql"
	"ginn/api/dao/redis"
	"ginn/core"
	"ginn/package/logger"
)

func main() {

	defer mysql.CloseDb()
	defer redis.CloseRedis()
	defer logger.Sync()

	core.Start()
}
