package lc242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqS := [26]int{}
	freqT := [26]int{}

	for i, r := range s {
		freqS[r-'a']++
		freqT[t[i]-'a']++
	}

	return freqS == freqT
}
