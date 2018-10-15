func countSegments(s string) int {
    if len(s) == 0 {
        return 0
    }
    flag := true
    ans := 0
    for _, c := range s {
        if c != rune(" "[0]) && flag {
            ans++
            flag = false
        } else if c == rune(" "[0]) {
            flag = true
        }
    }
    return ans
}