package _022_07

import "math"

// 最小绝对差

func MinimumAbsDifference(arr []int) [][]int {
	quickSort(arr, 0, len(arr)-1)
	ans := make([][]int, 0)
	ml := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i], arr[i+1]) == ml {
			ans = append(ans, []int{arr[i], arr[i+1]})
		} else if abs(arr[i], arr[i+1]) < ml {
			ans = append([][]int{}, []int{arr[i], arr[i+1]})
			ml = abs(arr[i], arr[i+1])
		}
	}
	return ans
}
func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func quickSort(arr []int, start, end int) {
	if len(arr) == 0 {
		return
	}
	if start < end {
		p := findP(arr, start, end)
		quickSort(arr, start, p-1)
		quickSort(arr, p+1, end)
	}
}

func findP(arr []int, start int, end int) int {
	p := arr[end]
	i := start
	j := start
	for ; j < end; j++ {
		if arr[j] < p {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
