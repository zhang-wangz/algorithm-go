package main

import "algorithm-go/offer/offer专项突击/61-78"

func main() {
	a := &_1_78.ListNode{Val: 4}
	a.Next = &_1_78.ListNode{Val: 2}
	a.Next.Next = &_1_78.ListNode{Val: 1}
	a.Next.Next.Next = &_1_78.ListNode{Val: 3}
	_1_78.SortList(a)
}
