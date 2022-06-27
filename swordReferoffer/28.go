package swordReferoffer

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	dfs(root)
	return root
}

func dfs(root *Node) *Node {
	cur := root
	var last *Node
	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			childLast := dfs(cur.Child)
			// 连接cur和child
			cur.Next = cur.Child
			cur.Child.Prev = cur
			// 连接child和next
			if next != nil {
				childLast.Next = next
				next.Prev = childLast
			}
			cur.Child = nil
			last = childLast
		} else {
			last = cur
		}
		cur = next
	}
	return last
}
