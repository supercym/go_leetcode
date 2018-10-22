func findAnagrams(s string, p string) []int {
	pChar := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		if _, ok := pChar[p[i]]; ok {
			pChar[p[i]]++
		} else {
			pChar[p[i]] = 1
		}
	}
	lenp := len(p)
	count := lenp
	ans := []int{}
	for i := 0; i < len(s); i++ {
		if _, ok := pChar[s[i]]; ok {
			if pChar[s[i]] >= 1 {
				count--
			}
			pChar[s[i]]--
		} else {
			pChar[s[i]] = -1
		}
		if i >= lenp {
			if pChar[s[i-lenp]] >= 0 {
				count++
			}
			pChar[s[i-lenp]]++
		}
		if count == 0 {
			ans = append(ans, i-lenp+1)
		}
	}
	return ans
}