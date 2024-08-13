package helpers

import (
	"encoding/json"
	"io"
)

func FormatJson(body []byte, jsonData interface{}) interface{} {
	json.Unmarshal(body, &jsonData)
	return jsonData
}

func FormatBody(body io.Reader, jsonData interface{}) error {
	return json.NewDecoder(body).Decode(&jsonData)
}
