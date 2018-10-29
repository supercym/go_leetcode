package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	partition := func(low, high int) int {
		ref := nums[low]
		index := low+1
		for i := low+1; i <= high; i++ {
			if nums[i] >= ref {
				nums[i], nums[index] = nums[index], nums[i]
				index++
			}
		}
		nums[low], nums[index-1] = nums[index-1], nums[low]
		return index-1
	}
	l, h := 0, len(nums)-1
	ref := 0
	for l <= h {
		ref = nums[l]
		pos := partition(l, h)
		if pos == k-1 {
			break
		} else if pos < k-1 {
			l = pos+1
		} else {
			h = pos-1
		}
	}
	return ref
}

func main() {
	c := []int{3,3,3,2,1,5,6,4}
	fmt.Println(findKthLargest(c, 2))

}
