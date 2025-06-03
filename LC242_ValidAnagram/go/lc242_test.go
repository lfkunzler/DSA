package lc242

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"listen", "silent", true},
		{"hello", "olleh", true},
		{"a", "a", true},
		{"", "", true},

		{"abc", "ab", false},
		{"ab", "abc", false},

		{"aabb", "baba", true},
		{"aabbc", "babba", false},
		{"aabb", "aaab", false},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("s=\"%s\", t=\"%s\"", test.s, test.t)
		t.Run(testName, func(t *testing.T) {
			result := isAnagram(test.s, test.t)
			if result != test.expected {
				t.Errorf("isAnagram(%q, %q) = %t; expected %t", test.s, test.t, result, test.expected)
			}
		})
	}
}
