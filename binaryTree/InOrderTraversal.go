package binaryTree

func InOrder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}

// Morris 空间复杂下降到1
func Morris(root *Node) {

}
