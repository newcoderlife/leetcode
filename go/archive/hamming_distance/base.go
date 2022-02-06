package hamming_distance

func hammingDistance(x int, y int) int {
	count := 0
	for delta := x ^ y; delta > 0; delta >>= 1 {
		if delta&1 == 1 {
			count++
		}
	}
	return count
}
