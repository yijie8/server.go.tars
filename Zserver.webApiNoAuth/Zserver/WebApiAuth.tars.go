// Package Zserver comment
// This file was generated by tars2go 1.1.4
// Generated from Web.tars
package Zserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8
var _ = unsafe.Pointer(nil)

//WebApiAuth struct
type WebApiAuth struct {
	s m.Servant
}

//LoginLog_Get is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_Get(req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_Get", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = (*res).ReadBlock(_is, 2, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_GetWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_GetWithContext(tarsCtx context.Context, req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_Get", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = (*res).ReadBlock(_is, 2, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_GetOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_GetOneWayWithContext(tarsCtx context.Context, req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "LoginLog_Get", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_GetPage is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_GetPage(pageSize int32, pageIndex int32, req *LoginLog, res *LoginLog_List, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(pageSize, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(pageIndex, 2)
	if err != nil {
		return err
	}

	err = req.WriteBlock(_os, 3)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 4)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_GetPage", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = (*res).ReadBlock(_is, 4, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_GetPageWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_GetPageWithContext(tarsCtx context.Context, pageSize int32, pageIndex int32, req *LoginLog, res *LoginLog_List, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(pageSize, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(pageIndex, 2)
	if err != nil {
		return err
	}

	err = req.WriteBlock(_os, 3)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 4)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_GetPage", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = (*res).ReadBlock(_is, 4, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_GetPageOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_GetPageOneWayWithContext(tarsCtx context.Context, pageSize int32, pageIndex int32, req *LoginLog, res *LoginLog_List, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(pageSize, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(pageIndex, 2)
	if err != nil {
		return err
	}

	err = req.WriteBlock(_os, 3)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 4)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "LoginLog_GetPage", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_Create is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_Create(req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_Create", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = (*res).ReadBlock(_is, 2, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_CreateWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_CreateWithContext(tarsCtx context.Context, req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_Create", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = (*res).ReadBlock(_is, 2, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_CreateOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_CreateOneWayWithContext(tarsCtx context.Context, req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = req.WriteBlock(_os, 1)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "LoginLog_Create", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_Update is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_Update(id int32, req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(id, 1)
	if err != nil {
		return err
	}

	err = req.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 3)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_Update", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = (*res).ReadBlock(_is, 3, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_UpdateWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_UpdateWithContext(tarsCtx context.Context, id int32, req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(id, 1)
	if err != nil {
		return err
	}

	err = req.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 3)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_Update", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = (*res).ReadBlock(_is, 3, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_UpdateOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_UpdateOneWayWithContext(tarsCtx context.Context, id int32, req *LoginLog, res *LoginLog, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(id, 1)
	if err != nil {
		return err
	}

	err = req.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	err = (*res).WriteBlock(_os, 3)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "LoginLog_Update", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_BatchDelete is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_BatchDelete(id []int32, req *LoginLog, res *bool, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(id)), 0)
	if err != nil {
		return err
	}

	for _, v := range id {

		err = _os.Write_int32(v, 0)
		if err != nil {
			return err
		}

	}

	err = req.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	err = _os.Write_bool((*res), 3)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_BatchDelete", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_bool(&(*res), 3, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_BatchDeleteWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_BatchDeleteWithContext(tarsCtx context.Context, id []int32, req *LoginLog, res *bool, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(id)), 0)
	if err != nil {
		return err
	}

	for _, v := range id {

		err = _os.Write_int32(v, 0)
		if err != nil {
			return err
		}

	}

	err = req.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	err = _os.Write_bool((*res), 3)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "LoginLog_BatchDelete", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_bool(&(*res), 3, true)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//LoginLog_BatchDeleteOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *WebApiAuth) LoginLog_BatchDeleteOneWayWithContext(tarsCtx context.Context, id []int32, req *LoginLog, res *bool, _opt ...map[string]string) (err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(int32(len(id)), 0)
	if err != nil {
		return err
	}

	for _, v := range id {

		err = _os.Write_int32(v, 0)
		if err != nil {
			return err
		}

	}

	err = req.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	err = _os.Write_bool((*res), 3)
	if err != nil {
		return err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "LoginLog_BatchDelete", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return nil
}

//SetServant sets servant for the service.
func (_obj *WebApiAuth) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *WebApiAuth) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol sets the protocol for the servant.
func (_obj *WebApiAuth) TarsSetProtocol(p m.Protocol) {
	_obj.s.TarsSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *WebApiAuth) AddServant(imp _impWebApiAuth, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *WebApiAuth) AddServantWithContext(imp _impWebApiAuthWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impWebApiAuth interface {
	LoginLog_Get(req *LoginLog, res *LoginLog) (err error)
	LoginLog_GetPage(pageSize int32, pageIndex int32, req *LoginLog, res *LoginLog_List) (err error)
	LoginLog_Create(req *LoginLog, res *LoginLog) (err error)
	LoginLog_Update(id int32, req *LoginLog, res *LoginLog) (err error)
	LoginLog_BatchDelete(id []int32, req *LoginLog, res *bool) (err error)
}
type _impWebApiAuthWithContext interface {
	LoginLog_Get(tarsCtx context.Context, req *LoginLog, res *LoginLog) (err error)
	LoginLog_GetPage(tarsCtx context.Context, pageSize int32, pageIndex int32, req *LoginLog, res *LoginLog_List) (err error)
	LoginLog_Create(tarsCtx context.Context, req *LoginLog, res *LoginLog) (err error)
	LoginLog_Update(tarsCtx context.Context, id int32, req *LoginLog, res *LoginLog) (err error)
	LoginLog_BatchDelete(tarsCtx context.Context, id []int32, req *LoginLog, res *bool) (err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the tars file. _withContext shows using context or not.
func (_obj *WebApiAuth) Dispatch(tarsCtx context.Context, _val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, _withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	_os := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "LoginLog_Get":
		var req LoginLog
		var res LoginLog

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		if _withContext == false {
			_imp := _val.(_impWebApiAuth)
			err = _imp.LoginLog_Get(&req, &res)
		} else {
			_imp := _val.(_impWebApiAuthWithContext)
			err = _imp.LoginLog_Get(tarsCtx, &req, &res)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = res.WriteBlock(_os, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			_os.Reset()
			err = res.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("res", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["res"] = res

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "LoginLog_GetPage":
		var pageSize int32
		var pageIndex int32
		var req LoginLog
		var res LoginLog_List

		if tarsReq.IVersion == basef.TARSVERSION {

			err = _is.Read_int32(&pageSize, 1, true)
			if err != nil {
				return err
			}

			err = _is.Read_int32(&pageIndex, 2, true)
			if err != nil {
				return err
			}

			err = req.ReadBlock(_is, 3, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("pageSize", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_int32(&pageSize, 0, true)
			if err != nil {
				return err
			}

			_reqTup_.GetBuffer("pageIndex", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_int32(&pageIndex, 0, true)
			if err != nil {
				return err
			}

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["pageSize"])
				if err = json.Unmarshal([]byte(_jsonStr_), &pageSize); err != nil {
					return err
				}
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["pageIndex"])
				if err = json.Unmarshal([]byte(_jsonStr_), &pageIndex); err != nil {
					return err
				}
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		if _withContext == false {
			_imp := _val.(_impWebApiAuth)
			err = _imp.LoginLog_GetPage(pageSize, pageIndex, &req, &res)
		} else {
			_imp := _val.(_impWebApiAuthWithContext)
			err = _imp.LoginLog_GetPage(tarsCtx, pageSize, pageIndex, &req, &res)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = res.WriteBlock(_os, 4)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			_os.Reset()
			err = res.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("res", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["res"] = res

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "LoginLog_Create":
		var req LoginLog
		var res LoginLog

		if tarsReq.IVersion == basef.TARSVERSION {

			err = req.ReadBlock(_is, 1, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		if _withContext == false {
			_imp := _val.(_impWebApiAuth)
			err = _imp.LoginLog_Create(&req, &res)
		} else {
			_imp := _val.(_impWebApiAuthWithContext)
			err = _imp.LoginLog_Create(tarsCtx, &req, &res)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = res.WriteBlock(_os, 2)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			_os.Reset()
			err = res.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("res", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["res"] = res

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "LoginLog_Update":
		var id int32
		var req LoginLog
		var res LoginLog

		if tarsReq.IVersion == basef.TARSVERSION {

			err = _is.Read_int32(&id, 1, true)
			if err != nil {
				return err
			}

			err = req.ReadBlock(_is, 2, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("id", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_int32(&id, 0, true)
			if err != nil {
				return err
			}

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["id"])
				if err = json.Unmarshal([]byte(_jsonStr_), &id); err != nil {
					return err
				}
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		if _withContext == false {
			_imp := _val.(_impWebApiAuth)
			err = _imp.LoginLog_Update(id, &req, &res)
		} else {
			_imp := _val.(_impWebApiAuthWithContext)
			err = _imp.LoginLog_Update(tarsCtx, id, &req, &res)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = res.WriteBlock(_os, 3)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			_os.Reset()
			err = res.WriteBlock(_os, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("res", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["res"] = res

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "LoginLog_BatchDelete":
		var id []int32
		var req LoginLog
		var res bool

		if tarsReq.IVersion == basef.TARSVERSION {

			err, have, ty = _is.SkipToNoCheck(1, true)
			if err != nil {
				return err
			}

			if ty == codec.LIST {
				err = _is.Read_int32(&length, 0, true)
				if err != nil {
					return err
				}

				id = make([]int32, length)
				for i1, e1 := int32(0), length; i1 < e1; i1++ {

					err = _is.Read_int32(&id[i1], 0, false)
					if err != nil {
						return err
					}

				}
			} else if ty == codec.SIMPLE_LIST {
				err = fmt.Errorf("not support simple_list type")
				if err != nil {
					return err
				}

			} else {
				err = fmt.Errorf("require vector, but not")
				if err != nil {
					return err
				}

			}

			err = req.ReadBlock(_is, 2, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("id", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err, have, ty = _is.SkipToNoCheck(0, true)
			if err != nil {
				return err
			}

			if ty == codec.LIST {
				err = _is.Read_int32(&length, 0, true)
				if err != nil {
					return err
				}

				id = make([]int32, length)
				for i2, e2 := int32(0), length; i2 < e2; i2++ {

					err = _is.Read_int32(&id[i2], 0, false)
					if err != nil {
						return err
					}

				}
			} else if ty == codec.SIMPLE_LIST {
				err = fmt.Errorf("not support simple_list type")
				if err != nil {
					return err
				}

			} else {
				err = fmt.Errorf("require vector, but not")
				if err != nil {
					return err
				}

			}

			_reqTup_.GetBuffer("req", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = req.ReadBlock(_is, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["id"])
				if err = json.Unmarshal([]byte(_jsonStr_), &id); err != nil {
					return err
				}
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["req"])
				if err = json.Unmarshal([]byte(_jsonStr_), &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		if _withContext == false {
			_imp := _val.(_impWebApiAuth)
			err = _imp.LoginLog_BatchDelete(id, &req, &res)
		} else {
			_imp := _val.(_impWebApiAuthWithContext)
			err = _imp.LoginLog_BatchDelete(tarsCtx, id, &req, &res)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_bool(res, 3)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			_os.Reset()
			err = _os.Write_bool(res, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("res", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["res"] = res

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(tarsCtx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(tarsCtx)
	if ok && c != nil {
		_context = c
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}

	_ = _is
	_ = _os
	_ = length
	_ = have
	_ = ty
	return nil
}
