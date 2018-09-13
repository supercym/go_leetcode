package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	d := make(map[int]int, 0)
	m := make(map[int]int, 0)
	for i, v := range nums {
		if _, ok := d[v]; ok {
			if _, ok2 := m[v]; ok2 {
				if m[v] > i-d[v] {
					m[v] = i - d[v]
				}
			} else {
				m[v] = i - d[v]
			}
		}
		d[v] = i
	}

	for _, v := range m {
		if v <= k {
			return true
		}
	}
	return false
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(a, 2))
}
