package config

import (
	"fmt"
	errmsg "ginn/utils/code"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	CONFIGPATH1 = "./config"
	CONFIGPATH2 = "."
	CONFIGNAME  = "config"
	CONFIGTYPE  = "yaml"
)

var Gconfig = NewWebConfig()

type WebConfig struct {
	Server ServerInfo `mapstructure:"ServerInfo"`
	Mysql  MysqlInfo  `mapstructure:"MysqlInfo"`
	Redis  RedisInfo  `mapstructure:"RedisInfo"`
	QiNiu  QiNiuInfo  `mapstructure:"QiNiuInfo"`
	Auth   Auth       `mapstructure:"Auth"`
	Email  Email      `mapstructure:"email"`
}

type ServerInfo struct {
	Name      string `mapstructure:"name"`
	Version   string `mapstructure:"version"`
	Desc      string `mapstructure:"desc"`
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Mode      string `mapstructure:"mode"`
	StartTime string `mapstructure:"startTime"`
	MachineId int64  `mapstructure:"machineId"`
	SecretKey string `mapstructure:"secretKey"`
}

type MysqlInfo struct {
	DbName   string `mapstructure:"dbName"`
	UserName string `mapstructure:"userName"`
	PassWord string `mapstructure:"passWord"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	CharSet  string `mapstructure:"charset"`
}

type RedisInfo struct {
	DbName   int    `mapstructure:"dbName"`
	PoolSize int    `mapstructure:"poolSize"`
	PassWord string `mapstructure:"passWord"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
}

type QiNiuInfo struct {
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
	ImageUrl  string `mapstructure:"imageUrl"`
	Bucket    string `mapstructure:"bucket"`
}

type Auth struct {
	JwtExpire int    `mapstructure:"jwtExpire"`
	SecretKey string `mapstructure:"secretKey"`
	CoxFlag   string `mapstructure:"coxFlag"`
}

type Email struct {
	Addr string     `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password   string `mapstructure:"password"`
	Host   string `mapstructure:"host"`
}

func NewWebConfig() (config *WebConfig) {
	config = new(WebConfig)
	viperPkg := viper.New()
	viperPkg.AddConfigPath(CONFIGPATH1)
	viperPkg.AddConfigPath(CONFIGPATH2)
	viperPkg.SetConfigName(CONFIGNAME)
	viperPkg.SetConfigType(CONFIGTYPE)
	err := viperPkg.ReadInConfig()
	errmsg.CheckMsgWithPanic(err)
	viperPkg.WatchConfig()
	viperPkg.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("change config data")
	})
	err = viperPkg.Unmarshal(config)
	errmsg.CheckMsgWithPanic(err)
	return
}

func SplicingMysqlRetDsn() (dsn string) {
	mysqlCfg := Gconfig.Mysql
	userName := mysqlCfg.UserName
	passWord := mysqlCfg.PassWord
	port := mysqlCfg.Port
	host := mysqlCfg.Host
	dbName := mysqlCfg.DbName
	charSet := mysqlCfg.CharSet
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		userName,
		passWord,
		host,
		port,
		dbName,
		charSet,
	)
	return
}
