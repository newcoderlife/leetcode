package build_tree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 1 && len(inorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:i+1], inorder[:i]),
				Right: buildTree(preorder[i+1:], inorder[i+1:]),
			}
		}
	}

	return nil
}
