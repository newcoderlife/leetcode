package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	for last := root; last != nil; last = last.Right {
		if last.Left != nil {
			var leftTail *TreeNode
			for leftTail = last.Left; leftTail.Right != nil; leftTail = leftTail.Right {
			}
			leftTail.Right = last.Right
			last.Right = last.Left
			last.Left = nil
		}
	}
}
