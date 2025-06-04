package lc049

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeGroups(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}

	sort.Slice(groups, func(i, j int) bool {
		groupA := groups[i]
		groupB := groups[j]

		minLen := len(groupA)
		if len(groupB) < minLen {
			minLen = len(groupB)
		}

		for k := 0; k < minLen; k++ {
			if groupA[k] < groupB[k] {
				return true
			}
			if groupA[k] > groupB[k] {
				return false
			}
		}
		return len(groupA) < len(groupB)
	})
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name:     "Example 1: Basic Anagrams",
			strs:     []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name:     "Example 2: Empty String",
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "Example 3: Single Character",
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
		{
			name:     "No Anagrams",
			strs:     []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:     "Mixed Anagrams and Non-Anagrams",
			strs:     []string{"listen", "silent", "go", "dog", "god", "hello"},
			expected: [][]string{{"dog", "god"}, {"go"}, {"hello"}, {"listen", "silent"}},
		},
		{
			name:     "Strings with Duplicate Characters",
			strs:     []string{"aabb", "bbaa", "baba", "cccc"},
			expected: [][]string{{"aabb", "baba", "bbaa"}, {"cccc"}},
		},
		{
			name:     "All Same String",
			strs:     []string{"abc", "abc", "abc"},
			expected: [][]string{{"abc", "abc", "abc"}},
		},
		{
			name:     "Empty Input Slice",
			strs:     []string{},
			expected: [][]string{},
		},
		{
			name:     "Longer Strings",
			strs:     []string{"topcoder", "redocpot", "coderpot", "programming"},
			expected: [][]string{{"coderpot", "redocpot", "topcoder"}, {"programming"}},
		},
		{
			name:     "Single element groups",
			strs:     []string{"a", "b", "c"},
			expected: [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name:     "Strings with different character counts",
			strs:     []string{"aa", "aaa"},
			expected: [][]string{{"aa"}, {"aaa"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := groupAnagrams(tt.strs)

			normalizeGroups(actual)
			normalizeGroups(tt.expected)

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("groupAnagrams(%v)\nGot:    %v\nWanted: %v", tt.strs, actual, tt.expected)
			}
		})
	}
}
