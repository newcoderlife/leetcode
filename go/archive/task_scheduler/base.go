package task_scheduler

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func leastInterval(tasks []byte, n int) int {
	counts := make([]int, 26)
	for _, c := range tasks {
		counts[int(c)-int('A')]++
	}
	sort.Slice(counts, func(i, j int) bool { return counts[i] > counts[j] })

	tail := 1
	for tail < len(counts) && counts[tail] == counts[0] {
		tail++
	}

	return max(len(tasks), tail+(counts[0]-1)*(n+1))
}
