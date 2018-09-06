package main

import (
	"fmt"
)

func hammingWeight(n int32) int {
	count := 0
	var mask int32 = 1

	var i uint
	for i = 0; i < 32; i++ {
		if (n&mask)>>i == 1 {
			count++
		}
		mask <<= 1
	}
	return count
}
func main() {

	fmt.Println(hammingWeight(255))
	fmt.Println(hammingWeight(128))
}
