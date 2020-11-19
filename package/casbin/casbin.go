package casbin

import (
	"ginn/config"
	errmsg "ginn/utils/code"
	gormAdapter "github.com/casbin/gorm-adapter"

	"github.com/casbin/casbin"
)

var Enforcer *casbin.Enforcer

func init() {
	dsn := config.SplicingMysqlRetDsn()
	adapter := gormAdapter.NewAdapter("mysql", dsn, true)
	Enforcer = casbin.NewEnforcer("../config/rbac_model.conf", adapter)
	err := Enforcer.LoadPolicy()
	errmsg.CheckMsgWithPanic(err)
}

/*
e.AddGroupingPolicy("a1", "admin")
e.RemovePolicy("admin", "/v1/index", "GET")
e.GetPolicy()
*/
