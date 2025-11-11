package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func fatalErr(err error, t *testing.T)  {
	t.Helper()
	if err != nil {
	  t.Fatal(err)
	}
  
}

func TestServer(t *testing.T)  {

	// this should be my router
	r := GetRouter()
	server := httptest.NewServer(r)
	defer server.Close()


	resp, err := http.Get(server.URL + "/health")
	fatalErr(err, t)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatal("not status 200")
	}

	body, _ := io.ReadAll(resp.Body)

	if string(body) != "healthy" {
		t.Fatalf("Expected body 'healthy' got %s", body)
	}
}
