package system

import (
	"github.com/gin-gonic/gin"
	"github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/app"
	"github.com/yijie8/zserver/tools/captcha"
)

func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := GetCaptcha_()
	tools.HasError(err, "验证码获取失败", 500)
	app.Custum(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}

func GetCaptcha_() (id string, b64s string, err error) {
	return captcha.DriverDigitFunc()
}
