package urlutil_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv/internal/urlutil"
)

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
			got, err := urlutil.ParseNextPageOffset(tt.url, tt.field)
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
