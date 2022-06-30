package _1_60

import "math"

func MaxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ml := -1 << 63
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left < 0 {
			left = 0
		}
		if right < 0 {
			right = 0
		}
		sum := left + right + node.Val
		ml = int(math.Max(float64(ml), float64(sum)))
		return node.Val + int(math.Max(float64(left), float64(right)))
	}
	dfs(root)
	return ml
}
