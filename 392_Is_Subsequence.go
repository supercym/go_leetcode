func isSubsequence(s string, t string) bool {
    if len(s) > len(t) {
        return false
    }
    index := 0
    for i := 0; i < len(s); i++ {
        for s[i] != t[index] {
            index++
            if index == len(t) {
                return false
            }
        }
        index++
        if index == len(t) && i != len(s)-1 {
            return false
        }
    }
    return true
}
