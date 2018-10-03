func findTheDifference(s string, t string) byte {
    ds := make(map[byte]int)
    for i := 0; i < len(s); i++ {
        if _, ok := ds[s[i]]; ok {
            ds[s[i]]++
        } else {
            ds[s[i]] = 1
        }
    }
    
    dt := make(map[byte]int)
    for i := 0; i < len(t); i++ {
        if _, ok := dt[t[i]]; ok {
            dt[t[i]]++
        } else {
            dt[t[i]] = 1
        }
    }
    
    var ans byte
    for k, v := range dt {
        if _, ok := ds[k]; !ok {
            ans = k
        } else {
            if v != ds[k] {
                ans = k
            }
        }
    }
    return ans
}