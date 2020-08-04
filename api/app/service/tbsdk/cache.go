package tbsdk

import (
	"TBapi/library/response"
	"TBapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"net/http"
)

func CacheTb(c *gin.Context, key string, f func() (interface{}, error)) {
	rr := gcache.GetOrSetFuncLock(key, func() interface{} {
		if res, err := f(); err != nil {
			return err
		} else {
			if _, ok := res.(error); ok {
				response.Json(c, http.StatusMethodNotAllowed, res.(error).Error())
			} else {
				switch res.(type) {
				case []OutCp:
					res = ToChangePic(res)
				case g.Map:
					ss := res.(g.Map)
					if _, ok := ss["cp"]; ok {
						if _, ok := ss["cp"].(OutCp); ok {
							ss["cp"] = ToChangePic(ss["cp"])
						}
					}
					if _, ok := ss["arr"]; ok {
						if _, ok := ss["arr"].(OutCpXg); ok {
							ss["arr"] = ToChangePic(ss["arr"])
						}
					}
				case GetCpListJson:
					res = ToChangePic(res)
				case GetDPXgListKeyJson:
					res = ToChangePic(res)
				case GetContentJson:
					res = ToChangePic(res)
				case []CpOnly:
					res = ToChangePic(res)
				case CpOnly:
					res = ToChangePic(res)
				}
			}
			return res
		}
	}, 0) //time.Second*viper.GetDuration("settings.redis.cacheDefaultTime")

	response.Json(c, http.StatusOK, "", rr)
}

func ToChangePic(rr interface{}) (res interface{}) {
	switch rr.(type) {
	case []OutCp:
		ss := rr.([]OutCp)
		if len(ss) > 0 {
			for k, v := range ss {
				ss[k].Pic = utils.GetPic(v.Pic)
			}
		}
		res = ss
	case OutCp:
		ss := rr.(OutCp)
		ss.Pic = utils.GetPic(ss.Pic)
		res = ss
	case OutCpXg:
		ss := rr.(OutCpXg)
		if len(ss.Items) > 0 {
			ss.Items = ToChangePic(ss.Items).([]OutCp)
		}
		res = ss
	case GetCpListJson:
		ss := rr.(GetCpListJson)
		if len(ss.ResultList.MapData) > 0 {
			for k, v := range ss.ResultList.MapData {
				ss.ResultList.MapData[k].PictURL = utils.GetPic(v.PictURL)
				if len(v.SmallImages.String) > 0 {
					for kk, vv := range ss.ResultList.MapData[k].SmallImages.String {
						ss.ResultList.MapData[k].SmallImages.String[kk] = utils.GetPic(vv)
					}
				}
			}
		}
		res = ss
	case GetDPXgListKeyJson:
		ss := rr.(GetDPXgListKeyJson)
		if len(ss.Results.NTbkShop) > 0 {
			for k, v := range ss.Results.NTbkShop {
				ss.Results.NTbkShop[k].PictURL = utils.GetPic(v.PictURL)
			}
		}
		res = ss
	case GetContentJson:
		ss := rr.(GetContentJson)
		if len(ss.Contents.MapData) > 0 {
			for k, v := range ss.Contents.MapData {
				if len(v.Images.String) > 0 {
					for kk, vv := range ss.Contents.MapData[k].Images.String {
						ss.Contents.MapData[k].Images.String[kk] = utils.GetPic(vv)
					}
				}
			}
		}
		res = ss
	case []CpOnly:
		ss := rr.([]CpOnly)
		if len(ss) > 0 {
			for k, v := range ss {
				ss[k].PictUrl = utils.GetPic(v.PictUrl)
			}
		}
		res = ss
	case CpOnly:
		ss := rr.(CpOnly)
		ss.PictUrl = utils.GetPic(ss.PictUrl)
		res = ss
	case string:
		res = utils.GetPic(gconv.String(rr))
	}

	return
}
