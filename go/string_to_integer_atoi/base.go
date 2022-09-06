package string_to_integer_atoi

import "math"

func myAtoi(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '-' {
			result := 0
			for j := i + 1; j < len(s); j++ {
				if s[j] < '0' || '9' < s[j] {
					break
				}
				result = result*10 + int(s[j]-'0')
				if -result < math.MinInt32 {
					return math.MinInt32
				}
			}
			return -result
		}
		if s[i] == '+' {
			result := 0
			for j := i + 1; j < len(s); j++ {
				if s[j] < '0' || '9' < s[j] {
					break
				}
				result = result*10 + int(s[j]-'0')
				if result > math.MaxInt32 {
					return math.MaxInt32
				}
			}
			return result
		}
		if '0' <= s[i] && s[i] <= '9' {
			result := 0
			for j := i; j < len(s); j++ {
				if s[j] < '0' || '9' < s[j] {
					break
				}
				result = result*10 + int(s[j]-'0')
				if result > math.MaxInt32 {
					return math.MaxInt32
				}
			}
			return result
		}

		break
	}

	return 0
}
