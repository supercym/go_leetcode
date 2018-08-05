func twoSum(nums []int, target int) []int {
    ans := []int{}
    for i, i_v := range nums{
        for j, j_v := range nums{
            if (i != j) && (i_v + j_v == target){
                ans = append(ans, i, j)
                return ans
            }
        }
    }
    return ans
}