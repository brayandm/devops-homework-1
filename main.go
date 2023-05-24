package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Simple struct {
	Name        string
	Description string
	Url         string
}

func SimpleFactory(host string) Simple {
	return Simple{"Hello", "Brayan Duran Medina", "http://localhost1"}
}

func handler(w http.ResponseWriter, r *http.Request) {
	simple := Simple{"Hello", "Brayan Duran Medina", r.Host}

	jsonOutput, _ := json.Marshal(simple)

	fmt.Fprintln(w, string(jsonOutput))
}

func main() {
	fmt.Println("Server started on port 4444")
	http.HandleFunc("/api", handler)
	log.Fatal(http.ListenAndServe(":4444", nil))
}