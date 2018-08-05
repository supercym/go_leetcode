func reverse(x int) int {
    flag := 1
	if x < 0 {
		x = -x
		flag = -1
	}
	s := strconv.Itoa(x)
	b := []byte(s)
	for from, to := 0, len(b)-1; from < to; from, to = from+1, to-1 {
		b[from], b[to] = b[to], b[from]
	}
	d := string(b[:])
	ans, _ := strconv.Atoi(d)
	ans = ans * flag
	if (ans < -(1 << 31)) || (ans > (1<<31)-1) {
		return 0
	}
	return ans
}