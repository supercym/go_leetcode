func firstUniqChar(s string) int {
    d := make(map[rune]int)
    length := len(s)
    for i, v := range s {
        if _, ok := d[v]; !ok {
            d[v] = i
        } else {
            d[v] = length
        }
    }
    min_index := length
    for _, v := range d {
        if v < min_index {
            min_index = v
        }
    }
    if min_index == length {
        return -1
    }
    return min_index
}
