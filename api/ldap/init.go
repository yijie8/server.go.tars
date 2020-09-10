package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/spf13/viper"
	"log"
)

type LdapConfig struct {
	Host   string
	Port   string
	BindDn string
	Passwd string
	BaseDn string
}

var (
	LdapConfig_e LdapConfig
	LdapConn     *ldap.Conn
)

func LdapConnInit() {
	var err error
	m := viper.Sub("settings.ldap")
	if m == nil {
		panic("ldap配置文件有错误")
	}

	LdapConfig_e.Host = m.GetString("host")
	LdapConfig_e.Port = m.GetString("port")
	LdapConfig_e.BindDn = m.GetString("binddn")
	LdapConfig_e.Passwd = m.GetString("passwd")
	LdapConfig_e.BaseDn = m.GetString("basedn")

	LdapConn, err = ldap.Dial("tcp", LdapConfig_e.Host+":"+LdapConfig_e.Port)

	if err != nil {
		panic("无法连接ldap服务器 " + err.Error())
	}
	//defer LdapConn.Close()
	err = LdapConn.Bind("uid=test,dc=tta,dc=cn", "testtest")
	//err = LdapConn.Bind(LdapConfig_e.BindDn, LdapConfig_e.Passwd)
	if err != nil {
		panic("用户名密码错 " + err.Error())
	}
}

func Add() bool {
	if LdapConn == nil {
		LdapConnInit()
	}

	var userDn = fmt.Sprintf("uid=%s,%s", "test2", LdapConfig_e.BaseDn)

	sql := ldap.NewAddRequest(userDn, nil)
	sql.Attribute("uidNumber", []string{"10446"})
	sql.Attribute("gidNumber", []string{"0"})
	sql.Attribute("userPassword", []string{"testtest"})
	sql.Attribute("homeDirectory", []string{"/var/lib/ldap/test2"})

	sql.Attribute("cn", []string{"test2"})
	sql.Attribute("uid", []string{"test2"})
	sql.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "account"})

	if err := LdapConn.Add(sql); err != nil {
		if ldap.IsErrorWithCode(err, 68) {
			log.Panic("User already exist." + err.Error())
		} else {
			log.Panic("User insert error." + err.Error())
		}
		return false
	}
	return true
}
