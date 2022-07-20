package main

import (
	"fmt"
	"math"
)

func slove(n int, arr []int) {

	left, right := make([]int, n+5), make([]int, n+5)
	m1 := map[int]int{}
	l := 0
	for i := 1; i < n; i++ {
		l = l + m1[i]
		left[i] = left[i-1] + l + int(math.Max(float64(0), float64(arr[i-1]-(n-1))))
		if arr[i-1] <= n-1 {
			m1[n+i+1-arr[i-1]-1] = m1[n+i+1-arr[i-1]-1] + 1
		} else {
			l++
		}
	}
	m2 := map[int]int{}
	l = 0
	for i := n; i >= 2; i-- {
		l = l + m2[i]
		right[i] = right[i+1] + l + int(math.Max(float64(0), float64(arr[i-1]-(n-1))))
		if arr[i-1] <= n-1 {
			m2[arr[i-1]+i-n-1+1] = m2[arr[i-1]+i-n-1+1] + 1
		} else {
			l++
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Print(left[i-1] + right[i+1])
		if i < n {
			fmt.Print(" ")
		}
	}
}
func main() {
	m := 0
	fmt.Scanf("%d", &m)
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	slove(m, arr)
}
