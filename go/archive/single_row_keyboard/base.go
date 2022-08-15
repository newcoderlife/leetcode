package single_row_keyboard

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func calculateTime(keyboard string, word string) (ans int) {
	index := make(map[int]int)
	for i, c := range keyboard {
		index[int(c)] = i
	}

	current := 0
	for _, c := range word {
		ans += abs(index[int(c)] - current)
		current = index[int(c)]
	}
	return ans
}
