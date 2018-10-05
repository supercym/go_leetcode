package main

import (
	"fmt"
	"strconv"
)

func geneCom(upperbound, n int) []int {
	ans := []int{}
	for i := 0; i < upperbound; i++ {
		mask := 1
		count := 0
		for j := 0; j < 6; j++ {
			if i&mask == mask {
				count++
			}
			mask <<= 1
		}
		if count == n {
			ans = append(ans, i)
		}
	}
	return ans
}

func geneHM(n int) map[int]int {
	ans := make(map[int]int)
	for i := 0; i < 4; i++ {
		if n-i >= 0 && n-i <= 5 {
			ans[i] = n - i
		}
	}
	return ans
}

func readBinaryWatch(num int) []string {
	ans := []string{}
	HM := geneHM(num)
	for k, v := range HM {
		h := geneCom(12, k)
		m := geneCom(60, v)
		for _, hc := range h {
			for _, mc := range m {
				var s string
				if mc < 10 {
					s = strconv.Itoa(hc) + ":" + "0" + strconv.Itoa(mc)
				} else {
					s = strconv.Itoa(hc) + ":" + strconv.Itoa(mc)
				}
				ans = append(ans, s)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(readBinaryWatch(1))
	// fmt.Println(geneHM(1))
	// fmt.Println(geneCom(60, 1))
}
