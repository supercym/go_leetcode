func judgeSquareSum(c int) bool {
    i, j := 0, int(math.Sqrt(float64(c)))
    res := false
    for i <= j {
        product := i*i+j*j
        if product < c {
            i++
        } else if product > c {
            j--
        } else {
            res = true
            break
        }
    }
    return res
}
