package reverse_words_in_a_string_ii

import (
	"strings"
)

func reverseWords(s []byte) {
	words := strings.Split(string(s), " ")
	reverse := make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		reverse = append(reverse, words[len(words)-1-i])
	}
	copy(s, []byte(strings.Join(reverse, " ")))
}

func reverse(s []byte, start, end int) {
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseWords1(s []byte) {
	reverse(s, 0, len(s))

	start, current := 0, 0
	for current < len(s) {
		if s[current] == ' ' {
			reverse(s, start, current)
			start = current + 1
		}
		current++
	}
	reverse(s, start, current)
}
