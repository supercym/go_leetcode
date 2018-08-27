package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	ans := []int{1}

	for rowIndex != 0 {
		ans = gene_row(ans)
		rowIndex--
	}
	return ans
}

func gene_row(row []int) []int {
	next_row := []int{1}

	for i := 0; i < len(row)-1; i++ {
		next_row = append(next_row, row[i]+row[i+1])
	}

	next_row = append(next_row, 1)

	return next_row
}

func main() {
	fmt.Println(getRow(0))
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(21))

}
