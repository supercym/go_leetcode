func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    j := 0
    val := nums[0]
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            count++
            if count <= 2 {
                nums[j] = nums[i]
                j++
            }
        } else {
            val = nums[i]
            nums[j] = nums[i]
            j++
            count = 1
        }
    }
    return j
}