package _022_06

import "algorithm-go/binaryTree"

// https://leetcode.cn/problems/most-frequent-subtree-sum/

func FindFrequentTreeSum(root *binaryTree.Node) []int {
	max := 1
	cnt := map[int]int{}
	res := make([]int, 0)
	dfs(root, cnt, &max, &res)
	return res
}

func dfs(root *binaryTree.Node, cnt map[int]int, m *int, res *[]int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, cnt, m, res)
	right := dfs(root.Right, cnt, m, res)
	sum := left + right + root.Val
	cnt[sum]++
	if *m == cnt[sum] {
		*res = append(*res, sum)
	} else if cnt[sum] > *m {
		*res = []int{sum}
		*m = cnt[sum]
	}
	return sum
}
