package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func whoAmI(w http.ResponseWriter, _ *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprintf(w, hostname)
	log.Print("Endpoint Hit: /")
}

func handleRequests() {
	http.HandleFunc("/", whoAmI)

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
