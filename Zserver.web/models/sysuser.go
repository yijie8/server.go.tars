package models

import (
	"ZserverWeb/Zserver"
	"ZserverWeb/client"
	"encoding/json"
	"github.com/yijie8/zserver/tools"
	"golang.org/x/crypto/bcrypt"
)

// User
type User struct {
	// key
	IdentityKey string
	// 用户名
	UserName  string
	FirstName string
	LastName  string
	// 角色
	Role string
}

type UserName struct {
	Username string `gorm:"type:varchar(64)" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type SysUserB struct {
	NickName  string `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	Phone     string `gorm:"type:varchar(11)" json:"phone"`     // 手机号
	RoleId    int    `gorm:"type:int(11)" json:"roleId"`        // 角色编码
	Salt      string `gorm:"type:varchar(255)" json:"salt"`     //盐
	Avatar    string `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Sex       string `gorm:"type:varchar(255)" json:"sex"`      //性别
	Email     string `gorm:"type:varchar(128)" json:"email"`    //邮箱
	DeptId    int    `gorm:"type:int(11)" json:"deptId"`        //部门编码
	PostId    int    `gorm:"type:int(11)" json:"postId"`        //职位编码
	CreateBy  string `gorm:"type:varchar(128)" json:"createBy"` //
	UpdateBy  string `gorm:"type:varchar(128)" json:"updateBy"` //
	Remark    string `gorm:"type:varchar(255)" json:"remark"`   //备注
	Status    string `gorm:"type:int(1);" json:"status"`
	DataScope string `gorm:"-" json:"dataScope"`
	Params    string `gorm:"-" json:"params"`

	BaseModel
}

type SysUser struct {
	SysUserId
	SysUserB
	LoginM
}

func (SysUser) TableName() string {
	return "sys_user"
}

type SysUserPwd struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type SysUserPage struct {
	SysUserId
	SysUserB
	LoginM
	DeptName string `gorm:"-" json:"deptName"`
}

type SysUserView struct {
	SysUserId
	SysUserB
	LoginM
	RoleName string `gorm:"column:role_name"  json:"role_name"`
}

var req_sysuser Zserver.SysUser
var res_sysuser Zserver.SysUser
var ress_sysuser Zserver.SysUser_List

// 获取用户数据
func (e *SysUser) Get() (SysUserView SysUserView, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_sysuser)
	if err != nil {
		return SysUserView, err
	}
	err = client.WebApiAuth().SysUser_Get(&req_sysuser, &res_sysuser)
	if err != nil {
		return SysUserView, err
	}
	err = json.Unmarshal(client.Struct2Json(res_sysuser), &SysUserView)
	if err != nil {
		return SysUserView, err
	}
	//SysUserView.CreatedAt = time.Now()
	//SysUserView.UpdatedAt = time.Now()
	//SysUserView.DeletedAt = time.Now()
	return SysUserView, nil
}

func (e *SysUser) GetPage(pageSize int, pageIndex int) ([]SysUserPage, int, error) {
	var count int
	var docs []SysUserPage
	err := json.Unmarshal(client.Struct2Json(e), &req_sysuser)
	if err != nil {
		return docs, count, err
	}
	err = client.WebApiAuth().SysUser_GetPage(int32(pageSize), int32(pageIndex), &req_sysuser, &ress_sysuser)
	if err != nil {
		return docs, count, err
	}
	count = int(ress_sysuser.Count)
	err = json.Unmarshal(client.Struct2Json(ress_sysuser.SysUserList), &docs)
	if err != nil {
		return docs, count, err
	}
	return docs, count, nil
}

//加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

//添加
func (e SysUser) Insert() (id int, err error) {
	var id32 int32
	err = json.Unmarshal(client.Struct2Json(e), &req_sysuser)
	if err != nil {
		return 0, err
	}

	err = client.WebApiAuth().SysUser_Insert(&req_sysuser, &id32)
	if err != nil {
		return 0, err
	}
	return int(id32), nil
}

//修改
func (e *SysUser) Update(id int) (update SysUser, err error) {
	if e.Password != "" {
		if err = e.Encrypt(); err != nil {
			return
		}
	}
	err = json.Unmarshal(client.Struct2Json(e), &req_sysuser)
	if err != nil {
		return update, err
	}
	err = client.WebApiAuth().SysUser_Update(int32(id), &req_sysuser, &res_sysuser)
	if err != nil {
		return update, err
	}
	err = json.Unmarshal(client.Struct2Json(res_sysuser), &update)
	if err != nil {
		return update, err
	}
	return
}

func (e *SysUser) BatchDelete(id []int) (Result bool, err error) {

	err = json.Unmarshal(client.Struct2Json(e), &req_sysuser)
	if err != nil {
		return false, err
	}
	err = client.WebApiAuth().SysUser_BatchDelete(client.Ar2Int32(id), &req_sysuser, &Result)
	if err != nil {
		return false, err
	}
	Result = true
	return
}

func (e *SysUser) SetPwd(pwd SysUserPwd) (Result bool, err error) {
	err = json.Unmarshal(client.Struct2Json(e), &req_sysuser)
	if err != nil {
		tools.HasError(err, "获取用户数据失败(代码202)", 500)
		return false, err
	}
	err = client.WebApiAuth().SysUser_SetPwd(pwd.OldPassword, pwd.NewPassword, &req_sysuser, &Result)
	if err != nil {
		tools.HasError(err, "获取用户数据失败(代码202)", 500)
		return false, err
	}
	return
}
