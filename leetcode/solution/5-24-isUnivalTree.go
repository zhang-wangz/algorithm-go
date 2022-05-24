package solution

import "algorithm-pattern/data_structure/binaryTree"

// https://leetcode.cn/problems/univalued-binary-tree/

func isUnivalTree(root *binaryTree.Node) bool {
	if root == nil {
		return true
	}
	val := root.Val
	return helperisUnivalTree(root, val)
}

func helperisUnivalTree(root *binaryTree.Node, val interface{}) bool {
	if root == nil {
		return true
	} else if root.Val != val {
		return false
	}
	return helperisUnivalTree(root.Left, val) && helperisUnivalTree(root.Right, val)
}
