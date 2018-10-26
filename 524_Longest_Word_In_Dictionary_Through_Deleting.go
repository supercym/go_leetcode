package main

import "fmt"

func findLongestWord(s string, d []string) string {
	f := func(target, t string) bool {
		if len(target) < len(t) {
			return false
		}
		i, j := 0, 0
		for i < len(target) && j < len(t) {
			if target[i] == t[j] {
				j++
			}
			i++
		}
		return j == len(t)
	}

	var res string
	maxLength := 0
	for _, t := range d {
		if len(t) < maxLength || (len(t) == maxLength && t > res) {
			continue
		}
		if f(s, t) {
			maxLength = len(t)
			res = t
		}
	}
	return res
}


func main() {
	s := "abpcplea"
	d := []string{"ale","apple","monkey","plea"}
	fmt.Println(findLongestWord(s, d))
}
