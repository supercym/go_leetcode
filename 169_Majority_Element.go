func majorityElement(nums []int) int {
    d := make(map[int]int, len(nums))
    for _, v := range(nums) {
        if _, ok := d[v]; ok {
            d[v]++
        } else {
            d[v] = 1
        }
    }
    ans := 0
    max_count := 0
    for i, v := range(d) {
        if v > max_count {
            max_count = v
            ans = i
        }
    }
    return ans
}
