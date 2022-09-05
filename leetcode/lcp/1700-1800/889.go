package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序+后序 三种方式都可以看一下
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
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
		for ; idx < len(postorder) && preorder[pl+1] != postorder[idx]; idx++ {
		}
		leftlen := idx - pl + 1
		root.Left = dfs(pl+1, pl+leftlen, ol, idx)
		root.Right = dfs(pl+leftlen+1, pr, idx+1, or-1)
		return root
	}
	return dfs(0, len(preorder)-1, 0, len(postorder)-1)
}
