package exist

var (
	matrix [][]byte
	visit  [][]bool
	target string
)

func dfs(i, j, index int) bool {
	if index+1 == len(target) {
		return true
	}

	visit[i][j] = true
	defer func() { visit[i][j] = false }()

	if i > 0 && !visit[i-1][j] && matrix[i-1][j] == target[index+1] {
		if dfs(i-1, j, index+1) {
			return true
		}
	}
	if j > 0 && !visit[i][j-1] && matrix[i][j-1] == target[index+1] {
		if dfs(i, j-1, index+1) {
			return true
		}
	}
	if i+1 < len(matrix) && !visit[i+1][j] && matrix[i+1][j] == target[index+1] {
		if dfs(i+1, j, index+1) {
			return true
		}
	}
	if j+1 < len(matrix[0]) && !visit[i][j+1] && matrix[i][j+1] == target[index+1] {
		if dfs(i, j+1, index+1) {
			return true
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	target = word
	matrix = board
	visit = make([][]bool, len(board))
	for i := 0; i < len(visit); i++ {
		visit[i] = make([]bool, len(matrix[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
