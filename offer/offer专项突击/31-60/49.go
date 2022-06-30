package _1_60

func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(root *TreeNode, s int)
	dfs = func(root *TreeNode, s int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			sum = sum + s*10 + root.Val
			return
		}
		dfs(root.Left, s*10+root.Val)
		dfs(root.Right, s*10+root.Val)
	}
	dfs(root, 0)
	return sum
}
