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

type GreetingRequest struct {
	Names []string `json:"names"`
}

type GreetingResponse struct {
	Greetings []string `json:"greetings"`
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

func greetAllNamesHandler(w http.ResponseWriter, r *http.Request) {
	// http verb method check
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST request are accepted!", http.StatusMethodNotAllowed)
		return
	}

	// process request
	var req GreetingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalied request body", http.StatusBadRequest)
		return
	}

	var res GreetingResponse

	for _, name := range req.Names {
		res.Greetings = append(res.Greetings, fmt.Sprintf("Hello, %s!", name))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", greetHandler)
	router.HandleFunc("/you", greetByNameHandler)
	router.HandleFunc("/greet", greetAllNamesHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("server listening on port:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Print("could not start server!")
	}

}
