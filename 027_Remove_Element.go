func removeElement(nums []int, val int) int {
    a, b := 0, 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            a += 1
        } else {
            nums[i-a] = nums[i]
            b += 1
        }
    }
    return b
}