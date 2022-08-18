package main

func deepestLeavesSum(root *TreeNode) (sum int) {
	height := -1
	var dfs func(curH int, node *TreeNode)
	dfs = func(curH int, node *TreeNode) {
		if node == nil {
			return
		}
		if curH > height {
			height = curH
			sum = node.Val
		} else if curH == height {
			sum += node.Val
		}
		dfs(curH+1, node.Left)
		dfs(curH+1, node.Right)
	}
	dfs(0, root)
	return sum
}
