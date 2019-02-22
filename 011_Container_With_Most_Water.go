package main

import (
	"fmt"
)

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var h int
	var flag bool
	if height[i] < height[j] {
		flag = true
		h = height[i]
	} else {
		flag = false
		h = height[j]
	}
	maxwater := h*(j-i)

	for {
		if flag {
			for i < len(height) && height[i] <= h {
				i++
			}
		} else {
			for j > 0 && height[j] <= h {
				j--
			}
		}
		if i >= j {
			break
		}
		if height[i] < height[j] {
			flag = true
			h = height[i]
		} else {
			flag = false
			h = height[j]
		}
		if h*(j-i) > maxwater {
			maxwater = h*(j-i)
		}
	}

	return maxwater
}

func main() {
	heights := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(heights))
}
