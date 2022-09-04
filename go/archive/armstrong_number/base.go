package armstrong_number

func isArmstrong(n int) bool {
	k := 0
	for num := n; num > 0; k, num = k+1, num/10 {
	}

	result := 0
	for i, num := 0, n; i < k; num, i = num/10, i+1 {
		count := 1
		for j, digit := 0, num%10; j < k; j++ {
			count *= digit
		}
		result += count
	}

	return result == n
}
