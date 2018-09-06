package main

import (
	"fmt"
)

func main() {

	var n int32 = 240
	var value, mask int32
	value, mask = 0, 1

	var i uint
	for i = 0; i < 32; i++ {
		value = (value << 1) | ((n & mask) >> i)
		mask <<= 1
	}

	fmt.Println(value)
}
