func twoSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers)-1
    
    for left < right {
        sum := numbers[left] + numbers[right]
        if sum == target {
            break
        }else if sum > target {
            right--
        }else{
            left++
        }
    }
    ans := []int{left+1, right+1}
    return ans
}
