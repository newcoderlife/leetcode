package longest_valid_parentheses

import (
	"container/list"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {
	stack := list.New()
	stack.PushBack(-1)
	result := 0
	for index, char := range s {
		if char == '(' {
			stack.PushBack(index)
		} else {
			stack.Remove(stack.Back())
			if stack.Len() == 0 {
				stack.PushBack(index)
			} else {
				result = max(result, index-stack.Back().Value.(int))
			}
		}
	}

	return result
}

func longestValidParentheses2(s string) int {
	result := 0
	for left, right, index := 0, 0, 0; index < len(s); index++ {
		if s[index] == '(' {
			left++
		} else if s[index] == ')' {
			right++
		}

		if right == left {
			result = max(result, left*2)
		} else if right > left {
			left, right = 0, 0
		}
	}

	for left, right, index := 0, 0, len(s)-1; index >= 0; index-- {
		if s[index] == '(' {
			left++
		} else if s[index] == ')' {
			right++
		}

		if right == left {
			result = max(result, left*2)
		} else if right < left {
			left, right = 0, 0
		}
	}

	return result
}
