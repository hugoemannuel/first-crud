package middleware

import (
	"bytes"
	"encoding/json"
	"first-crud/dto"
	"first-crud/helpers"
	"first-crud/service"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func ValidExclude(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		convertedId, _ := strconv.Atoi(id)
		user := service.FindOne(convertedId)
		if user.Name == "" {
			err := dto.DefaultError{Code: http.StatusNotFound, Message: "User not found"}
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
