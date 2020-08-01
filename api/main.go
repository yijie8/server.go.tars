package main

// TODO webApi
import (
	"TBapi/router"
	"github.com/yijie8/zserver/middleware"
	"log"
	"os"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/gin-gonic/gin"
	"github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/config"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()
	mux := &tars.TarsHttpMux{}

	// 添加gin
	r := gin.New()

	// 发布前必须改这里
	//gin.SetMode(gin.ReleaseMode) //正式

	// 1. 读取配置
	if PathExists("settings.yml") {
		config.ConfigSetup("settings.yml")
	} else {
		config.ConfigSetup("/usr/local/app/tars/tarsnode/data/Zserver_Web_settings.yml")
	}
	// 2. 设置日志
	tools.InitLogger()

	//gcache 本地缓存配置 TODO
	//_ = gcache.New()

	middleware.InitMiddleware(r)
	// the jwt middleware
	//authMiddleware, err := middleware_auth.AuthInit()
	//tools.HasError(err, "JWT Init Error", 500)

	// 注册系统路由
	router.InitSysRouter(r)

	// 托管给gin
	mux.HandleFunc("/", r.ServeHTTP)
	// ///////////////

	// Register http server
	tars.AddHttpServant(mux, cfg.App+"."+cfg.Server+".HttpObj")
	log.Println(tars.GetServerConfig())

	// Register Servant
	// app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".TcpObj")

	// Run application
	tars.Run()
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
