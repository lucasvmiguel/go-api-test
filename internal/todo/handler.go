package todo

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/lucasvmiguel/go-api-test/pkg/http/client"
	"github.com/lucasvmiguel/go-api-test/pkg/http/respond"
)

var (
	ErrEmptyBaseUrl                = errors.New("base url cannot be empty or null")
	ErrEmptyHTTPClient             = errors.New("http client cannot be empty or null")
	ErrFailedAPIStatusCodeResponse = errors.New("api returned with a failed status code response")
)

type Handler struct {
	baseUrl    string
	httpClient client.HTTP
}

func NewHandler(httpClient client.HTTP, baseUrl string) (*Handler, error) {
	if baseUrl == "" {
		return nil, ErrEmptyBaseUrl
	}

	if httpClient == nil {
		return nil, ErrEmptyHTTPClient
	}

	return &Handler{
		baseUrl:    baseUrl,
		httpClient: httpClient,
	}, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	request, _ := http.NewRequest(http.MethodGet, h.baseUrl, nil)
	resp, err := h.httpClient.Do(request)
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	if resp.StatusCode >= 400 {
		respond.HTTPError(w, http.StatusInternalServerError, ErrFailedAPIStatusCodeResponse)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		respond.HTTPError(w, http.StatusInternalServerError, err)
		return
	}

	w.Write(body)
}
