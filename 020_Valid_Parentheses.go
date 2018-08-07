package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	d := map[string]int{
		"(": 1,
		")": -1,
		"[": 2,
		"]": -2,
		"{": 3,
		"}": -3,
	}
	n := []int{}
	for _, v := range s {
		n = append(n, d[string(v)])
	}
	tmp := []int{}
	for i := 0; i < len(n); i++ {
		if len(tmp) != 0 && tmp[len(tmp)-1]+n[i] == 0 {
			tmp = tmp[:len(tmp)-1]
		} else {
			tmp = append(tmp, n[i])
		}
	}
	return len(tmp) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
