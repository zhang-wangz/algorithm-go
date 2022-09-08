package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 树形dp 学习一下
func longestZigZag(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		_, r1 := dfs(node.Left)
		l2, _ := dfs(node.Right)
		L, R := 0, 0
		if node.Left != nil {
			L = r1 + 1
		}
		if node.Right != nil {
			R = l2 + 1
		}
		ans = max(ans, max(L, R))
		return L, R
	}
	dfs(root)
	return
}
