package binaryTree

// 深度搜索从上到下
func PreDfs(root *Node) []int {
	res := make([]int, 0)
	dfs(root, &res)
	return res
}

func dfs(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	dfs(root.Left, res)
	dfs(root.Right, res)
}
