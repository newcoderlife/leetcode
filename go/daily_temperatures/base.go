package daily_temperatures

import "container/list"

type Node struct {
	Index       int
	Temperature int
}

func dailyTemperatures(temperatures []int) []int {
	stack := list.New()
	result := make([]int, len(temperatures))
	for index, temperature := range temperatures {
		for {
			back := stack.Back()
			if back == nil {
				break
			}

			node := back.Value.(*Node)
			if node.Temperature < temperature {
				result[node.Index] = index - node.Index
				stack.Remove(back)
			} else {
				break
			}
		}
		stack.PushBack(&Node{Index: index, Temperature: temperature})
	}
	return result
}
