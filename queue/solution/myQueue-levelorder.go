package solution

import "algorithm-go/binaryTree"

func levelOrder(root *binaryTree.Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*binaryTree.Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		lst := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			lst = append(lst, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, lst)
	}
	return res
}
