package group_anagrams

import (
	"container/list"
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	exist := make(map[string]*list.List)
	for _, str := range strs {
		hash := make([]int, 26)
		for _, c := range str {
			hash[c-'a'] += 1
		}

		var hashCode string
		for index, num := range hash {
			if num != 0 {
				hashCode += fmt.Sprintf("%c%d", 'a'+index, num)
			}
		}

		if exist[hashCode] == nil {
			exist[hashCode] = list.New()
		}
		exist[hashCode].PushBack(str)
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
