package simplify_path

import (
	"container/list"
	"strings"
)

func simplifyPath(path string) string {
	buffer := list.New()
	commands := strings.Split(path, "/")
	for _, command := range commands {
		content := strings.ReplaceAll(command, "/", "")
		if len(content) == 0 || content == "." { // 有可能是 '//' 等情况
			continue
		}
		if content == ".." {
			if buffer.Len() > 0 {
				buffer.Remove(buffer.Back())
			}
			continue
		}
		buffer.PushBack(content)
	}

	var result string
	for current := buffer.Front(); current != nil; current = current.Next() {
		result += "/" + current.Value.(string)
	}

	if len(result) == 0 {
		return "/"
	}
	return result
}
