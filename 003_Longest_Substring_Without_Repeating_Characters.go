func lengthOfLongestSubstring(s string) int {
    i, j := 0, -1
    chars := make(map[byte]bool, 0)
    maxlength := 0
    
    for i < len(s) {
        if j+1 < len(s) {
            if _, ok := chars[s[j+1]]; !ok {
                j++
                chars[s[j]] = true
            } else {
                delete(chars, s[i])
                i++
            }
        } else {
            delete(chars, s[i])
            i++
        }
        if maxlength < j-i+1 {
            maxlength = j-i+1
        }
    }
    return maxlength
}