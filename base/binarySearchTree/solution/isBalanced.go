package solution

import (
	"algorithm-go/base/binaryTree"
	"math"
)

// https://leetcode.cn/problems/balanced-binary-tree/

type R struct {
	valid bool
	h     int
}

func isBalanced(root *binaryTree.Node) bool {
	return helperB(root).valid
}
func helperB(root *binaryTree.Node) (res R) {
	if root == nil {
		res.valid = true
		res.h = 0
		return
	}
	left := helperB(root.Left)
	right := helperB(root.Right)
	if left.valid && right.valid && math.Abs(float64(left.h)-float64(right.h)) <= 1 {
		res.valid = true
	}
	res.h = max(left.h, right.h) + 1
	return
}
