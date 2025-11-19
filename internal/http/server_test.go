package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MagnumTrader/repforge/internal/domain"
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


func TestWorkoutsPage(t *testing.T) {

	workouts := []domain.Workout{}

	// Add a workout
	w := domain.Workout{
		Date:     "2025-11-11",
		Kind:     "Swimming",
		Duration: 50,
		Notes:    "Pool session",
	}
	workouts = append(workouts, w)

	// Start the router as a test server
	r := GetRouter() // production router
	server := httptest.NewServer(r)
	defer server.Close()

	resp, err := http.Get(server.URL + "/workouts")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	html := string(body)

	// Verify workout appears in rendered HTML
	if !strings.Contains(html, w.Date) ||
		!strings.Contains(html, w.Kind) ||
		!strings.Contains(html, "50") ||
		!strings.Contains(html, w.Notes) {
		t.Fatalf("Workout not found in rendered HTML: %v\nHTML:\n%s", w, html)
	}
}
