func threeSum(nums []int) [][]int {
    ans := make([][]int, 0)
    if len(nums) < 3 {
        return ans
    }
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        l, r := i+1, len(nums)-1      
        for l < r {
            s := nums[l] + nums[r]
            if s < -nums[i] {
                l++
            } else if s > -nums[i] {
                r--
            } else {
                if l > i+1 && nums[l] == nums[l-1] {
                    l, r = l+1, r-1
                } else {
                    ans = append(ans, []int{nums[i], nums[l], nums[r]})
                    l, r = l+1, r-1
                }
            }
        }
    }
    return ans
}