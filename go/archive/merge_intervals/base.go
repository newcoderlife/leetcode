package merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	results := make([][]int, 0)
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] <= right {
			continue
		} else if intervals[i][0] <= right {
			right = intervals[i][1]
		} else {
			results = append(results, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		}
	}

	results = append(results, []int{left, right})
	return results
}
