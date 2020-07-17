package client

import (
	"ZserverWeb/Zserver"
	"encoding/json"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/yijie8/zserver/tools"
)

var Conn *tars.Communicator
var Obj string

func init() {
	Conn = tars.NewCommunicator()
}

func WebApiNoAuth() *Zserver.WebApiNoAuth {
	Obj = tools.Conn_server("Zserver.WebApiNoAuth.TcpObj", "", Conn)
	App := new(Zserver.WebApiNoAuth)
	Conn.StringToProxy(Obj, App)
	return App
}

func WebApiAuth() *Zserver.WebApiAuth {
	Obj = tools.Conn_server("Zserver.WebApiAuth.TcpObj", "10015", Conn)
	//Obj = fmt.Sprintf("%v@tcp -h 127.0.0.1 -p %v -t 60000", Obj, "10015")
	App := new(Zserver.WebApiAuth)
	Conn.StringToProxy(Obj, App)
	return App
}
func Struct2Struct(form interface{}, to interface{}) interface{} {
	// TODO interface conversion: interface {} is map[string]interface {}, not Zserver.LoginLog
	jsonx, err := json.Marshal(form)
	if err != nil {
		return to
	}
	err = json.Unmarshal(jsonx, &to)
	if err != nil {
		return to
	}
	return to
}

func Struct2Json(form interface{}) []byte {
	jsonx, err := json.Marshal(form)
	if err != nil {
		jsonx, err = json.Marshal(map[string]string{"err": err.Error()})
		if err != nil {
			return []byte(err.Error())
		}
	}
	return jsonx
}

func Ar2Int(from interface{}) (to []int) {
	to = make([]int, 0)
	for k, v := range from.([]interface{}) {
		switch v.(type) {
		case int8:
			to[k] = int(v.(int8))
		case int16:
			to[k] = int(v.(int16))
		case int32:
			to[k] = int(v.(int32))
		case int64:
			to[k] = int(v.(int64))
		}
	}
	return to
}

func Ar2Int32(from interface{}) (to []int32) {
	to = make([]int32, 0)
	for k, v := range from.([]interface{}) {
		switch v.(type) {
		case int8:
			to[k] = int32(v.(int8))
		case int16:
			to[k] = int32(v.(int16))
		case int:
			to[k] = int32(v.(int))
		case int64:
			to[k] = int32(v.(int64))
		}
	}
	return to
}
