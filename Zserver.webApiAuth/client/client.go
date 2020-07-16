package client

import "encoding/json"

func Struct2Struct(form interface{}, to interface{}) interface{} {
	jsonx, err := json.Marshal(form)
	if err != nil {
		return to
	}
	err = json.Unmarshal(jsonx, &to)
	if err != nil {
		return to
	}
	return to
}

func Ar2Int(from interface{}) (to []int) {
	to = make([]int, 0)
	for k, v := range from.([]interface{}) {
		switch v.(type) {
		case int8:
			to[k] = int(v.(int8))
		case int16:
			to[k] = int(v.(int16))
		case int32:
			to[k] = int(v.(int32))
		case int64:
			to[k] = int(v.(int64))
		}
	}
	return to
}

func Struct2Json(form interface{}) []byte {
	jsonx, err := json.Marshal(form)
	if err != nil {
		jsonx, err = json.Marshal(map[string]string{"err": err.Error()})
		if err != nil {
			return []byte(err.Error())
		}
	}
	return jsonx
}
