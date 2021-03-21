package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func WhoAmI(w http.ResponseWriter, _ *http.Request) {
	log.Print("Endpoint Hit: /")

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	_, _ = fmt.Fprintf(w, hostname)
}
