package lc347

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeIntSlice(slice []int) {
	sort.Ints(slice)
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Basic case",
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Single element array",
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "Multiple elements, one k",
			nums:     []int{1, 2, 3, 1, 2, 1},
			k:        1,
			expected: []int{1},
		},
		{
			name:     "All unique elements, k = length",
			nums:     []int{10, 20, 30, 40, 50},
			k:        5,
			expected: []int{10, 20, 30, 40, 50},
		},
		{
			name:     "Mixed frequencies",
			nums:     []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4},
			k:        3,
			expected: []int{1, 2, 4},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -1, -1, -2, -2, -3},
			k:        2,
			expected: []int{-1, -2},
		},
		{
			name:     "Large numbers and k",
			nums:     []int{10000, 10000, 10000, 1, 1, 2},
			k:        2,
			expected: []int{1, 10000},
		},
		{
			name:     "More elements than unique values",
			nums:     []int{1, 1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 4, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Mix of small and large numbers",
			nums:     []int{-100, 5, 5, 5, -100, 1, 1, 1, 1, 1},
			k:        2,
			expected: []int{1, 5},
		},
		{
			name:     "failed",
			nums:     []int{5, 2, 5, 3, 5, 3, 1, 1, 3},
			k:        2,
			expected: []int{3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := topKFrequent(tt.nums, tt.k)

			normalizeIntSlice(actual)
			normalizeIntSlice(tt.expected)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("topKFrequent(%v, %d)\nGot:    %v\nWanted: %v", tt.nums, tt.k, actual, tt.expected)
			}
		})
	}
}
