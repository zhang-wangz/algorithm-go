package dfs_79_87

// 解析
// https://leetcode.cn/problems/IDBivT/solution/jian-dan-yi-dong-javac-pythonjs-pei-yang-mmhs/
// 抽象成树之后，左节点是(,右节点是)
func generateParenthesis(n int) (ans []string) {
	var dfs func(int, int, string)
	dfs = func(left, right int, path string) {
		if left > n || right > left {
			return
		}
		if len(path) == n*2 {
			ans = append(ans, path)
			return
		}
		dfs(left+1, right, path+"(")
		dfs(left, right+1, path+")")
	}
	dfs(0, 0, "")
	return
}
