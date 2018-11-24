func checkPossibility(nums []int) bool {
    if len(nums) <= 2 {
        return true
    }
    count := 0
    for i := 1; i < len(nums); i++ {
        if nums[i-1] > nums[i] {
            count++
            if (i-2) < 0 || nums[i] >= nums[i-2] {
                nums[i-1] = nums[i]
            } else {
                nums[i] = nums[i-1]
            }
        }
    }
    return count <= 1
}