package solution

// 中+后  左中右 左右中
func buildTree2(inorder []int, postorder []int) *TreeNode {
	var dfs func(pl, pr, ol, or int) *TreeNode
	dfs = func(pl, pr, ol, or int) *TreeNode {
		if pl > pr || ol > or {
			return nil
		}
		root := &TreeNode{Val: postorder[or]}
		if ol == or {
			return root
		}
		idx := pl
		for ; idx < len(inorder) && postorder[or] != inorder[idx]; idx++ {
		}
		lc := idx - 1 - pl + 1
		root.Left = dfs(pl, idx-1, ol, ol+lc-1)
		root.Right = dfs(idx+1, pr, ol+lc, or-1)
		return root
	}
	return dfs(0, len(inorder)-1, 0, len(postorder)-1)
}
