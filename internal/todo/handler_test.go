package todo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/lucasvmiguel/go-api-test/pkg/test"
)

var baseUrl = "http://test.com"

var apiSuccessResponse = `[{"userId": 1,	"id": 1, "title": "delectus aut autem", "completed": false}]`

type fetcherSuccessMock struct{}

func (f *fetcherSuccessMock) Do(*http.Request) (*http.Response, error) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(apiSuccessResponse)))

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       r,
	}, nil
}

type fetcherErrorMock struct{}

func (f *fetcherErrorMock) Do(*http.Request) (*http.Response, error) {
	r := ioutil.NopCloser(bytes.NewReader([]byte("")))

	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       r,
	}, nil
}

func TestSuccessTodoHandler(t *testing.T) {
	h, _ := NewHandler(&fetcherSuccessMock{}, baseUrl)
	res, resBody := test.HttpRequest(h, nil)
	defer res.Body.Close()

	statusCodeExpected := http.StatusOK
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	if resBody != apiSuccessResponse {
		t.Errorf("expected %s got %v", apiSuccessResponse, resBody)
	}
}

func TestErrorTodoHandler(t *testing.T) {
	h, _ := NewHandler(&fetcherErrorMock{}, baseUrl)
	res, _ := test.HttpRequest(h, nil)
	defer res.Body.Close()

	statusCodeExpected := http.StatusInternalServerError
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}
}
