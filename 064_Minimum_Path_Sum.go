package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			} else if i != 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			} else if i != 0 && j != 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	grid1 := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

	grid2 := [][]int{
		{8, 6, 5, 5, 9, 8, 4, 3, 6, 4, 5, 0, 6, 4},
		{0, 1, 9, 1, 0, 6, 7, 6, 5, 7, 0, 6, 4, 4},
		{3, 9, 0, 3, 6, 5, 8, 5, 3, 4, 0, 1, 0, 5},
		{1, 7, 1, 3, 8, 3, 4, 1, 9, 4, 7, 4, 1, 0},
		{1, 5, 4, 6, 7, 4, 1, 3, 9, 9, 0, 4, 5, 0},
		{8, 6, 7, 9, 5, 1, 5, 5, 4, 1, 6, 5, 5, 6},
		{1, 8, 6, 4, 6, 2, 0, 3, 8, 1, 8, 9, 2, 0},
		{5, 0, 0, 3, 8, 9, 5, 3, 2, 0, 8, 6, 3, 7},
		{1, 1, 3, 3, 9, 1, 7, 5, 5, 0, 0, 3, 3, 0},
		{4, 6, 0, 9, 8, 2, 3, 6, 4, 8, 9, 6, 0, 9},
		{3, 3, 6, 6, 4, 1, 9, 6, 2, 9, 3, 7, 9, 4},
		{2, 2, 6, 6, 0, 7, 8, 8, 1, 1, 4, 5, 1, 5},
		{4, 1, 7, 7, 6, 3, 5, 3, 5, 5, 9, 9, 6, 2},
	}

	fmt.Println(minPathSum(grid1))
	fmt.Println(minPathSum(grid2))

}
