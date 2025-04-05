package tst

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Here we're testing a client (request)
// the goal is to create a mock server to handle the request

func TestDurationClientGet(t *testing.T) {
	// prepare server

	srv := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"duration": 120}`))
			},
		),
	)
	defer srv.Close()

	// prepare client

	client := NewDurationClient()

	duration, err := client.GetDuration(
		srv.URL, 51.551261, -0.1221146, 51.57, -0.13,
	)
	if err != nil {
		t.Fatal(err)
	}

	if duration != 120*time.Second {
		t.Errorf("duration expected %v, got %v", 120*time.Second, duration)
	}
}
