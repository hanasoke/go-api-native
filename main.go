package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	log.Println("Server running on port 8080")
	http.ListenAndServe(fmt.Sprintf(":%v", "8080"), r)

	log.Println()
}
