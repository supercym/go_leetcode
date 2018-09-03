package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	d := make(map[int]byte, 26)
	for i := 1; i < 26; i++ {
		d[i] = byte(i + 64)
	}
	d[0] = byte(90)

	s := make([]byte, 0)
	for {
		a := n % 26
		b := int((n - a) / 26)
		s = append(s, d[a])
		if a == 0 {
			b--
		}

		if b == 0 {
			break
		} else if b < 26 {
			s = append(s, d[b])
			break
		} else if b == 26 {
			s = append(s, d[0])
			break
		} else {
			n = b
		}

	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	ans := string(s[:])
	return ans
}

func main() {
	fmt.Println(convertToTitle(77))

}
