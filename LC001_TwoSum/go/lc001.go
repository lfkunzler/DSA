package lc001

func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)

	for i, v := range nums {
		j, ok := lookup[target-v]
		if ok {
			return []int{i, j}
		}
		lookup[v] = i
	}

	return []int{}
}
