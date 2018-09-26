func isPowerOfThree(n int) bool {
    if n == 0 {
        return false
    } else if n == 1 {
        return true
    }
    for n % 3 == 0 {
        n /= 3
        if n == 1 {
            return true
        }
    }
    return false
}

