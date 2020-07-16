package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/yijie8/zserver/handler"
	jwt "github.com/yijie8/zserver/pkg/jwtauth"

	"ZserverWebApiNoAuth/apis/monitor"
	"ZserverWebApiNoAuth/apis/system"
	"ZserverWebApiNoAuth/apis/system/dict"
	. "ZserverWebApiNoAuth/apis/tools"
	_ "ZserverWebApiNoAuth/docs"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	//sysBaseRouter(g)
	// 静态文件
	//sysStaticFileRouter(g)

	// swagger；注意：生产环境可以注释掉
	sysSwaggerRouter(g)

	// 无需认证
	sysNoCheckRoleRouter(g)

	return g
}

func sysBaseRouter(r *gin.RouterGroup) {
	r.GET("/", system.HelloWorld)
	r.GET("/info", handler.Ping)
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	r.Static("/static", "./static")
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.GET("/monitor/server", monitor.ServerInfo)
	v1.GET("/getCaptcha", system.GenerateCaptchaHandler)
	v1.GET("/gen/preview/:tableId", Preview)
	v1.GET("/menuTreeselect", system.GetMenuTreeelect)
	v1.GET("/dict/databytype/:dictType", dict.GetDictDataByDictType)

	registerDBRouter(v1)

	registerSysTableRouter(v1)

}

func registerDBRouter(api *gin.RouterGroup) {
	db := api.Group("/db")
	{
		db.GET("/tables/page", GetDBTableList)
		db.GET("/columns/page", GetDBColumnList)
	}
}

func registerSysTableRouter(v1 *gin.RouterGroup) {
	systables := v1.Group("/sys/tables")
	{
		systables.GET("/page", GetSysTableList) // TODO 做到这了
		tablesinfo := systables.Group("/info")
		{
			tablesinfo.POST("", InsertSysTable)
			tablesinfo.PUT("", UpdateSysTable)
			tablesinfo.DELETE("/:tableId", DeleteSysTables)
			tablesinfo.GET("/:tableId", GetSysTables)
		}
	}
}
