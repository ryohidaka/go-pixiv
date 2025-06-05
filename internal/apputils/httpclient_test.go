package apputils_test

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/ryohidaka/go-pixiv/internal/apputils"
)

// TestSetHeaders tests the setHeaders function by checking if headers are correctly set.
func TestSetHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "PixivClient/1.0",
	}

	apputils.SetHeaders(req, headers)

	// Check if headers were set correctly
	for key, value := range headers {
		if req.Header.Get(key) != value {
			t.Errorf("expected header %s: %s, got %s", key, value, req.Header.Get(key))
		}
	}
}

// TestReadResponse tests the readResponse function.
func TestReadResponse(t *testing.T) {
	// Simulate an HTTP response
	resp := &http.Response{
		Body: io.NopCloser(bytes.NewReader([]byte(`{"status":"ok"}`))),
	}

	body, err := apputils.ReadResponse(resp)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedBody := []byte(`{"status":"ok"}`)
	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("expected body %v, got %v", expectedBody, body)
	}
}

// TestDecodeJSON tests the decodeJSON function.
func TestDecodeJSON(t *testing.T) {
	body := []byte(`{"status":"ok"}`)
	var result map[string]string

	// Call the function
	err := apputils.DecodeJSON(body, &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check if the result is as expected
	expected := map[string]string{"status": "ok"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
