package lc347

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	for _, v := range nums { // time: O(n), space: O(m)
		count[v]++
	}

	freq := make([][]int, len(nums)) // the freq bucket arrays has the length of nums
	// because we can have all elements the same.
	for num, cnt := range count { // time: O(m), space O(n)
		freq[cnt-1] = append(freq[cnt-1], num)
	}

	res := make([]int, 0, k)
	for i := len(nums) - 1; i >= 0; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
