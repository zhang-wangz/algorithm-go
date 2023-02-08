package binaryTree

// 深度搜索，从下到上，分治
// 左右中
func PostOrder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*Node, 0)
	var lastVisit *Node
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return res
}
