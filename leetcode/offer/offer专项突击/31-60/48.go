package _1_60

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor7() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (code *Codec) serialize(root *TreeNode) string {
	sb := strings.Builder{}
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteByte(',')
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (code *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		v := sp[0]
		if v == "null" {
			sp = sp[1:]
			return nil
		}
		n, _ := strconv.Atoi(v)
		sp = sp[1:]
		return &TreeNode{Val: n, Left: build(), Right: build()}
	}
	return build()
}
