package quick_sort

func partition(nums []int, left, right int) int {
	if right-left < 2 {
		return left
	}
	start, end, mid := left, right-1, nums[left]
	for index := 1; index <= end; {
		if mid < nums[index] {
			nums[index], nums[end] = nums[end], nums[index]
			end--
		} else if nums[index] < mid {
			nums[index], nums[start] = nums[start], nums[index]
			start++
			index++
		}
	}
	return start
}

func quickSort(nums []int, left, right int) {
	if right-left < 2 {
		return
	}
	index := partition(nums, left, right)
	quickSort(nums, left, index)
	quickSort(nums, index+1, right)
}
