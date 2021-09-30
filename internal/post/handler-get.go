package post

import (
	"context"
	"net/http"

	"github.com/lucasvmiguel/go-api-test/db"
	"github.com/lucasvmiguel/go-api-test/pkg/http/respond"
)

type HandlerGet struct {
	dbClient *db.PrismaClient
}

func NewHandlerGet(dbClient *db.PrismaClient) (*HandlerGet, error) {
	if ErrNilDBClient == nil {
		return nil, ErrNilDBClient
	}

	return &HandlerGet{dbClient: dbClient}, nil
}

func (h *HandlerGet) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
