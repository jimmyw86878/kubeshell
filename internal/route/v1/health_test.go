package v1_test

import (
	v1 "gatekeeper/internal/route/v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fail()
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(v1.Health)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("wrong status code: %v want %v", status, http.StatusNoContent)
	}
}
