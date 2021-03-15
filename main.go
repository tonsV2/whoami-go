package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func whoAmI(w http.ResponseWriter, _ *http.Request) {
	log.Print("Endpoint Hit: /")

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprintf(w, hostname)
}

func health(writer http.ResponseWriter, _ *http.Request) {
	log.Print("Endpoint Hit: /health")

	writer.Header().Set("Content-Type", "application/json")

	type Health struct {
		Status string
	}

	health := Health{
		Status: "UP",
	}

	_ = json.NewEncoder(writer).Encode(health)
}

func handleRequests() {
	http.HandleFunc("/", whoAmI)
	http.HandleFunc("/health", health)

	port := "8080"
	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}

	log.Printf("Server listening on port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	handleRequests()
}
