package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/yijie8/zserver/tools"
	"net/http"
	"ZserverWeb/Zserver"
	"ZserverWeb/client"
)

func ServerInfo(c *gin.Context) {
	var res Zserver.Monitor_server_res
	err := client.WebApiNoAuth().Monitor_server(&res)
	tools.HasError(err, "", 500)
	c.JSON(http.StatusOK, res)
}
