package main

// 1145
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	l, r := 0, 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lv := dfs(node.Left)
		rv := dfs(node.Right)
		if node.Val == x {
			l, r = lv, rv
		}
		return lv + rv + 1
	}
	tot := dfs(root)
	p := tot - l - r - 1
	if l > tot-l {
		return true
	}
	if r > tot-r {
		return true
	}
	if p > tot-p {
		return true
	}
	return false
}
