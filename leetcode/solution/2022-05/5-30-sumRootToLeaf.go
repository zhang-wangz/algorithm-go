package _022_05

import "algorithm-pattern/data_structure/binaryTree"

// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
// 给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。
//
//例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
//对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
//
//返回这些数字之和。题目数据保证答案是一个 32 位 整数。
func SumRootToLeaf(root *binaryTree.Node) int {
	if root == nil {
		return 0
	}
	//val := 0
	return helperSum2(root)
}
func helperSum(root *binaryTree.Node, val int) int {
	if root == nil {
		return 0
	}
	val = val<<1 | root.Val
	if root.Left == nil && root.Right == nil {
		return val
	}
	return helperSum(root.Left, val) + helperSum(root.Right, val)
}

func helperSum2(root *binaryTree.Node) int {
	if root == nil {
		return 0
	}
	val := 0
	res := 0
	lastNode := &binaryTree.Node{}
	stack := make([]*binaryTree.Node, 0)
	for len(stack) != 0 || root != nil {
		for root != nil {
			val = (val << 1) | root.Val
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastNode {
			if node.Left == nil && node.Right == nil {
				res += val
			}
			val >>= 1
			stack = stack[:len(stack)-1]
			lastNode = node
		} else {
			root = node.Right
		}
	}
	return res
}
