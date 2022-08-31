package add_digits

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	total := 0
	for num > 0 {
		total += num % 10
		num /= 10
	}

	return addDigits(total)
}
