package lc128

import "fmt"

func longestConsecutive(num []int) int {
	set := make(map[int]bool)

	for _, v := range num {
		set[v] = true
	}

	longest := 0
	for v := range set {
		if !set[v-1] { // it's the beggining of the sequence
			length := 1
			next := v + 1

			for set[next] {
				length += 1
				next += 1
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
