func partitionLabels(S string) []int {
	ans := []int{}
	for len(S) != 0 {
		k := subPartition(S)
		S = S[k:]
		ans = append(ans, k)
	}
	return ans
}

func subPartition(S string) int {
	index := make(map[byte]int)
	for i := 0; i < len(S); i++ {
		index[S[i]] = i
	}

	subStrEnd := index[S[0]]
	pointer := 0

	for index[S[pointer]] <= subStrEnd {
		pointer++
		if pointer > subStrEnd {
			break
		}
		if index[S[pointer]] > subStrEnd {
			subStrEnd = index[S[pointer]]
		}
	}
	return pointer
}