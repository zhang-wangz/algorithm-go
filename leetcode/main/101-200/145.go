package main

// l r m
func postorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return []int{}
	}
	st := []*TreeNode{}
	for len(st) != 0 || root != nil {
		for root != nil {
			st = append(st, root)
			root = root.Left
		}
		var last *TreeNode
		node := st[len(st)-1]
		if node.Right == nil || last == node.Right {
			ans = append(ans, node.Val)
			st = st[:len(st)-1]
			last = node
		} else {
			root = node.Right
		}
	}
	return
}
