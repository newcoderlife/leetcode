package find_order

import (
	"container/list"
)

var (
	result *list.List
	visit  []int
	path   map[int]*list.List
)

func dfs(current int) bool {
	if visit[current] == 1 {
		return false
	}
	if visit[current] == 2 {
		return true
	}

	visit[current] = 1
	if to := path[current]; to != nil {
		for current := to.Front(); current != nil; current = current.Next() {
			if !dfs(current.Value.(int)) {
				return false
			}
		}
	}

	visit[current] = 2
	result.PushBack(current)
	return true
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	result = list.New()
	visit = make([]int, numCourses)
	path = make(map[int]*list.List)
	for _, prerequisite := range prerequisites {
		if path[prerequisite[1]] == nil {
			path[prerequisite[1]] = list.New()
		}
		path[prerequisite[1]].PushBack(prerequisite[0])
	}

	for start := range path {
		if !dfs(start) {
			return nil
		}
	}

	content := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if visit[i] == 0 {
			content = append(content, i)
		}
	}
	for current := result.Back(); current != nil; current = current.Prev() {
		content = append(content, current.Value.(int))
	}
	return content
}
