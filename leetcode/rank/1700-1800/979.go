package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 先把移动剩下的移动到根节点，然后进行转移
func distributeCoins(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		ans += int(math.Abs(float64(l)) + math.Abs(float64(r)))
		return node.Val + l + r - 1
	}
	dfs(root)
	return
}
