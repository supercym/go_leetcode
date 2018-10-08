func longestPalindrome(s string) int {
    d := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        if _, ok := d[s[i]]; ok {
            d[s[i]]++
        } else {
            d[s[i]] = 1
        }
    }
    count := 0
    Odd := false
    for _, v := range d {
        if v % 2 == 0 {
            count += v
        } else if v >= 3 {
            count += v-1
            Odd = true
        } else {
            Odd = true
        }
    }
    if Odd {
        count++
    }
    return count
}