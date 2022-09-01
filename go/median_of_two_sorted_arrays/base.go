package median_of_two_sorted_arrays

import (
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)

	if mid := len(nums) / 2; mid*2 == len(nums) {
		return (float64(nums[mid]) + float64(nums[mid-1])) / 2
	} else {
		return float64(nums[mid])
	}
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	left, right, index1, index2, target := -1, -1, 0, 0, (len(nums1)+len(nums2))/2
	for i := 0; i <= target; i++ {
		left = right
		if index1 < len(nums1) && (index2 >= len(nums2) || nums1[index1] < nums2[index2]) {
			right = nums1[index1]
			index1++
		} else {
			right = nums2[index2]
			index2++
		}
	}

	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(right)
	}
	return (float64(left) + float64(right)) / 2
}
