package post

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucasvmiguel/go-api-test/db"
	"github.com/lucasvmiguel/go-api-test/pkg/test"
)

func TestHandlerGet(t *testing.T) {
	dbClient, mock, ensure := db.NewMock()
	defer ensure(t)

	expected := db.PostModel{
		InnerPost: db.InnerPost{
			ID:        "123",
			Title:     "some title",
			Published: true,
		},
	}

	mock.Post.Expect(
		dbClient.Post.FindMany(),
	).ReturnsMany([]db.PostModel{expected})

	body, _ := json.Marshal(requestBody)
	handler, _ := NewHandlerGet(dbClient)
	res, resBody := test.HttpRequest(handler, bytes.NewReader(body))
	defer res.Body.Close()

	statusCodeExpected := http.StatusOK
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	expectedAPIResponse := fmt.Sprintf(
		"[{\"id\":\"%s\",\"createdAt\":\"0001-01-01T00:00:00Z\",\"updatedAt\":\"0001-01-01T00:00:00Z\",\"title\":\"%s\",\"published\":%t,\"desc\":null}]",
		expected.InnerPost.ID, expected.InnerPost.Title, expected.InnerPost.Published)
	if resBody != expectedAPIResponse {
		spew.Dump(resBody)
		spew.Dump(expectedAPIResponse)
		t.Errorf("expected %s got %v", expectedAPIResponse, resBody)
	}
}

func TestHandlerGetDBError(t *testing.T) {
	dbClient, mock, ensure := db.NewMock()
	defer ensure(t)

	mock.Post.Expect(
		dbClient.Post.FindMany(),
	).Errors(db.ErrNotFound)

	body, _ := json.Marshal(requestBody)
	handler, _ := NewHandlerGet(dbClient)
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
