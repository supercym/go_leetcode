func isPowerOfTwo(n int) bool {
    mask := 1
    for i:= 0; i < 64; i++ {
        if n ^ mask == 0 {
            return true
        }
        mask <<= 1
    }
    return false
}
