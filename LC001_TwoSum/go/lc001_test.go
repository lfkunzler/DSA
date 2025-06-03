package lc001

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1: Basic positive numbers",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "Example 2: Different order",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "Example 3: Numbers with zero",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "Mixed positive and negative numbers",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2},
		},
		{
			name:     "Large numbers",
			nums:     []int{100, 200, 300, 400},
			target:   700,
			expected: []int{2, 3},
		},
		{
			name:     "Array with two elements",
			nums:     []int{5, 5},
			target:   10,
			expected: []int{0, 1},
		},
		{
			name:     "Target sum from middle elements",
			nums:     []int{1, 2, 8, 9, 10},
			target:   17,
			expected: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := twoSum(tt.nums, tt.target)

			sort.Ints(result)
			sort.Ints(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("twoSum(%v, %d) = %v; expected %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
