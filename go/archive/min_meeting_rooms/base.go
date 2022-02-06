package min_meeting_rooms

import "sort"

func minMeetingRooms(intervals [][]int) int {
	start, end := make([]int, len(intervals)), make([]int, len(intervals))
	for index, interval := range intervals {
		start[index] = interval[0]
		end[index] = interval[1]
	}
	sort.Ints(start)
	sort.Ints(end)

	result := 0
	for iStart, iEnd := 0, 0; iStart < len(start); iStart++ {
		if start[iStart] < end[iEnd] {
			result++
		} else {
			iEnd++
		}
	}
	return result
}
