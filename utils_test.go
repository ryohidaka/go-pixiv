package pixiv

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// TestGenClientHash tests the genClientHash function.
func TestGenClientHash(t *testing.T) {
	clientTime := "2025-05-01T12:00:00Z"

	// Manually calculate the expected hash value using the known secret
	// ClientHashSecret is assumed to be "yourSecretValue"
	expectedSecret := ClientHashSecret // This will be imported from constants.go
	h := md5.New()
	io.WriteString(h, clientTime)
	io.WriteString(h, expectedSecret)
	expectedHash := hex.EncodeToString(h.Sum(nil)) // Expected MD5 hash

	// Call the function
	result := genClientHash(clientTime)

	// Check if the result matches the expected value
	if result != expectedHash {
		t.Errorf("expected %s, got %s", expectedHash, result)
	}
}

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

	setHeaders(req, headers)

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

	body, err := readResponse(resp)
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
	err := decodeJSON(body, &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check if the result is as expected
	expected := map[string]string{"status": "ok"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// TestGetExpiresAt tests the getExpiresAt function.
func TestGetExpiresAt(t *testing.T) {
	expiredIn := 3600 // 1 hour
	expirationTime := getExpiresAt(expiredIn)

	// Check if the expiration time is approximately 1 hour from now
	if expirationTime.Before(time.Now().Add(time.Duration(expiredIn-5)*time.Second)) || expirationTime.After(time.Now().Add(time.Duration(expiredIn+5)*time.Second)) {
		t.Errorf("expected expiration time to be within 1 hour, got %v", expirationTime)
	}
}

// TestParseNextPageOffset tests the parseNextPageOffset function.
func TestParseNextPageOffset(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		field      string
		wantOffset int
		wantErr    bool
	}{
		{
			name:       "valid offset",
			url:        "https://example.com/api?page=2&offset=100",
			field:      "offset",
			wantOffset: 100,
			wantErr:    false,
		},
		{
			name:       "missing offset param",
			url:        "https://example.com/api?page=2",
			field:      "offset",
			wantOffset: 0,
			wantErr:    true,
		},
		{
			name:       "invalid offset value",
			url:        "https://example.com/api?offset=abc",
			field:      "offset",
			wantOffset: 0,
			wantErr:    true,
		},
		{
			name:       "empty url",
			url:        "",
			field:      "offset",
			wantOffset: 0,
			wantErr:    false,
		},
		{
			name:       "invalid url",
			url:        "%%%",
			field:      "offset",
			wantOffset: 0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseNextPageOffset(tt.url, tt.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseNextPageOffset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantOffset {
				t.Errorf("parseNextPageOffset() = %v, want %v", got, tt.wantOffset)
			}
		})
	}
}
