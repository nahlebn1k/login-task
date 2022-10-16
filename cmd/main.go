package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	//_ "github.com/swaggo/http-swagger"
	"log"
	//_ "login-task/docs"
	"login-task/pkg/configs"
	"login-task/pkg/user/handlers"
	"net/http"
)

// @title       LoginApp
// @version     1.0
// @description Login API

// @host     localhost:8000
// @BasePath /

// @securityDefinitions.apikey AuthKey
// @in                         header
// @name                       Authorization
func main() {
	config := configs.GetConfig()
	fs := http.FileServer(http.Dir("web/static"))
	router := mux.NewRouter()
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/signin", handlers.SingIn).Methods("POST")
	router.HandleFunc("/user", handlers.User).Methods("POST")
	router.HandleFunc("/refresh", handlers.RefreshTokens).Methods("POST")
	router.PathPrefix("/").Handler(fs)
	handler := cors.Default().Handler(router)
	if err := http.ListenAndServe(":"+config.AppPort, handler); err != nil {
		log.Fatal("Server error!")
	}
}
