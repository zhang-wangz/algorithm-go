package main

func main() {
	a := &ListNode{Val: 4}
	a.Next = &ListNode{Val: 2}
	a.Next.Next = &ListNode{Val: 1}
	a.Next.Next.Next = &ListNode{Val: 3}
	sortList(a)
}
