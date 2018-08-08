func removeDuplicates(nums []int) int {
    if len(nums) == 0 || len(nums) == 1 {
        return len(nums)
    }
    a := 0
    for b := 1; b < len(nums); b++ {
        if nums[b] == nums[a] {
            continue
        } else {
            nums[a+1] = nums[b]
            a += 1
        }           
    }
    return a+1
}