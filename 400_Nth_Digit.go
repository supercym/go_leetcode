func findNthDigit(n int) int {
	f := func(x, y int) int {
		tmp := 1
		if y == 0 {
			return 1
		}
		for y > 0 {
			tmp *= x
			y--
		}
		return tmp
	}
	digit := 1
	for n > digit*9*f(10, digit-1) {
		n -= digit * 9 * f(10, digit-1)
		digit += 1
	}
	a := int((n - 1) / digit)
	b := int((n - 1) % digit)
	s := strconv.Itoa(f(10, digit-1) + a)
	ans, _ := strconv.Atoi(s[b : b+1])
	return ans
}
