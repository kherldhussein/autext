package autex

import (
	"reflect"
	"testing"
)

func TestUp(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"this", "(up)", "is", "a", "test"},
			expected: []string{"THIS", "is", "a", "test"},
		},
		{
			input:    []string{"hello", "(up)", "world"},
			expected: []string{"HELLO", "world"},
		},
		{
			input:    []string{"this", "is", "a", "test", "(up,", "2)"},
			expected: []string{"this", "is", "A", "TEST"},
		},
		{
			input:    []string{"first", "(up)", "word"},
			expected: []string{"FIRST", "word"},
		},
	}
	for _, tt := range tests {
		result := Up(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Expected %v, but got %v", tt.expected, result)
		}
	}
}
