package webutils_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv"
	"github.com/ryohidaka/go-pixiv/internal/webutils"
)

// TestGetRestrict tests the GetRestrict function.
func TestGetRestrict(t *testing.T) {
	public := pixiv.Show
	private := pixiv.WebRestrict("hide")

	tests := []struct {
		name     string
		input    *pixiv.WebRestrict
		expected pixiv.WebRestrict
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: public,
		},
		{
			name:     "empty string",
			input:    func() *pixiv.WebRestrict { r := pixiv.WebRestrict(""); return &r }(),
			expected: public,
		},
		{
			name:     "valid restrict",
			input:    &private,
			expected: private,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := webutils.GetRestrict(tt.input)
			if result != tt.expected {
				t.Errorf("GetRestrict() = %v, want %v", result, tt.expected)
			}
		})
	}
}
