package tbsdk

import (
	"testing"
)

func TestGetXG_id(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    OutCpXg
		wantErr bool
	}{
		{
			name: "获取相关商品 id",
			args: args{id: 542504679550},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetXGId(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetXGId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// t.Log(got)
			_ = got
		})
	}
}

func TestGetXG_key(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    OutCpXg
		wantErr bool
	}{
		{
			name: "获取相关商品　by key",
			args: args{key: "欧实木床现代简约主卧婚床双人1.8米家用"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetXGKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetXGKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// t.Log(got)
			_ = got
		})
	}
}
