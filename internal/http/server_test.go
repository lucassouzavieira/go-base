package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helpers functions
func makeRequest(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	initRouter().ServeHTTP(res, req)
	return res
}

// Tests
func TestListUsers(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/users", nil)
	response := makeRequest(req)

	if response.Code != 200 {
		t.Error("Expecting HTTP 200. Got %i", response.Code)
	}
}
