package models

import (
	"ZserverWeb/Zserver"
	"ZserverWeb/client"
	"encoding/json"
	_ "time"
)

type SysConfig struct {
	ConfigId    int    `json:"configId" gorm:"primary_key;auto_increment;"` //编码
	ConfigName  string `json:"configName" gorm:"type:varchar(128);"`        //参数名称
	ConfigKey   string `json:"configKey" gorm:"type:varchar(128);"`         //参数键名
	ConfigValue string `json:"configValue" gorm:"type:varchar(255);"`       //参数键值
	ConfigType  string `json:"configType" gorm:"type:varchar(64);"`         //是否系统内置
	Remark      string `json:"remark" gorm:"type:varchar(128);"`            //备注
	CreateBy    string `json:"createBy" gorm:"type:varchar(128);"`
	UpdateBy    string `json:"updateBy" gorm:"type:varchar(128);"`
	DataScope   string `json:"dataScope" gorm:"-"`
	Params      string `json:"params"  gorm:"-"`
	BaseModel
}

var req_sysConfig Zserver.SysConfig
var res_sysConfig Zserver.SysConfig
var ress_sysConfig Zserver.SysConfig_List

func (SysConfig) TableName() string {
	return "sys_config"
}

// Config 创建
func (e *SysConfig) Create() (doc SysConfig, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_sysConfig)
	if err != nil {
		return doc, err
	}
	err = client.WebApiAuth().SysConfig_Create(&req_sysConfig, &res_sysConfig)
	if err != nil {
		return doc, err
	}
	err = json.Unmarshal(client.Struct2Json(res_sysConfig), &doc)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

// 获取 Config
func (e *SysConfig) Get() (doc SysConfig, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_sysConfig)
	if err != nil {
		return doc, err
	}
	err = client.WebApiAuth().SysConfig_Get(&req_sysConfig, &res_sysConfig)
	if err != nil {
		return doc, err
	}
	err = json.Unmarshal(client.Struct2Json(res_sysConfig), &doc)
	if err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *SysConfig) GetPage(pageSize int, pageIndex int) (doc []SysConfig, count int, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &ress_sysConfig)
	if err != nil {
		return doc, count, err
	}
	err = client.WebApiAuth().SysConfig_GetPage(int32(pageSize), int32(pageIndex), &req_sysConfig, &ress_sysConfig)
	if err != nil {
		return doc, count, err
	}
	err = json.Unmarshal(client.Struct2Json(ress_sysConfig.SysConfigList), &doc)
	if err != nil {
		return doc, count, err
	}
	return doc, int(ress_sysConfig.Count), nil
}

func (e *SysConfig) Update(id int) (update SysConfig, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_sysConfig)
	if err != nil {
		return update, err
	}
	err = client.WebApiAuth().SysConfig_Update(int32(id), &req_sysConfig, &res_sysConfig)
	if err != nil {
		return update, err
	}
	err = json.Unmarshal(client.Struct2Json(res_sysConfig), &update)
	if err != nil {
		return update, err
	}
	return
}

func (e *SysConfig) Delete() (success bool, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_loginlog)
	if err != nil {
		return false, err
	}
	err = client.WebApiAuth().SysConfig_Delete(&req_sysConfig, &success)
	if err != nil {
		return false, err
	}
	return
}

func (e *SysConfig) BatchDelete(id []int) (Result bool, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_sysConfig)
	if err != nil {
		return false, err
	}
	err = client.WebApiAuth().SysConfig_BatchDelete(client.Ar2Int32(id), &req_sysConfig, &Result)
	if err != nil {
		return false, err
	}
	return
}
