package main

import (
	"algorithm-pattern/data_structure/binaryTree"
)

func main() {
	root := new(binaryTree.Node)
	root.Val = 10
	root.Left = &binaryTree.Node{
		Val:  5,
		Left: &binaryTree.Node{Val: 4},
	}
	root.Right = &binaryTree.Node{Val: 11}

	//res := binaryTree.PreOrder(root)
	//for _, val := range res {
	//	println(val)
	//}

	//res := binaryTree.InOrder(root)
	//for _, val := range res {
	//	println(val)
	//}
}
