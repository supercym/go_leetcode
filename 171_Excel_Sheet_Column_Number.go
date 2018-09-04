package main

import (
	"fmt"
)

func titleToNumber(s string) int {
	d := make(map[byte]int, 26)
	for i := 1; i < 27; i++ {
		d[byte(i+64)] = i
	}

	val := 0
	index := 0

	for i := 0; i < len(s); i++ {
		val += d[s[len(s)-1-i]] * pow(26, index)
		index++
	}
	return val
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	ans := 1
	for i := 0; i < n; i++ {
		ans *= x
	}
	return ans
}

func main() {
	fmt.Println(titleToNumber("AB"))

}
