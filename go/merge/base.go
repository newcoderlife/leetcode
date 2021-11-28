package merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, index := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; index-- {
		if p1 == -1 {
			nums1[index] = nums2[p2]
			p2--
		} else if p2 == -1 {
			nums1[index] = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			nums1[index] = nums1[p1]
			p1--
		} else {
			nums1[index] = nums2[p2]
			p2--
		}
	}
}
