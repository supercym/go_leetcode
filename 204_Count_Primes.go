func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }
    count := make([]int, n)
    count[0], count[1] = 1, 1
	// 我们只需遍历[2,sqrt(n)],因为超过sqrt(n)部分如果不是素数，则作为因子在前面的数已经被删除了
    sqrt_n := int(math.Sqrt(float64(n)))+1
    for i := 2; i < sqrt_n; i++ {
        if count[i] == 0 {
            for j := i*i; j < n; j += i {
                count[j] = 1
            }
        }
    }
    ans := 0
    for _, v := range count {
        if v == 0 {
            ans += 1
        }
    }
    return ans
}
