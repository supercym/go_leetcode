func threeSumClosest(nums []int, target int) int {
    abs := func(x int) int{
        if x < 0 {
            return -x
        }
        return x
    }
    sort.Ints(nums)
    ans := nums[0]+nums[1]+nums[2]
    diff := abs(target-ans)
    for i := 0; i < len(nums); i++ {
        l, r := i+1, len(nums)-1
        for l < r {
            s := nums[i]+nums[l]+nums[r]
            if abs(target-s) < diff {
                diff = abs(target-s)
                ans = s
            }
            if s < target {
                l++
            } else if s > target {
                r--
            } else {
                return s
            }
        }
    }
    return ans
}