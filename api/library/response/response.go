package response

import (
	"github.com/gin-gonic/gin"
	"go/types"
)

// 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// err:  错误码(0:成功, 1:失败, >1:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func Json(c *gin.Context, code int, msg interface{}, data ...interface{}) {
	msgString := ""
	switch msg.(type) {
	case types.Nil:
		msgString = ""
	case string:
		msgString = msg.(string)
	case error:
		if msg.(error) != nil {
			msgString = msg.(error).Error()
		} else {
			msgString = ""
		}
	}

	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	//c.Header("Content-Type", "application/json")

	c.JSON(code, gin.H{
		"code": code,
		"msg":  msgString,
		"data": responseData,
	})
}

// 903参数错误　904验证错误　905返回数据出错
