package controller

import (
	"encoding/json"
	"first-crud/dto"
	"first-crud/helpers"
	"first-crud/service"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	users := service.GetAll()
	json.NewEncoder(w).Encode(users)
}

func UserFindOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user := service.FindOne(convertedId)
	json.NewEncoder(w).Encode(user)
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	var formatBody dto.User

	helpers.FormatJson(body, &formatBody)

	user, errQuery := service.Create(formatBody.Name, formatBody.Password, formatBody.Email)

	if errQuery != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)

}

func UserExclude(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	convertedId, _ := strconv.Atoi(id)
	user := service.FindOne(convertedId)
	if user.Name == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	service.Exclude(convertedId)
	w.WriteHeader(http.StatusNoContent)
}

func Update(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	vars := mux.Vars(r)
	id := vars["id"]
	convertedId, _ := strconv.Atoi(id)
	var formatBody dto.User

	helpers.FormatJson(body, &formatBody)
	user, errQuery := service.Update(convertedId, formatBody.Name, formatBody.Email, formatBody.Password)

	if errQuery != nil {
		http.Error(w, "Unable to read request body", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
