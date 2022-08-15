package remove_vowels_from_a_string

func anagramMappings(nums1 []int, nums2 []int) []int {
	exist := make([]bool, len(nums2))
	result := make([]int, len(nums1))
	for i, n := range nums1 {
		for j, m := range nums2 {
			if !exist[j] && n == m {
				result[i] = j
				exist[j] = true
				break
			}
		}
	}
	return result
}
