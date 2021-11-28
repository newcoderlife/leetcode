package subsets

import (
	"container/list"
	"sort"
)

func subsets(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	queue := list.New()
	results := [][]int{{}}
	for _, num := range nums {
		results = append(results, []int{num})
		queue.PushBack([]int{num})
	}

	for queue.Len() > 0 {
		current := queue.Front().Value.([]int)
		queue.Remove(queue.Front())

		last := current[len(current)-1]
		for _, num := range nums {
			if last < num {
				result := make([]int, len(current)+1)
				copy(result, current)
				result[len(result)-1] = num
				results = append(results, result)
				queue.PushBack(result)
			}
		}
	}
	return results
}
