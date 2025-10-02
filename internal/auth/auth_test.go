package auth

import (
	"errors"
	"net/http"
	"testing"
)

// no header
func TestGetApi_NoHeader(t *testing.T) {
	headers := http.Header{}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("expected erorr, got %v", err)
	}
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Errorf("expected %v , got %v", ErrNoAuthHeaderIncluded, err)
	}
}

//malformed header

func TestGetApi_malformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer token123")
	_, err := GetAPIKey(headers)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("expected malformed authorization header, got %v", err)
	}

}

//invalid header

func TestGetApi_valid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey token123")
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("expected no error got %v", err)
	}
	if key != "token123" {

		t.Errorf("expected key 'token123' got %v", key)
	}
}
