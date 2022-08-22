package reorganize_string

import "sort"

type Item struct {
	Value int
	Count int
}

func reorganizeString(s string) string {
	counts := make([]int, 26)
	for _, c := range s {
		counts[int(c)-int('a')]++
	}
	items := make([]*Item, 0, 26)
	for index, count := range counts {
		if count > 0 {
			items = append(items, &Item{Value: index, Count: count})
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Count > items[j].Count
	})

	if items[0].Count*2-1 > len(s) {
		return ""
	}

	results := make([]byte, len(s))
	for i, index := 0, 0; i < len(items); i++ {
		for items[i].Count > 0 {
			results[index] = byte('a' + items[i].Value)
			items[i].Count--
			index += 2
			if index >= len(s) {
				index = 1
			}
		}
	}

	return string(results)
}
