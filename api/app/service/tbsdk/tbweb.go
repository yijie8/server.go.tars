package tbsdk

import (
	"TBapi/utils"
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"github.com/gogf/gf/os/glog"
)

type OutCp struct {
	Id        int64   `json:"id"`
	Price     float64 `json:"price"`
	Yprice    float64 `json:"yprice"`
	Title     string  `json:"title"`
	Wangwang  string  `json:"wangwang"`
	Chengjiao int     `json:"chengjiao"`
	Pic       string  `json:"pic"`
	Uid       int     `json:"uid"`
}

type productDetail struct {
	DsCust          int     `json:"ds_cust"`
	DsDiscountPrice float64 `json:"ds_discount_price"` // 价格
	DsDiscountRate  float64 `json:"ds_discount_rate"`  // 提成比例
	// ds_exchange7 int
	// "ds_genuine": 1,
	DsImg struct {
		Src string `json:"src"`
	} `json:"ds_img"`

	// 	{
	// 	"src": "http:\/\/img02.taobaocdn.com\/bao\/uploaded\/i1\/3199953386\/O1CN01st9j7z1aspeN2KU8m_!!3199953386-0-pixelsss.jpg"
	// },
	// 	"ds_istmall": 1,
	// 	"ds_item_click": "",
	DsNick string `json:"ds_nick"`
	DsNid  int64  `json:"ds_nid"`

	// "ds_post24": 0,
	// "ds_postfee": 0.00,
	// "ds_provcity": "\u6d59\u6c5f \u676d\u5dde",
	// "ds_rank": 10,
	DsReservePrice float64 `json:"ds_reserve_price"` // 市场价
	DsSell         int     `json:"ds_sell"`
	// "ds_shipping": 1,
	// "ds_shop_click": "",
	// "ds_taoke": 1,
	DsTitle  string `json:"ds_title"`
	DsUserId int    `json:"ds_user_id"`
}
type JsonxCp struct {
	Code int
	Data struct{ Items []productDetail } // []map[string]interface{}
}

// 搜索商品
func GetSearch(key string) ([]OutCp, error) {
	glog.Println(key, "<<<<<<<<<<<<<<<<<<<<<<<<<GetSearch")
	if key == "" {
		return nil, nil
	}
	u := url.Values{}
	u.Add("cb", "jsonp_callback_03822715911819512")
	u.Add("pid", "mm_12462245_2219662_23014024")
	u.Add("wt", "2")
	u.Add("ti", "176")
	u.Add("tl", "790x675")
	u.Add("rd", "2")
	u.Add("st", "2")
	u.Add("rf", "http://www.quchaoshi.com")
	u.Add("pgid", "6ed70a3508d2e0539c43ab4c3bcb9615")
	u.Add("v", "2.0")
	u.Add("et", "")
	u.Add("ct", "keyword="+key)
	uri := u.Encode()
	txt := utils.Getpage("http://g.click.taobao.com/display?" + uri)
	txt = strings.ReplaceAll(txt, "jsonp_callback_03822715911819512(", "")
	txt = strings.ReplaceAll(txt, "})", "}")
	var jsonx JsonxCp
	if err := json.Unmarshal([]byte(txt), &jsonx); err != nil {
		return nil, err
	}
	if jsonx.Code != 200 || jsonx.Data.Items == nil {
		return nil, errors.New("读取出错：" + utils.ToStr(jsonx.Code))
	}
	var cpall []OutCp
	for _, v := range jsonx.Data.Items {
		cpall = append(cpall, OutCp{
			Id:        v.DsNid,
			Price:     utils.FloatRound(v.DsDiscountPrice/100, 2),
			Yprice:    utils.FloatRound(v.DsReservePrice/100, 2),
			Title:     v.DsTitle,
			Wangwang:  v.DsNick,
			Chengjiao: v.DsSell,
			Pic:       ToChangePic(v.DsImg.Src).(string),
			Uid:       v.DsUserId,
		})
	}
	// log.Fatal(json)
	return cpall, nil
}

