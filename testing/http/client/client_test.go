package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDoStuffWithTestServer(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		q := req.URL.Query()
		id := q["id"]
		//id := q["id"]
		// Send response to be tested
		fmt.Fprintf(w, "testing"+id[0])
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{server.Client(), server.URL + "?id=34"}
	body, err := api.DoStuff()
	if err != nil {
		t.Errorf("err : %s", err)
	} else {
		fmt.Println(string(body))
	}

	if string(body) != "testing" {
		t.Errorf("body not matching: %s", string(body))
	}
}
