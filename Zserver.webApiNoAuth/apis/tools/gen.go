package tools

import (
	"bytes"
	"github.com/gin-gonic/gin"
	tools2 "github.com/yijie8/zserver/tools"
	"github.com/yijie8/zserver/tools/app"
	"net/http"
	"ZserverWebApiNoAuth/Zserver"
	"ZserverWebApiNoAuth/models/tools"
	"text/template"
)

func Preview(c *gin.Context) {
	id, err := tools2.StringToInt(c.Param("tableId"))
	tools2.HasError(err, "", -1)
	mp, err := Preview_(id)
	tools2.HasError(err, "", -1)
	var res app.Response
	res.Data = mp

	c.JSON(http.StatusOK, res.ReturnOK())
}
func Preview_tars(id int, res *Zserver.Preview_res) error {
	mp, err := Preview_(id)
	res.Code = 200
	res.Msg = "success"
	res.Data = mp
	return err
}
func Preview_(id int) (map[string]string, error) {
	table := tools.SysTables{}
	table.TableId = id
	t1, err := template.ParseFiles("template/model.go.template")
	if err != nil {
		return nil, err
	}
	t2, err := template.ParseFiles("template/api.go.template")
	if err != nil {
		return nil, err
	}
	t3, err := template.ParseFiles("template/js.go.template")
	if err != nil {
		return nil, err
	}
	t4, err := template.ParseFiles("template/vue.go.template")
	if err != nil {
		return nil, err
	}
	t5, err := template.ParseFiles("template/router.go.template")
	if err != nil {
		return nil, err
	}
	tab, _ := table.Get()
	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)
	var b5 bytes.Buffer
	err = t5.Execute(&b5, tab)
	if err != nil {
		return nil, err
	}

	mp := make(map[string]string)
	mp["template/model.go.template"] = b1.String()
	mp["template/api.go.template"] = b2.String()
	mp["template/js.go.template"] = b3.String()
	mp["template/vue.go.template"] = b4.String()
	mp["template/router.go.template"] = b5.String()
	return mp, nil
}
