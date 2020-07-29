package tbsdk

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"

	"github.com/gogf/gf/os/gtime"
	"math/rand"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/nilorg/go-opentaobao"
)

type Tbsdk struct {
	Api     string                 // 接口名
	AppKey  string                 // 接口的appkey
	Args    map[string]interface{} // 参数
	Zid     string                 // 指定一个站点的key
	KeyWord string                 // 搜索关键词
	Result  *simplejson.Json
}

func NewTbsdk() *Tbsdk {
	return &Tbsdk{
		Api:    "",
		AppKey: "",
		Args:   nil,
		Zid:    "",
		Result: nil,
	}
}
func (t *Tbsdk) Init_(args ...string) string {
	// log.Warning()ing("tbsdk init")
	appkey := "23124448 6ae95ec943bd2668888da6f57fb99d8e 249286346 hahale"
	appkey += "|23124447 e76c9fcba4dbb745ebd40cde367663d8 24228091 pp130" //
	appkey += "|23131817 16d9080907f0d4e1085826fead471ab2 24364644 s56"
	appkey += "|23131818 5bad4c09ded622055c577366a27a7bd2 26980824 pp30"
	appkey += "|23131816 26d745e3e44ab3d07438fa04097fbea2 26802604 tta"
	appkey += "|23652288 7343f6712e04c300e367b16855aabf84 26826725 quchaoshi" //
	appkey += "|23652082 510ca6d15000a6f7b37a0033148a2249 109687600500 nb"
	appkey += "|23650905 1a655755ce79bf8612ce267241d3c524 109688450493 ufo" //
	appkey += "|25113003 49c06a762571822cc0c62f93da50e203 109621650188 62a"
	appkey += "|24842016 97b708543242d7fd8ae0371efa37e6a5 409028055 miangoujuan"
	appkey += "|24842017 e42742ebab07bbe81d23bd3b2c7e3129 403248213 mianfeijuan"

	keyar := strings.Split(appkey, "|")
	// 一条数据
	// keyone := ""

	// 指定一个key
	if len(args) > 0 && args[0] != "" {
		t.Zid = args[0]
		for _, value := range keyar {
			ks := strings.Split(value, " ")
			if ks[3] == t.Zid {
				t.AppKey = value
				break
			}
		}
	} else {
		// 随机一个key
		rand.Seed(time.Now().UnixNano())
		t.AppKey = keyar[rand.Intn(len(keyar)-1)]
	}

	keys := strings.Split(t.AppKey, " ")
	opentaobao.AppKey = keys[0]
	opentaobao.AppSecret = keys[1]
	opentaobao.Timeout = 3 * time.Second
	opentaobao.CacheExpiration = time.Hour * 24
	opentaobao.Router = "http://gw.api.taobao.com/router/rest"
	log.Debug(t.AppKey)
	return t.AppKey
}

// 取接口的底层类
func (t *Tbsdk) GetTB(args ...string) *simplejson.Json {
	t.Init_(args...)
	if t.Api == "taobao.tbk.dg.vegas.tlj.create" || t.Api == "taobao.tbk.dg.optimus.material" || t.Api == "taobao.tbk.dg.material.optional" || t.Api == "taobao.tbk.ju.tqg.get" || t.Api == "taobao.tbk.content.get" {
		t.Args["adzone_id"] = strings.Split(t.AppKey, " ")[2]
	}
	delete(t.Args, "zid")
	res, err := opentaobao.Execute(t.Api, t.Args)
	if err != nil {
		log.Debug(err)
	}
	//log.Debug(res.MarshalJSON())
	if res == nil {
		return &simplejson.Json{}
	}
	return res
}

// 取相关店铺 搜索
func (t *Tbsdk) GetDPXgListKey(key string, page ...int) *Tbsdk {
	t = NewTbsdk()
	p := 1
	if len(page) > 0 && page[0] != 0 {
		p = page[0]
	}
	t.KeyWord = key
	t.Api = "taobao.tbk.shop.get"
	t.Args = map[string]interface{}{
		"fields":    "user_id,shop_title,seller_nick,pict_url", // ,shop_type,shop_url
		"q":         key,
		"sort":      "total_auction_des",
		"page_no":   p,
		"page_size": 100,
	}
	t.Result = t.GetTB().Get("tbk_shop_get_response")
	if t.Result == nil {
		log.Warning("没有获取到数据")
		return nil
	}
	return t
}

// 取相关店铺 ＩＤ
func (t *Tbsdk) GetDPXgListId(dpid int) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.shop.recommend.get"
	t.Args = map[string]interface{}{
		"fields":  "user_id,shop_title,seller_nick,pict_url",
		"user_id": dpid,
		"count":   20,
	}
	t.Result = t.GetTB().Get("tbk_shop_recommend_get_response")
	if t.Result == nil {
		log.Warning("没有获取到数据")
		return nil
	}
	return t
}

