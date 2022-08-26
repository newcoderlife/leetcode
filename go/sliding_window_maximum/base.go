package sliding_window_maximum

import (
	"container/list"
)

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	buffer := list.New()
	for i := 0; i < len(nums); i++ {
		for buffer.Len() > 0 && nums[i] >= nums[buffer.Back().Value.(int)] {
			buffer.Remove(buffer.Back())
		}
		buffer.PushBack(i)

		for buffer.Front().Value.(int)+k <= i {
			buffer.Remove(buffer.Front())
		}
		if i+1 >= k {
			result = append(result, nums[buffer.Front().Value.(int)])
		}
	}

	return result
}
