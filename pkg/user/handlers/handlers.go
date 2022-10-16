package handlers

import (
	"encoding/json"
	"login-task/pkg/user/jwt"
	"login-task/pkg/user/models"
	"login-task/pkg/user/storage"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	res, err := jwt.TokenCheck(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "invalid auth", http.StatusUnauthorized)
	}
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
	var token models.Tokens
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := storage.GetUser(user.Login, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token.AccessToken, err = jwt.CreateAccessToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	token.RefreshToken, err = jwt.CreateRefreshToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	token.Id = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	var token models.Tokens
	var err error
	id := r.Header.Get("User")
	token.AccessToken, err = jwt.CreateAccessToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	token.RefreshToken, err = jwt.CreateRefreshToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	token.Id = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
