package tbsdk

import (
	"github.com/gogf/gf/frame/g"
	"strings"
	"testing"

	"git.tta.cn/all/serverAllWeb/utils"
)

func Test_getpage(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing getpage",
			args: args{url: "http://www.baidu.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.Getpage(tt.args.url); strings.Contains(got, "百度搜索") == false {
				t.Errorf("getpage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gets56qq(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []OutCp
		wantErr bool
	}{
		{
			name: "非接口取１０个商品",
			args: args{key: "中国"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSearch(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("gets56qq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("gets56qq() got = %v, want %v", got, tt.want)
			// }
			if len(got) != 10 {
				t.Errorf("gets56qq() got = %v, want %v", got, len(got))
			}
		})
	}
}

func TestGetsDP(t *testing.T) {
	type args struct {
		uid int
	}
	tests := []struct {
		name    string
		args    args
		want    []OutCp
		wantErr bool
	}{
		{
			name: "测试取shop",
			args: args{uid: 1633407959},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDP(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetsDP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g.Log().Warning(got)

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("GetsDP() got = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestGetCP(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "test read cp",
			args: args{id: 604271249091},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetCP(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
