func minSubArrayLen(s int, nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    i, j := 0, 0
    sum  := nums[0]
    arrlen := len(nums)+1
    
    for {
        if sum >= s {
            if arrlen > (j-i+1) {
                arrlen = j-i+1
            }
            if i == j {
                break
            }
            sum -= nums[i]
            i++
        } else {
            if j == len(nums)-1 {
                break
            }
            j++
            sum += nums[j]
        }
    }
    if arrlen == len(nums)+1 {
        return 0
    }
    return arrlen
}