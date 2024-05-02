package autex

import (
	"reflect"
	"testing"
)

func TestVow(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"It", "was", "a", "apple"},
			expected: []string{"It", "was", "an", "apple"},
		},
		{
			input:    []string{"He", "has", "a", "house"},
			expected: []string{"He", "has", "an", "house"},
		},
		{
			input:    []string{"She", "ate", "an", "apple"},
			expected: []string{"She", "ate", "an", "apple"},
		},
		{
			input:    []string{"It", "was", "A", "hour"},
			expected: []string{"It", "was", "An", "hour"},
		},
	}
	for _, tt := range tests {
		result := Vow(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Expected %v, but got %v", tt.expected, result)
		}
	}
}
