package lc049

func groupAnagrams(strs []string) [][]string {
	smap := make(map[[26]int][]string, len(strs))

	for _, s := range strs {
		freq := [26]int{}
		for _, r := range s {
			if r-'a' >= 0 && r-'a' < 26 {
				freq[r-'a']++
			}
		}
		smap[freq] = append(smap[freq], s)
	}

	result := make([][]string, 0, len(smap))
	for _, strSlice := range smap {
		result = append(result, strSlice)
	}

	return result
}
