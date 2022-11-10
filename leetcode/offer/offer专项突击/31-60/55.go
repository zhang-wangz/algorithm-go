package _1_60

type BSTIterator struct {
	inorder []*TreeNode
	root    *TreeNode
	cur     *TreeNode
}

// Constructor8 Constructor l i r
func Constructor8(root *TreeNode) BSTIterator {
	order := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)
	head := root
	if root == nil {
		return BSTIterator{inorder: order, root: root}
	}
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, node)
		root = node.Right
	}
	return BSTIterator{
		inorder: order,
		root:    head,
	}
}

func (this *BSTIterator) Next() int {
	n := this.inorder[0].Val
	this.inorder = this.inorder[1:]
	return n
}

func (this *BSTIterator) HasNext() bool {
	if len(this.inorder) != 0 {
		return true
	}
	return false
}
