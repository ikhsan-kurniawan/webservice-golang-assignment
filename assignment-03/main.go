package main

import (
	"fmt"
	"html/template"
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

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	tpl := template.Must(template.ParseFiles("index.html"))

	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	status := Status{
		Water: water,
		Wind:  wind,
	}

	response := Response {
		Status: status,
	}

	tpl.Execute(w, response)
	return
}

