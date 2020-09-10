package ldap

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func init() {
	viper.SetConfigFile("../settings.yml")
	content, err := ioutil.ReadFile("../settings.yml")
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	ldapc := viper.Sub("settings.ldap")
	if ldapc == nil {
		panic("config not found ldap")
	}
}

func Test_conn(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"ldap测试",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LdapConnInit()
		})
	}
}

func TestAdd(t *testing.T) {
	LdapConnInit()
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "添加测试",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}