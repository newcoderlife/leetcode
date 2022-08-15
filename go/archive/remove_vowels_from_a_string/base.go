package remove_vowels_from_a_string

import "strings"

func removeVowels(s string) string {
	result := new(strings.Builder)
	for _, c := range s {
		if c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u' {
			result.WriteRune(c)
		}
	}
	return result.String()
}
