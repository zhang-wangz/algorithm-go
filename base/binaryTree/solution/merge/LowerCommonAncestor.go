package merge

import "algorithm-go/base/binaryTree"

func LowerCommAncestor(root, p, q *binaryTree.Node) *binaryTree.Node {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := LowerCommAncestor(root.Left, p, q)
	right := LowerCommAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}

	if right != nil {
		return right
	}
	return nil
}
