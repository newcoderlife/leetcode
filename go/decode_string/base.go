package decode_string

import "strings"

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func findPattern(s string, start int) (int, string, int) {
	if isDigit(s[start]) {
		num := 0
		current := start
		for current < len(s) && isDigit(s[current]) {
			num = num*10 + int(s[current]-'0')
			current++
		}

		head := current + 1
		count := 1
		for current = head + 1; current < len(s); current++ {
			if s[current] == '[' {
				count++
			} else if s[current] == ']' {
				count--
			}
			if count == 0 {
				break
			}
		}
		return num, s[head:current], current + 1
	} else {
		current := start
		for current < len(s) && !isDigit(s[current]) {
			current++
		}
		return 1, s[start:current], current
	}
}

func decodeString(s string) string {
	if !strings.Contains(s, "[") {
		return s
	}

	result := new(strings.Builder)
	for current := 0; current < len(s); {
		repeat, content, next := findPattern(s, current)
		substr := decodeString(content)
		for i := 0; i < repeat; i++ {
			result.WriteString(substr)
		}
		current = next
	}
	return result.String()
}
