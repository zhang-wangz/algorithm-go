package _022_07

// https://leetcode.cn/problems/binary-tree-pruning/
// Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = pruneTree2(root.Left)
	root.Right = pruneTree2(root.Right)
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var dfs func(root *TreeNode) (*TreeNode, bool)
	dfs = func(root *TreeNode) (*TreeNode, bool) {
		if root == nil {
			return nil, true
		}
		if root.Val == 0 && root.Left == nil && root.Right == nil {
			return root, true
		}
		_, ok1 := dfs(root.Left)
		_, ok2 := dfs(root.Right)
		if ok1 {
			root.Left = nil
		}
		if ok2 {
			root.Right = nil
		}
		if root.Val == 0 && ok1 && ok2 {
			return root, true
		}
		return root, false
	}
	_, ok := dfs(root)
	if root.Val == 0 && ok {
		root = nil
	}
	return root
}
