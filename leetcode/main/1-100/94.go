package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// l m r
func inorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return []int{}
	}
	st := []*TreeNode{}
	for len(st) != 0 || root != nil {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}
		node := st[len(st)-1]
		st = st[:len(st)-1]
		ans = append(ans, node.Val)
		root = node.Right
	}
	return
}
