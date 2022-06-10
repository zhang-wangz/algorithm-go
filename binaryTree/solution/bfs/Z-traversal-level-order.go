package bfs

import "algorithm-go/binaryTree"

func ZLevelOrder(root *binaryTree.Node) [][]int {
	res := make([][]int, 0)
	queue := make([]*binaryTree.Node, 0)
	tag := false
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		list := make([]int, 0)
		list = append(list, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if tag {
			reverseLst(list)
		}
		res = append(res, list)
		tag = !tag
	}
	return res
}
func reverseLst(num []int) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}
