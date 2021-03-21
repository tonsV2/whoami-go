package test

import (
	. "github.com/tonsV2/whoami-go/internal"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestWhoAmIEndpoint(t *testing.T) {
	t.Run("returns server hostname", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		WhoAmI(response, request)

		got := response.Body.String()
		want, _ := os.Hostname()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
