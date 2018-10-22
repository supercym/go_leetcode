func arrangeCoins(n int) int {
    ans := 0
    k := int((math.Sqrt(float64(8*n+1))-1)/2)
    for {
        tempSum := int(k*(k+1)/2)
        if tempSum == n {
            ans = k
            break
        } else if tempSum > n {
            ans = k-1
            break
        }
        k++
    }
    return ans
}
