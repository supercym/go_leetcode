func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    ds := make(map[byte]int)
    dt := make(map[byte]int)
    for _, v := range []byte(s) {
        if _, ok := ds[v]; ok {
            ds[v]++
        } else {
            ds[v] = 1
        }
    }
    for _, v := range []byte(t) {
        if _, ok := dt[v]; ok {
            dt[v]++
        } else {
            dt[v] = 1
        }
    }
    for k, v := range ds {
        if _, ok := dt[k]; !ok {
            return false
        }
        if v != dt[k] {
            return false
        }
    }
    return true
}
