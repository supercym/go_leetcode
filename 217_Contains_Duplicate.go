func containsDuplicate(nums []int) bool {
    d := make(map[int]bool, 0)
    for _, v := range nums {
        if _, ok := d[v]; ok {
            return true
        } else {
            d[v] = true
        }
    }
    return false
}
