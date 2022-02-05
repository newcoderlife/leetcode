package decode_string

import (
	"fmt"
	"testing"
)

func Test_decodeString(t *testing.T) {
	fmt.Println(decodeString("3[a2[c]]") == "accaccacc")
	fmt.Println(decodeString("leetcode") == "leetcode")
	fmt.Println(decodeString("3[a]2[bc]") == "aaabcbc")
	fmt.Println(decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef")
	fmt.Println(decodeString("abc3[cd]xyz") == "abccdcdcdxyz")
}
