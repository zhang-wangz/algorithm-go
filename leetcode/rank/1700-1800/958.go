package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 958 判断是否是完全二叉树
// 层次遍历，如果遇到空的就把标志位设为true，然后继续判断如果有值就设置为false
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	f := false

	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		// 如果空，则置为true
		if n == nil {
			f = true
			continue
		} else {
			if f {
				return false
			}
		}
		q = append(q, n.Left)
		q = append(q, n.Right)
	}
	return true
}