// 取商品列表
func (t *Tbsdk) GetCpList(args map[string]interface{}, appkey ...string) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.dg.material.optional"
	// 初始默认值
	if _, ok := args["page_size"]; !ok || (ok && args["page_size"] == 0) {
		args["page_size"] = 100
	}
	if _, ok := args["page_no"]; !ok || (ok && args["page_no"] == 0) {
		args["page_no"] = 1
	}
	if args["page_no"] == 0 || args["page_no"] == "0" || args["page_no"] == "" {
		args["page_no"] = 1
	}
	if _, ok := args["platform"]; !ok || (ok && args["platform"] == 0) {
		args["platform"] = 1
	}
	// if _, ok := args["end_tk_rate"]; !ok {
	// 	args["end_tk_rate"] = 10000
	// }
	if _, ok := args["start_tk_rate"]; !ok {
		//args["start_tk_rate"] = 500
	}
	if _, ok := args["sort"]; !ok || (ok && args["sort"] == "") {
		args["sort"] = "total_sales_des"
	}
	if _, ok := args["material_id"]; !ok || (ok && args["material_id"] == 0) {
		args["material_id"] = 17004 // 默认物料id=2836 官方个性化算法优化的搜索物料id=17004
	}

	// 商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
	if _, ok := args["start_ka_tk_rate"]; !ok {
		// args["start_ka_tk_rate"] = 100
	}
	if _, ok := args["q"]; !ok {

	} else {
		t.KeyWord = args["q"].(string)
	}

	if _, ok := args["cat"]; ok {
		if args["cat"] == "0" {
			args["cat"] = ""
		}
	}

	t.Args = args
	t.Result = t.GetTB(appkey...)
	if t.Result == nil {
		log.Debug("没有获取到数据")
		return nil
	}
	t.Result = t.Result.Get("tbk_dg_material_optional_response") // .Get("result_list").Get("map_data")
	return t
}

// 取相关商品
func (t *Tbsdk) GetCpXgList(cpid int64) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.item.recommend.get"
	t.Args = map[string]interface{}{
		"fields":  "num_iid,title,pict_url,reserve_price,zk_final_price,nick,seller_id,volume",
		"num_iid": cpid,
		"count":   20,
	}
	t.Result = t.GetTB().Get("tbk_item_recommend_get_response").Get("results").Get("n_tbk_item")
	if t.Result == nil {
		log.Warning("没有获取到数据")
		return nil
	}
	return t
}

// 取单个商品
func (t *Tbsdk) GetCp(cpids string) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.item.info.get"
	t.Args = map[string]interface{}{
		"num_iids": cpids,
	}
	// TODO 没有数据
	t.Result = t.GetTB().Get("tbk_item_info_get_response").Get("results").Get("n_tbk_item")
	if t.Result == nil {
		log.Warning("没有获取到数据")
		return nil
	}
	return t
}

// 淘抢购
func (t *Tbsdk) GetTGG(args map[string]interface{}, appkey ...string) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.ju.tqg.get"
	// total_amount 总库存
	// sold_num 已抢购数量
	args["fields"] = "click_url,pic_url,reserve_price,zk_final_price,total_amount,sold_num,title,category_name,start_time,end_time"
	if _, ok := args["start_time"]; !ok || (ok && args["start_time"] == "") {
		args["start_time"] = gtime.Now().Format("Y-m-d H:i:s T")
	}
	if _, ok := args["end_time"]; !ok || (ok && args["end_time"] == "") {
		args["end_time"] = gtime.Now().AddDate(0, 0, +1).Format("Y-m-d H:i:s")
	}
	if _, ok := args["page_no"]; !ok || (ok && args["page_no"] == 0) {
		args["page_no"] = 1
	}
	if _, ok := args["page_size"]; !ok || (ok && args["page_size"] == 0) {
		args["page_size"] = 40
	}
	t.Args = args
	t.Result = t.GetTB(appkey...).Get("tbk_ju_tqg_get_response").Get("results").Get("results")
	if t.Result == nil {
		log.Warning("没有获取到数据")
		return nil
	}
	return t
}

// 获取淘口令
// test 口令弹框内容
// url 口令跳转目标页
func (t *Tbsdk) GetTKL(text string, url string, appkey ...string) (string, error) {
	t = NewTbsdk()
	t.Api = "taobao.tbk.tpwd.create"
	t.Args = map[string]interface{}{
		"text": text,
		"url":  url,
	}
	res, err := t.GetTB(appkey...).Get("tbk_tpwd_create_response").Get("data").Get("model").String()
	if err != nil {
		return "", err
	}
	return res, nil
}

