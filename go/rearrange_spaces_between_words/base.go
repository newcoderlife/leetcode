package rearrange_spaces_between_words

import (
	"strings"
)

func reorderSpaces(text string) string {
	words := strings.Fields(text)
	spaces := 0
	for index := 0; index < len(text); index++ {
		if text[index] == ' ' {
			spaces++
		}
	}

	join := make([]byte, 0)
	if len(words) > 1 {
		for i := 0; i < spaces/(len(words)-1); i++ {
			join = append(join, ' ')
		}
	}

	results := []byte(strings.Join(words, string(join)))
	for len(results) < len(text) {
		results = append(results, ' ')
	}
	return string(results)
}
