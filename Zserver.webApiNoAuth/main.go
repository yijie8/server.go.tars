package main

// TODO webApi 开始拆
import (
	"fmt"
	"os"

	"github.com/TarsCloud/TarsGo/tars"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/yijie8/zserver/middleware"

	"github.com/yijie8/zserver/database"
	"github.com/yijie8/zserver/global/orm"
	"github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/config"

	"ZserverWebApiNoAuth/Zserver"
	Imp "ZserverWebApiNoAuth/imp"
	middleware_auth "ZserverWebApiNoAuth/middleware"
	"ZserverWebApiNoAuth/models"
	"ZserverWebApiNoAuth/models/gorm"
	"ZserverWebApiNoAuth/router"
)

func main() {
	// Get server config
	cfg := tars.GetServerConfig()

	// New servant imp
	imp := new(Imp.WebGo_Imp)
	err := imp.Init()
	if err != nil {
		fmt.Printf("webGo_Imp init fail, err:(%s)\n", err)
		os.Exit(-1)
	}
	// New servant
	app := new(Zserver.WebApiNoAuth)
	mux := &tars.TarsHttpMux{}

	// 添加gin
	r := gin.New()

	// 发布前必须改这里
	// gin.SetMode(gin.ReleaseMode) // 正式

	// 1. 读取配置 上传配置文件解决
	if PathExists("settings.yml") {
		config.ConfigSetup("settings.yml")
	} else {
		config.ConfigSetup("/data/app/conf/Zserver_Web_settings.yml")
	}
	// if gin.Mode() == gin.ReleaseMode {
	//	config.ConfigSetup("/data/app/conf/Zserver_Web_settings.yml")
	// } else {
	//	config.ConfigSetup("settings.yml")
	// }
	// 2. 设置日志
	tools.InitLogger()
	// 3. 初始化数据库链接
	database.Setup()

	// 4. 数据库迁移
	if gin.Mode() == gin.ReleaseMode && false {
		_ = migrateModel()
		log.Println("数据库结构初始化成功！")
		// 5. 数据初始化完成
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库基础数据初始化失败！")
		}
	}

	middleware.InitMiddleware(r)

	// the jwt middleware
	authMiddleware, err := middleware_auth.AuthInit()
	tools.HasError(err, "JWT Init Error", 500)

	// 注册系统路由
	router.InitSysRouter(r, authMiddleware)

	// 注册业务路由
	// TODO: 这里可存放业务路由，里边并无实际路由是有演示代码
	// router.InitExamplesRouter(r, authMiddleware)

	// 托管给gin
	mux.HandleFunc("/", r.ServeHTTP)
	// ///////////////

	// Register http server
	tars.AddHttpServant(mux, cfg.App+"."+cfg.Server+".HttpObj")

	// Register Servant
	app.AddServantWithContext(imp, cfg.App+"."+cfg.Server+".TcpObj")

	// Run application
	tars.Run()
}

func migrateModel() error {
	if config.DatabaseConfig.Dbtype == "mysql" {
		orm.Eloquent = orm.Eloquent.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	return gorm.AutoMigrate(orm.Eloquent)
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
