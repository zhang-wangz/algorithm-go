package _022_08

import (
	"strconv"
)

func printTree(root *TreeNode) [][]string {
	var getH func(node *TreeNode) int
	getH = func(node *TreeNode) int {
		h := 0
		if node.Left != nil {
			h = getH(node.Left) + 1
		}
		if node.Right != nil {
			h = max(h, getH(node.Right)+1)
		}
		return h
	}
	h := getH(root)
	m := h + 1
	n := 1<<m + 1
	ans := make([][]string, m)
	for i := range ans {
		ans[i] = make([]string, n)
	}
	var dfs func(node *TreeNode, r, c int)
	dfs = func(node *TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfs(node.Left, r+1, c-1<<(h-r-1))
		}
		if node.Right != nil {
			dfs(node.Right, r+1, c+1<<(h-r-1))
		}
	}
	dfs(root, 0, (n-1)/2)
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
