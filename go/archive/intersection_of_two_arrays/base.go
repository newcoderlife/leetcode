package intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	exist1 := make(map[int]bool)
	for _, num := range nums1 {
		exist1[num] = true
	}
	exist2 := make(map[int]bool)
	for _, num := range nums2 {
		exist2[num] = true
	}

	result := make([]int, 0)
	for num := range exist1 {
		if exist2[num] {
			result = append(result, num)
		}
	}

	return result
}
