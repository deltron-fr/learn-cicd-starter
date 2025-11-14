package auth

import (
	"net/http"
	"testing"
)


func TestGetApiKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)

	req.Header.Set("Authorization", "ApiKey RAYY")
	got, _ := GetAPIKey(req.Header)
	want := "RAYY"
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}