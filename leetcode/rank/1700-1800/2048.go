package main

// 2048
// 枚举
func nextBeautifulNumber(n int) int {
next:
	for n++; ; n++ {
		cnt := [10]int{}
		for x := n; x > 0; x /= 10 {
			cnt[x%10]++
		}
		for x := n; x > 0; x /= 10 {
			if cnt[x%10] != x%10 {
				continue next
			}
		}
		return n
	}
}
