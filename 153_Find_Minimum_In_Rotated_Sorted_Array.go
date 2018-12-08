func findMin(nums []int) int {
    l, h := 0, len(nums)-1
    for l < h {
        m := int(l+(h-l)/2)
        if nums[m] < nums[h] {
            h = m
        } else if nums[m] > nums[l] {
            l = m
        } else {
            break
        }
    }
    return nums[h]
}
