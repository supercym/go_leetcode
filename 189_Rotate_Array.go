package main

import (
	"fmt"
)

// func rotate(nums []int, k int) {
// 	if k > len(nums) {
// 		k %= len(nums)
// 	}
// 	tmp := make([]int, k)
// 	for i := len(nums) - k; i < len(nums); i++ {
// 		tmp[i-(len(nums)-k)] = nums[i]
// 	}
// 	for i := len(nums) - k - 1; i >= 0; i-- {
// 		nums[i+k] = nums[i]
// 	}
// 	for i := 0; i < k; i++ {
// 		nums[i] = tmp[i]
// 	}
// }

// ******    reverse 3 times    ***********
func rotate(nums []int, k int) {
	n := len(nums)
	if k > n {
		k %= n
	}

	reverse(nums, 0, n-k-1)
	reverse(nums, n-k, n-1)
	reverse(nums, 0, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a)

}
