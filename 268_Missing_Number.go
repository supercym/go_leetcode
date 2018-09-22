func missingNumber(nums []int) int {
    ans := len(nums)
    for i := 0; i < len(nums); i++ {
        ans ^= i
        ans ^= nums[i]
    }
    return ans
}
