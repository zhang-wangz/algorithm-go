package main

import "fmt"

// dfs
func findDuplicateSubtrees(root *TreeNode) (ans []*TreeNode) {
	repeat := map[*TreeNode]struct{}{}
	vis := map[string]*TreeNode{}
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		serial := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
		if n, ok := vis[serial]; ok {
			repeat[n] = struct{}{}
		} else {
			vis[serial] = node
		}
		return serial
	}
	dfs(root)
	for node := range repeat {
		ans = append(ans, node)
	}
	return
}
