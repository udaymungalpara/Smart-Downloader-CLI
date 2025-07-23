package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRetrySucceedsAfterFailures(t *testing.T) {
	attempts := 0

	// Set up a mock server that fails twice, then succeeds
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			http.Error(w, "temporary failure", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}))
	defer server.Close()

	resp, err := retry(server.URL, 10*time.Millisecond, 5)

	if err != nil {
		t.Fatalf("Expected successful retry, got error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}
