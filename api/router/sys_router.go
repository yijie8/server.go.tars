package router

import (
	"TBapi/app/api/tbapi"
	_ "TBapi/docs"
	"github.com/gin-gonic/gin"
)

func InitSysRouter(r *gin.Engine) *gin.RouterGroup {
	g := r.Group("/tb_web_go")
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
		web.GET("/cplist", tbapi.Cplist)
		web.GET("/xgbyidkw", tbapi.XgByIdKw)
		web.GET("/xgbyid", tbapi.XgById)
		web.GET("/xgbykey", tbapi.XgByKey)
		web.GET("/cp10", tbapi.Cp10)
		web.GET("/dp", tbapi.Dp)
		web.GET("/cp", tbapi.Cpweb)
		web.GET("/gets56qq", tbapi.Gets56qq)
		web.GET("/gets56q", tbapi.Gets56q)
	}
	sdk := g.Group("/sdk")
	{
		sdk.GET("/cplist", tbapi.CpList)
		sdk.GET("/dpxgbykey", tbapi.DPXgByKey)
		sdk.GET("/dpxgbyid", tbapi.DPXgById)
		sdk.GET("/cpxglist", tbapi.CpXgList)
		sdk.GET("/cp", tbapi.Cp)
		sdk.GET("/tgg", tbapi.Tgg)
		sdk.GET("/tkl", tbapi.Tkl)
		sdk.GET("/content", tbapi.Content)
		sdk.GET("/cpgood", tbapi.CpGood)
		sdk.GET("/coupondesc", tbapi.CouponDesc)
		sdk.GET("/createtlj", tbapi.CreateTlj)
		sdk.GET("/tljdesc", tbapi.TljDesc)
	}
	return g
}
