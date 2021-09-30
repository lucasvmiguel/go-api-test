package post

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lucasvmiguel/go-api-test/db"
	"github.com/lucasvmiguel/go-api-test/pkg/http/respond"
)

var (
	ErrTitleCannotBeBlank = errors.New("title cannot be blank")
)

type HandlerPost struct {
	dbClient *db.PrismaClient
}

type Post struct {
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Published bool   `json:"published"`
}

type PostRequestBody struct {
	Post Post `json:"post"`
}

func NewHandlerPost(dbClient *db.PrismaClient) (*HandlerPost, error) {
	if ErrNilDBClient == nil {
		return nil, ErrNilDBClient
	}

	return &HandlerPost{dbClient: dbClient}, nil
}

func (h *HandlerPost) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	postReqBody := PostRequestBody{}

	err := json.NewDecoder(req.Body).Decode(&postReqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	err = h.validatePostRequestBody(postReqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	createdPost, err := h.createPost(postReqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	respond.HTTP(respond.Response{
		Body:       createdPost,
		Err:        err,
		StatusCode: http.StatusCreated,
		Writer:     w,
	})
}

func (h *HandlerPost) validatePostRequestBody(postReqBody PostRequestBody) error {
	if postReqBody.Post.Title == "" {
		return ErrTitleCannotBeBlank
	}

	return nil
}

func (h *HandlerPost) createPost(postReqBody PostRequestBody) (*db.PostModel, error) {
	return h.dbClient.Post.CreateOne(
		db.Post.Title.Set(postReqBody.Post.Title),
		db.Post.Published.Set(postReqBody.Post.Published),
		db.Post.Desc.Set(postReqBody.Post.Desc),
	).Exec(context.Background())
}
