package main

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start int, end int) (ans []*TreeNode) {
	if start > end {
		return []*TreeNode{nil}
	}
	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)
		for _, l := range left {
			for _, r := range right {
				n := &TreeNode{i, nil, nil}
				n.Left = l
				n.Right = r
				ans = append(ans, n)
			}
		}
	}
	return
}
