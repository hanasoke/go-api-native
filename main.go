package main

import (
	"fmt"
	"go-api-native/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()

	log.Println("Server running on port", config.ENV.PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)

	log.Println()
}
