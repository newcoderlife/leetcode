package convert_to_title

func convertToTitle(columnNumber int) string {
	result := ""
	for columnNumber > 0 {
		columnNumber--
		result = string(rune('A'+columnNumber%26))+result
		columnNumber/=26
	}
	return result
}
