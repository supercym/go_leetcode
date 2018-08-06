func romanToInt(s string) int {
    val := 0
	d := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	num := []int{}
	for _, v := range s {
		num = append(num, d[string(v)])
	}
	if len(num) == 1 {
		return num[0]
	}
	for i := 0; i < len(num)-1; i++ {
		if num[i] < num[i+1] {
			num[i] = -num[i]
		}
	}
	for _, v := range num {
		val += v
	}
	return val
}