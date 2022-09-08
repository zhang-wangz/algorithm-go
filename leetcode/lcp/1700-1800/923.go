package main

import "sort"

// 923 三数之和的多种可能
func threeSumMulti(arr []int, target int) (ans int) {
	const mod = 1e9 + 7
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		t := target - arr[i]
		j := i + 1
		k := len(arr) - 1
		for j < k {
			if arr[j]+arr[k] < t {
				j++
			} else if arr[j]+arr[k] < t {
				k--
			} else if arr[j] != arr[k] {
				l, r := 1, 1
				for j+1 < k && arr[j] == arr[j+1] {
					l++
					j++
				}
				for k-1 > j && arr[k] == arr[k-1] {
					r++
					k--
				}
				ans += l * r
				ans %= mod
				j++
				k--
			} else {
				// 如果相等，剩下的里面挑2个出来，相当于Cn 2
				ans += (k - j + 1) * (k - j) / 2
				ans %= mod
				break
			}
		}
	}
	return
}
