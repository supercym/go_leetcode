package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	n := float64(x)
	ans := 1.0
	tmp := 0.0
	for {
		tmp = 0.5 * (ans + (n / ans))
		if abs(tmp-ans) < 1 {
			break
		} else {
			ans = tmp
		}
	}
	return int(tmp)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	x := 1836583659
	fmt.Println(mySqrt(x))
}
