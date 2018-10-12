package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		if _, ok := set[v]; !ok {
			set[v] = false
		}
	}

	f := func() int {
		maxV := 0
		for k := range set {
			maxV = k
			break
		}
		for k := range set {
			if k > maxV {
				maxV = k
			}
		}
		return maxV
	}

	firstMax := f()
	delete(set, firstMax)
	if len(set) == 0 {
		return firstMax
	}

	secondMax := f()
	delete(set, secondMax)
	if len(set) == 0 {
		return firstMax
	}

	return f()
}

func main() {
	nums := []int{1, 2, 3, 0, 9, 3, 3, 3}
	fmt.Println(thirdMax(nums))
}
