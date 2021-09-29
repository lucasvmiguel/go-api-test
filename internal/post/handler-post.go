package post

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lucasvmiguel/go-api-test/db"
	"github.com/lucasvmiguel/go-api-test/pkg/http/respond"
)

type Post struct {
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	Published bool   `json:"published"`
}

type PostRequestBody struct {
	Post Post `json:"post"`
}

func (h *Handler) HandlePost(w http.ResponseWriter, req *http.Request) {
	postReqBody := PostRequestBody{}

	err := json.NewDecoder(req.Body).Decode(&postReqBody)
	if err != nil {
		respond.HTTPError(w, http.StatusBadRequest, err)
		return
	}

	createdPost, err := h.dbClient.Post.CreateOne(
		db.Post.Title.Set(postReqBody.Post.Title),
		db.Post.Published.Set(postReqBody.Post.Published),
		db.Post.Desc.Set(postReqBody.Post.Desc),
	).Exec(context.Background())
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
