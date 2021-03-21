package internal

import (
	"encoding/json"
	"log"
	"net/http"
)

func Health(writer http.ResponseWriter, _ *http.Request) {
	log.Print("Endpoint Hit: /health - on a feature branch")

	writer.Header().Set("Content-Type", "application/json")

	type Health struct {
		Status string
	}

	health := Health{
		Status: "UP",
	}

	_ = json.NewEncoder(writer).Encode(health)
}
