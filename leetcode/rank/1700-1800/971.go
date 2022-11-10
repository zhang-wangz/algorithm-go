package main

// 971
func flipMatchVoyage(root *TreeNode, voyage []int) (ans []int) {
	var dfs func(node *TreeNode)
	idx := 0
	dfs = func(node *TreeNode) {
		if node != nil {
			if node.Val != voyage[idx] {
				ans = []int{-1}
				return
			}
			idx++
			if idx < len(voyage) && node.Left != nil && node.Left.Val != voyage[idx] {
				ans = append(ans, node.Val)
				dfs(node.Right)
				dfs(node.Left)
			} else {
				dfs(node.Left)
				dfs(node.Right)
			}
		}
	}
	dfs(root)
	if len(ans) == 0 || ans[0] == -1 {
		ans = []int{-1}
	}
	return ans
}
