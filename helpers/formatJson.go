package helpers

import (
	"encoding/json"
)

func FormatJson(body []byte, jsonData interface{}) interface{} {
	json.Unmarshal(body, &jsonData)
	return jsonData
}
