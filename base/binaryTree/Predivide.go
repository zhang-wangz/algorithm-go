package binaryTree

// 深度搜索，从下到上
func PreDivide(root *Node) []int {
	res := make([]int, 0)
	res = divide(root)
	return res
}

func divide(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	left := divide(root.Left)
	right := divide(root.Right)

	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}
