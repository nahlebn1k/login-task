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
	router.PathPrefix("/").Handler(fs)
	router.HandleFunc("/sign-up", handlers.SignUp).Methods("POST")
	handler := cors.Default().Handler(router)

	if err := http.ListenAndServe(":"+config.App.Port, handler); err != nil {
		log.Fatal("Server error!")
	}
}
