package main

func strStr(haystack string, needle string) int {
	return BM(haystack, needle)
}

func BM(str, pattern string) int {
	if len(str) < len(pattern) {
		return -1
	}
	bc := genBC(pattern)
	sf, pf := genGS(pattern)
	i := 0
	for i <= len(str)-len(pattern) {
		j := len(pattern) - 1
		for ; j >= 0 && str[i+j] == pattern[j]; j-- {
		}
		if j == -1 {
			return i
		}
		bs := j - bc[int(str[i+j])]
		gs := 0
		if j < len(pattern)-1 {
			gs = getCS(j, len(pattern), sf, pf)
		}
		i += maxInt(gs, bs)
	}
	return -1
}

func getCS(j int, lenP int, sf []int, pf []bool) int {
	k := lenP - 1 - j // 好后缀个数 [j+1, lenP-1] lenP-1 - (j+1) + 1
	if sf[k] != -1 {
		return j + 1 - sf[k]
	}
	for i := j + 1; i < lenP; i++ {
		if pf[i] {
			return i
		}
	}
	return lenP
}

func genBC(pattern string) []int {
	bc := make([]int, 256)
	for i := 0; i < len(bc); i++ {
		bc[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		bc[int(pattern[i])] = i
	}
	return bc
}

func genGS(pattern string) (suffix []int, prefix []bool) {
	suffix = make([]int, len(pattern)+1)
	prefix = make([]bool, len(pattern)+1)
	for i := 0; i < len(suffix); i++ {
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < len(pattern)-1; i++ {
		j := i
		k := 0
		for j >= 0 && pattern[j] == pattern[len(pattern)-1-k] {
			k++
			suffix[k] = j
			j--
		}
		if j == -1 {
			prefix[k] = true
		}
	}
	return
}
