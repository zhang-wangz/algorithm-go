package solution

import "algorithm-go/binaryTree"

func deleteNode(root *binaryTree.Node, key int) *binaryTree.Node {
	if root == nil {
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Val == key {
		// 用比当前节点大的最小的值去替换
		// 如果没有左孩子
		if root.Left == nil {
			return root.Right
			// 如果没有右孩子，存在左孩子
		} else if root.Right == nil {
			return root.Left
			// 存在左孩子，右孩子
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			// 把右孩子里面的最小的左节点取出来
			// 把root的左子树放到当前树的左边
			cur.Left = root.Left
			return root.Right
		}
	}
	return root
}
