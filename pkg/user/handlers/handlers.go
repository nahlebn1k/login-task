package handlers

import (
	"encoding/json"
	"login-task/pkg/user/jwt"
	"login-task/pkg/user/models"
	"login-task/pkg/user/storage"
	"net/http"
	"strings"
)

func User(w http.ResponseWriter, r *http.Request) {
	res, err := jwt.TokenCheck(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "invalid auth", http.StatusUnauthorized)
		return
	}
	if _, err := w.Write([]byte(res)); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := storage.CreateUser(user.Login, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func SingIn(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin
	var token models.Tokens
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := storage.GetUser(user.Login, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token.AccessToken, err = jwt.CreateAccessToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token.RefreshToken, err = jwt.CreateRefreshToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	var token models.Tokens
	var err error
	var idResp string
	var tokenDB string
	_, err = jwt.TokenCheck(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "invalid auth", http.StatusUnauthorized)
		return
	}
	id := r.Header.Get("User")
	if err != nil || id == "" {
		http.Error(w, "invalid auth", http.StatusUnauthorized)
		return
	}
	tokenDB, err = storage.GetRefreshToken(id)
	if err != nil {
		http.Error(w, "invalid auth", http.StatusUnauthorized)
		return
	}
	tokenHeader := r.Header.Get("Authorization")
	tokenReq := strings.Split(tokenHeader, " ")

	if tokenDB != tokenReq[1] {
		http.Error(w, "invalid auth", http.StatusUnauthorized)
		return
	}

	token.AccessToken, err = jwt.CreateAccessToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token.RefreshToken, err = jwt.CreateRefreshToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token.ID = idResp
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
