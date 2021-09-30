package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

func HttpRequest(handler http.Handler, body io.Reader) (*http.Response, string) {
	req := httptest.NewRequest(http.MethodGet, "/", body)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)
	return w.Result(), strings.TrimSuffix(w.Body.String(), "\n")
}
