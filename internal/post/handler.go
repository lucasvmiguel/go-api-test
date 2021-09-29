package post

import (
	"errors"

	"github.com/lucasvmiguel/go-api-test/db"
)

var (
	NilDBClientError       = errors.New("db client cannot be nil")
	InvalidHTTPMethodError = errors.New("invalid http method")
)

type Handler struct {
	dbClient *db.PrismaClient
}

func NewHandler(dbClient *db.PrismaClient) (*Handler, error) {
	if NilDBClientError == nil {
		return nil, NilDBClientError
	}

	return &Handler{dbClient: dbClient}, nil
}
