package core

import (
	"context"
	"fmt"
	"ginn/api/dao/mysql"
	"ginn/config"
	"ginn/package/logger"
	"ginn/package/snowflake"
	"ginn/router"
	errmsg "ginn/utils/code"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "ginn/docs"
)

func NewEngine() *gin.Engine {
	return gin.New()
}

func Start() {
	var err error
	Engine.SetFuncMap(RegisterTemplateFunc())
	Engine.Static("/static", "static")
	Engine.LoadHTMLGlob("template/*")
	RegisterGlobalMiddle()
	RegisterRouters()
	err = mysql.RegisterModelsAndMigrate()
	errmsg.CheckMsgWithPanic(err)
	PrintServerInfo()
	RunEngine()
}

func RunEngine() {
	serverCfg := config.Gconfig.Server
	gin.SetMode(serverCfg.Mode)

	err := snowflake.Init(serverCfg.StartTime, serverCfg.MachineId)
	errmsg.CheckMsgWithPanic(err)

	srv := endless.NewServer(serverCfg.Port, Engine)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatalf("listen: %s\n", err)
			panic(fmt.Sprintf("ListenAndServe failed! err:%v\n", err))
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Logger.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("Server Shutdown:", err)
		panic(fmt.Sprintf("ShutDown failed! err:%v\n", err))
	}
	logger.Logger.Info("Server exiting")
}

//加载自定义模板函数
func RegisterTemplateFunc() template.FuncMap {
	return template.FuncMap{
		"safe": func(data string) template.HTML {
			return template.HTML(data)
		},
	}
}

func RegisterRouters() {
	Engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "NOT FOUND", "code": 404})
	})
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1 := Engine.Group("/v1")
	router.NewRegisterUser(v1)
	router.NewRegisterFile(v1)
}

func RegisterGlobalMiddle() {
	Engine.Use(
		LoggerCore(),
		GinRecovery(true),
		Cors(),
	)
}

func PrintServerInfo() {
	serverCfg := config.Gconfig.Server
	logger.Logger.Infof(`
   ___  __ __  __ __  __
  // \\ || ||\ || ||\ ||
 (( ___ || ||\\|| ||\\||
  \\_|| || || \|| || \||
    ____________________
                        `+
		"\nServerName: %s \nVersion: %s \nDesc: %s \nUrl: http://%s%s \nMode: %s",
		serverCfg.Name,
		serverCfg.Version,
		serverCfg.Desc,
		serverCfg.Host,
		serverCfg.Port,
		serverCfg.Mode,
	)
}
