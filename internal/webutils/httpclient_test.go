package webutils_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/ryohidaka/go-pixiv/internal/webutils"
	"github.com/stretchr/testify/assert"
)

// mockReader implements io.Reader and returns an error for testing.
type mockReader struct{}

func (m *mockReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

func TestDecodeJSON_Success(t *testing.T) {
	type Payload struct {
		Message string `json:"message"`
	}

	jsonBody := `{"message": "hello"}`
	reader := strings.NewReader(jsonBody)

	result, err := webutils.DecodeJSON[Payload](reader)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "hello", result.Message)
}

func TestDecodeJSON_ReadError(t *testing.T) {
	type Dummy struct{}

	_, err := webutils.DecodeJSON[Dummy](&mockReader{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read response body")
}

func TestDecodeJSON_UnmarshalError(t *testing.T) {
	type Payload struct {
		Count int `json:"count"`
	}

	invalidJSON := `{"count": "not-an-int"}`
	reader := bytes.NewReader([]byte(invalidJSON))

	_, err := webutils.DecodeJSON[Payload](reader)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal response body")
}
