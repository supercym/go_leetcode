func trailingZeroes(n int) int {
    count := 0
    tmp := int(n/5)
    for tmp != 0 {
        count += tmp
        tmp = int(tmp/5)
    }
    return count
}
