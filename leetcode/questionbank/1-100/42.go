package main

// 接雨水的问题
/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/
// 下一个较大的值 单调栈 只保留比当前高度高的值 递减
func trap(height []int) int {
	stk := []int{}
	ans := 0
	for i, h := range height {
		for len(stk) > 0 && h > height[stk[len(stk)-1]] {
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				break
			}
			left := stk[len(stk)-1]
			w := i - left - 1
			curh := _min(height[left], h) - height[top]
			ans += w * curh
		}
		stk = append(stk, i)
	}
	return ans
}

func _min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
