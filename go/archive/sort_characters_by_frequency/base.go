package sort_characters_by_frequency

import (
	"sort"
	"strings"
)

type Item struct {
	Value byte
	Count int
}

func frequencySort(s string) string {
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	items := make([]*Item, 0, len(count))
	for value, count := range count {
		items = append(items, &Item{Value: value, Count: count})
	}
	sort.Slice(items, func(i, j int) bool { return items[i].Count > items[j].Count })

	result := new(strings.Builder)
	for _, item := range items {
		for i := 0; i < item.Count; i++ {
			result.WriteByte(item.Value)
		}
	}
	return result.String()
}
