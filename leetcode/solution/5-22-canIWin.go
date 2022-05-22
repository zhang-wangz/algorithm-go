package solution

// https://leetcode.cn/problems/can-i-win/
// https://leetcode.cn/problems/can-i-win/solution/by-a-fei-8-jvlx/
// 在 "100 game" 这个游戏中，两名玩家轮流选择从 1 到 10 的任意整数，累计整数和，
// 先使得累计整数和 达到或超过  100 的玩家，即为胜者。
// 如果我们将游戏规则改为 “玩家 不能 重复使用整数” 呢？
// 例如，两个玩家可以轮流从公共整数池中抽取从 1 到 15 的整数（不放回），直到累计整数和 >= 100。
// 给定两个整数 maxChoosableInteger （整数池中可选择的最大数）和 desiredTotal（累计和），若先出手的玩家是否能稳赢则返回 true ，否则返回 false 。假设两位玩家游戏时都表现 最佳 。

// 回溯 + 状态压缩
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	n := maxChoosableInteger
	if n*(n+1)/2 < desiredTotal {
		return false
	}
	if desiredTotal <= maxChoosableInteger {
		return true
	}
	// 记录状态
	dp := map[int]bool{}
	var back func(des, state int) bool
	back = func(des, state int) bool {
		// 如果此时先手并且使用的数字已经有存储，直接返回
		if v, ok := dp[state]; ok {
			return v
		}
		for i := 1; i <= n; i++ {
			if (1<<i)&state == 0 {
				// 如果先手目标达成 或者 后手继续时候目标达不成，都返回true
				if des-i <= 0 || !back(des-i, state|1<<i) {
					dp[state] = true
					return true
				}
			}
		}
		dp[state] = false
		return dp[state]
	}
	return back(desiredTotal, 0)
}

// TLE //41/57 //case: //20 //210
func TLEcanIWin(maxChoosableInteger int, desiredTotal int) bool {
	n := maxChoosableInteger
	if n*(n+1) < desiredTotal {
		return false
	}
	if desiredTotal <= maxChoosableInteger {
		return true
	}
	var vis []bool
	for i := 1; i <= n; i++ {
		vis = make([]bool, 21)
		vis[i] = true
		if TLEback(maxChoosableInteger, desiredTotal, i, vis) {
			return true
		}
	}
	return false
}

func TLEback(n int, desTotal int, curTotal int, vis []bool) bool {
	if curTotal >= desTotal {
		return true
	}
	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		if TLEback(n, desTotal, curTotal+i, vis) {
			return false
		}
		vis[i] = false
	}
	return true
}
