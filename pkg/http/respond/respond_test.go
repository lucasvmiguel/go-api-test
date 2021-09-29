package respond

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var message = "testing"
var err = errors.New("testing")

type JSONBody struct {
	Message string
}

type HandlerSuccess struct{}

func (h *HandlerSuccess) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	HTTP(Response{
		Body:       JSONBody{Message: message},
		StatusCode: http.StatusCreated,
		Writer:     w,
	})
}

type HandlerError struct{}

func (h *HandlerError) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	HTTP(Response{
		Body:       JSONBody{Message: message},
		StatusCode: http.StatusBadRequest,
		Writer:     w,
		Err:        err,
	})
}

func TestHTTP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h := HandlerSuccess{}
	h.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	statusCodeExpected := http.StatusCreated
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	jsonBody := JSONBody{}
	json.Unmarshal(data, &jsonBody)
	expectedBody := JSONBody{Message: message}
	if !reflect.DeepEqual(jsonBody, expectedBody) {
		t.Errorf("expected %s got %s", expectedBody, jsonBody)
	}

	expectedHeader := "application/json"
	if res.Header.Get("Content-Type") != expectedHeader {
		t.Errorf("expected %s got %s", expectedHeader, res.Header.Get("Content-Type"))
	}
}

func TestHTTPError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h := HandlerError{}
	h.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	statusCodeExpected := http.StatusBadRequest
	if res.StatusCode != statusCodeExpected {
		t.Errorf("expected %d got %d", statusCodeExpected, res.StatusCode)
	}

	jsonBody := JSONBody{}
	json.Unmarshal(data, &jsonBody)
	expectedBody := JSONBody{Message: message}
	if !reflect.DeepEqual(jsonBody, expectedBody) {
		t.Errorf("expected %s got %s", expectedBody, jsonBody)
	}

	expectedHeader := "application/json"
	if res.Header.Get("Content-Type") != expectedHeader {
		t.Errorf("expected %s got %s", expectedHeader, res.Header.Get("Content-Type"))
	}
}
