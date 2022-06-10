package binary_search

import "algorithm-go/binaryTree"

// 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。 返回插入后二叉搜索树的根节点。
// https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/

func InsertTree(root *binaryTree.Node, val int) *binaryTree.Node {
	if root == nil {
		node := &binaryTree.Node{
			Val: val,
		}
		root = node
		return root
	}
	if root.Val > val {
		root.Left = InsertTree(root.Left, val)
	} else {
		root.Right = InsertTree(root.Right, val)
	}
	return root
}
