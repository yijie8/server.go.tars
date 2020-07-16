package tools

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	tools2 "github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/app"

	"ZserverWebApiNoAuth/Zserver"
	"ZserverWebApiNoAuth/models/tools"
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
	var data tools.SysTables
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools2.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools2.StrToInt(err, index)
	}

	data.TBName = c.Request.FormValue("tableName")
	data.TableComment = c.Request.FormValue("tableComment")

	result, count, err := GetSysTableList_(pageSize, pageIndex, data.TBName, data.TableComment)
	tools2.HasError(err, "", -1)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize

	var res app.Response
	res.Data = mp

	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetSysTableList_(pageSize int, pageIndex int, tableName string, tableComment string) (res app.Response, count int, err error) {
	var data tools.SysTables
	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex == 0 {
		pageIndex = 1
	}

	data.TBName = tableName
	data.TableComment = tableComment
	result, count, err := data.GetPage(pageSize, pageIndex)
	if err != nil {
		return res, 0, err
	}

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageSize"] = pageSize
	res.Data = mp
	return res, count, nil
}

func GetSysTableList_tars(req *Zserver.Table_req, res *Zserver.SysTables_res) (err error) {
	res.Code = 500
	if req.PageSize == "" {
		req.PageSize = "10"
	}
	if req.PageIndex == "" {
		req.PageIndex = "1"
	}
	result, count, err := GetSysTableList_(tools2.StrToInt(err, req.PageSize), tools2.StrToInt(err, req.PageIndex), req.TableName, req.TableComment)
	if err != nil {
		res.Msg = err.Error()
		return err
	}

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = tools2.StrToInt(err, req.PageIndex)
	mp["pageSize"] = tools2.StrToInt(err, req.PageSize)

	res.Code = 200
	jsonx, err := json.Marshal(mp)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonx, &res)
	if err != nil {
		return err
	}
	return nil
}

// @Summary 获取配置
// @Description 获取JSON
// @Tags 工具 - 生成表
// @Param configKey path int true "configKey"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/info/{tableId} [get]
// @Security
func GetSysTables(c *gin.Context) {
	res, err := GetSysTables_(c.Param("tableId"))
	tools2.HasError(err, "抱歉未找到相关信息", -1)
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetSysTables_(tableId string) (res app.Response, err error) {
	var data tools.SysTables
	data.TableId, _ = tools2.StringToInt(tableId)
	result, err := data.Get()
	if err != nil {
		return app.Response{}, err
	}
	res.Code = 200
	res.Data = result
	return res, nil
}

func GetSysTables_tars(tableId string, res *Zserver.SysTables_one_res) error {
	resx, err := GetSysTables_(tableId)
	if err != nil {
		res.Code = 400
		res.Msg = err.Error()
		return err
	}
	res.Code = 200
	jsonx, err := json.Marshal(resx)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonx, &res)
	if err != nil {
		return err
	}
	return nil
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
	err := InsertSysTable_(c.Request.FormValue("tables"))
	tools2.HasError(err, "", -1)

	var res app.Response
	res.Msg = "添加成功！"
	c.JSON(http.StatusOK, res.ReturnOK())
}

func InsertSysTable_(tables string) error {
	tablesList := strings.Split(tables, ",")
	for i := 0; i < len(tablesList); i++ {
		data, err := genTableInit(tablesList, i, nil)
		if err != nil {
			return err
		}
		_, err = data.Create()
		if err != nil {
			return err
		}
	}
	return nil
}
func InsertSysTable_tars(req *Zserver.InsertSysTable_req, res *Zserver.Response_res) error {
	err := InsertSysTable_(req.Tables)
	if err != nil {
		res.Msg = err.Error()
		return err
	}
	res.Msg = ""
	res.Code = 200
	return nil
}

