package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root != nil && root.Val != val {
			return false
		}
		if root == nil {
			return true
		}
		return dfs(root.Left) && dfs(root.Right)
	}
	return dfs(root)
}
