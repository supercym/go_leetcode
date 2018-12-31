package main

import (
	"fmt"
	"strconv"
)


func diffWaysToCompute(input string) []int {
	var ans []int
	for i := 0; i < len(input); i++ {
		if input[i] == '+' || input[i] == '-' || input[i] == '*' {
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch input[i] {
					case '+':
						ans = append(ans, l+r)
					case '-':
						ans = append(ans, l-r)
					case '*':
						ans = append(ans, l*r)
					}
				}
			}
		}
	}
	if len(ans) == 0 {
		n, _ := strconv.Atoi(input)
		ans = append(ans, n)
	}

	return ans
}

func main() {
	S := "2*3-4*5"

	fmt.Println(diffWaysToCompute(S))
}
