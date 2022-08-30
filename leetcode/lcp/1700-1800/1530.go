package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countPairs(root *TreeNode, distance int) (ans int) {
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) (res []int) {
		if node == nil {
			return []int{}
		}
		if node.Left == nil && node.Right == nil {
			return []int{0}
		}
		left := dfs(node.Left)
		for i := range left {
			left[i]++
			if left[i] <= distance {
				res = append(res, left[i])
			}
		}
		right := dfs(node.Right)
		for i := range right {
			if right[i] <= distance {
				res = append(res, right[i])
			}
		}
		for _, i := range left {
			for _, j := range right {
				if i+j <= distance {
					ans += 1
				}
			}
		}
		return
	}
	dfs(root)
	return
}

//func main() {
//	root := &TreeNode{Val: 1}
//	root.Left = &TreeNode{Val: 2}
//	root.Right = &TreeNode{Val: 3}
//	root.Left.Left = &TreeNode{Val: 4}
//	root.Left.Right = &TreeNode{Val: 5}
//	root.Right.Left = &TreeNode{Val: 6}
//	root.Right.Right = &TreeNode{Val: 7}
//	countPairs(root, 3)
//}
