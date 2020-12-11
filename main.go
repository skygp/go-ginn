package main

import (
	"ginn/api/dao/mysql"
	"ginn/api/dao/redis"
	"ginn/core"
	"ginn/package/logger"
)

/*
- Ginn-QQ群: 458314077
- 作者QQ:    1209371783
- 加群与加好友备注: ginn，谢谢！
 */
func main() {

	defer mysql.CloseDb()
	defer redis.CloseRedis()
	defer logger.Sync()

	core.Start()
}
