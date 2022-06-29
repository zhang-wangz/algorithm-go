package main

import "algorithm-go/offer"

func main() {
	root := &offer.TreeNode{Val: 2}
	root.Left = &offer.TreeNode{Val: -1}
	//root.Right = &offer.TreeNode{Val: 3}
	//root.Left.Right = &offer.TreeNode{Val: 5}
	//root.Right.Right = &offer.TreeNode{Val: 4}
	//offer.RightSideView(root)
	offer.MaxPathSum(root)
}
