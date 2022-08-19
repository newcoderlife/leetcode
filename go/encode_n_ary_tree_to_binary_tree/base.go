package encode_n_ary_tree_to_binary_tree

type Node struct {
	Val      int
	Children []*Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() *Codec { return nil }

func (this *Codec) encodeDFS(current *Node) *TreeNode {
	if current == nil || len(current.Children) == 0 {
		return nil
	}

	result := new(TreeNode)
	now := result
	for _, child := range current.Children {
		now.Right = this.encode(child)
		now = now.Right
	}
	return result.Right
}

func (this *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return nil
	}

	return &TreeNode{Val: root.Val, Left: this.encodeDFS(root)}
}

func (this *Codec) decodeDFS(root *TreeNode) []*Node {
	if root == nil {
		return nil
	}

	results := make([]*Node, 0)
	for root != nil {
		results = append(results, this.decode(root))
		root = root.Right
	}
	return results
}

func (this *Codec) decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}

	return &Node{Val: root.Val, Children: this.decodeDFS(root.Left)}
}
