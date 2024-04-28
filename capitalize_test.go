package autex

import (
	"reflect"
	"testing"
)

func TestWordZ(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"it", "(cap)", "was", "the", "best", "of", "times"},
			expected: []string{"It", "was", "the", "best", "of", "times"},
		},
		{
			input:    []string{"hello", "(cap)", "world"},
			expected: []string{"Hello", "world"},
		},
		{
			input:    []string{"this", "is", "a", "test", "(cap,", "2)"},
			expected: []string{"this", "is", "A", "Test"},
		},
		{
			input:    []string{"first", "(cap)", "word"},
			expected: []string{"First", "word"},
		},
	}
	for _, tt := range tests {
		result := WordZ(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Expected %v, but got %v", tt.expected, result)
		}
	}
}
