package tbsdk

import (
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gutil"
)

type OutCpXg struct {
	Items []OutCp `json:"items"`
	Xg    []int64 `json:"xg"`
}

// 取相关商品列表 by key
func GetXGKey(key string) (OutCpXg, error) {
	out := OutCpXg{}
	if ar, err := GetSearch(key); err != nil {
		return out, err
	} else {
		out.Items = ar
		for _, v := range ar {
			out.Xg = append(out.Xg, v.Id)
		}
		return out, nil
	}
}

// 取相关商品列表 by cpid
func GetXGId(id int64) (OutCpXg, error) {
	out := OutCpXg{}
	if ar, err := GetCP(id); err != nil {
		return out, err
	} else {
		if ar2, err := GetSearch(ar.Title); err != nil {
			return out, err
		} else {
			out.Items = ar2
			for _, v := range ar2 {
				out.Xg = append(out.Xg, v.Id)
			}
			return out, nil
		}
	}
}

// 从商品ID取相关
func GetCpXg(id string, kw ...string) (cp OutCp, arr OutCpXg, err error) {
	kws := ""
	if !gstr.IsNumeric(id) && len(kw) > 0 {
		kws = kw[0]
		id = ""
	}
	if !gutil.IsEmpty(id) {
		cp, err = GetCP(gconv.Int64(id))
		if err != nil {
			return
		} else {
			kws = cp.Title
		}
	}
	arr, err = GetXGKey(kws)
	if err != nil {
		return
	}
	return
}

func Cplist(ty, q, c, pv, t, p string) []OutCp {
	if gutil.IsEmpty(ty) {
		ty = "1"
	}
	if gutil.IsEmpty(c) {
		c = "0"
	}
	if gutil.IsEmpty(pv) {
		pv = "total_sales"
	}
	if gutil.IsEmpty(p) {
		p = "1"
	}
	// 列表
	if ty == "1" {
		args := make(map[string]interface{})
		if gutil.IsEmpty(t) {
			t = "2"
		}
		if !gutil.IsEmpty(q) {
			args["q"] = q
		}
		if c != "0" {
			args["cat"] = c
		}
		args["sort"] = pv
		if t == "1" {
			args["is_tmall"] = "true"
		}
		args["page_no"] = p
		res, err := NewTbsdk().GetCpList(args).ToStruct()
		if err != nil {
			glog.Error(err)
		}
		return Cpjson2OutCp(res.(GetCpListJson))
		// 单商品
	} else if ty == "2" {
		id := q
		cps, err := NewTbsdk().GetCp(id).ToStruct()
		if err != nil {
			glog.Error(err)
		}
		cp := cps.(CpOnly)
		return []OutCp{
			{
				Id:        cp.NumIid,
				Price:     gconv.Float64(cp.ZkFinalPrice),
				Yprice:    gconv.Float64(cp.ReservePrice),
				Title:     cp.Title,
				Wangwang:  cp.Nick,
				Chengjiao: cp.Volume,
				Pic:       ToChangePic(cp.PictUrl).(string),
				Uid:       cp.SellerId,
			},
		}
	} else if ty == "3" {
		id := q
		cps, err := NewTbsdk().GetCpXgList(gconv.Int64(id)).ToStruct()
		if err != nil {
			glog.Error(err)
		}
		cplist := cps.([]CpOnly)
		var cpall []OutCp
		for _, v := range cplist {
			cpall = append(cpall, OutCp{
				Id:        v.NumIid,
				Price:     gconv.Float64(v.ZkFinalPrice),
				Yprice:    gconv.Float64(v.ReservePrice),
				Title:     v.Title,
				Wangwang:  v.Nick,
				Chengjiao: v.Volume,
				Pic:       ToChangePic(v.PictUrl).(string),
				Uid:       v.SellerId,
			})
		}
		return cpall
	}
	return nil
}

func Cpjson2OutCp(json GetCpListJson) (cpall []OutCp) {
	for _, v := range json.ResultList.MapData {
		cpall = append(cpall, OutCp{
			Id:        v.NumIid,
			Price:     gconv.Float64(v.ZkFinalPrice),
			Yprice:    gconv.Float64(v.ReservePrice),
			Title:     v.Title,
			Wangwang:  v.Nick,
			Chengjiao: v.Volume,
			Pic:       ToChangePic(v.PictURL).(string),
			Uid:       v.SellerID,
		})
	}
	return
}
