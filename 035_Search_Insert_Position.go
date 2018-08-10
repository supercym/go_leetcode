package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	ans := 0
	for i, v := range nums {
		if v >= target {
			ans = i
			break
		}
	}
	return ans
}

func main() {
	s := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(s, 5))
	fmt.Println("###################")
	fmt.Println(searchInsert(s, 2))
	fmt.Println("###################")
	fmt.Println(searchInsert(s, 7))
	fmt.Println("###################")
	fmt.Println(searchInsert(s, 0))
	fmt.Println("###################")
	fmt.Println(searchInsert(s, 6))

}
