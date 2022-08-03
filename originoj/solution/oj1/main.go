package main

// 假设在游戏中，你当前有体力 T ；有M个活动，每个活动 i 有所需消耗的体力值 t，可完成次数 c，奖励经验 e
// 记为 play[i] = [t, c, e]。请问最多能获得多少经验。
//输入：
//T = 180
//play = [[30, 3, 25], [40, 2, 30], [20, 5, 20]]
//输出：160

func solve(play [][]int, t int) int {
	// 一共有j的体力，最多装多少活动获得的经验最大
	// 前i个游戏，j的体力能获得的经验最大是多少
	cnt := 0
	for _, p := range play {
		c := p[1]
		cnt += c
	}
	dp := make([][]int, cnt+1)
	for i := range dp {
		dp[i] = make([]int, t+1)
	}
	k := 1
	for i := 1; i <= len(play); i++ {
		t1, c, e := play[i-1][0], play[i-1][1], play[i-1][2]
		for c > 0 {
			for j := 1; j <= t; j++ {
				if j-t1 < 0 {
					// 可以完成多少次，j的体力完不成该次活动
					dp[k][j] = dp[k-1][j]
				} else {
					// 在玩这次和不玩这次取最大值
					dp[k][j] = max(dp[k-1][j], dp[k-1][j-t1]+e)
				}
			}
			c--
			k++
		}
	}
	return dp[cnt][t]
}

func main() {
	a := solve([][]int{{30, 3, 25}, {40, 2, 30}, {20, 5, 20}}, 180)
	println(a)
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
