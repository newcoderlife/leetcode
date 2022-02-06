package trie

type TreeNode struct {
	children map[int32]*TreeNode
	end      bool
}

type Trie struct {
	root *TreeNode
}

func Constructor() Trie {
	return Trie{root: &TreeNode{children: make(map[int32]*TreeNode)}}
}

func (this *Trie) Insert(word string) {
	current := this.root
	for _, b := range word {
		if current.children[b-'a'] == nil {
			current.children[b-'a'] = &TreeNode{
				children: make(map[int32]*TreeNode),
			}
		}
		current = current.children[b-'a']
	}
	current.end = true
}

func (this *Trie) Search(word string) bool {
	current := this.root
	for _, b := range word {
		next := current.children[b-'a']
		if next == nil {
			return false
		}
		current = next
	}
	return current.end
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root
	for _, b := range prefix {
		next := current.children[b-'a']
		if next == nil {
			return false
		}
		current = next
	}
	return true
}
