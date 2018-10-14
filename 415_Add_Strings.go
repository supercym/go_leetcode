package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	fn := func(a, b int) []int {
		res := make([]int, 2)
		res[0] = (a + b) % 10
		res[1] = (a + b) / 10
		return res
	}

	fs := func(length int, num string) string {
		s := make([]byte, length)
		for i := 0; i < length; i++ {
			if i < length-len(num) {
				s[i] = 48
			} else {
				s[i] = num[i-(length-len(num))]
			}
		}
		return string(s[:])
	}
	var length int
	if len(num1) > len(num2) {
		length = len(num1)
		num2 = fs(length, num2)
	} else {
		length = len(num2)
		num1 = fs(length, num1)
	}

	ans := make([]byte, length+1)
	k := 0
	for i := length - 1; i >= 0; i-- {
		s := fn(int(num1[i]-48), int(num2[i]-48))
		s1, s2 := s[0], s[1]
		s1 += k
		if s1 == 10 {
			s1 = 0
			s2 = 1
		}
		k = s2
		ans[i+1] = byte(s1 + 48)
	}
	if k != 0 {
		ans[0] = 49
		fmt.Println(ans)
		return string(ans[:])
	}
	fmt.Println(ans)
	return string(ans[1:])
}

func main() {
	num1 := "9999"
	num2 := "11"
	fmt.Println(addStrings(num1, num2))
}
