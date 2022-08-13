package top_k_frequent_elements

import "sort"

type Num struct {
	Number int
	Time   int
}

func topKFrequent(nums []int, k int) []int {
	exist := make(map[int]int)
	for _, num := range nums {
		exist[num] += 1
	}

	numbers := make([]*Num, 0, len(exist))
	for num, time := range exist {
		numbers = append(numbers, &Num{Number: num, Time: time})
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Time > numbers[j].Time
	})

	if k > len(nums) {
		k = len(nums)
	}
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, numbers[i].Number)
	}
	return result
}
