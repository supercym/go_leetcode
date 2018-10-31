func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	for _, v := range g {
		index := 0
		for index < len(s) {
			sLength := len(s)
			if s[index] >= v {
				res++
				s = append(s[:index], s[index+1:]...)
				break
			} else {
				index++
			}
			if len(s) == 0 || index == sLength {
				return res
			}
		}
	}
	return res
}