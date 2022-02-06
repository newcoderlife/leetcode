package generate_parenthesis

var (
	result []string
)

func dfs(add, minus int, current string) {
	if add > minus { // invalid
		return
	}
	if add == 0 {
		for index := 0; index < minus; index++ {
			current += ")"
		}
		result = append(result, current)
		return
	}
	dfs(add-1, minus, current+"(")
	dfs(add, minus-1, current+")")
}

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	dfs(n, n, "")
	return result
}
