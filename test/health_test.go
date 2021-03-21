package test

import (
	. "github.com/tonsV2/whoami-go/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	t.Run("returns status message", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		response := httptest.NewRecorder()

		Health(response, request)

		got := response.Body.String()
		want := "{\"Status\":\"UP\"}\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
