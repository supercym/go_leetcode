package main

import (
	"fmt"
)

func isHappy(n int) bool {
	s := make(map[int]bool, 0)
	s[n] = true
	for {
		val := 0
		for n != 0 {
			remainder := n % 10
			val += remainder * remainder
			n = int((n - remainder) / 10)
		}
		if val == 1 {
			return true
		}
		if _, ok := s[val]; ok {
			return false
		} else {
			s[val] = true
		}
		n = val
	}
}

func main() {

	fmt.Println(isHappy(19))
}
