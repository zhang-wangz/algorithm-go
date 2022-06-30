package _1_60

func findBottomLeftValue(root *TreeNode) int {
	curheight := 0
	val := -1
	var dfs = func(root *TreeNode, height int) {}
	dfs = func(root *TreeNode, height int) {
		if root == nil {
			return
		}
		height++
		dfs(root.Left, height)
		dfs(root.Right, height)
		if curheight < height {
			curheight = height
			val = root.Val
		}
	}
	dfs(root, 0)
	return val
}
