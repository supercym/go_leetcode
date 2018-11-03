/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func eraseOverlapIntervals(intervals []Interval) int {
    if len(intervals) == 0 {
        return 0
    }
    qsort(intervals, 0, len(intervals)-1)
    ans := []Interval{intervals[0]}
    for {
        lastInterval := ans[len(ans)-1]
        ansLength := len(ans)
        for _, t := range intervals {
            if t.Start >= lastInterval.End {
                ans = append(ans, t)
                break
            }
        }
        if len(ans) == ansLength {
            break
        }
    }
    return len(intervals) - len(ans)
}

func qsort(nums []Interval, start, end int) {
	if start <= end {
		k := partition(nums, start, end)
		qsort(nums, start, k-1)
		qsort(nums, k+1, end)
	}

}

func partition(nums []Interval, start, end int) int {
	ref := nums[end].End
	i := start-1
	for j := start; j < end; j++{
		if nums[j].End <= ref {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[end] = nums[end], nums[i+1]
	return i+1
}
