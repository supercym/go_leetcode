package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	m := len(digits)
	// d := make([]int, m)
	d := []int(digits)
	d[m-1] += 1

	for i := m - 1; i >= 0; i-- {
		if d[i] < 10 {
			break
		} else if i != 0 {
			d[i] = 0
			d[i-1] += 1
		} else {
			d[i] = 1
			d = append(d, 0)
		}
	}

	return d
}

func main() {
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{1, 2, 3, 4}))
	fmt.Println(plusOne([]int{1, 2, 3, 9}))
	fmt.Println(plusOne([]int{1, 2, 9, 9}))
	fmt.Println(plusOne([]int{1, 9, 9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))

}
