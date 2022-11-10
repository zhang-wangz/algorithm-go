package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中左右  左中右
func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func(pl, pr, ol, or int) *TreeNode
	dfs = func(pl, pr, ol, or int) *TreeNode {
		if pl > pr || ol > or {
			return nil
		}
		root := &TreeNode{Val: preorder[pl]}
		if pl == pr {
			return root
		}
		idx := ol
		for ; idx < len(inorder) && preorder[pl] != inorder[idx]; idx++ {
		}
		root.Left = dfs(pl+1, pl+idx-ol, ol, idx-1)
		root.Right = dfs(pl+idx-ol+1, pr, idx+1, or)
		return root
	}
	return dfs(0, len(preorder)-1, 0, len(inorder)-1)
}
