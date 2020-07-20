package models

import (
	"ZserverWeb/Zserver"
	"ZserverWeb/client"
	"encoding/json"
	"time"
)

type LoginLog struct {
	InfoId        int       `json:"infoId" gorm:"primary_key;AUTO_INCREMENT"` //主键
	Username      string    `json:"username" gorm:"type:varchar(128);"`       //用户名
	Status        string    `json:"status" gorm:"type:int(1);"`               //状态
	Ipaddr        string    `json:"ipaddr" gorm:"type:varchar(255);"`         //ip地址
	LoginLocation string    `json:"loginLocation" gorm:"type:varchar(255);"`  //归属地
	Browser       string    `json:"browser" gorm:"type:varchar(255);"`        //浏览器
	Os            string    `json:"os" gorm:"type:varchar(255);"`             //系统
	Platform      string    `json:"platform" gorm:"type:varchar(255);"`       // 固件
	LoginTime     time.Time `json:"loginTime" gorm:"type:timestamp;"`         //登录时间
	CreateBy      string    `json:"createBy" gorm:"type:varchar(128);"`       //创建人
	UpdateBy      string    `json:"updateBy" gorm:"type:varchar(128);"`       //更新者
	DataScope     string    `json:"dataScope" gorm:"-"`                       //数据
	Params        string    `json:"params" gorm:"-"`                          //
	Remark        string    `json:"remark" gorm:"type:varchar(255);"`         //备注
	Msg           string    `json:"msg" gorm:"type:varchar(255);"`
	BaseModel
}

var req_loginlog Zserver.LoginLog
var res_loginlog Zserver.LoginLog
var ress_loginLog_List Zserver.LoginLog_List
var doc_loginLog LoginLog
var docs_loginLog []LoginLog

func (LoginLog) TableName() string {
	return "sys_loginlog"
}

func (e *LoginLog) Get() (LoginLog, error) {
	err := json.Unmarshal(client.Struct2Json(e), &req_loginlog)
	if err != nil {
		return doc_loginLog, err
	}

	err = client.WebApiAuth().LoginLog_Get(&req_loginlog, &res_loginlog)
	if err != nil {
		return doc_loginLog, err
	}
	err = json.Unmarshal(client.Struct2Json(res_loginlog), &doc_loginLog)
	if err != nil {
		return doc_loginLog, err
	}

	return doc_loginLog, nil
}

func (e *LoginLog) GetPage(pageSize int, pageIndex int) ([]LoginLog, int, error) {
	var count int
	err := json.Unmarshal(client.Struct2Json(e), &req_loginlog)
	if err != nil {
		return docs_loginLog, count, err
	}
	err = client.WebApiAuth().LoginLog_GetPage(int32(pageSize), int32(pageIndex), &req_loginlog, &ress_loginLog_List)
	if err != nil {
		return docs_loginLog, count, err
	}
	err = json.Unmarshal(client.Struct2Json(ress_loginLog_List.LoginLogList), &docs_loginLog)
	if err != nil {
		return docs_loginLog, 0, err
	}
	return docs_loginLog, int(ress_loginLog_List.Count), nil
}

func (e *LoginLog) Create() (LoginLog, error) {
	var doc LoginLog
	e.CreateBy = "0"
	e.UpdateBy = "0"

	err := json.Unmarshal(client.Struct2Json(e), &req_loginlog)
	if err != nil {
		return doc, err
	}

	err = client.WebApiAuth().LoginLog_Create(&req_loginlog, &res_loginlog)
	if err != nil {
		return doc, err
	}
	err = json.Unmarshal(client.Struct2Json(res_loginlog), &doc)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *LoginLog) Update(id int) (update LoginLog, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_loginlog)
	if err != nil {
		return update, err
	}
	err = client.WebApiAuth().LoginLog_Update(int32(id), &req_loginlog, &res_loginlog)
	if err != nil {
		return update, err
	}
	err = json.Unmarshal(client.Struct2Json(res_loginlog), &update)
	if err != nil {
		return update, err
	}
	return
}

func (e *LoginLog) BatchDelete(id []int) (Result bool, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_loginlog)
	if err != nil {
		return false, err
	}
	err = client.WebApiAuth().LoginLog_BatchDelete(client.Ar2Int32(id), &req_loginlog, &Result)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(client.Struct2Json(res_loginlog), &doc_loginLog)
	if err != nil {
		return false, err
	}
	Result = true
	return
}
