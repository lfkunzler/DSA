package lc238

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Basic Case (Positive Numbers)",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Single Zero",
			nums:     []int{1, 2, 0, 4},
			expected: []int{0, 0, 8, 0},
		},
		{
			name:     "Multiple Zeros",
			nums:     []int{1, 0, 3, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Even Number of Negative Numbers",
			nums:     []int{-1, 2, -3, 4},
			expected: []int{-24, 12, -8, 6},
		},
		{
			name:     "Odd Number of Negative Numbers",
			nums:     []int{-1, -2, -3, 4},
			expected: []int{24, 12, 8, -6},
		},
		{
			name:     "Mixed Positive, Negative, and Zero",
			nums:     []int{-1, 0, 3, -4},
			expected: []int{0, 12, 0, 0},
		},
		{
			name:     "Smallest Array Length (Positive)",
			nums:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Smallest Array Length (Negative)",
			nums:     []int{-5, 10},
			expected: []int{10, -5},
		},
		{
			name:     "Largest Values (Near Constraint Limits)",
			nums:     []int{30, 30, 30},
			expected: []int{900, 900, 900},
		},
		{
			name:     "Smallest Values (Near Constraint Limits)",
			nums:     []int{-30, -30, -30},
			expected: []int{900, 900, 900},
		},
		{
			name:     "All Ones",
			nums:     []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			name:     "All Negative Ones",
			nums:     []int{-1, -1, -1, -1},
			expected: []int{-1, -1, -1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("productExceptSelf(%v) = %v; want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
