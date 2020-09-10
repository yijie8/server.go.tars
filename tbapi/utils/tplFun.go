package utils

import (
	"github.com/gogf/gf/text/gstr"
	"strings"
)

//模板函数

func GetPic(url string) string {

	// 已转换完的不转了
	if !strings.HasPrefix(strings.ToLower(url), "http://") && !strings.HasPrefix(strings.ToLower(url), "https://") {
		return url
	}

	url = gstr.Replace(url, "https://", "")
	url = gstr.Replace(url, "http://", "")
	url = gstr.Replace(url, "image.taobao.com", "img.taobaocdn.com")
	if gstr.Contains(url, ".taobaocdn.com/") {
		ua := gstr.Split(url, ".taobaocdn.com/")
		if len(ua) < 2 {
			return "http://" + url
		}
		return "/tbimg/" + ua[0] + "/" + ua[1]
	} else if gstr.Contains(url, ".alicdn.com/") {
		ua := gstr.Split(url, ".alicdn.com/")
		if len(ua) < 2 {
			return "http://" + url
		}
		ua[0] = gstr.Replace(ua[0], "gaitaobao", "img0")
		return "/tbimg/" + ua[0] + "/" + ua[1]
	} else if gstr.Contains(url, ".paipaiimg.com/") {
		ua := gstr.Split(url, ".paipaiimg.com/")
		if len(ua) < 2 {
			return "http://" + url
		}
		return "/ppimg/" + ua[0] + "/" + ua[1]
	} else if gstr.Contains(url, ".tbcdn.cn/") {
		ua := gstr.Split(url, ".tbcdn.cn/")
		if len(ua) < 2 {
			return "http://" + url
		}
		return "/tbimg/" + ua[0] + "/" + ua[1]
	} else {
		return "http://" + url
	}
}
