package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type CBTInserter struct {
	root  *TreeNode
	queue []*TreeNode
}

func Constructor6(root *TreeNode) CBTInserter {
	return CBTInserter{
		root:  root,
		queue: []*TreeNode{},
	}
}

func (this *CBTInserter) Insert(v int) int {
	if len(this.queue) == 0 {
		this.queue = append(this.queue, this.root)
	}
	for len(this.queue) != 0 {
		node := this.queue[0]
		this.queue = this.queue[1:]
		if node.Left == nil {
			node.Left = &TreeNode{Val: v}
			this.queue = append([]*TreeNode{node}, this.queue...)
		} else if node.Right == nil {
			node.Right = &TreeNode{Val: v}
			this.queue = append([]*TreeNode{node}, this.queue...)
		} else {
			this.queue = append(this.queue, node.Left)
			this.queue = append(this.queue, node.Right)
			continue
		}
		return node.Val
	}
	return -1
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
