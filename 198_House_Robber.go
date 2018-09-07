func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }
    val_2 := nums[0]
    val_1 := max(nums[0], nums[1])
    for i := 2; i < len(nums); i++ {
        max_cur := max(val_1, val_2 + nums[i])
        val_2 = val_1
        val_1 = max_cur      
    }
    return val_1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
