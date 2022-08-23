package main

// 每段中的最小值和最大值，最大值要小于下个组的最小值
// 求分割最多的次数
// 贪心
// 下一个大值
// 遇到大值直接加入
// 遇到小的就弹出
func maxChunksToSorted(arr []int) int {
	st := []int{}
	for _, v := range arr {
		if len(st) == 0 || st[len(st)-1] <= v {
			st = append(st, v)
		} else if len(st) > 0 && st[len(st)-1] > v {
			cur := st[len(st)-1]
			st = st[:len(st)-1]
			for len(st) > 0 && st[len(st)-1] > v {
				st = st[:len(st)-1]
			}
			st = append(st, cur)
		}
	}
	return len(st)
}
