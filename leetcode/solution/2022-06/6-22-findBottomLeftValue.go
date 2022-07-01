package _022_06

import "algorithm-go/binaryTree"

// https://leetcode.cn/problems/find-bottom-left-tree-value/
// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
// 写个层次遍历
func findBottomLeftValue(root *binaryTree.Node) int {
	queue := make([]*binaryTree.Node, 0)
	queue = append(queue, root)
	var ans int
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		ans = node.Val
	}
	return ans
}
