package _1_60

// 反中序 写一个非递归吧
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	ans := root
	stack := []*TreeNode{}
	sum := 0
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += node.Val
		node.Val = sum
		root = node.Left
	}
	return ans
}
