package context1

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// ----------------------------------------- mock

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}

// ----------------------------------------- mock

func TestServer(t *testing.T) {
	data := "hello world"
	svr := Server(&SpyStore{response: data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf("got %s and expected %s", response.Body.String(), data)
	}
}
