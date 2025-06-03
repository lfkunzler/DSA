package lc217

// Given an integer array nums,
// return true if any value appears at least twice in the array
// return false if every element is distinct.
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, v := range nums {
		_, ok := seen[v]
		if ok {
			return true
		}
		seen[v] = true
	}
	return false
}
