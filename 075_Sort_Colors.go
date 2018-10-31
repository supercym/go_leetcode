func sortColors(nums []int)  {
    low, high := -1, len(nums)
    index := 0
    for index < high {
        if nums[index] == 0 {
            low++
            nums[low], nums[index] = nums[index], nums[low]
            index = low+1
        } else if nums[index] == 2 {
            high--
            nums[high], nums[index] = nums[index], nums[high]
        } else {
            index++
        }
    }
    return
}