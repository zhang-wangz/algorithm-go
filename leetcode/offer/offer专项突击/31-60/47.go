package _1_60

// 给定一个二叉树 根节点root，树的每个节点的值要么是 0，要么是 1。请剪除该二叉树中所有节点的值为 0 的子树。
//
//节点 node 的子树为node 本身，以及所有 node的后代。

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left == true {
			root.Left = nil
		}
		if right == true {
			root.Right = nil
		}
		if root.Val == 0 && left && right {
			return true
		}
		return false
	}
	f := dfs(root)
	if f {
		return nil
	}
	return root
}
