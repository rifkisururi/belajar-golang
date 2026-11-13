package task

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes(t *testing.T) {
	h := NewHandler(NewService(NewMemoryRepository())).Routes()
	post := httptest.NewRecorder()
	h.ServeHTTP(post, httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(`{"title":"learn Go"}`)))
	if post.Code != http.StatusCreated {
		t.Fatalf("POST status = %d", post.Code)
	}
	get := httptest.NewRecorder()
	h.ServeHTTP(get, httptest.NewRequest(http.MethodGet, "/tasks", nil))
	if get.Code != http.StatusOK || !bytes.Contains(get.Body.Bytes(), []byte("learn Go")) {
		t.Fatalf("GET = %d %s", get.Code, get.Body.String())
	}
	bad := httptest.NewRecorder()
	h.ServeHTTP(bad, httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(`{"title":" "}`)))
	if bad.Code != http.StatusBadRequest {
		t.Fatalf("invalid POST status = %d", bad.Code)
	}
}
