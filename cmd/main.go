package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"login-task/pkg/configs"
	"net/http"
)

func main() {
	config := configs.GetConfig()
	router := mux.NewRouter()
	handler := cors.Default().Handler(router)

	if err := http.ListenAndServe(":"+config.App.Port, handler); err != nil {
		log.Fatal("Server error!")
	}
}
