func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	quicksort(points, 0, len(points)-1)
	span := points[0]
	cnt := 1
	for _, t := range points {
		if t[0] > span[1] {
			span = t
			cnt++
		}
	}
	return cnt
}

func quicksort(nums [][]int, start, end int)  {
	if start < end {
		k := partition(nums, start, end)
		quicksort(nums, start, k-1)
		quicksort(nums, k+1, end)
	}
}

func partition(nums [][]int, start, end int) int {
	ref := nums[end][1]
	i := start-1
	for j := start; j < end; j++ {
		if nums[j][1] <= ref {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = nums[end], nums[i+1]
	return i+1
}