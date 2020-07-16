package tools

import (
	"github.com/gin-gonic/gin"
	tools2 "github.com/yijie8/zserver/tools"
	"net/http"
	"ZserverWeb/Zserver"
	"ZserverWeb/client"
)

func Preview(c *gin.Context) {
	id, err := tools2.StringToInt(c.Param("tableId"))
	tools2.HasError(err, "", -1)
	var res Zserver.Preview_res
	err = client.WebApiNoAuth().Preview(int32(id), &res)
	tools2.HasError(err, "", -1)
	c.JSON(http.StatusOK, res)
}
