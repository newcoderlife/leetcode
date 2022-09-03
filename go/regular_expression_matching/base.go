package regular_expression_matching

import "fmt"

var (
	memory = make(map[string]bool)
)

func getNextPattern(p string) string {
	if len(p) < 2 {
		return p
	}

	if p[1] == '*' {
		return p[:2]
	}
	return p[:1]
}

func match(s, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(p) == 1 {
		return len(s) == 1 && (s[0] == p[0] || p[0] == '.')
	}

	for i := 0; i < len(s); i++ {
		if s[i] != p[0] && p[0] != '.' {
			return false
		}
	}
	return true
}

func isMatch(s string, p string) (result bool) {
	key := fmt.Sprintf("%s_%s", s, p)
	if result, ok := memory[key]; ok {
		return result
	}

	pattern := getNextPattern(p)
	if pattern == p {
		result = match(s, p)
		memory[key] = result
		return result
	}

	start := 0
	if len(pattern) == 1 {
		start = 1
	}

	for i := start; i <= len(s); i++ {
		if !match(s[:i], pattern) {
			break
		}
		if result = result || isMatch(s[i:], p[len(pattern):]); result {
			memory[key] = true
			return true
		}
	}

	memory[key] = false
	return false
}
