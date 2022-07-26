package _022_07

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	root      *TreeNode
	childNode []*TreeNode
}

func Constructor6(root *TreeNode) CBTInserter {
	r := root
	q := make([]*TreeNode, 0)
	cadi := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left == nil || node.Right == nil {
			cadi = append(cadi, node)
		}
	}
	return CBTInserter{root: r, childNode: cadi}
}

func (it *CBTInserter) Insert(val int) int {
	cadi := it.childNode
	new := &TreeNode{Val: val}
	res := 0
	for len(cadi) != 0 {
		p := cadi[0]
		if p.Left == nil {
			p.Left = new
			res = p.Val
			break
		} else if p.Right == nil {
			p.Right = new
			res = p.Val
			break
		} else {
			cadi = cadi[1:]
			continue
		}
	}
	cadi = append(cadi, new)
	it.childNode = cadi
	return res
}

func (it *CBTInserter) Get_root() *TreeNode {
	return it.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
