package group_anagrams

import (
	"container/list"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	exist := make(map[string]*list.List)
	for _, str := range strs {
		buffer := []byte(str)
		sort.Slice(buffer, func(i, j int) bool {
			return buffer[i] < buffer[j]
		})
		sorted := string(buffer)

		if exist[sorted] == nil {
			exist[sorted] = list.New()
		}
		exist[sorted].PushBack(str)
	}

	results := make([][]string, 0)
	for _, l := range exist {
		result := make([]string, 0)
		for current := l.Front(); current != nil; current = current.Next() {
			result = append(result, current.Value.(string))
		}
		results = append(results, result)
	}
	return results
}
