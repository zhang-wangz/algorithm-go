package binary_search

import (
	"algorithm-go/base/binaryTree"
)

// bfs 判断中序是否有序就行
func IsBinaryTree(root *binaryTree.Node) bool {
	if root == nil {
		return true
	}
	res := binaryTree.InOrder(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			return false
		}
	}
	return true
}

func InOrder(root *binaryTree.Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*binaryTree.Node, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}
