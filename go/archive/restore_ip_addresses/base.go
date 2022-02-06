package restore_ip_addresses

import (
	"fmt"
	"strconv"
)

func judge(a, b, c, d string) string {
	if (a != "0" && a[0] == '0') ||
		(b != "0" && b[0] == '0') ||
		(c != "0" && c[0] == '0') ||
		(d != "0" && d[0] == '0') {
		return ""
	}
	if num, _ := strconv.ParseInt(a, 10, 64); num > 255 {
		return ""
	}
	if num, _ := strconv.ParseInt(b, 10, 64); num > 255 {
		return ""
	}
	if num, _ := strconv.ParseInt(c, 10, 64); num > 255 {
		return ""
	}
	if num, _ := strconv.ParseInt(d, 10, 64); num > 255 {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s.%s", a, b, c, d)
}

func restoreIpAddresses(s string) []string {
	results := make([]string, 10)

	for i := 1; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			for k := j + 1; k < len(s); k++ {
				if result := judge(s[:i], s[i:j], s[j:k], s[k:]); len(result) > 0 {
					results = append(results, result)
				}
			}
		}
	}

	return results
}
