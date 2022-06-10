package binaryTree

// print proOrder num
func PreOrder(root *Node) []int {
	if root == nil {
		return make([]int, 0)
	}
	res := make([]int, 0)
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return res
}
