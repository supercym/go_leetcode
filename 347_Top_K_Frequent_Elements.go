func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int)
	maxCount := 0
	for _, n := range nums {
		if _, ok := counter[n]; ok {
			counter[n]++
		} else {
			counter[n] = 1
		}
		if counter[n] > maxCount {
			maxCount = counter[n]
		}
	}

	barrels := make([][]int, maxCount)
	for i := range barrels {
		barrels[i] = []int{}
	}
	for k, v := range counter {
		barrels[v-1] = append(barrels[v-1], k)
	}

	res := make([]int, 0)
	for i := maxCount-1; i >= 0; i-- {
		res = append(res, barrels[i]...)
		if len(res) == k {
			break
		}
	}
	return res
}
