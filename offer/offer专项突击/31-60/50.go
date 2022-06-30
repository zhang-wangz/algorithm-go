package _1_60

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func dfs(node *TreeNode, t int) (res int) {
	if node == nil {
		return
	}
	if node.Val == t {
		res++
	}
	res += dfs(node.Left, t-node.Val)
	res += dfs(node.Right, t-node.Val)
	return
}

// 前缀和
func pathSum2(root *TreeNode, t int) int {
	path := []int{0}
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum += node.Val
		path = append(path, sum)
		for i := 0; i < len(path)-1; i++ {
			if sum-path[i] == t {
				res++
			}
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		path = path[:len(path)-1]
	}
	dfs(root, 0)
	return res
}
