package post

import (
	"context"
	"net/http"

	"github.com/lucasvmiguel/go-api-test/pkg/http/respond"
)

func (h *Handler) HandleGet(w http.ResponseWriter, req *http.Request) {
	posts, err := h.dbClient.Post.FindMany().Exec(context.Background())
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	respond.HTTP(respond.Response{
		Body:       posts,
		Err:        err,
		StatusCode: http.StatusOK,
		Writer:     w,
	})
}
