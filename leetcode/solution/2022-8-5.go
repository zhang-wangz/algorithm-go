package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	level := 1
	for len(q) != 0 {
		if level == depth-1 {
			for i := 0; i < len(q); i++ {
				node := q[i]
				node.Left = &TreeNode{val, node.Left, nil}
				node.Right = &TreeNode{val, nil, node.Right}
			}
			return root
		}
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++
	}
	return root
}
