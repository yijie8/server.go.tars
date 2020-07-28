package tbsdk

import (
	"github.com/gogf/gf/os/glog"
	"reflect"
	"testing"

	"github.com/bitly/go-simplejson"
)

func TestTbsdk_GetTB(t1 *testing.T) {
	type fields struct {
		Api    string
		AppKey string
		Args   map[string]interface{}
		Zid    string
		Result *simplejson.Json
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *simplejson.Json
	}{
		{
			name: "主函数测试",
			args: args{[]string{"tta"}},
			fields: fields{
				Api: "taobao.tbk.item.recommend.get",
				Args: map[string]interface{}{
					"fields":   "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url",
					"num_iid":  590455411752,
					"count":    20,
					"platform": 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:    tt.fields.Api,
				AppKey: tt.fields.AppKey,
				Args:   tt.fields.Args,
				Zid:    tt.fields.Zid,
				Result: tt.fields.Result,
			}
			if got, err := t.GetTB(tt.args.args...).Get("tbk_item_recommend_get_response").Get("results").Get("n_tbk_item").Array(); err != nil {
				t1.Errorf("GetTB() = %v, err %v", got, err)
			}
		})
	}
}

func TestTbsdk_GetDPXgListKey(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		key  string
		page []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "取相关店铺key",
			args: args{key: "中国"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got, err := t.GetDPXgListKey(tt.args.key, tt.args.page...).ToStruct(); err != nil {
				t1.Errorf("GetDPXgListKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetContent(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		args   map[string]interface{}
		appkey []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "获取图文内容输出 ",
			args: args{
				args:   map[string]interface{}{},
				appkey: []string{"62a"},
			},
			want: "{}",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got, err := t.GetContent(tt.args.args, tt.args.appkey...).ToStruct(); err != nil {
				t1.Errorf("GetContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetCp(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		cpids string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		{
			name: "取单个商品的详细信息",
			args: args{cpids: "42745500604"},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got := t.GetCp(tt.args.cpids); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetCp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetCpList(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		args   map[string]interface{}
		appkey []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		{
			name: "商品列表测试",
			args: args{
				args: map[string]interface{}{
					"q": "中国",
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}

			got := t.GetCpList(tt.args.args, tt.args.appkey...)
			glog.Debug(got)
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetCpList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetCpListGood(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		args   map[string]interface{}
		appkey []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got := t.GetCpListGood(tt.args.args, tt.args.appkey...); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetCpListGood() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetCpQDesc(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		args map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got := t.GetCpQDesc(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetCpQDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetCpXgList(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		cpid int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got := t.GetCpXgList(tt.args.cpid); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetCpXgList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetCreateTLJ(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		args   map[string]interface{}
		appkey []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got := t.GetCreateTLJ(tt.args.args, tt.args.appkey...); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetCreateTLJ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetDPXgListId(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		dpid int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got := t.GetDPXgListId(tt.args.dpid); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetDPXgListId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetTGG(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		args   map[string]interface{}
		appkey []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got := t.GetTGG(tt.args.args, tt.args.appkey...); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTGG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetTKL(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		text string
		url  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			got, err := t.GetTKL(tt.args.text, tt.args.url)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetTKL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("GetTKL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTbsdk_GetTLJDesc(t1 *testing.T) {
	type fields struct {
		Api     string
		AppKey  string
		Args    map[string]interface{}
		Zid     string
		KeyWord string
		Result  *simplejson.Json
	}
	type args struct {
		rights_id string
		appkey    []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tbsdk
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tbsdk{
				Api:     tt.fields.Api,
				AppKey:  tt.fields.AppKey,
				Args:    tt.fields.Args,
				Zid:     tt.fields.Zid,
				KeyWord: tt.fields.KeyWord,
				Result:  tt.fields.Result,
			}
			if got := t.GetTLJDesc(tt.args.rights_id, tt.args.appkey...); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetTLJDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}
