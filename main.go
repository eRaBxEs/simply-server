package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func greetByNameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := Response{Message: fmt.Sprintf("Hello, %s", name)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", greetHandler)
	router.HandleFunc("/you", greetByNameHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("server listening on port:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Print("could not start server!")
	}

}
