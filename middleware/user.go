package middleware

import (
	"bytes"
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		handler.ServeHTTP(w, r)
	})

}
