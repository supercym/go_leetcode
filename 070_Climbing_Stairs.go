package main

import (
	"fmt"
)

func climbStairs(n int) int {
	count := make(map[int]int, 3)
	count[1] = 1
	count[2] = 2
	count[3] = 3
	if n <= 3 {
		return count[n]
	}
	m := 4
	for m <= n {
		count[m] = count[m-1] + count[m-2]
		m++
	}
	return count[n]
}

func main() {
	fmt.Println(climbStairs(35))

}
