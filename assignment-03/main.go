package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

var PORT = ":8080"
type Response struct {
	Status Status `json:"status"`
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/status", getStatus)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	status := Status{
		Water: water,
		Wind:  wind,
	}

	response := Response{
		Status: status,
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

