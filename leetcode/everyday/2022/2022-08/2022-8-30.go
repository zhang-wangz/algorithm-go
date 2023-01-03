package _022_08

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val {
		return &TreeNode{val, root, nil}
	}
	var p *TreeNode
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			if p == nil {
				return &TreeNode{val, cur, nil}
			}
			p.Right = &TreeNode{val, cur, nil}
			return root
		}
		p = cur
	}
	p.Right = &TreeNode{Val: val}
	return root
}
