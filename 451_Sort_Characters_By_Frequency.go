func frequencySort(s string) string {
	counter := make(map[rune]int)
	maxCount := 0
	for _, c := range s {
		if _, ok := counter[c]; ok {
			counter[c]++
		} else {
			counter[c] = 1
		}
		if counter[c] > maxCount {
			maxCount = counter[c]
		}
	}

	barrels := make([][]rune, maxCount)
	for i := range barrels {
		barrels[i] = []rune{}
	}
	for k, v := range counter {
		barrels[v-1] = append(barrels[v-1], k)
	}
	res := make([]rune, 0)
	for i := maxCount-1; i >= 0; i-- {
		for j := 0; j < len(barrels[i]); j++ {
			for m := 0; m < i+1; m++ {
				res = append(res, barrels[i][j])
			}
		}
	}
	return string(res)
}