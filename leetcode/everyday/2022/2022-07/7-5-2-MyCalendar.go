package _022_07

type MyNode struct {
	Left, Right *MyNode
	start, end  int
}
type MyCalendar2 struct {
	root *MyNode
}

func Constructo1r() MyCalendar2 {
	return MyCalendar2{&MyNode{}}
}

func (node *MyCalendar2) Book(start int, end int) bool {
	if node.root == nil {
		node.root = &MyNode{start: start, end: end}
		return true
	}
	cur := node.root
	for true {
		if end <= cur.start {
			if cur.Left == nil {
				cur.Left = &MyNode{start: start, end: end}
				return true
			}
			cur = cur.Left
		} else if start >= cur.end {
			if cur.Right == nil {
				cur.Right = &MyNode{start: start, end: end}
				return true
			}
			cur = cur.Right
		} else {
			return false
		}
	}
	return false
}
