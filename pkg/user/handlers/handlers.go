package handlers

import (
	"encoding/json"
	"login-task/pkg/user/jwt"
	"login-task/pkg/user/models"
	"login-task/pkg/user/storage"
	"net/http"
)

func UserAuth(w http.ResponseWriter, r *http.Request) {
	res, err := jwt.TokenCheck(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "invalid auth", http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res))
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	storage.CreateUser(user.Login, user.Password)
	w.Write([]byte("OK"))
}

func SingIn(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin
	var token models.Token
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id, err := storage.GetUser(user.Login, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token.TokenString, err = jwt.CreateToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//w.Write([]byte("OK"))
}
