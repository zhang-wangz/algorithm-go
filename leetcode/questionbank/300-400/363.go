package main

/**
给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
题目数据保证总会存在一个数值和不超过 k 的矩形区域。

**/
import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func maxSumSubmatrix(matrix [][]int, k int) int {
	ans := math.MinInt64
	n := len(matrix[0])
	for i := range matrix {
		sum := make([]int, n)
		for _, row := range matrix[i:] { // 枚举下边界
			for c, v := range row {
				sum[c] += v // 更新每列的元素和
			}
			t := redblacktree.NewWithIntComparator()
			t.Put(0, nil)
			s := 0
			for _, v := range sum {
				s += v
				if node, ok := t.Ceiling(s - k); ok {
					ans = max(ans, s-node.Key.(int))
				}
				t.Put(s, nil)
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
