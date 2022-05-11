package solution

import (
	"algorithm-pattern/data_structure/binaryTree"
	"strconv"
	"strings"
)

//  5-11
// https://leetcode.cn/problems/serialize-and-deserialize-bst/

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (_ *Codec) Serialize(root *binaryTree.Node) string {
	if root == nil {
		return "#"
	}
	res := []string{}
	q := make([]*binaryTree.Node, 0)
	q = append(q, root)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			res = append(res, "#")
		} else {
			res = append(res, strconv.Itoa(node.Val))
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (_ *Codec) Deserialize(data string) *binaryTree.Node {
	if data == "#" {
		return nil
	}
	values := strings.Split(data, ",")
	q := make([]*binaryTree.Node, 0)
	v := values[0]
	n, _ := strconv.Atoi(v)
	q = append(q, &binaryTree.Node{Val: n})
	res := q[0]
	idx := 0
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if values[idx*2+1] != "#" {
			n, _ := strconv.Atoi(values[idx*2+1])
			node.Left = &binaryTree.Node{
				Val: n,
			}
			q = append(q, node.Left)
		}
		if values[idx*2+2] != "#" {
			n, _ := strconv.Atoi(values[idx*2+2])
			node.Right = &binaryTree.Node{
				Val: n,
			}
			q = append(q, node.Right)
		}
		idx++
	}
	return res
}
