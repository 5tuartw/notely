package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	test := http.Header{}
	testKey := "TestingWithAWorkingKey123"
	test.Set("Authorization", "ApiKey "+testKey)

	result, err := GetAPIKey(test)
	if err != nil {
		t.Fatalf("Didn't get token from header %v: %v", test, err)
	}

	if result != testKey {
		t.Fatalf("Extracted API Key %s does not match expected value %s.", result, testKey)
	}

}

func TestGetAPIKeyNoAuthHeader(t *testing.T) {
	test := http.Header{}

	_, err := GetAPIKey(test)

	if err == nil || !(err == ErrNoAuthHeaderIncluded) {
		t.Fatalf("Expected no auth header error, got %v", err)
	}
}

func TestGetAPIKeyMalformed(t *testing.T) {

	test := http.Header{}
	testKey := "TestingWithAWorkingKey123"
	test.Set("Authorization", testKey)

	result, err := GetAPIKey(test)
	if result != "" {
		t.Fatalf("Expected no API key, got %v", result)
	}

	if err == nil || !strings.Contains(err.Error(), "malformed authorization header") {
		t.Fatalf("Expected malformed authorization header error, got %v", err)
	}
}
