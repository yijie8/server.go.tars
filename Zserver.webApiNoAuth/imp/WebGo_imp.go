package imp

import (
	"context"
	"ZserverWebApiNoAuth/Zserver"
	"ZserverWebApiNoAuth/apis/monitor"
	"ZserverWebApiNoAuth/apis/system"
	"ZserverWebApiNoAuth/apis/system/dict"
	"ZserverWebApiNoAuth/apis/tools"
)

// WebGo_Imp servant implementation  _impWebWithContext
type WebGo_Imp struct {
}

func (imp *WebGo_Imp) MenuTreeselect(tarsCtx context.Context, res *Zserver.GetMenuTreeelect_res) (err error) {
	return system.GetMenuTreeelect_tar(res)
}
func (imp *WebGo_Imp) Monitor_server(tarsCtx context.Context, res *Zserver.Monitor_server_res) (err error) {
	return monitor.ServerInfo_tars(res)
}

func (imp *WebGo_Imp) GetCaptcha(tarsCtx context.Context, res *Zserver.GetCaptcha_res) (err error) {
	return system.GetCaptcha_tars(res)
}
func (imp *WebGo_Imp) Preview(tarsCtx context.Context, id int32, res *Zserver.Preview_res) (err error) {
	return tools.Preview_tars(int(id), res)
}

func (imp *WebGo_Imp) GetDictDataByDictType(tarsCtx context.Context, dictType string, res *Zserver.DictData_res) (err error) {
	return dict.GetDictDataByDictType_tars(dictType, res)
}

func (imp *WebGo_Imp) GetDBTableList(tarsCtx context.Context, Req *Zserver.GetDBTableList_req, Res *Zserver.GetDBTableList_res) (err error) {
	return tools.GetDBTableList_tars(Req, Res)
}
func (imp *WebGo_Imp) GetDBColumnList(tarsCtx context.Context, Req *Zserver.GetDBColumnList_req, Res *Zserver.GetDBColumnList_res) (err error) {
	return tools.GetDBColumnList_tars(Req, Res)
}

func (imp *WebGo_Imp) GetSysTableList(tarsCtx context.Context, Req *Zserver.Table_req, Res *Zserver.SysTables_res) (err error) {
	return tools.GetSysTableList_tars(Req, Res)
}

func (imp *WebGo_Imp) InsertSysTable(tarsCtx context.Context, req *Zserver.InsertSysTable_req, res *Zserver.Response_res) (err error) {
	return tools.InsertSysTable_tars(req, res)
}

func (imp *WebGo_Imp) DeleteSysTables(tarsCtx context.Context, tableId string, res *Zserver.Response_res) (err error) {
	return tools.DeleteSysTables_tars(tableId, res)
}
func (imp *WebGo_Imp) UpdateSysTable(tarsCtx context.Context, uid string, req *Zserver.SysTables_List, res *Zserver.Response_res) (err error) {
	return tools.UpdateSysTable_tars(uid, req, res)
}

func (imp *WebGo_Imp) GetSysTables(tarsCtx context.Context, tableId string, res *Zserver.SysTables_one_res) (err error) {
	return tools.GetSysTables_tars(tableId, res)
}

// Init servant init
func (imp *WebGo_Imp) Init() error {
	//initialize servant here:
	//...
	return nil
}

// Destroy servant destory
func (imp *WebGo_Imp) Destroy() {
	//destroy servant here:
	//...
}
