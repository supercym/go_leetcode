func isPowerOfFour(num int) bool {
    mask := 1
    var i uint = 0
    for i < 32 {
        if num == mask << i {
            return true
        }
        i += 2
    }
    return false
}
