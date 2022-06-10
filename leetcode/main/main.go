package main

import dp2 "algorithm-go/leetcode/solution/dp/6-10-countPalindromicSubsequences"

func main() {
	//data := [][]int{{4, 2, 3}, {0, 0, 1}, {7, 6, 5}}
	//solution.CutOffTree(data)
	//s := "2001:0db8:85a3:00000:0:8A2E:0370:7334"
	//root := new(binaryTree.Node)
	//root.Val = 1
	//root.Left = &binaryTree.Node{
	//	Val:  0,
	//	Left: &binaryTree.Node{Val: 1},
	//}
	//root.Right = &binaryTree.Node{Val: 0}
	//res := solution.SumRootToLeaf(root)
	//fmt.Println(res)
	//s := []int{1, 1, 2, 2, 2}
	//s := []string{"fxggfzk.fo.q+e@uxbeyetxc.com", "fxggfzk.fo.q+h@uxbeyetxc.com", "uv+mw.lkw+ybe@yppz.com"}
	s := "bccb"
	res := dp2.CountPalindromicSubsequences(s)
	println(res)
}
