package utils

import (
	"testing"
)

func TestReverse(t *testing.T) {
	// Test cases for different types and scenarios
	tests := []struct {
		name     string
		input    []int // 使用具体类型，这里以 []int 为例
		expected []int
	}{
		{
			name:     "Reverse slice of int",
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "Reverse empty slice of int",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Reverse single element slice",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.input) // 直接操作和测试具体类型的切片
			if !equal(tt.input, tt.expected) {
				t.Errorf("Reverse() = %v, expected %v", tt.input, tt.expected)
			}
		})
	}
}

// Helper function to compare slices of ints
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
