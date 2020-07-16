package tools

import (
	"net/http"
	"ZserverWeb/client"

	"github.com/gin-gonic/gin"

	tools2 "github.com/yijie8/zserver/tools"
	"ZserverWeb/Zserver"
)

// @Summary 分页列表数据 / page list data
// @Description 数据库表列分页列表 / database table column page list
// @Tags 工具 / Tools
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/db/columns/page [get]
func GetDBColumnList(c *gin.Context) {
	var req Zserver.GetDBColumnList_req
	var res Zserver.GetDBColumnList_res
	req.PageSize = c.Request.FormValue("pageSize")
	req.PageIndex = c.Request.FormValue("pageIndex")
	req.TableName = c.Request.FormValue("tableName")
	err := client.WebApiNoAuth().GetDBColumnList(&req, &res)
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res)
}
