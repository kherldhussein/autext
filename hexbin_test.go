package autex

import (
	"reflect"
	"testing"
)

func TestHexBin(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"Simply", "add", "42", "(hex)", "and", "10", "(bin)", "and", "you", "will", "see", "the", "result", "is", "68."},
			expected: []string{"Simply", "add", "66", "and", "2", "and", "you", "will", "see", "the", "result", "is", "68."},
		},
		{
			input:    []string{"There", "are", "no", "hexadecimal", "or", "binary", "numbers", "here."},
			expected: []string{"There", "are", "no", "hexadecimal", "or", "binary", "numbers", "here."},
		},
	}
	for _, tt := range tests {
		result := Hexbin(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Expected %v, but got %v", tt.expected, result)
		}
	}
}
