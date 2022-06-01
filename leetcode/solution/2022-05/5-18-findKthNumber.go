package _022_05

// https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/solution/dong-tu-yan-shi-by-xiaohu9527-3k7s/
// 几乎每一个人都用 乘法表。但是你能在乘法表中快速找到第k小的数字吗？
// 给定高度m 、宽度n 的一张 m * n的乘法表，以及正整数k，你需要返回表中第k 小的数字。

func findKthNumber(m int, n int, k int) int {
	l := 1
	r := m * n
	// mid不一定在乘法表里面
	for l+1 < r {
		mid := l + (r-l)/2
		// count数 如果比k小，说明需要移到右边，扩大范围
		if count(m, n, mid) < k {
			l = mid
			// 如果大于，需要到左边
			// 如果等于，也继续缩小范围，因为mid不一定是在乘法表里面
			// 往左缩小
		} else {
			r = mid
		}
	}
	// 获取第一个大于等于k的值，不是l就是r
	// 先判断l，因为l相对r小，这题需要获w取左范围
	if count(m, n, l) >= k {
		return l
	}

	if count(m, n, r) >= k {
		return r
	}
	return 0
}

func count(m, n, mid int) int {
	cnt := 0
	i, j := m, 1
	for i >= 1 && j <= n {
		// 如果小移动到下一列，否则移动到上一行
		if i*j <= mid {
			cnt += i
			j++
		} else {
			i--
		}
	}
	return cnt
}
