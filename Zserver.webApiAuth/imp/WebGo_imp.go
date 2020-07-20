package imp

import (
	"context"
	"encoding/json"

	"ZserverWebApiAuth/Zserver"
	"ZserverWebApiAuth/client"
	"ZserverWebApiAuth/models"
)

// WebGo_Imp servant implementation  _impWebWithContext
type WebGo_Imp struct {
}

// Init servant init
func (imp *WebGo_Imp) Init() error {
	// initialize servant here:
	// ...
	return nil
}

var loginLog models.LoginLog

func (imp *WebGo_Imp) LoginLog_Get(tarsCtx context.Context, req *Zserver.LoginLog, res *Zserver.LoginLog) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &loginLog)
	if err != nil {
		return err
	}

	result, err := loginLog.Get()
	if err != nil {
		return err
	}
	err = json.Unmarshal(client.Struct2Json(result), res)
	if err != nil {
		return err
	}
	return nil
}
func (imp *WebGo_Imp) LoginLog_GetPage(tarsCtx context.Context, pageSize int32, pageIndex int32, req *Zserver.LoginLog, res *Zserver.LoginLog_List) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &loginLog)
	if err != nil {
		return err
	}

	result, count, err := loginLog.GetPage(int(pageSize), int(pageIndex))
	if err != nil {
		return err
	}
	res.Count = int32(count)

	err = json.Unmarshal(client.Struct2Json(result), &res.LoginLogList)
	if err != nil {
		return err
	}

	return nil
}
func (imp *WebGo_Imp) LoginLog_Create(tarsCtx context.Context, req *Zserver.LoginLog, res *Zserver.LoginLog) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &loginLog)
	if err != nil {
		return err
	}

	result, err := loginLog.Create()
	if err != nil {
		return err
	}
	err = json.Unmarshal(client.Struct2Json(result), res)
	if err != nil {
		return err
	}
	return nil
}
func (imp *WebGo_Imp) LoginLog_Update(tarsCtx context.Context, id int32, req *Zserver.LoginLog, res *Zserver.LoginLog) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &loginLog)
	if err != nil {
		return err
	}

	result, err := loginLog.Update(int(id))
	if err != nil {
		return err
	}
	err = json.Unmarshal(client.Struct2Json(result), res)
	if err != nil {
		return err
	}
	return nil
}

func (imp *WebGo_Imp) LoginLog_BatchDelete(tarsCtx context.Context, id []int32, req *Zserver.LoginLog, res *bool) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &loginLog)
	if err != nil {
		return err
	}
	result, err := loginLog.BatchDelete(client.Ar2Int(id))
	if err != nil {
		return err
	}
	res = &result
	return nil
}

var sysConfig models.SysConfig

func (imp *WebGo_Imp) SysConfig_Create(tarsCtx context.Context, req *Zserver.SysConfig, res *Zserver.SysConfig) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysConfig)
	if err != nil {
		return err
	}

	result, err := sysConfig.Create()
	if err != nil {
		return err
	}
	err = json.Unmarshal(client.Struct2Json(result), res)
	if err != nil {
		return err
	}
	return nil
}
func (imp *WebGo_Imp) SysConfig_Get(tarsCtx context.Context, req *Zserver.SysConfig, res *Zserver.SysConfig) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysConfig)
	if err != nil {
		return err
	}

	result, err := sysConfig.Get()
	if err != nil {
		return err
	}
	err = json.Unmarshal(client.Struct2Json(result), res)
	if err != nil {
		return err
	}
	return nil
}
func (imp *WebGo_Imp) SysConfig_GetPage(tarsCtx context.Context, pageSize int32, pageIndex int32, req *Zserver.SysConfig, res *Zserver.SysConfig_List) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysConfig)
	if err != nil {
		return err
	}

	result, count, err := sysConfig.GetPage(int(pageSize), int(pageIndex))
	if err != nil {
		return err
	}
	res.Count = int32(count)

	err = json.Unmarshal(client.Struct2Json(result), &res.SysConfigList)
	if err != nil {
		return err
	}

	return nil
}
func (imp *WebGo_Imp) SysConfig_Update(tarsCtx context.Context, id int32, req *Zserver.SysConfig, res *Zserver.SysConfig) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysConfig)
	if err != nil {
		return err
	}

	result, err := sysConfig.Update(int(id))
	if err != nil {
		return err
	}
	err = json.Unmarshal(client.Struct2Json(result), res)
	if err != nil {
		return err
	}
	return nil
}
func (imp *WebGo_Imp) SysConfig_Delete(tarsCtx context.Context, req *Zserver.SysConfig, res *bool) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysConfig)
	if err != nil {
		return err
	}

	result, err := sysConfig.Delete()
	if err != nil {
		return err
	}
	*res = result
	return nil
}
func (imp *WebGo_Imp) SysConfig_BatchDelete(tarsCtx context.Context, id []int32, req *Zserver.SysConfig, res *bool) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysConfig)
	if err != nil {
		return err
	}
	result, err := sysConfig.BatchDelete(client.Ar2Int(id))
	if err != nil {
		return err
	}
	res = &result
	return nil
}

var sysUser models.SysUser

func (imp *WebGo_Imp) SysUser_Get(tarsCtx context.Context, req *Zserver.SysUser, res *Zserver.SysUser) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysUser)
	if err != nil {
		return err
	}

	result, err := sysUser.Get()
	if err != nil {
		return err
	}
	err = json.Unmarshal(client.Struct2Json(result), res)
	if err != nil {
		return err
	}
	return nil
}
func (imp *WebGo_Imp) SysUser_GetPage(tarsCtx context.Context, pageSize int32, pageIndex int32, req *Zserver.SysUser, res *Zserver.SysUser_List) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysUser)
	if err != nil {
		return err
	}

	result, count, err := sysUser.GetPage(int(pageSize), int(pageIndex))
	if err != nil {
		return err
	}
	res.Count = int32(count)

	err = json.Unmarshal(client.Struct2Json(result), &res.SysUserList)
	if err != nil {
		return err
	}

	return nil
}
func (imp *WebGo_Imp) SysUser_Insert(tarsCtx context.Context, req *Zserver.SysUser, id *int32) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysUser)
	if err != nil {
		return err
	}
	result, err := sysUser.Insert()
	*id = int32(result)
	return nil
}
func (imp *WebGo_Imp) SysUser_Update(tarsCtx context.Context, id int32, req *Zserver.SysUser, res *Zserver.SysUser) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysUser)
	if err != nil {
		return err
	}

	result, err := sysUser.Update(int(id))
	if err != nil {
		return err
	}
	err = json.Unmarshal(client.Struct2Json(result), res)
	if err != nil {
		return err
	}
	return nil
}
func (imp *WebGo_Imp) SysUser_BatchDelete(tarsCtx context.Context, id []int32, req *Zserver.SysUser, res *bool) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysUser)
	if err != nil {
		return err
	}
	result, err := sysUser.BatchDelete(client.Ar2Int(id))
	if err != nil {
		return err
	}
	res = &result
	return nil
}
func (imp *WebGo_Imp) SysUser_SetPwd(tarsCtx context.Context, oldPassword string, newPassword string, req *Zserver.SysUser, res *bool) (err error) {
	err = json.Unmarshal(client.Struct2Json(req), &sysUser)
	if err != nil {
		return err
	}
	result, err := sysUser.SetPwd(models.SysUserPwd{
		OldPassword: oldPassword,
		NewPassword: newPassword,
	})
	*res = result
	return nil
}

// Destroy servant destory
func (imp *WebGo_Imp) Destroy() {
	// destroy servant here:
	// ...
}
