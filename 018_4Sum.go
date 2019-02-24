func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    length := len(nums)
    ans := make([][]int, 0)
    for i := 0; i < length; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        } 
        for j := i+1; j < length; j++ {
            if j > i+1 && nums[j] == nums[j-1] {
                continue
            }
            l, r := j+1, length-1
            for l < r {
                s := nums[l]+nums[r]
                if s < target-nums[i]-nums[j] {
                    l++
                } else if s > target-nums[i]-nums[j] {
                    r--
                } else {
                    if l > j+1 && nums[l] == nums[l-1] {
                        l, r = l+1, r-1
                    } else {
                        ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
                        l, r = l+1, r-1
                    }
                }
            }
        }
    }
    return ans
}