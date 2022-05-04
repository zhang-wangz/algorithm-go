package binaryTree

func LevelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	queue := make([]*Node, 0)

	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		list := make([]int, 0)
		list = append(list, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		res = append(res, list)
	}
	return res
}
