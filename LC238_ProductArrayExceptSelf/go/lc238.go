package lc238

func productExceptSelf(nums []int) []int {
	len_nums := len(nums)
	res := make([]int, len_nums)

	prefix := 1
	for i := 0; i < len_nums; i++ {
		res[i] = prefix
		prefix *= nums[i]
	}

	sufix := 1
	for i := len_nums; i > 0; i-- {
		res[i-1] *= sufix
		sufix *= nums[i-1]
	}

	return res
}
