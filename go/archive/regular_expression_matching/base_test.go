package regular_expression_matching

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("aaa", "aa"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("a", "ab"))
	fmt.Println(isMatch("mississippi", "mis*is*p*."))
	fmt.Println(isMatch("abcaaaaaaabaabcabac", ".*ab.a.*a*a*.*b*b*"))
	fmt.Println(isMatch("abcabac", "ab.a.*a*a*.*b*b*"))
	fmt.Println(isMatch("bac", ".*b*b*"))
	fmt.Println(isMatch("", "b*b*"))
}
