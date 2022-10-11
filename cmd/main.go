package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"login-task/pkg/configs"
	"login-task/pkg/user/handlers"
	"net/http"
)

func main() {
	config := configs.GetConfig()
	fs := http.FileServer(http.Dir("web/static"))
	router := mux.NewRouter()
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	router.PathPrefix("/").Handler(fs)
	handler := cors.Default().Handler(router)

	if err := http.ListenAndServe(":"+config.AppPort, handler); err != nil {
		log.Fatal("Server error!")
	}
}
