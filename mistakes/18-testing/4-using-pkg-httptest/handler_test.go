package tst

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Here we're testing server (a handler in the server)
// the goal is to generate a request & a writer to call the handler

func TestHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", strings.NewReader("There"))

	w := httptest.NewRecorder()

	Handler(w, req)

	if got := w.Result().Header.Get("X-API-VERSION"); got != "5.0" {
		t.Errorf("want X-API-VERSION=5.0 got %s", got)
	}

	body, _ := io.ReadAll(w.Result().Body)

	if string(body) != "Hello There" {
		t.Errorf("want 'Hello There' got %s", string(body))
	}

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("want status 200 OK got %d", w.Result().StatusCode)
	}
}
