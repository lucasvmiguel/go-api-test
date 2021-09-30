package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/lucasvmiguel/go-api-test/db"
	"github.com/lucasvmiguel/go-api-test/pkg/test"
)

var postID = "123"
var requestBody = PostRequestBody{
	Post: Post{
		Title:     "testing",
		Desc:      "testing desc",
		Published: true,
	},
}

func TestCreatePost(t *testing.T) {
	dbClient, mock, ensure := db.NewMock()
	defer ensure(t)

	expected := db.PostModel{
		InnerPost: db.InnerPost{
			ID:        postID,
			Title:     requestBody.Post.Title,
			Desc:      &requestBody.Post.Desc,
			Published: requestBody.Post.Published,
		},
	}

	mock.Post.Expect(
		dbClient.Post.CreateOne(
			db.Post.Title.Set(requestBody.Post.Title),
			db.Post.Published.Set(requestBody.Post.Published),
			db.Post.Desc.Set(requestBody.Post.Desc),
		),
	).Returns(expected)

	handler, _ := NewHandlerPost(dbClient)

	handler.createPost(requestBody)
}

func TestValidatePostRequestBody(t *testing.T) {
	tests := []struct {
		message  string
		param    PostRequestBody
		expected error
	}{
		{
			message:  "valid post request body",
			param:    PostRequestBody{Post: Post{Title: "some title"}},
			expected: nil,
		},
		{
			message:  "invalid post request body - title cannot be blank",
			param:    PostRequestBody{Post: Post{Title: ""}},
			expected: ErrTitleCannotBeBlank,
		},
	}

	dbClient, _, _ := db.NewMock()
	handler, _ := NewHandlerPost(dbClient)

	for _, test := range tests {
		result := handler.validatePostRequestBody(test.param)

		if result != test.expected {
			t.Errorf("expected %v got %v", result, test.expected)
		}
	}
}

func TestHandlerPost(t *testing.T) {
	dbClient, mock, ensure := db.NewMock()
	defer ensure(t)

	expected := db.PostModel{
		InnerPost: db.InnerPost{
			ID:        postID,
			Title:     requestBody.Post.Title,
			Desc:      &requestBody.Post.Desc,
			Published: requestBody.Post.Published,
		},
	}

	mock.Post.Expect(
		dbClient.Post.CreateOne(
			db.Post.Title.Set(requestBody.Post.Title),
			db.Post.Published.Set(requestBody.Post.Published),
			db.Post.Desc.Set(requestBody.Post.Desc),
		),
	).Returns(expected)

	body, _ := json.Marshal(requestBody)
	handler, _ := NewHandlerPost(dbClient)
	res, resBody := test.HttpRequest(handler, bytes.NewReader(body))
	defer res.Body.Close()

	statusCodeExpected := http.StatusCreated
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	expectedAPIResponse := `{"id":"123","createdAt":"0001-01-01T00:00:00Z","updatedAt":"0001-01-01T00:00:00Z","title":"testing","published":true,"desc":"testing desc"}`
	if resBody != expectedAPIResponse {
		t.Errorf("expected %s got %v", expectedAPIResponse, resBody)
	}
}

func TestHandlerPostDBError(t *testing.T) {
	dbClient, mock, ensure := db.NewMock()
	defer ensure(t)

	mock.Post.Expect(
		dbClient.Post.CreateOne(
			db.Post.Title.Set(requestBody.Post.Title),
			db.Post.Published.Set(requestBody.Post.Published),
			db.Post.Desc.Set(requestBody.Post.Desc),
		),
	).Errors(db.ErrNotFound)

	body, _ := json.Marshal(requestBody)
	handler, _ := NewHandlerPost(dbClient)
	res, resBody := test.HttpRequest(handler, bytes.NewReader(body))
	defer res.Body.Close()

	statusCodeExpected := http.StatusInternalServerError
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	expectedAPIResponse := fmt.Sprintf("{\"message\":\"%s\"}", db.ErrNotFound)
	if resBody != expectedAPIResponse {
		t.Errorf("expected %s got %v", expectedAPIResponse, resBody)
	}
}

func TestHandlerPostRequestError(t *testing.T) {
	dbClient, _, _ := db.NewMock()

	handler, _ := NewHandlerPost(dbClient)
	res, resBody := test.HttpRequest(handler, strings.NewReader("invalid json"))
	defer res.Body.Close()

	statusCodeExpected := http.StatusBadRequest
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	expectedAPIResponse := "{\"message\":\"invalid character 'i' looking for beginning of value\"}"
	if resBody != expectedAPIResponse {
		t.Errorf("expected %s got %v", expectedAPIResponse, resBody)
	}
}

func TestHandlerPostRequestValidateError(t *testing.T) {
	dbClient, _, _ := db.NewMock()

	handler, _ := NewHandlerPost(dbClient)
	res, resBody := test.HttpRequest(handler, strings.NewReader("{}"))
	defer res.Body.Close()

	statusCodeExpected := http.StatusBadRequest
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	expectedAPIResponse := fmt.Sprintf("{\"message\":\"%s\"}", ErrTitleCannotBeBlank)
	if resBody != expectedAPIResponse {
		t.Errorf("expected %s got %v", expectedAPIResponse, resBody)
	}
}
