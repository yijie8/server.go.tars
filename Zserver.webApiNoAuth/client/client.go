package main

import (
	"encoding/json"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/yijie8/zserver/tools"
	"ZserverWebApiNoAuth/Zserver"
)

func main() {
	comm := tars.NewCommunicator()
	obj := tools.Conn_server("Zserver.Web.TcpObj", "10015", comm)

	app := new(Zserver.Web)
	comm.StringToProxy(obj, app)

	//var out, i int32
	//i = 123
	//ret, err := app.Add(i, i*2, &out)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//var out Zserver.Monitor_server_res
	//err := app.Monitor_server(&out)

	var out Zserver.GetMenuTreeelect_res
	err := app.MenuTreeselect(&out)
	jsonx, _ := json.Marshal(out)
	fmt.Println(err, string(jsonx))
}
