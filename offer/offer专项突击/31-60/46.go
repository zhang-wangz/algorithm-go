package _1_60

func RightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	curhight := 0
	if root == nil {
		return ans
	}
	var dfs func(root *TreeNode, height int)
	dfs = func(root *TreeNode, height int) {
		if root == nil {
			return
		}
		height++
		if curhight < height {
			curhight = height
			ans = append(ans, root.Val)
		}
		dfs(root.Right, height)
		dfs(root.Left, height)

	}
	dfs(root, 0)
	return ans
}
