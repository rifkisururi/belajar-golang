package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name, target, wantBody string
		wantStatus             int
	}{
		{"health", "/healthz", "ok\n", http.StatusOK},
		{"greeting", "/greet?name=Rifki", "Hello, Rifki!\n", http.StatusOK},
		{"default greeting", "/greet", "Hello, Gopher!\n", http.StatusOK},
		{"missing route", "/missing", "404 page not found\n", http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := httptest.NewRecorder()
			New().ServeHTTP(res, httptest.NewRequest(http.MethodGet, tt.target, nil))
			if res.Code != tt.wantStatus {
				t.Fatalf("status = %d; want %d", res.Code, tt.wantStatus)
			}
			if res.Body.String() != tt.wantBody {
				t.Errorf("body = %q; want %q", res.Body.String(), tt.wantBody)
			}
		})
	}
}
