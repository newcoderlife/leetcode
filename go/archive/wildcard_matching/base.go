package wildcard_matching

import (
	"fmt"
	"strings"
)

var (
	memory = make(map[string]bool)
)

func isMatch(s string, p string) (result bool) {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 {
		return p == "*" || (p == "?" && len(s) == 1) || (p == s)
	}

	for strings.Contains(p, "**") {
		p = strings.ReplaceAll(p, "**", "*")
	}

	key := fmt.Sprintf("%s_%s", s, p)
	if result, ok := memory[key]; ok {
		return result
	}
	defer func() { memory[key] = result }()

	switch p[0] {
	case '*':
		pure := len(strings.ReplaceAll(p, "*", ""))
		for i := 0; i <= len(s) && i+pure <= len(s); i++ {
			if result = result || isMatch(s[i:], p[1:]); result {
				return true
			}
		}

		return false
	default:
		if len(s) == 0 {
			return false
		}

		index := 0
		for ; index < len(s) && index < len(p); index++ {
			if p[index] == '*' {
				break
			}

			if s[index] != p[index] && p[index] != '?' {
				return false
			}
		}
		return isMatch(s[index:], p[index:])
	}
}
