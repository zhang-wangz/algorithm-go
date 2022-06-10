package bfs

import "algorithm-go/binaryTree"

// 从底向上遍历
func LevelOrderReverse(root *binaryTree.Node) [][]int {
	res := make([][]int, 0)
	res = binaryTree.LevelOrder(root)
	reverse(res)
	return res
}

func reverse(num [][]int) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}
