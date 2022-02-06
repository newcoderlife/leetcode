package is_valid

import "container/list"

func isLeft(c byte) bool {
	return c == '[' || c == '{' || c == '('
}

func getLeft(c byte) byte {
	if c == ')' {
		return '('
	}
	if c == ']' {
		return '['
	}
	if c == '}' {
		return '{'
	}
	panic("invalid request")
}

func isValid(s string) bool {
	stack := list.New()
	for index := 0; index < len(s); index++ {
		c := s[index]
		if isLeft(c) {
			stack.PushBack(c)
			continue
		}

		if stack.Back() == nil || stack.Back().Value.(byte) != getLeft(c) {
			return false
		}
		stack.Remove(stack.Back())
	}
	return stack.Len() == 0
}
