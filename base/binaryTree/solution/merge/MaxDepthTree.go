package merge

import "algorithm-go/base/binaryTree"

// 分治法
func MaxDepth(root *binaryTree.Node) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}
