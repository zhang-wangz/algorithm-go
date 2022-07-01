package _022_06

import (
	"algorithm-go/binaryTree"
	"math"
)

// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
// 给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
func largestValues(root *binaryTree.Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	q := make([]*binaryTree.Node, 0)
	q = append(q, root)
	for len(q) != 0 {
		le := len(q)
		t := -1 << 63
		for i := 0; i < le; i++ {
			node := q[0]
			q = q[1:]
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			t = int(math.Max(float64(t), float64(node.Val)))
		}
		ans = append(ans, t)
	}
	return ans
}
