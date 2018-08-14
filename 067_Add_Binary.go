package main

import (
	"fmt"
	"strconv"
	"strings"
)

func string2Int(s string) []int {
	d := make([]int, len(s))
	for i, v := range s {
		if n, err := strconv.Atoi(string(v)); err == nil {
			d[i] = n
		} else {
			fmt.Println("string2Int error!")
			break
		}
	}
	return d
}

func addBinary(a string, b string) string {
	c := string2Int(a)
	d := string2Int(b)

	if len(c) >= len(d) {
		tmp := make([]int, len(c)-len(d))
		d = append(tmp, d...)
	} else {
		tmp := make([]int, len(d)-len(c))
		c = append(tmp, c...)
	}

	flag := 0
	n := 0
	val := make([]int, len(c))
	for i := len(c) - 1; i >= 0; i-- {
		n = c[i] + d[i] + flag
		switch n {
		case 0:
			val[i] = 0
			flag = 0
		case 1:
			val[i] = 1
			flag = 0
		case 2:
			val[i] = 0
			flag = 1
		case 3:
			val[i] = 1
			flag = 1
		}
	}
	if flag == 1 {
		val = append([]int{1}, val...)
	}

	s := make([]string, len(val))
	for i, v := range val {
		s[i] = strconv.Itoa(v)
	}

	ans := strings.Join(s, "")
	return ans
}

func main() {
	fmt.Println(addBinary("01", "11"))
	fmt.Println(addBinary("1010", "1011"))

}
