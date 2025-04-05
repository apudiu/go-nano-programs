package tst

import (
	"io"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "5.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("Hello "), b...))

	w.WriteHeader(http.StatusOK)
}
