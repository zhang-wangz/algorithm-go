package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) (ans int) {
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				ans += node.Left.Val
			} else {
				ans += dfs(node.Left)
			}
		}
		if node.Right != nil && !(node.Right.Left == nil && node.Right.Right == nil) {
			ans += dfs(node.Right)
		}
		return
	}
	return dfs(root)
}