func genTableInit(tablesList []string, i int, c *gin.Context) (tools.SysTables, error) {
	var data tools.SysTables
	var dbTable tools.DBTables
	var dbColumn tools.DBColumns
	data.TBName = tablesList[i]
	data.CreateBy = tools2.GetUserIdStr(c)

	dbTable.TableName = data.TBName
	dbtable, err := dbTable.Get()

	dbColumn.TableName = data.TBName
	tablenamelist := strings.Split(dbColumn.TableName, "_")
	for i := 0; i < len(tablenamelist); i++ {
		strStart := string([]byte(tablenamelist[i])[:1])
		strend := string([]byte(tablenamelist[i])[1:])
		data.ClassName += strings.ToUpper(strStart) + strend
		data.PackageName += strings.ToLower(strStart) + strings.ToLower(strend)
		data.ModuleName += strings.ToLower(strStart) + strings.ToLower(strend)
	}
	data.TplCategory = "crud"
	data.Crud = true

	dbcolumn, err := dbColumn.GetList()
	data.CreateBy = tools2.GetUserIdStr(c)
	data.TableComment = dbtable.TableComment
	if dbtable.TableComment == "" {
		data.TableComment = data.ClassName
	}

	data.FunctionName = data.TableComment
	data.BusinessName = data.ModuleName
	data.IsLogicalDelete = "1"
	data.LogicalDelete = true
	data.LogicalDeleteColumn = "is_del"

	data.FunctionAuthor = "wenjianzhang"
	for i := 0; i < len(dbcolumn); i++ {
		var column tools.SysColumns
		column.ColumnComment = dbcolumn[i].ColumnComment
		column.ColumnName = dbcolumn[i].ColumnName
		column.ColumnType = dbcolumn[i].ColumnType
		column.Sort = i + 1
		column.Insert = true
		column.IsInsert = "1"
		column.QueryType = "EQ"
		column.IsPk = "0"

		namelist := strings.Split(dbcolumn[i].ColumnName, "_")
		for i := 0; i < len(namelist); i++ {
			strStart := string([]byte(namelist[i])[:1])
			strend := string([]byte(namelist[i])[1:])
			column.GoField += strings.ToUpper(strStart) + strend
			if i == 0 {
				column.JsonField = strings.ToLower(strStart) + strend
			} else {
				column.JsonField += strings.ToUpper(strStart) + strend
			}
		}
		if strings.Contains(dbcolumn[i].ColumnKey, "PR") {
			column.IsPk = "1"
			column.Pk = true
			data.PkColumn = dbcolumn[i].ColumnName
			data.PkGoField = column.GoField
			data.PkJsonField = column.JsonField
		}
		column.IsRequired = "0"
		if strings.Contains(dbcolumn[i].IsNullable, "NO") {
			column.IsRequired = "1"
			column.Required = true
		}

		if strings.Contains(dbcolumn[i].ColumnType, "int") {
			column.GoType = "int"
			column.HtmlType = "input"
		} else {
			column.GoType = "string"
			column.HtmlType = "input"
		}

		data.Columns = append(data.Columns, column)
	}
	return data, err
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
	var data tools.SysTables
	var res app.Response
	uid := tools2.GetUserIdStr(c)
	err := c.Bind(&data)
	tools2.HasError(err, "数据解析失败", -1)
	err = UpdateSysTable_(uid, &data, &res)
	tools2.HasError(err, "修改错误", -1)
	c.JSON(http.StatusOK, res.ReturnOK())
}

func UpdateSysTable_(uid string, data *tools.SysTables, res *app.Response) error {
	data.UpdateBy = uid // 当前用户的ID tools2.GetUserIdStr(nil) // TODO
	result, err := data.Update()
	if err != nil {
		return err
	}
	res.Data = result
	res.Msg = "修改成功"
	return nil
}

func UpdateSysTable_tars(uid string, req *Zserver.SysTables_List, res *Zserver.Response_res) error {
	var data tools.SysTables
	var resx app.Response
	jsonx, err := json.Marshal(req)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonx, &data)
	if err != nil {
		return err
	}

	err = UpdateSysTable_(uid, &data, &resx)
	if err != nil {
		res.Msg = err.Error()
		return err
	}
	res.Msg = ""
	res.Code = 200
	return nil
}

// @Summary 删除表结构
// @Description 删除表结构
// @Tags 工具 - 生成表
// @Param tableId path int true "tableId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/sys/tables/info/{tableId} [delete]
func DeleteSysTables(c *gin.Context) {
	var res app.Response
	err := DeleteSysTables_(c.Param("tableId"), &res)
	tools2.HasError(err, "删除失败", 500)
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

func DeleteSysTables_(tableId string, res *app.Response) error {
	var data tools.SysTables
	IDS := IdsStrToIdsIntGroup_(tableId)
	_, err := data.BatchDelete(IDS)
	if err != nil {
		return err
	}
	res.Code = 200
	res.Msg = "删除成功"
	return nil
}
func DeleteSysTables_tars(tableId string, res *Zserver.Response_res) error {
	var resx app.Response
	err := DeleteSysTables_(tableId, &resx)
	if err != nil {
		return err
	}
	res.Code = int32(resx.Code)
	res.Msg = resx.Msg
	return nil
}

func IdsStrToIdsIntGroup_(keys string) []int {
	IDS := make([]int, 0)
	ids := strings.Split(keys, ",")
	for i := 0; i < len(ids); i++ {
		ID, _ := tools2.StringToInt(ids[i])
		IDS = append(IDS, ID)
	}
	return IDS
}
