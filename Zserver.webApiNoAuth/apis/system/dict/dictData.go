package dict

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/app"
	"net/http"
	"ZserverWebApiNoAuth/Zserver"
	"ZserverWebApiNoAuth/models"
)



// @Summary 通过字典类型获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param dictType path int true "dictType"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/databyType/{dictType} [get]
// @Security
func GetDictDataByDictType(c *gin.Context) {
	result, err := GetDictDataByDictType_(c.Param("dictType"))
	tools.HasError(err, "抱歉未找到相关信息", -1)
	var res app.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}
func GetDictDataByDictType_(dictType string) ([]models.DictData, error) {
	var DictData models.DictData
	DictData.DictType = dictType
	return DictData.Get()
}
func GetDictDataByDictType_tars(dictType string, res *Zserver.DictData_res) error {
	DictData, err := GetDictDataByDictType_(dictType)
	if err != nil {
		return err
	}
	jsonx, err := json.Marshal(DictData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonx, res)
	if err != nil {
		return err
	}
	return nil
}



