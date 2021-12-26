package num_trees

func numTrees(n int) int {
	result := []int{1, 2, 5, 14, 42, 132, 429, 1430, 4862, 16796, 58786, 208012, 742900, 2674440, 9694845, 35357670, 129644790, 477638700, 1767263190}
	return result[n-1]
}
