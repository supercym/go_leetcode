func isPerfectSquare(num int) bool {
    tmp := float64(num)
    k := tmp
    for abs(k*k - tmp) > 1 {
        k = (k + tmp/k)/2
    }
    n := int(k)
    if n*n == num || (n-1)*(n-1) == num || (n+1)*(n+1) == num {
        return true
    }
    return false
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}