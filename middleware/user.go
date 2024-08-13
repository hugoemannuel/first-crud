package middleware

import (
	"bytes"
	"encoding/json"
	"first-crud/dto"
	"first-crud/helpers"
	"io"
	"net/http"
)

func ValidCreate(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		var jsonData dto.User
		helpers.FormatJson(body, &jsonData)
		if jsonData.Name == "" || jsonData.Email == "" || jsonData.Password == "" {
			err := dto.DefaultError{Code: http.StatusBadRequest, Message: "Invalid Request"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		handler.ServeHTTP(w, r)
	})

}
