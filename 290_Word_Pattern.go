package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	p := []byte(pattern)
	s := strings.Fields(str)
	if len(p) != len(s) {
		return false
	}
	d := make(map[byte]string)
	set := make(map[string]bool)
	for i := 0; i < len(p); i++ {
		if v, ok := d[p[i]]; ok {
			if v != s[i] {
				return false
			}
		} else {
			d[p[i]] = s[i]
			if _, ok2 := set[s[i]]; ok2 {
				return false
			} else {
				set[s[i]] = false
			}
		}
	}
	return true
}

func main() {
	pattern := "abba"
	str := "dog cat cat dog"
	pattern2 := "abba"
	str2 := "dog dog dog dog"
	fmt.Println(wordPattern(pattern, str))
	fmt.Println(wordPattern(pattern2, str2))

}
