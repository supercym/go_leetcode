func searchRange(nums []int, target int) []int {
    l, h := 0, len(nums)-1
    start, end := -1, -1
    for l <= h {
        m := int(l+(h-l)/2)
        if nums[m] < target {
            l = m+1
        } else if nums[m] > target {
            h = m-1
        } else {
            start, end = m, m
            for start > l {
                if nums[start-1] != target {
                    break
                } else {
                    start--
                }
            }
            for end < h {
                if nums[end+1] != target {
                    break
                } else {
                    end++
                }
            }
            break
        }
    }
    ans := []int{start, end}
    return ans
}
