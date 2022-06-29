package offer

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := make([]*TreeNode, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		stack = append(stack, node)
		dfs(node.Right)
	}
	dfs(root)
	stack = bfs(root)
	head := stack[0]
	ans := head
	for i := 1; i < len(stack); i++ {
		ans.Right = stack[i]
		ans.Left = nil
		ans = ans.Right
	}
	ans.Left = nil
	return head
}

// 左中右
func bfs(node *TreeNode) []*TreeNode {
	stack := make([]*TreeNode, 0)
	ans := make([]*TreeNode, 0)
	if node == nil {
		return ans
	}
	for len(stack) != 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node1)
		node = node1.Right
	}
	return ans
}
