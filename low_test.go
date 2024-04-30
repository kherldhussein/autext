package autex

import (
	"reflect"
	"testing"
)

func TestLow(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"This", "(low)", "is", "a", "Test"},
			expected: []string{"this", "is", "a", "Test"},
		},
		{
			input:    []string{"Hello", "(low)", "WORLD"},
			expected: []string{"hello", "WORLD"},
		},
		{
			input:    []string{"This", "is", "A", "Test", "(low,", "2)"},
			expected: []string{"This", "is", "a", "test"},
		},
		{
			input:    []string{"First", "(low)", "word"},
			expected: []string{"first", "word"},
		},
	}
	for _, tt := range tests {
		result := Low(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Expected %v, but got %v", tt.expected, result)
		}
	}
}
