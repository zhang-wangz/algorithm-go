package merge

import "algorithm-go/base/binaryTree"

// 左子树最大 < 根 < 右子树最小
type Res struct {
	IsValid  bool
	Min, Max *binaryTree.Node
}

func IsBinaryTree(root *binaryTree.Node) bool {
	res := helperIsBinaryTree(root)
	return res.IsValid
}

func helperIsBinaryTree(root *binaryTree.Node) Res {
	r := Res{}
	if root == nil {
		r.IsValid = true
		return r
	}
	left := helperIsBinaryTree(root.Left)
	right := helperIsBinaryTree(root.Right)
	if left.Max != nil && left.Max.Val >= root.Val {
		r.IsValid = false
		return r
	}
	if right.Min != nil && right.Min.Val <= root.Val {
		r.IsValid = false
		return r
	}
	r.IsValid = true
	r.Min = root
	if left.Min != nil {
		r.Min = left.Min
	}
	if right.Max != nil {
		r.Max = right.Max
	}
	return r
}
