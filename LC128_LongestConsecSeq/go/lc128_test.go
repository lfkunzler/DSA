package lc128

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		}, {
			name:     "Example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		}, {
			name:     "Example 3",
			nums:     []int{1, 0, 1, 2},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestConsecutive(tt.nums)
			if got != tt.expected {
				t.Errorf("longestConsecutive() for <%s> returned <%d>, wanted <%d>", tt.name, got, tt.expected)
			}
		})
	}
}
