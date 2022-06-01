package _022_05

import "algorithm-pattern/data_structure/binaryTree"

// 设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
// 如果指定节点没有对应的“下一个”节点，则返回null。

func inorderSuccessor(root *binaryTree.Node, p *binaryTree.Node) *binaryTree.Node {
	if root == nil {
		return root
	}
	flag := false
	stack := make([]*binaryTree.Node, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == p {
			flag = true
		}
		if flag == true && node != p {
			return node
		}
		root = node.Right

	}
	return nil
}
