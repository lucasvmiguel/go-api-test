package todo

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/lucasvmiguel/go-api-test/pkg/http/client"
	"github.com/lucasvmiguel/go-api-test/pkg/http/respond"
)

var (
	EmptyBaseUrlError                = errors.New("base url cannot be empty or null")
	EmptyHTTPClientError             = errors.New("http client cannot be empty or null")
	FailedAPIStatusCodeResponseError = errors.New("api returned with a failed status code response")
)

type Handler struct {
	baseUrl    string
	httpClient client.HTTP
}

func NewHandler(httpClient client.HTTP, baseUrl string) (*Handler, error) {
	if baseUrl == "" {
		return nil, EmptyBaseUrlError
	}

	if httpClient == nil {
		return nil, EmptyHTTPClientError
	}

	return &Handler{
		baseUrl:    baseUrl,
		httpClient: httpClient,
	}, nil
}

func (h *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	request, _ := http.NewRequest(http.MethodGet, h.baseUrl, nil)
	resp, err := h.httpClient.Do(request)
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	if resp.StatusCode >= 400 {
		respond.HTTPError(w, http.StatusInternalServerError, FailedAPIStatusCodeResponseError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(body)
}
