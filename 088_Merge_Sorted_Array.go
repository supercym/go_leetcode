package main

import (
	"fmt"
)

// 这道题最让人讨厌的地方在于不让有返回值
// 等价于对nums1原地修改，只能对其进行元素值的修改，
// 任何插入操作或者append都会在一个新产生的底层数组上进行
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	} else {
		nums_tmp := make([]int, m+n, 2*(m+n))
		copy(nums_tmp, nums1)
		for i := 0; i < n; i++ {
			nums_tmp = insert_list(nums_tmp, m, nums2[i])
			m += 1
		}
		copy(nums1, nums_tmp)
	}
}

func insert_list(nums []int, m, d int) []int {
	low := 0
	high := m - 1
	for low <= high {
		if (high - low) < 2 {
			if d < nums[low] {
				nums = insert_slice(nums, low, d)
			} else if d > nums[high] {
				nums = insert_slice(nums, high+1, d)
			} else {
				nums = insert_slice(nums, high, d)
			}
			break
		}
		mid := int((low + high) / 2)
		tmp := nums[mid]
		if tmp == d {
			nums = insert_slice(nums, mid, d)
			break
		} else if tmp > d {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nums
}

func insert_slice(nums []int, index, val int) []int {
	tmp := append([]int{}, nums[index:]...)
	nums = append(nums[:index], val)
	nums = append(nums, tmp...)
	return nums
}

func main() {

	a := []int{4, 5, 7, 0, 0, 0}
	b := []int{1, 6, 99}

	merge(a, 3, b, 3)
	fmt.Println(a)

}
