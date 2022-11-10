package solution

import "algorithm-go/base/binaryTree"

// 返回二叉树的镜像树
func ReserveTree(root *binaryTree.Node) *binaryTree.Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		ReserveTree(root.Left)
	}
	if root.Right != nil {
		ReserveTree(root.Right)
	}
	return root
}
func ReserveTree2(root *binaryTree.Node) *binaryTree.Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	stack := make([]*binaryTree.Node, 0)
	res := root
	var preNode *binaryTree.Node
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if node.Right == nil || preNode == node.Right {
			node.Left, node.Right = node.Right, node.Left
			stack = stack[:len(stack)-1]
			preNode = node
		} else {
			root = node.Right
		}
	}
	return res
}
