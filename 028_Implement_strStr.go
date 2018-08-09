package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || haystack == needle {
		return 0
	}

	ans := -1
	for i := 0; i <= len(haystack)-len(needle); i++ {
		tmp := []byte{}
		for j := 0; j < len(needle); j++ {
			tmp = append(tmp, haystack[i+j])
		}
		str := string(tmp[:])
		if str == needle {
			ans = i
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("favorite", "te"))
	fmt.Println(strStr("helllllo", "ll"))
	fmt.Println(strStr("hello", "ad"))
	fmt.Println(strStr("hello", ""))

}
