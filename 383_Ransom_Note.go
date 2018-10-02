func canConstruct(ransomNote string, magazine string) bool {
    dRan := make(map[rune]int)
    for _, v := range ransomNote {
        if _, ok := dRan[v]; ok {
            dRan[v]++
        } else {
            dRan[v] = 1
        }
    }
    dMag := make(map[rune]int)
    for _, v := range magazine {
        if _, ok := dMag[v]; ok {
            dMag[v]++
        } else {
            dMag[v] = 1
        }
    }
    
    for k, v := range dRan {
        if _, ok := dMag[k]; !ok {
            return false
        } else {
            if v > dMag[k] {
                return false
            }
        }
    }
    return true
}