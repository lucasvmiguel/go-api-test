package ping

import "net/http"

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("pong"))
}
