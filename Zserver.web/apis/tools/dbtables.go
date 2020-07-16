package tools

import (
	"net/http"
	"ZserverWeb/client"

	"github.com/gin-gonic/gin"

	tools2 "github.com/yijie8/zserver/tools"
	config2 "github.com/yijie8/zserver/tools/config"
	"ZserverWeb/Zserver"
)

// @Summary 分页列表数据 / page list data
// @Description 数据库表分页列表 / database table page list
// @Tags 工具 / Tools
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/db/tables/page [get]
func GetDBTableList(c *gin.Context) {
	var req Zserver.GetDBTableList_req
	var res Zserver.GetDBTableList_res
	if config2.DatabaseConfig.Dbtype == "sqlite3" {
		res.Msg = "对不起，sqlite3 暂不支持代码生成！"
		res.Code = 500
		c.JSON(http.StatusOK, res)
		return
	}
	req.PageSize = c.Request.FormValue("pageSize")
	req.PageIndex = c.Request.FormValue("pageIndex")
	req.TableName = c.Request.FormValue("tableName")
	err := client.WebApiNoAuth().GetDBTableList(&req, &res)
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res)
}
