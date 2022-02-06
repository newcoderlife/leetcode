package reconstruct_queue

import "sort"

func insert(position int, people []int, result [][]int) [][]int {
	if position == 0 {
		return append([][]int{people}, result...)
	} else if position == len(result) {
		return append(result, people)
	} else {
		return append(result[:position], append([][]int{people}, result[position:]...)...)
	}
}

func reconstructQueue(people [][]int) [][]int {
	result := make([][]int, 0, len(people))
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	for _, people := range people {
		result = insert(people[1], people, result)
	}
	return result
}
