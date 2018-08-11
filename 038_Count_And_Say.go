package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	tmp := []int{1}
	for i := 0; i < n-1; i++ {
		tmp = geneNext(tmp)
	}
	s := []string{}
	for i := 0; i < len(tmp); i++ {
		s = append(s, strconv.Itoa(tmp[i]))
	}
	ans := strings.Join(s, "")
	return ans
}

func geneNext(nums []int) []int {
	count := 0
	val := nums[0]
	ans := []int{}

	for i, v := range nums {
		if v == val {
			count += 1
		} else {
			ans = append(ans, count)
			ans = append(ans, val)
			count = 1
			val = v
		}

		if i == len(nums)-1 {
			ans = append(ans, count)
			ans = append(ans, val)
		}
	}
	return ans
}

func main() {

	fmt.Println(1, countAndSay(1))
	fmt.Println(2, countAndSay(2))
	fmt.Println(3, countAndSay(3))
	fmt.Println(4, countAndSay(4))
	fmt.Println(5, countAndSay(5))
	// fmt.Println("###################")
	// fmt.Println(searchInsert(s, 2))
	// fmt.Println("###################")
	// fmt.Println(searchInsert(s, 7))
	// fmt.Println("###################")
	// fmt.Println(searchInsert(s, 0))
	// fmt.Println("###################")
	// fmt.Println(searchInsert(s, 6))

}
