package solution

import "algorithm-go/binaryTree"

func insertIntoBST(root *binaryTree.Node, val int) *binaryTree.Node {
	if root == nil {
		return &binaryTree.Node{Val: val}
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
