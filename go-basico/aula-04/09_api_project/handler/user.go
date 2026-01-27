package handler

//Camada de manipulador (handler) - Lida com as requisições HTTP
//Normalmente interage com a camada de serviço (service)

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/diego/go-basico-01/service"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := service.GetAllUsers()
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	user, err := service.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

