package system

import (
	"github.com/gin-gonic/gin"
	"github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/app"
	"ZserverWebApiNoAuth/Zserver"
	"ZserverWebApiNoAuth/models"
)

// @Summary 获取菜单树
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/menuTreeselect [get]
// @Security Bearer
func GetMenuTreeelect(c *gin.Context) {
	result, err := GetMenuTreeelect_()
	tools.HasError(err, "抱歉未找到相关信息", -1)
	app.OK(c, result, "")
}

func GetMenuTreeelect_() ([]models.MenuLable, error) {
	var data models.Menu
	result, err := data.SetMenuLable()
	return result, err
}
func GetMenuTreeelect_tar(res *Zserver.GetMenuTreeelect_res) error {
	result, err := GetMenuTreeelect_()
	if err != nil {
		res.Code = 400
		res.Msg = err.Error()
		res.Data = nil
	} else {
		res.Code = 200
		res.Msg = ""
		for _, v := range result {
			res.Data = append(res.Data, Zserver.GetMenuTreeelect_Children{
				Id:       int32(v.Id),
				Label:    v.Label,
				Children: DiguiMenuLable(&v.Children),
			})
		}
	}
	return err
}

func DiguiMenuLable(menu *[]models.MenuLable) (res []Zserver.GetMenuTreeelect_Children) {
	if menu == nil {
		return []Zserver.GetMenuTreeelect_Children{}
	}
	for _, v := range *menu {
		res = append(res, Zserver.GetMenuTreeelect_Children{
			Id:       int32(v.Id),
			Label:    v.Label,
			Children: DiguiMenuLable(&v.Children),
		})
	}
	return res
}