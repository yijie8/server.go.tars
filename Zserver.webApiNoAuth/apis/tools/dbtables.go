package tools

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	tools2 "github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/app"
	config2 "github.com/yijie8/zserver/tools/config"
	"ZserverWebApiNoAuth/Zserver"
	"ZserverWebApiNoAuth/models/tools"
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
	var res app.Response
	var err error
	var pageSize = 10
	var pageIndex = 1
	if config2.DatabaseConfig.Dbtype == "sqlite3" {
		res.Msg = "对不起，sqlite3 暂不支持代码生成！"
		c.JSON(http.StatusOK, res.ReturnError(500))
		return
	}
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools2.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools2.StrToInt(err, index)
	}
	res, err = GetDBTableList_(pageSize, pageIndex, c.Request.FormValue("tableName"))
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetDBTableList_(pageSize int, pageIndex int, tableName string) (res app.Response, err error) {
	var data tools.DBTables
	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex == 0 {
		pageIndex = 1
	}
	if config2.DatabaseConfig.Dbtype == "sqlite3" {
		res.Msg = "对不起，sqlite3 暂不支持代码生成！"
		err = errors.New(res.Msg)
		return
	}
	data.TableName = tableName
	result, count, err := data.GetPage(pageSize, pageIndex)
	if err != nil {
		return
	}
	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize
	res.Data = mp
	return
}

func GetDBTableList_tars(req *Zserver.GetDBTableList_req, res *Zserver.GetDBTableList_res) (err error) {
	resx, err := GetDBTableList_(tools2.StrToInt(err, req.PageSize), tools2.StrToInt(err, req.PageIndex), req.TableName)
	if err != nil {
		return err
	}
	resx_map := resx.Data.(map[string]interface{})
	res.Msg = resx.Msg
	res.Code = int32(resx.Code)
	res.Data.PageIndex = int32(resx_map["pageIndex"].(int))
	res.Data.PageSize = int32(resx_map["pageSize"].(int))
	res.Data.Count = int32(resx_map["count"].(int))

	jsonx, err := json.Marshal(resx_map["list"])
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonx, &res.Data.List)
	if err != nil {
		return err
	}

	return
}
