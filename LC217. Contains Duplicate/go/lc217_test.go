package lc217

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5, 1}, true},
		{[]int{1, 2, 3, 4, 5, 6}, false},
		{[]int{}, false},
		{[]int{1}, false},
	}

	for _, test := range tests {
		result := ContainsDuplicate(test.nums)
		if result != test.expected {
			t.Errorf("For nums: %v | expected: %t, received: %t", test.nums, test.expected, result)
		}
	}
}
