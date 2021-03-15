package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func whoAmI(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprintf(w, hostname)
	log.Print("Endpoint Hit: /")
}

func handleRequests() {
	http.HandleFunc("/", whoAmI)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
