package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// m l r
func preorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return []int{}
	}
	st := []*TreeNode{}
	for len(st) != 0 || root != nil {
		for root != nil {
			st = append(st, root)
			ans = append(ans, root.Val)
			root = root.Left
		}
		node := st[len(st)-1]
		st = st[:len(st)-1]
		root = node.Right
	}
	return
}
