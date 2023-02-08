package main

import "algorithm-go/base/binaryTree"

func main() {
	root := new(binaryTree.Node)
	root.Val = 10
	root.Left = &binaryTree.Node{
		Val:  5,
		Left: &binaryTree.Node{Val: 4},
	}
	root.Right = &binaryTree.Node{Val: 11}

	//res := binaryTree.InOrder(root)
	//for _, val := range res {
	//	println(val)
	//}

	//node := everyday.ReserveTree2(root)
	//res = binaryTree.InOrder(node)
	//for _, val := range res {
	//	println(val)
	//}

	//res := binary_search.IsBinaryTree(root)
	//for _, val := range res {
	//	for _, v := range val {
	//		println(v)
	//	}
	//}
	res := &_022_05.Codec{}
	prin := binaryTree.InOrder(root)
	for _, val := range prin {
		println(val)
	}

	s1 := res.Serialize2(root)
	println("序列化", s1)
	s := res.Deserialize2(s1)
	prin = binaryTree.InOrder(s)
	for _, val := range prin {
		println(val)
	}
	println(res)
}
