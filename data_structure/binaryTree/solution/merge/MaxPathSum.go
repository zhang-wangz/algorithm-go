package merge

import (
	"algorithm-pattern/data_structure/binaryTree"
)

// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/
// 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
// 同一个节点在一条路径序列中至多出现一次。该路径至少包含一个节点，且不一定经过根节点。
// 路径和是路径中各节点值的总和。
// 给你一个二叉树的根节点root，返回其最大路径和 。

func MaxPathSum(root *binaryTree.Node) int {
	_, b := helper(root)
	return b
}

func helper(root *binaryTree.Node) (int, int) {
	if root == nil {
		return 0, -(1 << 31)
	}
	lSingle, lMax := helper(root.Left)
	rSingle, rMax := helper(root.Right)
	var single int
	if lSingle > rSingle {
		single = max(lSingle+root.Val, 0)
	} else {
		single = max(rSingle+root.Val, 0)
	}

	maxPath := max(lMax, rMax)
	resMax := max(maxPath, lSingle+rSingle+root.Val)
	return single, resMax
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
