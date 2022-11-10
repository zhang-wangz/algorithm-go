package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	q := []*TreeNode{root}
	inorderIdx := 0
	for i := 1; i < len(preorder); i++ {
		preVal := preorder[i]
		node := q[len(q)-1]
		if node.Val != inorder[inorderIdx] {
			node.Left = &TreeNode{Val: preVal}
			q = append(q, node.Left)
		} else {
			for len(q) != 0 && q[len(q)-1].Val == inorder[inorderIdx] {
				node = q[len(q)-1]
				q = q[:len(q)-1]
				inorderIdx++
			}
			node.Right = &TreeNode{Val: preVal}
			q = append(q, node.Right)
		}
	}
	return root
}

// 重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 中左右
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