// 获取图文内容输出
func (t *Tbsdk) GetContent(args map[string]interface{}, appkey ...string) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.content.get"
	if _, ok := args["type"]; !ok || (ok && args["type"] == 0) {
		args["type"] = 1 // 内容类型，1:图文、2: 图集、3: 短视频
	}
	if _, ok := args["before_timestamp"]; !ok || (ok && args["before_timestamp"] == 0) {
		args["before_timestamp"] = time.Now().UnixNano() // 表示取这个时间点以前的数据，以毫秒为单位（出参中的last_timestamp是指本次返回的内容中最早一条的数据，下一次作为before_timestamp传过来，即可实现翻页）
	}
	if _, ok := args["count"]; !ok || (ok && args["count"] == 0) {
		args["count"] = 10 // 表示期望获取条数
	}
	if _, ok := args["cid"]; !ok || (ok && args["cid"] == "") { // 类目

	}

	if _, ok := args["image_width"]; !ok || (ok && args["image_width"] == 0) {
		args["image_width"] = 300 // 图片宽度
	}
	if _, ok := args["image_height"]; !ok || (ok && args["image_height"] == 0) {
		args["image_height"] = 300 // 图片高度
	}
	t.Args = args

	t.Result = t.GetTB(appkey...).Get("tbk_content_get_response").Get("result").Get("data")
	if t.Result == nil {
		log.Warning(t.Api + "没有获取到数据")
		return nil
	}
	return t
}

// 获取优质商品 物料精选
func (t *Tbsdk) GetCpListGood(args map[string]interface{}, appkey ...string) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.dg.optimus.material"
	if _, ok := args["material_id"]; !ok || (ok && args["material_id"] == 0) {
		args["material_id"] = 6708
		// 6708含全部商品	28017营销商品库商品
		// 好券直播(综合3756 鞋包配饰3762 母婴3760 女装3767 美妆个护3763 食品3761 家居家装3758 男装3764 运动户外3766 数码家电3759 内衣3765)
		// 大额券(综合9660 // 鞋包配饰	9648// 母婴	 9650// 女装	9658// 美妆个护	9653// 食品	 9649// 家居家装	9655// 男装	9654// 运动户外	 9651// 数码家电	9656// 内衣	9652)
		// 高佣榜(综合	13366// 鞋包配饰	13370// 母婴	13374// 女装	13367// 美妆个护	13371// 食品	13375// 家居家装	13368// 男装	13372// 运动户外	13376// 数码家电	13369// 内衣	13373)
		// 品牌券(综合3786// 鞋包配饰	 3796// 母婴	 3789// 女装3788// 美妆个护	 3794// 食品	 3791// 家居家装3792// 男装	 3790// 运动户外	 3795// 数码家电3793// 内衣	 3787)
		// 母婴主题(母婴_备孕 4040// 母婴_0至6个月 4041// 母婴_4至6岁 4044// 母婴_7至12个月 4042// 母婴_1至3岁 4043// 母婴_7至12岁 4045)
		// 有好货(4092)
		// 潮流范(4093)
		// 特惠(4094)
	}
	if _, ok := args["content_id"]; !ok || (ok && args["content_id"] == 0) {
		args["content_id"] = 0 // 内容专用-内容详情ID
	}
	if _, ok := args["content_source"]; !ok || (ok && args["content_source"] == "") {
		args["content_source"] = "" // 内容专用-内容渠道信息
	}
	if _, ok := args["item_id"]; !ok || (ok && args["item_id"] == "") { // 商品ID，用于相似商品推荐

	}

	if _, ok := args["page_size"]; !ok || (ok && args["page_size"] == 0) {
		args["page_size"] = 40 // 页大小
	}
	if _, ok := args["page_no"]; !ok || (ok && args["page_no"] == 0) {
		args["page_no"] = 1 // 第几页
	}
	t.Args = args
	t.Result = t.GetTB(appkey...).Get("tbk_dg_optimus_material_response").Get("result_list").Get("map_data")
	if t.Result == nil {
		log.Warning(t.Api + "没有获取到数据")
		return nil
	}
	return t
}

// 没有权限 链接解析出商品id　从长链接或短链接中解析出open_iid
// func GetCpTuiID(click_url string) (open_iid ,item_id string , err error) {
// 	if click_url=="" {
// 		return "","", errors.New("链接必须有")
// 	}
// 	res, err := GetTB("taobao.tbk.item.click.extract", map[string]interface{}{"click_url":click_url})
// 	if err != nil {
// 		return "","", err
// 	}
// 	if open_iid, err = res.Get("tbk_item_click_extract_response").Get("open_iid").String(); err != nil {
// 		return "","", errors.New("没有数据：" + err.Error())
// 	}
// 	if item_id, err = res.Get("tbk_item_click_extract_response").Get("item_id").String(); err != nil {
// 		return "","", errors.New("没有数据：" + err.Error())
// 	}
//
// 	return
// }

