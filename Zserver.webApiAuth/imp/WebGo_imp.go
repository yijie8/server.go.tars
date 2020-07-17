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

// Destroy servant destory
func (imp *WebGo_Imp) Destroy() {
	// destroy servant here:
	// ...
}
