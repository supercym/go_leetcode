package main

import (
	"fmt"
)

type node struct {
	index     int
	val       int
	cost      int
	processed bool
	neighbors []int
}

func find_lowest_cost(nodes []*node) *node {
	lowest_node := &node{}
	lowest_cost := 1 << 31
	flag := false
	for _, v := range nodes {
		if v.cost < lowest_cost && v.processed == false {
			lowest_node = v
			lowest_cost = v.cost
			flag = true
		}
	}
	if flag == false {
		return nil
	}
	return lowest_node
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	nodes := make([]*node, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nodes[i*n+j] = &node{index: i*n + j, val: grid[i][j], cost: 1 << 31}
			if i < m-1 && j < n-1 {
				nodes[i*n+j].neighbors = []int{(i+1)*n + j, i*n + (j + 1)}
			} else if i == m-1 && j < n-1 {
				nodes[i*n+j].neighbors = []int{i*n + (j + 1)}
			} else if i < m-1 && j == n-1 {
				nodes[i*n+j].neighbors = []int{(i+1)*n + j}
			}
		}
	}

	nodes[0].cost = grid[0][0]
	a := find_lowest_cost(nodes)
	for a != nil {
		cost := a.cost
		for _, b := range a.neighbors {
			new_cost := cost + nodes[b].val
			if new_cost < nodes[b].cost {
				nodes[b].cost = new_cost
			}
		}
		a.processed = true
		a = find_lowest_cost(nodes)
	}

	return nodes[m*n-1].cost
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
