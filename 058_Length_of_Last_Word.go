package main

import (
	"fmt"
	// "sort"
	// "strconv"
	// "strings"
)

func lengthOfLastWord(s string) int {
	a := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	count := 0
	flag := false
	for _, v := range a {
		if v != 32 {
			count += 1
			flag = true
		} else if flag == true {
			break
		}
	}
	return count

}

func main() {

	fmt.Println(lengthOfLastWord("hello go"))
	fmt.Println(lengthOfLastWord(" hello go "))
	fmt.Println(lengthOfLastWord("hello"))
	fmt.Println(lengthOfLastWord(" hello "))
	fmt.Println(lengthOfLastWord("hello "))
	fmt.Println(lengthOfLastWord("hello "))
	fmt.Println(lengthOfLastWord(" "))
}
