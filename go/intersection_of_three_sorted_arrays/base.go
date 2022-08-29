package intersection_of_three_sorted_arrays

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	results := make([]int, 0)

	for i1, i2, i3 := 0, 0, 0; i1 < len(arr1); i1++ {
		for {
			if i2 == len(arr2) {
				return results
			}
			if arr2[i2] >= arr1[i1] {
				break
			}
			i2++
		}
		for {
			if i3 == len(arr3) {
				return results
			}
			if arr3[i3] >= arr1[i1] {
				break
			}
			i3++
		}

		if arr1[i1] == arr2[i2] && arr1[i1] == arr3[i3] {
			results = append(results, arr1[i1])
		}
	}
	return results
}
