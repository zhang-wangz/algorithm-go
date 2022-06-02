package solution

import (
	"algorithm-pattern/data_structure/binaryTree"
)

// https://leetcode.cn/problems/validate-binary-search-tree/

type Res struct {
	max, min int
	valid    bool
}

func isValidBST(root *binaryTree.Node) bool {
	if root == nil {
		return true
	}
	return dfs(root).valid
}

func dfs(root *binaryTree.Node) (res Res) {
	if root == nil {
		res.max = -1 << 63
		res.min = 1<<63 - 1
		res.valid = true
		return
	}
	left := dfs(root.Left)
	right := dfs(root.Right)

	if left.max < root.Val && right.min > root.Val && left.valid && right.valid {
		res.valid = true
	}
	res.max = max(right.max, root.Val)
	res.min = min(left.min, root.Val)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
