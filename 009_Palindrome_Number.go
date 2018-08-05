func isPalindrome(x int) bool {
    if x < 0{
        return false
    }
    s := strconv.Itoa(x)
    b := []byte(s)
    for from, to := 0, len(b)-1; from < to; from, to = from+1, to-1{
        if b[from] != b[to]{
            return false
        }
    }
    return true
}