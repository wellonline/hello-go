package handler_test

import (
	"hello-go/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Error(err)
		return
	}

	resp := httptest.NewRecorder()
	h := handler.New()
	h.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusOK {
		t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
	}
}