type JsonDp struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"error_msg"`
	Data     interface{}
}

// GetDP TODO 目前不好使．
func GetDP(uid int) (interface{}, error) {
	return nil, nil
	//u := url.Values{}
	//u.Add("cb", "jsonp_callback_007466789392105988")
	//u.Add("pid", "mm_24447735_7348472_24414817")
	//u.Add("wt", "1")
	//u.Add("ti", "3")
	//u.Add("tl", "140x190")
	//u.Add("rd", "2")
	//u.Add("st", "2")
	//u.Add("rf", "http://tta.cn/sss-%E4%B8%89%E6%98%9F.html")
	//u.Add("pgid", "c90c4586a8016545d3485ef2df953388")
	//u.Add("v", "2.0")
	//u.Add("et", "02255812")
	//u.Add("ct", "sellerid="+utils.ToStr(uid))
	//uri := u.Encode()
	//txt := utils.Getpage("http://g.click.taobao.com/display?" + uri)
	//txt = strings.ReplaceAll(txt, "jsonp_callback_007466789392105988(", "")
	//txt = strings.ReplaceAll(txt, "})", "}")
	//var jsonx JsonDp
	//if err := json.Unmarshal([]byte(txt), &jsonx); err != nil {
	//	return nil, err
	//}
	//if jsonx.Code != 200 {
	//	return nil, errors.New(jsonx.ErrorMsg)
	//}
	//// if jsonx.Code != 200 || jsonx.Data.Items == nil {
	//// 	return nil, errors.New("读取出错：" + utils.ToStr(jsonx.Code))
	//// }
	//// var cpall []OutCp
	//// for _, v := range jsonx.Data.Items {
	//// 	cpall = append(cpall, OutCp{
	//// 		Id:        v.DsNid,
	//// 		Price:     v.DsDiscountPrice,
	//// 		Yprice:    v.DsReservePrice,
	//// 		Title:     v.DsTitle,
	//// 		Wangwang:  v.DsNick,
	//// 		Chengjiao: v.DsSell,
	//// 		Pic:       v.DsImg.Src,
	//// 		Uid:       v.DsUserId,
	//// 	})
	//// }
	//// log.Fatal(json)
	//return jsonx, nil
}

// 获取商品详情
func GetCP(id int64) (OutCp, error) {
	u := url.Values{}
	u.Add("cb", "jsonp_callback_03822715911819512")
	u.Add("pid", "mm_12462245_2219662_23014024")
	u.Add("wt", "0")
	u.Add("tl", "290x380")
	u.Add("rd", "1")
	u.Add("st", "2")
	u.Add("rf", "http://www.quchaoshi.com")
	u.Add("pgid", "6ed70a3508d2e0539c43ab4c3bcb9615")
	u.Add("v", "2.0")
	u.Add("et", "")
	u.Add("ct", "itemid="+utils.ToStr(id))
	uri := u.Encode()
	txt := utils.Getpage("http://g.click.taobao.com/display?" + uri)
	txt = strings.ReplaceAll(txt, "jsonp_callback_03822715911819512(", "")
	txt = strings.ReplaceAll(txt, "})", "}")
	var jsonx JsonxCp
	if err := json.Unmarshal([]byte(txt), &jsonx); err != nil {
		return OutCp{}, err
	}
	if (jsonx.Code != 201 && jsonx.Code != 200) || jsonx.Data.Items == nil || len(jsonx.Data.Items) == 0 {
		return OutCp{}, errors.New("读取出错：" + utils.ToStr(jsonx.Code))
	}
	cpall := OutCp{}
	for k, v := range jsonx.Data.Items {
		if k == 0 {
			cpall = OutCp{
				Id:        v.DsNid,
				Price:     utils.FloatRound(v.DsDiscountPrice/100, 2),
				Yprice:    utils.FloatRound(v.DsReservePrice/100, 2),
				Title:     v.DsTitle,
				Wangwang:  v.DsNick,
				Chengjiao: v.DsSell,
				Pic:       ToChangePic(v.DsImg.Src).(string),
				Uid:       v.DsUserId,
			}
		}
	}
	return cpall, nil
}
