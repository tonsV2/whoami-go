package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func WhoAmI(w http.ResponseWriter, r *http.Request) {
	log.Print("Endpoint Hit: " + r.URL.String())

	hostname, err := os.Hostname()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	_, _ = fmt.Fprint(w, hostname)
}
