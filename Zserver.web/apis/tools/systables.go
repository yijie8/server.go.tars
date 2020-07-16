package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ZserverWeb/client"
	"strings"

	tools2 "github.com/yijie8/zserver/tools"
	"ZserverWeb/Zserver"
)

// @Summary 分页列表数据
// @Description 生成表分页列表
// @Tags 工具 - 生成表
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/page [get]
func GetSysTableList(c *gin.Context) {
	var req Zserver.Table_req
	var res Zserver.SysTables_res
	req.PageSize = c.Request.FormValue("pageSize")
	req.PageIndex = c.Request.FormValue("pageIndex")
	req.TableName = c.Request.FormValue("tableName")
	req.TableComment = c.Request.FormValue("tableComment")
	err := client.WebApiNoAuth().GetSysTableList(&req, &res)
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res)
}

// @Summary 获取配置
// @Description 获取JSON
// @Tags 工具 - 生成表
// @Param configKey path int true "configKey"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/info/{tableId} [get]
// @Security
func GetSysTables(c *gin.Context) {
	var res Zserver.SysTables_one_res
	err := client.WebApiNoAuth().GetSysTables(c.Param("tableId"), &res)
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res)
}

// @Summary 添加表结构
// @Description 添加表结构
// @Tags 工具 - 生成表
// @Accept  application/json
// @Product application/json
// @Param tables query string false "tableName / 数据表名称"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sys/tables/info [post]
// @Security Bearer
func InsertSysTable(c *gin.Context) {
	var req Zserver.InsertSysTable_req
	var res Zserver.Response_res
	req.Tables = c.Request.FormValue("tables")
	err := client.WebApiNoAuth().InsertSysTable(&req, &res)
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res)
}

// @Summary 修改表结构
// @Description 修改表结构
// @Tags 工具 - 生成表
// @Accept  application/json
// @Product application/json
// @Param data body models.Dept true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sys/tables/info [put]
// @Security Bearer
func UpdateSysTable(c *gin.Context) {
	var req Zserver.SysTables_List
	var res Zserver.Response_res
	err := c.Bind(&res)
	tools2.HasError(err, "数据解析失败", -1)
	uid := tools2.GetUserIdStr(c)
	err = client.WebApiNoAuth().UpdateSysTable(uid, &req, &res)
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res)
}

// @Summary 删除表结构
// @Description 删除表结构
// @Tags 工具 - 生成表
// @Param tableId path int true "tableId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/sys/tables/info/{tableId} [delete]
func DeleteSysTables(c *gin.Context) {
	var res Zserver.Response_res
	err := c.Bind(&res)
	tools2.HasError(err, "数据解析失败", -1)
	IDS := tools2.IdsStrToIdsIntGroup("tableId", c)

	// []int to []string
	ids_string := []string{}
	for _, v := range IDS {
		ids_string = append(ids_string, tools2.IntToString(v))
	}
	err = client.WebApiNoAuth().DeleteSysTables(strings.Join(ids_string, ","), &res)
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res)
}