// 获取coupon详细
func (t *Tbsdk) GetCpQDesc(args map[string]interface{}) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.coupon.get"
	if _, ok := args["activity_id"]; !ok || (ok && args["activity_id"] == "") {
		log.Warning("activity_id error")

	} else if _, ok := args["item_id"]; !ok || (ok && args["item_id"] == "") {
		log.Warning("item_id error")
		return t
	}
	t.Args = args
	t.Result = t.GetTB().Get("tbk_coupon_get_response").Get("data")
	if t.Result == nil {
		log.Warning(t.Api + "没有获取到数据")
		return nil
	}
	return t
}

// 创建淘礼金
func (t *Tbsdk) GetCreateTLJ(args map[string]interface{}, appkey ...string) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.dg.vegas.tlj.create"
	if _, ok := args["total_num"]; !ok || (ok && args["total_num"] == 0) {
		args["total_num"] = 1
	}
	if _, ok := args["item_id"]; !ok || (ok && args["item_id"] == "") {
		log.Warning("item_id error")
		return t
	}

	if _, ok := args["name"]; !ok || (ok && args["name"] == "") { // 淘礼金名称
		args["name"] = "淘礼金来啦"
	}
	if _, ok := args["user_total_win_num_limit"]; !ok || (ok && args["user_total_win_num_limit"] == 0) { // 单用户累计中奖次数上限
		args["user_total_win_num_limit"] = 1
	}
	if _, ok := args["security_switch"]; !ok { // 安全开关
		args["security_switch"] = true
	}
	if _, ok := args["per_face"]; !ok || (ok && args["per_face"] == 0) { // 单个淘礼金面额，支持两位小数，单位元
		args["per_face"] = 1
	}
	if _, ok := args["send_start_time"]; !ok || (ok && args["send_start_time"] == "") { // 发放开始时间
		args["send_start_time"] = gtime.Now().Format("Y-m-d H:i:s")
	}
	if _, ok := args["send_end_time"]; !ok || (ok && args["send_end_time"] == "") { // 发放截止时间
		args["send_end_time"] = gtime.Now().AddDate(0, 0, +7).Format("Y-m-d H:i:s")
	}
	t.Args = args

	t.Result = t.GetTB(appkey...).Get("tbk_dg_vegas_tlj_create_response").Get("result")
	if t.Result == nil {
		log.Warning(t.Api + "没有获取到数据")
		return nil
	}
	return t
}

// 淘礼金使用情况
// 实例ID
func (t *Tbsdk) GetTLJDesc(rightsId string, appkey ...string) *Tbsdk {
	t = NewTbsdk()
	t.Api = "taobao.tbk.dg.vegas.tlj.instance.report"
	t.Args = map[string]interface{}{
		"rightsId": rightsId,
	}
	if rightsId == "" {
		log.Warning("实例ID 不能空")
		return t
	}

	t.Result = t.GetTB(appkey...).Get("result").Get("model")
	if t.Result == nil {
		log.Warning(t.Api + "没有获取到数据")
		return nil
	}
	return t
}

func (t *Tbsdk) ToJson() (string, error) {
	if r, err := t.Result.MarshalJSON(); err != nil {
		return "", errors.New("没有数据：" + err.Error())
	} else {
		return string(r), nil
	}
}

func (t *Tbsdk) ToStruct() (interface{}, error) {
	if t == nil {
		return nil, errors.New("没有数据")
	}
	if r, err := t.Result.MarshalJSON(); err != nil {
		return nil, errors.New("没有数据：" + err.Error())
	} else {
		if t.Api == "taobao.tbk.dg.material.optional" {
			var resType GetCpListJson
			if e := json.Unmarshal(r, &resType); e != nil {
				return nil, e
			} else {
				resType.SearchKey = t.KeyWord
				return resType, nil
			}
		} else if t.Api == "taobao.tbk.shop.get" {
			var resType GetDPXgListKeyJson
			if e := json.Unmarshal(r, &resType); e != nil {
				return nil, e
			} else {
				resType.SearchKey = t.KeyWord
				return resType, nil
			}
		} else if t.Api == "taobao.tbk.content.get" {
			var resType GetContentJson
			if e := json.Unmarshal(r, &resType); e != nil {
				return nil, e
			} else {
				return resType, nil
			}
		} else if t.Api == "taobao.tbk.item.info.get" {
			var resType []CpOnly
			if e := json.Unmarshal(r, &resType); e != nil {
				return nil, e
			} else {
				return resType, nil
			}
		} else if t.Api == "taobao.tbk.item.recommend.get" {
			var resType []CpOnly
			if e := json.Unmarshal(r, &resType); e != nil {
				return nil, e
			} else {
				return resType, nil
			}
		} else {
			return string(r), nil
		}
	}
}
