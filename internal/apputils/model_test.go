package apputils_test

import (
	"testing"

	"github.com/ryohidaka/go-pixiv/internal/apputils"
	"github.com/ryohidaka/go-pixiv/models"
)

// TestGetRestrict tests the GetRestrict function.
func TestGetRestrict(t *testing.T) {
	public := models.Public
	private := models.Restrict("private")

	tests := []struct {
		name     string
		input    *models.Restrict
		expected models.Restrict
	}{
		{
			name:     "nil input",
			input:    nil,
			expected: public,
		},
		{
			name:     "empty string",
			input:    func() *models.Restrict { r := models.Restrict(""); return &r }(),
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
			result := apputils.GetRestrict(tt.input)
			if result != tt.expected {
				t.Errorf("GetRestrict() = %v, want %v", result, tt.expected)
			}
		})
	}
}
