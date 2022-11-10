package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 意思就是要不该节点为空，要不该节点就包括2个子节点
// 递归构造
var m map[int][]*TreeNode

func init() {
	m = map[int][]*TreeNode{}
}
func allPossibleFBT(n int) (ans []*TreeNode) {
	if _, ok := m[n]; !ok {
		ans = []*TreeNode{}
		if n == 1 {
			ans = append(ans, &TreeNode{Val: 0})
		} else if n%2 == 1 {
			for l := 0; l < n; l++ {
				r := n - 1 - l
				for _, l := range allPossibleFBT(l) {
					for _, r := range allPossibleFBT(r) {
						ans = append(ans, &TreeNode{0, l, r})
					}
				}
			}
		}
		m[n] = ans
	}
	return m[n]
}
