package coreutils_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv/internal/coreutils"
	"github.com/stretchr/testify/assert"
)

// dummy query struct for testing
type queryParams struct {
	ID    int    `url:"id"`
	Query string `url:"query"`
}

func TestBuildRequestURL_Success(t *testing.T) {
	base := "https://example.com/api"
	params := queryParams{ID: 42, Query: "pixiv"}

	got, err := coreutils.BuildRequestURL(base, params)

	assert.NoError(t, err)
	assert.NotNil(t, got)

	expectedURL := "https://example.com/api?id=42&query=pixiv"
	assert.Equal(t, expectedURL, got.String())
}

func TestBuildRequestURL_NoQueryParams(t *testing.T) {
	base := "https://example.com/api"

	got, err := coreutils.BuildRequestURL(base, nil)

	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, base, got.String())
}

func TestBuildRequestURL_InvalidBaseURL(t *testing.T) {
	base := "://bad-url"

	got, err := coreutils.BuildRequestURL(base, nil)

	assert.Error(t, err)
	assert.Nil(t, got)
}
