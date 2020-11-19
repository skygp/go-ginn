package mysql

import (
	"database/sql"
	"ginn/api/models"
	"ginn/config"
	errmsg "ginn/utils/code"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db    *gorm.DB
	sqlDb *sql.DB
)

func init() {
	db, sqlDb = NewDB()
}

func NewDB() (db *gorm.DB, sqlDb *sql.DB) {
	var err error
	dsn := config.SplicingMysqlRetDsn()
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errmsg.CheckMsgWithPanic(err)
	sqlDb, err = db.DB()
	errmsg.CheckMsgWithPanic(err)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Second * 10)
	return
}

func RegisterModelsAndMigrate() (err error) {
	err = db.AutoMigrate(
		models.DbUser{},
	)
	return
}

func CloseDb() {
	_ = sqlDb.Close()
}
