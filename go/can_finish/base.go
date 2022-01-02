package can_finish

import "container/list"

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	inCap := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		edges[prerequisite[1]] = append(edges[prerequisite[1]], prerequisite[0])
		inCap[prerequisite[0]]++
	}

	queue := list.New()
	for i := 0; i < numCourses; i++ {
		if inCap[i] == 0 {
			queue.PushBack(i)
		}
	}
	for current := queue.Front(); current != nil; current = current.Next() {
		for _, end := range edges[current.Value.(int)] {
			inCap[end]--
			if inCap[end] == 0 {
				queue.PushBack(end)
			}
		}
	}

	return queue.Len() == numCourses
}
