package system

import (
	"github.com/gin-gonic/gin"
	"github.com/yijie8/zserver/tools"
	"net/http"
	"ZserverWeb/Zserver"
	"ZserverWeb/client"
)

func GenerateCaptchaHandler(c *gin.Context) {
	var res Zserver.GetCaptcha_res
	err := client.WebApiNoAuth().GetCaptcha(&res)
	tools.HasError(err, "验证码获取失败", 500)
	c.JSON(http.StatusOK, res)
}
