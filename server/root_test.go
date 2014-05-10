package server

import (
	"log"
	"net/http"
	"testing"

	"github.com/sourcegraph/vcsstore/client"
)

func TestHandler_serveRoot(t *testing.T) {
	setupHandlerTest()
	defer teardownHandlerTest()

	resp, err := http.Get(server.URL + router.URLTo(client.RouteRoot).String())
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("got code %d, want %d", got, want)
	}
}