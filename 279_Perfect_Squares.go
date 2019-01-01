package main

import (
	"fmt"
)


func numSquares(n int) int {
	var sqrtNums []int
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return 1
		} else {
			sqrtNums = append(sqrtNums, i*i)
		}
	}
	tree := sqrtNums
	ans := 1
	for {
		ans++
		var tmpTree []int
		for _, v := range tree {
			for _, k := range sqrtNums {
				if v+k == n {
					return ans
				} else if v+k < n {
					tmpTree = append(tmpTree, v+k)
				}
			}
		}
		tree = tmpTree
	}
}

func main()  {
	fmt.Println(numSquares(12))
}


