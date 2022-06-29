package offer

import "math"

func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		le := len(q)
		max := -1 << 31
		for i := 0; i < le; i++ {
			node := q[0]
			max = int(math.Max(float64(max), float64(node.Val)))
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, max)
	}
	return ans
}
