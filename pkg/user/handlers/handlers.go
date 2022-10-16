package handlers

import (
	"encoding/json"
	"login-task/pkg/user/jwt"
	"login-task/pkg/user/models"
	"login-task/pkg/user/storage"
	"net/http"
)

// @Summary     TestUserAuth
// @Description Check user auth with token
// @ID          test-auth
// @Router      /user

func User(w http.ResponseWriter, r *http.Request) {
	res, err := jwt.TokenCheck(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "invalid auth", http.StatusUnauthorized)
	}
	w.Write([]byte(res))
}

// @Summary     SignUp
// @Description Signing up user and creates user in db
// @ID          create-user
// @Accept      json
// @Router      /signup

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	storage.CreateUser(user.Login, user.Password)
	w.Write([]byte("OK"))
}

// @Summary     SignIn
// @Description Signing in user and give tokens
// @ID          sign-in
// @Accept      json
// @Produce     json
// @Router      /signin

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

// @Summary     Refresh expired tokens
// @Description generate new access tokens
// @ID          refresh-token
// @Produce     json
// @Router      /refresh

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
