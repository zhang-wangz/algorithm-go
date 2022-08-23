package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// m l r
//相邻的节点不是左节点，就是该节点往上某个节点的右节点
func recoverFromPreorder(traversal string) *TreeNode {
	pos, path := 0, make([]*TreeNode, 0)
	for pos < len(traversal) {
		level := 0
		for pos < len(traversal) && traversal[pos] == '-' {
			level++
			pos++
		}
		value := 0
		for pos < len(traversal) && traversal[pos] >= '0' && traversal[pos] <= '9' {
			value = value*10 + int(traversal[pos]-'0')
			pos++
		}
		node := &TreeNode{Val: value}
		if level == len(path) {
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			path = path[:level]
			path[len(path)-1].Right = node
		}
		path = append(path, node)
	}
	return path[0]
}
