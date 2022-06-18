package main

import (
	"algorithm-go/leetcode/solution"
	"algorithm-go/linkedList"
)

func main() {
	//"fooleetbar"
	//"f00l"s
	//[["o","0"]]
	//words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	//pattern := "abb"
	// [62,100,4]
	//2
	nums := []int{1, 0, 0, 0, 0, 0, 0}
	a := &linkedList.ListNode{Val: 1}
	a.Next = &linkedList.ListNode{
		Val: 3,
		Next: &linkedList.ListNode{
			Val:  4,
			Next: a,
		}}
	solution.Insert(a, 2)
	println(nums)
}
