package remove_invalid_parentheses

var (
	target  string
	results map[string]bool
	longest int
)

func correct(s []byte) bool {
	cnt := 0
	for _, char := range s {
		if char == '(' {
			cnt++
		} else if char == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
}

// func IfThen[T any](cond bool, v1, v2 T) T {
// 	if cond {
// 		return v1
// 	}
// 	return v2
// }

func IfThen(cond bool, v1, v2 int) int {
	if cond {
		return v1
	}
	return v2
}

func dfs(current []byte, index int, left, right int) {
	if left < right {
		return
	}
	if left > right+len(target)-index {
		return
	}

	if index == len(target) {
		if correct(current) {
			if len(current) >= longest {
				results[string(current)] = true
				longest = len(current)
			}
		}

		return
	}

	dfs(append(current, target[index]), index+1, left+IfThen(target[index] == '(', 1, 0), right+IfThen(target[index] == ')', 1, 0))
	dfs(current, index+1, left, right)
}

func removeInvalidParentheses(s string) []string {
	results = make(map[string]bool)
	target = s
	longest = 0

	dfs([]byte{}, 0, 0, 0)

	result := make([]string, 0)
	for key := range results {
		if len(key) == longest {
			result = append(result, key)
		}
	}
	return result
}
