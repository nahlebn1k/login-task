package handlers

import (
	"encoding/json"
	"login-task/pkg/user/models"
	"login-task/pkg/user/storage"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	storage.CreateUser(user.Login, user.Password)
	w.Write([]byte("OK"))
}
