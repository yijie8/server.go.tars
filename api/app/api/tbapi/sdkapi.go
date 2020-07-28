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
		C           string `json:"c"` //分类
		Q           string `json:"q" binding:"required_without=C"`
		CC          string `json:"cc"`            //排序
		P           string `json:"p"`             //分页
		Zid         string `json:"zid"`           //站点的ID 62a等
		MaterialId  string `json:"material_id"`   //物料ID
		StartTkRate string `json:"start_tk_rate"` //开始的提成
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		Key  string `json:"key" binding:"required"`
		Page string `json:"p"` //分页
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetDPXgListKey(queryValid.Key, gconv.Int(queryValid.Page)).ToStruct()
	})
}

func DPXgById(c *gin.Context) {
	type queryValid_ struct {
		Id string `json:"id" binding:"required"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		Id int64 `json:"id" binding:"required"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		Ids string `json:"id" binding:"required"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		PageNo    int    `json:"page_no"`
		PageSize  int    `json:"page_size"`
		Zid       string `json:"zid"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		Text string `json:"text" binding:"required"`
		Url  string `json:"url" binding:"required"`
		Zid  string `json:"zid"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		Type            string `json:"type"`
		BeforeTimestamp int64  `json:"before_timestamp"`
		Count           int    `json:"count"`
		Cid             int    `json:"cid"`
		ImageWidth      int    `json:"image_width"`
		ImageHeight     int    `json:"image_height"`
		Zid             string `json:"zid"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		MaterialId    int    `json:"type"`
		ContentId     int    `json:"content_id"`
		ContentSource string `json:"content_source"`
		ItemId        string `json:"item_id"`
		PageSize      int    `json:"page_size"`
		PageNo        int    `json:"page_no"`
		Zid           string `json:"zid"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		ActivityId string `json:"activity_id" binding:"required"`
		ItemId     string `json:"item_id" binding:"required"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		TotalNum             int    `json:"total_num"`
		ItemId               string `json:"item_id" binding:"required"`
		Name                 string `json:"name"`
		UserTotalWinNumLimit int    `json:"user_total_win_num_limit"`
		PerFace              int    `json:"per_face"`
		SendStartTime        string `json:"send_start_time"`
		SendEndTime          string `json:"send_end_time"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
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
		RightsId string `json:"rights_id" binding:"required"`
		Zid      string `json:"zid"`
	}
	queryValid := (*queryValid_)(nil)
	if err := c.ShouldBindJSON(&queryValid); err != nil {
		response.Json(c, http.StatusBadRequest, err.Error())
	}
	//if e := gvalid.CheckStruct(&queryValid, nil); e != nil {
	//	response.Json(r, http.StatusRequestedRangeNotSatisfiable, e.String())
	//}
	tbsdk.CacheTb(c, utils.RunFuncName()+"_"+utils.JoinAny(queryValid, "_"), func() (interface{}, error) {
		return tbsdk.NewTbsdk().GetTLJDesc(queryValid.RightsId, queryValid.Zid).ToStruct()
	})
}
