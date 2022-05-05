package merge

import "algorithm-pattern/data_structure/binaryTree"

func IsBalanced(root *binaryTree.Node) bool {
	if root == nil {
		return true
	}
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

// 分治
func maxDepth(root *binaryTree.Node) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
