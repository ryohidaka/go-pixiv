package crawler

import (
	"testing"
	"time"
)

func TestGetSleepDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected time.Duration
	}{
		{
			name:     "No input (default duration)",
			input:    []int{},
			expected: 1000 * time.Millisecond,
		},
		{
			name:     "Custom duration 500ms",
			input:    []int{500},
			expected: 500 * time.Millisecond,
		},
		{
			name:     "Custom duration 2000ms",
			input:    []int{2000},
			expected: 2000 * time.Millisecond,
		},
		{
			name:     "Multiple values (only first used)",
			input:    []int{300, 1000, 5000},
			expected: 300 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSleepDuration(tt.input...)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
