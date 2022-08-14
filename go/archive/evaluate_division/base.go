package evaluate_division

type Edge struct {
	To    string
	Value float64
}

type Node struct {
	Value string
	Next  []*Edge
}

var (
	graph  map[string]*Node
	target string
	exist  map[string]bool
)

func dfs(current string) float64 {
	exist[current] = true
	node := graph[current]
	if node == nil {
		return 0
	}
	if node.Value == target {
		return 1
	}

	for _, edge := range node.Next {
		if !exist[edge.To] {
			if result := dfs(edge.To); result != 0 {
				return result * edge.Value
			}
		}
	}

	return 0
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph = make(map[string]*Node)
	for index, equation := range equations {
		node := graph[equation[0]]
		if node == nil {
			node = &Node{Value: equation[0], Next: make([]*Edge, 0)}
			graph[equation[0]] = node
		}
		node.Next = append(node.Next, &Edge{To: equation[1], Value: values[index]})

		node = graph[equation[1]]
		if node == nil {
			node = &Node{Value: equation[1], Next: make([]*Edge, 0)}
			graph[equation[1]] = node
		}
		node.Next = append(node.Next, &Edge{To: equation[0], Value: 1 / values[index]})
	}

	results := make([]float64, 0, len(queries))
	for _, query := range queries {
		target = query[1]
		exist = make(map[string]bool)
		result := dfs(query[0])

		if result == 0 {
			results = append(results, -1)
		} else {
			results = append(results, result)
		}
	}
	return results
}
