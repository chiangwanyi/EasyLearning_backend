package util

import (
	"encoding/json"
)

// Json 返回 JSON 数据
func Json(data interface{}) []byte {
	res, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return res
}
