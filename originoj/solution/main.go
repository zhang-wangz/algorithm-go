package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(root, node *TreeNode) int {
	h := 0
	var dfs func(root, node *TreeNode, h *int)
	dfs = func(root, node *TreeNode, h *int) {
		if root == nil {
			return
		}
		if root == node {
			return
		}
		*h += 1
		dfs(root.Left, node, h)
		dfs(root.Right, node, h)
	}
	dfs(root, node, &h)
	return h
}

func main() {
	a := &TreeNode{Val: 5}
	a.Left = &TreeNode{Val: 4}
	a.Right = &TreeNode{Val: 9}
	c := &TreeNode{Val: 10}
	a.Left.Left = c
	println(solve(a, c))
}
