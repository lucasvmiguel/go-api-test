package todo

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var baseUrl = "http://test.com"

type fetcherSuccessMock struct{}

var apiSuccessResponse = `[{"userId": 1,	"id": 1, "title": "delectus aut autem", "completed": false}]`

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
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h, _ := NewHandler(&fetcherSuccessMock{}, baseUrl)
	h.Handle(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	statusCodeExpected := http.StatusOK
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	if string(data) != apiSuccessResponse {
		t.Errorf("expected %s got %v", apiSuccessResponse, string(data))
	}
}

func TestErrorTodoHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h, _ := NewHandler(&fetcherErrorMock{}, baseUrl)
	h.Handle(w, req)
	res := w.Result()
	defer res.Body.Close()

	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	statusCodeExpected := http.StatusInternalServerError
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}
}
