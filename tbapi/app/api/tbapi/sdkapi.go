package tbapi

import (
	"TBapi/app/service/tbsdk"
	"TBapi/library/response"
	"TBapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"net/http"
)

func CpList(c *gin.Context) {
	type queryValid_ struct {
		C           string `form:"c"` //分类
		Q           string `form:"q" binding:"required_without=C"`
		CC          string `form:"cc"`            //排序
		P           string `form:"p"`             //分页
		Zid         string `form:"zid"`           //站点的ID 62a等
		MaterialId  string `form:"material_id"`   //物料ID
		StartTkRate string `form:"start_tk_rate"` //开始的提成
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(c, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}

	argsMap := make(map[string]interface{})
	argsMap["q"] = queryValid.Q
	argsMap["page_no"] = gconv.Int(queryValid.P)
	if queryValid.CC != "" {
		argsMap["sort"] = queryValid.CC
	}
	if queryValid.C != "" {
		argsMap["cat"] = queryValid.C
	}
	if queryValid.MaterialId != "" {
		argsMap["material_id"] = gconv.Int(queryValid.MaterialId)
	}
	if queryValid.StartTkRate != "" {
		argsMap["start_tk_rate"] = gconv.Int(queryValid.StartTkRate)
	}

	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetCpList(argsMap, queryValid.Zid).ToStruct()
	})
}

func DPXgByKey(c *gin.Context) {
	type queryValid_ struct {
		Key  string `form:"key" binding:"required"`
		Page string `form:"p"` //分页
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetDPXgListKey(queryValid.Key, gconv.Int(queryValid.Page)).ToStruct()
	})
}

// 没有数据 TODO
func DPXgById(c *gin.Context) {
	type queryValid_ struct {
		Id string `form:"id" binding:"required"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}

	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetDPXgListId(gconv.Int(queryValid.Id)).ToStruct()
	})
}

func CpXgList(c *gin.Context) {
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
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetCpXgList(queryValid.Id).ToStruct()
	})
}

func Cp(c *gin.Context) {
	type queryValid_ struct {
		Ids string `form:"id" binding:"required"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}

	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetCp(queryValid.Ids).ToStruct()
	})
}

func Tgg(c *gin.Context) {
	type queryValid_ struct {
		StartTime string `form:"start_time"`
		EndTime   string `form:"end_time"`
		PageNo    int    `form:"page_no"`
		PageSize  int    `form:"page_size"`
		Zid       string `form:"zid"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetTGG(gconv.Map(queryValid), queryValid.Zid).ToStruct()
	})
}

func Tkl(c *gin.Context) {
	type queryValid_ struct {
		Text string `form:"text" binding:"required"`
		Url  string `form:"url" binding:"required"`
		Zid  string `form:"zid"`
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
			return tbsdk.NewTbsdk().GetTKL(queryValid.Text, queryValid.Url, queryValid.Zid)
		})
}

func Content(c *gin.Context) {
	type queryValid_ struct {
		Type            string `form:"type"`
		BeforeTimestamp int64  `form:"before_timestamp"`
		Count           int    `form:"count"`
		Cid             int    `form:"cid"`
		ImageWidth      int    `form:"image_width"`
		ImageHeight     int    `form:"image_height"`
		Zid             string `form:"zid"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetContent(gconv.Map(queryValid), queryValid.Zid).ToStruct()
	})
}

func CpGood(c *gin.Context) {
	type queryValid_ struct {
		MaterialId    int    `form:"type"`
		ContentId     int    `form:"content_id"`
		ContentSource string `form:"content_source"`
		ItemId        string `form:"item_id"`
		PageSize      int    `form:"page_size"`
		PageNo        int    `form:"page_no"`
		Zid           string `form:"zid"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetCpListGood(gconv.Map(queryValid), queryValid.Zid).ToStruct()
	})
}

func CouponDesc(c *gin.Context) {
	type queryValid_ struct {
		ActivityId string `form:"activity_id" binding:"required"`
		ItemId     string `form:"item_id" binding:"required"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetCpListGood(gconv.Map(queryValid)).ToStruct()
	})
}

func CreateTlj(c *gin.Context) {
	type queryValid_ struct {
		TotalNum             int    `form:"total_num"`
		ItemId               string `form:"item_id" binding:"required"`
		Name                 string `form:"name"`
		UserTotalWinNumLimit int    `form:"user_total_win_num_limit"`
		PerFace              int    `form:"per_face"`
		SendStartTime        string `form:"send_start_time"`
		SendEndTime          string `form:"send_end_time"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetCpListGood(gconv.Map(queryValid)).ToStruct()
	})
}

func TljDesc(c *gin.Context) {
	type queryValid_ struct {
		RightsId string `form:"rights_id" binding:"required"`
		Zid      string `form:"zid"`
	}
	queryValid := queryValid_{}
	if err := c.ShouldBind(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
		return
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetTLJDesc(queryValid.RightsId, queryValid.Zid).ToStruct()
	})
}
