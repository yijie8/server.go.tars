package tbapi

import (
	"TBapi/app/service/tbsdk"
	"TBapi/library/response"
	"TBapi/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"net/http"
)

// 配合页面里的cplist
func Cplist(c *gin.Context) {
	type queryValid_ struct {
		C  string `form:"c"`
		Ty string `form:"ty"`
		Q  string `form:"q" binding:"required_without=C"`
		Pv string `form:"pv"`
		T  string `form:"t"`
		P  string `form:"p"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}

	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"),
		func() (interface{}, error) {
			res := tbsdk.Cplist(queryValid.Ty, queryValid.Q, queryValid.C, queryValid.Pv, queryValid.T, queryValid.P)
			if res == nil {
				return res, errors.New("无数据")
			} else {
				return res, nil
			}
		})
}

// 从商品ID取相关
func XgByIdKw(c *gin.Context) {
	type queryValid_ struct {
		Id string `form:"id" binding:"required"`
		Kw string
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}

	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"),
		func() (interface{}, error) {
			if cp, arr, err := tbsdk.GetCpXg(queryValid.Id, queryValid.Kw); err != nil {
				return g.Map{
					"cp":  cp,
					"arr": arr,
				}, err
			} else {
				return g.Map{
					"cp":  cp,
					"arr": arr,
				}, nil
			}
		})
}

// 取相关商品列表 by cpid
func XgById(c *gin.Context) {
	type queryValid_ struct {
		Id int64 `form:"id" binding:"required"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"),
		func() (interface{}, error) {
			return tbsdk.GetXGId(queryValid.Id)
		})
}

// 取相关商品列表 by key
func XgByKey(c *gin.Context) {
	type queryValid_ struct {
		Key string `form:"key" binding:"required"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}

	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"),
		func() (interface{}, error) {
			return tbsdk.GetXGKey(queryValid.Key)
		})
}

// 搜索商品返１０条
func Cp10(c *gin.Context) {
	type queryValid_ struct {
		Key string `form:"key" binding:"required"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}

	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"),
		func() (interface{}, error) {
			return tbsdk.GetSearch(queryValid.Key)
		})
}

// 取店铺详情
func Dp(c *gin.Context) {
	type queryValid_ struct {
		Id int `form:"id" binding:"required"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"),
		func() (interface{}, error) {
			return tbsdk.GetDP(queryValid.Id)
		})
}

// 取商品详情
func Cpweb(c *gin.Context) {
	type queryValid_ struct {
		Id int64 `form:"id" binding:"required"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"),
		func() (interface{}, error) {
			return tbsdk.GetCP(queryValid.Id)
		})
}