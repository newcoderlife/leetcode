package trailing_zeroes

func trailingZeroes(n int) int {
	count2 := 0
	count5 := 0
	count10 := 0
	for i := 2; i <= n; i++ {
		num := i
		for ; num%10 == 0; num /= 10 {
			count10++
		}
		for ; num%2 == 0; num /= 2 {
			count2++
		}
		for ; num%5 == 0; num /= 5 {
			count5++
		}
	}

	if count2 < count5 {
		return count2 + count10
	}
	return count5 + count10
}
