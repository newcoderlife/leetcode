package word_break

import "strings"

var (
	candidates []string
	memory     map[string]bool
)

func dfs(current string) bool {
	if len(current) == 0 {
		return true
	}
	if result, ok := memory[current]; ok {
		return result
	}

	for _, candidate := range candidates {
		if strings.HasPrefix(current, candidate) {
			if dfs(strings.TrimPrefix(current, candidate)) {
				memory[current] = true
				return true
			}
		}
	}
	memory[current] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	candidates = wordDict
	memory = make(map[string]bool)
	return dfs(s)
}
