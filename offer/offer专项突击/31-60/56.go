package _1_60

func findTarget(root *TreeNode, k int) bool {
	set := map[int]struct{}{}
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, t int) bool {
		if node == nil {
			return false
		}
		if _, ok := set[t-node.Val]; ok {
			return true
		}
		set[node.Val] = struct{}{}
		return dfs(node.Left, t) || dfs(node.Right, t)
	}
	return dfs(root, k)
}
