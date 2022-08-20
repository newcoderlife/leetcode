package employee_free_time

import "sort"

type Interval struct {
	Start int
	End   int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func employeeFreeTime(schedules [][]*Interval) []*Interval {
	buffer := make([]*Interval, 0)
	for _, people := range schedules {
		buffer = append(buffer, people...)
	}
	sort.Slice(buffer, func(i, j int) bool {
		if buffer[i].Start == buffer[j].Start {
			return buffer[i].End < buffer[j].End
		}
		return buffer[i].Start < buffer[j].Start
	})

	results := make([]*Interval, 0)
	current := buffer[0].End
	for index := 1; index < len(buffer); index++ {
		if buffer[index].Start > current {
			results = append(results, &Interval{Start: current, End: buffer[index].Start})
		}
		current = max(current, buffer[index].End)
	}

	return results
}
