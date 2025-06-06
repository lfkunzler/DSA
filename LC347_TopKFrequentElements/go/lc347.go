package lc347

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	topE := make([]int, k)
	topF := make([]int, k)
	for s, v := range freq {
		for count := 0; count < k; count++ {
			if topF[count] < v {
				topF[count] = v
				topE[count] = s
				break
			}
		}
	}

	return topE
}
