func reverseVowels(s string) string {
	a := []byte("aeiouAEIOU")
	m := make(map[byte]bool)
	for _, v := range a {
		m[v] = true
	}
	b := []byte(s)
	i, j := 0, len(b)-1
	for i < j {
		if _, ok := m[b[i]]; !ok {
			i++
		} else if _, ok2 := m[b[j]]; !ok2 {
			j--
		} else {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}
	return string(b[:])
}
