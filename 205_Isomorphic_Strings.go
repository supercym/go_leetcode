package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) == 0 && len(t) == 0 {
		return true
	}
	return sliceEqual(string2int(s), string2int(t))
}

func string2int(s string) []int {
	sint := []int{}
	index := 0
	d := make(map[rune]int, 0)

	for _, v := range s {
		if _, ok := d[v]; !ok {
			d[v] = index
			sint = append(sint, index)
			index++
		} else {
			sint = append(sint, d[v])
		}
	}
	return sint
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	// 留意这种情况: []int{} != []int(nil)
	// []int{} 和 []int(nil)长度都是0，但是二者不相等
	if (a == nil) != (b == nil) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func main() {

	fmt.Println(isIsomorphic("paper", "title"))
}
