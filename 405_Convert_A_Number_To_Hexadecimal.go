package main

import (
	"fmt"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = -num
		num = ^num
		num++
	}
	n := make([]int, 8)
	for i := 0; i < 8; i++ {
		mask := 15
		n[7-i] = mask & num
		num >>= 4
	}
	d := make(map[int]byte)
	s := "0123456789abcdef"
	for i := 0; i < len(s); i++ {
		d[i] = s[i]
	}
	ans := []byte{}
	for i := 0; i < 8; i++ {
		if n[i] != 0 {
			for j := i; j < 8; j++ {
				ans = append(ans, d[n[j]])
			}
			break
		}
	}
	return string(ans[:])
}

func main() {
	fmt.Println(toHex(-1))
	fmt.Println(toHex(16))
}
