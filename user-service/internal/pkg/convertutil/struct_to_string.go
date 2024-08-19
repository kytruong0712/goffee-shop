package convertutil

import "encoding/json"

func ConvertStructToString(v interface{}) string {
	r, _ := json.Marshal(v)

	return string(r)
}
