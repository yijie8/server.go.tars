package router

import (
	"TBapi/app/api/tbapi"
	_ "TBapi/docs"
	"github.com/gin-gonic/gin"
)

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("")
	//g.GET("/tbimg/*", func(c *gin.Context) {
	//	url := c.Request.URL.Path
	//	url = strings.ReplaceAll(url, "/tbimg/", "")
	//	urlAr := strings.Split(url, "/")
	//	s1 := urlAr[0]
	//	s1 = gstr.ReplaceByArray(s1, []string{
	//		"img1", "img01",
	//		"img2", "img02",
	//		"img3", "img03",
	//		"img4", "img04",
	//		"img5", "img05",
	//		"gaitaobao1", "img01",
	//		"gaitaobao2", "img02",
	//		"gaitaobao3", "img03",
	//		"gaitaobao4", "img04",
	//		"gaitaobao5", "img05",
	//	})
	//	s2 := url[len(s1+"/"):]
	//	if c.DefaultQuery("p", "") != "" {
	//		c.Redirect(200, "http://"+s1+".alicdn.com/"+s2)
	//	} else {
	//		c.Redirect(200, "http://"+s1+".taobaocdn.com/"+s2)
	//	}
	//})

	// tb接口
	web := g.Group("/web")
	{
		web.GET("/Cplist", tbapi.Cplist)
		web.GET("/XgByIdKw", tbapi.XgByIdKw)
		web.GET("/XgById", tbapi.XgById)
		web.GET("/XgByKey", tbapi.XgByKey)
		web.GET("/Cp10", tbapi.Cp10)
		web.GET("/Dp", tbapi.Dp)
		web.GET("/Cp", tbapi.Cpweb)
	}
	sdk := g.Group("/sdk")
	{
		sdk.GET("/CpList", tbapi.CpList)
		sdk.GET("/DPXgByKey", tbapi.DPXgByKey)
		sdk.GET("/DPXgById", tbapi.DPXgById)
		sdk.GET("/CpXgList", tbapi.CpXgList)
		sdk.GET("/Cp", tbapi.Cp)
		sdk.GET("/Tgg", tbapi.Tgg)
		sdk.GET("/Tkl", tbapi.Tkl)
		sdk.GET("/Content", tbapi.Content)
		sdk.GET("/CpGood", tbapi.CpGood)
		sdk.GET("/CouponDesc", tbapi.CouponDesc)
		sdk.GET("/CreateTlj", tbapi.CreateTlj)
		sdk.GET("/TljDesc", tbapi.TljDesc)
	}
	return g
}
