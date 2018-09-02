package main

import (
	"fmt"
)

// *******    Kadene Algorithm   ***********
// func maxSubArray(nums []int) int {
// 	max_sum := nums[0]
// 	if len(nums) > 1 {
// 		tmp_sum := 0
// 		for i := 0; i < len(nums); i++ {
// 			tmp_sum += nums[i]
// 			if tmp_sum > max_sum {
// 				max_sum = tmp_sum
// 			}
// 			if tmp_sum < 0 {
// 				tmp_sum = 0
// 			}
// 		}
// 	}
// 	return max_sum
// }

// *******    Divide & Conquer Approach   ***********
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}

	mid := int(length / 2)
	left := maxSubArray(nums[:mid])
	right := maxSubArray(nums[mid:])

	s1 := nums[mid-1]
	left_sum := 0
	for i := mid - 1; i >= 0; i-- {
		left_sum += nums[i]
		if left_sum > s1 {
			s1 = left_sum
		}
	}

	s2 := nums[mid]
	right_sum := 0
	for i := mid; i < length; i++ {
		right_sum += nums[i]
		if right_sum > s2 {
			s2 = right_sum
		}
	}

	sum_val := s1 + s2
	if left > sum_val {
		sum_val = left
	}
	if right > sum_val {
		sum_val = right
	}
	return sum_val
}

func main() {
	s := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(s))

}
