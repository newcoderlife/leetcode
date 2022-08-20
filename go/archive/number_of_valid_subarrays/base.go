package number_of_valid_subarrays

type Item struct {
	Index int
	Value int
}

func validSubarrays(nums []int) (ans int) {
	stack := []*Item{{Index: len(nums), Value: -1}}
	for index := len(nums) - 1; index >= 0; index-- {
		for stack[len(stack)-1].Value >= nums[index] {
			stack = stack[:len(stack)-1]
		}
		ans += stack[len(stack)-1].Index - index
		stack = append(stack, &Item{Index: index, Value: nums[index]})
	}
	return ans
}
