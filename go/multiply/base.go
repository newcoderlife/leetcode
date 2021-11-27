package multiply

func getNum(num string, index int) uint8 {
	return num[len(num)-index-1] - '0'
}

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return "0"
	}

	buffer := make([]byte, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			a, b := getNum(num1, i), getNum(num2, j)
			buffer[i+j] += a * b
			buffer[i+j+1] += buffer[i+j] / 10
			buffer[i+j] %= 10
		}
	}

	last := -1
	result := make([]byte, len(buffer))
	for index := 0; index < len(buffer); index++ {
		if buffer[index] != 0 {
			last = index
		}
		result[len(buffer)-index-1] = buffer[index] + '0'
	}
	if last == -1 {
		return "0"
	}
	return string(result[len(result)-last-1:])
}
