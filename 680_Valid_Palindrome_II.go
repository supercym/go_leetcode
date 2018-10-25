func validPalindrome(s string) bool {
    f := func(t string) bool {
        i, j := 0, len(t)-1
        for i < j {
            if t[i] != t[j] {
                return false
            }
            i++
            j--
        }
        return true
    }
    
    res := true
    i, j := 0, len(s)-1
    for i < j {
        if s[i] != s[j] {
            if f(string(s[i+1:j+1])) || f(string(s[i:j])) {
                res = true
            } else {
                res = false
            }
            break
        }
        i++
        j--
    }
    return res
}
